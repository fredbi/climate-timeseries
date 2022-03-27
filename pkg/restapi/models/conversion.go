// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Conversion Unit conversions.
//
//
// swagger:model conversion
type Conversion struct {

	// audit trail
	AuditTrail *Audit `json:"auditTrail,omitempty"`

	// from unit
	// Required: true
	FromUnit Munit `json:"fromUnit"`

	// to unit
	// Required: true
	ToUnit Munit `json:"toUnit"`

	ConversionSpec
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Conversion) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		AuditTrail *Audit `json:"auditTrail,omitempty"`

		FromUnit Munit `json:"fromUnit"`

		ToUnit Munit `json:"toUnit"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.AuditTrail = dataAO0.AuditTrail

	m.FromUnit = dataAO0.FromUnit

	m.ToUnit = dataAO0.ToUnit

	// AO1
	var aO1 ConversionSpec
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ConversionSpec = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Conversion) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		AuditTrail *Audit `json:"auditTrail,omitempty"`

		FromUnit Munit `json:"fromUnit"`

		ToUnit Munit `json:"toUnit"`
	}

	dataAO0.AuditTrail = m.AuditTrail

	dataAO0.FromUnit = m.FromUnit

	dataAO0.ToUnit = m.ToUnit

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(m.ConversionSpec)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this conversion
func (m *Conversion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditTrail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToUnit(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with ConversionSpec
	if err := m.ConversionSpec.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Conversion) validateAuditTrail(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditTrail) { // not required
		return nil
	}

	if m.AuditTrail != nil {
		if err := m.AuditTrail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auditTrail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auditTrail")
			}
			return err
		}
	}

	return nil
}

func (m *Conversion) validateFromUnit(formats strfmt.Registry) error {

	if err := m.FromUnit.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fromUnit")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fromUnit")
		}
		return err
	}

	return nil
}

func (m *Conversion) validateToUnit(formats strfmt.Registry) error {

	if err := m.ToUnit.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("toUnit")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("toUnit")
		}
		return err
	}

	return nil
}

// ContextValidate validate this conversion based on the context it is used
func (m *Conversion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditTrail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFromUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateToUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with ConversionSpec
	if err := m.ConversionSpec.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Conversion) contextValidateAuditTrail(ctx context.Context, formats strfmt.Registry) error {

	if m.AuditTrail != nil {
		if err := m.AuditTrail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auditTrail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auditTrail")
			}
			return err
		}
	}

	return nil
}

func (m *Conversion) contextValidateFromUnit(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FromUnit.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fromUnit")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("fromUnit")
		}
		return err
	}

	return nil
}

func (m *Conversion) contextValidateToUnit(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ToUnit.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("toUnit")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("toUnit")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Conversion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Conversion) UnmarshalBinary(b []byte) error {
	var res Conversion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}