# Climate time series API
#
# The API allows the public to search and consult time series about climate change research conducted by TheShiftProject.  The climate time series API allows contributors to upload time series about their climate change models and studies.  Other secured endpoints allows admins to maintain the nomenclatures used by the climate time series, such as units etc. Timeseries publication status and ownership is for now managed by admins. 
#
# OpenAPI spec version: v0.0.1
# Contact: fredbi@yahoo.com
# Generated by: https://github.com/swagger-api/swagger-codegen.git


#' ClassNomenclature Class
#'
#' @field id 
#' @field shortCode 
#' @field title 
#' @field descriptionShort 
#' @field descriptionLong 
#' @field auditTrail 
#'
#' @importFrom R6 R6Class
#' @importFrom jsonlite fromJSON toJSON
#' @export
ClassNomenclature <- R6::R6Class(
  'ClassNomenclature',
  public = list(
    `id` = NULL,
    `shortCode` = NULL,
    `title` = NULL,
    `descriptionShort` = NULL,
    `descriptionLong` = NULL,
    `auditTrail` = NULL,
    initialize = function(`id`, `shortCode`, `title`, `descriptionShort`, `descriptionLong`, `auditTrail`){
      if (!missing(`id`)) {
        stopifnot(is.numeric(`id`), length(`id`) == 1)
        self$`id` <- `id`
      }
      if (!missing(`shortCode`)) {
        stopifnot(R6::is.R6(`shortCode`))
        self$`shortCode` <- `shortCode`
      }
      if (!missing(`title`)) {
        stopifnot(R6::is.R6(`title`))
        self$`title` <- `title`
      }
      if (!missing(`descriptionShort`)) {
        stopifnot(R6::is.R6(`descriptionShort`))
        self$`descriptionShort` <- `descriptionShort`
      }
      if (!missing(`descriptionLong`)) {
        stopifnot(R6::is.R6(`descriptionLong`))
        self$`descriptionLong` <- `descriptionLong`
      }
      if (!missing(`auditTrail`)) {
        stopifnot(R6::is.R6(`auditTrail`))
        self$`auditTrail` <- `auditTrail`
      }
    },
    toJSON = function() {
      ClassNomenclatureObject <- list()
      if (!is.null(self$`id`)) {
        ClassNomenclatureObject[['id']] <- self$`id`
      }
      if (!is.null(self$`shortCode`)) {
        ClassNomenclatureObject[['shortCode']] <- self$`shortCode`$toJSON()
      }
      if (!is.null(self$`title`)) {
        ClassNomenclatureObject[['title']] <- self$`title`$toJSON()
      }
      if (!is.null(self$`descriptionShort`)) {
        ClassNomenclatureObject[['descriptionShort']] <- self$`descriptionShort`$toJSON()
      }
      if (!is.null(self$`descriptionLong`)) {
        ClassNomenclatureObject[['descriptionLong']] <- self$`descriptionLong`$toJSON()
      }
      if (!is.null(self$`auditTrail`)) {
        ClassNomenclatureObject[['auditTrail']] <- self$`auditTrail`$toJSON()
      }

      ClassNomenclatureObject
    },
    fromJSON = function(ClassNomenclatureJson) {
      ClassNomenclatureObject <- jsonlite::fromJSON(ClassNomenclatureJson)
      if (!is.null(ClassNomenclatureObject$`id`)) {
        self$`id` <- ClassNomenclatureObject$`id`
      }
      if (!is.null(ClassNomenclatureObject$`shortCode`)) {
        shortCodeObject <- ClassNomenclatureName$new()
        shortCodeObject$fromJSON(jsonlite::toJSON(ClassNomenclatureObject$shortCode, auto_unbox = TRUE))
        self$`shortCode` <- shortCodeObject
      }
      if (!is.null(ClassNomenclatureObject$`title`)) {
        titleObject <- Translation$new()
        titleObject$fromJSON(jsonlite::toJSON(ClassNomenclatureObject$title, auto_unbox = TRUE))
        self$`title` <- titleObject
      }
      if (!is.null(ClassNomenclatureObject$`descriptionShort`)) {
        descriptionShortObject <- Translation$new()
        descriptionShortObject$fromJSON(jsonlite::toJSON(ClassNomenclatureObject$descriptionShort, auto_unbox = TRUE))
        self$`descriptionShort` <- descriptionShortObject
      }
      if (!is.null(ClassNomenclatureObject$`descriptionLong`)) {
        descriptionLongObject <- Translation$new()
        descriptionLongObject$fromJSON(jsonlite::toJSON(ClassNomenclatureObject$descriptionLong, auto_unbox = TRUE))
        self$`descriptionLong` <- descriptionLongObject
      }
      if (!is.null(ClassNomenclatureObject$`auditTrail`)) {
        auditTrailObject <- Audit$new()
        auditTrailObject$fromJSON(jsonlite::toJSON(ClassNomenclatureObject$auditTrail, auto_unbox = TRUE))
        self$`auditTrail` <- auditTrailObject
      }
    },
    toJSONString = function() {
       sprintf(
        '{
           "id": %d,
           "shortCode": %s,
           "title": %s,
           "descriptionShort": %s,
           "descriptionLong": %s,
           "auditTrail": %s
        }',
        self$`id`,
        self$`shortCode`$toJSON(),
        self$`title`$toJSON(),
        self$`descriptionShort`$toJSON(),
        self$`descriptionLong`$toJSON(),
        self$`auditTrail`$toJSON()
      )
    },
    fromJSONString = function(ClassNomenclatureJson) {
      ClassNomenclatureObject <- jsonlite::fromJSON(ClassNomenclatureJson)
      self$`id` <- ClassNomenclatureObject$`id`
      ClassNomenclatureNameObject <- ClassNomenclatureName$new()
      self$`shortCode` <- ClassNomenclatureNameObject$fromJSON(jsonlite::toJSON(ClassNomenclatureObject$shortCode, auto_unbox = TRUE))
      TranslationObject <- Translation$new()
      self$`title` <- TranslationObject$fromJSON(jsonlite::toJSON(ClassNomenclatureObject$title, auto_unbox = TRUE))
      TranslationObject <- Translation$new()
      self$`descriptionShort` <- TranslationObject$fromJSON(jsonlite::toJSON(ClassNomenclatureObject$descriptionShort, auto_unbox = TRUE))
      TranslationObject <- Translation$new()
      self$`descriptionLong` <- TranslationObject$fromJSON(jsonlite::toJSON(ClassNomenclatureObject$descriptionLong, auto_unbox = TRUE))
      AuditObject <- Audit$new()
      self$`auditTrail` <- AuditObject$fromJSON(jsonlite::toJSON(ClassNomenclatureObject$auditTrail, auto_unbox = TRUE))
    }
  )
)
