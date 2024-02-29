package order

import (
	"context"

	"nasdaqvfs/internal/models"
	"nasdaqvfs/pkg/utils"
)

// Repository user interface
//
//go:generate mockery --name Repository
type Repository interface {
	Create(ctx context.Context, order models.Order) error
	// GetByID get user by id
	GetByID(ctx context.Context, id int64) (models.Order, error)
	// UpdateUserByID update user by ID
	UpdateByID(ctx context.Context, id int64, order models.Order) error
	// GetOrders list order
	GetOrders(ctx context.Context, pq *utils.PaginationQuery) ([]models.Order, error)
}
