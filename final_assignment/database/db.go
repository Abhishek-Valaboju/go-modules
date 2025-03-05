package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"library/models"
)

var DB *gorm.DB

type BookRepository interface {
	GetAllBooks() ([]*models.Book, error)
	AddBook(book *models.Book) error
}
type BookRepo struct {
	db *gorm.DB
}

func DbConn() *BookRepo {
	var err error
	dbStr := ""
	DB, err = gorm.Open(postgres.Open(dbStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&models.Book{})

	return &BookRepo{db: DB}
}
func (b *BookRepo) GetAllBooks() ([]*models.Book, error) {
	var books []*models.Book
	err := b.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}
func (b *BookRepo) AddBook(book *models.Book) error {
	return b.db.Create(book).Error
}
