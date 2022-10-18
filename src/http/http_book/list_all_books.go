package http_book

import (
	"clean_architecture_in_go/src/utils/http_response"
	"net/http"
)

func (bookHandlers *BookHandlers) ListAllBooks(w http.ResponseWriter, r *http.Request) {
	books, status, err := bookHandlers.BookUsecases.ListAllBooks()
	if err != nil {
		http_response.Respond(w, status, nil, err)
		return
	}

	http_response.Respond(w, status, books, nil)
}
