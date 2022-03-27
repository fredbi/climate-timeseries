// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewDeleteSeriesVersionsVersionedSeriesIDValuesParams creates a new DeleteSeriesVersionsVersionedSeriesIDValuesParams object
//
// There are no default values defined in the spec.
func NewDeleteSeriesVersionsVersionedSeriesIDValuesParams() DeleteSeriesVersionsVersionedSeriesIDValuesParams {

	return DeleteSeriesVersionsVersionedSeriesIDValuesParams{}
}

// DeleteSeriesVersionsVersionedSeriesIDValuesParams contains all the bound params for the delete series versions versioned series ID values operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteSeriesVersionsVersionedSeriesIDValues
type DeleteSeriesVersionsVersionedSeriesIDValuesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The UUID of a version of the time series.
	  Required: true
	  In: path
	*/
	VersionedSeriesID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteSeriesVersionsVersionedSeriesIDValuesParams() beforehand.
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rVersionedSeriesID, rhkVersionedSeriesID, _ := route.Params.GetOK("versionedSeriesId")
	if err := o.bindVersionedSeriesID(rVersionedSeriesID, rhkVersionedSeriesID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindVersionedSeriesID binds and validates parameter VersionedSeriesID from path.
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesParams) bindVersionedSeriesID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("versionedSeriesId", "path", "strfmt.UUID", raw)
	}
	o.VersionedSeriesID = *(value.(*strfmt.UUID))

	if err := o.validateVersionedSeriesID(formats); err != nil {
		return err
	}

	return nil
}

// validateVersionedSeriesID carries on validations for parameter VersionedSeriesID
func (o *DeleteSeriesVersionsVersionedSeriesIDValuesParams) validateVersionedSeriesID(formats strfmt.Registry) error {

	if err := validate.FormatOf("versionedSeriesId", "path", "uuid", o.VersionedSeriesID.String(), formats); err != nil {
		return err
	}
	return nil
}