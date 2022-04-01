package service

import (
	"github.com/HalukErd/Week5Assignment/domain/author"
	"sync"
)

type AuthorService struct {
	authorRepo *author.AuthorRepo
}

var authorService *AuthorService
var onceAuthorService sync.Once

func NewAuthorService(repo *author.AuthorRepo) *AuthorService {
	onceAuthorService.Do(func() {
		authorService = &AuthorService{authorRepo: repo}
	})
	return authorService
}
