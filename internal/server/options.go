package server

import (
	"nasdaqvfs/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Option -.
type Option func(*Server)

func FiberEngine(gin *gin.Engine) Option {
	return func(s *Server) {
		s.gin = gin
	}
}

func Logger(logger logger.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}
