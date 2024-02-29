package http

import (
	"nasdaqvfs/internal/middleware"
	"nasdaqvfs/internal/order"

	"github.com/gin-gonic/gin"
)

// MapOrderRoutes Map order routes
func MapOrderRoutes(userGroup *gin.RouterGroup, h order.Handlers, mw *middleware.MiddlewareManager) {
	//userGroup.Use(mw.AuthAdminMiddleware())
	userGroup.POST("", h.Create)
	userGroup.PUT("/:id", h.UpdateByID)
	userGroup.GET("/:id", h.GetByID)
	userGroup.GET("", h.GetOrders)
}
