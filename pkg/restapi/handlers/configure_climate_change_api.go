// This file is safe to edit. Once it exists it will not be overwritten

package handlers

import (
	"crypto/tls"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/fredbi/climate-timeseries/pkg/restapi/handlers/operations"
	"github.com/fredbi/climate-timeseries/pkg/restapi/handlers/operations/classes"
	"github.com/fredbi/climate-timeseries/pkg/restapi/handlers/operations/series"
	"github.com/fredbi/climate-timeseries/pkg/restapi/handlers/operations/tags"
)

//go:generate swagger generate server --target ../../restapi --name ClimateChangeAPI --spec ../../../api/swagger/timeseries.yaml --server-package handlers --principal interface{} --default-scheme https --skip-models --exclude-main --strict-responders

func configureFlags(api *operations.ClimateChangeAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ClimateChangeAPIAPI) http.Handler {
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

	api.CsvConsumer = runtime.ConsumerFunc(func(r io.Reader, target interface{}) error {
		return errors.NotImplemented("csv consumer has not yet been implemented")
	})
	api.JSONConsumer = runtime.JSONConsumer()

	api.CsvProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("csv producer has not yet been implemented")
	})
	api.JSONProducer = runtime.JSONProducer()

	if api.SeriesDeleteSeriesSeriesIDHandler == nil {
		api.SeriesDeleteSeriesSeriesIDHandler = series.DeleteSeriesSeriesIDHandlerFunc(func(params series.DeleteSeriesSeriesIDParams) series.DeleteSeriesSeriesIDResponder {
			return series.DeleteSeriesSeriesIDNotImplemented()
		})
	}
	if api.SeriesDeleteSeriesVersionsVersionedSeriesIDHandler == nil {
		api.SeriesDeleteSeriesVersionsVersionedSeriesIDHandler = series.DeleteSeriesVersionsVersionedSeriesIDHandlerFunc(func(params series.DeleteSeriesVersionsVersionedSeriesIDParams) series.DeleteSeriesVersionsVersionedSeriesIDResponder {
			return series.DeleteSeriesVersionsVersionedSeriesIDNotImplemented()
		})
	}
	if api.SeriesDeleteSeriesVersionsVersionedSeriesIDValuesHandler == nil {
		api.SeriesDeleteSeriesVersionsVersionedSeriesIDValuesHandler = series.DeleteSeriesVersionsVersionedSeriesIDValuesHandlerFunc(func(params series.DeleteSeriesVersionsVersionedSeriesIDValuesParams) series.DeleteSeriesVersionsVersionedSeriesIDValuesResponder {
			return series.DeleteSeriesVersionsVersionedSeriesIDValuesNotImplemented()
		})
	}
	if api.ClassesGetClassesHandler == nil {
		api.ClassesGetClassesHandler = classes.GetClassesHandlerFunc(func(params classes.GetClassesParams) classes.GetClassesResponder {
			return classes.GetClassesNotImplemented()
		})
	}
	if api.ClassesGetClassesClassIDHandler == nil {
		api.ClassesGetClassesClassIDHandler = classes.GetClassesClassIDHandlerFunc(func(params classes.GetClassesClassIDParams) classes.GetClassesClassIDResponder {
			return classes.GetClassesClassIDNotImplemented()
		})
	}
	if api.ClassesGetClassesClassIDMembersHandler == nil {
		api.ClassesGetClassesClassIDMembersHandler = classes.GetClassesClassIDMembersHandlerFunc(func(params classes.GetClassesClassIDMembersParams) classes.GetClassesClassIDMembersResponder {
			return classes.GetClassesClassIDMembersNotImplemented()
		})
	}
	if api.TagsGetSearchTagsHandler == nil {
		api.TagsGetSearchTagsHandler = tags.GetSearchTagsHandlerFunc(func(params tags.GetSearchTagsParams) tags.GetSearchTagsResponder {
			return tags.GetSearchTagsNotImplemented()
		})
	}
	if api.TagsGetSearchTagsTagHandler == nil {
		api.TagsGetSearchTagsTagHandler = tags.GetSearchTagsTagHandlerFunc(func(params tags.GetSearchTagsTagParams) tags.GetSearchTagsTagResponder {
			return tags.GetSearchTagsTagNotImplemented()
		})
	}
	if api.SeriesGetSeriesHandler == nil {
		api.SeriesGetSeriesHandler = series.GetSeriesHandlerFunc(func(params series.GetSeriesParams) series.GetSeriesResponder {
			return series.GetSeriesNotImplemented()
		})
	}
	if api.SeriesGetSeriesSeriesIDHandler == nil {
		api.SeriesGetSeriesSeriesIDHandler = series.GetSeriesSeriesIDHandlerFunc(func(params series.GetSeriesSeriesIDParams) series.GetSeriesSeriesIDResponder {
			return series.GetSeriesSeriesIDNotImplemented()
		})
	}
	if api.SeriesGetSeriesSeriesIDLatestHandler == nil {
		api.SeriesGetSeriesSeriesIDLatestHandler = series.GetSeriesSeriesIDLatestHandlerFunc(func(params series.GetSeriesSeriesIDLatestParams) series.GetSeriesSeriesIDLatestResponder {
			return series.GetSeriesSeriesIDLatestNotImplemented()
		})
	}
	if api.SeriesGetSeriesSeriesIDLatestValuesHandler == nil {
		api.SeriesGetSeriesSeriesIDLatestValuesHandler = series.GetSeriesSeriesIDLatestValuesHandlerFunc(func(params series.GetSeriesSeriesIDLatestValuesParams) series.GetSeriesSeriesIDLatestValuesResponder {
			return series.GetSeriesSeriesIDLatestValuesNotImplemented()
		})
	}
	if api.SeriesGetSeriesSeriesIDSemverHandler == nil {
		api.SeriesGetSeriesSeriesIDSemverHandler = series.GetSeriesSeriesIDSemverHandlerFunc(func(params series.GetSeriesSeriesIDSemverParams) series.GetSeriesSeriesIDSemverResponder {
			return series.GetSeriesSeriesIDSemverNotImplemented()
		})
	}
	if api.SeriesGetSeriesSeriesIDSemverSemverHandler == nil {
		api.SeriesGetSeriesSeriesIDSemverSemverHandler = series.GetSeriesSeriesIDSemverSemverHandlerFunc(func(params series.GetSeriesSeriesIDSemverSemverParams) series.GetSeriesSeriesIDSemverSemverResponder {
			return series.GetSeriesSeriesIDSemverSemverNotImplemented()
		})
	}
	if api.SeriesGetSeriesSeriesIDSemverSemverValuesHandler == nil {
		api.SeriesGetSeriesSeriesIDSemverSemverValuesHandler = series.GetSeriesSeriesIDSemverSemverValuesHandlerFunc(func(params series.GetSeriesSeriesIDSemverSemverValuesParams) series.GetSeriesSeriesIDSemverSemverValuesResponder {
			return series.GetSeriesSeriesIDSemverSemverValuesNotImplemented()
		})
	}
	if api.SeriesGetSeriesVersionsVersionedSeriesIDHandler == nil {
		api.SeriesGetSeriesVersionsVersionedSeriesIDHandler = series.GetSeriesVersionsVersionedSeriesIDHandlerFunc(func(params series.GetSeriesVersionsVersionedSeriesIDParams) series.GetSeriesVersionsVersionedSeriesIDResponder {
			return series.GetSeriesVersionsVersionedSeriesIDNotImplemented()
		})
	}
	if api.SeriesGetSeriesVersionsVersionedSeriesIDValuesHandler == nil {
		api.SeriesGetSeriesVersionsVersionedSeriesIDValuesHandler = series.GetSeriesVersionsVersionedSeriesIDValuesHandlerFunc(func(params series.GetSeriesVersionsVersionedSeriesIDValuesParams) series.GetSeriesVersionsVersionedSeriesIDValuesResponder {
			return series.GetSeriesVersionsVersionedSeriesIDValuesNotImplemented()
		})
	}
	if api.SeriesPostSeriesHandler == nil {
		api.SeriesPostSeriesHandler = series.PostSeriesHandlerFunc(func(params series.PostSeriesParams) series.PostSeriesResponder {
			return series.PostSeriesNotImplemented()
		})
	}
	if api.SeriesPostSeriesSeriesIDHandler == nil {
		api.SeriesPostSeriesSeriesIDHandler = series.PostSeriesSeriesIDHandlerFunc(func(params series.PostSeriesSeriesIDParams) series.PostSeriesSeriesIDResponder {
			return series.PostSeriesSeriesIDNotImplemented()
		})
	}
	if api.SeriesPostSeriesVersionsVersionedSeriesIDHandler == nil {
		api.SeriesPostSeriesVersionsVersionedSeriesIDHandler = series.PostSeriesVersionsVersionedSeriesIDHandlerFunc(func(params series.PostSeriesVersionsVersionedSeriesIDParams) series.PostSeriesVersionsVersionedSeriesIDResponder {
			return series.PostSeriesVersionsVersionedSeriesIDNotImplemented()
		})
	}
	if api.SeriesPostSeriesVersionsVersionedSeriesIDValuesHandler == nil {
		api.SeriesPostSeriesVersionsVersionedSeriesIDValuesHandler = series.PostSeriesVersionsVersionedSeriesIDValuesHandlerFunc(func(params series.PostSeriesVersionsVersionedSeriesIDValuesParams) series.PostSeriesVersionsVersionedSeriesIDValuesResponder {
			return series.PostSeriesVersionsVersionedSeriesIDValuesNotImplemented()
		})
	}
	if api.SeriesPutSeriesSeriesIDHandler == nil {
		api.SeriesPutSeriesSeriesIDHandler = series.PutSeriesSeriesIDHandlerFunc(func(params series.PutSeriesSeriesIDParams) series.PutSeriesSeriesIDResponder {
			return series.PutSeriesSeriesIDNotImplemented()
		})
	}
	if api.SeriesPutSeriesVersionsVersionedSeriesIDHandler == nil {
		api.SeriesPutSeriesVersionsVersionedSeriesIDHandler = series.PutSeriesVersionsVersionedSeriesIDHandlerFunc(func(params series.PutSeriesVersionsVersionedSeriesIDParams) series.PutSeriesVersionsVersionedSeriesIDResponder {
			return series.PutSeriesVersionsVersionedSeriesIDNotImplemented()
		})
	}
	if api.SeriesPutSeriesVersionsVersionedSeriesIDValuesHandler == nil {
		api.SeriesPutSeriesVersionsVersionedSeriesIDValuesHandler = series.PutSeriesVersionsVersionedSeriesIDValuesHandlerFunc(func(params series.PutSeriesVersionsVersionedSeriesIDValuesParams) series.PutSeriesVersionsVersionedSeriesIDValuesResponder {
			return series.PutSeriesVersionsVersionedSeriesIDValuesNotImplemented()
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
