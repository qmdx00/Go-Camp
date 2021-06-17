package biz

import (
	"context"
	"week04/app/book/internal/data/ent"
	where "week04/app/book/internal/data/ent/book"
)

type BookService struct{}

// CreateBook ...
func (*BookService) CreateBook(ctx context.Context, client *ent.Client) (*ent.Book, error) {
	return client.Book.
		Create().
		SetName("Hello World").
		SetAuthor("Wimi").
		Save(ctx)
}

// DeleteBookById ...
func (*BookService) DeleteBookById(ctx context.Context, client *ent.Client, id int) error {
	return client.Book.
		DeleteOneID(id).
		Exec(ctx)
}

// QueryBookById ...
func (*BookService) QueryBookById(ctx context.Context, client *ent.Client, id int) (*ent.Book, error) {
	return client.Book.
		Query().
		Where(where.IDEQ(id)).
		Only(ctx)
}

// UpdateBookNameById ...
func (*BookService) UpdateBookNameById(ctx context.Context, client *ent.Client, id int, name string) (*ent.Book, error) {
	return client.Book.
		UpdateOneID(id).
		SetName(name).
		Save(ctx)
}
