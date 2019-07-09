package service

import (
	"context"
	"pronto-hubble/api/v1/models"
	"pronto-hubble/repo"
)

type ClusterService struct {
	clusterRepo repo.Cluster
}

func(service *ClusterService) CreateCluster(ctx context.Context, cluster *models.Cluster) (string, error) {
	return service.clusterRepo.Store(ctx, cluster)
}

func(service *ClusterService) FindAllClusters(ctx context.Context) ([]*models.Cluster, error) {
	return service.clusterRepo.FindAll(ctx)
}

func(service *ClusterService) FindCluster(ctx context.Context, id string) (*models.Cluster, error) {
	return service.clusterRepo.FindOne(ctx, id)
}

func(service *ClusterService) AddDevices(ctx context.Context, id string, deviceIds []string) error {
	return service.clusterRepo.AddDevices(ctx, id, deviceIds)
}

func(service *ClusterService) RemoveDevices(ctx context.Context, id string, deviceIds []string) error {
	return service.clusterRepo.RemoveDevices(ctx, id, deviceIds)
}