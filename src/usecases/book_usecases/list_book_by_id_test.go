package book_usecases

import (
	"clean_architecture_in_go/src/entities"
	"clean_architecture_in_go/src/utils/statuses"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListBookWithNegativeID(t *testing.T) {
	bookUsecases := BookUsecases{}
	//
	book, status, err := bookUsecases.ListBookByID(-5)
	//
	assert.Empty(t, book)
	assert.Equal(t, status, statuses.INPUT_ERROR)
	assert.EqualError(t, err, "the id must be a positive integer")
}

func TestListBookWithError(t *testing.T) {
	mockRepo := new(MockBookRepository)
	mockRepo.On("ListBookByID").Return(entities.Book{}, errors.New("db error"))
	bookUsecases := BookUsecases{BookRepository: mockRepo}
	//
	book, status, err := bookUsecases.ListBookByID(14)
	//
	assert.Empty(t, book)
	assert.Equal(t, status, statuses.INTERNAL_ERROR)
	assert.EqualError(t, err, "db error")
}

func TestListNonexistentBook(t *testing.T) {
	mockRepo := new(MockBookRepository)
	mockRepo.On("ListBookByID").Return(entities.Book{}, nil)
	bookUsecases := BookUsecases{BookRepository: mockRepo}
	//
	book, status, err := bookUsecases.ListBookByID(27)
	//
	assert.Empty(t, book)
	assert.Equal(t, status, statuses.NOT_FOUND)
	assert.EqualError(t, err, "the book with id: 27 does not exist")
}

func TestListExistentBook(t *testing.T) {
	expectedBook := entities.Book{
		ID:     4,
		Title:  "100 años de soledad",
		Author: "Gabriel García Márquez",
		Year:   1982,
	}
	mockRepo := new(MockBookRepository)
	mockRepo.On("ListBookByID").Return(expectedBook, nil)
	bookUsecases := BookUsecases{BookRepository: mockRepo}
	//
	book, status, err := bookUsecases.ListBookByID(4)
	//
	assert.Equal(t, expectedBook, book)
	assert.Equal(t, status, statuses.SUCCESS)
	assert.Nil(t, err)
}
