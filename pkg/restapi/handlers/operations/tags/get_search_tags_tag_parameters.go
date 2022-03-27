// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetSearchTagsTagParams creates a new GetSearchTagsTagParams object
//
// There are no default values defined in the spec.
func NewGetSearchTagsTagParams() GetSearchTagsTagParams {

	return GetSearchTagsTagParams{}
}

// GetSearchTagsTagParams contains all the bound params for the get search tags tag operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetSearchTagsTag
type GetSearchTagsTagParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*When audit is specified, the response will also contain the audit trail field.

	  In: query
	*/
	Audit *bool
	/*When brief is specified, the response will only contain essential data and strip long descriptions.

	  In: query
	*/
	Brief *bool
	/*When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.

	  In: query
	*/
	Deep *bool
	/*Refers to some data owner by email.

	  Max Length: 100
	  Min Length: 1
	  In: query
	*/
	Email *string
	/*The ID of the contributor.

	  In: query
	*/
	OwnerID *strfmt.UUID
	/*Filter the result according to a given status.

	Only series in the "PUBLISHED" status are returned to the public.

	Series owner may consult their series in any status.

	  In: query
	*/
	Status *string
	/*Filter the result by search for a given keyword, whenever tag search is applicable.

	This parameter has no effect on objects where tag search is not applicable.

	  Required: true
	  Max Length: 100
	  Min Length: 1
	  In: path
	*/
	Tag string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSearchTagsTagParams() beforehand.
func (o *GetSearchTagsTagParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAudit, qhkAudit, _ := qs.GetOK("audit")
	if err := o.bindAudit(qAudit, qhkAudit, route.Formats); err != nil {
		res = append(res, err)
	}

	qBrief, qhkBrief, _ := qs.GetOK("brief")
	if err := o.bindBrief(qBrief, qhkBrief, route.Formats); err != nil {
		res = append(res, err)
	}

	qDeep, qhkDeep, _ := qs.GetOK("deep")
	if err := o.bindDeep(qDeep, qhkDeep, route.Formats); err != nil {
		res = append(res, err)
	}

	qEmail, qhkEmail, _ := qs.GetOK("email")
	if err := o.bindEmail(qEmail, qhkEmail, route.Formats); err != nil {
		res = append(res, err)
	}

	qOwnerID, qhkOwnerID, _ := qs.GetOK("ownerId")
	if err := o.bindOwnerID(qOwnerID, qhkOwnerID, route.Formats); err != nil {
		res = append(res, err)
	}

	qStatus, qhkStatus, _ := qs.GetOK("status")
	if err := o.bindStatus(qStatus, qhkStatus, route.Formats); err != nil {
		res = append(res, err)
	}

	rTag, rhkTag, _ := route.Params.GetOK("tag")
	if err := o.bindTag(rTag, rhkTag, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAudit binds and validates parameter Audit from query.
func (o *GetSearchTagsTagParams) bindAudit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("audit", "query", "bool", raw)
	}
	o.Audit = &value

	return nil
}

// bindBrief binds and validates parameter Brief from query.
func (o *GetSearchTagsTagParams) bindBrief(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("brief", "query", "bool", raw)
	}
	o.Brief = &value

	return nil
}

// bindDeep binds and validates parameter Deep from query.
func (o *GetSearchTagsTagParams) bindDeep(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("deep", "query", "bool", raw)
	}
	o.Deep = &value

	return nil
}

// bindEmail binds and validates parameter Email from query.
func (o *GetSearchTagsTagParams) bindEmail(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Email = &raw

	if err := o.validateEmail(formats); err != nil {
		return err
	}

	return nil
}

// validateEmail carries on validations for parameter Email
func (o *GetSearchTagsTagParams) validateEmail(formats strfmt.Registry) error {

	if err := validate.MinLength("email", "query", *o.Email, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("email", "query", *o.Email, 100); err != nil {
		return err
	}

	return nil
}

// bindOwnerID binds and validates parameter OwnerID from query.
func (o *GetSearchTagsTagParams) bindOwnerID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("ownerId", "query", "strfmt.UUID", raw)
	}
	o.OwnerID = (value.(*strfmt.UUID))

	if err := o.validateOwnerID(formats); err != nil {
		return err
	}

	return nil
}

// validateOwnerID carries on validations for parameter OwnerID
func (o *GetSearchTagsTagParams) validateOwnerID(formats strfmt.Registry) error {

	if err := validate.FormatOf("ownerId", "query", "uuid", o.OwnerID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindStatus binds and validates parameter Status from query.
func (o *GetSearchTagsTagParams) bindStatus(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Status = &raw

	if err := o.validateStatus(formats); err != nil {
		return err
	}

	return nil
}

// validateStatus carries on validations for parameter Status
func (o *GetSearchTagsTagParams) validateStatus(formats strfmt.Registry) error {

	if err := validate.EnumCase("status", "query", *o.Status, []interface{}{"VALIDATED", "PUBLISHED", "REJECTED", "PENDINGV", "PENDINGP"}, true); err != nil {
		return err
	}

	return nil
}

// bindTag binds and validates parameter Tag from path.
func (o *GetSearchTagsTagParams) bindTag(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Tag = raw

	if err := o.validateTag(formats); err != nil {
		return err
	}

	return nil
}

// validateTag carries on validations for parameter Tag
func (o *GetSearchTagsTagParams) validateTag(formats strfmt.Registry) error {

	if err := validate.MinLength("tag", "path", o.Tag, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("tag", "path", o.Tag, 100); err != nil {
		return err
	}

	return nil
}
