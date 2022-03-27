// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSeriesHandlerFunc turns a function with the right signature into a get series handler
type GetSeriesHandlerFunc func(GetSeriesParams) GetSeriesResponder

// Handle executing the request and returning a response
func (fn GetSeriesHandlerFunc) Handle(params GetSeriesParams) GetSeriesResponder {
	return fn(params)
}

// GetSeriesHandler interface for that can handle valid get series params
type GetSeriesHandler interface {
	Handle(GetSeriesParams) GetSeriesResponder
}

// NewGetSeries creates a new http.Handler for the get series operation
func NewGetSeries(ctx *middleware.Context, handler GetSeriesHandler) *GetSeries {
	return &GetSeries{Context: ctx, Handler: handler}
}

/* GetSeries swagger:route GET /series series getSeries

Get all available series, with search filters

*/
type GetSeries struct {
	Context *middleware.Context
	Handler GetSeriesHandler
}

func (o *GetSeries) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSeriesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}