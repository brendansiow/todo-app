package models

type TodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
