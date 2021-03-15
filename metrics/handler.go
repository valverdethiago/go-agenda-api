package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	objectives   = map[float64]float64{0: 0.001, 0.5: 0.05, 0.9: 0.01, 0.99: 0.0001}
	responseTime = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "response_time",
		Help:       "response time of requests received by API endpoint",
		Objectives: objectives,
	}, []string{"status_code", "handler"})
	methodTime = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "method_time",
		Help:       "response time of requests received by the API endpoint",
		Objectives: objectives,
	}, []string{"method"})
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (controller *Controller) SetupRoutes(router *gin.Engine) {
	router.GET("/metrics", controller.prometheusHandler())
}

func (controller *Controller) prometheusHandler() gin.HandlerFunc {
	handler := promhttp.Handler()
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
