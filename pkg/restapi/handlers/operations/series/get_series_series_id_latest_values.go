// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSeriesSeriesIDLatestValuesHandlerFunc turns a function with the right signature into a get series series ID latest values handler
type GetSeriesSeriesIDLatestValuesHandlerFunc func(GetSeriesSeriesIDLatestValuesParams) GetSeriesSeriesIDLatestValuesResponder

// Handle executing the request and returning a response
func (fn GetSeriesSeriesIDLatestValuesHandlerFunc) Handle(params GetSeriesSeriesIDLatestValuesParams) GetSeriesSeriesIDLatestValuesResponder {
	return fn(params)
}

// GetSeriesSeriesIDLatestValuesHandler interface for that can handle valid get series series ID latest values params
type GetSeriesSeriesIDLatestValuesHandler interface {
	Handle(GetSeriesSeriesIDLatestValuesParams) GetSeriesSeriesIDLatestValuesResponder
}

// NewGetSeriesSeriesIDLatestValues creates a new http.Handler for the get series series ID latest values operation
func NewGetSeriesSeriesIDLatestValues(ctx *middleware.Context, handler GetSeriesSeriesIDLatestValuesHandler) *GetSeriesSeriesIDLatestValues {
	return &GetSeriesSeriesIDLatestValues{Context: ctx, Handler: handler}
}

/* GetSeriesSeriesIDLatestValues swagger:route GET /series/{seriesId}/latest/values series getSeriesSeriesIdLatestValues

Get the values from the latest version of a time series

The latest version is determined according to semantic versioning rules (e.g. v1.2.3 < v1.2.4).

Values are returned in chronological order.

If the request negotiates a response MIME type with text/csv (with the Accept header),
this endpoint produces a CSV file.


*/
type GetSeriesSeriesIDLatestValues struct {
	Context *middleware.Context
	Handler GetSeriesSeriesIDLatestValuesHandler
}

func (o *GetSeriesSeriesIDLatestValues) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSeriesSeriesIDLatestValuesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
