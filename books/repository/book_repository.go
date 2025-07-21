package repository

import (
	"simplelibraryapp/books/entity"
	"encoding/json"
	"errors"
	"os"
)

type FileBookRepository struct {
	filePath string
}

func NewFileBookRepository(path string) *FileBookRepository {
	return &FileBookRepository{filePath: path}
}

func (r *FileBookRepository) readFromFile() ([]entity.Book, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var books []entity.Book
	if len(data) == 0 {
		return books, nil
	}
	err = json.Unmarshal(data, &books)
	return books, err
}

func (r *FileBookRepository) writeToFile(books []entity.Book) error {
	data, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		return err
	}
	return  os.WriteFile(r.filePath, data, 0644)
}

func (r *FileBookRepository) GetAll() ([]entity.Book, error) {
	return r.readFromFile()
}

func (r *FileBookRepository) GetByID(id string) (*entity.Book, error) {
	books, err := r.readFromFile()
	if err != nil {
		return nil, err
	}

	for _, book := range books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, errors.New("book not found")
}

func (r *FileBookRepository) Create(book entity.Book) error {
	books, _ := r.readFromFile()
	books = append(books, book)
	return r.writeToFile(books)
}

func (r *FileBookRepository) Update(book entity.Book) error {
	books, _ := r.readFromFile()
	for i, b := range books {
		if b.ID == book.ID {
			books[i] = book
			return r.writeToFile(books)
		}
	}
	return errors.New("book not found")
}

func (r *FileBookRepository) Delete(id string) error {
	books, _ := r.readFromFile()
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return r.writeToFile(books)
		}
	}
	return errors.New("book not found")
}