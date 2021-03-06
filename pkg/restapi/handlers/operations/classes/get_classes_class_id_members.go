// Code generated by go-swagger; DO NOT EDIT.

package classes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetClassesClassIDMembersHandlerFunc turns a function with the right signature into a get classes class ID members handler
type GetClassesClassIDMembersHandlerFunc func(GetClassesClassIDMembersParams) GetClassesClassIDMembersResponder

// Handle executing the request and returning a response
func (fn GetClassesClassIDMembersHandlerFunc) Handle(params GetClassesClassIDMembersParams) GetClassesClassIDMembersResponder {
	return fn(params)
}

// GetClassesClassIDMembersHandler interface for that can handle valid get classes class ID members params
type GetClassesClassIDMembersHandler interface {
	Handle(GetClassesClassIDMembersParams) GetClassesClassIDMembersResponder
}

// NewGetClassesClassIDMembers creates a new http.Handler for the get classes class ID members operation
func NewGetClassesClassIDMembers(ctx *middleware.Context, handler GetClassesClassIDMembersHandler) *GetClassesClassIDMembers {
	return &GetClassesClassIDMembers{Context: ctx, Handler: handler}
}

/* GetClassesClassIDMembers swagger:route GET /classes/{classId}/members classes getClassesClassIdMembers

Get all the members of a nomenclature class

Returns a a collection of nomenclature class members.

Example:
  GET /classes/munit/members returns all measurement units.

If the request negotiates a response MIME type with text/csv (with the Accept header),
this endpoint produces a CSV file.


*/
type GetClassesClassIDMembers struct {
	Context *middleware.Context
	Handler GetClassesClassIDMembersHandler
}

func (o *GetClassesClassIDMembers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetClassesClassIDMembersParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
