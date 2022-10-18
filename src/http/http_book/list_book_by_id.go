package http_book

import (
	"clean_architecture_in_go/src/utils/http_response"
	"clean_architecture_in_go/src/utils/request_input"
	"clean_architecture_in_go/src/utils/statuses"
	"net/http"
)

func (bookHandlers *BookHandlers) ListBookByID(w http.ResponseWriter, r *http.Request) {
	ID, err := request_input.GetIDFromPath(r)
	if err != nil {
		http_response.Respond(w, statuses.INPUT_ERROR, nil, err)
		return
	}

	book, status, err := bookHandlers.BookUsecases.ListBookByID(ID)
	if err != nil {
		http_response.Respond(w, status, nil, err)
		return
	}

	http_response.Respond(w, status, book, nil)
}
