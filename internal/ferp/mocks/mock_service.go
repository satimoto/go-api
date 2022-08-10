package mocks

import (
	"context"
	"errors"
	"sync"

	"github.com/satimoto/go-ferp/pkg/rate"
)

type MockFerpService struct{
	getRateMockData []*rate.CurrencyRate
}

func NewService() *MockFerpService {
	return &MockFerpService{}
}

func (s *MockFerpService) Start(shutdownCtx context.Context, waitGroup *sync.WaitGroup) {}

func (s *MockFerpService) GetRate(currency string) (*rate.CurrencyRate, error) {
	if len(s.getRateMockData) == 0 {
		return &rate.CurrencyRate{}, errors.New("NotFound")
	}

	response := s.getRateMockData[0]
	s.getRateMockData = s.getRateMockData[1:]
	return response, nil
}

func (s *MockFerpService) SetGetRateMockData(currencyRate *rate.CurrencyRate) {
	s.getRateMockData = append(s.getRateMockData, currencyRate)
}