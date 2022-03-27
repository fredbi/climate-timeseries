# Climate time series API
#
# The API allows the public to search and consult time series about climate change research conducted by TheShiftProject.  The climate time series API allows contributors to upload time series about their climate change models and studies.  Other secured endpoints allows admins to maintain the nomenclatures used by the climate time series, such as units etc. Timeseries publication status and ownership is for now managed by admins. 
#
# OpenAPI spec version: v0.0.1
# Contact: fredbi@yahoo.com
# Generated by: https://github.com/swagger-api/swagger-codegen.git


#' Source Class
#'
#' @field id 
#' @field shortCode 
#' @field title 
#' @field descriptionShort 
#' @field descriptionLong 
#' @field auditTrail 
#' @field rating 
#' @field tags 
#'
#' @importFrom R6 R6Class
#' @importFrom jsonlite fromJSON toJSON
#' @export
Source <- R6::R6Class(
  'Source',
  public = list(
    `id` = NULL,
    `shortCode` = NULL,
    `title` = NULL,
    `descriptionShort` = NULL,
    `descriptionLong` = NULL,
    `auditTrail` = NULL,
    `rating` = NULL,
    `tags` = NULL,
    initialize = function(`id`, `shortCode`, `title`, `descriptionShort`, `descriptionLong`, `auditTrail`, `rating`, `tags`){
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
      if (!missing(`rating`)) {
        stopifnot(is.numeric(`rating`), length(`rating`) == 1)
        self$`rating` <- `rating`
      }
      if (!missing(`tags`)) {
        stopifnot(is.list(`tags`), length(`tags`) != 0)
        lapply(`tags`, function(x) stopifnot(is.character(x)))
        self$`tags` <- `tags`
      }
    },
    toJSON = function() {
      SourceObject <- list()
      if (!is.null(self$`id`)) {
        SourceObject[['id']] <- self$`id`
      }
      if (!is.null(self$`shortCode`)) {
        SourceObject[['shortCode']] <- self$`shortCode`$toJSON()
      }
      if (!is.null(self$`title`)) {
        SourceObject[['title']] <- self$`title`$toJSON()
      }
      if (!is.null(self$`descriptionShort`)) {
        SourceObject[['descriptionShort']] <- self$`descriptionShort`$toJSON()
      }
      if (!is.null(self$`descriptionLong`)) {
        SourceObject[['descriptionLong']] <- self$`descriptionLong`$toJSON()
      }
      if (!is.null(self$`auditTrail`)) {
        SourceObject[['auditTrail']] <- self$`auditTrail`$toJSON()
      }
      if (!is.null(self$`rating`)) {
        SourceObject[['rating']] <- self$`rating`
      }
      if (!is.null(self$`tags`)) {
        SourceObject[['tags']] <- self$`tags`
      }

      SourceObject
    },
    fromJSON = function(SourceJson) {
      SourceObject <- jsonlite::fromJSON(SourceJson)
      if (!is.null(SourceObject$`id`)) {
        self$`id` <- SourceObject$`id`
      }
      if (!is.null(SourceObject$`shortCode`)) {
        shortCodeObject <- ClassNomenclatureName$new()
        shortCodeObject$fromJSON(jsonlite::toJSON(SourceObject$shortCode, auto_unbox = TRUE))
        self$`shortCode` <- shortCodeObject
      }
      if (!is.null(SourceObject$`title`)) {
        titleObject <- Translation$new()
        titleObject$fromJSON(jsonlite::toJSON(SourceObject$title, auto_unbox = TRUE))
        self$`title` <- titleObject
      }
      if (!is.null(SourceObject$`descriptionShort`)) {
        descriptionShortObject <- Translation$new()
        descriptionShortObject$fromJSON(jsonlite::toJSON(SourceObject$descriptionShort, auto_unbox = TRUE))
        self$`descriptionShort` <- descriptionShortObject
      }
      if (!is.null(SourceObject$`descriptionLong`)) {
        descriptionLongObject <- Translation$new()
        descriptionLongObject$fromJSON(jsonlite::toJSON(SourceObject$descriptionLong, auto_unbox = TRUE))
        self$`descriptionLong` <- descriptionLongObject
      }
      if (!is.null(SourceObject$`auditTrail`)) {
        auditTrailObject <- Audit$new()
        auditTrailObject$fromJSON(jsonlite::toJSON(SourceObject$auditTrail, auto_unbox = TRUE))
        self$`auditTrail` <- auditTrailObject
      }
      if (!is.null(SourceObject$`rating`)) {
        self$`rating` <- SourceObject$`rating`
      }
      if (!is.null(SourceObject$`tags`)) {
        self$`tags` <- SourceObject$`tags`
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
           "rating": %d,
           "tags": [%s]
        }',
        self$`id`,
        self$`shortCode`$toJSON(),
        self$`title`$toJSON(),
        self$`descriptionShort`$toJSON(),
        self$`descriptionLong`$toJSON(),
        self$`auditTrail`$toJSON(),
        self$`rating`,
        lapply(self$`tags`, function(x) paste(paste0('"', x, '"'), sep=","))
      )
    },
    fromJSONString = function(SourceJson) {
      SourceObject <- jsonlite::fromJSON(SourceJson)
      self$`id` <- SourceObject$`id`
      ClassNomenclatureNameObject <- ClassNomenclatureName$new()
      self$`shortCode` <- ClassNomenclatureNameObject$fromJSON(jsonlite::toJSON(SourceObject$shortCode, auto_unbox = TRUE))
      TranslationObject <- Translation$new()
      self$`title` <- TranslationObject$fromJSON(jsonlite::toJSON(SourceObject$title, auto_unbox = TRUE))
      TranslationObject <- Translation$new()
      self$`descriptionShort` <- TranslationObject$fromJSON(jsonlite::toJSON(SourceObject$descriptionShort, auto_unbox = TRUE))
      TranslationObject <- Translation$new()
      self$`descriptionLong` <- TranslationObject$fromJSON(jsonlite::toJSON(SourceObject$descriptionLong, auto_unbox = TRUE))
      AuditObject <- Audit$new()
      self$`auditTrail` <- AuditObject$fromJSON(jsonlite::toJSON(SourceObject$auditTrail, auto_unbox = TRUE))
      self$`rating` <- SourceObject$`rating`
      self$`tags` <- SourceObject$`tags`
    }
  )
)