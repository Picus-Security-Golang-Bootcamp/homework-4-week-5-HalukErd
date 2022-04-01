package book

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primaryKey"`
	Name      string    `gorm:"uniqueIndex"`
	Genre     string
	Pages     int
	Publisher string
	AuthorID  uuid.UUID `gorm:"type:varchar(100);column:author_id"`
}

type Books []Book

// String method overrides default string method
func (b Book) String() string {
	return fmt.Sprintf("Name:%s,  Genre:%s,  Pages:%d,  Publisher:%s\n",
		b.Name,
		b.Genre,
		b.Pages,
		b.Publisher)
}

// ToString method to get all data as string
func (b Book) ToString() string {
	return fmt.Sprintf("Name:%s,  Genre:%s,  Pages:%d,  Publisher:%s,  CreatedAt:%s,  UpdatedAt:%s\n",
		b.Name,
		b.Genre,
		b.Pages,
		b.Publisher,
		b.CreatedAt.Format("2006-12-30 15:04:05"),
		b.UpdatedAt.Format("2006-12-30 15:04:05"))
}
