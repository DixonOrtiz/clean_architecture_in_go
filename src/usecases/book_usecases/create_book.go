package book_usecases

import (
	"clean_architecture_in_go/src/entities"
	"clean_architecture_in_go/src/utils/input_validation"
	"clean_architecture_in_go/src/utils/logging"
	"clean_architecture_in_go/src/utils/statuses"
	"fmt"
)

func (bookUsecases *BookUsecases) CreateBook(UUID string, book entities.Book) (entities.Book, string, error) {
	log := logging.NewLog(UUID, "usecases", "CreateBook")
	emptyBook := entities.Book{}

	err := input_validation.ValidateBook(book)
	if err != nil {
		log.Warn(err.Error())
		return emptyBook, statuses.INPUT_ERROR, err
	}

	foundBook, err := bookUsecases.BookRepository.ListBookByFields(book)
	if err != nil {
		log.Error(err.Error())
		return emptyBook, statuses.INTERNAL_ERROR, err
	}

	if foundBook.ID != 0 {
		err = fmt.Errorf("the book '%s' already exists in our records", book.Title)
		log.Warn(err.Error())
		return emptyBook, statuses.CONFLICT, err
	}

	createdBook, err := bookUsecases.BookRepository.CreateBook(book)
	if err != nil {
		log.Error(err.Error())
		return emptyBook, statuses.INTERNAL_ERROR, err
	}

	log.Info(fmt.Sprintf("book created: %d) %s", createdBook.ID, createdBook.Title))
	return createdBook, statuses.SUCCESS, nil
}
