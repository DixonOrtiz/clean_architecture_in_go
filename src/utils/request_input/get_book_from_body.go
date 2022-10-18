package request_input

import (
	"clean_architecture_in_go/src/entities"
	"encoding/json"
	"net/http"
)

func GetBookFromBody(r *http.Request) (entities.Book, error) {
	var book entities.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		return entities.Book{}, err
	}

	return book, nil
}
