package main

import (
	"github.com/brendansiow/todo-app/core"
	"github.com/brendansiow/todo-app/models"
)

func init() {
	core.Migration = true
	core.Initialize()
}

func main() {
	core.DB.AutoMigrate(&models.Todo{})
}
