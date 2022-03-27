// Code generated by go-swagger; DO NOT EDIT.

package conversions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/fredbi/climate-timeseries/pkg/restapi/models"
)

// NewPostConversionFromUnitToUnitParams creates a new PostConversionFromUnitToUnitParams object
//
// There are no default values defined in the spec.
func NewPostConversionFromUnitToUnitParams() PostConversionFromUnitToUnitParams {

	return PostConversionFromUnitToUnitParams{}
}

// PostConversionFromUnitToUnitParams contains all the bound params for the post conversion from unit to unit operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostConversionFromUnitToUnit
type PostConversionFromUnitToUnitParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Unit conversion specification.

	  Required: true
	  In: body
	*/
	ConversionSpec models.ConversionSpec
	/*Original unit to be converted.

	  Required: true
	  Max Length: 100
	  Min Length: 1
	  In: path
	*/
	FromUnit string
	/*Original unit to be converted.

	  Required: true
	  Max Length: 100
	  Min Length: 1
	  In: path
	*/
	ToUnit string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostConversionFromUnitToUnitParams() beforehand.
func (o *PostConversionFromUnitToUnitParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ConversionSpec
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("conversionSpec", "body", ""))
			} else {
				res = append(res, errors.NewParseError("conversionSpec", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ConversionSpec = body
			}
		}
	} else {
		res = append(res, errors.Required("conversionSpec", "body", ""))
	}

	rFromUnit, rhkFromUnit, _ := route.Params.GetOK("fromUnit")
	if err := o.bindFromUnit(rFromUnit, rhkFromUnit, route.Formats); err != nil {
		res = append(res, err)
	}

	rToUnit, rhkToUnit, _ := route.Params.GetOK("toUnit")
	if err := o.bindToUnit(rToUnit, rhkToUnit, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFromUnit binds and validates parameter FromUnit from path.
func (o *PostConversionFromUnitToUnitParams) bindFromUnit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.FromUnit = raw

	if err := o.validateFromUnit(formats); err != nil {
		return err
	}

	return nil
}

// validateFromUnit carries on validations for parameter FromUnit
func (o *PostConversionFromUnitToUnitParams) validateFromUnit(formats strfmt.Registry) error {

	if err := validate.MinLength("fromUnit", "path", o.FromUnit, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("fromUnit", "path", o.FromUnit, 100); err != nil {
		return err
	}

	return nil
}

// bindToUnit binds and validates parameter ToUnit from path.
func (o *PostConversionFromUnitToUnitParams) bindToUnit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ToUnit = raw

	if err := o.validateToUnit(formats); err != nil {
		return err
	}

	return nil
}

// validateToUnit carries on validations for parameter ToUnit
func (o *PostConversionFromUnitToUnitParams) validateToUnit(formats strfmt.Registry) error {

	if err := validate.MinLength("toUnit", "path", o.ToUnit, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("toUnit", "path", o.ToUnit, 100); err != nil {
		return err
	}

	return nil
}
