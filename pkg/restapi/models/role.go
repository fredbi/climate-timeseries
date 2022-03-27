// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Role Roles qualify how owners contribute to this or that piece of data.
//
//
// swagger:model role
type Role struct {
	auditTrailField *Audit

	descriptionLongField Translation

	descriptionShortField Translation

	idField int64

	titleField Translation

	Extra
}

// AuditTrail gets the audit trail of this subtype
func (m *Role) AuditTrail() *Audit {
	return m.auditTrailField
}

// SetAuditTrail sets the audit trail of this subtype
func (m *Role) SetAuditTrail(val *Audit) {
	m.auditTrailField = val
}

// DescriptionLong gets the description long of this subtype
func (m *Role) DescriptionLong() Translation {
	return m.descriptionLongField
}

// SetDescriptionLong sets the description long of this subtype
func (m *Role) SetDescriptionLong(val Translation) {
	m.descriptionLongField = val
}

// DescriptionShort gets the description short of this subtype
func (m *Role) DescriptionShort() Translation {
	return m.descriptionShortField
}

// SetDescriptionShort sets the description short of this subtype
func (m *Role) SetDescriptionShort(val Translation) {
	m.descriptionShortField = val
}

// ID gets the id of this subtype
func (m *Role) ID() int64 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *Role) SetID(val int64) {
	m.idField = val
}

// ShortCode gets the short code of this subtype
func (m *Role) ShortCode() ClassNomenclatureName {
	return "role"
}

// SetShortCode sets the short code of this subtype
func (m *Role) SetShortCode(val ClassNomenclatureName) {
}

// Title gets the title of this subtype
func (m *Role) Title() Translation {
	return m.titleField
}

// SetTitle sets the title of this subtype
func (m *Role) SetTitle(val Translation) {
	m.titleField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *Role) UnmarshalJSON(raw []byte) error {
	var data struct {
		Extra
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		AuditTrail *Audit `json:"auditTrail,omitempty"`

		DescriptionLong Translation `json:"descriptionLong,omitempty"`

		DescriptionShort Translation `json:"descriptionShort,omitempty"`

		ID int64 `json:"id"`

		ShortCode ClassNomenclatureName `json:"shortCode"`

		Title Translation `json:"title"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result Role

	result.auditTrailField = base.AuditTrail

	result.descriptionLongField = base.DescriptionLong

	result.descriptionShortField = base.DescriptionShort

	result.idField = base.ID

	if base.ShortCode != result.ShortCode() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid shortCode value: %q", base.ShortCode)
	}
	result.titleField = base.Title

	result.Extra = data.Extra

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m Role) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		Extra
	}{

		Extra: m.Extra,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		AuditTrail *Audit `json:"auditTrail,omitempty"`

		DescriptionLong Translation `json:"descriptionLong,omitempty"`

		DescriptionShort Translation `json:"descriptionShort,omitempty"`

		ID int64 `json:"id"`

		ShortCode ClassNomenclatureName `json:"shortCode"`

		Title Translation `json:"title"`
	}{

		AuditTrail: m.AuditTrail(),

		DescriptionLong: m.DescriptionLong(),

		DescriptionShort: m.DescriptionShort(),

		ID: m.ID(),

		ShortCode: m.ShortCode(),

		Title: m.Title(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this role
func (m *Role) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuditTrail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescriptionLong(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescriptionShort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with Extra

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Role) validateAuditTrail(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditTrail()) { // not required
		return nil
	}

	if m.AuditTrail() != nil {
		if err := m.AuditTrail().Validate(formats); err != nil {
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

func (m *Role) validateDescriptionLong(formats strfmt.Registry) error {

	if swag.IsZero(m.DescriptionLong()) { // not required
		return nil
	}

	if m.DescriptionLong() != nil {
		if err := m.DescriptionLong().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("descriptionLong")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("descriptionLong")
			}
			return err
		}
	}

	return nil
}

func (m *Role) validateDescriptionShort(formats strfmt.Registry) error {

	if swag.IsZero(m.DescriptionShort()) { // not required
		return nil
	}

	if m.DescriptionShort() != nil {
		if err := m.DescriptionShort().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("descriptionShort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("descriptionShort")
			}
			return err
		}
	}

	return nil
}

func (m *Role) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *Role) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title()); err != nil {
		return err
	}

	if m.Title() != nil {
		if err := m.Title().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("title")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("title")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this role based on the context it is used
func (m *Role) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditTrail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescriptionLong(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescriptionShort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTitle(ctx, formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with Extra

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Role) contextValidateAuditTrail(ctx context.Context, formats strfmt.Registry) error {

	if m.AuditTrail() != nil {
		if err := m.AuditTrail().ContextValidate(ctx, formats); err != nil {
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

func (m *Role) contextValidateDescriptionLong(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DescriptionLong().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("descriptionLong")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("descriptionLong")
		}
		return err
	}

	return nil
}

func (m *Role) contextValidateDescriptionShort(ctx context.Context, formats strfmt.Registry) error {

	if err := m.DescriptionShort().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("descriptionShort")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("descriptionShort")
		}
		return err
	}

	return nil
}

func (m *Role) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *Role) contextValidateShortCode(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ShortCode().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("shortCode")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("shortCode")
		}
		return err
	}

	return nil
}

func (m *Role) contextValidateTitle(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Title().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("title")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("title")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Role) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Role) UnmarshalBinary(b []byte) error {
	var res Role
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
