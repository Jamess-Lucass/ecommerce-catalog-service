package main

import (
	"github.com/Jamess-Lucass/ecommerce-catalog-service/database"
	"github.com/Jamess-Lucass/ecommerce-catalog-service/handlers"
	"github.com/Jamess-Lucass/ecommerce-catalog-service/services"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Sugar().Warnf("could not flush: %v", err)
		}
	}()

	db := database.Connect(logger)

	if err := database.Migrate(db); err != nil {
		logger.Sugar().Errorf("error occured migrating database: %v", err)
	}

	catalogService := services.NewCatalogService(db)
	server := handlers.NewServer(logger, catalogService)

	if err := server.Start(); err != nil {
		logger.Sugar().Fatalf("error starting web server: %v", err)
	}
}
