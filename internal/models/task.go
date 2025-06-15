package models

import "gorm.io/gorm"

type Task struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	Status    string         `json:"status"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
