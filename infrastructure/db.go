package infrastructure

import (
	"bookstore/model"
	"context"
)

type DB interface {
	Create(ctx context.Context, book model.Book) (model.Book, error)
	SelectAll(ctx context.Context) ([]model.Book, error)
	Update(ctx context.Context, id string, book model.Book) (model.Book, error)
	Delete(ctx context.Context, id string) error
}
