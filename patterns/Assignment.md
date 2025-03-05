# Assignment

## Objective
The goal of this assignment is to implement a real-world application using appropriate design patterns in Golang. You will demonstrate your ability to choose and apply the right design patterns based on specific requirements.

## Problem Statement
You are tasked with building a simple library management system. This system should allow users to manage books, including adding new books, searching for books, and tracking borrowed books. You will need to use various design patterns to structure your code effectively

## Requirements
### Implement the Following Features:
- **Add Books**: Users should be able to add new books to the library
- **Search Books**: Users should be able to search for books by title, author, or genre
- **Borrow Books**: Users should be able to borrow books, and the system should keep track of borrowed books

### Use Appropriate Design Patterns:
- **Factory Method:** Use this pattern to create book objects. Define a Book interface and implement different types of books (e.g., EBook, PrintedBook)
- **Singleton**:  Use the Singleton pattern to implement a centralized datastore that holds the collection of books
- **Observer**: Implement an observer pattern to notify users when a book they are interested in becomes available
- **Strategy**: Use this pattern to allow different search algorithms (e.g., search by title, author, or genre)
