// Code generated by go-swagger; DO NOT EDIT.

package classes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/fredbi/climate-timeseries/pkg/auth"
)

// PutClassesClassIDHandlerFunc turns a function with the right signature into a put classes class ID handler
type PutClassesClassIDHandlerFunc func(PutClassesClassIDParams, auth.Principal) PutClassesClassIDResponder

// Handle executing the request and returning a response
func (fn PutClassesClassIDHandlerFunc) Handle(params PutClassesClassIDParams, principal auth.Principal) PutClassesClassIDResponder {
	return fn(params, principal)
}

// PutClassesClassIDHandler interface for that can handle valid put classes class ID params
type PutClassesClassIDHandler interface {
	Handle(PutClassesClassIDParams, auth.Principal) PutClassesClassIDResponder
}

// NewPutClassesClassID creates a new http.Handler for the put classes class ID operation
func NewPutClassesClassID(ctx *middleware.Context, handler PutClassesClassIDHandler) *PutClassesClassID {
	return &PutClassesClassID{Context: ctx, Handler: handler}
}

/* PutClassesClassID swagger:route PUT /classes/{classId} classes putClassesClassId

Update descriptive metadata about a nomenclature class

*/
type PutClassesClassID struct {
	Context *middleware.Context
	Handler PutClassesClassIDHandler
}

func (o *PutClassesClassID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutClassesClassIDParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal auth.Principal
	if uprinc != nil {
		principal = uprinc.(auth.Principal) // this is really a auth.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
