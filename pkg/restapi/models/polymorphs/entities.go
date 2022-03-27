package polymorphs

import (
	"context"

	"github.com/fredbi/climate-timeseries/pkg/restapi/models"
	"github.com/go-openapi/strfmt"
)

type Entities []models.ClassNomenclature

func (m Entities) Validate(formats strfmt.Registry) error {
	return nil
}

func (m Entities) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

func (m Entities) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (m *Entities) UnmarshalJSON([]byte) error {
	return nil
}
