package handler

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"pronto-hubble/api/v1/models"
	"pronto-hubble/api/v1/restapi/operations/users"
	"pronto-hubble/config"
	"pronto-hubble/service"
)

var (
	userService = getUserService()
)

func getUserService() *service.UserService {
	serviceFactory := service.InitServiceFactory(context.Background(), config.GetDbConfig())
	return serviceFactory.UserService
}

func CreateUser(params users.CreateUserParams) middleware.Responder {
	id, err := userService.CreateUser(context.Background(), params.Body)
	if id != "" {
		return users.NewCreateUserCreated().WithPayload(&models.ID{ID: &id})
	} else {
		errMsg := err.Error()
		return users.NewCreateUserDefault(500).WithPayload(&models.Error{Message: &errMsg})
	}
}

func ListAllUsers(params users.ListUsersParams) middleware.Responder {
	usersList, err := userService.FindAllUsers(context.Background())
	if err != nil {
		errMsg := err.Error()
		return users.NewListUsersDefault(500).WithPayload(&models.Error{Message: &errMsg})
	} else {
		return users.NewListUsersOK().WithPayload(usersList)
	}
}

func GetUser(params users.GetUserParams) middleware.Responder {
	id := params.ID
	user, err := userService.FindUser(context.Background(), id)
	if err != nil {
		errMsg := err.Error()
		return users.NewGetUserDefault(500).WithPayload(&models.Error{Message: &errMsg})
	} else {
		return users.NewGetUserOK().WithPayload(user)
	}
}
