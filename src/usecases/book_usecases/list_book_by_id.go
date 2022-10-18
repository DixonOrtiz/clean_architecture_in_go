package book_usecases

import (
	"clean_architecture_in_go/src/entities"
	"clean_architecture_in_go/src/utils/statuses"
	"errors"
	"fmt"
)

func (bookUsecases *BookUsecases) ListBookByID(ID int) (entities.Book, string, error) {
	emptyBook := entities.Book{}

	if ID <= 0 {
		return emptyBook, statuses.INPUT_ERROR, errors.New("the id must be a positive integer")
	}

	book, err := bookUsecases.BookRepository.ListBookByID(ID)
	if err != nil {
		return emptyBook, statuses.INTERNAL_ERROR, err
	}

	if book.ID == 0 {
		return emptyBook, statuses.NOT_FOUND, fmt.Errorf("the book with id: %d does not exist", ID)
	}

	return book, statuses.SUCCESS, nil
}
