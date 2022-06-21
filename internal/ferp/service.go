package ferp

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ferp/ferprpc"
	"github.com/satimoto/go-ferp/pkg/ferp"
	"github.com/satimoto/go-ferp/pkg/rate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Ferp interface {
	Start(shutdownCtx context.Context, waitGroup *sync.WaitGroup)
	GetRate(currency string) (*rate.CurrencyRate, error)
}

type FerpService struct {
	FerpRpc       ferp.Ferp
	RatesClient   ferprpc.RateService_SubscribeRatesClient
	currencyRates rate.LatestCurrencyRates
}

func NewService(address string) Ferp {
	return &FerpService{
		FerpRpc:       ferp.NewService(address),
		currencyRates: make(rate.LatestCurrencyRates),
	}
}

func (s *FerpService) Start(shutdownCtx context.Context, waitGroup *sync.WaitGroup) {
	log.Printf("Starting up FERP client")
	ratesChan := make(chan ferprpc.SubscribeRatesResponse)

	go s.waitForRates(shutdownCtx, waitGroup, ratesChan)
	go s.subscribeRates(shutdownCtx, ratesChan)
}

func (s *FerpService) GetRate(currency string) (*rate.CurrencyRate, error) {
	if currencyRate, ok := s.currencyRates[currency]; ok {
		return &currencyRate, nil
	}

	return nil, errors.New("no currency rate available")
}

func (s *FerpService) handleRate(currencyRate ferprpc.SubscribeRatesResponse) {
	/** Rate received.
	 *
	 */
	s.currencyRates[currencyRate.Currency] = rate.CurrencyRate{
		Rate:        currencyRate.Rate,
		RateMsat:    currencyRate.RateMsat,
		LastUpdated: time.Unix(currencyRate.LastUpdated, 0),
	}
}

func (s *FerpService) subscribeRates(shutdownCtx context.Context, ratesChan chan<- ferprpc.SubscribeRatesResponse) {
	ratesClient, err := s.waitForSubscribeRatesClient(shutdownCtx, 0, 1000)
	util.PanicOnError("API001", "Error creating FERP client", err)
	s.RatesClient = ratesClient

	for {
		subscribeRatesResponse, err := s.RatesClient.Recv()

		if err == nil {
			ratesChan <- *subscribeRatesResponse
		} else {
			if errors.Is(shutdownCtx.Err(), context.Canceled) {
				log.Printf("Context is cancelled")
				return
			}

			s.RatesClient, err = s.waitForSubscribeRatesClient(shutdownCtx, 100, 1000)
			util.PanicOnError("API002", "Error creating FERP client", err)
		}
	}
}

func (s *FerpService) waitForRates(shutdownCtx context.Context, waitGroup *sync.WaitGroup, ratesChan chan ferprpc.SubscribeRatesResponse) {
	waitGroup.Add(1)

waitLoop:
	for {
		select {
		case <-shutdownCtx.Done():
			log.Printf("Shutting down FERP client")
			break waitLoop
		case subscribeRatesResponse := <-ratesChan:
			s.handleRate(subscribeRatesResponse)
		}
	}

	log.Printf("FERP client shut down")
	close(ratesChan)
	waitGroup.Done()
}

func (s *FerpService) waitForSubscribeRatesClient(shutdownCtx context.Context, initialDelay, retryDelay time.Duration) (ferprpc.RateService_SubscribeRatesClient, error) {
	for {
		if initialDelay > 0 {
			time.Sleep(retryDelay * time.Millisecond)
		}

		subscribeRatesClient, err := s.FerpRpc.SubscribeRates(shutdownCtx, &ferprpc.SubscribeRatesRequest{})

		if err == nil {
			return subscribeRatesClient, nil
		} else if status.Code(err) != codes.Unavailable {
			return nil, err
		}

		log.Print("Waiting for FERP client")
		time.Sleep(retryDelay * time.Millisecond)
	}
}
