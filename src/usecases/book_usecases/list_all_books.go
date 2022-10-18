package book_usecases

import (
	"clean_architecture_in_go/src/entities"
	"clean_architecture_in_go/src/utils/statuses"
)

func (bookUsecases *BookUsecases) ListAllBooks() ([]entities.Book, string, error) {
	books, err := bookUsecases.BookRepository.ListAllBooks()
	if err != nil {
		return nil, statuses.INTERNAL_ERROR, err
	}

	return books, statuses.SUCCESS, nil
}
