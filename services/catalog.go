package services

import (
	"github.com/Jamess-Lucass/ecommerce-catalog-service/middleware"
	"github.com/Jamess-Lucass/ecommerce-catalog-service/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CatalogService struct {
	db *gorm.DB
}

func NewCatalogService(db *gorm.DB) *CatalogService {
	return &CatalogService{
		db: db,
	}
}

func (s *CatalogService) List(user *middleware.Claim) *gorm.DB {
	if user == nil {
		return s.db.
			Model(models.Catalog{}).
			Preload("Images").
			Where("is_deleted <> ?", true)
	}

	return s.db.
		Model(models.Catalog{}).
		Preload("Images").
		Where("is_deleted <> ?", true).
		Select("EXISTS(select * from catalog_user_likes l where l.catalog_id = catalogs.id and l.user_id = ?) AS is_liked, catalogs.*", user.Subject)
}

func (s *CatalogService) Get(user *middleware.Claim, id uuid.UUID) (*models.Catalog, error) {
	var catalog models.Catalog
	if err := s.List(user).First(&catalog, id).Error; err != nil {
		return nil, err
	}

	return &catalog, nil
}

func (s *CatalogService) Create(c *models.Catalog) error {
	return s.db.Create(&c).Error
}
