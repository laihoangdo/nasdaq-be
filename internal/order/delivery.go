package order

import "github.com/gin-gonic/gin"

// Handlers User handler interface
//
//go:generate mockery --name Handlers
type Handlers interface {
	Create(c *gin.Context)
	// UpdateByID update user
	UpdateByID(c *gin.Context)
	// GetByID get order by id
	GetByID(c *gin.Context)
	// GetOrders get list user
	GetOrders(c *gin.Context)
}
