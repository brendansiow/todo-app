package apis

import (
	"net/http"

	"github.com/brendansiow/todo-app/core"
	"github.com/brendansiow/todo-app/models"
	"github.com/gin-gonic/gin"
)

func BindTodoApi(router *gin.Engine) {
	router.POST("/todo/create", create)
	router.GET("/todo/list", list)
}

func create(c *gin.Context) {
	var requestBody struct {
		Title       string
		Description string
	}
	c.Bind(&requestBody)

	todo := models.Todo{
		Title:       requestBody.Title,
		Description: requestBody.Description,
	}
	result := core.DB.Create(&todo)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func list(c *gin.Context) {
	todos := []models.Todo{}

	result := core.DB.Find(&todos)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todos})
}
