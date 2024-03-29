// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// LeaveClusterHandlerFunc turns a function with the right signature into a leave cluster handler
type LeaveClusterHandlerFunc func(LeaveClusterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LeaveClusterHandlerFunc) Handle(params LeaveClusterParams) middleware.Responder {
	return fn(params)
}

// LeaveClusterHandler interface for that can handle valid leave cluster params
type LeaveClusterHandler interface {
	Handle(LeaveClusterParams) middleware.Responder
}

// NewLeaveCluster creates a new http.Handler for the leave cluster operation
func NewLeaveCluster(ctx *middleware.Context, handler LeaveClusterHandler) *LeaveCluster {
	return &LeaveCluster{Context: ctx, Handler: handler}
}

/*LeaveCluster swagger:route PUT /clusters/{id}/leave clusters leaveCluster

LeaveCluster leave cluster API

*/
type LeaveCluster struct {
	Context *middleware.Context
	Handler LeaveClusterHandler
}

func (o *LeaveCluster) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewLeaveClusterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
