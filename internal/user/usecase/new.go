package usecase

import (
	"nasdaqvfs/config"
	"nasdaqvfs/internal/user"
	"nasdaqvfs/pkg/logger"
)

// User UseCase
type userUC struct {
	cfg      *config.Config
	userRepo user.Repository
	logger   logger.Logger
}

// NewUserUseCase constructor
func NewUserUseCase(cfg *config.Config, userRepo user.Repository, logger logger.Logger) user.UseCase {
	return &userUC{cfg: cfg, userRepo: userRepo, logger: logger}

}
