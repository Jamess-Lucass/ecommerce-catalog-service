package main

import (
	"errors"

	"github.com/Jamess-Lucass/ecommerce-catalog-service/database"
	"github.com/Jamess-Lucass/ecommerce-catalog-service/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	logger, _ := zap.NewProduction()
	db := database.Connect(logger)

	if err := db.First(&models.Catalog{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Sugar().Infof("seeding data...")

		db.Create(&models.Catalog{Name: "T-Shirt", Description: "A very cool new T-Shirt to purchase.", Price: 20.50})

		db.Create(&models.Catalog{Name: "Jeans", Description: "Some brand spanking new jeans!", Price: 32.99})
	}
}
