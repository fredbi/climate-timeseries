# Climate time series API
#
# The API allows the public to search and consult time series about climate change research conducted by TheShiftProject.  The climate time series API allows contributors to upload time series about their climate change models and studies.  Other secured endpoints allows admins to maintain the nomenclatures used by the climate time series, such as units etc. Timeseries publication status and ownership is for now managed by admins. 
#
# OpenAPI spec version: v0.0.1
# Contact: fredbi@yahoo.com
# Generated by: https://github.com/swagger-api/swagger-codegen.git


#' Series Class
#'
#' @field id 
#' @field title 
#' @field descriptionShort 
#' @field descriptionLong 
#' @field timePeriod 
#' @field measurementUnit 
#' @field inputComposedUnit 
#' @field status 
#' @field statusChangeReason 
#' @field zone 
#' @field dataSource 
#' @field tags 
#' @field linkedDocuments 
#' @field auditTrail 
#' @field seriesHasThemes 
#' @field seriesHasOwners 
#' @field seriesHasVersions 
#'
#' @importFrom R6 R6Class
#' @importFrom jsonlite fromJSON toJSON
#' @export
Series <- R6::R6Class(
  'Series',
  public = list(
    `id` = NULL,
    `title` = NULL,
    `descriptionShort` = NULL,
    `descriptionLong` = NULL,
    `timePeriod` = NULL,
    `measurementUnit` = NULL,
    `inputComposedUnit` = NULL,
    `status` = NULL,
    `statusChangeReason` = NULL,
    `zone` = NULL,
    `dataSource` = NULL,
    `tags` = NULL,
    `linkedDocuments` = NULL,
    `auditTrail` = NULL,
    `seriesHasThemes` = NULL,
    `seriesHasOwners` = NULL,
    `seriesHasVersions` = NULL,
    initialize = function(`id`, `title`, `descriptionShort`, `descriptionLong`, `timePeriod`, `measurementUnit`, `inputComposedUnit`, `status`, `statusChangeReason`, `zone`, `dataSource`, `tags`, `linkedDocuments`, `auditTrail`, `seriesHasThemes`, `seriesHasOwners`, `seriesHasVersions`){
      if (!missing(`id`)) {
        stopifnot(is.numeric(`id`), length(`id`) == 1)
        self$`id` <- `id`
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
      if (!missing(`timePeriod`)) {
        stopifnot(R6::is.R6(`timePeriod`))
        self$`timePeriod` <- `timePeriod`
      }
      if (!missing(`measurementUnit`)) {
        stopifnot(R6::is.R6(`measurementUnit`))
        self$`measurementUnit` <- `measurementUnit`
      }
      if (!missing(`inputComposedUnit`)) {
        stopifnot(R6::is.R6(`inputComposedUnit`))
        self$`inputComposedUnit` <- `inputComposedUnit`
      }
      if (!missing(`status`)) {
        stopifnot(R6::is.R6(`status`))
        self$`status` <- `status`
      }
      if (!missing(`statusChangeReason`)) {
        stopifnot(R6::is.R6(`statusChangeReason`))
        self$`statusChangeReason` <- `statusChangeReason`
      }
      if (!missing(`zone`)) {
        stopifnot(R6::is.R6(`zone`))
        self$`zone` <- `zone`
      }
      if (!missing(`dataSource`)) {
        stopifnot(R6::is.R6(`dataSource`))
        self$`dataSource` <- `dataSource`
      }
      if (!missing(`tags`)) {
        stopifnot(R6::is.R6(`tags`))
        self$`tags` <- `tags`
      }
      if (!missing(`linkedDocuments`)) {
        stopifnot(R6::is.R6(`linkedDocuments`))
        self$`linkedDocuments` <- `linkedDocuments`
      }
      if (!missing(`auditTrail`)) {
        stopifnot(R6::is.R6(`auditTrail`))
        self$`auditTrail` <- `auditTrail`
      }
      if (!missing(`seriesHasThemes`)) {
        stopifnot(is.list(`seriesHasThemes`), length(`seriesHasThemes`) != 0)
        lapply(`seriesHasThemes`, function(x) stopifnot(R6::is.R6(x)))
        self$`seriesHasThemes` <- `seriesHasThemes`
      }
      if (!missing(`seriesHasOwners`)) {
        stopifnot(is.list(`seriesHasOwners`), length(`seriesHasOwners`) != 0)
        lapply(`seriesHasOwners`, function(x) stopifnot(R6::is.R6(x)))
        self$`seriesHasOwners` <- `seriesHasOwners`
      }
      if (!missing(`seriesHasVersions`)) {
        stopifnot(is.list(`seriesHasVersions`), length(`seriesHasVersions`) != 0)
        lapply(`seriesHasVersions`, function(x) stopifnot(R6::is.R6(x)))
        self$`seriesHasVersions` <- `seriesHasVersions`
      }
    },
    toJSON = function() {
      SeriesObject <- list()
      if (!is.null(self$`id`)) {
        SeriesObject[['id']] <- self$`id`
      }
      if (!is.null(self$`title`)) {
        SeriesObject[['title']] <- self$`title`$toJSON()
      }
      if (!is.null(self$`descriptionShort`)) {
        SeriesObject[['descriptionShort']] <- self$`descriptionShort`$toJSON()
      }
      if (!is.null(self$`descriptionLong`)) {
        SeriesObject[['descriptionLong']] <- self$`descriptionLong`$toJSON()
      }
      if (!is.null(self$`timePeriod`)) {
        SeriesObject[['timePeriod']] <- self$`timePeriod`$toJSON()
      }
      if (!is.null(self$`measurementUnit`)) {
        SeriesObject[['measurementUnit']] <- self$`measurementUnit`$toJSON()
      }
      if (!is.null(self$`inputComposedUnit`)) {
        SeriesObject[['inputComposedUnit']] <- self$`inputComposedUnit`$toJSON()
      }
      if (!is.null(self$`status`)) {
        SeriesObject[['status']] <- self$`status`$toJSON()
      }
      if (!is.null(self$`statusChangeReason`)) {
        SeriesObject[['statusChangeReason']] <- self$`statusChangeReason`$toJSON()
      }
      if (!is.null(self$`zone`)) {
        SeriesObject[['zone']] <- self$`zone`$toJSON()
      }
      if (!is.null(self$`dataSource`)) {
        SeriesObject[['dataSource']] <- self$`dataSource`$toJSON()
      }
      if (!is.null(self$`tags`)) {
        SeriesObject[['tags']] <- self$`tags`$toJSON()
      }
      if (!is.null(self$`linkedDocuments`)) {
        SeriesObject[['linkedDocuments']] <- self$`linkedDocuments`$toJSON()
      }
      if (!is.null(self$`auditTrail`)) {
        SeriesObject[['auditTrail']] <- self$`auditTrail`$toJSON()
      }
      if (!is.null(self$`seriesHasThemes`)) {
        SeriesObject[['seriesHasThemes']] <- lapply(self$`seriesHasThemes`, function(x) x$toJSON())
      }
      if (!is.null(self$`seriesHasOwners`)) {
        SeriesObject[['seriesHasOwners']] <- lapply(self$`seriesHasOwners`, function(x) x$toJSON())
      }
      if (!is.null(self$`seriesHasVersions`)) {
        SeriesObject[['seriesHasVersions']] <- lapply(self$`seriesHasVersions`, function(x) x$toJSON())
      }

      SeriesObject
    },
    fromJSON = function(SeriesJson) {
      SeriesObject <- jsonlite::fromJSON(SeriesJson)
      if (!is.null(SeriesObject$`id`)) {
        self$`id` <- SeriesObject$`id`
      }
      if (!is.null(SeriesObject$`title`)) {
        titleObject <- Translation$new()
        titleObject$fromJSON(jsonlite::toJSON(SeriesObject$title, auto_unbox = TRUE))
        self$`title` <- titleObject
      }
      if (!is.null(SeriesObject$`descriptionShort`)) {
        descriptionShortObject <- Translation$new()
        descriptionShortObject$fromJSON(jsonlite::toJSON(SeriesObject$descriptionShort, auto_unbox = TRUE))
        self$`descriptionShort` <- descriptionShortObject
      }
      if (!is.null(SeriesObject$`descriptionLong`)) {
        descriptionLongObject <- Translation$new()
        descriptionLongObject$fromJSON(jsonlite::toJSON(SeriesObject$descriptionLong, auto_unbox = TRUE))
        self$`descriptionLong` <- descriptionLongObject
      }
      if (!is.null(SeriesObject$`timePeriod`)) {
        timePeriodObject <- Period$new()
        timePeriodObject$fromJSON(jsonlite::toJSON(SeriesObject$timePeriod, auto_unbox = TRUE))
        self$`timePeriod` <- timePeriodObject
      }
      if (!is.null(SeriesObject$`measurementUnit`)) {
        measurementUnitObject <- Munit$new()
        measurementUnitObject$fromJSON(jsonlite::toJSON(SeriesObject$measurementUnit, auto_unbox = TRUE))
        self$`measurementUnit` <- measurementUnitObject
      }
      if (!is.null(SeriesObject$`inputComposedUnit`)) {
        inputComposedUnitObject <- CompositeUnit$new()
        inputComposedUnitObject$fromJSON(jsonlite::toJSON(SeriesObject$inputComposedUnit, auto_unbox = TRUE))
        self$`inputComposedUnit` <- inputComposedUnitObject
      }
      if (!is.null(SeriesObject$`status`)) {
        statusObject <- Vstatus$new()
        statusObject$fromJSON(jsonlite::toJSON(SeriesObject$status, auto_unbox = TRUE))
        self$`status` <- statusObject
      }
      if (!is.null(SeriesObject$`statusChangeReason`)) {
        statusChangeReasonObject <- Translation$new()
        statusChangeReasonObject$fromJSON(jsonlite::toJSON(SeriesObject$statusChangeReason, auto_unbox = TRUE))
        self$`statusChangeReason` <- statusChangeReasonObject
      }
      if (!is.null(SeriesObject$`zone`)) {
        zoneObject <- Zone$new()
        zoneObject$fromJSON(jsonlite::toJSON(SeriesObject$zone, auto_unbox = TRUE))
        self$`zone` <- zoneObject
      }
      if (!is.null(SeriesObject$`dataSource`)) {
        dataSourceObject <- Source$new()
        dataSourceObject$fromJSON(jsonlite::toJSON(SeriesObject$dataSource, auto_unbox = TRUE))
        self$`dataSource` <- dataSourceObject
      }
      if (!is.null(SeriesObject$`tags`)) {
        tagsObject <- SearchTags$new()
        tagsObject$fromJSON(jsonlite::toJSON(SeriesObject$tags, auto_unbox = TRUE))
        self$`tags` <- tagsObject
      }
      if (!is.null(SeriesObject$`linkedDocuments`)) {
        linkedDocumentsObject <- Documents$new()
        linkedDocumentsObject$fromJSON(jsonlite::toJSON(SeriesObject$linkedDocuments, auto_unbox = TRUE))
        self$`linkedDocuments` <- linkedDocumentsObject
      }
      if (!is.null(SeriesObject$`auditTrail`)) {
        auditTrailObject <- Audit$new()
        auditTrailObject$fromJSON(jsonlite::toJSON(SeriesObject$auditTrail, auto_unbox = TRUE))
        self$`auditTrail` <- auditTrailObject
      }
      if (!is.null(SeriesObject$`seriesHasThemes`)) {
        self$`seriesHasThemes` <- lapply(SeriesObject$`seriesHasThemes`, function(x) {
          seriesHasThemesObject <- Theme$new()
          seriesHasThemesObject$fromJSON(jsonlite::toJSON(x, auto_unbox = TRUE))
          seriesHasThemesObject
        })
      }
      if (!is.null(SeriesObject$`seriesHasOwners`)) {
        self$`seriesHasOwners` <- lapply(SeriesObject$`seriesHasOwners`, function(x) {
          seriesHasOwnersObject <- Owner$new()
          seriesHasOwnersObject$fromJSON(jsonlite::toJSON(x, auto_unbox = TRUE))
          seriesHasOwnersObject
        })
      }
      if (!is.null(SeriesObject$`seriesHasVersions`)) {
        self$`seriesHasVersions` <- lapply(SeriesObject$`seriesHasVersions`, function(x) {
          seriesHasVersionsObject <- VersionedSeries$new()
          seriesHasVersionsObject$fromJSON(jsonlite::toJSON(x, auto_unbox = TRUE))
          seriesHasVersionsObject
        })
      }
    },
    toJSONString = function() {
       sprintf(
        '{
           "id": %d,
           "title": %s,
           "descriptionShort": %s,
           "descriptionLong": %s,
           "timePeriod": %s,
           "measurementUnit": %s,
           "inputComposedUnit": %s,
           "status": %s,
           "statusChangeReason": %s,
           "zone": %s,
           "dataSource": %s,
           "tags": %s,
           "linkedDocuments": %s,
           "auditTrail": %s,
           "seriesHasThemes": [%s],
           "seriesHasOwners": [%s],
           "seriesHasVersions": [%s]
        }',
        self$`id`,
        self$`title`$toJSON(),
        self$`descriptionShort`$toJSON(),
        self$`descriptionLong`$toJSON(),
        self$`timePeriod`$toJSON(),
        self$`measurementUnit`$toJSON(),
        self$`inputComposedUnit`$toJSON(),
        self$`status`$toJSON(),
        self$`statusChangeReason`$toJSON(),
        self$`zone`$toJSON(),
        self$`dataSource`$toJSON(),
        self$`tags`$toJSON(),
        self$`linkedDocuments`$toJSON(),
        self$`auditTrail`$toJSON(),
        lapply(self$`seriesHasThemes`, function(x) paste(x$toJSON(), sep=",")),
        lapply(self$`seriesHasOwners`, function(x) paste(x$toJSON(), sep=",")),
        lapply(self$`seriesHasVersions`, function(x) paste(x$toJSON(), sep=","))
      )
    },
    fromJSONString = function(SeriesJson) {
      SeriesObject <- jsonlite::fromJSON(SeriesJson)
      self$`id` <- SeriesObject$`id`
      TranslationObject <- Translation$new()
      self$`title` <- TranslationObject$fromJSON(jsonlite::toJSON(SeriesObject$title, auto_unbox = TRUE))
      TranslationObject <- Translation$new()
      self$`descriptionShort` <- TranslationObject$fromJSON(jsonlite::toJSON(SeriesObject$descriptionShort, auto_unbox = TRUE))
      TranslationObject <- Translation$new()
      self$`descriptionLong` <- TranslationObject$fromJSON(jsonlite::toJSON(SeriesObject$descriptionLong, auto_unbox = TRUE))
      PeriodObject <- Period$new()
      self$`timePeriod` <- PeriodObject$fromJSON(jsonlite::toJSON(SeriesObject$timePeriod, auto_unbox = TRUE))
      MunitObject <- Munit$new()
      self$`measurementUnit` <- MunitObject$fromJSON(jsonlite::toJSON(SeriesObject$measurementUnit, auto_unbox = TRUE))
      CompositeUnitObject <- CompositeUnit$new()
      self$`inputComposedUnit` <- CompositeUnitObject$fromJSON(jsonlite::toJSON(SeriesObject$inputComposedUnit, auto_unbox = TRUE))
      VstatusObject <- Vstatus$new()
      self$`status` <- VstatusObject$fromJSON(jsonlite::toJSON(SeriesObject$status, auto_unbox = TRUE))
      TranslationObject <- Translation$new()
      self$`statusChangeReason` <- TranslationObject$fromJSON(jsonlite::toJSON(SeriesObject$statusChangeReason, auto_unbox = TRUE))
      ZoneObject <- Zone$new()
      self$`zone` <- ZoneObject$fromJSON(jsonlite::toJSON(SeriesObject$zone, auto_unbox = TRUE))
      SourceObject <- Source$new()
      self$`dataSource` <- SourceObject$fromJSON(jsonlite::toJSON(SeriesObject$dataSource, auto_unbox = TRUE))
      SearchTagsObject <- SearchTags$new()
      self$`tags` <- SearchTagsObject$fromJSON(jsonlite::toJSON(SeriesObject$tags, auto_unbox = TRUE))
      DocumentsObject <- Documents$new()
      self$`linkedDocuments` <- DocumentsObject$fromJSON(jsonlite::toJSON(SeriesObject$linkedDocuments, auto_unbox = TRUE))
      AuditObject <- Audit$new()
      self$`auditTrail` <- AuditObject$fromJSON(jsonlite::toJSON(SeriesObject$auditTrail, auto_unbox = TRUE))
      self$`seriesHasThemes` <- lapply(SeriesObject$`seriesHasThemes`, function(x) Theme$new()$fromJSON(jsonlite::toJSON(x, auto_unbox = TRUE)))
      self$`seriesHasOwners` <- lapply(SeriesObject$`seriesHasOwners`, function(x) Owner$new()$fromJSON(jsonlite::toJSON(x, auto_unbox = TRUE)))
      self$`seriesHasVersions` <- lapply(SeriesObject$`seriesHasVersions`, function(x) VersionedSeries$new()$fromJSON(jsonlite::toJSON(x, auto_unbox = TRUE)))
    }
  )
)
