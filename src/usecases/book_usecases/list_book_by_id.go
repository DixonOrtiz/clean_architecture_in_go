package book_usecases

import (
	"clean_architecture_in_go/src/entities"
	"errors"
	"fmt"
	"net/http"
)

func (bookUsecases *BookUsecases) ListBookByID(ID int) (entities.Book, int, error) {
	emptyBook := entities.Book{}

	if ID <= 0 {
		return emptyBook, http.StatusBadRequest, errors.New("the id must be a positive integer")
	}

	book, err := bookUsecases.BookRepository.ListBookByID(ID)
	if err != nil {
		return emptyBook, http.StatusInternalServerError, err
	}

	if book.ID == 0 {
		return emptyBook, http.StatusNotFound, fmt.Errorf("the book with id: %d does not exist", ID)
	}

	return book, http.StatusOK, nil
}
