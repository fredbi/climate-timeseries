// Code generated by go-swagger; DO NOT EDIT.

package conversions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/fredbi/climate-timeseries/pkg/auth"
)

// PutConversionsFromUnitToUnitHandlerFunc turns a function with the right signature into a put conversions from unit to unit handler
type PutConversionsFromUnitToUnitHandlerFunc func(PutConversionsFromUnitToUnitParams, auth.Principal) PutConversionsFromUnitToUnitResponder

// Handle executing the request and returning a response
func (fn PutConversionsFromUnitToUnitHandlerFunc) Handle(params PutConversionsFromUnitToUnitParams, principal auth.Principal) PutConversionsFromUnitToUnitResponder {
	return fn(params, principal)
}

// PutConversionsFromUnitToUnitHandler interface for that can handle valid put conversions from unit to unit params
type PutConversionsFromUnitToUnitHandler interface {
	Handle(PutConversionsFromUnitToUnitParams, auth.Principal) PutConversionsFromUnitToUnitResponder
}

// NewPutConversionsFromUnitToUnit creates a new http.Handler for the put conversions from unit to unit operation
func NewPutConversionsFromUnitToUnit(ctx *middleware.Context, handler PutConversionsFromUnitToUnitHandler) *PutConversionsFromUnitToUnit {
	return &PutConversionsFromUnitToUnit{Context: ctx, Handler: handler}
}

/* PutConversionsFromUnitToUnit swagger:route PUT /conversions/{fromUnit}/{toUnit} conversions putConversionsFromUnitToUnit

Update an existing conversion specification from unit to unit.

*/
type PutConversionsFromUnitToUnit struct {
	Context *middleware.Context
	Handler PutConversionsFromUnitToUnitHandler
}

func (o *PutConversionsFromUnitToUnit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutConversionsFromUnitToUnitParams()
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
