package main

import (
	"errors"
	"fmt"

	"github.com/Jamess-Lucass/ecommerce-catalog-service/database"
	"github.com/Jamess-Lucass/ecommerce-catalog-service/models"
	"github.com/brianvoe/gofakeit/v6"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	logger, _ := zap.NewProduction()
	db := database.Connect(logger)

	if err := database.Migrate(db); err != nil {
		logger.Sugar().Fatalf("error occured migrating database: %v", err)
	}

	gofakeit.Seed(8675309)

	if err := db.First(&models.Catalog{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		items := []models.Catalog{}

		for i := 0; i < 100_000; i++ {
			item := models.Catalog{
				Name:        gofakeit.LoremIpsumWord(),
				Description: gofakeit.LoremIpsumSentence(10),
				Price:       float32(gofakeit.Price(5, 100)),
			}

			for j := 0; j < 3; j++ {
				image := models.CatalogImage{
					URL: fmt.Sprintf("https://picsum.photos/seed/%s/500/500", gofakeit.UUID()),
				}

				item.Images = append(item.Images, image)
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
