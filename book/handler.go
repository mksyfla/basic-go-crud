package book

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler interface {
	FindAll()
	FindByID()
	Create()
	Update()
	DeleteBookByID()
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service}
}

func (h *handler) GetBooks(c *gin.Context) {
	books, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	var booksResponse []BooksResponse

	for _, book := range books {
		bookResponse := convertToBookResponse(book)

		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *handler) GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	book, err := h.service.FindByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "book not found",
		})
		return
	}

	booksResponse := convertToBookResponse(book)

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *handler) PostBook(c *gin.Context) {
	var bookRequest BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		var errorMessage []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage = append(errorMessage, fmt.Sprintf("error on field %s, condition, %s", e.Field(), e.ActualTag()))
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errorMessage,
		})
		return
	}

	book, err := h.service.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": book,
	})
}

func (h *handler) UpdateBookByID(c *gin.Context) {
	var bookRequest BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		var errorMessage []string
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage = append(errorMessage, fmt.Sprintf("error on field %s, condition, %s", e.Field(), e.ActualTag()))
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errorMessage,
		})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	book, err := h.service.Update(bookRequest, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": book,
	})
}

func (h *handler) DeleteBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	book, err := h.service.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": book,
	})
}

func convertToBookResponse(bookProperty Book) BooksResponse {
	return BooksResponse{
		ID:          bookProperty.ID,
		Title:       bookProperty.Title,
		Price:       bookProperty.Price,
		Description: bookProperty.Description,
		Rating:      bookProperty.Rating,
		Discount:    bookProperty.Discount,
	}
}
