package mongo

import (
	"context"
	"pronto-hubble/config"
	"pronto-hubble/repo"
)

type MongoRepoFactory struct {
	config config.DbConfig
	userRepo repo.User
}

func InitMongoRepoFactory(ctx context.Context, dbConfig config.DbConfig) *MongoRepoFactory {
	mongoClient := Init(context.Background(), dbConfig)

	return &MongoRepoFactory{
		config: dbConfig,
		userRepo: &UserRepo{client: mongoClient},
	}
}

func(f *MongoRepoFactory) User() repo.User {
	return f.userRepo
}


