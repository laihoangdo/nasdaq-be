package user

import (
	"context"

	"nasdaqvfs/internal/models"
	"nasdaqvfs/pkg/utils"
)

// Repository user interface
//
//go:generate mockery --name Repository
type Repository interface {
	Create(ctx context.Context, user models.User) error
	GetByUsername(ctx context.Context, username string) (record models.User, err error)
	// GetUserByID get user by id
	GetUserByID(ctx context.Context, id int64) (models.User, error)
	// UpdateUserByID update user by ID
	UpdateUserByID(ctx context.Context, id int64, user models.User) error
	// GetUsers list user
	GetUsers(ctx context.Context, pq *utils.PaginationQuery) ([]models.User, error)
}
