package book_usecases

import "clean_architecture_in_go/src/entities"

type BookUsecaseInterface interface {
	ListAllBooks(UUID string) ([]entities.Book, string, error)
	ListBookByID(UUID string, ID int) (entities.Book, string, error)
	CreateBook(UUID string, book entities.Book) (entities.Book, string, error)
}
