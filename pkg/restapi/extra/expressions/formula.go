package expressions

import (
	"context"

	"github.com/go-openapi/strfmt"
)

type Formula struct {
}

func (m Formula) Validate(formats strfmt.Registry) error {
	return nil
}

func (m Formula) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

func (m Formula) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (m *Formula) UnmarshalJSON([]byte) error {
	return nil
}
