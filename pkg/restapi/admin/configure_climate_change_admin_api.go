// This file is safe to edit. Once it exists it will not be overwritten

package admin

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/fredbi/climate-timeseries/pkg/restapi/admin/operations"
	"github.com/fredbi/climate-timeseries/pkg/restapi/admin/operations/classes"
	"github.com/fredbi/climate-timeseries/pkg/restapi/admin/operations/conversions"
)

//go:generate swagger generate server --target ../../restapi --name ClimateChangeAdminAPI --spec ../../../api/swagger/admin.yaml --server-package admin --principal interface{} --default-scheme https --skip-models --exclude-main --strict-responders

func configureFlags(api *operations.ClimateChangeAdminAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ClimateChangeAdminAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.BearerTokenAuth == nil {
		api.BearerTokenAuth = func(token string, scopes []string) (interface{}, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (bearerToken) has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.ClassesDeleteClassesClassIDMembersClassMemberIDHandler == nil {
		api.ClassesDeleteClassesClassIDMembersClassMemberIDHandler = classes.DeleteClassesClassIDMembersClassMemberIDHandlerFunc(func(params classes.DeleteClassesClassIDMembersClassMemberIDParams, principal interface{}) classes.DeleteClassesClassIDMembersClassMemberIDResponder {
			return classes.DeleteClassesClassIDMembersClassMemberIDNotImplemented()
		})
	}
	if api.ClassesPostClassesClassIDMembersHandler == nil {
		api.ClassesPostClassesClassIDMembersHandler = classes.PostClassesClassIDMembersHandlerFunc(func(params classes.PostClassesClassIDMembersParams, principal interface{}) classes.PostClassesClassIDMembersResponder {
			return classes.PostClassesClassIDMembersNotImplemented()
		})
	}
	if api.ConversionsPostConversionHandler == nil {
		api.ConversionsPostConversionHandler = conversions.PostConversionHandlerFunc(func(params conversions.PostConversionParams, principal interface{}) conversions.PostConversionResponder {
			return conversions.PostConversionNotImplemented()
		})
	}
	if api.ConversionsPostConversionFromUnitToUnitHandler == nil {
		api.ConversionsPostConversionFromUnitToUnitHandler = conversions.PostConversionFromUnitToUnitHandlerFunc(func(params conversions.PostConversionFromUnitToUnitParams, principal interface{}) conversions.PostConversionFromUnitToUnitResponder {
			return conversions.PostConversionFromUnitToUnitNotImplemented()
		})
	}
	if api.ClassesPutClassesClassIDHandler == nil {
		api.ClassesPutClassesClassIDHandler = classes.PutClassesClassIDHandlerFunc(func(params classes.PutClassesClassIDParams, principal interface{}) classes.PutClassesClassIDResponder {
			return classes.PutClassesClassIDNotImplemented()
		})
	}
	if api.ClassesPutClassesClassIDMembersClassMemberIDHandler == nil {
		api.ClassesPutClassesClassIDMembersClassMemberIDHandler = classes.PutClassesClassIDMembersClassMemberIDHandlerFunc(func(params classes.PutClassesClassIDMembersClassMemberIDParams, principal interface{}) classes.PutClassesClassIDMembersClassMemberIDResponder {
			return classes.PutClassesClassIDMembersClassMemberIDNotImplemented()
		})
	}
	if api.ConversionsPutConversionHandler == nil {
		api.ConversionsPutConversionHandler = conversions.PutConversionHandlerFunc(func(params conversions.PutConversionParams, principal interface{}) conversions.PutConversionResponder {
			return conversions.PutConversionNotImplemented()
		})
	}
	if api.ConversionsPutConversionFromUnitToUnitHandler == nil {
		api.ConversionsPutConversionFromUnitToUnitHandler = conversions.PutConversionFromUnitToUnitHandlerFunc(func(params conversions.PutConversionFromUnitToUnitParams, principal interface{}) conversions.PutConversionFromUnitToUnitResponder {
			return conversions.PutConversionFromUnitToUnitNotImplemented()
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
