package handlers

import (
	"github.com/gin-gonic/gin"
	"library/database"
	"library/models"
	"net/http"
)

type BookHandler struct {
	DbFunc database.BookRepository
}

func (b *BookHandler) GetBooks(c *gin.Context) {
	books, err := b.DbFunc.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get books"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
}
func (b *BookHandler) AddBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := b.DbFunc.AddBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": book})
}
