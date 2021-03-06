// Code generated by go-swagger; DO NOT EDIT.

package classes

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

// NewPutClassesClassIDParams creates a new PutClassesClassIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutClassesClassIDParams() *PutClassesClassIDParams {
	return &PutClassesClassIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutClassesClassIDParamsWithTimeout creates a new PutClassesClassIDParams object
// with the ability to set a timeout on a request.
func NewPutClassesClassIDParamsWithTimeout(timeout time.Duration) *PutClassesClassIDParams {
	return &PutClassesClassIDParams{
		timeout: timeout,
	}
}

// NewPutClassesClassIDParamsWithContext creates a new PutClassesClassIDParams object
// with the ability to set a context for a request.
func NewPutClassesClassIDParamsWithContext(ctx context.Context) *PutClassesClassIDParams {
	return &PutClassesClassIDParams{
		Context: ctx,
	}
}

// NewPutClassesClassIDParamsWithHTTPClient creates a new PutClassesClassIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutClassesClassIDParamsWithHTTPClient(client *http.Client) *PutClassesClassIDParams {
	return &PutClassesClassIDParams{
		HTTPClient: client,
	}
}

/* PutClassesClassIDParams contains all the parameters to send to the API endpoint
   for the put classes class ID operation.

   Typically these are written to a http.Request.
*/
type PutClassesClassIDParams struct {

	/* ClassDescription.

	   Class descriptive metadata.

	*/
	ClassDescription *models.ClassDescription

	/* ClassID.

	     The internal name of a nomenclature class.

	Valid classes are:
	  * constant: mathematical and physical constants
	  * mdimension: base measured dimensions
	  * mdomain: domains that pertain to measurements
	  * measurement: physical and economic measurements
	  * multiplier: unit multipliers aka prefixes (e.g. k,M,G...)
	  * munit: measurement units
	  * musystem: systems of measurement units
	  * ostatus: owner statuses
	  * owner: series owner
	  * period: time series periods (e.g. monthly, yearly...)
	  * role: series owner role
	  * source: data sources
	  * status: series and versions statuses
	  * theme: climate themes
	  * zone: geographical zones
	  * ztype: zone types

	*/
	ClassID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put classes class ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutClassesClassIDParams) WithDefaults() *PutClassesClassIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put classes class ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutClassesClassIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put classes class ID params
func (o *PutClassesClassIDParams) WithTimeout(timeout time.Duration) *PutClassesClassIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put classes class ID params
func (o *PutClassesClassIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put classes class ID params
func (o *PutClassesClassIDParams) WithContext(ctx context.Context) *PutClassesClassIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put classes class ID params
func (o *PutClassesClassIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put classes class ID params
func (o *PutClassesClassIDParams) WithHTTPClient(client *http.Client) *PutClassesClassIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put classes class ID params
func (o *PutClassesClassIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClassDescription adds the classDescription to the put classes class ID params
func (o *PutClassesClassIDParams) WithClassDescription(classDescription *models.ClassDescription) *PutClassesClassIDParams {
	o.SetClassDescription(classDescription)
	return o
}

// SetClassDescription adds the classDescription to the put classes class ID params
func (o *PutClassesClassIDParams) SetClassDescription(classDescription *models.ClassDescription) {
	o.ClassDescription = classDescription
}

// WithClassID adds the classID to the put classes class ID params
func (o *PutClassesClassIDParams) WithClassID(classID string) *PutClassesClassIDParams {
	o.SetClassID(classID)
	return o
}

// SetClassID adds the classId to the put classes class ID params
func (o *PutClassesClassIDParams) SetClassID(classID string) {
	o.ClassID = classID
}

// WriteToRequest writes these params to a swagger request
func (o *PutClassesClassIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ClassDescription != nil {
		if err := r.SetBodyParam(o.ClassDescription); err != nil {
			return err
		}
	}

	// path param classId
	if err := r.SetPathParam("classId", o.ClassID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
