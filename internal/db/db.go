package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"taskServer/internal/models"
)

var db *gorm.DB // глобальная переменная для доступа к БД

func initDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=taskdb port=5435 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	if err := db.AutoMigrate(&models.Task{}); err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

}
