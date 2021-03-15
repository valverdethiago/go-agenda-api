package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Service Prometheus Service
type Service struct {
	prometheusUrl        string
	httpRequestHistogram *prometheus.HistogramVec
}

//NewPrometheusService create a new prometheus service
func NewPrometheusService(prometheusUrl string) (*Service, error) {
	http := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Name:      "request_duration_seconds",
		Help:      "The latency of the HTTP requests.",
		Buckets:   prometheus.DefBuckets,
	}, []string{"handler", "method", "code"})

	s := &Service{
		prometheusUrl:        prometheusUrl,
		httpRequestHistogram: http,
	}
	err := prometheus.Register(s.httpRequestHistogram)
	if err != nil && err.Error() != "duplicate metrics collector registration attempted" {
		return nil, err
	}
	return s, nil
}

//SaveHTTP send metrics to server
func (s *Service) SaveHTTP(h *HTTP) {
	s.httpRequestHistogram.WithLabelValues(h.Handler, h.Method, h.StatusCode).Observe(h.Duration)
}
