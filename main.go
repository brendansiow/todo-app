package main

import (
	"github.com/brendansiow/todo-app/apis"
	"github.com/brendansiow/todo-app/core"
	"github.com/gin-gonic/gin"
)

func init() {
	core.Initialize()
}

func main() {
	router := gin.Default()
	apis.BindTodoApi(router)
	router.Run()
}
