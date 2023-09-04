package services

import (
	"context"

	"github.com/Jamess-Lucass/ecommerce-catalog-service/models"
	"gorm.io/gorm"
)

type CatalogUserLikesService struct {
	db *gorm.DB
}

func NewCatalogUserLikesService(db *gorm.DB) *CatalogUserLikesService {
	return &CatalogUserLikesService{
		db: db,
	}
}

func (s *CatalogUserLikesService) Create(ctx context.Context, c *models.CatalogUserLike) error {
	db := s.db.WithContext(ctx)

	return db.Create(&c).Error
}

func (s *CatalogUserLikesService) Delete(ctx context.Context, c *models.CatalogUserLike) error {
	db := s.db.WithContext(ctx)

	return db.Delete(&c).Error
}
