package service

import (
	"context"
	"pronto-hubble/api/v1/models"
	"pronto-hubble/repo"
)

type DeviceService struct {
	deviceRepo repo.Device
}

func(service *DeviceService) RegisterDevice(ctx context.Context, device *models.Device) (string, error) {
	return service.deviceRepo.Store(ctx, device)
}

func(service *DeviceService) FindAllDevices(ctx context.Context) ([]*models.Device, error) {
	return service.deviceRepo.FindAll(ctx)
}

func(service *DeviceService) FindDevice(ctx context.Context, id string) (*models.Device, error) {
	return service.deviceRepo.FindOne(ctx, id)
}
