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

// Constant Mathematical and physical constants.
//
//
// swagger:model constant
type Constant struct {
	auditTrailField *Audit

	descriptionLongField Translation

	descriptionShortField Translation

	idField int64

	titleField Translation

	// constant has measurement domains
	ConstantHasMeasurementDomains []Mdomain `json:"constantHasMeasurementDomains"`

	// measurement unit
	MeasurementUnit Munit `json:"measurementUnit,omitempty"`

	// metadata
	// Required: true
	Metadata *Mumeta `json:"metadata"`

	// symbol
	// Required: true
	// Max Length: 10
	// Min Length: 1
	Symbol *string `json:"symbol"`

	// value
	// Required: true
	Value *float32 `json:"value"`
}

// AuditTrail gets the audit trail of this subtype
func (m *Constant) AuditTrail() *Audit {
	return m.auditTrailField
}

// SetAuditTrail sets the audit trail of this subtype
func (m *Constant) SetAuditTrail(val *Audit) {
	m.auditTrailField = val
}

// DescriptionLong gets the description long of this subtype
func (m *Constant) DescriptionLong() Translation {
	return m.descriptionLongField
}

// SetDescriptionLong sets the description long of this subtype
func (m *Constant) SetDescriptionLong(val Translation) {
	m.descriptionLongField = val
}

// DescriptionShort gets the description short of this subtype
func (m *Constant) DescriptionShort() Translation {
	return m.descriptionShortField
}

// SetDescriptionShort sets the description short of this subtype
func (m *Constant) SetDescriptionShort(val Translation) {
	m.descriptionShortField = val
}

// ID gets the id of this subtype
func (m *Constant) ID() int64 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *Constant) SetID(val int64) {
	m.idField = val
}

// ShortCode gets the short code of this subtype
func (m *Constant) ShortCode() ClassName {
	return "constant"
}

// SetShortCode sets the short code of this subtype
func (m *Constant) SetShortCode(val ClassName) {
}

// Title gets the title of this subtype
func (m *Constant) Title() Translation {
	return m.titleField
}

// SetTitle sets the title of this subtype
func (m *Constant) SetTitle(val Translation) {
	m.titleField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *Constant) UnmarshalJSON(raw []byte) error {
	var data struct {

		// constant has measurement domains
		ConstantHasMeasurementDomains []Mdomain `json:"constantHasMeasurementDomains"`

		// measurement unit
		MeasurementUnit Munit `json:"measurementUnit,omitempty"`

		// metadata
		// Required: true
		Metadata *Mumeta `json:"metadata"`

		// symbol
		// Required: true
		// Max Length: 10
		// Min Length: 1
		Symbol *string `json:"symbol"`

		// value
		// Required: true
		Value *float32 `json:"value"`
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

	var result Constant

	result.auditTrailField = base.AuditTrail

	result.descriptionLongField = base.DescriptionLong

	result.descriptionShortField = base.DescriptionShort

	result.idField = base.ID

	if base.ShortCode != result.ShortCode() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid shortCode value: %q", base.ShortCode)
	}
	result.titleField = base.Title

	result.ConstantHasMeasurementDomains = data.ConstantHasMeasurementDomains
	result.MeasurementUnit = data.MeasurementUnit
	result.Metadata = data.Metadata
	result.Symbol = data.Symbol
	result.Value = data.Value

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m Constant) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// constant has measurement domains
		ConstantHasMeasurementDomains []Mdomain `json:"constantHasMeasurementDomains"`

		// measurement unit
		MeasurementUnit Munit `json:"measurementUnit,omitempty"`

		// metadata
		// Required: true
		Metadata *Mumeta `json:"metadata"`

		// symbol
		// Required: true
		// Max Length: 10
		// Min Length: 1
		Symbol *string `json:"symbol"`

		// value
		// Required: true
		Value *float32 `json:"value"`
	}{

		ConstantHasMeasurementDomains: m.ConstantHasMeasurementDomains,

		MeasurementUnit: m.MeasurementUnit,

		Metadata: m.Metadata,

		Symbol: m.Symbol,

		Value: m.Value,
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

// Validate validates this constant
func (m *Constant) Validate(formats strfmt.Registry) error {
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

	if err := m.validateConstantHasMeasurementDomains(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeasurementUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSymbol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Constant) validateAuditTrail(formats strfmt.Registry) error {

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

func (m *Constant) validateDescriptionLong(formats strfmt.Registry) error {

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

func (m *Constant) validateDescriptionShort(formats strfmt.Registry) error {

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

func (m *Constant) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *Constant) validateTitle(formats strfmt.Registry) error {

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

func (m *Constant) validateConstantHasMeasurementDomains(formats strfmt.Registry) error {

	if swag.IsZero(m.ConstantHasMeasurementDomains) { // not required
		return nil
	}

	for i := 0; i < len(m.ConstantHasMeasurementDomains); i++ {

		if err := m.ConstantHasMeasurementDomains[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("constantHasMeasurementDomains" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("constantHasMeasurementDomains" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Constant) validateMeasurementUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.MeasurementUnit) { // not required
		return nil
	}

	if err := m.MeasurementUnit.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("measurementUnit")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("measurementUnit")
		}
		return err
	}

	return nil
}

func (m *Constant) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *Constant) validateSymbol(formats strfmt.Registry) error {

	if err := validate.Required("symbol", "body", m.Symbol); err != nil {
		return err
	}

	if err := validate.MinLength("symbol", "body", *m.Symbol, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("symbol", "body", *m.Symbol, 10); err != nil {
		return err
	}

	return nil
}

func (m *Constant) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this constant based on the context it is used
func (m *Constant) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateConstantHasMeasurementDomains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMeasurementUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Constant) contextValidateAuditTrail(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Constant) contextValidateDescriptionLong(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Constant) contextValidateDescriptionShort(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Constant) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *Constant) contextValidateTitle(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Constant) contextValidateConstantHasMeasurementDomains(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConstantHasMeasurementDomains); i++ {

		if err := m.ConstantHasMeasurementDomains[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("constantHasMeasurementDomains" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("constantHasMeasurementDomains" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Constant) contextValidateMeasurementUnit(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MeasurementUnit.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("measurementUnit")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("measurementUnit")
		}
		return err
	}

	return nil
}

func (m *Constant) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Constant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Constant) UnmarshalBinary(b []byte) error {
	var res Constant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
