package usecase

import (
	"bookstore/infrastructure"
	"bookstore/model"
	"context"
)

type BookStore struct {
	db infrastructure.DB
}

func NewBookStore(db infrastructure.DB) *BookStore {
	return &BookStore{
		db: db,
	}
}

func (bs *BookStore) Add(ctx context.Context, book model.Book) (model.Book, error) {
	panic("implement me")
}

func (bs *BookStore) Get(ctx context.Context) ([]model.Book, error) {
	return bs.db.SelectAll(ctx)
}

func (bs *BookStore) Update(ctx context.Context, id string, book model.Book) (model.Book, error) {
	panic("implement me")
}

func (bs *BookStore) Remove(ctx context.Context, id string) (model.Book, error) {
	panic("implement me")
}
