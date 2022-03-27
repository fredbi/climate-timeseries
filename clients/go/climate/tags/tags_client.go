// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tags API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tags API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetSearchTags(params *GetSearchTagsParams, opts ...ClientOption) (*GetSearchTagsOK, error)

	GetSearchTagsTag(params *GetSearchTagsTagParams, opts ...ClientOption) (*GetSearchTagsTagOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetSearchTags lists all known tags
*/
func (a *Client) GetSearchTags(params *GetSearchTagsParams, opts ...ClientOption) (*GetSearchTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSearchTags",
		Method:             "GET",
		PathPattern:        "/search/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchTagsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSearchTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSearchTagsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSearchTagsTag searches all entities with some tag

  This endpoint returns mulitple entities, such as classes or series.

*/
func (a *Client) GetSearchTagsTag(params *GetSearchTagsTagParams, opts ...ClientOption) (*GetSearchTagsTagOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchTagsTagParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSearchTagsTag",
		Method:             "GET",
		PathPattern:        "/search/tags/{tag}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchTagsTagReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSearchTagsTagOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSearchTagsTagDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}