package repo

import (
	"context"
	"pronto-hubble/api/v1/models"
)

type RepoFactory interface {
	User() User
	Device() Device
	Cluster() Cluster
}

type User interface {

	Store(ctx context.Context, user *models.User) (string, error)
	FindOne(ctx context.Context, id string) (*models.User, error)
	FindAll(ctx context.Context,) ([]*models.User, error)
	Update(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
}

type Device interface {

	Store(ctx context.Context, user *models.Device) (string, error)
	FindOne(ctx context.Context, id string) (*models.Device, error)
	FindAll(ctx context.Context,) ([]*models.Device, error)
	Update(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
}

type Cluster interface {
	Store(ctx context.Context, user *models.Cluster) (string, error)
	FindOne(ctx context.Context, id string) (*models.Cluster, error)
	FindAll(ctx context.Context,) ([]*models.Cluster, error)
	Update(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
	AddDevices(ctx context.Context, id string, deviceIds []string) error
	RemoveDevices(ctx context.Context, id string, deviceIds []string) error
}