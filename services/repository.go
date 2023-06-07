package services

import (
	"book-store/models"
	"context"
	"sync"

	helperModel "git.innovasive.co.th/backend/models"
)

type RepositoryInterface interface {
	FetchListBooks(ctx context.Context, args *sync.Map, paginator *helperModel.Paginator) ([]*models.Books, error)
}
