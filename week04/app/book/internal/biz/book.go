package biz

import (
	"context"
	"week04/app/book/internal/data/ent"
	where "week04/app/book/internal/data/ent/book"
)

type BookRepo interface {
	CreateBook(context.Context, *ent.Client) (*ent.Book, error)
	QueryBookById(context.Context, *ent.Client, int) (*ent.Book, error)
	UpdateBookNameById(context.Context, *ent.Client, int, string) (*ent.Book, error)
	DeleteBookById(context.Context, *ent.Client, int) error
}

type BookBusiness struct{}

// CreateBook ...
func (*BookBusiness) CreateBook(ctx context.Context, client *ent.Client) (*ent.Book, error) {
	return client.Book.
		Create().
		SetName("Hello World").
		SetAuthor("Wimi").
		Save(ctx)
}

// DeleteBookById ...
func (*BookBusiness) DeleteBookById(ctx context.Context, client *ent.Client, id int) error {
	return client.Book.
		DeleteOneID(id).
		Exec(ctx)
}

// QueryBookById ...
func (*BookBusiness) QueryBookById(ctx context.Context, client *ent.Client, id int) (*ent.Book, error) {
	return client.Book.
		Query().
		Where(where.IDEQ(id)).
		Only(ctx)
}

// UpdateBookNameById ...
func (*BookBusiness) UpdateBookNameById(ctx context.Context, client *ent.Client, id int, name string) (*ent.Book, error) {
	return client.Book.
		UpdateOneID(id).
		SetName(name).
		Save(ctx)
}
