// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new series API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for series API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteSeriesSeriesID(params *DeleteSeriesSeriesIDParams, opts ...ClientOption) (*DeleteSeriesSeriesIDNoContent, error)

	DeleteSeriesVersionsVersionedSeriesID(params *DeleteSeriesVersionsVersionedSeriesIDParams, opts ...ClientOption) (*DeleteSeriesVersionsVersionedSeriesIDNoContent, error)

	DeleteSeriesVersionsVersionedSeriesIDValues(params *DeleteSeriesVersionsVersionedSeriesIDValuesParams, opts ...ClientOption) (*DeleteSeriesVersionsVersionedSeriesIDValuesNoContent, error)

	GetSeries(params *GetSeriesParams, opts ...ClientOption) (*GetSeriesOK, error)

	GetSeriesSeriesID(params *GetSeriesSeriesIDParams, opts ...ClientOption) (*GetSeriesSeriesIDOK, error)

	GetSeriesSeriesIDLatest(params *GetSeriesSeriesIDLatestParams, opts ...ClientOption) (*GetSeriesSeriesIDLatestOK, error)

	GetSeriesSeriesIDLatestValues(params *GetSeriesSeriesIDLatestValuesParams, opts ...ClientOption) (*GetSeriesSeriesIDLatestValuesOK, error)

	GetSeriesSeriesIDSemver(params *GetSeriesSeriesIDSemverParams, opts ...ClientOption) (*GetSeriesSeriesIDSemverOK, error)

	GetSeriesSeriesIDSemverSemver(params *GetSeriesSeriesIDSemverSemverParams, opts ...ClientOption) (*GetSeriesSeriesIDSemverSemverOK, error)

	GetSeriesSeriesIDSemverSemverValues(params *GetSeriesSeriesIDSemverSemverValuesParams, opts ...ClientOption) (*GetSeriesSeriesIDSemverSemverValuesOK, error)

	GetSeriesVersionsVersionedSeriesID(params *GetSeriesVersionsVersionedSeriesIDParams, opts ...ClientOption) (*GetSeriesVersionsVersionedSeriesIDOK, error)

	GetSeriesVersionsVersionedSeriesIDValues(params *GetSeriesVersionsVersionedSeriesIDValuesParams, opts ...ClientOption) (*GetSeriesVersionsVersionedSeriesIDValuesOK, error)

	PostSeries(params *PostSeriesParams, opts ...ClientOption) (*PostSeriesCreated, error)

	PostSeriesSeriesID(params *PostSeriesSeriesIDParams, opts ...ClientOption) (*PostSeriesSeriesIDCreated, error)

	PostSeriesVersionsVersionedSeriesID(params *PostSeriesVersionsVersionedSeriesIDParams, opts ...ClientOption) (*PostSeriesVersionsVersionedSeriesIDCreated, error)

	PostSeriesVersionsVersionedSeriesIDValues(params *PostSeriesVersionsVersionedSeriesIDValuesParams, opts ...ClientOption) (*PostSeriesVersionsVersionedSeriesIDValuesCreated, error)

	PutSeriesSeriesID(params *PutSeriesSeriesIDParams, opts ...ClientOption) (*PutSeriesSeriesIDNoContent, error)

	PutSeriesVersionsVersionedSeriesID(params *PutSeriesVersionsVersionedSeriesIDParams, opts ...ClientOption) (*PutSeriesVersionsVersionedSeriesIDNoContent, error)

	PutSeriesVersionsVersionedSeriesIDValues(params *PutSeriesVersionsVersionedSeriesIDValuesParams, opts ...ClientOption) (*PutSeriesVersionsVersionedSeriesIDValuesNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteSeriesSeriesID deletes a time series

  All versions and values associated to this series are deleted.

Only the original owner of a series may delete it.

*/
func (a *Client) DeleteSeriesSeriesID(params *DeleteSeriesSeriesIDParams, opts ...ClientOption) (*DeleteSeriesSeriesIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSeriesSeriesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteSeriesSeriesID",
		Method:             "DELETE",
		PathPattern:        "/series/{seriesId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSeriesSeriesIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteSeriesSeriesIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSeriesSeriesIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteSeriesVersionsVersionedSeriesID deletes a version of a time series

  Only the original owner of this version of the series may delete it.

All values associated to that version are deleted.

*/
func (a *Client) DeleteSeriesVersionsVersionedSeriesID(params *DeleteSeriesVersionsVersionedSeriesIDParams, opts ...ClientOption) (*DeleteSeriesVersionsVersionedSeriesIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSeriesVersionsVersionedSeriesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteSeriesVersionsVersionedSeriesID",
		Method:             "DELETE",
		PathPattern:        "/series/versions/{versionedSeriesId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSeriesVersionsVersionedSeriesIDReader{formats: a.formats},
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
	success, ok := result.(*DeleteSeriesVersionsVersionedSeriesIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSeriesVersionsVersionedSeriesIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteSeriesVersionsVersionedSeriesIDValues deletes the values attached to a version of a time series

  Only the original owner of this version of the series may delete it.

The version is not deleted. Only the values associated to that version are deleted.

*/
func (a *Client) DeleteSeriesVersionsVersionedSeriesIDValues(params *DeleteSeriesVersionsVersionedSeriesIDValuesParams, opts ...ClientOption) (*DeleteSeriesVersionsVersionedSeriesIDValuesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSeriesVersionsVersionedSeriesIDValuesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteSeriesVersionsVersionedSeriesIDValues",
		Method:             "DELETE",
		PathPattern:        "/series/versions/{versionedSeriesId}/values",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSeriesVersionsVersionedSeriesIDValuesReader{formats: a.formats},
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
	success, ok := result.(*DeleteSeriesVersionsVersionedSeriesIDValuesNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSeriesVersionsVersionedSeriesIDValuesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSeries gets all available series with search filters
*/
func (a *Client) GetSeries(params *GetSeriesParams, opts ...ClientOption) (*GetSeriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSeriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSeries",
		Method:             "GET",
		PathPattern:        "/series",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSeriesReader{formats: a.formats},
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
	success, ok := result.(*GetSeriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSeriesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSeriesSeriesID gets all versions of a time series
*/
func (a *Client) GetSeriesSeriesID(params *GetSeriesSeriesIDParams, opts ...ClientOption) (*GetSeriesSeriesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSeriesSeriesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSeriesSeriesID",
		Method:             "GET",
		PathPattern:        "/series/{seriesId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSeriesSeriesIDReader{formats: a.formats},
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
	success, ok := result.(*GetSeriesSeriesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSeriesSeriesIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSeriesSeriesIDLatest gets the latest version of the description of a time series

  The latest version is determined according to semantic versioning rules (e.g. v1.2.3 < v1.2.4).

Unless requested by the query parameters, the time series values are not returned by default.

*/
func (a *Client) GetSeriesSeriesIDLatest(params *GetSeriesSeriesIDLatestParams, opts ...ClientOption) (*GetSeriesSeriesIDLatestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSeriesSeriesIDLatestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSeriesSeriesIDLatest",
		Method:             "GET",
		PathPattern:        "/series/{seriesId}/latest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSeriesSeriesIDLatestReader{formats: a.formats},
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
	success, ok := result.(*GetSeriesSeriesIDLatestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSeriesSeriesIDLatestDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSeriesSeriesIDLatestValues gets the values from the latest version of a time series

  The latest version is determined according to semantic versioning rules (e.g. v1.2.3 < v1.2.4).

Values are returned in chronological order.

If the request negotiates a response MIME type with text/csv (with the Accept header),
this endpoint produces a CSV file.

*/
func (a *Client) GetSeriesSeriesIDLatestValues(params *GetSeriesSeriesIDLatestValuesParams, opts ...ClientOption) (*GetSeriesSeriesIDLatestValuesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSeriesSeriesIDLatestValuesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSeriesSeriesIDLatestValues",
		Method:             "GET",
		PathPattern:        "/series/{seriesId}/latest/values",
		ProducesMediaTypes: []string{"application/json", "text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSeriesSeriesIDLatestValuesReader{formats: a.formats},
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
	success, ok := result.(*GetSeriesSeriesIDLatestValuesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSeriesSeriesIDLatestValuesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSeriesSeriesIDSemver gets all semver tags associated to a series
*/
func (a *Client) GetSeriesSeriesIDSemver(params *GetSeriesSeriesIDSemverParams, opts ...ClientOption) (*GetSeriesSeriesIDSemverOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSeriesSeriesIDSemverParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSeriesSeriesIDSemver",
		Method:             "GET",
		PathPattern:        "/series/{seriesId}/semver",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSeriesSeriesIDSemverReader{formats: a.formats},
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
	success, ok := result.(*GetSeriesSeriesIDSemverOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSeriesSeriesIDSemverDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSeriesSeriesIDSemverSemver gets a version of a time series with a semver tag

  Semver tags follow semantic versioning rules (e.g. v1.2.3, v1.2.3-rc1, v1.3.0 ...).

Unless requested by the query parameters, the time series values are not returned by default.

*/
func (a *Client) GetSeriesSeriesIDSemverSemver(params *GetSeriesSeriesIDSemverSemverParams, opts ...ClientOption) (*GetSeriesSeriesIDSemverSemverOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSeriesSeriesIDSemverSemverParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSeriesSeriesIDSemverSemver",
		Method:             "GET",
		PathPattern:        "/series/{seriesId}/semver/{semver}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSeriesSeriesIDSemverSemverReader{formats: a.formats},
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
	success, ok := result.(*GetSeriesSeriesIDSemverSemverOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSeriesSeriesIDSemverSemverDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSeriesSeriesIDSemverSemverValues gets the values of version of a time series with a semver tag


If the request negotiates a response MIME type with text/csv (with the Accept header),
this endpoint produces a CSV file.

*/
func (a *Client) GetSeriesSeriesIDSemverSemverValues(params *GetSeriesSeriesIDSemverSemverValuesParams, opts ...ClientOption) (*GetSeriesSeriesIDSemverSemverValuesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSeriesSeriesIDSemverSemverValuesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSeriesSeriesIDSemverSemverValues",
		Method:             "GET",
		PathPattern:        "/series/{seriesId}/semver/{semver}/values",
		ProducesMediaTypes: []string{"application/json", "text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSeriesSeriesIDSemverSemverValuesReader{formats: a.formats},
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
	success, ok := result.(*GetSeriesSeriesIDSemverSemverValuesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSeriesSeriesIDSemverSemverValuesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSeriesVersionsVersionedSeriesID gets a version of a time series

  Unless requested by the query parameters, the time series values are not returned by default.

*/
func (a *Client) GetSeriesVersionsVersionedSeriesID(params *GetSeriesVersionsVersionedSeriesIDParams, opts ...ClientOption) (*GetSeriesVersionsVersionedSeriesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSeriesVersionsVersionedSeriesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSeriesVersionsVersionedSeriesID",
		Method:             "GET",
		PathPattern:        "/series/versions/{versionedSeriesId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSeriesVersionsVersionedSeriesIDReader{formats: a.formats},
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
	success, ok := result.(*GetSeriesVersionsVersionedSeriesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSeriesVersionsVersionedSeriesIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSeriesVersionsVersionedSeriesIDValues gets the values of version of a time series

  If the request negotiates a response MIME type with text/csv (with the Accept header),
this endpoint produces a CSV file.

*/
func (a *Client) GetSeriesVersionsVersionedSeriesIDValues(params *GetSeriesVersionsVersionedSeriesIDValuesParams, opts ...ClientOption) (*GetSeriesVersionsVersionedSeriesIDValuesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSeriesVersionsVersionedSeriesIDValuesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSeriesVersionsVersionedSeriesIDValues",
		Method:             "GET",
		PathPattern:        "/series/versions/{versionedSeriesId}/values",
		ProducesMediaTypes: []string{"application/json", "text/csv"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSeriesVersionsVersionedSeriesIDValuesReader{formats: a.formats},
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
	success, ok := result.(*GetSeriesVersionsVersionedSeriesIDValuesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSeriesVersionsVersionedSeriesIDValuesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostSeries creates a new time series

  This creates a new time time series.

If some values are specified, the initial version for the new series is created with these values.

*/
func (a *Client) PostSeries(params *PostSeriesParams, opts ...ClientOption) (*PostSeriesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSeriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostSeries",
		Method:             "POST",
		PathPattern:        "/series",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostSeriesReader{formats: a.formats},
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
	success, ok := result.(*PostSeriesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostSeriesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostSeriesSeriesID creates a new version of a time series
*/
func (a *Client) PostSeriesSeriesID(params *PostSeriesSeriesIDParams, opts ...ClientOption) (*PostSeriesSeriesIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSeriesSeriesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostSeriesSeriesID",
		Method:             "POST",
		PathPattern:        "/series/{seriesId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostSeriesSeriesIDReader{formats: a.formats},
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
	success, ok := result.(*PostSeriesSeriesIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostSeriesSeriesIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostSeriesVersionsVersionedSeriesID creates a new version of a time series

  A conflict is reported if the version was already attributed to a version of the same series. Use PUT to update existing values.

*/
func (a *Client) PostSeriesVersionsVersionedSeriesID(params *PostSeriesVersionsVersionedSeriesIDParams, opts ...ClientOption) (*PostSeriesVersionsVersionedSeriesIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSeriesVersionsVersionedSeriesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostSeriesVersionsVersionedSeriesID",
		Method:             "POST",
		PathPattern:        "/series/versions/{versionedSeriesId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostSeriesVersionsVersionedSeriesIDReader{formats: a.formats},
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
	success, ok := result.(*PostSeriesVersionsVersionedSeriesIDCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostSeriesVersionsVersionedSeriesIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PostSeriesVersionsVersionedSeriesIDValues uploads values to the version of a time series

  This endpoint creates time series values for the latest version.

Only the owners of the time series may add values.

A conflict is reported if values were already attributed to this version. Use PUT to update existing values.

If the request negotiates a request MIME type with text/csv (with the Content-Type header),
this endpoint consumes a CSV file.

*/
func (a *Client) PostSeriesVersionsVersionedSeriesIDValues(params *PostSeriesVersionsVersionedSeriesIDValuesParams, opts ...ClientOption) (*PostSeriesVersionsVersionedSeriesIDValuesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSeriesVersionsVersionedSeriesIDValuesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostSeriesVersionsVersionedSeriesIDValues",
		Method:             "POST",
		PathPattern:        "/series/versions/{versionedSeriesId}/values",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/csv"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostSeriesVersionsVersionedSeriesIDValuesReader{formats: a.formats},
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
	success, ok := result.(*PostSeriesVersionsVersionedSeriesIDValuesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostSeriesVersionsVersionedSeriesIDValuesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutSeriesSeriesID updates metadata about a time series

  This action only updates metadata about the series (such as associated themes, tags, etc) and does not create a new version.

Only owners registered for a series may update the series.

Any time series values specified in the input are ignored. One must update a specific version to modify the values of a time series.

*/
func (a *Client) PutSeriesSeriesID(params *PutSeriesSeriesIDParams, opts ...ClientOption) (*PutSeriesSeriesIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutSeriesSeriesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutSeriesSeriesID",
		Method:             "PUT",
		PathPattern:        "/series/{seriesId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutSeriesSeriesIDReader{formats: a.formats},
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
	success, ok := result.(*PutSeriesSeriesIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutSeriesSeriesIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutSeriesVersionsVersionedSeriesID updates the metadata of version of a time series

  This endpoint replaces the metadata for the requested version, without creating a new version.

Only the owners of the time series may modify values.

*/
func (a *Client) PutSeriesVersionsVersionedSeriesID(params *PutSeriesVersionsVersionedSeriesIDParams, opts ...ClientOption) (*PutSeriesVersionsVersionedSeriesIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutSeriesVersionsVersionedSeriesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutSeriesVersionsVersionedSeriesID",
		Method:             "PUT",
		PathPattern:        "/series/versions/{versionedSeriesId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutSeriesVersionsVersionedSeriesIDReader{formats: a.formats},
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
	success, ok := result.(*PutSeriesVersionsVersionedSeriesIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutSeriesVersionsVersionedSeriesIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PutSeriesVersionsVersionedSeriesIDValues replaces the values of version of a time series

  This endpoint replaces time series values for the requested version, without creating a new version.

Only the owners of the time series may modify values.

A conflict is reported if values were already attributed to this version. Use PUT to update existing values.

If the request negotiates a request MIME type with text/csv (with the Content-Type header),
this endpoint consumes a CSV file.

*/
func (a *Client) PutSeriesVersionsVersionedSeriesIDValues(params *PutSeriesVersionsVersionedSeriesIDValuesParams, opts ...ClientOption) (*PutSeriesVersionsVersionedSeriesIDValuesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutSeriesVersionsVersionedSeriesIDValuesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PutSeriesVersionsVersionedSeriesIDValues",
		Method:             "PUT",
		PathPattern:        "/series/versions/{versionedSeriesId}/values",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/csv"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutSeriesVersionsVersionedSeriesIDValuesReader{formats: a.formats},
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
	success, ok := result.(*PutSeriesVersionsVersionedSeriesIDValuesNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutSeriesVersionsVersionedSeriesIDValuesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}