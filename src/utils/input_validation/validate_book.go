package input_validation

import (
	"clean_architecture_in_go/src/entities"
	"errors"
	"fmt"
)

func ValidateBook(book entities.Book) error {
	messages := []string{}
	errorText := ""

	if book.Author == "" {
		messages = append(messages, "author is required")
	}

	if book.Title == "" {
		messages = append(messages, "title is required")
	}

	if book.Year == 0 {
		messages = append(messages, "year is required")
	}

	for _, message := range messages {
		if errorText == "" {
			errorText = message
			continue
		}
		errorText += fmt.Sprintf("; %s", message)
	}

	if errorText != "" {
		return errors.New(errorText)
	}

	return nil
}
