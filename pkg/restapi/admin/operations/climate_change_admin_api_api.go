// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
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

	"github.com/fredbi/climate-timeseries/pkg/restapi/admin/operations/classes"
	"github.com/fredbi/climate-timeseries/pkg/restapi/admin/operations/conversions"
)

// NewClimateChangeAdminAPIAPI creates a new ClimateChangeAdminAPI instance
func NewClimateChangeAdminAPIAPI(spec *loads.Document) *ClimateChangeAdminAPIAPI {
	return &ClimateChangeAdminAPIAPI{
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

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		ClassesDeleteClassesClassIDMembersClassMemberIDHandler: classes.DeleteClassesClassIDMembersClassMemberIDHandlerFunc(func(params classes.DeleteClassesClassIDMembersClassMemberIDParams, principal interface{}) classes.DeleteClassesClassIDMembersClassMemberIDResponder {
			return classes.DeleteClassesClassIDMembersClassMemberIDNotImplemented()
		}),
		ClassesPostClassesClassIDMembersHandler: classes.PostClassesClassIDMembersHandlerFunc(func(params classes.PostClassesClassIDMembersParams, principal interface{}) classes.PostClassesClassIDMembersResponder {
			return classes.PostClassesClassIDMembersNotImplemented()
		}),
		ConversionsPostConversionHandler: conversions.PostConversionHandlerFunc(func(params conversions.PostConversionParams, principal interface{}) conversions.PostConversionResponder {
			return conversions.PostConversionNotImplemented()
		}),
		ConversionsPostConversionFromUnitToUnitHandler: conversions.PostConversionFromUnitToUnitHandlerFunc(func(params conversions.PostConversionFromUnitToUnitParams, principal interface{}) conversions.PostConversionFromUnitToUnitResponder {
			return conversions.PostConversionFromUnitToUnitNotImplemented()
		}),
		ClassesPutClassesClassIDHandler: classes.PutClassesClassIDHandlerFunc(func(params classes.PutClassesClassIDParams, principal interface{}) classes.PutClassesClassIDResponder {
			return classes.PutClassesClassIDNotImplemented()
		}),
		ClassesPutClassesClassIDMembersClassMemberIDHandler: classes.PutClassesClassIDMembersClassMemberIDHandlerFunc(func(params classes.PutClassesClassIDMembersClassMemberIDParams, principal interface{}) classes.PutClassesClassIDMembersClassMemberIDResponder {
			return classes.PutClassesClassIDMembersClassMemberIDNotImplemented()
		}),
		ConversionsPutConversionHandler: conversions.PutConversionHandlerFunc(func(params conversions.PutConversionParams, principal interface{}) conversions.PutConversionResponder {
			return conversions.PutConversionNotImplemented()
		}),
		ConversionsPutConversionFromUnitToUnitHandler: conversions.PutConversionFromUnitToUnitHandlerFunc(func(params conversions.PutConversionFromUnitToUnitParams, principal interface{}) conversions.PutConversionFromUnitToUnitResponder {
			return conversions.PutConversionFromUnitToUnitNotImplemented()
		}),

		BearerTokenAuth: func(token string, scopes []string) (interface{}, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (bearerToken) has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*ClimateChangeAdminAPIAPI The timeseries admin API allows to maintain the nomenclatures used by Climate Timeseries,
such as units etc.

Timeseries publication status and ownership may be modified.

Only extra administrative operations are exposed here (e.g. updating nomenclatures).

Regular data retrieval remains carried out by the public API.
*/
type ClimateChangeAdminAPIAPI struct {
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

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// BearerTokenAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	BearerTokenAuth func(string, []string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// ClassesDeleteClassesClassIDMembersClassMemberIDHandler sets the operation handler for the delete classes class ID members class member ID operation
	ClassesDeleteClassesClassIDMembersClassMemberIDHandler classes.DeleteClassesClassIDMembersClassMemberIDHandler
	// ClassesPostClassesClassIDMembersHandler sets the operation handler for the post classes class ID members operation
	ClassesPostClassesClassIDMembersHandler classes.PostClassesClassIDMembersHandler
	// ConversionsPostConversionHandler sets the operation handler for the post conversion operation
	ConversionsPostConversionHandler conversions.PostConversionHandler
	// ConversionsPostConversionFromUnitToUnitHandler sets the operation handler for the post conversion from unit to unit operation
	ConversionsPostConversionFromUnitToUnitHandler conversions.PostConversionFromUnitToUnitHandler
	// ClassesPutClassesClassIDHandler sets the operation handler for the put classes class ID operation
	ClassesPutClassesClassIDHandler classes.PutClassesClassIDHandler
	// ClassesPutClassesClassIDMembersClassMemberIDHandler sets the operation handler for the put classes class ID members class member ID operation
	ClassesPutClassesClassIDMembersClassMemberIDHandler classes.PutClassesClassIDMembersClassMemberIDHandler
	// ConversionsPutConversionHandler sets the operation handler for the put conversion operation
	ConversionsPutConversionHandler conversions.PutConversionHandler
	// ConversionsPutConversionFromUnitToUnitHandler sets the operation handler for the put conversion from unit to unit operation
	ConversionsPutConversionFromUnitToUnitHandler conversions.PutConversionFromUnitToUnitHandler

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
func (o *ClimateChangeAdminAPIAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *ClimateChangeAdminAPIAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *ClimateChangeAdminAPIAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ClimateChangeAdminAPIAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ClimateChangeAdminAPIAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ClimateChangeAdminAPIAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ClimateChangeAdminAPIAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ClimateChangeAdminAPIAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ClimateChangeAdminAPIAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ClimateChangeAdminAPIAPI
func (o *ClimateChangeAdminAPIAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BearerTokenAuth == nil {
		unregistered = append(unregistered, "BearerTokenAuth")
	}

	if o.ClassesDeleteClassesClassIDMembersClassMemberIDHandler == nil {
		unregistered = append(unregistered, "classes.DeleteClassesClassIDMembersClassMemberIDHandler")
	}
	if o.ClassesPostClassesClassIDMembersHandler == nil {
		unregistered = append(unregistered, "classes.PostClassesClassIDMembersHandler")
	}
	if o.ConversionsPostConversionHandler == nil {
		unregistered = append(unregistered, "conversions.PostConversionHandler")
	}
	if o.ConversionsPostConversionFromUnitToUnitHandler == nil {
		unregistered = append(unregistered, "conversions.PostConversionFromUnitToUnitHandler")
	}
	if o.ClassesPutClassesClassIDHandler == nil {
		unregistered = append(unregistered, "classes.PutClassesClassIDHandler")
	}
	if o.ClassesPutClassesClassIDMembersClassMemberIDHandler == nil {
		unregistered = append(unregistered, "classes.PutClassesClassIDMembersClassMemberIDHandler")
	}
	if o.ConversionsPutConversionHandler == nil {
		unregistered = append(unregistered, "conversions.PutConversionHandler")
	}
	if o.ConversionsPutConversionFromUnitToUnitHandler == nil {
		unregistered = append(unregistered, "conversions.PutConversionFromUnitToUnitHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ClimateChangeAdminAPIAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ClimateChangeAdminAPIAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "bearerToken":
			result[name] = o.BearerAuthenticator(name, o.BearerTokenAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *ClimateChangeAdminAPIAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *ClimateChangeAdminAPIAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
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
func (o *ClimateChangeAdminAPIAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
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
func (o *ClimateChangeAdminAPIAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the climate change admin API API
func (o *ClimateChangeAdminAPIAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ClimateChangeAdminAPIAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/classes/{classId}/members/{classMemberId}"] = classes.NewDeleteClassesClassIDMembersClassMemberID(o.context, o.ClassesDeleteClassesClassIDMembersClassMemberIDHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/classes/{classId}/members"] = classes.NewPostClassesClassIDMembers(o.context, o.ClassesPostClassesClassIDMembersHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/conversion"] = conversions.NewPostConversion(o.context, o.ConversionsPostConversionHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/conversion/{fromUnit}/{toUnit}"] = conversions.NewPostConversionFromUnitToUnit(o.context, o.ConversionsPostConversionFromUnitToUnitHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/classes/{classId}"] = classes.NewPutClassesClassID(o.context, o.ClassesPutClassesClassIDHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/classes/{classId}/members/{classMemberId}"] = classes.NewPutClassesClassIDMembersClassMemberID(o.context, o.ClassesPutClassesClassIDMembersClassMemberIDHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/conversion"] = conversions.NewPutConversion(o.context, o.ConversionsPutConversionHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/conversion/{fromUnit}/{toUnit}"] = conversions.NewPutConversionFromUnitToUnit(o.context, o.ConversionsPutConversionFromUnitToUnitHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ClimateChangeAdminAPIAPI) Serve(builder middleware.Builder) http.Handler {
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
func (o *ClimateChangeAdminAPIAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *ClimateChangeAdminAPIAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *ClimateChangeAdminAPIAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *ClimateChangeAdminAPIAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
