package repo

import (
	"context"
	"pronto-hubble/api/v1/models"
)

type RepoFactory interface {
	User() User
}

type User interface {

	Store(ctx context.Context, user *models.User) (string, error)
	FindOne(ctx context.Context, id string) (*models.User, error)
	FindAll(ctx context.Context,) ([]*models.User, error)
	Update(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
}