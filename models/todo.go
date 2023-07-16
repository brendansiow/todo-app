package models

type Todo struct {
	GormBase
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Completed   bool   `gorm:"default:false" db:"completed" json:"completed"`
}
