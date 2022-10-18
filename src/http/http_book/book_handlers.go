package http_book

import "clean_architecture_in_go/src/usecases/book_usecases"

type BookHandlers struct {
	BookUsecases book_usecases.BookUsecaseInterface
}
