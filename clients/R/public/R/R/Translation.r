# Climate time series API
#
# The API allows the public to search and consult time series about climate change research conducted by TheShiftProject.  The climate time series API allows contributors to upload time series about their climate change models and studies.  Other secured endpoints allows admins to maintain the nomenclatures used by the climate time series, such as units etc. Timeseries publication status and ownership is for now managed by admins. 
#
# OpenAPI spec version: v0.0.1
# Contact: fredbi@yahoo.com
# Generated by: https://github.com/swagger-api/swagger-codegen.git


#' Translation Class
#'
#'
#' @importFrom R6 R6Class
#' @importFrom jsonlite fromJSON toJSON
#' @export
Translation <- R6::R6Class(
  'Translation',
  public = list(
    initialize = function(){
    },
    toJSON = function() {
      TranslationObject <- list()

      TranslationObject
    },
    fromJSON = function(TranslationJson) {
      TranslationObject <- jsonlite::fromJSON(TranslationJson)
    },
    toJSONString = function() {
       sprintf(
        '{
        }',
      )
    },
    fromJSONString = function(TranslationJson) {
      TranslationObject <- jsonlite::fromJSON(TranslationJson)
    }
  )
)
