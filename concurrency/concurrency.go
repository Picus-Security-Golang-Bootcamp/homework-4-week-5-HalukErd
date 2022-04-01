package concurrency

import (
	"github.com/HalukErd/Week5Assignment/csv"
	"github.com/HalukErd/Week5Assignment/domain/author"
	"github.com/HalukErd/Week5Assignment/domain/book"
	"github.com/HalukErd/Week5Assignment/mapper/bookCsvLineMapper"
	"github.com/HalukErd/Week5Assignment/models"
	"log"
	"sync"
)

func InitializeInsertion(bookRepo *book.BookRepo, authorRepo *author.AuthorRepo) {
	const N = 5
	recordJobs := make(chan []string, N)
	bookCsvLineJobs := make(chan models.BookCsvLine, N)
	bookMapperJobChan := make(chan bookCsvLineMapper.BookMapperJob, N)
	authorMapperJobChan := make(chan bookCsvLineMapper.AuthorMapperJob, 2*N)
	bookInsertChan := make(chan book.Book, N)
	authorInsertChan := make(chan author.Author, N)
	errorChan := make(chan error, N*3)
	wg := sync.WaitGroup{}

	mapper := bookCsvLineMapper.NewBookCsvLineMapper()

	for w := 1; w <= 3; w++ {
		wg.Add(6)

		go authorRepo.InsertAuthor(authorInsertChan, errorChan, &wg)
		go bookRepo.InsertBook(bookInsertChan, errorChan, &wg)
		go mapper.ToBookStruct(bookMapperJobChan, bookInsertChan, &wg)
		go mapper.ToAuthorStruct(authorMapperJobChan, authorInsertChan, &wg)
		go mapper.ToBookAndAuthorFromBookCsvLine(bookCsvLineJobs, bookMapperJobChan, authorMapperJobChan, &wg)
		go mapper.ToBookCsvLineStruct(recordJobs, bookCsvLineJobs, &wg)
	}

	records, err := csv.ReadCsvRecords("assets/books.csv")
	if err != nil {
		log.Fatalf("ReadCsvRecords could not be initialized. Err: %v ", err)
	}

	go func() {
		for _, line := range records[1:] {
			recordJobs <- line
		}

		close(recordJobs)
	}()

	go func() {
		wg.Wait()
		close(errorChan)
	}()

	for err := range errorChan {
		if err != nil {
			log.Println(err)
		}
	}

}
