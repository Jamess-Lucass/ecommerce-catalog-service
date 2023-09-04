package database

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Jamess-Lucass/ecommerce-catalog-service/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
)

func Connect(log *zap.Logger) *gorm.DB {
	server := os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic("Could not parse PORT to an integar")
	}
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%d", server, user, pass, name, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(log.Level())),
	})

	if err != nil {
		log.Sugar().Fatalf("failed to connect to the database. %v", err)
	}

	return db
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Catalog{},
		&models.CatalogImage{},
		&models.CatalogUserLike{},
	)
}
