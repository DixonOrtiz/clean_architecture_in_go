package book_repository

import "clean_architecture_in_go/src/entities"

type BookRepositoryInterface interface {
	ListBookByID(ID int) (entities.Book, error)
	ListAllBooks() ([]entities.Book, error)
	ListBookByFields(book entities.Book) (entities.Book, error)
	CreateBook(book entities.Book) (entities.Book, error)
}
