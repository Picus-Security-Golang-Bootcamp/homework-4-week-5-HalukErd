package bookCsvLineMapper

import (
	"github.com/HalukErd/Week5Assignment/domain/author"
	"github.com/HalukErd/Week5Assignment/domain/book"
	"github.com/HalukErd/Week5Assignment/models"
	"github.com/google/uuid"
	"log"
	"strconv"
	"sync"
)

type BookCsvLineMapper struct {
}

type BookMapperJob struct {
	source models.BookCsvLine
	target book.Book
}

type AuthorMapperJob struct {
	source models.BookCsvLine
	target author.Author
}

func NewBookCsvLineMapper() *BookCsvLineMapper {
	return &BookCsvLineMapper{}
}

// GetBooksAndAuthors creates Books and Authors from CsvLines
func (mapper *BookCsvLineMapper) GetBooksAndAuthors(lines models.BookCsvLines) (book.Books, author.Authors, error) {
	var books book.Books
	var authors author.Authors
	for _, line := range lines {
		author := author.Author{
			ID:   uuid.New(),
			Name: line.Author,
		}
		pages, err := strconv.Atoi(line.Height)
		if err != nil {
			return nil, nil, err
		}
		book := book.Book{
			ID:        uuid.New(),
			Name:      line.Title,
			Genre:     line.Genre,
			Pages:     pages,
			Publisher: line.Publisher,
			AuthorID:  author.ID,
		}
		authors = append(authors, author)
		books = append(books, book)
	}
	return books, authors, nil
}

func (mapper *BookCsvLineMapper) ToBookCsvLineStruct(jobs <-chan []string, results chan<- models.BookCsvLine, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		bookCsvLine := models.BookCsvLine{
			Title:     j[0],
			Author:    j[1],
			Genre:     j[2],
			Height:    j[3],
			Publisher: j[4],
		}

		results <- bookCsvLine
	}
}

func (mapper *BookCsvLineMapper) ToBookAndAuthorFromBookCsvLine(jobs <-chan models.BookCsvLine,
	bookMapperJobChan chan<- BookMapperJob,
	authorMapperJobChan chan<- AuthorMapperJob,
	wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		authorUUID := uuid.New()
		bookMapperJobChan <- BookMapperJob{source: j, target: book.Book{AuthorID: authorUUID}}
		authorMapperJobChan <- AuthorMapperJob{source: j, target: author.Author{ID: authorUUID}}
	}
}

func (mapper *BookCsvLineMapper) ToBookStruct(jobs <-chan BookMapperJob, results chan<- book.Book, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		bookCsvLine := j.source
		book := j.target
		book.ID = uuid.New()
		book.Name = bookCsvLine.Title
		pages, err := strconv.Atoi(bookCsvLine.Height)
		if err != nil {
			log.Fatal(err)
		}
		book.Pages = pages
		book.Genre = bookCsvLine.Genre
		book.Publisher = bookCsvLine.Publisher

		results <- book
	}
}

func (mapper *BookCsvLineMapper) ToAuthorStruct(jobs <-chan AuthorMapperJob, results chan<- author.Author, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		bookCsvLine := j.source
		author := j.target
		author.Name = bookCsvLine.Author

		results <- author
	}
}
