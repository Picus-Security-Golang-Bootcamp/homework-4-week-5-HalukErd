package book

import (
	"fmt"
	"github.com/HalukErd/Week5Assignment/pkg"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math"
	"sync"
)

type BookRepo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (b *BookRepo) Migrations() {
	b.db.AutoMigrate(&Book{})
}

// InsertSampleData inserts data from csv file if there is no such Book
func (b *BookRepo) InsertSampleData(books Books) {
	for _, book := range books {
		b.db.Where(Book{Name: book.Name}).FirstOrCreate(&book)
	}
}

func (b *BookRepo) InsertBook(jobs <-chan Book, results chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	for book := range jobs {
		fmt.Println("data to insert ---->>>> book's authorId:", book.AuthorID)
		err := b.db.Where(Book{Name: book.Name}).FirstOrCreate(&book).Error

		results <- err
	}
}

// FindAll finds all books :)
func (b *BookRepo) FindAll() Books {
	var books Books
	b.db.Find(&books)
	return books
}

// FindBookByNameWithRawSql this is an example for raw sql usage
func (b *BookRepo) FindBookByNameWithRawSql(name string) Books {
	var books Books
	b.db.Raw("SELECT * FROM books WHERE name LIKE ?", "%"+name+"%").Scan(&books)
	return books
}

func (b *BookRepo) GetBooksWithAuthor(authorID uuid.UUID) (Books, error) {
	var books Books
	if err := b.db.Where("author_id = ?", authorID).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// GetBooksOrderedByPages orders by pages asc
func (b *BookRepo) GetBooksOrderedByPages() (Books, error) {
	var books Books
	if err := b.db.Order("pages asc").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// GetCountOfLongBooksGreaterThan gets a count to try count()
func (b *BookRepo) GetCountOfLongBooksGreaterThan(minPage int) (int64, error) {
	var count int64
	if err := b.db.Model(&Book{}).Where("pages > ?", minPage).Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

// GetOnlyNamesOfBookAndAuthorWith gets only book's name and author's name to try Rows()
func (b *BookRepo) GetOnlyNamesOfBookAndAuthorWith() (map[string]string, error) {
	m := make(map[string]string)
	rows, err := b.db.Table("books").Joins("JOIN authors ON books.author_id=authors.id").Select("books.name, authors.name").Rows()
	if err != nil {
		return nil, err
	}
	var bookName string
	var authorName string
	for rows.Next() {
		rows.Scan(&bookName, &authorName)
		m[bookName] = authorName
	}
	return m, nil
}

// GetBooksOrderedByPagesWithPaginationLol get books with pagination ordered by pages
func (b *BookRepo) GetBooksOrderedByPagesWithPaginationLol(pagination pkg.Pagination) (Books, error) {
	var books Books
	var count int64
	b.db.Model(&Book{}).Count(&count)
	pagination.TotalPages = int(math.Ceil(float64(count) / float64(pagination.PageSize)))
	offSet := (pagination.Page - 1) * pagination.PageSize
	if err := b.db.Order("pages asc").Offset(offSet).Limit(pagination.PageSize).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
