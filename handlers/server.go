package handlers

import (
	"github.com/Jamess-Lucass/ecommerce-catalog-service/services"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Server struct {
	validator               *validator.Validate
	logger                  *zap.Logger
	healthService           *services.HealthService
	catalogService          *services.CatalogService
	catalogUserLikesService *services.CatalogUserLikesService
}

func NewServer(
	logger *zap.Logger,
	healthService *services.HealthService,
	catalogService *services.CatalogService,
	catalogUserLikesService *services.CatalogUserLikesService,
) *Server {
	return &Server{
		validator:               validator.New(),
		logger:                  logger,
		healthService:           healthService,
		catalogService:          catalogService,
		catalogUserLikesService: catalogUserLikesService,
	}
}
