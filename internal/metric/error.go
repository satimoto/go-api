package metrics

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	metricErrorsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "ocpi_errors_total",
		Help: "The total number of errors",
	}, []string{"code"})
)

func RecordError(code, message string, err error) {
	if err != nil {
		metricErrorsTotal.WithLabelValues(code).Inc()
		
		log.Printf("%s: %s", code, message)
		log.Printf("%s: %v", code, err)
	}
}
