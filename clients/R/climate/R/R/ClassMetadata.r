# Climate time series API
#
# The API allows the public to search and consult time series about climate change research conducted by TheShiftProject.  The climate time series API allows contributors to upload time series about their climate change models and studies.  Other secured endpoints allows admins to maintain the nomenclatures used by the climate time series, such as units etc. Timeseries publication status and ownership is for now managed by admins. 
#
# OpenAPI spec version: v0.0.1
# Contact: fredbi@yahoo.com
# Generated by: https://github.com/swagger-api/swagger-codegen.git


#' ClassMetadata Class
#'
#' @field fromTemplate 
#' @field tagSearch 
#' @field extraFields 
#' @field hasOneClass 
#' @field hasZeroOneClass 
#' @field hasZeroManyClass 
#'
#' @importFrom R6 R6Class
#' @importFrom jsonlite fromJSON toJSON
#' @export
ClassMetadata <- R6::R6Class(
  'ClassMetadata',
  public = list(
    `fromTemplate` = NULL,
    `tagSearch` = NULL,
    `extraFields` = NULL,
    `hasOneClass` = NULL,
    `hasZeroOneClass` = NULL,
    `hasZeroManyClass` = NULL,
    initialize = function(`fromTemplate`, `tagSearch`, `extraFields`, `hasOneClass`, `hasZeroOneClass`, `hasZeroManyClass`){
      if (!missing(`fromTemplate`)) {
        self$`fromTemplate` <- `fromTemplate`
      }
      if (!missing(`tagSearch`)) {
        self$`tagSearch` <- `tagSearch`
      }
      if (!missing(`extraFields`)) {
        stopifnot(is.list(`extraFields`), length(`extraFields`) != 0)
        lapply(`extraFields`, function(x) stopifnot(is.character(x)))
        self$`extraFields` <- `extraFields`
      }
      if (!missing(`hasOneClass`)) {
        self$`hasOneClass` <- `hasOneClass`
      }
      if (!missing(`hasZeroOneClass`)) {
        self$`hasZeroOneClass` <- `hasZeroOneClass`
      }
      if (!missing(`hasZeroManyClass`)) {
        self$`hasZeroManyClass` <- `hasZeroManyClass`
      }
    },
    toJSON = function() {
      ClassMetadataObject <- list()
      if (!is.null(self$`fromTemplate`)) {
        ClassMetadataObject[['fromTemplate']] <- self$`fromTemplate`
      }
      if (!is.null(self$`tagSearch`)) {
        ClassMetadataObject[['tagSearch']] <- self$`tagSearch`
      }
      if (!is.null(self$`extraFields`)) {
        ClassMetadataObject[['extraFields']] <- self$`extraFields`
      }
      if (!is.null(self$`hasOneClass`)) {
        ClassMetadataObject[['hasOneClass']] <- self$`hasOneClass`
      }
      if (!is.null(self$`hasZeroOneClass`)) {
        ClassMetadataObject[['hasZeroOneClass']] <- self$`hasZeroOneClass`
      }
      if (!is.null(self$`hasZeroManyClass`)) {
        ClassMetadataObject[['hasZeroManyClass']] <- self$`hasZeroManyClass`
      }

      ClassMetadataObject
    },
    fromJSON = function(ClassMetadataJson) {
      ClassMetadataObject <- jsonlite::fromJSON(ClassMetadataJson)
      if (!is.null(ClassMetadataObject$`fromTemplate`)) {
        self$`fromTemplate` <- ClassMetadataObject$`fromTemplate`
      }
      if (!is.null(ClassMetadataObject$`tagSearch`)) {
        self$`tagSearch` <- ClassMetadataObject$`tagSearch`
      }
      if (!is.null(ClassMetadataObject$`extraFields`)) {
        self$`extraFields` <- ClassMetadataObject$`extraFields`
      }
      if (!is.null(ClassMetadataObject$`hasOneClass`)) {
        self$`hasOneClass` <- ClassMetadataObject$`hasOneClass`
      }
      if (!is.null(ClassMetadataObject$`hasZeroOneClass`)) {
        self$`hasZeroOneClass` <- ClassMetadataObject$`hasZeroOneClass`
      }
      if (!is.null(ClassMetadataObject$`hasZeroManyClass`)) {
        self$`hasZeroManyClass` <- ClassMetadataObject$`hasZeroManyClass`
      }
    },
    toJSONString = function() {
       sprintf(
        '{
           "fromTemplate": %s,
           "tagSearch": %s,
           "extraFields": [%s],
           "hasOneClass": %s,
           "hasZeroOneClass": %s,
           "hasZeroManyClass": %s
        }',
        self$`fromTemplate`,
        self$`tagSearch`,
        lapply(self$`extraFields`, function(x) paste(paste0('"', x, '"'), sep=",")),
        self$`hasOneClass`,
        self$`hasZeroOneClass`,
        self$`hasZeroManyClass`
      )
    },
    fromJSONString = function(ClassMetadataJson) {
      ClassMetadataObject <- jsonlite::fromJSON(ClassMetadataJson)
      self$`fromTemplate` <- ClassMetadataObject$`fromTemplate`
      self$`tagSearch` <- ClassMetadataObject$`tagSearch`
      self$`extraFields` <- ClassMetadataObject$`extraFields`
      self$`hasOneClass` <- ClassMetadataObject$`hasOneClass`
      self$`hasZeroOneClass` <- ClassMetadataObject$`hasZeroOneClass`
      self$`hasZeroManyClass` <- ClassMetadataObject$`hasZeroManyClass`
    }
  )
)