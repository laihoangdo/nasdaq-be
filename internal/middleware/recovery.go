package middleware

import (
	pkgLogger "nasdaqvfs/pkg/logger"

	"github.com/gin-gonic/gin"
)

func (mw *MiddlewareManager) Recover(l pkgLogger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx := c.Request.Context()
				l.Error(ctx, err)
				c.Next()
			}
		}()
	}
}
