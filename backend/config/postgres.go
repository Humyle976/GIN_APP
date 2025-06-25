package config

import (
	"gin_app/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres() {
	dsn := os.Getenv("POSTGRES_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatal("Couldn't connect to the database: ", err)
	} else {
		log.Println("Connected To The Database")
	}
	DB = db
}

func Migrate() {
	
	DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Likes{})
}
