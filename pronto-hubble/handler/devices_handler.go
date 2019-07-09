package handler

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"pronto-hubble/api/v1/models"
	"pronto-hubble/api/v1/restapi/operations/devices"
	"pronto-hubble/config"
	"pronto-hubble/service"
)

var (
	deviceService = getDeviceService()
)

func getDeviceService() *service.DeviceService {
	serviceFactory := service.InitServiceFactory(context.Background(), config.GetDbConfig())
	return serviceFactory.DeviceService
}

func RegisterDevice(params devices.RegisterDeviceParams) middleware.Responder {
	id, err := deviceService.RegisterDevice(context.Background(), params.Body)
	if id != "" {
		return devices.NewCreateDeviceCreated().WithPayload(&models.ID{ID: &id})
	} else {
		errMsg := err.Error()
		return devices.NewCreateDeviceDefault(500).WithPayload(&models.Error{Message: &errMsg})
	}
}

func ListAllDevices(params devices.ListDevicesParams) middleware.Responder {
	devicesList, err := deviceService.FindAllDevices(context.Background())
	if err != nil {
		errMsg := err.Error()
		return devices.NewListDevicesDefault(500).WithPayload(&models.Error{Message: &errMsg})
	} else {
		return devices.NewListDevicesOK().WithPayload(devicesList)
	}
}

func GetDevice(params devices.GetDeviceParams) middleware.Responder {
	id := params.ID
	user, err := deviceService.FindDevice(context.Background(), id)
	if err != nil {
		errMsg := err.Error()
		return devices.NewGetDeviceDefault(500).WithPayload(&models.Error{Message: &errMsg})
	} else {
		return devices.NewGetDeviceOK().WithPayload(user)
	}
}
