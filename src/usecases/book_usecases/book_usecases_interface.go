package book_usecases

import "clean_architecture_in_go/src/entities"

type BookUsecaseInterface interface {
	ListAllBooks() ([]entities.Book, string, error)
	ListBookByID(ID int) (entities.Book, string, error)
	CreateBook(book entities.Book) (entities.Book, string, error)
}
