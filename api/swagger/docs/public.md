


# Climate time series API
The timeseries API allows contributors to upload time series about their climate change models and studies.

The API allows the public to search and consult time series about climate change research.

  

## Informations

### Version

v0.0.1

### License

[Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0.html)

### Contact

Frédéric BIDON (Shifter, volunteer) fredbi@yahoo.com https://github.com/fredbi

### Terms Of Service

TODO


## Tags

  ### <span id="tag-classes"></span>[classes](https://github.io/wiki/TODO "API concepts: classes.\n")

Classes represent the various nomenclatures used by the time series, such as
measurement units and data sources.


  ### <span id="tag-series"></span>[series](https://github.io/wiki/TODO "API concepts: series.\n")

Series represent time series of measurements.


  ### <span id="tag-tags"></span>[tags](https://github.io/wiki/TODO#tags "API concepts: tags.\n")

Tags are keywords that may be used to search and classify series, data source and climate themes.


  ### <span id="tag-versioning"></span>[versioning](https://github.io/wiki/TODO#versions "API concepts: series versioning\n")

Timeseries are versioned.


  ### <span id="tag-semver"></span>[semver](https://github.io/wiki/TODO#semver "API concepts: semver.\n")

Semantic version tags ("semver") identify a checked-in time series.


  ### <span id="tag-conversions"></span>[conversions](https://github.io/wiki/TODO#unitconversions "API concepts: unit conversions.\n")

Unit conversion logic.


## Content negotiation

### URI Schemes
  * https

### Consumes
  * text/csv
  * application/json

### Produces
  * text/csv
  * application/json

## All endpoints

###  classes

  
> [API concepts: classes.](https://github.io/wiki/TODO)

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /v1/classes | [get classes](#get-classes) | Get all valid nomenclature classes |
| GET | /v1/classes/{classId} | [get classes class ID](#get-classes-class-id) | Get metadata about a nomenclature class |
| GET | /v1/classes/{classId}/members | [get classes class ID members](#get-classes-class-id-members) | Get all the members of a nomenclature class |
  


###  conversions

  
> [API concepts: unit conversions.](https://github.io/wiki/TODO#unitconversions)

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /v1/conversion | [get conversion](#get-conversion) | Get the conversion specification from unit to unit. |
| GET | /v1/conversion/{fromUnit}/{toUnit} | [get conversion from unit to unit](#get-conversion-from-unit-to-unit) | Get the conversion specification from unit to unit. |
  


###  series

  
> [API concepts: series.](https://github.io/wiki/TODO)

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| DELETE | /v1/series/{seriesId} | [delete series series ID](#delete-series-series-id) | Deletes a time series |
| DELETE | /v1/series/versions/{versionedSeriesId} | [delete series versions versioned series ID](#delete-series-versions-versioned-series-id) | Deletes a version of a time series |
| DELETE | /v1/series/versions/{versionedSeriesId}/values | [delete series versions versioned series ID values](#delete-series-versions-versioned-series-id-values) | Deletes the values attached to a version of a time series |
| GET | /v1/series | [get series](#get-series) | Get all available series, with search filters |
| GET | /v1/series/{seriesId} | [get series series ID](#get-series-series-id) | Get all versions of a time series. |
| GET | /v1/series/{seriesId}/latest | [get series series ID latest](#get-series-series-id-latest) | Get the latest version of the description of a time series |
| GET | /v1/series/{seriesId}/latest/values | [get series series ID latest values](#get-series-series-id-latest-values) | Get the values from the latest version of a time series |
| GET | /v1/series/{seriesId}/semver | [get series series ID semver](#get-series-series-id-semver) | Get all semver tags associated to a series |
| GET | /v1/series/{seriesId}/semver/{semver} | [get series series ID semver semver](#get-series-series-id-semver-semver) | Get a version of a time series with a semver tag |
| GET | /v1/series/{seriesId}/semver/{semver}/values | [get series series ID semver semver values](#get-series-series-id-semver-semver-values) | Get the values of version of a time series with a semver tag |
| GET | /v1/series/versions/{versionedSeriesId} | [get series versions versioned series ID](#get-series-versions-versioned-series-id) | Get a version of a time series |
| GET | /v1/series/versions/{versionedSeriesId}/values | [get series versions versioned series ID values](#get-series-versions-versioned-series-id-values) | Get the values of version of a time series |
| POST | /v1/series | [post series](#post-series) | Creates a new time series |
| POST | /v1/series/{seriesId} | [post series series ID](#post-series-series-id) | Creates a new version of a time series |
| POST | /v1/series/versions/{versionedSeriesId} | [post series versions versioned series ID](#post-series-versions-versioned-series-id) | creates a new version of a time series |
| POST | /v1/series/versions/{versionedSeriesId}/values | [post series versions versioned series ID values](#post-series-versions-versioned-series-id-values) | Uploads values to the version of a time series |
| PUT | /v1/series/{seriesId} | [put series series ID](#put-series-series-id) | Updates metadata about a time series |
| PUT | /v1/series/versions/{versionedSeriesId} | [put series versions versioned series ID](#put-series-versions-versioned-series-id) | Updates the metadata of version of a time series |
| PUT | /v1/series/versions/{versionedSeriesId}/values | [put series versions versioned series ID values](#put-series-versions-versioned-series-id-values) | Replaces the values of version of a time series |
  


###  tags

  
> [API concepts: tags.](https://github.io/wiki/TODO#tags)

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /v1/search/tags | [get search tags](#get-search-tags) | List all known tags |
| GET | /v1/search/tags/{tag} | [get search tags tag](#get-search-tags-tag) | Search all entities with some tag |
  


## Paths

### <span id="delete-series-series-id"></span> Deletes a time series (*DeleteSeriesSeriesID*)

```
DELETE /v1/series/{seriesId}
```

All versions and values associated to this series are deleted.

Only the original owner of a series may delete it.


#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| seriesId | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | The unique ID of a time series. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-series-series-id-204) | No Content | Series successfully deleted. |  | [schema](#delete-series-series-id-204-schema) |
| [400](#delete-series-series-id-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#delete-series-series-id-400-schema) |
| [401](#delete-series-series-id-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#delete-series-series-id-401-schema) |
| [403](#delete-series-series-id-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#delete-series-series-id-403-schema) |
| [404](#delete-series-series-id-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#delete-series-series-id-404-schema) |
| [405](#delete-series-series-id-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#delete-series-series-id-405-schema) |
| [500](#delete-series-series-id-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#delete-series-series-id-500-schema) |
| [default](#delete-series-series-id-default) | | Other error. See error details. |  | [schema](#delete-series-series-id-default-schema) |

#### Responses


##### <span id="delete-series-series-id-204"></span> 204 - Series successfully deleted.
Status: No Content

###### <span id="delete-series-series-id-204-schema"></span> Schema

##### <span id="delete-series-series-id-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="delete-series-series-id-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="delete-series-series-id-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="delete-series-series-id-401-schema"></span> Schema

##### <span id="delete-series-series-id-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="delete-series-series-id-403-schema"></span> Schema

##### <span id="delete-series-series-id-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="delete-series-series-id-404-schema"></span> Schema

##### <span id="delete-series-series-id-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="delete-series-series-id-405-schema"></span> Schema

##### <span id="delete-series-series-id-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="delete-series-series-id-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="delete-series-series-id-default"></span> Default Response
Other error. See error details.

###### <span id="delete-series-series-id-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="delete-series-versions-versioned-series-id"></span> Deletes a version of a time series (*DeleteSeriesVersionsVersionedSeriesID*)

```
DELETE /v1/series/versions/{versionedSeriesId}
```

Only the original owner of this version of the series may delete it.

All values associated to that version are deleted.


#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| versionedSeriesId | `path` | uuid (formatted string) | `strfmt.UUID` |  | ✓ |  | The UUID of a version of the time series. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-series-versions-versioned-series-id-204) | No Content | Version of the time series successfully deleted. |  | [schema](#delete-series-versions-versioned-series-id-204-schema) |
| [400](#delete-series-versions-versioned-series-id-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#delete-series-versions-versioned-series-id-400-schema) |
| [401](#delete-series-versions-versioned-series-id-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#delete-series-versions-versioned-series-id-401-schema) |
| [403](#delete-series-versions-versioned-series-id-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#delete-series-versions-versioned-series-id-403-schema) |
| [404](#delete-series-versions-versioned-series-id-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#delete-series-versions-versioned-series-id-404-schema) |
| [405](#delete-series-versions-versioned-series-id-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#delete-series-versions-versioned-series-id-405-schema) |
| [500](#delete-series-versions-versioned-series-id-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#delete-series-versions-versioned-series-id-500-schema) |
| [default](#delete-series-versions-versioned-series-id-default) | | Other error. See error details. |  | [schema](#delete-series-versions-versioned-series-id-default-schema) |

#### Responses


##### <span id="delete-series-versions-versioned-series-id-204"></span> 204 - Version of the time series successfully deleted.
Status: No Content

###### <span id="delete-series-versions-versioned-series-id-204-schema"></span> Schema

##### <span id="delete-series-versions-versioned-series-id-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="delete-series-versions-versioned-series-id-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="delete-series-versions-versioned-series-id-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="delete-series-versions-versioned-series-id-401-schema"></span> Schema

##### <span id="delete-series-versions-versioned-series-id-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="delete-series-versions-versioned-series-id-403-schema"></span> Schema

##### <span id="delete-series-versions-versioned-series-id-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="delete-series-versions-versioned-series-id-404-schema"></span> Schema

##### <span id="delete-series-versions-versioned-series-id-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="delete-series-versions-versioned-series-id-405-schema"></span> Schema

##### <span id="delete-series-versions-versioned-series-id-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="delete-series-versions-versioned-series-id-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="delete-series-versions-versioned-series-id-default"></span> Default Response
Other error. See error details.

###### <span id="delete-series-versions-versioned-series-id-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="delete-series-versions-versioned-series-id-values"></span> Deletes the values attached to a version of a time series (*DeleteSeriesVersionsVersionedSeriesIDValues*)

```
DELETE /v1/series/versions/{versionedSeriesId}/values
```

Only the original owner of this version of the series may delete it.

The version is not deleted. Only the values associated to that version are deleted.


#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| versionedSeriesId | `path` | uuid (formatted string) | `strfmt.UUID` |  | ✓ |  | The UUID of a version of the time series. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-series-versions-versioned-series-id-values-204) | No Content | Values of the versioned time series successfully deleted. |  | [schema](#delete-series-versions-versioned-series-id-values-204-schema) |
| [400](#delete-series-versions-versioned-series-id-values-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#delete-series-versions-versioned-series-id-values-400-schema) |
| [401](#delete-series-versions-versioned-series-id-values-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#delete-series-versions-versioned-series-id-values-401-schema) |
| [403](#delete-series-versions-versioned-series-id-values-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#delete-series-versions-versioned-series-id-values-403-schema) |
| [404](#delete-series-versions-versioned-series-id-values-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#delete-series-versions-versioned-series-id-values-404-schema) |
| [405](#delete-series-versions-versioned-series-id-values-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#delete-series-versions-versioned-series-id-values-405-schema) |
| [500](#delete-series-versions-versioned-series-id-values-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#delete-series-versions-versioned-series-id-values-500-schema) |
| [default](#delete-series-versions-versioned-series-id-values-default) | | Other error. See error details. |  | [schema](#delete-series-versions-versioned-series-id-values-default-schema) |

#### Responses


##### <span id="delete-series-versions-versioned-series-id-values-204"></span> 204 - Values of the versioned time series successfully deleted.
Status: No Content

###### <span id="delete-series-versions-versioned-series-id-values-204-schema"></span> Schema

##### <span id="delete-series-versions-versioned-series-id-values-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="delete-series-versions-versioned-series-id-values-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="delete-series-versions-versioned-series-id-values-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="delete-series-versions-versioned-series-id-values-401-schema"></span> Schema

##### <span id="delete-series-versions-versioned-series-id-values-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="delete-series-versions-versioned-series-id-values-403-schema"></span> Schema

##### <span id="delete-series-versions-versioned-series-id-values-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="delete-series-versions-versioned-series-id-values-404-schema"></span> Schema

##### <span id="delete-series-versions-versioned-series-id-values-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="delete-series-versions-versioned-series-id-values-405-schema"></span> Schema

##### <span id="delete-series-versions-versioned-series-id-values-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="delete-series-versions-versioned-series-id-values-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="delete-series-versions-versioned-series-id-values-default"></span> Default Response
Other error. See error details.

###### <span id="delete-series-versions-versioned-series-id-values-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="get-classes"></span> Get all valid nomenclature classes (*GetClasses*)

```
GET /v1/classes
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| audit | `query` | boolean | `bool` |  |  |  | When audit is specified, the response will also contain the audit trail field. |
| brief | `query` | boolean | `bool` |  |  |  | When brief is specified, the response will only contain essential data and strip long descriptions. |
| deep | `query` | boolean | `bool` |  |  |  | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description. |
| glob | `query` | string | `string` |  |  |  | Filters the result according to a glob pattern applied on the short name of the requested object. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-classes-200) | OK | A collection of nomenclature classes. |  | [schema](#get-classes-200-schema) |
| [400](#get-classes-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#get-classes-400-schema) |
| [401](#get-classes-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#get-classes-401-schema) |
| [403](#get-classes-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#get-classes-403-schema) |
| [404](#get-classes-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#get-classes-404-schema) |
| [405](#get-classes-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#get-classes-405-schema) |
| [500](#get-classes-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#get-classes-500-schema) |
| [default](#get-classes-default) | | Other error. See error details. |  | [schema](#get-classes-default-schema) |

#### Responses


##### <span id="get-classes-200"></span> 200 - A collection of nomenclature classes.
Status: OK

###### <span id="get-classes-200-schema"></span> Schema
   
  

[][Class](#class)

##### <span id="get-classes-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="get-classes-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-classes-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="get-classes-401-schema"></span> Schema

##### <span id="get-classes-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="get-classes-403-schema"></span> Schema

##### <span id="get-classes-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="get-classes-404-schema"></span> Schema

##### <span id="get-classes-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="get-classes-405-schema"></span> Schema

##### <span id="get-classes-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="get-classes-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-classes-default"></span> Default Response
Other error. See error details.

###### <span id="get-classes-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="get-classes-class-id"></span> Get metadata about a nomenclature class (*GetClassesClassID*)

```
GET /v1/classes/{classId}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| classId | `path` | string | `string` |  | ✓ |  | The internal name of a nomenclature class.

Valid classes are:
  * constant: mathematical and physical constants
  * mdimension: base measured dimensions
  * mdomain: domains that pertain to measurements
  * measurement: physical and economic measurements
  * multiplier: unit multipliers aka prefixes (e.g. k,M,G...)
  * munit: measurement units
  * musystem: systems of measurement units
  * ostatus: owner statuses
  * owner: series owner
  * period: time series periods (e.g. monthly, yearly...)
  * role: series owner role
  * source: data sources
  * status: series and versions statuses
  * theme: climate themes
  * zone: geographical zones
  * ztype: zone types |
| audit | `query` | boolean | `bool` |  |  |  | When audit is specified, the response will also contain the audit trail field. |
| brief | `query` | boolean | `bool` |  |  |  | When brief is specified, the response will only contain essential data and strip long descriptions. |
| deep | `query` | boolean | `bool` |  |  |  | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-classes-class-id-200) | OK | Descriptive metadata about nomenclature classes. |  | [schema](#get-classes-class-id-200-schema) |
| [400](#get-classes-class-id-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#get-classes-class-id-400-schema) |
| [401](#get-classes-class-id-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#get-classes-class-id-401-schema) |
| [403](#get-classes-class-id-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#get-classes-class-id-403-schema) |
| [404](#get-classes-class-id-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#get-classes-class-id-404-schema) |
| [405](#get-classes-class-id-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#get-classes-class-id-405-schema) |
| [500](#get-classes-class-id-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#get-classes-class-id-500-schema) |
| [default](#get-classes-class-id-default) | | Other error. See error details. |  | [schema](#get-classes-class-id-default-schema) |

#### Responses


##### <span id="get-classes-class-id-200"></span> 200 - Descriptive metadata about nomenclature classes.
Status: OK

###### <span id="get-classes-class-id-200-schema"></span> Schema
   
  

[ClassDescription](#class-description)

##### <span id="get-classes-class-id-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="get-classes-class-id-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-classes-class-id-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="get-classes-class-id-401-schema"></span> Schema

##### <span id="get-classes-class-id-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="get-classes-class-id-403-schema"></span> Schema

##### <span id="get-classes-class-id-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="get-classes-class-id-404-schema"></span> Schema

##### <span id="get-classes-class-id-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="get-classes-class-id-405-schema"></span> Schema

##### <span id="get-classes-class-id-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="get-classes-class-id-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-classes-class-id-default"></span> Default Response
Other error. See error details.

###### <span id="get-classes-class-id-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="get-classes-class-id-members"></span> Get all the members of a nomenclature class (*GetClassesClassIDMembers*)

```
GET /v1/classes/{classId}/members
```

Returns a a collection of nomenclature class members.

Example:
  GET /classes/munit/members returns all measurement units.

If the request negotiates a response MIME type with text/csv (with the Accept header),
this endpoint produces a CSV file.


> [Read more](https://github.io/wiki/TODO#classcsv "API responses: nomenclature values
")

#### Produces
  * application/json
  * text/csv

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| classId | `path` | string | `string` |  | ✓ |  | The internal name of a nomenclature class.

Valid classes are:
  * constant: mathematical and physical constants
  * mdimension: base measured dimensions
  * mdomain: domains that pertain to measurements
  * measurement: physical and economic measurements
  * multiplier: unit multipliers aka prefixes (e.g. k,M,G...)
  * munit: measurement units
  * musystem: systems of measurement units
  * ostatus: owner statuses
  * owner: series owner
  * period: time series periods (e.g. monthly, yearly...)
  * role: series owner role
  * source: data sources
  * status: series and versions statuses
  * theme: climate themes
  * zone: geographical zones
  * ztype: zone types |
| audit | `query` | boolean | `bool` |  |  |  | When audit is specified, the response will also contain the audit trail field. |
| brief | `query` | boolean | `bool` |  |  |  | When brief is specified, the response will only contain essential data and strip long descriptions. |
| deep | `query` | boolean | `bool` |  |  |  | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description. |
| glob | `query` | string | `string` |  |  |  | Filters the result according to a glob pattern applied on the short name of the requested object. |
| tag | `query` | string | `string` |  |  |  | Filter the result by search for a given keyword, whenever tag search is applicable.

This parameter has no effect on objects where tag search is not applicable. |
| tags | `query` | []string | `[]string` | `csv` |  |  | Filter the result by search for one tag in the given list of keywords, whenever tag search is applicable.

Up to 5 tags can be searched.

This parameter has no effect on objects where tag search is not applicable. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-classes-class-id-members-200) | OK | Nomenclature class members. |  | [schema](#get-classes-class-id-members-200-schema) |
| [400](#get-classes-class-id-members-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#get-classes-class-id-members-400-schema) |
| [401](#get-classes-class-id-members-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#get-classes-class-id-members-401-schema) |
| [403](#get-classes-class-id-members-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#get-classes-class-id-members-403-schema) |
| [404](#get-classes-class-id-members-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#get-classes-class-id-members-404-schema) |
| [405](#get-classes-class-id-members-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#get-classes-class-id-members-405-schema) |
| [500](#get-classes-class-id-members-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#get-classes-class-id-members-500-schema) |
| [default](#get-classes-class-id-members-default) | | Other error. See error details. |  | [schema](#get-classes-class-id-members-default-schema) |

#### Responses


##### <span id="get-classes-class-id-members-200"></span> 200 - Nomenclature class members.
Status: OK

###### <span id="get-classes-class-id-members-200-schema"></span> Schema
   
  

[][Class](#class)

##### <span id="get-classes-class-id-members-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="get-classes-class-id-members-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-classes-class-id-members-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="get-classes-class-id-members-401-schema"></span> Schema

##### <span id="get-classes-class-id-members-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="get-classes-class-id-members-403-schema"></span> Schema

##### <span id="get-classes-class-id-members-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="get-classes-class-id-members-404-schema"></span> Schema

##### <span id="get-classes-class-id-members-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="get-classes-class-id-members-405-schema"></span> Schema

##### <span id="get-classes-class-id-members-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="get-classes-class-id-members-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-classes-class-id-members-default"></span> Default Response
Other error. See error details.

###### <span id="get-classes-class-id-members-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="get-conversion"></span> Get the conversion specification from unit to unit. (*GetConversion*)

```
GET /v1/conversion
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| audit | `query` | boolean | `bool` |  |  |  | When audit is specified, the response will also contain the audit trail field. |
| brief | `query` | boolean | `bool` |  |  |  | When brief is specified, the response will only contain essential data and strip long descriptions. |
| deep | `query` | boolean | `bool` |  |  |  | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description. |
| fromUnit | `query` | string | `string` |  | ✓ |  | Original unit to be converted. |
| toUnit | `query` | string | `string` |  | ✓ |  | Original unit to be converted. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-conversion-200) | OK | A unit conversion specification |  | [schema](#get-conversion-200-schema) |
| [400](#get-conversion-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#get-conversion-400-schema) |
| [401](#get-conversion-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#get-conversion-401-schema) |
| [403](#get-conversion-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#get-conversion-403-schema) |
| [404](#get-conversion-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#get-conversion-404-schema) |
| [405](#get-conversion-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#get-conversion-405-schema) |
| [500](#get-conversion-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#get-conversion-500-schema) |
| [default](#get-conversion-default) | | Other error. See error details. |  | [schema](#get-conversion-default-schema) |

#### Responses


##### <span id="get-conversion-200"></span> 200 - A unit conversion specification
Status: OK

###### <span id="get-conversion-200-schema"></span> Schema
   
  

[Conversion](#conversion)

##### <span id="get-conversion-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="get-conversion-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-conversion-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="get-conversion-401-schema"></span> Schema

##### <span id="get-conversion-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="get-conversion-403-schema"></span> Schema

##### <span id="get-conversion-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="get-conversion-404-schema"></span> Schema

##### <span id="get-conversion-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="get-conversion-405-schema"></span> Schema

##### <span id="get-conversion-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="get-conversion-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-conversion-default"></span> Default Response
Other error. See error details.

###### <span id="get-conversion-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="get-conversion-from-unit-to-unit"></span> Get the conversion specification from unit to unit. (*GetConversionFromUnitToUnit*)

```
GET /v1/conversion/{fromUnit}/{toUnit}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| fromUnit | `path` | string | `string` |  | ✓ |  | Original unit to be converted. |
| toUnit | `path` | string | `string` |  | ✓ |  | Original unit to be converted. |
| audit | `query` | boolean | `bool` |  |  |  | When audit is specified, the response will also contain the audit trail field. |
| brief | `query` | boolean | `bool` |  |  |  | When brief is specified, the response will only contain essential data and strip long descriptions. |
| deep | `query` | boolean | `bool` |  |  |  | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-conversion-from-unit-to-unit-200) | OK | A unit conversion specification |  | [schema](#get-conversion-from-unit-to-unit-200-schema) |
| [400](#get-conversion-from-unit-to-unit-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#get-conversion-from-unit-to-unit-400-schema) |
| [401](#get-conversion-from-unit-to-unit-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#get-conversion-from-unit-to-unit-401-schema) |
| [403](#get-conversion-from-unit-to-unit-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#get-conversion-from-unit-to-unit-403-schema) |
| [404](#get-conversion-from-unit-to-unit-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#get-conversion-from-unit-to-unit-404-schema) |
| [405](#get-conversion-from-unit-to-unit-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#get-conversion-from-unit-to-unit-405-schema) |
| [500](#get-conversion-from-unit-to-unit-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#get-conversion-from-unit-to-unit-500-schema) |
| [default](#get-conversion-from-unit-to-unit-default) | | Other error. See error details. |  | [schema](#get-conversion-from-unit-to-unit-default-schema) |

#### Responses


##### <span id="get-conversion-from-unit-to-unit-200"></span> 200 - A unit conversion specification
Status: OK

###### <span id="get-conversion-from-unit-to-unit-200-schema"></span> Schema
   
  

[ConversionSpec](#conversion-spec)

##### <span id="get-conversion-from-unit-to-unit-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="get-conversion-from-unit-to-unit-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-conversion-from-unit-to-unit-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="get-conversion-from-unit-to-unit-401-schema"></span> Schema

##### <span id="get-conversion-from-unit-to-unit-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="get-conversion-from-unit-to-unit-403-schema"></span> Schema

##### <span id="get-conversion-from-unit-to-unit-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="get-conversion-from-unit-to-unit-404-schema"></span> Schema

##### <span id="get-conversion-from-unit-to-unit-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="get-conversion-from-unit-to-unit-405-schema"></span> Schema

##### <span id="get-conversion-from-unit-to-unit-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="get-conversion-from-unit-to-unit-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-conversion-from-unit-to-unit-default"></span> Default Response
Other error. See error details.

###### <span id="get-conversion-from-unit-to-unit-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="get-search-tags"></span> List all known tags (*GetSearchTags*)

```
GET /v1/search/tags
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-search-tags-200) | OK | All known search tags. |  | [schema](#get-search-tags-200-schema) |
| [400](#get-search-tags-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#get-search-tags-400-schema) |
| [401](#get-search-tags-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#get-search-tags-401-schema) |
| [403](#get-search-tags-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#get-search-tags-403-schema) |
| [404](#get-search-tags-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#get-search-tags-404-schema) |
| [405](#get-search-tags-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#get-search-tags-405-schema) |
| [500](#get-search-tags-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#get-search-tags-500-schema) |
| [default](#get-search-tags-default) | | Other error. See error details. |  | [schema](#get-search-tags-default-schema) |

#### Responses


##### <span id="get-search-tags-200"></span> 200 - All known search tags.
Status: OK

###### <span id="get-search-tags-200-schema"></span> Schema
   
  

[]string

##### <span id="get-search-tags-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="get-search-tags-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-search-tags-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="get-search-tags-401-schema"></span> Schema

##### <span id="get-search-tags-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="get-search-tags-403-schema"></span> Schema

##### <span id="get-search-tags-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="get-search-tags-404-schema"></span> Schema

##### <span id="get-search-tags-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="get-search-tags-405-schema"></span> Schema

##### <span id="get-search-tags-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="get-search-tags-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-search-tags-default"></span> Default Response
Other error. See error details.

###### <span id="get-search-tags-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="get-search-tags-tag"></span> Search all entities with some tag (*GetSearchTagsTag*)

```
GET /v1/search/tags/{tag}
```

This endpoint returns mulitple entities, such as classes or series.


#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| tag | `path` | string | `string` |  | ✓ |  | Filter the result by search for a given keyword, whenever tag search is applicable.

This parameter has no effect on objects where tag search is not applicable. |
| audit | `query` | boolean | `bool` |  |  |  | When audit is specified, the response will also contain the audit trail field. |
| brief | `query` | boolean | `bool` |  |  |  | When brief is specified, the response will only contain essential data and strip long descriptions. |
| deep | `query` | boolean | `bool` |  |  |  | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description. |
| email | `query` | string | `string` |  |  |  | Refers to some data owner by email. |
| ownerId | `query` | uuid (formatted string) | `strfmt.UUID` |  |  |  | The ID of the contributor. |
| status | `query` | string | `string` |  |  |  | Filter the result according to a given status.

Only series in the "PUBLISHED" status are returned to the public.

Series owner may consult their series in any status. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-search-tags-tag-200) | OK | All entities that match the searched tag. |  | [schema](#get-search-tags-tag-200-schema) |
| [400](#get-search-tags-tag-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#get-search-tags-tag-400-schema) |
| [401](#get-search-tags-tag-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#get-search-tags-tag-401-schema) |
| [403](#get-search-tags-tag-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#get-search-tags-tag-403-schema) |
| [404](#get-search-tags-tag-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#get-search-tags-tag-404-schema) |
| [405](#get-search-tags-tag-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#get-search-tags-tag-405-schema) |
| [500](#get-search-tags-tag-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#get-search-tags-tag-500-schema) |
| [default](#get-search-tags-tag-default) | | Other error. See error details. |  | [schema](#get-search-tags-tag-default-schema) |

#### Responses


##### <span id="get-search-tags-tag-200"></span> 200 - All entities that match the searched tag.
Status: OK

###### <span id="get-search-tags-tag-200-schema"></span> Schema
   
  


 [Entities](#entities)

##### <span id="get-search-tags-tag-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="get-search-tags-tag-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-search-tags-tag-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="get-search-tags-tag-401-schema"></span> Schema

##### <span id="get-search-tags-tag-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="get-search-tags-tag-403-schema"></span> Schema

##### <span id="get-search-tags-tag-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="get-search-tags-tag-404-schema"></span> Schema

##### <span id="get-search-tags-tag-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="get-search-tags-tag-405-schema"></span> Schema

##### <span id="get-search-tags-tag-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="get-search-tags-tag-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-search-tags-tag-default"></span> Default Response
Other error. See error details.

###### <span id="get-search-tags-tag-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="get-series"></span> Get all available series, with search filters (*GetSeries*)

```
GET /v1/series
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| datasource | `query` | string | `string` |  |  |  | A datasource short code to be searched for. |
| email | `query` | string | `string` |  |  |  | Refers to some data owner by email. |
| glob | `query` | string | `string` |  |  |  | Filters the result according to a glob pattern applied on the short name of the requested object. |
| ownerId | `query` | uuid (formatted string) | `strfmt.UUID` |  |  |  | The ID of the contributor. |
| status | `query` | string | `string` |  |  |  | Filter the result according to a given status.

Only series in the "PUBLISHED" status are returned to the public.

Series owner may consult their series in any status. |
| tag | `query` | string | `string` |  |  |  | Filter the result by search for a given keyword, whenever tag search is applicable.

This parameter has no effect on objects where tag search is not applicable. |
| tags | `query` | []string | `[]string` | `csv` |  |  | Filter the result by search for one tag in the given list of keywords, whenever tag search is applicable.

Up to 5 tags can be searched.

This parameter has no effect on objects where tag search is not applicable. |
| theme | `query` | string | `string` |  |  |  | A theme path (url-encoded) to be searched for. |
| zone | `query` | string | `string` |  |  |  | A zone short code (url-encoded) to be searched for. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-series-200) | OK | Available climate timeseries. |  | [schema](#get-series-200-schema) |
| [400](#get-series-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#get-series-400-schema) |
| [401](#get-series-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#get-series-401-schema) |
| [403](#get-series-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#get-series-403-schema) |
| [404](#get-series-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#get-series-404-schema) |
| [405](#get-series-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#get-series-405-schema) |
| [500](#get-series-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#get-series-500-schema) |
| [default](#get-series-default) | | Other error. See error details. |  | [schema](#get-series-default-schema) |

#### Responses


##### <span id="get-series-200"></span> 200 - Available climate timeseries.
Status: OK

###### <span id="get-series-200-schema"></span> Schema
   
  

[][Series](#series)

##### <span id="get-series-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="get-series-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="get-series-401-schema"></span> Schema

##### <span id="get-series-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="get-series-403-schema"></span> Schema

##### <span id="get-series-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="get-series-404-schema"></span> Schema

##### <span id="get-series-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="get-series-405-schema"></span> Schema

##### <span id="get-series-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="get-series-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-default"></span> Default Response
Other error. See error details.

###### <span id="get-series-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="get-series-series-id"></span> Get all versions of a time series. (*GetSeriesSeriesID*)

```
GET /v1/series/{seriesId}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| seriesId | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | The unique ID of a time series. |
| audit | `query` | boolean | `bool` |  |  |  | When audit is specified, the response will also contain the audit trail field. |
| brief | `query` | boolean | `bool` |  |  |  | When brief is specified, the response will only contain essential data and strip long descriptions. |
| deep | `query` | boolean | `bool` |  |  |  | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description. |
| glob | `query` | string | `string` |  |  |  | Filters the result according to a glob pattern applied on the short name of the requested object. |
| status | `query` | string | `string` |  |  |  | Filter the result according to a given status.

Only series in the "PUBLISHED" status are returned to the public.

Series owner may consult their series in any status. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-series-series-id-200) | OK | All available versions of a climate timeseries. |  | [schema](#get-series-series-id-200-schema) |
| [400](#get-series-series-id-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#get-series-series-id-400-schema) |
| [401](#get-series-series-id-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#get-series-series-id-401-schema) |
| [403](#get-series-series-id-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#get-series-series-id-403-schema) |
| [404](#get-series-series-id-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#get-series-series-id-404-schema) |
| [405](#get-series-series-id-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#get-series-series-id-405-schema) |
| [500](#get-series-series-id-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#get-series-series-id-500-schema) |
| [default](#get-series-series-id-default) | | Other error. See error details. |  | [schema](#get-series-series-id-default-schema) |

#### Responses


##### <span id="get-series-series-id-200"></span> 200 - All available versions of a climate timeseries.
Status: OK

###### <span id="get-series-series-id-200-schema"></span> Schema
   
  

[][VersionedSeries](#versioned-series)

##### <span id="get-series-series-id-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="get-series-series-id-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-series-id-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="get-series-series-id-401-schema"></span> Schema

##### <span id="get-series-series-id-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="get-series-series-id-403-schema"></span> Schema

##### <span id="get-series-series-id-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="get-series-series-id-404-schema"></span> Schema

##### <span id="get-series-series-id-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="get-series-series-id-405-schema"></span> Schema

##### <span id="get-series-series-id-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="get-series-series-id-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-series-id-default"></span> Default Response
Other error. See error details.

###### <span id="get-series-series-id-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="get-series-series-id-latest"></span> Get the latest version of the description of a time series (*GetSeriesSeriesIDLatest*)

```
GET /v1/series/{seriesId}/latest
```

The latest version is determined according to semantic versioning rules (e.g. v1.2.3 < v1.2.4).

Unless requested by the query parameters, the time series values are not returned by default.


#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| seriesId | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | The unique ID of a time series. |
| audit | `query` | boolean | `bool` |  |  |  | When audit is specified, the response will also contain the audit trail field. |
| brief | `query` | boolean | `bool` |  |  |  | When brief is specified, the response will only contain essential data and strip long descriptions. |
| deep | `query` | boolean | `bool` |  |  |  | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description. |
| status | `query` | string | `string` |  |  |  | Filter the result according to a given status.

Only series in the "PUBLISHED" status are returned to the public.

Series owner may consult their series in any status. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-series-series-id-latest-200) | OK | Latest available version of a time series. |  | [schema](#get-series-series-id-latest-200-schema) |
| [400](#get-series-series-id-latest-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#get-series-series-id-latest-400-schema) |
| [401](#get-series-series-id-latest-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#get-series-series-id-latest-401-schema) |
| [403](#get-series-series-id-latest-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#get-series-series-id-latest-403-schema) |
| [404](#get-series-series-id-latest-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#get-series-series-id-latest-404-schema) |
| [405](#get-series-series-id-latest-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#get-series-series-id-latest-405-schema) |
| [500](#get-series-series-id-latest-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#get-series-series-id-latest-500-schema) |
| [default](#get-series-series-id-latest-default) | | Other error. See error details. |  | [schema](#get-series-series-id-latest-default-schema) |

#### Responses


##### <span id="get-series-series-id-latest-200"></span> 200 - Latest available version of a time series.
Status: OK

###### <span id="get-series-series-id-latest-200-schema"></span> Schema
   
  

[VersionedSeries](#versioned-series)

##### <span id="get-series-series-id-latest-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="get-series-series-id-latest-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-series-id-latest-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="get-series-series-id-latest-401-schema"></span> Schema

##### <span id="get-series-series-id-latest-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="get-series-series-id-latest-403-schema"></span> Schema

##### <span id="get-series-series-id-latest-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="get-series-series-id-latest-404-schema"></span> Schema

##### <span id="get-series-series-id-latest-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="get-series-series-id-latest-405-schema"></span> Schema

##### <span id="get-series-series-id-latest-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="get-series-series-id-latest-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-series-id-latest-default"></span> Default Response
Other error. See error details.

###### <span id="get-series-series-id-latest-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="get-series-series-id-latest-values"></span> Get the values from the latest version of a time series (*GetSeriesSeriesIDLatestValues*)

```
GET /v1/series/{seriesId}/latest/values
```

The latest version is determined according to semantic versioning rules (e.g. v1.2.3 < v1.2.4).

Values are returned in chronological order.

If the request negotiates a response MIME type with text/csv (with the Accept header),
this endpoint produces a CSV file.


> [Read more](https://github.io/wiki/TODO#seriescsv "API responses: time series values
")

#### Produces
  * application/json
  * text/csv

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| seriesId | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | The unique ID of a time series. |
| convert | `query` | string | `string` |  |  |  | Converts values in some other compatible measurement unit, possibly specified as a composite measurement unit. |
| from | `query` | date-time (formatted string) | `strfmt.DateTime` |  |  |  | Truncates values starting from this date (boundary is included). |
| to | `query` | date-time (formatted string) | `strfmt.DateTime` |  |  |  | Truncates values up to this date (boundary is included). |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-series-series-id-latest-values-200) | OK | Time series values. |  | [schema](#get-series-series-id-latest-values-200-schema) |
| [400](#get-series-series-id-latest-values-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#get-series-series-id-latest-values-400-schema) |
| [401](#get-series-series-id-latest-values-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#get-series-series-id-latest-values-401-schema) |
| [403](#get-series-series-id-latest-values-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#get-series-series-id-latest-values-403-schema) |
| [404](#get-series-series-id-latest-values-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#get-series-series-id-latest-values-404-schema) |
| [405](#get-series-series-id-latest-values-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#get-series-series-id-latest-values-405-schema) |
| [500](#get-series-series-id-latest-values-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#get-series-series-id-latest-values-500-schema) |
| [default](#get-series-series-id-latest-values-default) | | Other error. See error details. |  | [schema](#get-series-series-id-latest-values-default-schema) |

#### Responses


##### <span id="get-series-series-id-latest-values-200"></span> 200 - Time series values.
Status: OK

###### <span id="get-series-series-id-latest-values-200-schema"></span> Schema
   
  


 [TsValues](#ts-values)

##### <span id="get-series-series-id-latest-values-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="get-series-series-id-latest-values-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-series-id-latest-values-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="get-series-series-id-latest-values-401-schema"></span> Schema

##### <span id="get-series-series-id-latest-values-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="get-series-series-id-latest-values-403-schema"></span> Schema

##### <span id="get-series-series-id-latest-values-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="get-series-series-id-latest-values-404-schema"></span> Schema

##### <span id="get-series-series-id-latest-values-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="get-series-series-id-latest-values-405-schema"></span> Schema

##### <span id="get-series-series-id-latest-values-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="get-series-series-id-latest-values-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-series-id-latest-values-default"></span> Default Response
Other error. See error details.

###### <span id="get-series-series-id-latest-values-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="get-series-series-id-semver"></span> Get all semver tags associated to a series (*GetSeriesSeriesIDSemver*)

```
GET /v1/series/{seriesId}/semver
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| seriesId | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | The unique ID of a time series. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-series-series-id-semver-200) | OK | All semver tagged versions available for a time series. |  | [schema](#get-series-series-id-semver-200-schema) |
| [400](#get-series-series-id-semver-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#get-series-series-id-semver-400-schema) |
| [401](#get-series-series-id-semver-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#get-series-series-id-semver-401-schema) |
| [403](#get-series-series-id-semver-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#get-series-series-id-semver-403-schema) |
| [404](#get-series-series-id-semver-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#get-series-series-id-semver-404-schema) |
| [405](#get-series-series-id-semver-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#get-series-series-id-semver-405-schema) |
| [500](#get-series-series-id-semver-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#get-series-series-id-semver-500-schema) |
| [default](#get-series-series-id-semver-default) | | Other error. See error details. |  | [schema](#get-series-series-id-semver-default-schema) |

#### Responses


##### <span id="get-series-series-id-semver-200"></span> 200 - All semver tagged versions available for a time series.
Status: OK

###### <span id="get-series-series-id-semver-200-schema"></span> Schema
   
  

[]string

##### <span id="get-series-series-id-semver-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="get-series-series-id-semver-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-series-id-semver-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="get-series-series-id-semver-401-schema"></span> Schema

##### <span id="get-series-series-id-semver-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="get-series-series-id-semver-403-schema"></span> Schema

##### <span id="get-series-series-id-semver-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="get-series-series-id-semver-404-schema"></span> Schema

##### <span id="get-series-series-id-semver-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="get-series-series-id-semver-405-schema"></span> Schema

##### <span id="get-series-series-id-semver-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="get-series-series-id-semver-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-series-id-semver-default"></span> Default Response
Other error. See error details.

###### <span id="get-series-series-id-semver-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="get-series-series-id-semver-semver"></span> Get a version of a time series with a semver tag (*GetSeriesSeriesIDSemverSemver*)

```
GET /v1/series/{seriesId}/semver/{semver}
```

Semver tags follow semantic versioning rules (e.g. v1.2.3, v1.2.3-rc1, v1.3.0 ...).

Unless requested by the query parameters, the time series values are not returned by default.


#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| semver | `path` | uuid (formatted string) | `strfmt.UUID` |  | ✓ |  | The semver tag of a series' version. |
| seriesId | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | The unique ID of a time series. |
| audit | `query` | boolean | `bool` |  |  |  | When audit is specified, the response will also contain the audit trail field. |
| brief | `query` | boolean | `bool` |  |  |  | When brief is specified, the response will only contain essential data and strip long descriptions. |
| deep | `query` | boolean | `bool` |  |  |  | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-series-series-id-semver-semver-200) | OK | Versioned time series for a given semver tag. |  | [schema](#get-series-series-id-semver-semver-200-schema) |
| [400](#get-series-series-id-semver-semver-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#get-series-series-id-semver-semver-400-schema) |
| [401](#get-series-series-id-semver-semver-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#get-series-series-id-semver-semver-401-schema) |
| [403](#get-series-series-id-semver-semver-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#get-series-series-id-semver-semver-403-schema) |
| [404](#get-series-series-id-semver-semver-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#get-series-series-id-semver-semver-404-schema) |
| [405](#get-series-series-id-semver-semver-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#get-series-series-id-semver-semver-405-schema) |
| [500](#get-series-series-id-semver-semver-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#get-series-series-id-semver-semver-500-schema) |
| [default](#get-series-series-id-semver-semver-default) | | Other error. See error details. |  | [schema](#get-series-series-id-semver-semver-default-schema) |

#### Responses


##### <span id="get-series-series-id-semver-semver-200"></span> 200 - Versioned time series for a given semver tag.
Status: OK

###### <span id="get-series-series-id-semver-semver-200-schema"></span> Schema
   
  

[VersionedSeries](#versioned-series)

##### <span id="get-series-series-id-semver-semver-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="get-series-series-id-semver-semver-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-series-id-semver-semver-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="get-series-series-id-semver-semver-401-schema"></span> Schema

##### <span id="get-series-series-id-semver-semver-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="get-series-series-id-semver-semver-403-schema"></span> Schema

##### <span id="get-series-series-id-semver-semver-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="get-series-series-id-semver-semver-404-schema"></span> Schema

##### <span id="get-series-series-id-semver-semver-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="get-series-series-id-semver-semver-405-schema"></span> Schema

##### <span id="get-series-series-id-semver-semver-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="get-series-series-id-semver-semver-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-series-id-semver-semver-default"></span> Default Response
Other error. See error details.

###### <span id="get-series-series-id-semver-semver-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="get-series-series-id-semver-semver-values"></span> Get the values of version of a time series with a semver tag (*GetSeriesSeriesIDSemverSemverValues*)

```
GET /v1/series/{seriesId}/semver/{semver}/values
```


If the request negotiates a response MIME type with text/csv (with the Accept header),
this endpoint produces a CSV file.


> [Read more](https://github.io/wiki/TODO#seriescsv "API responses: time series values
")

#### Produces
  * application/json
  * text/csv

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| semver | `path` | uuid (formatted string) | `strfmt.UUID` |  | ✓ |  | The semver tag of a series' version. |
| seriesId | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | The unique ID of a time series. |
| convert | `query` | string | `string` |  |  |  | Converts values in some other compatible measurement unit, possibly specified as a composite measurement unit. |
| from | `query` | date-time (formatted string) | `strfmt.DateTime` |  |  |  | Truncates values starting from this date (boundary is included). |
| to | `query` | date-time (formatted string) | `strfmt.DateTime` |  |  |  | Truncates values up to this date (boundary is included). |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-series-series-id-semver-semver-values-200) | OK | Values of a given version of a time series. |  | [schema](#get-series-series-id-semver-semver-values-200-schema) |
| [400](#get-series-series-id-semver-semver-values-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#get-series-series-id-semver-semver-values-400-schema) |
| [401](#get-series-series-id-semver-semver-values-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#get-series-series-id-semver-semver-values-401-schema) |
| [403](#get-series-series-id-semver-semver-values-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#get-series-series-id-semver-semver-values-403-schema) |
| [404](#get-series-series-id-semver-semver-values-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#get-series-series-id-semver-semver-values-404-schema) |
| [405](#get-series-series-id-semver-semver-values-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#get-series-series-id-semver-semver-values-405-schema) |
| [500](#get-series-series-id-semver-semver-values-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#get-series-series-id-semver-semver-values-500-schema) |
| [default](#get-series-series-id-semver-semver-values-default) | | Other error. See error details. |  | [schema](#get-series-series-id-semver-semver-values-default-schema) |

#### Responses


##### <span id="get-series-series-id-semver-semver-values-200"></span> 200 - Values of a given version of a time series.
Status: OK

###### <span id="get-series-series-id-semver-semver-values-200-schema"></span> Schema
   
  


 [TsValues](#ts-values)

##### <span id="get-series-series-id-semver-semver-values-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="get-series-series-id-semver-semver-values-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-series-id-semver-semver-values-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="get-series-series-id-semver-semver-values-401-schema"></span> Schema

##### <span id="get-series-series-id-semver-semver-values-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="get-series-series-id-semver-semver-values-403-schema"></span> Schema

##### <span id="get-series-series-id-semver-semver-values-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="get-series-series-id-semver-semver-values-404-schema"></span> Schema

##### <span id="get-series-series-id-semver-semver-values-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="get-series-series-id-semver-semver-values-405-schema"></span> Schema

##### <span id="get-series-series-id-semver-semver-values-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="get-series-series-id-semver-semver-values-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-series-id-semver-semver-values-default"></span> Default Response
Other error. See error details.

###### <span id="get-series-series-id-semver-semver-values-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="get-series-versions-versioned-series-id"></span> Get a version of a time series (*GetSeriesVersionsVersionedSeriesID*)

```
GET /v1/series/versions/{versionedSeriesId}
```

Unless requested by the query parameters, the time series values are not returned by default.


#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| versionedSeriesId | `path` | uuid (formatted string) | `strfmt.UUID` |  | ✓ |  | The UUID of a version of the time series. |
| audit | `query` | boolean | `bool` |  |  |  | When audit is specified, the response will also contain the audit trail field. |
| brief | `query` | boolean | `bool` |  |  |  | When brief is specified, the response will only contain essential data and strip long descriptions. |
| deep | `query` | boolean | `bool` |  |  |  | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-series-versions-versioned-series-id-200) | OK | Description of a given version of a time series. |  | [schema](#get-series-versions-versioned-series-id-200-schema) |
| [400](#get-series-versions-versioned-series-id-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#get-series-versions-versioned-series-id-400-schema) |
| [401](#get-series-versions-versioned-series-id-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#get-series-versions-versioned-series-id-401-schema) |
| [403](#get-series-versions-versioned-series-id-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#get-series-versions-versioned-series-id-403-schema) |
| [404](#get-series-versions-versioned-series-id-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#get-series-versions-versioned-series-id-404-schema) |
| [405](#get-series-versions-versioned-series-id-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#get-series-versions-versioned-series-id-405-schema) |
| [500](#get-series-versions-versioned-series-id-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#get-series-versions-versioned-series-id-500-schema) |
| [default](#get-series-versions-versioned-series-id-default) | | Other error. See error details. |  | [schema](#get-series-versions-versioned-series-id-default-schema) |

#### Responses


##### <span id="get-series-versions-versioned-series-id-200"></span> 200 - Description of a given version of a time series.
Status: OK

###### <span id="get-series-versions-versioned-series-id-200-schema"></span> Schema
   
  

[VersionedSeries](#versioned-series)

##### <span id="get-series-versions-versioned-series-id-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="get-series-versions-versioned-series-id-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-versions-versioned-series-id-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="get-series-versions-versioned-series-id-401-schema"></span> Schema

##### <span id="get-series-versions-versioned-series-id-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="get-series-versions-versioned-series-id-403-schema"></span> Schema

##### <span id="get-series-versions-versioned-series-id-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="get-series-versions-versioned-series-id-404-schema"></span> Schema

##### <span id="get-series-versions-versioned-series-id-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="get-series-versions-versioned-series-id-405-schema"></span> Schema

##### <span id="get-series-versions-versioned-series-id-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="get-series-versions-versioned-series-id-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-versions-versioned-series-id-default"></span> Default Response
Other error. See error details.

###### <span id="get-series-versions-versioned-series-id-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="get-series-versions-versioned-series-id-values"></span> Get the values of version of a time series (*GetSeriesVersionsVersionedSeriesIDValues*)

```
GET /v1/series/versions/{versionedSeriesId}/values
```

If the request negotiates a response MIME type with text/csv (with the Accept header),
this endpoint produces a CSV file.


> [Read more](https://github.io/wiki/TODO#seriescsv "API responses: time series values.
")

#### Produces
  * application/json
  * text/csv

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| versionedSeriesId | `path` | uuid (formatted string) | `strfmt.UUID` |  | ✓ |  | The UUID of a version of the time series. |
| convert | `query` | string | `string` |  |  |  | Converts values in some other compatible measurement unit, possibly specified as a composite measurement unit. |
| from | `query` | date-time (formatted string) | `strfmt.DateTime` |  |  |  | Truncates values starting from this date (boundary is included). |
| to | `query` | date-time (formatted string) | `strfmt.DateTime` |  |  |  | Truncates values up to this date (boundary is included). |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-series-versions-versioned-series-id-values-200) | OK | Values of a given version of a time series. |  | [schema](#get-series-versions-versioned-series-id-values-200-schema) |
| [400](#get-series-versions-versioned-series-id-values-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#get-series-versions-versioned-series-id-values-400-schema) |
| [401](#get-series-versions-versioned-series-id-values-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#get-series-versions-versioned-series-id-values-401-schema) |
| [403](#get-series-versions-versioned-series-id-values-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#get-series-versions-versioned-series-id-values-403-schema) |
| [404](#get-series-versions-versioned-series-id-values-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#get-series-versions-versioned-series-id-values-404-schema) |
| [405](#get-series-versions-versioned-series-id-values-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#get-series-versions-versioned-series-id-values-405-schema) |
| [500](#get-series-versions-versioned-series-id-values-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#get-series-versions-versioned-series-id-values-500-schema) |
| [default](#get-series-versions-versioned-series-id-values-default) | | Other error. See error details. |  | [schema](#get-series-versions-versioned-series-id-values-default-schema) |

#### Responses


##### <span id="get-series-versions-versioned-series-id-values-200"></span> 200 - Values of a given version of a time series.
Status: OK

###### <span id="get-series-versions-versioned-series-id-values-200-schema"></span> Schema
   
  


 [TsValues](#ts-values)

##### <span id="get-series-versions-versioned-series-id-values-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="get-series-versions-versioned-series-id-values-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-versions-versioned-series-id-values-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="get-series-versions-versioned-series-id-values-401-schema"></span> Schema

##### <span id="get-series-versions-versioned-series-id-values-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="get-series-versions-versioned-series-id-values-403-schema"></span> Schema

##### <span id="get-series-versions-versioned-series-id-values-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="get-series-versions-versioned-series-id-values-404-schema"></span> Schema

##### <span id="get-series-versions-versioned-series-id-values-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="get-series-versions-versioned-series-id-values-405-schema"></span> Schema

##### <span id="get-series-versions-versioned-series-id-values-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="get-series-versions-versioned-series-id-values-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="get-series-versions-versioned-series-id-values-default"></span> Default Response
Other error. See error details.

###### <span id="get-series-versions-versioned-series-id-values-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="post-series"></span> Creates a new time series (*PostSeries*)

```
POST /v1/series
```

This creates a new time time series.

If some values are specified, the initial version for the new series is created with these values.


#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| series | `body` | [Series](#series) | `models.Series` | | ✓ | | The description of a time series. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#post-series-201) | Created | Series successfully created.

Check the response headers to retrieve this resource. | ✓ | [schema](#post-series-201-schema) |
| [400](#post-series-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#post-series-400-schema) |
| [401](#post-series-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#post-series-401-schema) |
| [403](#post-series-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#post-series-403-schema) |
| [404](#post-series-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#post-series-404-schema) |
| [405](#post-series-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#post-series-405-schema) |
| [409](#post-series-409) | Conflict | Resource already exists. An object creation was requested, but this object was already existing. |  | [schema](#post-series-409-schema) |
| [500](#post-series-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#post-series-500-schema) |
| [default](#post-series-default) | | Other error. See error details. |  | [schema](#post-series-default-schema) |

#### Responses


##### <span id="post-series-201"></span> 201 - Series successfully created.

Check the response headers to retrieve this resource.
Status: Created

###### <span id="post-series-201-schema"></span> Schema

###### Response headers

| Name | Type | Go type | Separator | Default | Description |
|------|------|---------|-----------|---------|-------------|
| Location | uri (formatted string) | `strfmt.URI` |  |  | The URI of the newly created resource. |
| X-ID | int64 (formatted integer) | `int64` |  |  | The ID of the newly created resource. |
| X-LocationVersion | uri (formatted string) | `strfmt.URI` |  |  | The URI of the newly created versioned resource, whenever applicable. |
| X-VersionID | uuid (formatted string) | `strfmt.UUID` |  |  | The UUID of the newly created versioned resource, whenever applicable. |

##### <span id="post-series-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="post-series-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="post-series-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="post-series-401-schema"></span> Schema

##### <span id="post-series-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="post-series-403-schema"></span> Schema

##### <span id="post-series-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="post-series-404-schema"></span> Schema

##### <span id="post-series-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="post-series-405-schema"></span> Schema

##### <span id="post-series-409"></span> 409 - Resource already exists. An object creation was requested, but this object was already existing.
Status: Conflict

###### <span id="post-series-409-schema"></span> Schema

##### <span id="post-series-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="post-series-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="post-series-default"></span> Default Response
Other error. See error details.

###### <span id="post-series-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="post-series-series-id"></span> Creates a new version of a time series (*PostSeriesSeriesID*)

```
POST /v1/series/{seriesId}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| seriesId | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | The unique ID of a time series. |
| series | `body` | [Series](#series) | `models.Series` | | ✓ | | The description of a time series. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#post-series-series-id-201) | Created | Series version successfully created.

Check the response headers to retrieve this resource. | ✓ | [schema](#post-series-series-id-201-schema) |
| [400](#post-series-series-id-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#post-series-series-id-400-schema) |
| [401](#post-series-series-id-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#post-series-series-id-401-schema) |
| [403](#post-series-series-id-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#post-series-series-id-403-schema) |
| [404](#post-series-series-id-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#post-series-series-id-404-schema) |
| [405](#post-series-series-id-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#post-series-series-id-405-schema) |
| [409](#post-series-series-id-409) | Conflict | Resource already exists. An object creation was requested, but this object was already existing. |  | [schema](#post-series-series-id-409-schema) |
| [500](#post-series-series-id-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#post-series-series-id-500-schema) |
| [default](#post-series-series-id-default) | | Other error. See error details. |  | [schema](#post-series-series-id-default-schema) |

#### Responses


##### <span id="post-series-series-id-201"></span> 201 - Series version successfully created.

Check the response headers to retrieve this resource.
Status: Created

###### <span id="post-series-series-id-201-schema"></span> Schema

###### Response headers

| Name | Type | Go type | Separator | Default | Description |
|------|------|---------|-----------|---------|-------------|
| Location | uri (formatted string) | `strfmt.URI` |  |  | The URI of the newly created versioned resource. |
| X-LocationVersion | uri (formatted string) | `strfmt.URI` |  |  | The URI of the newly created versioned resource. |
| X-VersionID | uuid (formatted string) | `strfmt.UUID` |  |  | The UUID of the newly created versioned resource. |

##### <span id="post-series-series-id-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="post-series-series-id-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="post-series-series-id-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="post-series-series-id-401-schema"></span> Schema

##### <span id="post-series-series-id-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="post-series-series-id-403-schema"></span> Schema

##### <span id="post-series-series-id-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="post-series-series-id-404-schema"></span> Schema

##### <span id="post-series-series-id-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="post-series-series-id-405-schema"></span> Schema

##### <span id="post-series-series-id-409"></span> 409 - Resource already exists. An object creation was requested, but this object was already existing.
Status: Conflict

###### <span id="post-series-series-id-409-schema"></span> Schema

##### <span id="post-series-series-id-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="post-series-series-id-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="post-series-series-id-default"></span> Default Response
Other error. See error details.

###### <span id="post-series-series-id-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="post-series-versions-versioned-series-id"></span> creates a new version of a time series (*PostSeriesVersionsVersionedSeriesID*)

```
POST /v1/series/versions/{versionedSeriesId}
```

A conflict is reported if the version was already attributed to a version of the same series. Use PUT to update existing values.


> [Read more](https://github.io/wiki/TODO#seriesversion "API inputs: time series version.
")

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| versionedSeriesId | `path` | uuid (formatted string) | `strfmt.UUID` |  | ✓ |  | The UUID of a version of the time series. |
| versionSeries | `body` | [VersionedSeries](#versioned-series) | `models.VersionedSeries` | | ✓ | | The description of a version of a time series. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#post-series-versions-versioned-series-id-201) | Created | Time series version successfully created.

Check the response headers to retrieve this resource. | ✓ | [schema](#post-series-versions-versioned-series-id-201-schema) |
| [400](#post-series-versions-versioned-series-id-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#post-series-versions-versioned-series-id-400-schema) |
| [401](#post-series-versions-versioned-series-id-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#post-series-versions-versioned-series-id-401-schema) |
| [403](#post-series-versions-versioned-series-id-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#post-series-versions-versioned-series-id-403-schema) |
| [404](#post-series-versions-versioned-series-id-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#post-series-versions-versioned-series-id-404-schema) |
| [405](#post-series-versions-versioned-series-id-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#post-series-versions-versioned-series-id-405-schema) |
| [409](#post-series-versions-versioned-series-id-409) | Conflict | Resource already exists. An object creation was requested, but this object was already existing. |  | [schema](#post-series-versions-versioned-series-id-409-schema) |
| [500](#post-series-versions-versioned-series-id-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#post-series-versions-versioned-series-id-500-schema) |
| [default](#post-series-versions-versioned-series-id-default) | | Other error. See error details. |  | [schema](#post-series-versions-versioned-series-id-default-schema) |

#### Responses


##### <span id="post-series-versions-versioned-series-id-201"></span> 201 - Time series version successfully created.

Check the response headers to retrieve this resource.
Status: Created

###### <span id="post-series-versions-versioned-series-id-201-schema"></span> Schema

###### Response headers

| Name | Type | Go type | Separator | Default | Description |
|------|------|---------|-----------|---------|-------------|
| Location | uri (formatted string) | `strfmt.URI` |  |  | The URI of the newly created resource. |
| X-ID | int64 (formatted integer) | `int64` |  |  | The ID of the updated resource. |

##### <span id="post-series-versions-versioned-series-id-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="post-series-versions-versioned-series-id-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="post-series-versions-versioned-series-id-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="post-series-versions-versioned-series-id-401-schema"></span> Schema

##### <span id="post-series-versions-versioned-series-id-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="post-series-versions-versioned-series-id-403-schema"></span> Schema

##### <span id="post-series-versions-versioned-series-id-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="post-series-versions-versioned-series-id-404-schema"></span> Schema

##### <span id="post-series-versions-versioned-series-id-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="post-series-versions-versioned-series-id-405-schema"></span> Schema

##### <span id="post-series-versions-versioned-series-id-409"></span> 409 - Resource already exists. An object creation was requested, but this object was already existing.
Status: Conflict

###### <span id="post-series-versions-versioned-series-id-409-schema"></span> Schema

##### <span id="post-series-versions-versioned-series-id-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="post-series-versions-versioned-series-id-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="post-series-versions-versioned-series-id-default"></span> Default Response
Other error. See error details.

###### <span id="post-series-versions-versioned-series-id-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="post-series-versions-versioned-series-id-values"></span> Uploads values to the version of a time series (*PostSeriesVersionsVersionedSeriesIDValues*)

```
POST /v1/series/versions/{versionedSeriesId}/values
```

This endpoint creates time series values for the latest version.

Only the owners of the time series may add values.

A conflict is reported if values were already attributed to this version. Use PUT to update existing values.

If the request negotiates a request MIME type with text/csv (with the Content-Type header),
this endpoint consumes a CSV file.


> [Read more](https://github.io/wiki/TODO#seriescsv "API inputs: time series values.
")

#### Consumes
  * application/json
  * text/csv

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| versionedSeriesId | `path` | uuid (formatted string) | `strfmt.UUID` |  | ✓ |  | The UUID of a version of the time series. |
| values | `body` | [TsValues](#ts-values) | `models.TsValues` | | ✓ | | The values of a time series |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#post-series-versions-versioned-series-id-values-201) | Created | Time series values successfully created.

Check the response headers to retrieve this resource. | ✓ | [schema](#post-series-versions-versioned-series-id-values-201-schema) |
| [400](#post-series-versions-versioned-series-id-values-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#post-series-versions-versioned-series-id-values-400-schema) |
| [401](#post-series-versions-versioned-series-id-values-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#post-series-versions-versioned-series-id-values-401-schema) |
| [403](#post-series-versions-versioned-series-id-values-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#post-series-versions-versioned-series-id-values-403-schema) |
| [404](#post-series-versions-versioned-series-id-values-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#post-series-versions-versioned-series-id-values-404-schema) |
| [405](#post-series-versions-versioned-series-id-values-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#post-series-versions-versioned-series-id-values-405-schema) |
| [409](#post-series-versions-versioned-series-id-values-409) | Conflict | Resource already exists. An object creation was requested, but this object was already existing. |  | [schema](#post-series-versions-versioned-series-id-values-409-schema) |
| [500](#post-series-versions-versioned-series-id-values-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#post-series-versions-versioned-series-id-values-500-schema) |
| [default](#post-series-versions-versioned-series-id-values-default) | | Other error. See error details. |  | [schema](#post-series-versions-versioned-series-id-values-default-schema) |

#### Responses


##### <span id="post-series-versions-versioned-series-id-values-201"></span> 201 - Time series values successfully created.

Check the response headers to retrieve this resource.
Status: Created

###### <span id="post-series-versions-versioned-series-id-values-201-schema"></span> Schema

###### Response headers

| Name | Type | Go type | Separator | Default | Description |
|------|------|---------|-----------|---------|-------------|
| Location | uri (formatted string) | `strfmt.URI` |  |  | The URI of the newly created resource. |
| X-ID | int64 (formatted integer) | `int64` |  |  | The ID of the updated resource. |

##### <span id="post-series-versions-versioned-series-id-values-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="post-series-versions-versioned-series-id-values-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="post-series-versions-versioned-series-id-values-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="post-series-versions-versioned-series-id-values-401-schema"></span> Schema

##### <span id="post-series-versions-versioned-series-id-values-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="post-series-versions-versioned-series-id-values-403-schema"></span> Schema

##### <span id="post-series-versions-versioned-series-id-values-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="post-series-versions-versioned-series-id-values-404-schema"></span> Schema

##### <span id="post-series-versions-versioned-series-id-values-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="post-series-versions-versioned-series-id-values-405-schema"></span> Schema

##### <span id="post-series-versions-versioned-series-id-values-409"></span> 409 - Resource already exists. An object creation was requested, but this object was already existing.
Status: Conflict

###### <span id="post-series-versions-versioned-series-id-values-409-schema"></span> Schema

##### <span id="post-series-versions-versioned-series-id-values-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="post-series-versions-versioned-series-id-values-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="post-series-versions-versioned-series-id-values-default"></span> Default Response
Other error. See error details.

###### <span id="post-series-versions-versioned-series-id-values-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="put-series-series-id"></span> Updates metadata about a time series (*PutSeriesSeriesID*)

```
PUT /v1/series/{seriesId}
```

This action only updates metadata about the series (such as associated themes, tags, etc) and does not create a new version.

Only owners registered for a series may update the series.

Any time series values specified in the input are ignored. One must update a specific version to modify the values of a time series.


#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| seriesId | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | The unique ID of a time series. |
| series | `body` | [Series](#series) | `models.Series` | | ✓ | | The description of a time series. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#put-series-series-id-204) | No Content | Series successfully updated.

Check the response headers to retrieve this resource. | ✓ | [schema](#put-series-series-id-204-schema) |
| [400](#put-series-series-id-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#put-series-series-id-400-schema) |
| [401](#put-series-series-id-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#put-series-series-id-401-schema) |
| [403](#put-series-series-id-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#put-series-series-id-403-schema) |
| [404](#put-series-series-id-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#put-series-series-id-404-schema) |
| [405](#put-series-series-id-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#put-series-series-id-405-schema) |
| [500](#put-series-series-id-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#put-series-series-id-500-schema) |
| [default](#put-series-series-id-default) | | Other error. See error details. |  | [schema](#put-series-series-id-default-schema) |

#### Responses


##### <span id="put-series-series-id-204"></span> 204 - Series successfully updated.

Check the response headers to retrieve this resource.
Status: No Content

###### <span id="put-series-series-id-204-schema"></span> Schema

###### Response headers

| Name | Type | Go type | Separator | Default | Description |
|------|------|---------|-----------|---------|-------------|
| Location | uri (formatted string) | `strfmt.URI` |  |  | The URI of the updated resource. |
| X-ID | int64 (formatted integer) | `int64` |  |  | The ID of the updated resource. |

##### <span id="put-series-series-id-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="put-series-series-id-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="put-series-series-id-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="put-series-series-id-401-schema"></span> Schema

##### <span id="put-series-series-id-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="put-series-series-id-403-schema"></span> Schema

##### <span id="put-series-series-id-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="put-series-series-id-404-schema"></span> Schema

##### <span id="put-series-series-id-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="put-series-series-id-405-schema"></span> Schema

##### <span id="put-series-series-id-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="put-series-series-id-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="put-series-series-id-default"></span> Default Response
Other error. See error details.

###### <span id="put-series-series-id-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="put-series-versions-versioned-series-id"></span> Updates the metadata of version of a time series (*PutSeriesVersionsVersionedSeriesID*)

```
PUT /v1/series/versions/{versionedSeriesId}
```

This endpoint replaces the metadata for the requested version, without creating a new version.

Only the owners of the time series may modify values.


> [Read more](https://github.io/wiki/TODO#seriesversion "API inputs: time series version.
")

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| versionedSeriesId | `path` | uuid (formatted string) | `strfmt.UUID` |  | ✓ |  | The UUID of a version of the time series. |
| versionSeries | `body` | [VersionedSeries](#versioned-series) | `models.VersionedSeries` | | ✓ | | The description of a version of a time series. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#put-series-versions-versioned-series-id-204) | No Content | Time series values successfully updated.

Check the response headers to retrieve this resource. | ✓ | [schema](#put-series-versions-versioned-series-id-204-schema) |
| [400](#put-series-versions-versioned-series-id-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#put-series-versions-versioned-series-id-400-schema) |
| [401](#put-series-versions-versioned-series-id-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#put-series-versions-versioned-series-id-401-schema) |
| [403](#put-series-versions-versioned-series-id-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#put-series-versions-versioned-series-id-403-schema) |
| [404](#put-series-versions-versioned-series-id-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#put-series-versions-versioned-series-id-404-schema) |
| [405](#put-series-versions-versioned-series-id-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#put-series-versions-versioned-series-id-405-schema) |
| [500](#put-series-versions-versioned-series-id-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#put-series-versions-versioned-series-id-500-schema) |
| [default](#put-series-versions-versioned-series-id-default) | | Other error. See error details. |  | [schema](#put-series-versions-versioned-series-id-default-schema) |

#### Responses


##### <span id="put-series-versions-versioned-series-id-204"></span> 204 - Time series values successfully updated.

Check the response headers to retrieve this resource.
Status: No Content

###### <span id="put-series-versions-versioned-series-id-204-schema"></span> Schema

###### Response headers

| Name | Type | Go type | Separator | Default | Description |
|------|------|---------|-----------|---------|-------------|
| Location | uri (formatted string) | `strfmt.URI` |  |  | The URI of the updated resource. |
| X-LocationVersion | uri (formatted string) | `strfmt.URI` |  |  | The URI of the updated resource. |
| X-VersionID | int64 (formatted integer) | `int64` |  |  | The ID of the updated resource. |

##### <span id="put-series-versions-versioned-series-id-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="put-series-versions-versioned-series-id-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="put-series-versions-versioned-series-id-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="put-series-versions-versioned-series-id-401-schema"></span> Schema

##### <span id="put-series-versions-versioned-series-id-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="put-series-versions-versioned-series-id-403-schema"></span> Schema

##### <span id="put-series-versions-versioned-series-id-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="put-series-versions-versioned-series-id-404-schema"></span> Schema

##### <span id="put-series-versions-versioned-series-id-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="put-series-versions-versioned-series-id-405-schema"></span> Schema

##### <span id="put-series-versions-versioned-series-id-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="put-series-versions-versioned-series-id-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="put-series-versions-versioned-series-id-default"></span> Default Response
Other error. See error details.

###### <span id="put-series-versions-versioned-series-id-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="put-series-versions-versioned-series-id-values"></span> Replaces the values of version of a time series (*PutSeriesVersionsVersionedSeriesIDValues*)

```
PUT /v1/series/versions/{versionedSeriesId}/values
```

This endpoint replaces time series values for the requested version, without creating a new version.

Only the owners of the time series may modify values.

A conflict is reported if values were already attributed to this version. Use PUT to update existing values.

If the request negotiates a request MIME type with text/csv (with the Content-Type header),
this endpoint consumes a CSV file.


> [Read more](https://github.io/wiki/TODO#seriescsv "API responses: time series values.
")

#### Consumes
  * application/json
  * text/csv

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| versionedSeriesId | `path` | uuid (formatted string) | `strfmt.UUID` |  | ✓ |  | The UUID of a version of the time series. |
| values | `body` | [TsValues](#ts-values) | `models.TsValues` | | ✓ | | The values of a time series |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#put-series-versions-versioned-series-id-values-204) | No Content | Time series values successfully updated.

Check the response headers to retrieve this resource. | ✓ | [schema](#put-series-versions-versioned-series-id-values-204-schema) |
| [400](#put-series-versions-versioned-series-id-values-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#put-series-versions-versioned-series-id-values-400-schema) |
| [401](#put-series-versions-versioned-series-id-values-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#put-series-versions-versioned-series-id-values-401-schema) |
| [403](#put-series-versions-versioned-series-id-values-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#put-series-versions-versioned-series-id-values-403-schema) |
| [404](#put-series-versions-versioned-series-id-values-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#put-series-versions-versioned-series-id-values-404-schema) |
| [405](#put-series-versions-versioned-series-id-values-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#put-series-versions-versioned-series-id-values-405-schema) |
| [500](#put-series-versions-versioned-series-id-values-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#put-series-versions-versioned-series-id-values-500-schema) |
| [default](#put-series-versions-versioned-series-id-values-default) | | Other error. See error details. |  | [schema](#put-series-versions-versioned-series-id-values-default-schema) |

#### Responses


##### <span id="put-series-versions-versioned-series-id-values-204"></span> 204 - Time series values successfully updated.

Check the response headers to retrieve this resource.
Status: No Content

###### <span id="put-series-versions-versioned-series-id-values-204-schema"></span> Schema

###### Response headers

| Name | Type | Go type | Separator | Default | Description |
|------|------|---------|-----------|---------|-------------|
| Location | uri (formatted string) | `strfmt.URI` |  |  | The URI of the updated resource. |
| X-LocationVersion | uri (formatted string) | `strfmt.URI` |  |  | The URI of the updated resource. |
| X-VersionID | int64 (formatted integer) | `int64` |  |  | The ID of the updated resource. |

##### <span id="put-series-versions-versioned-series-id-values-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="put-series-versions-versioned-series-id-values-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="put-series-versions-versioned-series-id-values-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="put-series-versions-versioned-series-id-values-401-schema"></span> Schema

##### <span id="put-series-versions-versioned-series-id-values-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="put-series-versions-versioned-series-id-values-403-schema"></span> Schema

##### <span id="put-series-versions-versioned-series-id-values-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="put-series-versions-versioned-series-id-values-404-schema"></span> Schema

##### <span id="put-series-versions-versioned-series-id-values-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="put-series-versions-versioned-series-id-values-405-schema"></span> Schema

##### <span id="put-series-versions-versioned-series-id-values-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="put-series-versions-versioned-series-id-values-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="put-series-versions-versioned-series-id-values-default"></span> Default Response
Other error. See error details.

###### <span id="put-series-versions-versioned-series-id-values-default-schema"></span> Schema

  

[APIError](#api-error)

## Models

### <span id="api-error"></span> apiError


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| code | string| `string` |  | |  |  |
| message | string| `string` |  | |  |  |



### <span id="audit"></span> audit


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| createdAt | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| createdBy | string| `string` |  | |  |  |
| lastUpdatedAt | date-time (formatted string)| `strfmt.DateTime` |  | |  |  |
| lastUpdatedBy | string| `string` |  | |  |  |



### <span id="class"></span> class


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| auditTrail | [Audit](#audit)| `Audit` |  | |  |  |
| descriptionLong | [Translation](#translation)| `Translation` |  | |  |  |
| descriptionShort | [Translation](#translation)| `Translation` |  | |  |  |
| id | integer| `int64` | ✓ | |  |  |
| shortCode | [ClassName](#class-name)| `ClassName` | ✓ | |  |  |
| title | [Translation](#translation)| `Translation` | ✓ | |  |  |



### <span id="class-description"></span> classDescription


> Metadata about nomenture classes.

  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| class | [ClassName](#class-name)| `ClassName` | ✓ | |  |  |
| descriptionLong | [Translation](#translation)| `Translation` | ✓ | |  |  |
| descriptionShort | [Translation](#translation)| `Translation` | ✓ | |  |  |
| metadata | [ClassMetadata](#class-metadata)| `ClassMetadata` | ✓ | |  |  |
| tableName | string| `string` | ✓ | |  |  |
| title | [Translation](#translation)| `Translation` | ✓ | |  |  |



### <span id="class-metadata"></span> classMetadata


> Metadata used by UIs to render classes.

  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| extraFields | []string| `[]string` |  | |  |  |
| fromTemplate | boolean| `bool` |  | |  |  |
| hasOneClass | map of string| `map[string]string` |  | |  |  |
| hasZeroManyClass | map of string| `map[string]string` |  | |  |  |
| hasZeroOneClass | map of string| `map[string]string` |  | |  |  |
| tagSearch | boolean| `bool` |  | |  |  |



### <span id="class-name"></span> className


  

[interface{}](#interface)

### <span id="composite-unit"></span> compositeUnit


> A composite measurement unit is built from base units with a formula.
TODO doc.

  



[interface{}](#interface)

### <span id="conversion"></span> conversion


> Unit conversions.

  




* inlined member (*AO0*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| auditTrail | [Audit](#audit)| `Audit` |  | |  |  |
| fromUnit | [Munit](#munit)| `Munit` | ✓ | |  |  |
| toUnit | [Munit](#munit)| `Munit` | ✓ | |  |  |


* composed type [ConversionSpec](#conversion-spec)

### <span id="conversion-spec"></span> conversionSpec


> Unit conversion specification.

  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| factor | float (formatted number)| `float32` |  | |  |  |
| formula | [Formula](#formula)| `expressions.Formula` |  | |  |  |
| intercept | float (formatted number)| `float32` |  | |  |  |
| reverse_formula | [Formula](#formula)| `expressions.Formula` |  | |  |  |
| toUnitCode | string| `string` | ✓ | |  |  |



### <span id="document"></span> document


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| url | uri (formatted string)| `strfmt.URI` |  | |  |  |



### <span id="documents"></span> documents


  

[][Document](#document)

### <span id="extra"></span> extra


> reserved for future use
  



[interface{}](#interface)

### <span id="geometry"></span> geometry


> A geojson geometry.

  



[interface{}](#interface)

### <span id="mdomain"></span> mdomain


> Measurement domains: this is a domain in which a kind of measurement applies.

Examples:
  Measuring speed is useful in aeronautics and mechanics, but not so much in construction.

  Torque is used in mechanics but not so much in chemistry.

Domains also apply to constants.

  




* composed type [Class](#class)
* [Extra](#extra)

### <span id="measurement"></span> measurement


> A measurement refers to any physical or economic measurement, the unit notwithstanding.

Examples:
  * speed
  * temperature

We introduce "quantity" to take into account measurement that count things, as well as "currency"
to take into account measurements of economic facts.

Measurements usually refer to one or several physical (or economic) dimensions.

Example:
  speed is homogeneous to "L^T-1"

However, some measurements are without dimension, such as "QUANTITY" and "ANGLE".

Series of ratios (e.g. % increase of GDP) are not considered measurements. They shall use special units which do not refer to measurements.

  




* composed type [Class](#class)
* inlined member (*AO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| dimensions | [DimensionsFormula](#dimensions-formula)| `expressions.DimensionsFormula` |  | |  |  |
| measurementHasMeasurementDomains | [][Mdomain](#mdomain)| `[]Mdomain` |  | |  |  |



### <span id="multiplier"></span> multiplier


> Multipliers using to prefix measurement units (e.g. m,c,d,da,h,k,M,G...).

  




* composed type [Class](#class)
* [Extra](#extra)

### <span id="mumeta"></span> mumeta


> Measurement units metadata, for rendering (symbol glyphs, aliases...)

  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| extra | [Extra](#extra)| `Extra` |  | |  |  |



### <span id="munit"></span> munit


> Measurement units are used to quantify values provided by time series.

  




* composed type [Class](#class)
* inlined member (*AO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| includedMultiplier | [Multiplier](#multiplier)| `Multiplier` |  | |  |  |
| isStandard | boolean| `bool` | ✓ | |  |  |
| measurementUnitHasConversions | [][ConversionSpec](#conversion-spec)| `[]ConversionSpec` |  | |  |  |
| metadata | [Mumeta](#mumeta)| `Mumeta` |  | |  |  |
| unitMeasurement | [Measurement](#measurement)| `Measurement` |  | |  |  |
| unitMeasurementSystem | [Musystem](#musystem)| `Musystem` |  | |  |  |



### <span id="musystem"></span> musystem


> Measurement unit systems.

  




* composed type [Class](#class)
* [Extra](#extra)

### <span id="ostatus"></span> ostatus


> Status of owners.

  




* composed type [Class](#class)
* [Extra](#extra)

### <span id="owner"></span> owner


> Owners are the maintainers of timeseries and nomenclatures.

  




* composed type [Class](#class)
* inlined member (*AO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| email | email (formatted string)| `strfmt.Email` | ✓ | |  |  |
| name | string| `string` | ✓ | |  |  |
| ownerStatus | [Ostatus](#ostatus)| `Ostatus` | ✓ | |  |  |
| uuid | uuid (formatted string)| `strfmt.UUID` | ✓ | |  |  |



### <span id="period"></span> period


> A period describe the periodicity of a time series, e.g. quarterly, yearly, ...

  




* composed type [Class](#class)
* [Extra](#extra)

### <span id="search-tags"></span> searchTags


  

[]string

### <span id="series"></span> series


> A time series.

The brief description omits descriptionLong, statusChangeReason, tags, linkedDocuments and auditTrail.

The field auditTrail is read-only and requires the "audit" query parameter to be retrieved.

The field seriesHasVersions requires the "deep" query parameter to be retrieved.

  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| auditTrail | [Audit](#audit)| `Audit` |  | |  |  |
| dataSource | [Source](#source)| `Source` | ✓ | |  |  |
| descriptionLong | [Translation](#translation)| `Translation` |  | |  |  |
| descriptionShort | [Translation](#translation)| `Translation` |  | |  |  |
| id | integer| `int64` | ✓ | |  |  |
| inputComposedUnit | [CompositeUnit](#composite-unit)| `CompositeUnit` |  | |  |  |
| linkedDocuments | [Documents](#documents)| `Documents` |  | |  |  |
| measurementUnit | [Munit](#munit)| `Munit` | ✓ | |  |  |
| seriesHasOwners | [][Owner](#owner)| `[]Owner` |  | |  |  |
| seriesHasThemes | [][Theme](#theme)| `[]Theme` |  | |  |  |
| seriesHasVersions | [][VersionedSeries](#versioned-series)| `[]VersionedSeries` |  | |  |  |
| status | [Vstatus](#vstatus)| `Vstatus` | ✓ | |  |  |
| statusChangeReason | [Translation](#translation)| `Translation` |  | |  |  |
| tags | [SearchTags](#search-tags)| `SearchTags` |  | |  |  |
| timePeriod | [Period](#period)| `Period` | ✓ | |  |  |
| title | [Translation](#translation)| `Translation` | ✓ | |  |  |
| zone | [Zone](#zone)| `Zone` | ✓ | |  |  |



### <span id="source"></span> source


> Data sources.

  




* composed type [Class](#class)
* inlined member (*AO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| rating | int64 (formatted integer)| `int64` |  | |  |  |
| tags | []string| `[]string` |  | |  |  |



### <span id="theme"></span> theme


> Climate themes.

  




* composed type [Class](#class)
* inlined member (*AO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| linkedDocuments | [Documents](#documents)| `Documents` |  | |  |  |
| tags | []string| `[]string` |  | |  |  |



### <span id="translation"></span> translation


> Translation is a multi-language string.

  



[Translation](#translation)

### <span id="ts-value"></span> tsValue


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| t | [Timestamp](#timestamp)| `extra.Timestamp` |  | |  |  |
| v | float (formatted number)| `float32` |  | |  |  |



### <span id="ts-values"></span> tsValues


  

[][TsValue](#ts-value)

### <span id="versioned-series"></span> versionedSeries


> A version of a time series.

The brief description omits versionDescriptionLong, versionStatusChangeReason, tags, linkedDocuments and auditTrail.

The field auditTrail is read-only and requires the "audit" query parameter to be retrieved.

The field timeseries requires the "deep" query parameter to be retrieved.

  





**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| auditTrail | [Audit](#audit)| `Audit` |  | |  |  |
| formula | string| `string` |  | |  |  |
| parentVersionedId | uuid (formatted string)| `strfmt.UUID` |  | |  |  |
| timeseries | [TsValues](#ts-values)| `TsValues` |  | |  |  |
| version | [Semver](#semver)| `extra.Semver` | ✓ | |  |  |
| versionDescriptionLong | [Translation](#translation)| `Translation` |  | |  |  |
| versionDescriptionShort | [Translation](#translation)| `Translation` |  | |  |  |
| versionOwner | uuid (formatted string)| `strfmt.UUID` |  | |  |  |
| versionStatus | [Vstatus](#vstatus)| `Vstatus` | ✓ | |  |  |
| versionStatusChangeReason | [Translation](#translation)| `Translation` |  | |  |  |
| versionTitle | [Translation](#translation)| `Translation` |  | |  |  |
| versionedId | uuid (formatted string)| `strfmt.UUID` | ✓ | |  |  |



### <span id="vstatus"></span> vstatus


  

| Name | Type | Go type | Default | Description | Example |
|------|------|---------| ------- |-------------|---------|
| vstatus | string| string | |  |  |



### <span id="zone"></span> zone


> Geographical zones.

  




* composed type [Class](#class)
* inlined member (*AO1*)



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| zoneGeometry | [Geometry](#geometry)| `Geometry` |  | |  |  |
| zoneType | [Ztype](#ztype)| `Ztype` | ✓ | |  |  |



### <span id="ztype"></span> ztype


> Typology of geographical zones (e.g. world, country, ...)

  




* composed type [Class](#class)
* [Extra](#extra)
