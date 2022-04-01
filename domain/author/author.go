package author

import (
	"fmt"
	"github.com/HalukErd/Week5Assignment/domain/book"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type Author struct {
	gorm.Model
	ID    uuid.UUID   `gorm:"primaryKey"`
	Name  string      `gorm:"uniqueIndex"`
	Books []book.Book `gorm:"foreignKey:AuthorID;references:ID"`
}

type Authors []Author

// String method overrides default string method
func (b Author) String() string {
	return fmt.Sprintf("Name:%s \n", b.Name)
}

// ToString method to get all data
func (b Author) ToString() string {
	return fmt.Sprintf("Name:%s, CreatedAt:%s,  UpdatedAt:%s\n",
		b.Name,
		b.CreatedAt.Format("2006-12-30 15:04:05"),
		b.UpdatedAt.Format("2006-12-30 15:04:05"))
}

// ToStringWithBooks method to get all data with books
func (b Author) ToStringWithBooks() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("*** %s ***\n", b.Name))
	for i, book := range b.Books {
		sb.WriteString(fmt.Sprintf("Book %d: %s\n", i, book))
	}
	return sb.String()
}
