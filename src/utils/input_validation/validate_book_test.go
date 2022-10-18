package input_validation

import (
	"clean_architecture_in_go/src/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAuthorIsRequired(t *testing.T) {
	err := ValidateBook(entities.Book{
		Title: "Demian",
		Year:  1919,
	})

	assert.EqualError(t, err, "author is required")
}

func TestValidateTitleIsRequired(t *testing.T) {
	err := ValidateBook(entities.Book{
		Author: "Hermann Hesse",
		Year:   1919,
	})
	assert.EqualError(t, err, "title is required")
}

func TestValidateYearIsRequired(t *testing.T) {
	err := ValidateBook(entities.Book{
		Author: "Hermann Hesse",
		Title:  "Demian",
	})
	assert.EqualError(t, err, "year is required")
}

func TestValidateTwoErrors(t *testing.T) {
	err := ValidateBook(entities.Book{
		Author: "Hermann Hesse",
	})
	assert.EqualError(t, err, "title is required; year is required")
}

func TestValidateThreeErrors(t *testing.T) {
	err := ValidateBook(entities.Book{})
	assert.EqualError(t, err, "author is required; title is required; year is required")
}

func TestValidateValidBook(t *testing.T) {
	err := ValidateBook(entities.Book{
		Author: "Hermann Hesse",
		Title:  "Demian",
		Year:   1919,
	})
	assert.Nil(t, err)
}
