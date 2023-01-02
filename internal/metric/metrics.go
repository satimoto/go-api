package metrics

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

type Metrics interface {
	Handler() *chi.Mux
	StartMetrics(context.Context, *sync.WaitGroup)
}

type MetricsService struct {
	Server *http.Server
}

func NewMetrics() Metrics {
	return &MetricsService{}
}

func (rs *MetricsService) Handler() *chi.Mux {
	router := chi.NewRouter()
	router.Mount("/metrics", promhttp.Handler())

	return router
}

func (rs *MetricsService) StartMetrics(ctx context.Context, waitGroup *sync.WaitGroup) {
	if rs.Server == nil {
		rs.Server = &http.Server{
			Addr:    fmt.Sprintf(":%s", os.Getenv("METRIC_PORT")),
			Handler: rs.Handler(),
		}
	}

	log.Printf("Starting Metrics service")
	waitGroup.Add(1)

	go rs.listenAndServe()

	go func() {
		<-ctx.Done()
		log.Printf("Shutting down Metrics service")

		rs.shutdown()

		log.Printf("Metrics service shut down")
		waitGroup.Done()
	}()
}

func (rs *MetricsService) listenAndServe() {
	err := rs.Server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		RecordError("API044", "Error in Metrics service", err)
	}
}

func (rs *MetricsService) shutdown() {
	timeout := util.GetEnvInt32("SHUTDOWN_TIMEOUT", 20)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	err := rs.Server.Shutdown(ctx)

	if err != nil {
		RecordError("API045", "Error shutting down Metrics service", err)
	}
}
