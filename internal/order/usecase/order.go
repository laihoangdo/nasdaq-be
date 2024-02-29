package usecase

import (
	"context"
	"nasdaqvfs/internal/models"
	"nasdaqvfs/pkg/utils"
)

func (o orderUC) Create(ctx context.Context, order models.Order) error {
	return nil
}

func (o orderUC) GetByID(ctx context.Context, orderID int64) (models.Order, error) {
	return models.Order{}, nil

}
func (o orderUC) UpdateByID(ctx context.Context, orderID int64, order models.Order) error {
	return nil

}
func (o orderUC) GetOrders(ctx context.Context, pq *utils.PaginationQuery) ([]models.Order, error) {
	return nil, nil

}
