package database

import (
	"Assignment-2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbHost     = "localhost"
	dbPort     = 5434
	dbUser     = "postgres"
	dbPassword = "gendut167"
	dbName     = "Assignment-2"
	db         *gorm.DB
)

func StartDB() {
	Config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	var err error

	// Initialize the database using GORM.
	db, err = gorm.Open(postgres.Open(Config), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.Debug().AutoMigrate(&models.Orders{}, &models.Items{})

	fmt.Println("Connected to the database ")
}

func GetDB() *gorm.DB {
	return db
}
