package order

import (
	"nasdaqvfs/config"
	"nasdaqvfs/internal/order"
	"nasdaqvfs/pkg/logger"
)

// Order UseCase
type orderUC struct {
	cfg       *config.Config
	orderRepo order.Repository
	logger    logger.Logger
}

// NewUserUseCase constructor
func NewUserUseCase(cfg *config.Config, orderRepo order.Repository, logger logger.Logger) order.UseCase {
	return &orderUC{cfg: cfg, orderRepo: orderRepo, logger: logger}
}
