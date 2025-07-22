package main

import (
	"github.com/gin-gonic/gin"
	"simplelibraryapp/books/handler"
	"simplelibraryapp/books/repository"
	"simplelibraryapp/books/usecase"
)

func main() {
	r := gin.Default()

	repo := repository.NewFileBookRepository("data.json")
	useCase := usecase.NewBookUseCase(repo)
	bookHandler := handler.NewBookHandler(useCase)

	bookHandler.RegisterRoutes(r)

	r.Run(":8080")
}