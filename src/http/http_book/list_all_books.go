package http_book

import (
	"clean_architecture_in_go/src/utils/http_response"
	"clean_architecture_in_go/src/utils/logging"
	"net/http"
)

const ListAllBooks = "ListAllBooks"

func (bookHandlers *BookHandlers) ListAllBooks(w http.ResponseWriter, r *http.Request) {
	UUID := logging.GenerateOperationID()

	books, status, err := bookHandlers.BookUsecases.ListAllBooks(UUID)
	if err != nil {
		logging.Log(UUID, LAYER, ListAllBooks, logging.ALERT, err.Error())
		http_response.Respond(w, status, nil, err)
		return
	}

	logging.Log(UUID, LAYER, ListAllBooks, logging.INFO, "books listed")
	http_response.Respond(w, status, books, nil)
}
