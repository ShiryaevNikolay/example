package db

import (
	"context"

	"github.com/ShiryaevNikolay/example/internal/user"
	"github.com/ShiryaevNikolay/example/pkg/client/postgresql"
	"github.com/ShiryaevNikolay/example/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r *repository) Create(ctx context.Context, user user.User) (string, error) {
	panic("implement me")
}

func (r *repository) FindAll(ctx context.Context) (u []user.User, err error) {
	panic("implement me")
}

func (r *repository) FindOne(ctx context.Context, id string) (user.User, error) {
	panic("implement me")
}

func (r *repository) Update(ctx context.Context, user user.User) error {
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) user.Storage {
	return &repository{
		client: client,
		logger: logger,
	}
}
