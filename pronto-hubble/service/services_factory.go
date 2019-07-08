package service

import (
	"context"
	"pronto-hubble/config"
	"pronto-hubble/repo/mongo"
)

type ServiceFactory struct {
	UserService *UserService
}

func InitServiceFactory(ctx context.Context, dbConfig config.DbConfig) *ServiceFactory {

	repoFactory := mongo.InitMongoRepoFactory(ctx, dbConfig) //db factory should be based on dbType

	return &ServiceFactory{
		UserService: &UserService{
			repoFactory.User(),
		},
	}
}
