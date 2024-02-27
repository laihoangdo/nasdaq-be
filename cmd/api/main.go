package main

import (
	"context"
	"log"

	"nasdaqvfs/config"
	"nasdaqvfs/internal/server"
	"nasdaqvfs/pkg/database/mysql"
	"nasdaqvfs/pkg/logger"
	"nasdaqvfs/pkg/metric"
)

func main() {
	log.Println("Starting api server")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	ctx := context.Background()
	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof(ctx, "AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)

	// Repository
	mysqlDB, err := mysql.New(config.DefaultDB, &cfg.MySQL)
	if err != nil {
		appLogger.Fatalf(ctx, "MySQL init: %s", err)
	}

	//init metrics
	metrics, err := metric.NewMetrics(cfg.Metrics.ServiceName)
	if err != nil || metrics == nil {
		appLogger.Fatalf(ctx, "CreateMetrics Error: %s", err)
	}

	s := server.NewServer(
		cfg,
		mysqlDB,
		metrics,
		server.Logger(appLogger),
	)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
