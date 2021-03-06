# Climate time series API
#
# The API allows the public to search and consult time series about climate change research conducted by TheShiftProject.  The climate time series API allows contributors to upload time series about their climate change models and studies.  Other secured endpoints allows admins to maintain the nomenclatures used by the climate time series, such as units etc. Timeseries publication status and ownership is for now managed by admins. 
#
# OpenAPI spec version: v0.0.1
# Contact: fredbi@yahoo.com
# Generated by: https://github.com/swagger-api/swagger-codegen.git


#' Audit Class
#'
#' @field createdAt 
#' @field createdBy 
#' @field lastUpdatedAt 
#' @field lastUpdatedBy 
#'
#' @importFrom R6 R6Class
#' @importFrom jsonlite fromJSON toJSON
#' @export
Audit <- R6::R6Class(
  'Audit',
  public = list(
    `createdAt` = NULL,
    `createdBy` = NULL,
    `lastUpdatedAt` = NULL,
    `lastUpdatedBy` = NULL,
    initialize = function(`createdAt`, `createdBy`, `lastUpdatedAt`, `lastUpdatedBy`){
      if (!missing(`createdAt`)) {
        stopifnot(is.character(`createdAt`), length(`createdAt`) == 1)
        self$`createdAt` <- `createdAt`
      }
      if (!missing(`createdBy`)) {
        stopifnot(is.character(`createdBy`), length(`createdBy`) == 1)
        self$`createdBy` <- `createdBy`
      }
      if (!missing(`lastUpdatedAt`)) {
        stopifnot(is.character(`lastUpdatedAt`), length(`lastUpdatedAt`) == 1)
        self$`lastUpdatedAt` <- `lastUpdatedAt`
      }
      if (!missing(`lastUpdatedBy`)) {
        stopifnot(is.character(`lastUpdatedBy`), length(`lastUpdatedBy`) == 1)
        self$`lastUpdatedBy` <- `lastUpdatedBy`
      }
    },
    toJSON = function() {
      AuditObject <- list()
      if (!is.null(self$`createdAt`)) {
        AuditObject[['createdAt']] <- self$`createdAt`
      }
      if (!is.null(self$`createdBy`)) {
        AuditObject[['createdBy']] <- self$`createdBy`
      }
      if (!is.null(self$`lastUpdatedAt`)) {
        AuditObject[['lastUpdatedAt']] <- self$`lastUpdatedAt`
      }
      if (!is.null(self$`lastUpdatedBy`)) {
        AuditObject[['lastUpdatedBy']] <- self$`lastUpdatedBy`
      }

      AuditObject
    },
    fromJSON = function(AuditJson) {
      AuditObject <- jsonlite::fromJSON(AuditJson)
      if (!is.null(AuditObject$`createdAt`)) {
        self$`createdAt` <- AuditObject$`createdAt`
      }
      if (!is.null(AuditObject$`createdBy`)) {
        self$`createdBy` <- AuditObject$`createdBy`
      }
      if (!is.null(AuditObject$`lastUpdatedAt`)) {
        self$`lastUpdatedAt` <- AuditObject$`lastUpdatedAt`
      }
      if (!is.null(AuditObject$`lastUpdatedBy`)) {
        self$`lastUpdatedBy` <- AuditObject$`lastUpdatedBy`
      }
    },
    toJSONString = function() {
       sprintf(
        '{
           "createdAt": %s,
           "createdBy": %s,
           "lastUpdatedAt": %s,
           "lastUpdatedBy": %s
        }',
        self$`createdAt`,
        self$`createdBy`,
        self$`lastUpdatedAt`,
        self$`lastUpdatedBy`
      )
    },
    fromJSONString = function(AuditJson) {
      AuditObject <- jsonlite::fromJSON(AuditJson)
      self$`createdAt` <- AuditObject$`createdAt`
      self$`createdBy` <- AuditObject$`createdBy`
      self$`lastUpdatedAt` <- AuditObject$`lastUpdatedAt`
      self$`lastUpdatedBy` <- AuditObject$`lastUpdatedBy`
    }
  )
)
