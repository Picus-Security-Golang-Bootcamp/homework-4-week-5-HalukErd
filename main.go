package main

import (
	"fmt"
	postgres "github.com/HalukErd/Week5Assignment/common/db"
	"github.com/HalukErd/Week5Assignment/csv"
	"github.com/HalukErd/Week5Assignment/domain/author"
	"github.com/HalukErd/Week5Assignment/domain/book"
	"github.com/HalukErd/Week5Assignment/mapper/bookCsvLineMapper"
	"github.com/HalukErd/Week5Assignment/pkg"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file. Error:%v", err)
	}

	db, err := postgres.NewPsqlDB()
	if err != nil {
		log.Fatal("Postgres cannot init:", err)
	}
	log.Println("Postgres connected")

	bookRepo := book.NewBookRepo(db)
	authorRepo := author.NewAuthorRepo(db)

	// TODO I have implemented workerPool but has violates foreign key constraint errors.
	// probably we need priority for insert author
	//authorRepo.Migrations()
	//bookRepo.Migrations()
	//concurrency.InitializeInsertion(bookRepo, authorRepo) //has Errors

	readFromCsvAndInsertSampleData(authorRepo, bookRepo)

	fmt.Println("------1-------")
	fmt.Println(bookRepo.FindAll())

	fmt.Println("------2-------")
	fmt.Println(authorRepo.GetAllAuthorsWithoutBooks())

	fmt.Println("------3-------")
	fmt.Println(authorRepo.FindByName("Steinbeck"))

	fmt.Println("------4-------")
	uuidForAuthor, _ := uuid.Parse("bf0eda30-7d43-4146-a6b1-6104c36cff6d")
	fmt.Println(authorRepo.GetByID(uuidForAuthor))

	fmt.Println("------5-------")
	fmt.Println(bookRepo.FindBookByNameWithRawSql("More"))

	allAuthorsWithBookInfo, err := authorRepo.GetAllAuthorsWithBookInformation()
	if err != nil {
		log.Fatal("could not be get all authors with book info")
	}
	for _, authorWithBooks := range allAuthorsWithBookInfo {
		fmt.Println("*** ", authorWithBooks.Name, " ***")
		for _, b := range authorWithBooks.Books {
			fmt.Println(b.ToString())
		}
	}
	fmt.Println("------6-------")
	uuidForBook, _ := uuid.Parse("02169129-7e87-441d-bae3-d25c86c21c93")
	author, err := authorRepo.GetAuthorWithBooks(uuidForBook)

	fmt.Println(author)

	fmt.Println("------7-------")
	authorUUID, err := uuid.Parse("37eb3801-077f-4f83-bcf7-c69f44561838")
	books, err := bookRepo.GetBooksWithAuthor(authorUUID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(books)

	fmt.Println("------8-------")
	booksOrderedByPages, err := bookRepo.GetBooksOrderedByPages()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(&booksOrderedByPages)

	fmt.Println("------9-------")
	minPage := 200
	count, err := bookRepo.GetCountOfLongBooksGreaterThan(minPage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("There are %d books which has more pages than %d", count, minPage)

	fmt.Println("------10-------")
	booksAndAuthors, err := bookRepo.GetOnlyNamesOfBookAndAuthorWith()
	for k, v := range booksAndAuthors {
		fmt.Printf("Book: %s -> Author:%s\n", k, v)
	}

	fmt.Println("------11-------")
	pagination := pkg.Pagination{
		PageSize:   10,
		Page:       2,
		TotalPages: 0,
	}
	booksWithPagination, err := bookRepo.GetBooksOrderedByPagesWithPaginationLol(pagination)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(booksWithPagination)

	fmt.Println("Done!")
}

func readFromCsvAndInsertSampleData(authorRepo *author.AuthorRepo, bookRepo *book.BookRepo) {
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
