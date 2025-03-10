package db

import (
	"log"

	"github.com/rs-karasal/bitlab_go/lesson1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable search_path=mood"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&models.Mood{})

	log.Print("Successfully connected to database!")
	DB = database
}
