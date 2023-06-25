package models

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null;default:current_timestamp"`
	CreatedBy uuid.UUID `json:"createdBy" gorm:"type:uuid"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy uuid.UUID `json:"updatedBy" gorm:"type:uuid"`
	IsDeleted bool      `json:"-" gorm:"not null;default:false"`
}
