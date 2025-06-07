package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"taskServer/internal/models"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=taskdb port=5435 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	if err := db.AutoMigrate(&models.Task{}); err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}
	return db
}
