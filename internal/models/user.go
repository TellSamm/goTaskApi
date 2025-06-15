package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
