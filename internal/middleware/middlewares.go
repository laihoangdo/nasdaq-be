package middleware

import (
	"nasdaqvfs/config"
	"nasdaqvfs/pkg/logger"
	"nasdaqvfs/pkg/token"
)

// Middleware manager
type MiddlewareManager struct {
	cfg        *config.Config
	origins    []string
	logger     logger.Logger
	tokenMaker token.Maker
}

// Middleware manager constructor
func NewMiddlewareManager(cfg *config.Config, origins []string, logger logger.Logger, tokenMaker token.Maker) *MiddlewareManager {
	return &MiddlewareManager{
		cfg:        cfg,
		origins:    origins,
		logger:     logger,
		tokenMaker: tokenMaker,
	}
}
