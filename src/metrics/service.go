package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

// MetricService service to collect metrics
type MetricService struct {
	path     string
	router   *gin.Engine
	duration []float64
}

// NewMetricService creates a new instance of MetricService
func NewMetricService(path string, router *gin.Engine) *MetricService {
	return &MetricService{
		path:     path,
		router:   router,
		duration: []float64{0.1, 0.3, 1.2, 5, 10},
	}
}

// Configure the Metric Service and start collecting metrics
func (service *MetricService) Configure() {
	service.configureMetrics()
	service.init()
}

func (service *MetricService) init() {
	monitor := ginmetrics.GetMonitor()
	monitor.SetMetricPath(service.path)
	monitor.SetDuration(service.duration)
	monitor.Use(service.router)
}

func (service *MetricService) configureMetrics() {

}
