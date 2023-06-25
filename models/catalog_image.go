package models

import "github.com/google/uuid"

type CatalogImage struct {
	Base

	URL string `json:"url" gorm:"not null;type:varchar(1024)"`

	CatalogID uuid.UUID `json:"-"`
}
