package db

import (
	"log"

	"github.com/stckrz/go-stckrz-site/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error){
	database, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db", err)
	}
	database.AutoMigrate(&models.Song{})
	database.AutoMigrate(&models.Blogpost{})
	return database, nil
}
