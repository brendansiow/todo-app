package main

import (
	"github.com/brendansiow/todo-app/apis"
	"github.com/brendansiow/todo-app/core"
	"github.com/brendansiow/todo-app/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	core.Initialize()
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//Public API Endpoints
	public := router.Group("/")
	apis.BindLoginApi(public)

	//Protected API Endpoints
	protected := router.Group("/")
	protected.Use(middlewares.JwtAuthMiddleware)
	apis.BindTodoApi(protected)

	router.Run()
}
