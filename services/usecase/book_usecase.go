package usecase

import (
	"book-store/services"
)

type bookUsecase struct {
	bookRepo services.RepositoryInterface
}

func NewBookUsecase(bookRepo services.RepositoryInterface) services.UsecaseInterface {
	return &bookUsecase{
		bookRepo: bookRepo,
	}
}
