package book

import "context"

type Repository interface {
	FindAll(ctx context.Context) ([]Book, error)
}
