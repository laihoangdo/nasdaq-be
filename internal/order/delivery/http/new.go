package http

import (
	"nasdaqvfs/config"
	"nasdaqvfs/internal/order"
	"nasdaqvfs/pkg/logger"
)

// orderHandlers handlerss
type orderHandlers struct {
	cfg     *config.Config
	orderUC order.UseCase
	logger  logger.Logger
}

// NewOrderHandlers order handlers constructor
func NewOrderHandlers(cfg *config.Config, orderUC order.UseCase, logger logger.Logger) order.Handlers {
	return &orderHandlers{cfg: cfg, orderUC: orderUC, logger: logger}
}
