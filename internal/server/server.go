package server

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"

	"nasdaqvfs/config"
	"nasdaqvfs/pkg/logger"
	"nasdaqvfs/pkg/metric"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Server struct
type Server struct {
	gin     *gin.Engine
	cfg     *config.Config
	db      *gorm.DB
	logger  logger.Logger
	metrics metric.Metrics
}

// NewServer New Server constructor
func NewServer(cfg *config.Config, db *gorm.DB, metrics metric.Metrics, opts ...Option) *Server {
	s := &Server{
		gin:     gin.New(),
		cfg:     cfg,
		db:      db,
		metrics: metrics,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) Run() error {
	if err := s.MapHandlers(); err != nil {
		return err
	}

	ctx := context.Background()
	go func() {
		s.logger.Infof(ctx, "Server is listening on PORT: %s", s.cfg.Server.Port)
		ln, _ := net.Listen("tcp", ":"+s.cfg.Server.Port)
		err := s.gin.RunListener(ln)
		if err != nil {
			s.logger.Fatalf(ctx, "Error starting Server: ", err)
		}
	}()

	go func() {
		s.metrics.RunServer(s.cfg.Metrics.URL)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	s.logger.Info(ctx, "Server Exited Properly")
	return nil
}
