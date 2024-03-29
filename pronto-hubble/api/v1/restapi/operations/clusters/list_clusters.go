// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ListClustersHandlerFunc turns a function with the right signature into a list clusters handler
type ListClustersHandlerFunc func(ListClustersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListClustersHandlerFunc) Handle(params ListClustersParams) middleware.Responder {
	return fn(params)
}

// ListClustersHandler interface for that can handle valid list clusters params
type ListClustersHandler interface {
	Handle(ListClustersParams) middleware.Responder
}

// NewListClusters creates a new http.Handler for the list clusters operation
func NewListClusters(ctx *middleware.Context, handler ListClustersHandler) *ListClusters {
	return &ListClusters{Context: ctx, Handler: handler}
}

/*ListClusters swagger:route GET /clusters clusters listClusters

ListClusters list clusters API

*/
type ListClusters struct {
	Context *middleware.Context
	Handler ListClustersHandler
}

func (o *ListClusters) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListClustersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
