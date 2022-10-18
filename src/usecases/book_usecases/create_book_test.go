package book_usecases

import (
	"clean_architecture_in_go/src/entities"
	"clean_architecture_in_go/src/utils/statuses"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBookWithInvalidInput(t *testing.T) {
	bookUsecases := BookUsecases{}
	//
	book, status, err := bookUsecases.CreateBook("", entities.Book{
		Title:  "",
		Author: "",
		Year:   0,
	})
	//
	assert.Empty(t, book)
	assert.Equal(t, status, statuses.INPUT_ERROR)
	assert.EqualError(t, err, "author is required; title is required; year is required")
}

func TestCreateDuplicatedBook(t *testing.T) {
	mockRepo := new(MockBookRepository)
	mockRepo.On("ListBookByFields").Return(entities.Book{
		ID:     4,
		Title:  "100 años de soledad",
		Author: "Gabriel García Márquez",
		Year:   1982,
	}, nil)
	bookUsecases := BookUsecases{BookRepository: mockRepo}
	//
	book, status, err := bookUsecases.CreateBook("", entities.Book{
		Title:  "100 años de soledad",
		Author: "Gabriel García Márquez",
		Year:   1982,
	})
	//
	assert.Empty(t, book)
	assert.Equal(t, status, statuses.CONFLICT)
	assert.EqualError(t, err, "the book '100 años de soledad' already exists in our records")
}

func TestCreateDuplicatedBookWithError(t *testing.T) {
	mockRepo := new(MockBookRepository)
	mockRepo.On("ListBookByFields").Return(entities.Book{}, errors.New("db error validating duplicate"))
	bookUsecases := BookUsecases{BookRepository: mockRepo}
	//
	book, status, err := bookUsecases.CreateBook("", entities.Book{
		Title:  "100 años de soledad",
		Author: "Gabriel García Márquez",
		Year:   1982,
	})
	//
	assert.Empty(t, book)
	assert.Equal(t, status, statuses.INTERNAL_ERROR)
	assert.EqualError(t, err, "db error validating duplicate")
}

func TestCreateBookWithError(t *testing.T) {
	mockRepo := new(MockBookRepository)
	mockRepo.On("ListBookByFields").Return(entities.Book{}, nil)
	mockRepo.On("CreateBook").Return(entities.Book{}, errors.New("db error creating book"))
	bookUsecases := BookUsecases{BookRepository: mockRepo}
	//
	book, status, err := bookUsecases.CreateBook("", entities.Book{
		Title:  "Demian",
		Author: "Hermann Hesse",
		Year:   1919,
	})
	//
	assert.Empty(t, book)
	assert.Equal(t, status, statuses.INTERNAL_ERROR)
	assert.EqualError(t, err, "db error creating book")
}

func TestCreateBook(t *testing.T) {
	expectedBook := entities.Book{
		ID:     26,
		Title:  "Demian",
		Author: "Hermann Hesse",
		Year:   1919,
	}
	mockRepo := new(MockBookRepository)
	mockRepo.On("ListBookByFields").Return(entities.Book{}, nil)
	mockRepo.On("CreateBook").Return(expectedBook, nil)
	bookUsecases := BookUsecases{BookRepository: mockRepo}
	//
	book, status, err := bookUsecases.CreateBook("", entities.Book{
		Title:  "Demian",
		Author: "Hermann Hesse",
		Year:   1919,
	})
	//
	assert.Equal(t, expectedBook, book)
	assert.Equal(t, status, statuses.SUCCESS)
	assert.Nil(t, err)
}
