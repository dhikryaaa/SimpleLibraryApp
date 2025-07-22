package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"simplelibraryapp/books/entity"
	"simplelibraryapp/books/usecase"
)

type BookHandler struct {
	usecase *usecase.BookUseCase
}

func NewBookHandler(u *usecase.BookUseCase) *BookHandler {
	return &BookHandler{usecase: u}
}

func (h *BookHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/books", h.GetAll)
	router.GET("/books/:id", h.GetByID)
	router.POST("/books", h.Create)
	router.PUT("/books/:id", h.Update)
	router.DELETE("/books/:id", h.Delete)
}

func (h *BookHandler) GetAll(c *gin.Context) {
	books, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	book, err := h.usecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Create(c *gin.Context) {
	var book entity.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.usecase.Create(book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var book entity.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.usecase.Update(id, book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}