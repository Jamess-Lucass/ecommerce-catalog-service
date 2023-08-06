package models

import (
	"time"

	"github.com/google/uuid"
)

type CatalogUserLike struct {
	CatalogID uuid.UUID `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time `gorm:"not null;default:current_timestamp"`
}
