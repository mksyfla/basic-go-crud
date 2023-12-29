package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	v1 := router.Group("v1")

	v1.GET("/", rootHandler)
	v1.GET("/hello", helloHandler)
	v1.GET("/books/", booksHandler)
	v1.GET("/books/:id", bookHandler)

	v1.POST("/books", postBooksHandler)

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Kasyfil",
		"bio":  "Golang Gin Routing Test",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Hello World",
		"bio":  "Golang Gin Routing Test (Root /hello)",
	})
}

func bookHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func booksHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.QueryArray("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

type bookInput struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}

func postBooksHandler(c *gin.Context) {
	var book bookInput

	err := c.ShouldBindJSON(&book)
	if err != nil {
		var errorMessage []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage = append(errorMessage, fmt.Sprintf("error on field %s, condition, %s", e.Field(), e.ActualTag()))
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessage,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": book.Title,
		"price": book.Price,
	})
}
