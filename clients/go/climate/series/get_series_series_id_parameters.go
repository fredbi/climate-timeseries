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
	"github.com/go-openapi/swag"
)

// NewGetSeriesSeriesIDParams creates a new GetSeriesSeriesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSeriesSeriesIDParams() *GetSeriesSeriesIDParams {
	return &GetSeriesSeriesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSeriesSeriesIDParamsWithTimeout creates a new GetSeriesSeriesIDParams object
// with the ability to set a timeout on a request.
func NewGetSeriesSeriesIDParamsWithTimeout(timeout time.Duration) *GetSeriesSeriesIDParams {
	return &GetSeriesSeriesIDParams{
		timeout: timeout,
	}
}

// NewGetSeriesSeriesIDParamsWithContext creates a new GetSeriesSeriesIDParams object
// with the ability to set a context for a request.
func NewGetSeriesSeriesIDParamsWithContext(ctx context.Context) *GetSeriesSeriesIDParams {
	return &GetSeriesSeriesIDParams{
		Context: ctx,
	}
}

// NewGetSeriesSeriesIDParamsWithHTTPClient creates a new GetSeriesSeriesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSeriesSeriesIDParamsWithHTTPClient(client *http.Client) *GetSeriesSeriesIDParams {
	return &GetSeriesSeriesIDParams{
		HTTPClient: client,
	}
}

/* GetSeriesSeriesIDParams contains all the parameters to send to the API endpoint
   for the get series series ID operation.

   Typically these are written to a http.Request.
*/
type GetSeriesSeriesIDParams struct {

	/* Audit.

	   When audit is specified, the response will also contain the audit trail field.

	*/
	Audit *bool

	/* Brief.

	   When brief is specified, the response will only contain essential data and strip long descriptions.

	*/
	Brief *bool

	/* Deep.

	   When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.

	*/
	Deep *bool

	/* Glob.

	   Filters the result according to a glob pattern applied on the short name of the requested object.

	*/
	Glob *string

	/* SeriesID.

	   The unique ID of a time series.

	   Format: int64
	*/
	SeriesID int64

	/* Status.

	     Filter the result according to a given status.

	Only series in the "PUBLISHED" status are returned to the public.

	Series owner may consult their series in any status.

	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get series series ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSeriesSeriesIDParams) WithDefaults() *GetSeriesSeriesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get series series ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSeriesSeriesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get series series ID params
func (o *GetSeriesSeriesIDParams) WithTimeout(timeout time.Duration) *GetSeriesSeriesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get series series ID params
func (o *GetSeriesSeriesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get series series ID params
func (o *GetSeriesSeriesIDParams) WithContext(ctx context.Context) *GetSeriesSeriesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get series series ID params
func (o *GetSeriesSeriesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get series series ID params
func (o *GetSeriesSeriesIDParams) WithHTTPClient(client *http.Client) *GetSeriesSeriesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get series series ID params
func (o *GetSeriesSeriesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get series series ID params
func (o *GetSeriesSeriesIDParams) WithAudit(audit *bool) *GetSeriesSeriesIDParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get series series ID params
func (o *GetSeriesSeriesIDParams) SetAudit(audit *bool) {
	o.Audit = audit
}

// WithBrief adds the brief to the get series series ID params
func (o *GetSeriesSeriesIDParams) WithBrief(brief *bool) *GetSeriesSeriesIDParams {
	o.SetBrief(brief)
	return o
}

// SetBrief adds the brief to the get series series ID params
func (o *GetSeriesSeriesIDParams) SetBrief(brief *bool) {
	o.Brief = brief
}

// WithDeep adds the deep to the get series series ID params
func (o *GetSeriesSeriesIDParams) WithDeep(deep *bool) *GetSeriesSeriesIDParams {
	o.SetDeep(deep)
	return o
}

// SetDeep adds the deep to the get series series ID params
func (o *GetSeriesSeriesIDParams) SetDeep(deep *bool) {
	o.Deep = deep
}

// WithGlob adds the glob to the get series series ID params
func (o *GetSeriesSeriesIDParams) WithGlob(glob *string) *GetSeriesSeriesIDParams {
	o.SetGlob(glob)
	return o
}

// SetGlob adds the glob to the get series series ID params
func (o *GetSeriesSeriesIDParams) SetGlob(glob *string) {
	o.Glob = glob
}

// WithSeriesID adds the seriesID to the get series series ID params
func (o *GetSeriesSeriesIDParams) WithSeriesID(seriesID int64) *GetSeriesSeriesIDParams {
	o.SetSeriesID(seriesID)
	return o
}

// SetSeriesID adds the seriesId to the get series series ID params
func (o *GetSeriesSeriesIDParams) SetSeriesID(seriesID int64) {
	o.SeriesID = seriesID
}

// WithStatus adds the status to the get series series ID params
func (o *GetSeriesSeriesIDParams) WithStatus(status *string) *GetSeriesSeriesIDParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get series series ID params
func (o *GetSeriesSeriesIDParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetSeriesSeriesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Audit != nil {

		// query param audit
		var qrAudit bool

		if o.Audit != nil {
			qrAudit = *o.Audit
		}
		qAudit := swag.FormatBool(qrAudit)
		if qAudit != "" {

			if err := r.SetQueryParam("audit", qAudit); err != nil {
				return err
			}
		}
	}

	if o.Brief != nil {

		// query param brief
		var qrBrief bool

		if o.Brief != nil {
			qrBrief = *o.Brief
		}
		qBrief := swag.FormatBool(qrBrief)
		if qBrief != "" {

			if err := r.SetQueryParam("brief", qBrief); err != nil {
				return err
			}
		}
	}

	if o.Deep != nil {

		// query param deep
		var qrDeep bool

		if o.Deep != nil {
			qrDeep = *o.Deep
		}
		qDeep := swag.FormatBool(qrDeep)
		if qDeep != "" {

			if err := r.SetQueryParam("deep", qDeep); err != nil {
				return err
			}
		}
	}

	if o.Glob != nil {

		// query param glob
		var qrGlob string

		if o.Glob != nil {
			qrGlob = *o.Glob
		}
		qGlob := qrGlob
		if qGlob != "" {

			if err := r.SetQueryParam("glob", qGlob); err != nil {
				return err
			}
		}
	}

	// path param seriesId
	if err := r.SetPathParam("seriesId", swag.FormatInt64(o.SeriesID)); err != nil {
		return err
	}

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}