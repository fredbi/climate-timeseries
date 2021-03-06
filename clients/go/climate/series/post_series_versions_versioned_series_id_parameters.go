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

	"github.com/fredbi/climate-timeseries/pkg/restapi/models"
)

// NewPostSeriesVersionsVersionedSeriesIDParams creates a new PostSeriesVersionsVersionedSeriesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSeriesVersionsVersionedSeriesIDParams() *PostSeriesVersionsVersionedSeriesIDParams {
	return &PostSeriesVersionsVersionedSeriesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSeriesVersionsVersionedSeriesIDParamsWithTimeout creates a new PostSeriesVersionsVersionedSeriesIDParams object
// with the ability to set a timeout on a request.
func NewPostSeriesVersionsVersionedSeriesIDParamsWithTimeout(timeout time.Duration) *PostSeriesVersionsVersionedSeriesIDParams {
	return &PostSeriesVersionsVersionedSeriesIDParams{
		timeout: timeout,
	}
}

// NewPostSeriesVersionsVersionedSeriesIDParamsWithContext creates a new PostSeriesVersionsVersionedSeriesIDParams object
// with the ability to set a context for a request.
func NewPostSeriesVersionsVersionedSeriesIDParamsWithContext(ctx context.Context) *PostSeriesVersionsVersionedSeriesIDParams {
	return &PostSeriesVersionsVersionedSeriesIDParams{
		Context: ctx,
	}
}

// NewPostSeriesVersionsVersionedSeriesIDParamsWithHTTPClient creates a new PostSeriesVersionsVersionedSeriesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSeriesVersionsVersionedSeriesIDParamsWithHTTPClient(client *http.Client) *PostSeriesVersionsVersionedSeriesIDParams {
	return &PostSeriesVersionsVersionedSeriesIDParams{
		HTTPClient: client,
	}
}

/* PostSeriesVersionsVersionedSeriesIDParams contains all the parameters to send to the API endpoint
   for the post series versions versioned series ID operation.

   Typically these are written to a http.Request.
*/
type PostSeriesVersionsVersionedSeriesIDParams struct {

	/* VersionSeries.

	   The description of a version of a time series.

	*/
	VersionSeries models.VersionedSeries

	/* VersionedSeriesID.

	   The UUID of a version of the time series.

	   Format: uuid
	*/
	VersionedSeriesID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post series versions versioned series ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSeriesVersionsVersionedSeriesIDParams) WithDefaults() *PostSeriesVersionsVersionedSeriesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post series versions versioned series ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSeriesVersionsVersionedSeriesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post series versions versioned series ID params
func (o *PostSeriesVersionsVersionedSeriesIDParams) WithTimeout(timeout time.Duration) *PostSeriesVersionsVersionedSeriesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post series versions versioned series ID params
func (o *PostSeriesVersionsVersionedSeriesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post series versions versioned series ID params
func (o *PostSeriesVersionsVersionedSeriesIDParams) WithContext(ctx context.Context) *PostSeriesVersionsVersionedSeriesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post series versions versioned series ID params
func (o *PostSeriesVersionsVersionedSeriesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post series versions versioned series ID params
func (o *PostSeriesVersionsVersionedSeriesIDParams) WithHTTPClient(client *http.Client) *PostSeriesVersionsVersionedSeriesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post series versions versioned series ID params
func (o *PostSeriesVersionsVersionedSeriesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVersionSeries adds the versionSeries to the post series versions versioned series ID params
func (o *PostSeriesVersionsVersionedSeriesIDParams) WithVersionSeries(versionSeries models.VersionedSeries) *PostSeriesVersionsVersionedSeriesIDParams {
	o.SetVersionSeries(versionSeries)
	return o
}

// SetVersionSeries adds the versionSeries to the post series versions versioned series ID params
func (o *PostSeriesVersionsVersionedSeriesIDParams) SetVersionSeries(versionSeries models.VersionedSeries) {
	o.VersionSeries = versionSeries
}

// WithVersionedSeriesID adds the versionedSeriesID to the post series versions versioned series ID params
func (o *PostSeriesVersionsVersionedSeriesIDParams) WithVersionedSeriesID(versionedSeriesID strfmt.UUID) *PostSeriesVersionsVersionedSeriesIDParams {
	o.SetVersionedSeriesID(versionedSeriesID)
	return o
}

// SetVersionedSeriesID adds the versionedSeriesId to the post series versions versioned series ID params
func (o *PostSeriesVersionsVersionedSeriesIDParams) SetVersionedSeriesID(versionedSeriesID strfmt.UUID) {
	o.VersionedSeriesID = versionedSeriesID
}

// WriteToRequest writes these params to a swagger request
func (o *PostSeriesVersionsVersionedSeriesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.VersionSeries); err != nil {
		return err
	}

	// path param versionedSeriesId
	if err := r.SetPathParam("versionedSeriesId", o.VersionedSeriesID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
