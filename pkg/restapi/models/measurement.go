// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"

	"github.com/fredbi/climate-timeseries/pkg/restapi/extra/expressions"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Measurement A measurement refers to any physical or economic measurement, the unit notwithstanding.
//
// Examples:
//   * speed
//   * temperature
//
// We introduce "quantity" to take into account measurement that count things, as well as "currency"
// to take into account measurements of economic facts.
//
// Measurements usually refer to one or several physical (or economic) dimensions.
//
// Example:
//   speed is homogeneous to "L^T-1"
//
// However, some measurements are without dimension, such as "QUANTITY" and "ANGLE".
//
// Series of ratios (e.g. % increase of GDP) are not considered measurements. They shall use special units which do not refer to measurements.
//
//
// swagger:model measurement
type Measurement struct {
	auditTrailField *Audit

	descriptionLongField Translation

	descriptionShortField Translation

	idField int64

	titleField Translation

	// dimensions
	Dimensions expressions.DimensionsFormula `json:"dimensions,omitempty"`

	// measurement has measurement domains
	MeasurementHasMeasurementDomains []Mdomain `json:"measurementHasMeasurementDomains"`
}

// AuditTrail gets the audit trail of this subtype
func (m *Measurement) AuditTrail() *Audit {
	return m.auditTrailField
}

// SetAuditTrail sets the audit trail of this subtype
func (m *Measurement) SetAuditTrail(val *Audit) {
	m.auditTrailField = val
}

// DescriptionLong gets the description long of this subtype
func (m *Measurement) DescriptionLong() Translation {
	return m.descriptionLongField
}

// SetDescriptionLong sets the description long of this subtype
func (m *Measurement) SetDescriptionLong(val Translation) {
	m.descriptionLongField = val
}

// DescriptionShort gets the description short of this subtype
func (m *Measurement) DescriptionShort() Translation {
	return m.descriptionShortField
}

// SetDescriptionShort sets the description short of this subtype
func (m *Measurement) SetDescriptionShort(val Translation) {
	m.descriptionShortField = val
}

// ID gets the id of this subtype
func (m *Measurement) ID() int64 {
	return m.idField
}

// SetID sets the id of this subtype
func (m *Measurement) SetID(val int64) {
	m.idField = val
}

// ShortCode gets the short code of this subtype
func (m *Measurement) ShortCode() ClassNomenclatureName {
	return "measurement"
}

// SetShortCode sets the short code of this subtype
func (m *Measurement) SetShortCode(val ClassNomenclatureName) {
}

// Title gets the title of this subtype
func (m *Measurement) Title() Translation {
	return m.titleField
}

// SetTitle sets the title of this subtype
func (m *Measurement) SetTitle(val Translation) {
	m.titleField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *Measurement) UnmarshalJSON(raw []byte) error {
	var data struct {

		// dimensions
		Dimensions expressions.DimensionsFormula `json:"dimensions,omitempty"`

		// measurement has measurement domains
		MeasurementHasMeasurementDomains []Mdomain `json:"measurementHasMeasurementDomains"`
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

	var result Measurement

	result.auditTrailField = base.AuditTrail

	result.descriptionLongField = base.DescriptionLong

	result.descriptionShortField = base.DescriptionShort

	result.idField = base.ID

	if base.ShortCode != result.ShortCode() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid shortCode value: %q", base.ShortCode)
	}
	result.titleField = base.Title

	result.Dimensions = data.Dimensions
	result.MeasurementHasMeasurementDomains = data.MeasurementHasMeasurementDomains

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m Measurement) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// dimensions
		Dimensions expressions.DimensionsFormula `json:"dimensions,omitempty"`

		// measurement has measurement domains
		MeasurementHasMeasurementDomains []Mdomain `json:"measurementHasMeasurementDomains"`
	}{

		Dimensions: m.Dimensions,

		MeasurementHasMeasurementDomains: m.MeasurementHasMeasurementDomains,
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

// Validate validates this measurement
func (m *Measurement) Validate(formats strfmt.Registry) error {
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

	if err := m.validateDimensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeasurementHasMeasurementDomains(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Measurement) validateAuditTrail(formats strfmt.Registry) error {

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

func (m *Measurement) validateDescriptionLong(formats strfmt.Registry) error {

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

func (m *Measurement) validateDescriptionShort(formats strfmt.Registry) error {

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

func (m *Measurement) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *Measurement) validateTitle(formats strfmt.Registry) error {

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

func (m *Measurement) validateDimensions(formats strfmt.Registry) error {

	if swag.IsZero(m.Dimensions) { // not required
		return nil
	}

	if err := m.Dimensions.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("dimensions")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("dimensions")
		}
		return err
	}

	return nil
}

func (m *Measurement) validateMeasurementHasMeasurementDomains(formats strfmt.Registry) error {

	if swag.IsZero(m.MeasurementHasMeasurementDomains) { // not required
		return nil
	}

	for i := 0; i < len(m.MeasurementHasMeasurementDomains); i++ {

		if err := m.MeasurementHasMeasurementDomains[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("measurementHasMeasurementDomains" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("measurementHasMeasurementDomains" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this measurement based on the context it is used
func (m *Measurement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateMeasurementHasMeasurementDomains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Measurement) contextValidateAuditTrail(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Measurement) contextValidateDescriptionLong(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Measurement) contextValidateDescriptionShort(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Measurement) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID())); err != nil {
		return err
	}

	return nil
}

func (m *Measurement) contextValidateShortCode(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Measurement) contextValidateTitle(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Measurement) contextValidateMeasurementHasMeasurementDomains(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MeasurementHasMeasurementDomains); i++ {

		if err := m.MeasurementHasMeasurementDomains[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("measurementHasMeasurementDomains" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("measurementHasMeasurementDomains" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Measurement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Measurement) UnmarshalBinary(b []byte) error {
	var res Measurement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
