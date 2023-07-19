package apis

import (
	"net/http"

	"github.com/brendansiow/todo-app/core"
	"github.com/brendansiow/todo-app/helper"
	"github.com/brendansiow/todo-app/models"
	"github.com/gin-gonic/gin"
)

func BindTodoApi(router *gin.RouterGroup) {
	router.POST("/todo/create", create)
	router.GET("/todo/list", list)
	router.PUT("/todo/complete", complete)
	router.DELETE("/todo/delete", delete)
}

func create(c *gin.Context) {
	todoRequest := models.TodoRequest{}
	if err := c.ShouldBindJSON(&todoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Arguments"})
		return
	}
	if todoRequest.Title == "" || todoRequest.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Arguments"})
		return
	}

	todo := getTodoModelWithUserId(c)
	todo.Title = todoRequest.Title
	todo.Description = todoRequest.Description

	result := core.DB.Create(&todo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func list(c *gin.Context) {
	todos := []models.Todo{}
	findModel := getTodoModelWithUserId(c)

	result := core.DB.Where(&findModel).Find(&todos)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func complete(c *gin.Context) {
	todoId := c.Query("todoId")

	if todoId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Todo ID"})
		return
	}

	var todo models.Todo
	findModel := getTodoModelWithUserId(c)

	firstResult := core.DB.Where(&findModel).First(&todo, todoId)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "The Todo does not exist."})
		return
	}

	if firstResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": firstResult.Error.Error()})
		return
	}

	todo.Completed = !todo.Completed
	saveResult := core.DB.Save(&todo)

	if saveResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": saveResult.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todo})

}

func delete(c *gin.Context) {
	todoId := c.Query("todoId")

	if todoId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Todo ID"})
		return
	}

	findModel := getTodoModelWithUserId(c)

	result := core.DB.Where(&findModel).Delete(&models.Todo{}, todoId)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "The Todo does not exist."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todoId})
}

func getTodoModelWithUserId(c *gin.Context) models.Todo {
	userId := helper.GetID(c)
	return models.Todo{UserID: int(userId)}
}
