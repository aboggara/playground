// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"net/http"
	"pronto-hubble/api/v1/restapi/operations/clusters"
	"pronto-hubble/api/v1/restapi/operations/devices"
	"pronto-hubble/handler"

	"pronto-hubble/api/v1/restapi/operations"
	"pronto-hubble/api/v1/restapi/operations/users"
)

//go:generate swagger generate server --target ../../v1 --name ProntoHubble --spec ../spec/swagger.yml

func configureFlags(api *operations.ProntoHubbleAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ProntoHubbleAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	//users
	api.UsersCreateUserHandler = users.CreateUserHandlerFunc(handler.CreateUser)
	api.UsersListUsersHandler = users.ListUsersHandlerFunc(handler.ListAllUsers)
	api.UsersGetUserHandler = users.GetUserHandlerFunc(handler.GetUser)

	//devices
	api.DevicesRegisterDeviceHandler = devices.RegisterDeviceHandlerFunc(handler.RegisterDevice)
	api.DevicesListDevicesHandler = devices.ListDevicesHandlerFunc(handler.ListAllDevices)
	api.DevicesGetDeviceHandler = devices.GetDeviceHandlerFunc(handler.GetDevice)

	//clusters
	api.ClustersCreateClusterHandler = clusters.CreateClusterHandlerFunc(handler.CreateCluster)
	api.ClustersListClustersHandler = clusters.ListClustersHandlerFunc(handler.ListAllClusters)
	api.ClustersGetClusterHandler = clusters.GetClusterHandlerFunc(handler.GetCluster)
	api.ClustersJoinClusterHandler = clusters.JoinClusterHandlerFunc(handler.JoinCluster)
	api.ClustersLeaveClusterHandler = clusters.LeaveClusterHandlerFunc(handler.LeaveCluster)

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
