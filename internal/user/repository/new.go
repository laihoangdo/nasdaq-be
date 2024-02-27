package repository

import (
	"nasdaqvfs/internal/user"

	"gorm.io/gorm"
)

// user Repository
type userRepo struct {
	db *gorm.DB
}

// NewUserRepository constructor
func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepo{db: db}
}
