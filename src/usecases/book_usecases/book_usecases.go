package book_usecases

import "clean_architecture_in_go/src/repositories/book_repository"

type BookUsecases struct {
	BookRepository book_repository.BookRepositoryInterface
}
