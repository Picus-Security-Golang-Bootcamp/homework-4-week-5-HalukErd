package service

import (
	"github.com/HalukErd/Week5Assignment/domain/book"
	"github.com/HalukErd/Week5Assignment/models/generated/api"
	"github.com/go-openapi/strfmt"
	"sync"
)

type BookService struct {
	bookRepo *book.BookRepo
}

var bookService *BookService
var onceBookService sync.Once

func NewBookService(repo *book.BookRepo) *BookService {
	onceBookService.Do(func() {
		bookService = &BookService{bookRepo: repo}
	})
	return bookService
}

func (serv *BookService) GetAllBooksOrderedByPageLength() ([]api.Book, error) {
	bookEntities, err := serv.bookRepo.GetBooksOrderedByPageLength()
	if err != nil {
		return nil, err
	}
	var books []api.Book
	for _, bookEntity := range bookEntities {
		books = append(books, api.Book{
			AuthorID:  strfmt.UUID(bookEntity.AuthorID.String()),
			Genre:     bookEntity.Genre,
			ID:        strfmt.UUID(bookEntity.ID.String()),
			Name:      bookEntity.Name,
			Pages:     int32(bookEntity.Pages),
			Publisher: bookEntity.Publisher,
		})
	}
	return books, nil
}
