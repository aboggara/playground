package handler

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"pronto-hubble/api/v1/models"
	"pronto-hubble/api/v1/restapi/operations/clusters"
	"pronto-hubble/config"
	"pronto-hubble/service"
)

var (
	clusterService = getClusterService()
)

func getClusterService() *service.ClusterService {
	serviceFactory := service.InitServiceFactory(context.Background(), config.GetDbConfig())
	return serviceFactory.ClusterService
}

func CreateCluster(params clusters.CreateClusterParams) middleware.Responder {
	id, err := clusterService.CreateCluster(context.Background(), params.Body)
	if id != "" {
		return clusters.NewCreateClusterCreated().WithPayload(&models.ID{ID: &id})
	} else {
		errMsg := err.Error()
		return clusters.NewCreateClusterDefault(500).WithPayload(&models.Error{Message: &errMsg})
	}
}

func ListAllClusters(params clusters.ListClustersParams) middleware.Responder {
	clustersList, err := clusterService.FindAllClusters(context.Background())
	if err != nil {
		errMsg := err.Error()
		return clusters.NewListClustersDefault(500).WithPayload(&models.Error{Message: &errMsg})
	} else {
		return clusters.NewListClustersOK().WithPayload(clustersList)
	}
}

func GetCluster(params clusters.GetClusterParams) middleware.Responder {
	id := params.ID
	user, err := clusterService.FindCluster(context.Background(), id)
	if err != nil {
		errMsg := err.Error()
		return clusters.NewGetClusterDefault(500).WithPayload(&models.Error{Message: &errMsg})
	} else {
		return clusters.NewGetClusterOK().WithPayload(user)
	}
}

func JoinCluster(params clusters.JoinClusterParams) middleware.Responder {
	clusterId := params.ID
	deviceIds := params.Body.DeviceIds

	err := clusterService.AddDevices(context.Background(), clusterId, deviceIds)
	if err != nil {
		errMsg := err.Error()
		return clusters.NewJoinClusterDefault(500).WithPayload(&models.Error{Message: &errMsg})
	} else {
		return clusters.NewJoinClusterOK()
	}
}

func LeaveCluster(params clusters.LeaveClusterParams) middleware.Responder {
	clusterId := params.ID
	deviceIds := params.Body.DeviceIds

	err := clusterService.RemoveDevices(context.Background(), clusterId, deviceIds)
	if err != nil {
		errMsg := err.Error()
		return clusters.NewLeaveClusterDefault(500).WithPayload(&models.Error{Message: &errMsg})
	} else {
		return clusters.NewLeaveOK()
	}
}