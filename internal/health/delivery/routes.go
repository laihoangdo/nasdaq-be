package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Map routes
func MapRoutes(gr *gin.RouterGroup) {
	gr.GET("/liveness", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": true})
	})
}
