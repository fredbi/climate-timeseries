# Climate time series API
#
# The API allows the public to search and consult time series about climate change research conducted by TheShiftProject.  The climate time series API allows contributors to upload time series about their climate change models and studies.  Other secured endpoints allows admins to maintain the nomenclatures used by the climate time series, such as units etc. Timeseries publication status and ownership is for now managed by admins. 
#
# OpenAPI spec version: v0.0.1
# Contact: fredbi@yahoo.com
# Generated by: https://github.com/swagger-api/swagger-codegen.git


#' Zone Class
#'
#' @field id 
#' @field shortCode 
#' @field title 
#' @field descriptionShort 
#' @field descriptionLong 
#' @field auditTrail 
#' @field zoneType 
#' @field zoneGeometry 
#'
#' @importFrom R6 R6Class
#' @importFrom jsonlite fromJSON toJSON
#' @export
Zone <- R6::R6Class(
  'Zone',
  public = list(
    `id` = NULL,
    `shortCode` = NULL,
    `title` = NULL,
    `descriptionShort` = NULL,
    `descriptionLong` = NULL,
    `auditTrail` = NULL,
    `zoneType` = NULL,
    `zoneGeometry` = NULL,
    initialize = function(`id`, `shortCode`, `title`, `descriptionShort`, `descriptionLong`, `auditTrail`, `zoneType`, `zoneGeometry`){
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
      if (!missing(`zoneType`)) {
        stopifnot(R6::is.R6(`zoneType`))
        self$`zoneType` <- `zoneType`
      }
      if (!missing(`zoneGeometry`)) {
        stopifnot(R6::is.R6(`zoneGeometry`))
        self$`zoneGeometry` <- `zoneGeometry`
      }
    },
    toJSON = function() {
      ZoneObject <- list()
      if (!is.null(self$`id`)) {
        ZoneObject[['id']] <- self$`id`
      }
      if (!is.null(self$`shortCode`)) {
        ZoneObject[['shortCode']] <- self$`shortCode`$toJSON()
      }
      if (!is.null(self$`title`)) {
        ZoneObject[['title']] <- self$`title`$toJSON()
      }
      if (!is.null(self$`descriptionShort`)) {
        ZoneObject[['descriptionShort']] <- self$`descriptionShort`$toJSON()
      }
      if (!is.null(self$`descriptionLong`)) {
        ZoneObject[['descriptionLong']] <- self$`descriptionLong`$toJSON()
      }
      if (!is.null(self$`auditTrail`)) {
        ZoneObject[['auditTrail']] <- self$`auditTrail`$toJSON()
      }
      if (!is.null(self$`zoneType`)) {
        ZoneObject[['zoneType']] <- self$`zoneType`$toJSON()
      }
      if (!is.null(self$`zoneGeometry`)) {
        ZoneObject[['zoneGeometry']] <- self$`zoneGeometry`$toJSON()
      }

      ZoneObject
    },
    fromJSON = function(ZoneJson) {
      ZoneObject <- jsonlite::fromJSON(ZoneJson)
      if (!is.null(ZoneObject$`id`)) {
        self$`id` <- ZoneObject$`id`
      }
      if (!is.null(ZoneObject$`shortCode`)) {
        shortCodeObject <- ClassNomenclatureName$new()
        shortCodeObject$fromJSON(jsonlite::toJSON(ZoneObject$shortCode, auto_unbox = TRUE))
        self$`shortCode` <- shortCodeObject
      }
      if (!is.null(ZoneObject$`title`)) {
        titleObject <- Translation$new()
        titleObject$fromJSON(jsonlite::toJSON(ZoneObject$title, auto_unbox = TRUE))
        self$`title` <- titleObject
      }
      if (!is.null(ZoneObject$`descriptionShort`)) {
        descriptionShortObject <- Translation$new()
        descriptionShortObject$fromJSON(jsonlite::toJSON(ZoneObject$descriptionShort, auto_unbox = TRUE))
        self$`descriptionShort` <- descriptionShortObject
      }
      if (!is.null(ZoneObject$`descriptionLong`)) {
        descriptionLongObject <- Translation$new()
        descriptionLongObject$fromJSON(jsonlite::toJSON(ZoneObject$descriptionLong, auto_unbox = TRUE))
        self$`descriptionLong` <- descriptionLongObject
      }
      if (!is.null(ZoneObject$`auditTrail`)) {
        auditTrailObject <- Audit$new()
        auditTrailObject$fromJSON(jsonlite::toJSON(ZoneObject$auditTrail, auto_unbox = TRUE))
        self$`auditTrail` <- auditTrailObject
      }
      if (!is.null(ZoneObject$`zoneType`)) {
        zoneTypeObject <- Ztype$new()
        zoneTypeObject$fromJSON(jsonlite::toJSON(ZoneObject$zoneType, auto_unbox = TRUE))
        self$`zoneType` <- zoneTypeObject
      }
      if (!is.null(ZoneObject$`zoneGeometry`)) {
        zoneGeometryObject <- Geometry$new()
        zoneGeometryObject$fromJSON(jsonlite::toJSON(ZoneObject$zoneGeometry, auto_unbox = TRUE))
        self$`zoneGeometry` <- zoneGeometryObject
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
           "auditTrail": %s,
           "zoneType": %s,
           "zoneGeometry": %s
        }',
        self$`id`,
        self$`shortCode`$toJSON(),
        self$`title`$toJSON(),
        self$`descriptionShort`$toJSON(),
        self$`descriptionLong`$toJSON(),
        self$`auditTrail`$toJSON(),
        self$`zoneType`$toJSON(),
        self$`zoneGeometry`$toJSON()
      )
    },
    fromJSONString = function(ZoneJson) {
      ZoneObject <- jsonlite::fromJSON(ZoneJson)
      self$`id` <- ZoneObject$`id`
      ClassNomenclatureNameObject <- ClassNomenclatureName$new()
      self$`shortCode` <- ClassNomenclatureNameObject$fromJSON(jsonlite::toJSON(ZoneObject$shortCode, auto_unbox = TRUE))
      TranslationObject <- Translation$new()
      self$`title` <- TranslationObject$fromJSON(jsonlite::toJSON(ZoneObject$title, auto_unbox = TRUE))
      TranslationObject <- Translation$new()
      self$`descriptionShort` <- TranslationObject$fromJSON(jsonlite::toJSON(ZoneObject$descriptionShort, auto_unbox = TRUE))
      TranslationObject <- Translation$new()
      self$`descriptionLong` <- TranslationObject$fromJSON(jsonlite::toJSON(ZoneObject$descriptionLong, auto_unbox = TRUE))
      AuditObject <- Audit$new()
      self$`auditTrail` <- AuditObject$fromJSON(jsonlite::toJSON(ZoneObject$auditTrail, auto_unbox = TRUE))
      ZtypeObject <- Ztype$new()
      self$`zoneType` <- ZtypeObject$fromJSON(jsonlite::toJSON(ZoneObject$zoneType, auto_unbox = TRUE))
      GeometryObject <- Geometry$new()
      self$`zoneGeometry` <- GeometryObject$fromJSON(jsonlite::toJSON(ZoneObject$zoneGeometry, auto_unbox = TRUE))
    }
  )
)
