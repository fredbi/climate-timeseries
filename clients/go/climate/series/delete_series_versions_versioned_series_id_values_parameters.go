// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteSeriesVersionsVersionedSeriesIDValuesParams creates a new DeleteSeriesVersionsVersionedSeriesIDValuesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSeriesVersionsVersionedSeriesIDValuesParams() *DeleteSeriesVersionsVersionedSeriesIDValuesParams {
	return &DeleteSeriesVersionsVersionedSeriesIDValuesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSeriesVersionsVersionedSeriesIDValuesParamsWithTimeout creates a new DeleteSeriesVersionsVersionedSeriesIDValuesParams object
// with the ability to set a timeout on a request.
func NewDeleteSeriesVersionsVersionedSeriesIDValuesParamsWithTimeout(timeout time.Duration) *DeleteSeriesVersionsVersionedSeriesIDValuesParams {
	return &DeleteSeriesVersionsVersionedSeriesIDValuesParams{
		timeout: timeout,
	}
}

// NewDeleteSeriesVersionsVersionedSeriesIDValuesParamsWithContext creates a new DeleteSeriesVersionsVersionedSeriesIDValuesParams object
// with the ability to set a context for a request.
func NewDeleteSeriesVersionsVersionedSeriesIDValuesParamsWithContext(ctx context.Context) *DeleteSeriesVersionsVersionedSeriesIDValuesParams {
	return &DeleteSeriesVersionsVersionedSeriesIDValuesParams{
		Context: ctx,
	}
}

// NewDeleteSeriesVersionsVersionedSeriesIDValuesParamsWithHTTPClient creates a new DeleteSeriesVersionsVersionedSeriesIDValuesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSeriesVersionsVersionedSeriesIDValuesParamsWithHTTPClient(client *http.Client) *DeleteSeriesVersionsVersionedSeriesIDValuesParams {
	return &DeleteSeriesVersionsVersionedSeriesIDValuesParams{
		HTTPClient: client,
	}
}

/* DeleteSeriesVersionsVersionedSeriesIDValuesParams contains all the parameters to send to the API endpoint
   for the delete series versions versioned series ID values operation.

   Typically these are written to a http.Request.
*/
type DeleteSeriesVersionsVersionedSeriesIDValuesParams struct {

	/* VersionedSeriesID.

	   The UUID of a version of the time series.

	   Format: uuid
	*/
	VersionedSeriesID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete series versions versioned series ID values params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesParams) WithDefaults() *DeleteSeriesVersionsVersionedSeriesIDValuesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete series versions versioned series ID values params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete series versions versioned series ID values params
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesParams) WithTimeout(timeout time.Duration) *DeleteSeriesVersionsVersionedSeriesIDValuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete series versions versioned series ID values params
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete series versions versioned series ID values params
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesParams) WithContext(ctx context.Context) *DeleteSeriesVersionsVersionedSeriesIDValuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete series versions versioned series ID values params
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete series versions versioned series ID values params
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesParams) WithHTTPClient(client *http.Client) *DeleteSeriesVersionsVersionedSeriesIDValuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete series versions versioned series ID values params
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVersionedSeriesID adds the versionedSeriesID to the delete series versions versioned series ID values params
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesParams) WithVersionedSeriesID(versionedSeriesID strfmt.UUID) *DeleteSeriesVersionsVersionedSeriesIDValuesParams {
	o.SetVersionedSeriesID(versionedSeriesID)
	return o
}

// SetVersionedSeriesID adds the versionedSeriesId to the delete series versions versioned series ID values params
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesParams) SetVersionedSeriesID(versionedSeriesID strfmt.UUID) {
	o.VersionedSeriesID = versionedSeriesID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param versionedSeriesId
	if err := r.SetPathParam("versionedSeriesId", o.VersionedSeriesID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}