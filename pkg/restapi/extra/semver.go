package extra

import (
	"context"

	"github.com/go-openapi/strfmt"
)

type Semver struct {
}

func (m Semver) Validate(formats strfmt.Registry) error {
	return nil
}

func (m Semver) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

func (m Semver) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (m *Semver) UnmarshalJSON([]byte) error {
	return nil
}
