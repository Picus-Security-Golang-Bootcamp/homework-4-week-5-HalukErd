package main

import (
	postgres "github.com/HalukErd/Week5Assignment/common/db"
	"github.com/HalukErd/Week5Assignment/domain/author"
	"github.com/HalukErd/Week5Assignment/domain/book"
	"github.com/HalukErd/Week5Assignment/routers"
	"github.com/HalukErd/Week5Assignment/service"
)

func main() {
	db := postgres.InitDb()
	bookRepo := book.NewBookRepo(db)
	authorRepo := author.NewAuthorRepo(db)
	bookService := service.NewBookService(bookRepo)
	authorService := service.NewAuthorService(authorRepo)

	postgres.ReadFromCsvAndInsertSampleData(authorRepo, bookRepo)
	controller := routers.NewApiController(bookService, authorService)
	controller.InitRouter()
}
