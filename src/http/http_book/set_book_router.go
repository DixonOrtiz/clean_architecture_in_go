package http_book

import "github.com/gorilla/mux"

func SetBookRouter(handlers BookHandlers) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", handlers.ListAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.ListBookByID).Methods("GET")
	router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	return router
}
