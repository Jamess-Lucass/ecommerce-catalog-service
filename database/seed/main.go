package main

import (
	"errors"

	"github.com/Jamess-Lucass/ecommerce-catalog-service/database"
	"github.com/Jamess-Lucass/ecommerce-catalog-service/models"
	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	logger, _ := zap.NewProduction()
	db := database.Connect(logger)

	gofakeit.Seed(8675309)

	if err := db.First(&models.Catalog{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		items := []models.Catalog{}

		for i := 1; i < 100_000; i++ {
			item := models.Catalog{
				Name:        gofakeit.LoremIpsumWord(),
				Description: gofakeit.LoremIpsumSentence(10),
				Price:       float32(gofakeit.Price(5, 100)),
			}

			items = append(items, item)
		}

		if err := db.CreateInBatches(items, 1_000).Error; err != nil {
			logger.Sugar().Errorf("error occured while seeding data: %v", err)
			return
		}

		logger.Sugar().Infof("Seeded data...")
	}
}
