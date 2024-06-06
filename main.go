package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Autor string  `json:"autor"`
	Price float64 `json:"price"`
}

var books = []book{
	{ID: "1", Title: "1984", Autor: "George Orwell", Price: 56.99},
	{ID: "2", Title: "Harry Potter and the Philosopher's Stone", Autor: "J.K. Rowling", Price: 17.99},
	{ID: "3", Title: "To Kill a Mockingbird", Autor: "Harper Lee", Price: 39.99},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", postBooks)
	router.Run("localhost:8080")
}

func postBooks(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
