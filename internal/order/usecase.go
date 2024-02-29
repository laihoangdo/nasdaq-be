package order

import (
	"context"

	"nasdaqvfs/internal/models"
	"nasdaqvfs/pkg/utils"
)

// UseCase user use case interface
//
//go:generate mockery --name UseCase
type UseCase interface {
	Create(ctx context.Context, order models.Order) error
	GetByID(ctx context.Context, orderID int64) (models.Order, error)
	// UpdateByID update order by id
	UpdateByID(ctx context.Context, orderID int64, order models.Order) error
	// GetOrders list users
	GetOrders(ctx context.Context, pq *utils.PaginationQuery) ([]models.Order, error)
}
