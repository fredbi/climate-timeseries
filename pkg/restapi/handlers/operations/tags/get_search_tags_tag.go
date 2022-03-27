// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSearchTagsTagHandlerFunc turns a function with the right signature into a get search tags tag handler
type GetSearchTagsTagHandlerFunc func(GetSearchTagsTagParams) GetSearchTagsTagResponder

// Handle executing the request and returning a response
func (fn GetSearchTagsTagHandlerFunc) Handle(params GetSearchTagsTagParams) GetSearchTagsTagResponder {
	return fn(params)
}

// GetSearchTagsTagHandler interface for that can handle valid get search tags tag params
type GetSearchTagsTagHandler interface {
	Handle(GetSearchTagsTagParams) GetSearchTagsTagResponder
}

// NewGetSearchTagsTag creates a new http.Handler for the get search tags tag operation
func NewGetSearchTagsTag(ctx *middleware.Context, handler GetSearchTagsTagHandler) *GetSearchTagsTag {
	return &GetSearchTagsTag{Context: ctx, Handler: handler}
}

/* GetSearchTagsTag swagger:route GET /search/tags/{tag} tags getSearchTagsTag

Search all entities with some tag

This endpoint returns mulitple entities, such as classes or series.


*/
type GetSearchTagsTag struct {
	Context *middleware.Context
	Handler GetSearchTagsTagHandler
}

func (o *GetSearchTagsTag) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSearchTagsTagParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}