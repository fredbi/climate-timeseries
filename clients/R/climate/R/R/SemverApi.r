# Climate time series API
#
# The API allows the public to search and consult time series about climate change research conducted by TheShiftProject.  The climate time series API allows contributors to upload time series about their climate change models and studies.  Other secured endpoints allows admins to maintain the nomenclatures used by the climate time series, such as units etc. Timeseries publication status and ownership is for now managed by admins. 
#
# OpenAPI spec version: v0.0.1
# Contact: fredbi@yahoo.com
# Generated by: https://github.com/swagger-api/swagger-codegen.git

#' @title Semver operations
#' @description swagger.Semver
#'
#' @field path Stores url path of the request.
#' @field apiClient Handles the client-server communication.
#' @field userAgent Set the user agent of the request.
#'
#' @importFrom R6 R6Class
#'
#' @section Methods:
#' \describe{
#'
#' series_series_id_semver_get Get all semver tags associated to a series
#'
#'
#' series_series_id_semver_semver_get Get a version of a time series with a semver tag
#'
#'
#' series_series_id_semver_semver_values_get Get the values of version of a time series with a semver tag
#'
#' }
#'
#' @export
SemverApi <- R6::R6Class(
  'SemverApi',
  public = list(
    userAgent = "Swagger-Codegen/1.0.0/r",
    apiClient = NULL,
    initialize = function(apiClient){
      if (!missing(apiClient)) {
        self$apiClient <- apiClient
      }
      else {
        self$apiClient <- ApiClient$new()
      }
    },
    series_series_id_semver_get = function(series_id, ...){
      args <- list(...)
      queryParams <- list()
      headerParams <- character()

      urlPath <- "/series/{seriesId}/semver"
      if (!missing(`series_id`)) {
        urlPath <- gsub(paste0("\\{", "seriesId", "\\}"), `series_id`, urlPath)
      }

      resp <- self$apiClient$callApi(url = paste0(self$apiClient$basePath, urlPath),
                                 method = "GET",
                                 queryParams = queryParams,
                                 headerParams = headerParams,
                                 body = body,
                                 ...)
      
      if (httr::status_code(resp) >= 200 && httr::status_code(resp) <= 299) {
        returnObject <- Character$new()
        result <- returnObject$fromJSON(httr::content(resp, "text", encoding = "UTF-8"))
        Response$new(returnObject, resp)
      } else if (httr::status_code(resp) >= 400 && httr::status_code(resp) <= 499) {
        Response$new("API client error", resp)
      } else if (httr::status_code(resp) >= 500 && httr::status_code(resp) <= 599) {
        Response$new("API server error", resp)
      }

    },
    series_series_id_semver_semver_get = function(series_id, semver, deep, brief, audit, ...){
      args <- list(...)
      queryParams <- list()
      headerParams <- character()

      if (!missing(`deep`)) {
        queryParams['deep'] <- deep
      }

      if (!missing(`brief`)) {
        queryParams['brief'] <- brief
      }

      if (!missing(`audit`)) {
        queryParams['audit'] <- audit
      }

      urlPath <- "/series/{seriesId}/semver/{semver}"
      if (!missing(`series_id`)) {
        urlPath <- gsub(paste0("\\{", "seriesId", "\\}"), `series_id`, urlPath)
      }

      if (!missing(`semver`)) {
        urlPath <- gsub(paste0("\\{", "semver", "\\}"), `semver`, urlPath)
      }

      resp <- self$apiClient$callApi(url = paste0(self$apiClient$basePath, urlPath),
                                 method = "GET",
                                 queryParams = queryParams,
                                 headerParams = headerParams,
                                 body = body,
                                 ...)
      
      if (httr::status_code(resp) >= 200 && httr::status_code(resp) <= 299) {
        returnObject <- VersionedSeries$new()
        result <- returnObject$fromJSON(httr::content(resp, "text", encoding = "UTF-8"))
        Response$new(returnObject, resp)
      } else if (httr::status_code(resp) >= 400 && httr::status_code(resp) <= 499) {
        Response$new("API client error", resp)
      } else if (httr::status_code(resp) >= 500 && httr::status_code(resp) <= 599) {
        Response$new("API server error", resp)
      }

    },
    series_series_id_semver_semver_values_get = function(series_id, semver, from, to, convert, ...){
      args <- list(...)
      queryParams <- list()
      headerParams <- character()

      if (!missing(`from`)) {
        queryParams['from'] <- from
      }

      if (!missing(`to`)) {
        queryParams['to'] <- to
      }

      if (!missing(`convert`)) {
        queryParams['convert'] <- convert
      }

      urlPath <- "/series/{seriesId}/semver/{semver}/values"
      if (!missing(`series_id`)) {
        urlPath <- gsub(paste0("\\{", "seriesId", "\\}"), `series_id`, urlPath)
      }

      if (!missing(`semver`)) {
        urlPath <- gsub(paste0("\\{", "semver", "\\}"), `semver`, urlPath)
      }

      resp <- self$apiClient$callApi(url = paste0(self$apiClient$basePath, urlPath),
                                 method = "GET",
                                 queryParams = queryParams,
                                 headerParams = headerParams,
                                 body = body,
                                 ...)
      
      if (httr::status_code(resp) >= 200 && httr::status_code(resp) <= 299) {
        returnObject <- TsValues$new()
        result <- returnObject$fromJSON(httr::content(resp, "text", encoding = "UTF-8"))
        Response$new(returnObject, resp)
      } else if (httr::status_code(resp) >= 400 && httr::status_code(resp) <= 499) {
        Response$new("API client error", resp)
      } else if (httr::status_code(resp) >= 500 && httr::status_code(resp) <= 599) {
        Response$new("API server error", resp)
      }

    }
  )
)
