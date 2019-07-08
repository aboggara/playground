package mongo

import (
	"context"
	"pronto-hubble/config"
	"pronto-hubble/repo"
)

type MongoRepoFactory struct {
	config config.DbConfig
	userRepo repo.User
	deviceRepo repo.Device
	clusterRepo repo.Cluster
}

func InitMongoRepoFactory(ctx context.Context, dbConfig config.DbConfig) *MongoRepoFactory {
	mongoClient := Init(context.Background(), dbConfig)

	return &MongoRepoFactory{
		config: dbConfig,
		userRepo: &UserRepo{client: mongoClient},
		deviceRepo: &DeviceRepo{client: mongoClient},
		clusterRepo: &ClusterRepo{client: mongoClient},
	}
}

func(f *MongoRepoFactory) User() repo.User {
	return f.userRepo
}

func(f *MongoRepoFactory) Device() repo.Device {
	return f.deviceRepo
}

func(f *MongoRepoFactory) Cluster() repo.Cluster {
	return f.clusterRepo
}


