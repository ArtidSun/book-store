package usecase

import (
	"book-store/models"
	"book-store/services"
	"context"
	"sync"

	helperModel "git.innovasive.co.th/backend/models"
)

type bookUsecase struct {
	bookRepo services.RepositoryInterface
}

func NewBookUsecase(bookRepo services.RepositoryInterface) services.UsecaseInterface {
	return &bookUsecase{
		bookRepo: bookRepo,
	}
}

func (b bookUsecase) FetchListBooks(ctx context.Context, args *sync.Map, paginator *helperModel.Paginator) ([]*models.Books, error) {
	return b.bookRepo.FetchListBooks(ctx, args, paginator)
}
