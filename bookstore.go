package tai

import "context"

type BookStore struct{}

func NewBookStore() *BookStore {
	return &BookStore{}
}

func (*BookStore) Create(ctx context.Context, book Book) (Book, error) {
	panic("implement me")
}

func (*BookStore) Read(ctx context.Context, id string) (Book, error) {
	panic("implement me")
}

func (*BookStore) Update(ctx context.Context, id string, book Book) (Book, error) {
	panic("implement me")
}

func (*BookStore) Delete(ctx context.Context, id string) (Book, error) {
	panic("implement me")
}

type Book struct {
	ID     string
	Name   string
	Author string
}
