package middleware

import (
	"time"

	"nasdaqvfs/pkg/metric"

	"github.com/gin-gonic/gin"
)

// Prometheus metrics middleware
func (mw *MiddlewareManager) MetricsMiddleware(metrics metric.Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		// before req
		t := time.Now()

		c.Next()

		// after req
		endpoint := c.FullPath()

		if endpoint == "" {
			endpoint = c.Request.URL.Path
		}

		if !metrics.IsSkipPath(endpoint) {
			statusCode := c.Writer.Status()
			method := c.Request.Method
			latency := time.Since(t)

			metrics.IncHits(statusCode, method, endpoint)
			metrics.ObserveResponseTime(statusCode, method, endpoint, float64(latency.Seconds()))
		}
	}
}
