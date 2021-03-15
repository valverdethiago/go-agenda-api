package metrics

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
func Metrics(service UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		appMetric := NewHTTP(r.URL.Path, r.Method)
		appMetric.Started()
		next(w, r)
		res := w.(http.ResponseWriter)
		appMetric.Finished()
		appMetric.StatusCode = strconv.Itoa(res.Status())
		mService.SaveHTTP(appMetric)
	}
}*/

// HandlerFunc defines handler function for middleware
func Metrics(service UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		appMetric := NewHTTP(c.Request.URL.Path, c.Request.Method)
		appMetric.Started()
		c.Next()
		status := c.Writer.Status()
		appMetric.Finished()
		appMetric.StatusCode = strconv.Itoa(status)
		service.SaveHTTP(appMetric)
	}
}
