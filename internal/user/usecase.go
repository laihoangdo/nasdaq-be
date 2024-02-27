package user

import (
	"context"

	"nasdaqvfs/internal/models"
	"nasdaqvfs/pkg/utils"
)

// UseCase user use case interface
//
//go:generate mockery --name UseCase
type UseCase interface {
	Create(ctx context.Context, user models.User) error
	// Login  admin user to the system and return credentials
	Login(ctx context.Context, username string, password string) (models.User, error)
	// GetUserByID get user by id
	GetUserByID(ctx context.Context, userID int64) (models.User, error)
	// UpdateUserByID update user by id
	UpdateUserByID(ctx context.Context, userID int64, user models.User) error
	// GetUsers list users
	GetUsers(ctx context.Context, pq *utils.PaginationQuery) ([]models.User, error)

	//GetByUsername(ctx context.Context, username string) (models.User, error)
}
