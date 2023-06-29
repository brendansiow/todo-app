package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	router := gin.Default() //new gin router initialization
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "Hello World!"})
	}) // first endpoint returns Hello World
	router.Run(":8080") //running application, Default port is 8080
}
