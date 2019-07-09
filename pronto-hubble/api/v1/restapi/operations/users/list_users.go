// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListUsersHandlerFunc turns a function with the right signature into a list users handler
type ListUsersHandlerFunc func(ListUsersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListUsersHandlerFunc) Handle(params ListUsersParams) middleware.Responder {
	return fn(params)
}

// ListUsersHandler interface for that can handle valid list users params
type ListUsersHandler interface {
	Handle(ListUsersParams) middleware.Responder
}

// NewListUsers creates a new http.Handler for the list users operation
func NewListUsers(ctx *middleware.Context, handler ListUsersHandler) *ListUsers {
	return &ListUsers{Context: ctx, Handler: handler}
}

/*ListUsers swagger:route GET /users users listUsers

ListUsers list users API

*/
type ListUsers struct {
	Context *middleware.Context
	Handler ListUsersHandler
}

func (o *ListUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListUsersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
