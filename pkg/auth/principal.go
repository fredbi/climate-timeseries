package auth

// Principal identifies the connected user.
type Principal interface {
	Username() string
	Email() string
	Roles() []Role
}

func NewAPIPrincipal() *APIPrincipal {
	return &APIPrincipal{}
}

type (
	APIPrincipal struct {
	}

	Role string
)

const (
	RoleNone        Role = "none"
	RoleAdmin       Role = "admin"
	RoleContributor Role = "contributor"
)

var _ Principal = &APIPrincipal{}

func (p *APIPrincipal) Username() string {
	return ""
}

func (p *APIPrincipal) Email() string {
	return ""
}

func (p *APIPrincipal) Roles() []Role {
	return nil
}
