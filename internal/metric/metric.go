package metric

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/satimoto/go-datastore/pkg/util"
)

type Metric interface {
	Handler() *chi.Mux
	StartMetric(context.Context, *sync.WaitGroup)
}

type MetricService struct {
	Server *http.Server
}

func NewMetric() Metric {
	return &MetricService{}
}

func (rs *MetricService) Handler() *chi.Mux {
	router := chi.NewRouter()
	router.Mount("/metrics", promhttp.Handler())

	return router
}

func (rs *MetricService) StartMetric(ctx context.Context, waitGroup *sync.WaitGroup) {
	if rs.Server == nil {
		rs.Server = &http.Server{
			Addr:    fmt.Sprintf(":%s", os.Getenv("METRIC_PORT")),
			Handler: rs.Handler(),
		}
	}

	log.Printf("Starting Metric service")
	waitGroup.Add(1)

	go rs.listenAndServe()

	go func() {
		<-ctx.Done()
		log.Printf("Shutting down Metric service")

		rs.shutdown()

		log.Printf("Metric service shut down")
		waitGroup.Done()
	}()
}

func (rs *MetricService) listenAndServe() {
	err := rs.Server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		util.LogOnError("API044", "Error in Metric service", err)
	}
}

func (rs *MetricService) shutdown() {
	timeout := util.GetEnvInt32("SHUTDOWN_TIMEOUT", 20)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	err := rs.Server.Shutdown(ctx)

	if err != nil {
		util.LogOnError("API045", "Error shutting down Metric service", err)
	}
}
