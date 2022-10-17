package book_usecases

import (
	"clean_architecture_in_go/src/entities"

	"github.com/stretchr/testify/mock"
)

type MockBookRepository struct {
	mock.Mock
}

func (mock MockBookRepository) ListBookByID(ID int) (entities.Book, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entities.Book), args.Error(1)
}

func (mock MockBookRepository) ListAllBooks() ([]entities.Book, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entities.Book), args.Error(1)
}
