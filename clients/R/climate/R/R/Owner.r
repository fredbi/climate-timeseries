# Climate time series API
#
# The API allows the public to search and consult time series about climate change research conducted by TheShiftProject.  The climate time series API allows contributors to upload time series about their climate change models and studies.  Other secured endpoints allows admins to maintain the nomenclatures used by the climate time series, such as units etc. Timeseries publication status and ownership is for now managed by admins. 
#
# OpenAPI spec version: v0.0.1
# Contact: fredbi@yahoo.com
# Generated by: https://github.com/swagger-api/swagger-codegen.git


#' Owner Class
#'
#' @field id 
#' @field shortCode 
#' @field title 
#' @field descriptionShort 
#' @field descriptionLong 
#' @field auditTrail 
#' @field uuid 
#' @field name 
#' @field email 
#' @field ownerStatus 
#'
#' @importFrom R6 R6Class
#' @importFrom jsonlite fromJSON toJSON
#' @export
Owner <- R6::R6Class(
  'Owner',
  public = list(
    `id` = NULL,
    `shortCode` = NULL,
    `title` = NULL,
    `descriptionShort` = NULL,
    `descriptionLong` = NULL,
    `auditTrail` = NULL,
    `uuid` = NULL,
    `name` = NULL,
    `email` = NULL,
    `ownerStatus` = NULL,
    initialize = function(`id`, `shortCode`, `title`, `descriptionShort`, `descriptionLong`, `auditTrail`, `uuid`, `name`, `email`, `ownerStatus`){
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
      if (!missing(`uuid`)) {
        stopifnot(is.character(`uuid`), length(`uuid`) == 1)
        self$`uuid` <- `uuid`
      }
      if (!missing(`name`)) {
        stopifnot(is.character(`name`), length(`name`) == 1)
        self$`name` <- `name`
      }
      if (!missing(`email`)) {
        stopifnot(is.character(`email`), length(`email`) == 1)
        self$`email` <- `email`
      }
      if (!missing(`ownerStatus`)) {
        stopifnot(R6::is.R6(`ownerStatus`))
        self$`ownerStatus` <- `ownerStatus`
      }
    },
    toJSON = function() {
      OwnerObject <- list()
      if (!is.null(self$`id`)) {
        OwnerObject[['id']] <- self$`id`
      }
      if (!is.null(self$`shortCode`)) {
        OwnerObject[['shortCode']] <- self$`shortCode`$toJSON()
      }
      if (!is.null(self$`title`)) {
        OwnerObject[['title']] <- self$`title`$toJSON()
      }
      if (!is.null(self$`descriptionShort`)) {
        OwnerObject[['descriptionShort']] <- self$`descriptionShort`$toJSON()
      }
      if (!is.null(self$`descriptionLong`)) {
        OwnerObject[['descriptionLong']] <- self$`descriptionLong`$toJSON()
      }
      if (!is.null(self$`auditTrail`)) {
        OwnerObject[['auditTrail']] <- self$`auditTrail`$toJSON()
      }
      if (!is.null(self$`uuid`)) {
        OwnerObject[['uuid']] <- self$`uuid`
      }
      if (!is.null(self$`name`)) {
        OwnerObject[['name']] <- self$`name`
      }
      if (!is.null(self$`email`)) {
        OwnerObject[['email']] <- self$`email`
      }
      if (!is.null(self$`ownerStatus`)) {
        OwnerObject[['ownerStatus']] <- self$`ownerStatus`$toJSON()
      }

      OwnerObject
    },
    fromJSON = function(OwnerJson) {
      OwnerObject <- jsonlite::fromJSON(OwnerJson)
      if (!is.null(OwnerObject$`id`)) {
        self$`id` <- OwnerObject$`id`
      }
      if (!is.null(OwnerObject$`shortCode`)) {
        shortCodeObject <- ClassNomenclatureName$new()
        shortCodeObject$fromJSON(jsonlite::toJSON(OwnerObject$shortCode, auto_unbox = TRUE))
        self$`shortCode` <- shortCodeObject
      }
      if (!is.null(OwnerObject$`title`)) {
        titleObject <- Translation$new()
        titleObject$fromJSON(jsonlite::toJSON(OwnerObject$title, auto_unbox = TRUE))
        self$`title` <- titleObject
      }
      if (!is.null(OwnerObject$`descriptionShort`)) {
        descriptionShortObject <- Translation$new()
        descriptionShortObject$fromJSON(jsonlite::toJSON(OwnerObject$descriptionShort, auto_unbox = TRUE))
        self$`descriptionShort` <- descriptionShortObject
      }
      if (!is.null(OwnerObject$`descriptionLong`)) {
        descriptionLongObject <- Translation$new()
        descriptionLongObject$fromJSON(jsonlite::toJSON(OwnerObject$descriptionLong, auto_unbox = TRUE))
        self$`descriptionLong` <- descriptionLongObject
      }
      if (!is.null(OwnerObject$`auditTrail`)) {
        auditTrailObject <- Audit$new()
        auditTrailObject$fromJSON(jsonlite::toJSON(OwnerObject$auditTrail, auto_unbox = TRUE))
        self$`auditTrail` <- auditTrailObject
      }
      if (!is.null(OwnerObject$`uuid`)) {
        self$`uuid` <- OwnerObject$`uuid`
      }
      if (!is.null(OwnerObject$`name`)) {
        self$`name` <- OwnerObject$`name`
      }
      if (!is.null(OwnerObject$`email`)) {
        self$`email` <- OwnerObject$`email`
      }
      if (!is.null(OwnerObject$`ownerStatus`)) {
        ownerStatusObject <- Ostatus$new()
        ownerStatusObject$fromJSON(jsonlite::toJSON(OwnerObject$ownerStatus, auto_unbox = TRUE))
        self$`ownerStatus` <- ownerStatusObject
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
           "uuid": %s,
           "name": %s,
           "email": %s,
           "ownerStatus": %s
        }',
        self$`id`,
        self$`shortCode`$toJSON(),
        self$`title`$toJSON(),
        self$`descriptionShort`$toJSON(),
        self$`descriptionLong`$toJSON(),
        self$`auditTrail`$toJSON(),
        self$`uuid`,
        self$`name`,
        self$`email`,
        self$`ownerStatus`$toJSON()
      )
    },
    fromJSONString = function(OwnerJson) {
      OwnerObject <- jsonlite::fromJSON(OwnerJson)
      self$`id` <- OwnerObject$`id`
      ClassNomenclatureNameObject <- ClassNomenclatureName$new()
      self$`shortCode` <- ClassNomenclatureNameObject$fromJSON(jsonlite::toJSON(OwnerObject$shortCode, auto_unbox = TRUE))
      TranslationObject <- Translation$new()
      self$`title` <- TranslationObject$fromJSON(jsonlite::toJSON(OwnerObject$title, auto_unbox = TRUE))
      TranslationObject <- Translation$new()
      self$`descriptionShort` <- TranslationObject$fromJSON(jsonlite::toJSON(OwnerObject$descriptionShort, auto_unbox = TRUE))
      TranslationObject <- Translation$new()
      self$`descriptionLong` <- TranslationObject$fromJSON(jsonlite::toJSON(OwnerObject$descriptionLong, auto_unbox = TRUE))
      AuditObject <- Audit$new()
      self$`auditTrail` <- AuditObject$fromJSON(jsonlite::toJSON(OwnerObject$auditTrail, auto_unbox = TRUE))
      self$`uuid` <- OwnerObject$`uuid`
      self$`name` <- OwnerObject$`name`
      self$`email` <- OwnerObject$`email`
      OstatusObject <- Ostatus$new()
      self$`ownerStatus` <- OstatusObject$fromJSON(jsonlite::toJSON(OwnerObject$ownerStatus, auto_unbox = TRUE))
    }
  )
)
