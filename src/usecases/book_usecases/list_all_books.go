package book_usecases

import (
	"clean_architecture_in_go/src/entities"
	"clean_architecture_in_go/src/utils/logging"
	"clean_architecture_in_go/src/utils/statuses"
)

const ListAllBooks = "ListAllBooks"

func (bookUsecases *BookUsecases) ListAllBooks(UUID string) ([]entities.Book, string, error) {
	books, err := bookUsecases.BookRepository.ListAllBooks()
	if err != nil {
		logging.Log(UUID, LAYER, ListAllBooks, logging.ALERT, err.Error())
		return nil, statuses.INTERNAL_ERROR, err
	}

	logging.Log(UUID, LAYER, ListAllBooks, logging.INFO, "books listed")
	return books, statuses.SUCCESS, nil
}
