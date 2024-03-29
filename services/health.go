package services

import (
	"context"

	"gorm.io/gorm"
)

type HealthService struct {
	db *gorm.DB
}

func NewHealthService(db *gorm.DB) *HealthService {
	return &HealthService{
		db: db,
	}
}

func (s *HealthService) Ping(ctx context.Context) error {
	db := s.db.WithContext(ctx)

	sql, err := db.DB()
	if err != nil {
		return err
	}

	if err := sql.Ping(); err != nil {
		return err
	}

	return nil
}
