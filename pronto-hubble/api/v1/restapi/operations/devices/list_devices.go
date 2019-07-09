// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListDevicesHandlerFunc turns a function with the right signature into a list devices handler
type ListDevicesHandlerFunc func(ListDevicesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListDevicesHandlerFunc) Handle(params ListDevicesParams) middleware.Responder {
	return fn(params)
}

// ListDevicesHandler interface for that can handle valid list devices params
type ListDevicesHandler interface {
	Handle(ListDevicesParams) middleware.Responder
}

// NewListDevices creates a new http.Handler for the list devices operation
func NewListDevices(ctx *middleware.Context, handler ListDevicesHandler) *ListDevices {
	return &ListDevices{Context: ctx, Handler: handler}
}

/*ListDevices swagger:route GET /devices devices listDevices

ListDevices list devices API

*/
type ListDevices struct {
	Context *middleware.Context
	Handler ListDevicesHandler
}

func (o *ListDevices) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListDevicesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}