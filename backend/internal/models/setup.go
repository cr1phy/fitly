package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_URL")), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database.")
	}
	if err := db.AutoMigrate(&Product{}, &Dish{}); err != nil {
		log.Fatalln("something went wrong with migration:", err)
	}
	DB = db
}
