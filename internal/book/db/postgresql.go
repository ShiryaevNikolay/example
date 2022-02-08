package db

import (
	"context"

	"github.com/ShiryaevNikolay/example/internal/book"
	"github.com/ShiryaevNikolay/example/pkg/client/postgresql"
	"github.com/ShiryaevNikolay/example/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r *repository) FindAll(ctx context.Context) ([]book.Book, error) {
	q := `
		SELECT id, name FROM public.book;
	`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	books := make([]book.Book, 0)

	for rows.Next() {
		var bk Book

		err = rows.Scan(&bk.ID, &bk.Name)
		if err != nil {
			return nil, err
		}

		books = append(books, bk.ToDomain())
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) book.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
