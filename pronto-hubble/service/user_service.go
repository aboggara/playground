package service

import (
	"context"
	"pronto-hubble/api/v1/models"
	"pronto-hubble/repo"
)

type UserService struct {
	userRepo repo.User
}

func InitUserService(userRepo repo.User) *UserService {
	return &UserService {userRepo: userRepo}
}

func(service *UserService) CreateUser(ctx context.Context, user *models.User) (string, error) {
	return service.userRepo.Store(ctx, user)
}

func(service *UserService) FindAllUsers(ctx context.Context) ([]*models.User, error) {
	return service.userRepo.FindAll(ctx)
}

func(service *UserService) FindUser(ctx context.Context, id string) (*models.User, error) {
	return service.userRepo.FindOne(ctx, id)
}
