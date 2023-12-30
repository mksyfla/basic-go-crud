package main

import (
	"log"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection fail")
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := book.NewHandler(bookService)

	router := gin.Default()

	v1 := router.Group("v1")

	v1.GET("/books/", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBookByID)
	v1.POST("/books", bookHandler.PostBook)
	v1.PUT("/books/:id", bookHandler.UpdateBookByID)
	v1.DELETE("/books/:id", bookHandler.DeleteBookByID)

	router.Run()
}
