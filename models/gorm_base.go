package models

import (
	"time"

	"gorm.io/gorm"
)

type GormBase struct {
	ID        uint           `gorm:"primarykey" db:"id" json:"id"`
	CreatedAt time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt time.Time      `db:"updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" db:"deleted_at" json:"deleted_at"`
}
