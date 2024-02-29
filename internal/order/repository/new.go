package repository

import (
	"nasdaqvfs/internal/order"

	"gorm.io/gorm"
)

// order Repository
type orderRepo struct {
	db *gorm.DB
}

// NewOrderRepository constructor
func NewOrderRepository(db *gorm.DB) order.Repository {
	return &orderRepo{db: db}
}
