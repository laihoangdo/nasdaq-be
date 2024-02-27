package http

import (
	"nasdaqvfs/config"
	"nasdaqvfs/internal/user"
	"nasdaqvfs/pkg/logger"
)

// userHandlers handlers
type userHandlers struct {
	cfg    *config.Config
	userUC user.UseCase
	logger logger.Logger
}

// NewUserHandlers user handlers constructor
func NewUserHandlers(cfg *config.Config, userUC user.UseCase, logger logger.Logger) user.Handlers {
	return &userHandlers{cfg: cfg, userUC: userUC, logger: logger}
}
