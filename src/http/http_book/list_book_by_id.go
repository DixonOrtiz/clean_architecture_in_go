package http_book

import (
	"clean_architecture_in_go/src/utils/http_response"
	"clean_architecture_in_go/src/utils/logging"
	"clean_architecture_in_go/src/utils/request_input"
	"clean_architecture_in_go/src/utils/statuses"
	"fmt"
	"net/http"
)

const ListBookByID = "ListBookByID"

func (bookHandlers *BookHandlers) ListBookByID(w http.ResponseWriter, r *http.Request) {
	UUID := logging.GenerateTransactionID()
	log := logging.NewLog(UUID, "handlers", "ListBookByID")

	ID, err := request_input.GetIDFromPath(r)
	if err != nil {
		log.Warn(err.Error())
		http_response.Respond(w, statuses.INPUT_ERROR, nil, err)
		return
	}

	book, status, err := bookHandlers.BookUsecases.ListBookByID(UUID, ID)
	if err != nil {
		log.Error(err.Error())
		http_response.Respond(w, status, nil, err)
		return
	}

	log.Info(fmt.Sprintf("listed book with id: %d", book.ID))
	http_response.Respond(w, status, book, nil)
}
