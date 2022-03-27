// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostSeriesVersionsVersionedSeriesIDHandlerFunc turns a function with the right signature into a post series versions versioned series ID handler
type PostSeriesVersionsVersionedSeriesIDHandlerFunc func(PostSeriesVersionsVersionedSeriesIDParams) PostSeriesVersionsVersionedSeriesIDResponder

// Handle executing the request and returning a response
func (fn PostSeriesVersionsVersionedSeriesIDHandlerFunc) Handle(params PostSeriesVersionsVersionedSeriesIDParams) PostSeriesVersionsVersionedSeriesIDResponder {
	return fn(params)
}

// PostSeriesVersionsVersionedSeriesIDHandler interface for that can handle valid post series versions versioned series ID params
type PostSeriesVersionsVersionedSeriesIDHandler interface {
	Handle(PostSeriesVersionsVersionedSeriesIDParams) PostSeriesVersionsVersionedSeriesIDResponder
}

// NewPostSeriesVersionsVersionedSeriesID creates a new http.Handler for the post series versions versioned series ID operation
func NewPostSeriesVersionsVersionedSeriesID(ctx *middleware.Context, handler PostSeriesVersionsVersionedSeriesIDHandler) *PostSeriesVersionsVersionedSeriesID {
	return &PostSeriesVersionsVersionedSeriesID{Context: ctx, Handler: handler}
}

/* PostSeriesVersionsVersionedSeriesID swagger:route POST /series/versions/{versionedSeriesId} series versioning postSeriesVersionsVersionedSeriesId

creates a new version of a time series

A conflict is reported if the version was already attributed to a version of the same series. Use PUT to update existing values.


*/
type PostSeriesVersionsVersionedSeriesID struct {
	Context *middleware.Context
	Handler PostSeriesVersionsVersionedSeriesIDHandler
}

func (o *PostSeriesVersionsVersionedSeriesID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostSeriesVersionsVersionedSeriesIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}