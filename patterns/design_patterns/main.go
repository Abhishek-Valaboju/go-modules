package main

import (
	"fmt"
	"sync"
)

type Book interface {
	GetTitle() string
	GetAuthor() string
	GetGenre() string
}
type EBook struct {
	Title  string
	Author string
	Genre  string
}

func (e *EBook) GetTitle() string  { return e.Title }
func (e *EBook) GetAuthor() string { return e.Author }
func (e *EBook) GetGenre() string  { return e.Genre }

type PrintedBook struct {
	Title  string
	Author string
	Genre  string
}

func (p *PrintedBook) GetTitle() string  { return p.Title }
func (p *PrintedBook) GetAuthor() string { return p.Author }
func (p *PrintedBook) GetGenre() string  { return p.Genre }

// Factory Method for Book Creation
func CreateBook(bookType, title, author, genre string) Book {
	if bookType == "ebook" {
		return &EBook{Title: title, Author: author, Genre: genre}
	} else if bookType == "printed" {
		return &PrintedBook{Title: title, Author: author, Genre: genre}
	}
	return nil
}

// ====================== Singleton (Library Datastore) ======================
type Library struct {
	books     []Book
	observers []Observer
	mu        sync.Mutex
}

func GetLibraryInstance() *Library {
	return &Library{}
}
func (l *Library) AddBook(book Book) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.books = append(l.books, book)
	fmt.Println("Added book:", book.GetTitle())
	l.NotifyObservers(book.GetTitle())
}
func (l *Library) GetBooks() []Book {
	return l.books
}

// ====================== Observer (User Notification) ======================
// Observer Interface
type Observer interface {
	Notify(bookTitle string)
}

// User Struct (Observer)
type User struct {
	Name string
}

func (u *User) Notify(bookTitle string) {
	fmt.Println("Notification for", u.Name, ": The book", bookTitle, "is now available.")
}

// Methods to Register and Notify Observers
func (l *Library) RegisterObserver(observer Observer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.observers = append(l.observers, observer)
}
func (l *Library) NotifyObservers(bookTitle string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, observer := range l.observers {
		observer.Notify(bookTitle)
	}
}

// ====================== Strategy (Search Books by Different Criteria) ======================
// Search Strategy Interface
type SearchStrategy interface {
	Search(books []Book, query string) []Book
}

// Search by Title
type SearchByTitle struct{}

func (s *SearchByTitle) Search(books []Book, query string) []Book {
	var results []Book
	for _, book := range books {
		if book.GetTitle() == query {
			results = append(results, book)
		}
	}
	return results
}

// Search by Author
type SearchByAuthor struct{}

func (s *SearchByAuthor) Search(books []Book, query string) []Book {
	var results []Book
	for _, book := range books {
		if book.GetAuthor() == query {
			results = append(results, book)
		}
	}
	return results
}

// Search by Genre
type SearchByGenre struct{}

func (s *SearchByGenre) Search(books []Book, query string) []Book {
	var results []Book
	for _, book := range books {
		if book.GetGenre() == query {
			results = append(results, book)
		}
	}
	return results
}

// Search Context
type SearchContext struct {
	strategy SearchStrategy
}

func (c *SearchContext) SetStrategy(strategy SearchStrategy) {
	c.strategy = strategy
}
func (c *SearchContext) SearchBooks(books []Book, query string) []Book {
	return c.strategy.Search(books, query)
}

func main() {
	// Get Singleton Library Instance
	library := GetLibraryInstance()
	// Users Subscribing for Notifications
	user1 := &User{Name: "Alice"}
	user2 := &User{Name: "Bob"}
	library.RegisterObserver(user1)
	library.RegisterObserver(user2)
	// Adding Books using Factory Method
	library.AddBook(CreateBook("ebook", "Golang Basics", "Alice", "Tech"))
	library.AddBook(CreateBook("printed", "Python for Beginners", "Bob", "Tech"))
	library.AddBook(CreateBook("ebook", "Advanced Rust", "Alice", "Programming"))
	// Searching for Books using Strategy Pattern
	searchContext := &SearchContext{}
	// Search by Author
	searchContext.SetStrategy(&SearchByAuthor{})
	results := searchContext.SearchBooks(library.GetBooks(), "Alice")
	fmt.Println("\nBooks by Alice:")
	for _, book := range results {
		fmt.Println(book.GetTitle(), "-", book.GetGenre())
	}
	// Search by Genre
	searchContext.SetStrategy(&SearchByGenre{})
	results = searchContext.SearchBooks(library.GetBooks(), "Tech")
	fmt.Println("\nBooks in Tech Genre:")
	for _, book := range results {
		fmt.Println(book.GetTitle(), "-", book.GetAuthor())
	}
}
