package main

import (
	"log"
	"net/http"

	"github.com/Lenton-Losper/go-books/database"
	"github.com/Lenton-Losper/go-books/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../database")
	if err != nil {
		log.Println("No .env file found. Using default values.")
	}
}

func main() {
	defer database.Disconnect()

	router := gin.Default()

	// Create a wrapper function for CreateBook handler
	createBookHandler := func(c *gin.Context) {
		var book models.Book
		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id := database.CreateBook(book)
		c.JSON(http.StatusOK, gin.H{"id": id})
	}

	// Create a wrapper function for ListBooks handler
	listBooksHandler := func(c *gin.Context) {
		books := database.ListBooks()
		c.JSON(http.StatusOK, books)
	}

	// Create a wrapper function for FindBook handler
	findBookHandler := func(c *gin.Context) {
		name := c.Param("name")
		book := database.FindBook(name)
		if book == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusOK, book)
	}

	router.POST("/books", createBookHandler)
	router.GET("/books", listBooksHandler)
	router.GET("/books/:name", findBookHandler)

	router.Run("localhost:8080")
}
