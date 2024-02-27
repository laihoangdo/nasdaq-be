package middleware

import (
	pkgErrors "nasdaqvfs/pkg/errors"
	"nasdaqvfs/pkg/utils"

	"github.com/gin-gonic/gin"
)

const (
	internalAPIKeyHeader    = "x-internal-api-key"
	internalAPIOriginHeader = "x-internal-api-origin"
)

var (
	errEmptyOrigin = pkgErrors.NewError(401, "origin is empty", nil)
	errEmptyKey    = pkgErrors.NewError(401, "key is empty", nil)
	errInvalidKey  = pkgErrors.NewError(401, "key is invalid", nil)
)

// ValidateInternalService validates internal service
func (mw MiddlewareManager) ValidateInternalService() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader(internalAPIOriginHeader)
		if origin == "" {
			mw.logger.Error(c, "middleware.ValidateInternalService: origin is empty")
			c.JSON(401, errEmptyOrigin)
			c.Abort()
			return
		}

		key := c.GetHeader(internalAPIKeyHeader)
		if key == "" {
			mw.logger.Error(c, "middleware.ValidateInternalService: key is empty")
			c.JSON(401, errEmptyKey)
			c.Abort()
			return
		}

		if !utils.Contains(mw.cfg.InternalAPI.AcceptedKeys, key) {
			mw.logger.Error(c, "middleware.ValidateInternalService: key is invalid")
			c.JSON(401, errInvalidKey)
			c.Abort()
			return
		}

		c.Next()
	}
}
