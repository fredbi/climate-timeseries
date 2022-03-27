# Climate time series API
#
# The API allows the public to search and consult time series about climate change research conducted by TheShiftProject.  The climate time series API allows contributors to upload time series about their climate change models and studies.  Other secured endpoints allows admins to maintain the nomenclatures used by the climate time series, such as units etc. Timeseries publication status and ownership is for now managed by admins. 
#
# OpenAPI spec version: v0.0.1
# Contact: fredbi@yahoo.com
# Generated by: https://github.com/swagger-api/swagger-codegen.git


#' TsValues Class
#'
#'
#' @importFrom R6 R6Class
#' @importFrom jsonlite fromJSON toJSON
#' @export
TsValues <- R6::R6Class(
  'TsValues',
  public = list(
    initialize = function(){
    },
    toJSON = function() {
      TsValuesObject <- list()

      TsValuesObject
    },
    fromJSON = function(TsValuesJson) {
      TsValuesObject <- jsonlite::fromJSON(TsValuesJson)
    },
    toJSONString = function() {
       sprintf(
        '{
        }',
      )
    },
    fromJSONString = function(TsValuesJson) {
      TsValuesObject <- jsonlite::fromJSON(TsValuesJson)
    }
  )
)