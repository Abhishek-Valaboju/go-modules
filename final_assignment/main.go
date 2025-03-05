package main

import (
	"github.com/gin-gonic/gin"
	"library/database"
	"library/handlers"
)

func main() {

	bookRepo := database.DbConn()
	bookHandler := handlers.BookHandler{DbFunc: bookRepo}
	r := gin.Default()
	r.GET("/books", bookHandler.GetBooks)
	r.POST("/books", bookHandler.AddBook)
	r.Run(":8080")
}
