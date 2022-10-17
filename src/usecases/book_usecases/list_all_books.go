package book_usecases

import (
	"clean_architecture_in_go/src/entities"
	"net/http"
)

func (bookUsecases *BookUsecases) ListAllBooks() ([]entities.Book, int, error) {
	books, err := bookUsecases.BookRepository.ListAllBooks()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return books, http.StatusOK, nil
}
