// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/fredbi/climate-timeseries/pkg/restapi/handlers/operations/classes"
	"github.com/fredbi/climate-timeseries/pkg/restapi/handlers/operations/series"
	"github.com/fredbi/climate-timeseries/pkg/restapi/handlers/operations/tags"
)

// NewClimateChangeAPIAPI creates a new ClimateChangeAPI instance
func NewClimateChangeAPIAPI(spec *loads.Document) *ClimateChangeAPIAPI {
	return &ClimateChangeAPIAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		CsvConsumer: runtime.ConsumerFunc(func(r io.Reader, target interface{}) error {
			return errors.NotImplemented("csv consumer has not yet been implemented")
		}),
		JSONConsumer: runtime.JSONConsumer(),

		CsvProducer: runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
			return errors.NotImplemented("csv producer has not yet been implemented")
		}),
		JSONProducer: runtime.JSONProducer(),

		SeriesDeleteSeriesSeriesIDHandler: series.DeleteSeriesSeriesIDHandlerFunc(func(params series.DeleteSeriesSeriesIDParams) series.DeleteSeriesSeriesIDResponder {
			return series.DeleteSeriesSeriesIDNotImplemented()
		}),
		SeriesDeleteSeriesVersionsVersionedSeriesIDHandler: series.DeleteSeriesVersionsVersionedSeriesIDHandlerFunc(func(params series.DeleteSeriesVersionsVersionedSeriesIDParams) series.DeleteSeriesVersionsVersionedSeriesIDResponder {
			return series.DeleteSeriesVersionsVersionedSeriesIDNotImplemented()
		}),
		SeriesDeleteSeriesVersionsVersionedSeriesIDValuesHandler: series.DeleteSeriesVersionsVersionedSeriesIDValuesHandlerFunc(func(params series.DeleteSeriesVersionsVersionedSeriesIDValuesParams) series.DeleteSeriesVersionsVersionedSeriesIDValuesResponder {
			return series.DeleteSeriesVersionsVersionedSeriesIDValuesNotImplemented()
		}),
		ClassesGetClassesHandler: classes.GetClassesHandlerFunc(func(params classes.GetClassesParams) classes.GetClassesResponder {
			return classes.GetClassesNotImplemented()
		}),
		ClassesGetClassesClassIDHandler: classes.GetClassesClassIDHandlerFunc(func(params classes.GetClassesClassIDParams) classes.GetClassesClassIDResponder {
			return classes.GetClassesClassIDNotImplemented()
		}),
		ClassesGetClassesClassIDMembersHandler: classes.GetClassesClassIDMembersHandlerFunc(func(params classes.GetClassesClassIDMembersParams) classes.GetClassesClassIDMembersResponder {
			return classes.GetClassesClassIDMembersNotImplemented()
		}),
		TagsGetSearchTagsHandler: tags.GetSearchTagsHandlerFunc(func(params tags.GetSearchTagsParams) tags.GetSearchTagsResponder {
			return tags.GetSearchTagsNotImplemented()
		}),
		TagsGetSearchTagsTagHandler: tags.GetSearchTagsTagHandlerFunc(func(params tags.GetSearchTagsTagParams) tags.GetSearchTagsTagResponder {
			return tags.GetSearchTagsTagNotImplemented()
		}),
		SeriesGetSeriesHandler: series.GetSeriesHandlerFunc(func(params series.GetSeriesParams) series.GetSeriesResponder {
			return series.GetSeriesNotImplemented()
		}),
		SeriesGetSeriesSeriesIDHandler: series.GetSeriesSeriesIDHandlerFunc(func(params series.GetSeriesSeriesIDParams) series.GetSeriesSeriesIDResponder {
			return series.GetSeriesSeriesIDNotImplemented()
		}),
		SeriesGetSeriesSeriesIDLatestHandler: series.GetSeriesSeriesIDLatestHandlerFunc(func(params series.GetSeriesSeriesIDLatestParams) series.GetSeriesSeriesIDLatestResponder {
			return series.GetSeriesSeriesIDLatestNotImplemented()
		}),
		SeriesGetSeriesSeriesIDLatestValuesHandler: series.GetSeriesSeriesIDLatestValuesHandlerFunc(func(params series.GetSeriesSeriesIDLatestValuesParams) series.GetSeriesSeriesIDLatestValuesResponder {
			return series.GetSeriesSeriesIDLatestValuesNotImplemented()
		}),
		SeriesGetSeriesSeriesIDSemverHandler: series.GetSeriesSeriesIDSemverHandlerFunc(func(params series.GetSeriesSeriesIDSemverParams) series.GetSeriesSeriesIDSemverResponder {
			return series.GetSeriesSeriesIDSemverNotImplemented()
		}),
		SeriesGetSeriesSeriesIDSemverSemverHandler: series.GetSeriesSeriesIDSemverSemverHandlerFunc(func(params series.GetSeriesSeriesIDSemverSemverParams) series.GetSeriesSeriesIDSemverSemverResponder {
			return series.GetSeriesSeriesIDSemverSemverNotImplemented()
		}),
		SeriesGetSeriesSeriesIDSemverSemverValuesHandler: series.GetSeriesSeriesIDSemverSemverValuesHandlerFunc(func(params series.GetSeriesSeriesIDSemverSemverValuesParams) series.GetSeriesSeriesIDSemverSemverValuesResponder {
			return series.GetSeriesSeriesIDSemverSemverValuesNotImplemented()
		}),
		SeriesGetSeriesVersionsVersionedSeriesIDHandler: series.GetSeriesVersionsVersionedSeriesIDHandlerFunc(func(params series.GetSeriesVersionsVersionedSeriesIDParams) series.GetSeriesVersionsVersionedSeriesIDResponder {
			return series.GetSeriesVersionsVersionedSeriesIDNotImplemented()
		}),
		SeriesGetSeriesVersionsVersionedSeriesIDValuesHandler: series.GetSeriesVersionsVersionedSeriesIDValuesHandlerFunc(func(params series.GetSeriesVersionsVersionedSeriesIDValuesParams) series.GetSeriesVersionsVersionedSeriesIDValuesResponder {
			return series.GetSeriesVersionsVersionedSeriesIDValuesNotImplemented()
		}),
		SeriesPostSeriesHandler: series.PostSeriesHandlerFunc(func(params series.PostSeriesParams) series.PostSeriesResponder {
			return series.PostSeriesNotImplemented()
		}),
		SeriesPostSeriesSeriesIDHandler: series.PostSeriesSeriesIDHandlerFunc(func(params series.PostSeriesSeriesIDParams) series.PostSeriesSeriesIDResponder {
			return series.PostSeriesSeriesIDNotImplemented()
		}),
		SeriesPostSeriesVersionsVersionedSeriesIDHandler: series.PostSeriesVersionsVersionedSeriesIDHandlerFunc(func(params series.PostSeriesVersionsVersionedSeriesIDParams) series.PostSeriesVersionsVersionedSeriesIDResponder {
			return series.PostSeriesVersionsVersionedSeriesIDNotImplemented()
		}),
		SeriesPostSeriesVersionsVersionedSeriesIDValuesHandler: series.PostSeriesVersionsVersionedSeriesIDValuesHandlerFunc(func(params series.PostSeriesVersionsVersionedSeriesIDValuesParams) series.PostSeriesVersionsVersionedSeriesIDValuesResponder {
			return series.PostSeriesVersionsVersionedSeriesIDValuesNotImplemented()
		}),
		SeriesPutSeriesSeriesIDHandler: series.PutSeriesSeriesIDHandlerFunc(func(params series.PutSeriesSeriesIDParams) series.PutSeriesSeriesIDResponder {
			return series.PutSeriesSeriesIDNotImplemented()
		}),
		SeriesPutSeriesVersionsVersionedSeriesIDHandler: series.PutSeriesVersionsVersionedSeriesIDHandlerFunc(func(params series.PutSeriesVersionsVersionedSeriesIDParams) series.PutSeriesVersionsVersionedSeriesIDResponder {
			return series.PutSeriesVersionsVersionedSeriesIDNotImplemented()
		}),
		SeriesPutSeriesVersionsVersionedSeriesIDValuesHandler: series.PutSeriesVersionsVersionedSeriesIDValuesHandlerFunc(func(params series.PutSeriesVersionsVersionedSeriesIDValuesParams) series.PutSeriesVersionsVersionedSeriesIDValuesResponder {
			return series.PutSeriesVersionsVersionedSeriesIDValuesNotImplemented()
		}),
	}
}

/*ClimateChangeAPIAPI The timeseries API allows contributors to upload time series about their climate change models and studies.

The API allows the public to search and consult time series about climate change research.
*/
type ClimateChangeAPIAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// CsvConsumer registers a consumer for the following mime types:
	//   - text/csv
	CsvConsumer runtime.Consumer
	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// CsvProducer registers a producer for the following mime types:
	//   - text/csv
	CsvProducer runtime.Producer
	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// SeriesDeleteSeriesSeriesIDHandler sets the operation handler for the delete series series ID operation
	SeriesDeleteSeriesSeriesIDHandler series.DeleteSeriesSeriesIDHandler
	// SeriesDeleteSeriesVersionsVersionedSeriesIDHandler sets the operation handler for the delete series versions versioned series ID operation
	SeriesDeleteSeriesVersionsVersionedSeriesIDHandler series.DeleteSeriesVersionsVersionedSeriesIDHandler
	// SeriesDeleteSeriesVersionsVersionedSeriesIDValuesHandler sets the operation handler for the delete series versions versioned series ID values operation
	SeriesDeleteSeriesVersionsVersionedSeriesIDValuesHandler series.DeleteSeriesVersionsVersionedSeriesIDValuesHandler
	// ClassesGetClassesHandler sets the operation handler for the get classes operation
	ClassesGetClassesHandler classes.GetClassesHandler
	// ClassesGetClassesClassIDHandler sets the operation handler for the get classes class ID operation
	ClassesGetClassesClassIDHandler classes.GetClassesClassIDHandler
	// ClassesGetClassesClassIDMembersHandler sets the operation handler for the get classes class ID members operation
	ClassesGetClassesClassIDMembersHandler classes.GetClassesClassIDMembersHandler
	// TagsGetSearchTagsHandler sets the operation handler for the get search tags operation
	TagsGetSearchTagsHandler tags.GetSearchTagsHandler
	// TagsGetSearchTagsTagHandler sets the operation handler for the get search tags tag operation
	TagsGetSearchTagsTagHandler tags.GetSearchTagsTagHandler
	// SeriesGetSeriesHandler sets the operation handler for the get series operation
	SeriesGetSeriesHandler series.GetSeriesHandler
	// SeriesGetSeriesSeriesIDHandler sets the operation handler for the get series series ID operation
	SeriesGetSeriesSeriesIDHandler series.GetSeriesSeriesIDHandler
	// SeriesGetSeriesSeriesIDLatestHandler sets the operation handler for the get series series ID latest operation
	SeriesGetSeriesSeriesIDLatestHandler series.GetSeriesSeriesIDLatestHandler
	// SeriesGetSeriesSeriesIDLatestValuesHandler sets the operation handler for the get series series ID latest values operation
	SeriesGetSeriesSeriesIDLatestValuesHandler series.GetSeriesSeriesIDLatestValuesHandler
	// SeriesGetSeriesSeriesIDSemverHandler sets the operation handler for the get series series ID semver operation
	SeriesGetSeriesSeriesIDSemverHandler series.GetSeriesSeriesIDSemverHandler
	// SeriesGetSeriesSeriesIDSemverSemverHandler sets the operation handler for the get series series ID semver semver operation
	SeriesGetSeriesSeriesIDSemverSemverHandler series.GetSeriesSeriesIDSemverSemverHandler
	// SeriesGetSeriesSeriesIDSemverSemverValuesHandler sets the operation handler for the get series series ID semver semver values operation
	SeriesGetSeriesSeriesIDSemverSemverValuesHandler series.GetSeriesSeriesIDSemverSemverValuesHandler
	// SeriesGetSeriesVersionsVersionedSeriesIDHandler sets the operation handler for the get series versions versioned series ID operation
	SeriesGetSeriesVersionsVersionedSeriesIDHandler series.GetSeriesVersionsVersionedSeriesIDHandler
	// SeriesGetSeriesVersionsVersionedSeriesIDValuesHandler sets the operation handler for the get series versions versioned series ID values operation
	SeriesGetSeriesVersionsVersionedSeriesIDValuesHandler series.GetSeriesVersionsVersionedSeriesIDValuesHandler
	// SeriesPostSeriesHandler sets the operation handler for the post series operation
	SeriesPostSeriesHandler series.PostSeriesHandler
	// SeriesPostSeriesSeriesIDHandler sets the operation handler for the post series series ID operation
	SeriesPostSeriesSeriesIDHandler series.PostSeriesSeriesIDHandler
	// SeriesPostSeriesVersionsVersionedSeriesIDHandler sets the operation handler for the post series versions versioned series ID operation
	SeriesPostSeriesVersionsVersionedSeriesIDHandler series.PostSeriesVersionsVersionedSeriesIDHandler
	// SeriesPostSeriesVersionsVersionedSeriesIDValuesHandler sets the operation handler for the post series versions versioned series ID values operation
	SeriesPostSeriesVersionsVersionedSeriesIDValuesHandler series.PostSeriesVersionsVersionedSeriesIDValuesHandler
	// SeriesPutSeriesSeriesIDHandler sets the operation handler for the put series series ID operation
	SeriesPutSeriesSeriesIDHandler series.PutSeriesSeriesIDHandler
	// SeriesPutSeriesVersionsVersionedSeriesIDHandler sets the operation handler for the put series versions versioned series ID operation
	SeriesPutSeriesVersionsVersionedSeriesIDHandler series.PutSeriesVersionsVersionedSeriesIDHandler
	// SeriesPutSeriesVersionsVersionedSeriesIDValuesHandler sets the operation handler for the put series versions versioned series ID values operation
	SeriesPutSeriesVersionsVersionedSeriesIDValuesHandler series.PutSeriesVersionsVersionedSeriesIDValuesHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *ClimateChangeAPIAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *ClimateChangeAPIAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *ClimateChangeAPIAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ClimateChangeAPIAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ClimateChangeAPIAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ClimateChangeAPIAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ClimateChangeAPIAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ClimateChangeAPIAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ClimateChangeAPIAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ClimateChangeAPIAPI
func (o *ClimateChangeAPIAPI) Validate() error {
	var unregistered []string

	if o.CsvConsumer == nil {
		unregistered = append(unregistered, "CsvConsumer")
	}
	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.CsvProducer == nil {
		unregistered = append(unregistered, "CsvProducer")
	}
	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.SeriesDeleteSeriesSeriesIDHandler == nil {
		unregistered = append(unregistered, "series.DeleteSeriesSeriesIDHandler")
	}
	if o.SeriesDeleteSeriesVersionsVersionedSeriesIDHandler == nil {
		unregistered = append(unregistered, "series.DeleteSeriesVersionsVersionedSeriesIDHandler")
	}
	if o.SeriesDeleteSeriesVersionsVersionedSeriesIDValuesHandler == nil {
		unregistered = append(unregistered, "series.DeleteSeriesVersionsVersionedSeriesIDValuesHandler")
	}
	if o.ClassesGetClassesHandler == nil {
		unregistered = append(unregistered, "classes.GetClassesHandler")
	}
	if o.ClassesGetClassesClassIDHandler == nil {
		unregistered = append(unregistered, "classes.GetClassesClassIDHandler")
	}
	if o.ClassesGetClassesClassIDMembersHandler == nil {
		unregistered = append(unregistered, "classes.GetClassesClassIDMembersHandler")
	}
	if o.TagsGetSearchTagsHandler == nil {
		unregistered = append(unregistered, "tags.GetSearchTagsHandler")
	}
	if o.TagsGetSearchTagsTagHandler == nil {
		unregistered = append(unregistered, "tags.GetSearchTagsTagHandler")
	}
	if o.SeriesGetSeriesHandler == nil {
		unregistered = append(unregistered, "series.GetSeriesHandler")
	}
	if o.SeriesGetSeriesSeriesIDHandler == nil {
		unregistered = append(unregistered, "series.GetSeriesSeriesIDHandler")
	}
	if o.SeriesGetSeriesSeriesIDLatestHandler == nil {
		unregistered = append(unregistered, "series.GetSeriesSeriesIDLatestHandler")
	}
	if o.SeriesGetSeriesSeriesIDLatestValuesHandler == nil {
		unregistered = append(unregistered, "series.GetSeriesSeriesIDLatestValuesHandler")
	}
	if o.SeriesGetSeriesSeriesIDSemverHandler == nil {
		unregistered = append(unregistered, "series.GetSeriesSeriesIDSemverHandler")
	}
	if o.SeriesGetSeriesSeriesIDSemverSemverHandler == nil {
		unregistered = append(unregistered, "series.GetSeriesSeriesIDSemverSemverHandler")
	}
	if o.SeriesGetSeriesSeriesIDSemverSemverValuesHandler == nil {
		unregistered = append(unregistered, "series.GetSeriesSeriesIDSemverSemverValuesHandler")
	}
	if o.SeriesGetSeriesVersionsVersionedSeriesIDHandler == nil {
		unregistered = append(unregistered, "series.GetSeriesVersionsVersionedSeriesIDHandler")
	}
	if o.SeriesGetSeriesVersionsVersionedSeriesIDValuesHandler == nil {
		unregistered = append(unregistered, "series.GetSeriesVersionsVersionedSeriesIDValuesHandler")
	}
	if o.SeriesPostSeriesHandler == nil {
		unregistered = append(unregistered, "series.PostSeriesHandler")
	}
	if o.SeriesPostSeriesSeriesIDHandler == nil {
		unregistered = append(unregistered, "series.PostSeriesSeriesIDHandler")
	}
	if o.SeriesPostSeriesVersionsVersionedSeriesIDHandler == nil {
		unregistered = append(unregistered, "series.PostSeriesVersionsVersionedSeriesIDHandler")
	}
	if o.SeriesPostSeriesVersionsVersionedSeriesIDValuesHandler == nil {
		unregistered = append(unregistered, "series.PostSeriesVersionsVersionedSeriesIDValuesHandler")
	}
	if o.SeriesPutSeriesSeriesIDHandler == nil {
		unregistered = append(unregistered, "series.PutSeriesSeriesIDHandler")
	}
	if o.SeriesPutSeriesVersionsVersionedSeriesIDHandler == nil {
		unregistered = append(unregistered, "series.PutSeriesVersionsVersionedSeriesIDHandler")
	}
	if o.SeriesPutSeriesVersionsVersionedSeriesIDValuesHandler == nil {
		unregistered = append(unregistered, "series.PutSeriesVersionsVersionedSeriesIDValuesHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ClimateChangeAPIAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ClimateChangeAPIAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *ClimateChangeAPIAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *ClimateChangeAPIAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "text/csv":
			result["text/csv"] = o.CsvConsumer
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *ClimateChangeAPIAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "text/csv":
			result["text/csv"] = o.CsvProducer
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ClimateChangeAPIAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the climate change API API
func (o *ClimateChangeAPIAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ClimateChangeAPIAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/series/{seriesId}"] = series.NewDeleteSeriesSeriesID(o.context, o.SeriesDeleteSeriesSeriesIDHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/series/versions/{versionedSeriesId}"] = series.NewDeleteSeriesVersionsVersionedSeriesID(o.context, o.SeriesDeleteSeriesVersionsVersionedSeriesIDHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/series/versions/{versionedSeriesId}/values"] = series.NewDeleteSeriesVersionsVersionedSeriesIDValues(o.context, o.SeriesDeleteSeriesVersionsVersionedSeriesIDValuesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/classes"] = classes.NewGetClasses(o.context, o.ClassesGetClassesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/classes/{classId}"] = classes.NewGetClassesClassID(o.context, o.ClassesGetClassesClassIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/classes/{classId}/members"] = classes.NewGetClassesClassIDMembers(o.context, o.ClassesGetClassesClassIDMembersHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/search/tags"] = tags.NewGetSearchTags(o.context, o.TagsGetSearchTagsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/search/tags/{tag}"] = tags.NewGetSearchTagsTag(o.context, o.TagsGetSearchTagsTagHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/series"] = series.NewGetSeries(o.context, o.SeriesGetSeriesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/series/{seriesId}"] = series.NewGetSeriesSeriesID(o.context, o.SeriesGetSeriesSeriesIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/series/{seriesId}/latest"] = series.NewGetSeriesSeriesIDLatest(o.context, o.SeriesGetSeriesSeriesIDLatestHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/series/{seriesId}/latest/values"] = series.NewGetSeriesSeriesIDLatestValues(o.context, o.SeriesGetSeriesSeriesIDLatestValuesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/series/{seriesId}/semver"] = series.NewGetSeriesSeriesIDSemver(o.context, o.SeriesGetSeriesSeriesIDSemverHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/series/{seriesId}/semver/{semver}"] = series.NewGetSeriesSeriesIDSemverSemver(o.context, o.SeriesGetSeriesSeriesIDSemverSemverHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/series/{seriesId}/semver/{semver}/values"] = series.NewGetSeriesSeriesIDSemverSemverValues(o.context, o.SeriesGetSeriesSeriesIDSemverSemverValuesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/series/versions/{versionedSeriesId}"] = series.NewGetSeriesVersionsVersionedSeriesID(o.context, o.SeriesGetSeriesVersionsVersionedSeriesIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/series/versions/{versionedSeriesId}/values"] = series.NewGetSeriesVersionsVersionedSeriesIDValues(o.context, o.SeriesGetSeriesVersionsVersionedSeriesIDValuesHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/series"] = series.NewPostSeries(o.context, o.SeriesPostSeriesHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/series/{seriesId}"] = series.NewPostSeriesSeriesID(o.context, o.SeriesPostSeriesSeriesIDHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/series/versions/{versionedSeriesId}"] = series.NewPostSeriesVersionsVersionedSeriesID(o.context, o.SeriesPostSeriesVersionsVersionedSeriesIDHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/series/versions/{versionedSeriesId}/values"] = series.NewPostSeriesVersionsVersionedSeriesIDValues(o.context, o.SeriesPostSeriesVersionsVersionedSeriesIDValuesHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/series/{seriesId}"] = series.NewPutSeriesSeriesID(o.context, o.SeriesPutSeriesSeriesIDHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/series/versions/{versionedSeriesId}"] = series.NewPutSeriesVersionsVersionedSeriesID(o.context, o.SeriesPutSeriesVersionsVersionedSeriesIDHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/series/versions/{versionedSeriesId}/values"] = series.NewPutSeriesVersionsVersionedSeriesIDValues(o.context, o.SeriesPutSeriesVersionsVersionedSeriesIDValuesHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ClimateChangeAPIAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *ClimateChangeAPIAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *ClimateChangeAPIAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *ClimateChangeAPIAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *ClimateChangeAPIAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}