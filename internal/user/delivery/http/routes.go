package http

import (
	"nasdaqvfs/internal/middleware"
	"nasdaqvfs/internal/user"

	"github.com/gin-gonic/gin"
)

// MapUserRoutes Map menu routes
func MapUserRoutes(userGroup *gin.RouterGroup, h user.Handlers, mw *middleware.MiddlewareManager) {
	//userGroup.Use(mw.AuthAdminMiddleware())
	userGroup.POST("/register", h.Create)
	userGroup.POST("/login", h.Login)
	userGroup.PUT("/:id", h.UpdateUserByID)
	userGroup.GET("/:id", h.GetUserByID)
	userGroup.GET("", h.GetUsers)
}
