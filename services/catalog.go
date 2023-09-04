package services

import (
	"context"

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

func (s *CatalogService) List(ctx context.Context, user *middleware.Claim) *gorm.DB {
	db := s.db.WithContext(ctx)

	if user == nil {
		return db.
			Model(models.Catalog{}).
			Preload("Images").
			Where("is_deleted <> ?", true)
	}

	return db.
		Model(models.Catalog{}).
		Preload("Images").
		Where("is_deleted <> ?", true).
		Select("EXISTS(select * from catalog_user_likes l where l.catalog_id = catalogs.id and l.user_id = ?) AS is_liked, catalogs.*", user.Subject)
}

func (s *CatalogService) Get(ctx context.Context, user *middleware.Claim, id uuid.UUID) (*models.Catalog, error) {
	var catalog models.Catalog
	if err := s.List(ctx, user).First(&catalog, id).Error; err != nil {
		return nil, err
	}

	return &catalog, nil
}

func (s *CatalogService) Create(ctx context.Context, c *models.Catalog) error {
	db := s.db.WithContext(ctx)

	return db.Create(&c).Error
}
