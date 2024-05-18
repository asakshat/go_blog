package db

import (
	"log"
	"os"

	"github.com/asakshat/go_blog/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dbURL), config)
	if err != nil {
		return nil, err
	}

	DB = db
	DB.Config.DisableForeignKeyConstraintWhenMigrating = true

	if err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Like{}); err != nil {
		log.Println("Error migrating models:", err)
		return nil, err
	}

	// Auto migrate gorm models

	return db, nil

}
