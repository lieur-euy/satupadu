package database

import (
	"log"
	"os"

	"web/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error .env gaes")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("tidak konek ke database ")
	} else {
		log.Println("connect database sukses")
	}
	DB = database
	DB.AutoMigrate(
		&models.Users{}, &models.Product{},
	)
}
