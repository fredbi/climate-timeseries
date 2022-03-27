// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Theme Climate themes.
//
//
// swagger:model theme
type Theme struct {
	auditTrailField *Audit

	descriptionLongField Translation

	descriptionShortField Translation

	idField int64

	titleField Translation

	// linked documents
	LinkedDocuments Documents `json:"linkedDocuments,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// AuditTrail gets the audit trail of this subtype
func (m *Theme) AuditTrail() *Audit {
	return m.auditTrailField
}

// SetAuditTrail sets the audit trail of this subtype
func (m *Theme) SetAuditTrail(val *Audit) {
	m.auditTrailField = val
}

// DescriptionLong gets the description long of this subtype
func (m *Theme) DescriptionLong() Translation {
	return m.descriptionLongField
}

// SetDescriptionLong sets the description long of this subtype
func (m *Theme) SetDescriptionLong(val Translation) {
	m.descriptionLongField = val
}

// DescriptionShort gets the description short of this subtype
func (m *Theme) DescriptionShort() Translation {
	return m.descriptionShortField
}

// SetDescriptionShort sets the description short of this subtype
func (m *Theme) SetDescriptionShort(val Translation) {
	m.descriptionShortField = val
}

// ID gets the id of this subtype
func (m *Theme) ID() int64 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *Theme) SetID(val int64) {
	m.idField = val
}

// ShortCode gets the short code of this subtype
func (m *Theme) ShortCode() ClassName {
	return "theme"
}

// SetShortCode sets the short code of this subtype
func (m *Theme) SetShortCode(val ClassName) {
}

// Title gets the title of this subtype
func (m *Theme) Title() Translation {
	return m.titleField
}

// SetTitle sets the title of this subtype
func (m *Theme) SetTitle(val Translation) {
	m.titleField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *Theme) UnmarshalJSON(raw []byte) error {
	var data struct {

		// linked documents
		LinkedDocuments Documents `json:"linkedDocuments,omitempty"`

		// tags
		Tags []string `json:"tags"`
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

		ShortCode ClassName `json:"shortCode"`

		Title Translation `json:"title"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result Theme

	result.auditTrailField = base.AuditTrail

	result.descriptionLongField = base.DescriptionLong

	result.descriptionShortField = base.DescriptionShort

	result.idField = base.ID

	if base.ShortCode != result.ShortCode() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid shortCode value: %q", base.ShortCode)
	}
	result.titleField = base.Title

	result.LinkedDocuments = data.LinkedDocuments
	result.Tags = data.Tags

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m Theme) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// linked documents
		LinkedDocuments Documents `json:"linkedDocuments,omitempty"`

		// tags
		Tags []string `json:"tags"`
	}{

		LinkedDocuments: m.LinkedDocuments,

		Tags: m.Tags,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		AuditTrail *Audit `json:"auditTrail,omitempty"`

		DescriptionLong Translation `json:"descriptionLong,omitempty"`

		DescriptionShort Translation `json:"descriptionShort,omitempty"`

		ID int64 `json:"id"`

		ShortCode ClassName `json:"shortCode"`

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

// Validate validates this theme
func (m *Theme) Validate(formats strfmt.Registry) error {
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

	if err := m.validateLinkedDocuments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Theme) validateAuditTrail(formats strfmt.Registry) error {

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

func (m *Theme) validateDescriptionLong(formats strfmt.Registry) error {

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

func (m *Theme) validateDescriptionShort(formats strfmt.Registry) error {

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

func (m *Theme) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *Theme) validateTitle(formats strfmt.Registry) error {

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

func (m *Theme) validateLinkedDocuments(formats strfmt.Registry) error {

	if swag.IsZero(m.LinkedDocuments) { // not required
		return nil
	}

	if err := m.LinkedDocuments.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("linkedDocuments")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("linkedDocuments")
		}
		return err
	}

	return nil
}

func (m *Theme) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.MinLength("tags"+"."+strconv.Itoa(i), "body", m.Tags[i], 1); err != nil {
			return err
		}

		if err := validate.MaxLength("tags"+"."+strconv.Itoa(i), "body", m.Tags[i], 100); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validate this theme based on the context it is used
func (m *Theme) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateLinkedDocuments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Theme) contextValidateAuditTrail(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Theme) contextValidateDescriptionLong(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Theme) contextValidateDescriptionShort(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Theme) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *Theme) contextValidateTitle(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Theme) contextValidateLinkedDocuments(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LinkedDocuments.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("linkedDocuments")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("linkedDocuments")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Theme) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Theme) UnmarshalBinary(b []byte) error {
	var res Theme
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
