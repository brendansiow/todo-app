package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string
	Description string
	Completed   bool `gorm:"default:false"`
}
