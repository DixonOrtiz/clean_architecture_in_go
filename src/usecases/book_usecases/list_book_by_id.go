package book_usecases

import (
	"clean_architecture_in_go/src/entities"
	"clean_architecture_in_go/src/utils/logging"
	"clean_architecture_in_go/src/utils/statuses"
	"errors"
	"fmt"
)

const ListBookByID = "ListBookByID"

func (bookUsecases *BookUsecases) ListBookByID(UUID string, ID int) (entities.Book, string, error) {
	emptyBook := entities.Book{}

	if ID <= 0 {
		message := "the id must be a positive integer"
		logging.Log(UUID, LAYER, ListBookByID, logging.WARNING, message)
		return emptyBook, statuses.INPUT_ERROR, errors.New(message)
	}

	book, err := bookUsecases.BookRepository.ListBookByID(ID)
	if err != nil {
		logging.Log(UUID, LAYER, ListBookByID, logging.ALERT, err.Error())
		return emptyBook, statuses.INTERNAL_ERROR, err
	}

	if book.ID == 0 {
		err = fmt.Errorf("the book with id: %d does not exist", ID)
		logging.Log(UUID, LAYER, ListBookByID, logging.WARNING, err.Error())
		return emptyBook, statuses.NOT_FOUND, err
	}

	logging.Log(UUID, LAYER, ListBookByID, logging.INFO,
		fmt.Sprintf("listed book with id: %d", book.ID),
	)
	return book, statuses.SUCCESS, nil
}
