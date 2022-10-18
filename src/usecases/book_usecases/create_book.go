package book_usecases

import (
	"clean_architecture_in_go/src/entities"
	"clean_architecture_in_go/src/utils/input_validation"
	"clean_architecture_in_go/src/utils/statuses"
	"fmt"
)

func (bookUsecases *BookUsecases) CreateBook(book entities.Book) (entities.Book, string, error) {
	emptyBook := entities.Book{}

	err := input_validation.ValidateBook(book)
	if err != nil {
		return emptyBook, statuses.INPUT_ERROR, err
	}

	foundBook, err := bookUsecases.BookRepository.ListBookByFields(book)
	if err != nil {
		return emptyBook, statuses.INTERNAL_ERROR, err
	}

	if foundBook.ID != 0 {
		return emptyBook, statuses.CONFLICT, fmt.Errorf("the book '%s' already exists in our records", book.Title)
	}

	createdBook, err := bookUsecases.BookRepository.CreateBook(book)
	if err != nil {
		return emptyBook, statuses.INTERNAL_ERROR, err
	}

	return createdBook, statuses.SUCCESS, nil
}
