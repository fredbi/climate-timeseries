// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSeriesSeriesIDSemverSemverValuesHandlerFunc turns a function with the right signature into a get series series ID semver semver values handler
type GetSeriesSeriesIDSemverSemverValuesHandlerFunc func(GetSeriesSeriesIDSemverSemverValuesParams) GetSeriesSeriesIDSemverSemverValuesResponder

// Handle executing the request and returning a response
func (fn GetSeriesSeriesIDSemverSemverValuesHandlerFunc) Handle(params GetSeriesSeriesIDSemverSemverValuesParams) GetSeriesSeriesIDSemverSemverValuesResponder {
	return fn(params)
}

// GetSeriesSeriesIDSemverSemverValuesHandler interface for that can handle valid get series series ID semver semver values params
type GetSeriesSeriesIDSemverSemverValuesHandler interface {
	Handle(GetSeriesSeriesIDSemverSemverValuesParams) GetSeriesSeriesIDSemverSemverValuesResponder
}

// NewGetSeriesSeriesIDSemverSemverValues creates a new http.Handler for the get series series ID semver semver values operation
func NewGetSeriesSeriesIDSemverSemverValues(ctx *middleware.Context, handler GetSeriesSeriesIDSemverSemverValuesHandler) *GetSeriesSeriesIDSemverSemverValues {
	return &GetSeriesSeriesIDSemverSemverValues{Context: ctx, Handler: handler}
}

/* GetSeriesSeriesIDSemverSemverValues swagger:route GET /series/{seriesId}/semver/{semver}/values series semver getSeriesSeriesIdSemverSemverValues

Get the values of version of a time series with a semver tag


If the request negotiates a response MIME type with text/csv (with the Accept header),
this endpoint produces a CSV file.


*/
type GetSeriesSeriesIDSemverSemverValues struct {
	Context *middleware.Context
	Handler GetSeriesSeriesIDSemverSemverValuesHandler
}

func (o *GetSeriesSeriesIDSemverSemverValues) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSeriesSeriesIDSemverSemverValuesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
