// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSeriesSeriesIDHandlerFunc turns a function with the right signature into a get series series ID handler
type GetSeriesSeriesIDHandlerFunc func(GetSeriesSeriesIDParams) GetSeriesSeriesIDResponder

// Handle executing the request and returning a response
func (fn GetSeriesSeriesIDHandlerFunc) Handle(params GetSeriesSeriesIDParams) GetSeriesSeriesIDResponder {
	return fn(params)
}

// GetSeriesSeriesIDHandler interface for that can handle valid get series series ID params
type GetSeriesSeriesIDHandler interface {
	Handle(GetSeriesSeriesIDParams) GetSeriesSeriesIDResponder
}

// NewGetSeriesSeriesID creates a new http.Handler for the get series series ID operation
func NewGetSeriesSeriesID(ctx *middleware.Context, handler GetSeriesSeriesIDHandler) *GetSeriesSeriesID {
	return &GetSeriesSeriesID{Context: ctx, Handler: handler}
}

/* GetSeriesSeriesID swagger:route GET /series/{seriesId} series getSeriesSeriesId

Get all versions of a time series.

*/
type GetSeriesSeriesID struct {
	Context *middleware.Context
	Handler GetSeriesSeriesIDHandler
}

func (o *GetSeriesSeriesID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSeriesSeriesIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}