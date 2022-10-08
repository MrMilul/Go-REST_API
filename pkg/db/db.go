package db

import (
	"example/restapi/pkg/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	URL := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(URL), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Book{}, &models.User{})
	return db
}
