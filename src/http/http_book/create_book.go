package http_book

import (
	"clean_architecture_in_go/src/utils/http_response"
	"clean_architecture_in_go/src/utils/logging"
	"clean_architecture_in_go/src/utils/request_input"
	"clean_architecture_in_go/src/utils/statuses"
	"fmt"
	"net/http"
)

const CreateBook = "CreateBook"

func (bookHandlers *BookHandlers) CreateBook(w http.ResponseWriter, r *http.Request) {
	UUID := logging.GenerateTransactionID()
	log := logging.NewLog(UUID, "handlers", "CreateBook")

	book, err := request_input.GetBookFromBody(r)
	if err != nil {
		log.Warn(err.Error())
		http_response.Respond(w, statuses.INPUT_ERROR, nil, err)
		return
	}

	createdBook, status, err := bookHandlers.BookUsecases.CreateBook(UUID, book)
	if err != nil {
		log.Error(err.Error())
		http_response.Respond(w, status, nil, err)
		return
	}

	log.Info(fmt.Sprintf("book created: %d) %s", createdBook.ID, createdBook.Title))
	http_response.Respond(w, status, createdBook, nil)
}
