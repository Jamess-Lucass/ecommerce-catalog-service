package services

import (
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

func (s *CatalogUserLikesService) Create(c *models.CatalogUserLike) error {
	return s.db.Create(&c).Error
}

func (s *CatalogUserLikesService) Delete(c *models.CatalogUserLike) error {
	return s.db.Delete(&c).Error
}
