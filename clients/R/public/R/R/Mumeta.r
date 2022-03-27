# Climate time series API
#
# The API allows the public to search and consult time series about climate change research conducted by TheShiftProject.  The climate time series API allows contributors to upload time series about their climate change models and studies.  Other secured endpoints allows admins to maintain the nomenclatures used by the climate time series, such as units etc. Timeseries publication status and ownership is for now managed by admins. 
#
# OpenAPI spec version: v0.0.1
# Contact: fredbi@yahoo.com
# Generated by: https://github.com/swagger-api/swagger-codegen.git


#' Mumeta Class
#'
#' @field extra 
#'
#' @importFrom R6 R6Class
#' @importFrom jsonlite fromJSON toJSON
#' @export
Mumeta <- R6::R6Class(
  'Mumeta',
  public = list(
    `extra` = NULL,
    initialize = function(`extra`){
      if (!missing(`extra`)) {
        stopifnot(R6::is.R6(`extra`))
        self$`extra` <- `extra`
      }
    },
    toJSON = function() {
      MumetaObject <- list()
      if (!is.null(self$`extra`)) {
        MumetaObject[['extra']] <- self$`extra`$toJSON()
      }

      MumetaObject
    },
    fromJSON = function(MumetaJson) {
      MumetaObject <- jsonlite::fromJSON(MumetaJson)
      if (!is.null(MumetaObject$`extra`)) {
        extraObject <- Extra$new()
        extraObject$fromJSON(jsonlite::toJSON(MumetaObject$extra, auto_unbox = TRUE))
        self$`extra` <- extraObject
      }
    },
    toJSONString = function() {
       sprintf(
        '{
           "extra": %s
        }',
        self$`extra`$toJSON()
      )
    },
    fromJSONString = function(MumetaJson) {
      MumetaObject <- jsonlite::fromJSON(MumetaJson)
      ExtraObject <- Extra$new()
      self$`extra` <- ExtraObject$fromJSON(jsonlite::toJSON(MumetaObject$extra, auto_unbox = TRUE))
    }
  )
)
