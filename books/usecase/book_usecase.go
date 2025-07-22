package usecase

import (
	"simplelibraryapp/books/entity"
)

type BookRepository interface {
	GetAll() ([]entity.Book, error)
	GetByID(id string) (*entity.Book, error)
	Create(book entity.Book) error
	Update(id string, book entity.Book) error
	Delete(id string) error
}

type BookUseCase struct {
	repo BookRepository
}

func NewBookUseCase(r BookRepository) *BookUseCase {
	return &BookUseCase{repo: r}
}

func (u *BookUseCase) GetAll() ([]entity.Book, error) {
	return u.repo.GetAll()
}

func (u *BookUseCase) GetByID(id string) (*entity.Book, error) {
	return u.repo.GetByID(id)
}

func (u *BookUseCase) Create(book entity.Book) error {
	return u.repo.Create(book)
}

func (u *BookUseCase) Update(id string, book entity.Book) error {
	return u.repo.Update(id, book)
}

func (u *BookUseCase) Delete(id string) error {
	return u.repo.Delete(id)
}