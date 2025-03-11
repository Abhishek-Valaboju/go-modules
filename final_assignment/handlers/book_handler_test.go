package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"library/database"
	"library/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setupRouter() *gin.Engine {
	b := BookHandler{}
	r := gin.Default()
	r.GET("/books", b.GetBooks)
	r.POST("/books", b.AddBook)
	return r
}
func TestAddBook(t *testing.T) {
	database.DbConn()
	database.DB.AutoMigrate(&models.Book{})
	router := setupRouter()
	bookJson := `{"title":"Go","auther":"Test"}`
	req, _ := http.NewRequest("POST", "/books", strings.NewReader(bookJson))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetBooks(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/books", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
