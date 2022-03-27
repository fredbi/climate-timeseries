// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteSeriesVersionsVersionedSeriesIDValuesHandlerFunc turns a function with the right signature into a delete series versions versioned series ID values handler
type DeleteSeriesVersionsVersionedSeriesIDValuesHandlerFunc func(DeleteSeriesVersionsVersionedSeriesIDValuesParams) DeleteSeriesVersionsVersionedSeriesIDValuesResponder

// Handle executing the request and returning a response
func (fn DeleteSeriesVersionsVersionedSeriesIDValuesHandlerFunc) Handle(params DeleteSeriesVersionsVersionedSeriesIDValuesParams) DeleteSeriesVersionsVersionedSeriesIDValuesResponder {
	return fn(params)
}

// DeleteSeriesVersionsVersionedSeriesIDValuesHandler interface for that can handle valid delete series versions versioned series ID values params
type DeleteSeriesVersionsVersionedSeriesIDValuesHandler interface {
	Handle(DeleteSeriesVersionsVersionedSeriesIDValuesParams) DeleteSeriesVersionsVersionedSeriesIDValuesResponder
}

// NewDeleteSeriesVersionsVersionedSeriesIDValues creates a new http.Handler for the delete series versions versioned series ID values operation
func NewDeleteSeriesVersionsVersionedSeriesIDValues(ctx *middleware.Context, handler DeleteSeriesVersionsVersionedSeriesIDValuesHandler) *DeleteSeriesVersionsVersionedSeriesIDValues {
	return &DeleteSeriesVersionsVersionedSeriesIDValues{Context: ctx, Handler: handler}
}

/* DeleteSeriesVersionsVersionedSeriesIDValues swagger:route DELETE /series/versions/{versionedSeriesId}/values series versioning deleteSeriesVersionsVersionedSeriesIdValues

Deletes the values attached to a version of a time series

Only the original owner of this version of the series may delete it.

The version is not deleted. Only the values associated to that version are deleted.


*/
type DeleteSeriesVersionsVersionedSeriesIDValues struct {
	Context *middleware.Context
	Handler DeleteSeriesVersionsVersionedSeriesIDValuesHandler
}

func (o *DeleteSeriesVersionsVersionedSeriesIDValues) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteSeriesVersionsVersionedSeriesIDValuesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}