# Climate time series API
#
# The API allows the public to search and consult time series about climate change research conducted by TheShiftProject.  The climate time series API allows contributors to upload time series about their climate change models and studies.  Other secured endpoints allows admins to maintain the nomenclatures used by the climate time series, such as units etc. Timeseries publication status and ownership is for now managed by admins. 
#
# OpenAPI spec version: v0.0.1
# Contact: fredbi@yahoo.com
# Generated by: https://github.com/swagger-api/swagger-codegen.git


#' TimeseriesModelsYamldefinitionsconversionSpec Class
#'
#' @field toUnitCode 
#' @field factor 
#' @field intercept 
#' @field formula 
#' @field reverse_formula 
#'
#' @importFrom R6 R6Class
#' @importFrom jsonlite fromJSON toJSON
#' @export
TimeseriesModelsYamldefinitionsconversionSpec <- R6::R6Class(
  'TimeseriesModelsYamldefinitionsconversionSpec',
  public = list(
    `toUnitCode` = NULL,
    `factor` = NULL,
    `intercept` = NULL,
    `formula` = NULL,
    `reverse_formula` = NULL,
    initialize = function(`toUnitCode`, `factor`, `intercept`, `formula`, `reverse_formula`){
      if (!missing(`toUnitCode`)) {
        stopifnot(is.character(`toUnitCode`), length(`toUnitCode`) == 1)
        self$`toUnitCode` <- `toUnitCode`
      }
      if (!missing(`factor`)) {
        stopifnot(is.numeric(`factor`), length(`factor`) == 1)
        self$`factor` <- `factor`
      }
      if (!missing(`intercept`)) {
        stopifnot(is.numeric(`intercept`), length(`intercept`) == 1)
        self$`intercept` <- `intercept`
      }
      if (!missing(`formula`)) {
        stopifnot(is.character(`formula`), length(`formula`) == 1)
        self$`formula` <- `formula`
      }
      if (!missing(`reverse_formula`)) {
        stopifnot(is.character(`reverse_formula`), length(`reverse_formula`) == 1)
        self$`reverse_formula` <- `reverse_formula`
      }
    },
    toJSON = function() {
      TimeseriesModelsYamldefinitionsconversionSpecObject <- list()
      if (!is.null(self$`toUnitCode`)) {
        TimeseriesModelsYamldefinitionsconversionSpecObject[['toUnitCode']] <- self$`toUnitCode`
      }
      if (!is.null(self$`factor`)) {
        TimeseriesModelsYamldefinitionsconversionSpecObject[['factor']] <- self$`factor`
      }
      if (!is.null(self$`intercept`)) {
        TimeseriesModelsYamldefinitionsconversionSpecObject[['intercept']] <- self$`intercept`
      }
      if (!is.null(self$`formula`)) {
        TimeseriesModelsYamldefinitionsconversionSpecObject[['formula']] <- self$`formula`
      }
      if (!is.null(self$`reverse_formula`)) {
        TimeseriesModelsYamldefinitionsconversionSpecObject[['reverse_formula']] <- self$`reverse_formula`
      }

      TimeseriesModelsYamldefinitionsconversionSpecObject
    },
    fromJSON = function(TimeseriesModelsYamldefinitionsconversionSpecJson) {
      TimeseriesModelsYamldefinitionsconversionSpecObject <- jsonlite::fromJSON(TimeseriesModelsYamldefinitionsconversionSpecJson)
      if (!is.null(TimeseriesModelsYamldefinitionsconversionSpecObject$`toUnitCode`)) {
        self$`toUnitCode` <- TimeseriesModelsYamldefinitionsconversionSpecObject$`toUnitCode`
      }
      if (!is.null(TimeseriesModelsYamldefinitionsconversionSpecObject$`factor`)) {
        self$`factor` <- TimeseriesModelsYamldefinitionsconversionSpecObject$`factor`
      }
      if (!is.null(TimeseriesModelsYamldefinitionsconversionSpecObject$`intercept`)) {
        self$`intercept` <- TimeseriesModelsYamldefinitionsconversionSpecObject$`intercept`
      }
      if (!is.null(TimeseriesModelsYamldefinitionsconversionSpecObject$`formula`)) {
        self$`formula` <- TimeseriesModelsYamldefinitionsconversionSpecObject$`formula`
      }
      if (!is.null(TimeseriesModelsYamldefinitionsconversionSpecObject$`reverse_formula`)) {
        self$`reverse_formula` <- TimeseriesModelsYamldefinitionsconversionSpecObject$`reverse_formula`
      }
    },
    toJSONString = function() {
       sprintf(
        '{
           "toUnitCode": %s,
           "factor": %d,
           "intercept": %d,
           "formula": %s,
           "reverse_formula": %s
        }',
        self$`toUnitCode`,
        self$`factor`,
        self$`intercept`,
        self$`formula`,
        self$`reverse_formula`
      )
    },
    fromJSONString = function(TimeseriesModelsYamldefinitionsconversionSpecJson) {
      TimeseriesModelsYamldefinitionsconversionSpecObject <- jsonlite::fromJSON(TimeseriesModelsYamldefinitionsconversionSpecJson)
      self$`toUnitCode` <- TimeseriesModelsYamldefinitionsconversionSpecObject$`toUnitCode`
      self$`factor` <- TimeseriesModelsYamldefinitionsconversionSpecObject$`factor`
      self$`intercept` <- TimeseriesModelsYamldefinitionsconversionSpecObject$`intercept`
      self$`formula` <- TimeseriesModelsYamldefinitionsconversionSpecObject$`formula`
      self$`reverse_formula` <- TimeseriesModelsYamldefinitionsconversionSpecObject$`reverse_formula`
    }
  )
)