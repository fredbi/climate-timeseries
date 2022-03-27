package expressions

import (
	"context"

	"github.com/go-openapi/strfmt"
)

type DimensionsFormula struct {
}

func (m DimensionsFormula) Validate(formats strfmt.Registry) error {
	return nil
}

func (m DimensionsFormula) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

func (m DimensionsFormula) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (m *DimensionsFormula) UnmarshalJSON([]byte) error {
	return nil
}
