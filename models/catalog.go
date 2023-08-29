package models

type Catalog struct {
	Base

	Name        string  `json:"name" gorm:"not null;type:varchar(128)"`
	Description string  `json:"description" gorm:"not null;type:varchar(1024)"`
	Price       float32 `json:"price" gorm:"not null;type:decimal(10,2)"`

	Images []CatalogImage    `json:"images"`
	Likes  []CatalogUserLike `json:"-"`

	IsLiked bool `json:"isLiked" gorm:"->"`
}
