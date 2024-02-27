package server

import (
	"fmt"

	"nasdaqvfs/pkg/token"

	healthH "nasdaqvfs/internal/health/delivery"
	apiMiddlewares "nasdaqvfs/internal/middleware"

	userHttp "nasdaqvfs/internal/user/delivery/http"
	userRepository "nasdaqvfs/internal/user/repository"
	userUseCase "nasdaqvfs/internal/user/usecase"

	"github.com/gin-contrib/requestid"
)

// Map Server Handlers
func (s *Server) MapHandlers() error {
	// setup token
	tokenMaker, err := token.NewJWTMaker(s.cfg.Server.JwtSecretKey)
	if err != nil {
		return fmt.Errorf("cannot create token maker: %w", err)
	}

	// Init Kafka

	// Init other services
	//imageService := uploader.NewImageService(s.cfg.ImgSvc)
	//seSvc := media.NewScheduleExporterService(imageService)

	// Init repositories
	userRepo := userRepository.NewUserRepository(s.db)

	// Init useCases

	userProfileUC := userUseCase.NewUserUseCase(s.cfg, userRepo, s.logger)

	//Init province usecase

	//Passcode

	// Init handlers
	userHandlers := userHttp.NewUserHandlers(s.cfg, userProfileUC, s.logger)

	// init middleware

	mw := apiMiddlewares.NewMiddlewareManager(s.cfg, []string{"*"}, s.logger, tokenMaker)

	s.gin.Use(requestid.New())
	s.gin.Use(mw.MetricsMiddleware(s.metrics))
	s.gin.Use(mw.LoggerMiddleware(s.logger))
	s.gin.Use(mw.Recover(s.logger))

	v1 := s.gin.Group("/api/v1")
	healthH.MapRoutes(v1)

	userGroup := v1.Group("/user")

	userHttp.MapUserRoutes(userGroup, userHandlers, mw)

	//internal := s.gin.Group("/internal")
	//
	//internal.Use(mw.ValidateInternalService())

	return nil
}
