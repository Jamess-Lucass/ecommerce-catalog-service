package main

import (
	"net/http"
	"os"

	"github.com/Jamess-Lucass/ecommerce-catalog-service/database"
	"github.com/Jamess-Lucass/ecommerce-catalog-service/handlers"
	"github.com/Jamess-Lucass/ecommerce-catalog-service/services"
	"go.elastic.co/apm/module/apmhttp/v2"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var LOG_LEVEL = os.Getenv("LOG_LEVEL")
var LOG_LEVELS = map[string]zapcore.Level{
	"DEBUG": zap.DebugLevel,
	"INFO":  zap.InfoLevel,
	"WARN":  zap.WarnLevel,
	"ERROR": zap.ErrorLevel,
}

func main() {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, LOG_LEVELS[LOG_LEVEL])
	logger := zap.New(core, zap.AddCaller())

	db := database.Connect(logger)

	if err := database.Migrate(db); err != nil {
		logger.Sugar().Errorf("error occured migrating database: %v", err)
	}

	http.DefaultTransport = apmhttp.WrapRoundTripper(http.DefaultTransport, apmhttp.WithClientTrace())

	healthService := services.NewHealthService(db)
	catalogService := services.NewCatalogService(db)
	catalogUserLikesService := services.NewCatalogUserLikesService(db)

	server := handlers.NewServer(logger, healthService, catalogService, catalogUserLikesService)

	if err := server.Start(); err != nil {
		logger.Sugar().Fatalf("error starting web server: %v", err)
	}
}
