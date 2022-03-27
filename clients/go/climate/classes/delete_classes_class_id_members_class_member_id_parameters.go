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
	"github.com/go-openapi/swag"

	"github.com/fredbi/climate-timeseries/pkg/restapi/models"
)

// NewDeleteClassesClassIDMembersClassMemberIDParams creates a new DeleteClassesClassIDMembersClassMemberIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteClassesClassIDMembersClassMemberIDParams() *DeleteClassesClassIDMembersClassMemberIDParams {
	return &DeleteClassesClassIDMembersClassMemberIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClassesClassIDMembersClassMemberIDParamsWithTimeout creates a new DeleteClassesClassIDMembersClassMemberIDParams object
// with the ability to set a timeout on a request.
func NewDeleteClassesClassIDMembersClassMemberIDParamsWithTimeout(timeout time.Duration) *DeleteClassesClassIDMembersClassMemberIDParams {
	return &DeleteClassesClassIDMembersClassMemberIDParams{
		timeout: timeout,
	}
}

// NewDeleteClassesClassIDMembersClassMemberIDParamsWithContext creates a new DeleteClassesClassIDMembersClassMemberIDParams object
// with the ability to set a context for a request.
func NewDeleteClassesClassIDMembersClassMemberIDParamsWithContext(ctx context.Context) *DeleteClassesClassIDMembersClassMemberIDParams {
	return &DeleteClassesClassIDMembersClassMemberIDParams{
		Context: ctx,
	}
}

// NewDeleteClassesClassIDMembersClassMemberIDParamsWithHTTPClient creates a new DeleteClassesClassIDMembersClassMemberIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteClassesClassIDMembersClassMemberIDParamsWithHTTPClient(client *http.Client) *DeleteClassesClassIDMembersClassMemberIDParams {
	return &DeleteClassesClassIDMembersClassMemberIDParams{
		HTTPClient: client,
	}
}

/* DeleteClassesClassIDMembersClassMemberIDParams contains all the parameters to send to the API endpoint
   for the delete classes class ID members class member ID operation.

   Typically these are written to a http.Request.
*/
type DeleteClassesClassIDMembersClassMemberIDParams struct {

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

	/* ClassMember.

	   Class member metadata.

	*/
	ClassMember models.ClassNomenclature

	/* ClassMemberID.

	   The unique identifier of a class member.


	   Format: int64
	*/
	ClassMemberID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete classes class ID members class member ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteClassesClassIDMembersClassMemberIDParams) WithDefaults() *DeleteClassesClassIDMembersClassMemberIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete classes class ID members class member ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteClassesClassIDMembersClassMemberIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete classes class ID members class member ID params
func (o *DeleteClassesClassIDMembersClassMemberIDParams) WithTimeout(timeout time.Duration) *DeleteClassesClassIDMembersClassMemberIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete classes class ID members class member ID params
func (o *DeleteClassesClassIDMembersClassMemberIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete classes class ID members class member ID params
func (o *DeleteClassesClassIDMembersClassMemberIDParams) WithContext(ctx context.Context) *DeleteClassesClassIDMembersClassMemberIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete classes class ID members class member ID params
func (o *DeleteClassesClassIDMembersClassMemberIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete classes class ID members class member ID params
func (o *DeleteClassesClassIDMembersClassMemberIDParams) WithHTTPClient(client *http.Client) *DeleteClassesClassIDMembersClassMemberIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete classes class ID members class member ID params
func (o *DeleteClassesClassIDMembersClassMemberIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClassID adds the classID to the delete classes class ID members class member ID params
func (o *DeleteClassesClassIDMembersClassMemberIDParams) WithClassID(classID string) *DeleteClassesClassIDMembersClassMemberIDParams {
	o.SetClassID(classID)
	return o
}

// SetClassID adds the classId to the delete classes class ID members class member ID params
func (o *DeleteClassesClassIDMembersClassMemberIDParams) SetClassID(classID string) {
	o.ClassID = classID
}

// WithClassMember adds the classMember to the delete classes class ID members class member ID params
func (o *DeleteClassesClassIDMembersClassMemberIDParams) WithClassMember(classMember models.ClassNomenclature) *DeleteClassesClassIDMembersClassMemberIDParams {
	o.SetClassMember(classMember)
	return o
}

// SetClassMember adds the classMember to the delete classes class ID members class member ID params
func (o *DeleteClassesClassIDMembersClassMemberIDParams) SetClassMember(classMember models.ClassNomenclature) {
	o.ClassMember = classMember
}

// WithClassMemberID adds the classMemberID to the delete classes class ID members class member ID params
func (o *DeleteClassesClassIDMembersClassMemberIDParams) WithClassMemberID(classMemberID int64) *DeleteClassesClassIDMembersClassMemberIDParams {
	o.SetClassMemberID(classMemberID)
	return o
}

// SetClassMemberID adds the classMemberId to the delete classes class ID members class member ID params
func (o *DeleteClassesClassIDMembersClassMemberIDParams) SetClassMemberID(classMemberID int64) {
	o.ClassMemberID = classMemberID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClassesClassIDMembersClassMemberIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param classId
	if err := r.SetPathParam("classId", o.ClassID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.ClassMember); err != nil {
		return err
	}

	// path param classMemberId
	if err := r.SetPathParam("classMemberId", swag.FormatInt64(o.ClassMemberID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}