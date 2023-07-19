package models

type User struct {
	GormBase
	Email       string `json:"email"`
	LoginType   string `json:"login_type"`
	LoginTypeId int    `json:"login_type_id"`
}
