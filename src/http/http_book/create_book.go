package http_book

import (
	"clean_architecture_in_go/src/utils/http_response"
	"clean_architecture_in_go/src/utils/request_input"
	"clean_architecture_in_go/src/utils/statuses"
	"net/http"
)

func (bookHandlers *BookHandlers) CreateBook(w http.ResponseWriter, r *http.Request) {
	book, err := request_input.GetBookFromBody(r)
	if err != nil {
		http_response.Respond(w, statuses.INPUT_ERROR, nil, err)
		return
	}

	createdBook, status, err := bookHandlers.BookUsecases.CreateBook(book)
	if err != nil {
		http_response.Respond(w, status, nil, err)
		return
	}

	http_response.Respond(w, status, createdBook, nil)
}
