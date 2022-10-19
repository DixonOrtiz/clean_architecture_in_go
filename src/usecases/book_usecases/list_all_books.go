package book_usecases

import (
	"clean_architecture_in_go/src/entities"
	"clean_architecture_in_go/src/utils/logging"
	"clean_architecture_in_go/src/utils/statuses"
)

const ListAllBooks = "ListAllBooks"

func (bookUsecases *BookUsecases) ListAllBooks(UUID string) ([]entities.Book, string, error) {
	log := logging.NewLog(UUID, "usecases", "ListAllBooks")

	books, err := bookUsecases.BookRepository.ListAllBooks()
	if err != nil {
		log.Error(err.Error())
		return nil, statuses.INTERNAL_ERROR, err
	}

	log.Info("books listed")
	return books, statuses.SUCCESS, nil
}
