package book_usecases

import (
	"clean_architecture_in_go/src/entities"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAllBooksWithError(t *testing.T) {
	mockRepo := new(MockBookRepository)
	mockRepo.On("ListAllBooks").Return([]entities.Book{}, errors.New("db error"))
	bookUsecases := BookUsecases{BookRepository: mockRepo}
	//
	books, status, err := bookUsecases.ListAllBooks()
	//
	assert.Nil(t, books)
	assert.Equal(t, status, http.StatusInternalServerError)
	assert.EqualError(t, err, "db error")
}

func TestListAllBooks(t *testing.T) {
	mockBooks := []entities.Book{
		{
			ID:     4,
			Title:  "100 años de soledad",
			Author: "Gabriel García Márquez",
			Year:   1982,
		},
		{
			ID:     5,
			Title:  "12 cuentos peregrinos",
			Author: "Gabriel García Márquez",
			Year:   1992,
		},
	}
	mockRepo := new(MockBookRepository)
	mockRepo.On("ListAllBooks").Return(mockBooks, nil)
	bookUsecases := BookUsecases{BookRepository: mockRepo}
	//
	books, status, err := bookUsecases.ListAllBooks()
	//
	assert.Equal(t, mockBooks, books)
	assert.Equal(t, status, http.StatusOK)
	assert.Nil(t, err)
}
