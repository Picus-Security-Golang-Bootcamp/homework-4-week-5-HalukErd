package postgres

import (
	"github.com/HalukErd/Week5Assignment/csv"
	"github.com/HalukErd/Week5Assignment/domain/author"
	"github.com/HalukErd/Week5Assignment/domain/book"
	"github.com/HalukErd/Week5Assignment/mapper/bookCsvLineMapper"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
)

func InitDb() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file. Error:%v", err)
	}

	db := NewPsqlDB()
	log.Println("Postgres connected")
	return db
}

func ReadFromCsvAndInsertSampleData(authorRepo *author.AuthorRepo, bookRepo *book.BookRepo) {
	bookCsvLines, err := csv.ReadCsvToBookAndAuthor("assets/books.csv")
	if err != nil {
		log.Fatal(err)
	}

	mapper := bookCsvLineMapper.NewBookCsvLineMapper()
	books, authors, _ := mapper.GetBooksAndAuthors(bookCsvLines)

	authorRepo.Migrations()
	authorRepo.InsertSampleData(authors)

	bookRepo.Migrations()
	bookRepo.InsertSampleData(books)
}
