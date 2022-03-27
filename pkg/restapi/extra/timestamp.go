package extra

import (
	"context"

	"github.com/go-openapi/strfmt"
)

type Timestamp struct {
}

func (m Timestamp) Validate(formats strfmt.Registry) error {
	return nil
}

func (m Timestamp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

func (m Timestamp) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (m *Timestamp) UnmarshalJSON([]byte) error {
	return nil
}
