package user

import (
	"github.com/gin-gonic/gin"
)

// Handlers User handler interface
//
//go:generate mockery --name Handlers
type Handlers interface {
	Create(c *gin.Context)
	Login(c *gin.Context)
	// UpdateUserByID update user
	UpdateUserByID(c *gin.Context)
	// GetUserByID get user by id
	GetUserByID(c *gin.Context)
	// GetUsers get list user
	GetUsers(c *gin.Context)
}
