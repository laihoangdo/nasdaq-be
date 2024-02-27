package middleware

import (
	pkgLogger "nasdaqvfs/pkg/logger"
	"nasdaqvfs/pkg/utils"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware set the logger with some fields inside the logger.
func (mw *MiddlewareManager) LoggerMiddleware(l pkgLogger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = l.WithFields(ctx, pkgLogger.Fields{
			"METHOD":     c.Request.Method,
			"PATH":       c.Request.URL.Path,
			"REQUEST_ID": utils.GetRequestID(c),
		})
		c.Request.WithContext(ctx)
		c.Next()
	}
}
