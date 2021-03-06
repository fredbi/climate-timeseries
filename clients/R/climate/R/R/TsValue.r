# Climate time series API
#
# The API allows the public to search and consult time series about climate change research conducted by TheShiftProject.  The climate time series API allows contributors to upload time series about their climate change models and studies.  Other secured endpoints allows admins to maintain the nomenclatures used by the climate time series, such as units etc. Timeseries publication status and ownership is for now managed by admins. 
#
# OpenAPI spec version: v0.0.1
# Contact: fredbi@yahoo.com
# Generated by: https://github.com/swagger-api/swagger-codegen.git


#' TsValue Class
#'
#' @field t 
#' @field v 
#'
#' @importFrom R6 R6Class
#' @importFrom jsonlite fromJSON toJSON
#' @export
TsValue <- R6::R6Class(
  'TsValue',
  public = list(
    `t` = NULL,
    `v` = NULL,
    initialize = function(`t`, `v`){
      if (!missing(`t`)) {
        stopifnot(R6::is.R6(`t`))
        self$`t` <- `t`
      }
      if (!missing(`v`)) {
        stopifnot(is.numeric(`v`), length(`v`) == 1)
        self$`v` <- `v`
      }
    },
    toJSON = function() {
      TsValueObject <- list()
      if (!is.null(self$`t`)) {
        TsValueObject[['t']] <- self$`t`$toJSON()
      }
      if (!is.null(self$`v`)) {
        TsValueObject[['v']] <- self$`v`
      }

      TsValueObject
    },
    fromJSON = function(TsValueJson) {
      TsValueObject <- jsonlite::fromJSON(TsValueJson)
      if (!is.null(TsValueObject$`t`)) {
        tObject <- Ts$new()
        tObject$fromJSON(jsonlite::toJSON(TsValueObject$t, auto_unbox = TRUE))
        self$`t` <- tObject
      }
      if (!is.null(TsValueObject$`v`)) {
        self$`v` <- TsValueObject$`v`
      }
    },
    toJSONString = function() {
       sprintf(
        '{
           "t": %s,
           "v": %d
        }',
        self$`t`$toJSON(),
        self$`v`
      )
    },
    fromJSONString = function(TsValueJson) {
      TsValueObject <- jsonlite::fromJSON(TsValueJson)
      TsObject <- Ts$new()
      self$`t` <- TsObject$fromJSON(jsonlite::toJSON(TsValueObject$t, auto_unbox = TRUE))
      self$`v` <- TsValueObject$`v`
    }
  )
)
