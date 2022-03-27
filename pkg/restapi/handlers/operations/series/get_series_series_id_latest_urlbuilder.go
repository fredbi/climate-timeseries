// Code generated by go-swagger; DO NOT EDIT.

package series

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetSeriesSeriesIDLatestURL generates an URL for the get series series ID latest operation
type GetSeriesSeriesIDLatestURL struct {
	SeriesID int64

	Audit  *bool
	Brief  *bool
	Deep   *bool
	Status *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetSeriesSeriesIDLatestURL) WithBasePath(bp string) *GetSeriesSeriesIDLatestURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetSeriesSeriesIDLatestURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetSeriesSeriesIDLatestURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/series/{seriesId}/latest"

	seriesID := swag.FormatInt64(o.SeriesID)
	if seriesID != "" {
		_path = strings.Replace(_path, "{seriesId}", seriesID, -1)
	} else {
		return nil, errors.New("seriesId is required on GetSeriesSeriesIDLatestURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1/"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var auditQ string
	if o.Audit != nil {
		auditQ = swag.FormatBool(*o.Audit)
	}
	if auditQ != "" {
		qs.Set("audit", auditQ)
	}

	var briefQ string
	if o.Brief != nil {
		briefQ = swag.FormatBool(*o.Brief)
	}
	if briefQ != "" {
		qs.Set("brief", briefQ)
	}

	var deepQ string
	if o.Deep != nil {
		deepQ = swag.FormatBool(*o.Deep)
	}
	if deepQ != "" {
		qs.Set("deep", deepQ)
	}

	var statusQ string
	if o.Status != nil {
		statusQ = *o.Status
	}
	if statusQ != "" {
		qs.Set("status", statusQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetSeriesSeriesIDLatestURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetSeriesSeriesIDLatestURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetSeriesSeriesIDLatestURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetSeriesSeriesIDLatestURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetSeriesSeriesIDLatestURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetSeriesSeriesIDLatestURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
