package repository

import (
	"context"
	"nasdaqvfs/internal/models"
	"nasdaqvfs/pkg/utils"
)

func (r orderRepo) Create(ctx context.Context, order models.Order) error {
	return nil
}

func (r orderRepo) GetByID(ctx context.Context, id int64) (models.Order, error) {
	return models.Order{}, nil
}

func (r orderRepo) UpdateByID(ctx context.Context, id int64, order models.Order) error {
	return nil
}

func (r orderRepo) GetOrders(ctx context.Context, pq *utils.PaginationQuery) ([]models.Order, error) {
	return nil, nil
}
