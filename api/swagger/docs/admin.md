


# Climate timeseries - Admin API
The timeseries admin API allows to maintain the nomenclatures used by Climate Timeseries,
such as units etc.

Timeseries publication status and ownership may be modified.

Only extra administrative operations are exposed here (e.g. updating nomenclatures).

Regular data retrieval remains carried out by the public API.

  
> [Climate Timeseries API concepts](https://github.io/wiki/TODO)

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


  ### <span id="tag-conversions"></span>conversions

Unit conversion logic.


## Content negotiation

### URI Schemes
  * https

### Consumes
  * application/json

### Produces
  * application/json

## Access control

### Security Schemes

#### bearerToken



> **Type**: oauth2
>
> **Flow**: accessCode
>
> **Authorization URL**: https://myidentityprovider.org/protocol/openid-connect/authorize
>
> **Token URL**: https://myidentityprovider/protocol/openid-connect/token
      

##### Scopes

Name | Description
-----|-------------
openid | Authorize the API to inquire about standard OpenID claims
email | Authorize the API to inquire about the user's email

### Security Requirements
  * bearerToken

## All endpoints

###  classes

  
> [API concepts: classes.](https://github.io/wiki/TODO)

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| DELETE | /admin/v1/classes/{classId}/members/{classMemberId} | [delete classes class ID members class member ID](#delete-classes-class-id-members-class-member-id) | Removes a member from a nomenclature class |
| POST | /admin/v1/classes/{classId}/members | [post classes class ID members](#post-classes-class-id-members) | Add a new member in a nomenclature class |
| PUT | /admin/v1/classes/{classId} | [put classes class ID](#put-classes-class-id) | Update descriptive metadata about a nomenclature class |
| PUT | /admin/v1/classes/{classId}/members/{classMemberId} | [put classes class ID members class member ID](#put-classes-class-id-members-class-member-id) | Update a member of a nomenclature class |
  


###  conversions

  

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /admin/v1/conversion | [post conversion](#post-conversion) | Creates add new conversion specification from unit to unit. |
| POST | /admin/v1/conversion/{fromUnit}/{toUnit} | [post conversion from unit to unit](#post-conversion-from-unit-to-unit) | Creates add new conversion specification from unit to unit. |
| PUT | /admin/v1/conversion | [put conversion](#put-conversion) | Update an existing conversion specification from unit to unit. |
| PUT | /admin/v1/conversion/{fromUnit}/{toUnit} | [put conversion from unit to unit](#put-conversion-from-unit-to-unit) | Update an existing conversion specification from unit to unit. |
  


## Paths

### <span id="delete-classes-class-id-members-class-member-id"></span> Removes a member from a nomenclature class (*DeleteClassesClassIDMembersClassMemberID*)

```
DELETE /admin/v1/classes/{classId}/members/{classMemberId}
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
| classMemberId | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | The unique identifier of a class member. |
| classMember | `body` | [Class](#class) | `models.Class` | | ✓ | | Class member metadata. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#delete-classes-class-id-members-class-member-id-204) | No Content | Class members successfully deleted. |  | [schema](#delete-classes-class-id-members-class-member-id-204-schema) |
| [400](#delete-classes-class-id-members-class-member-id-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#delete-classes-class-id-members-class-member-id-400-schema) |
| [401](#delete-classes-class-id-members-class-member-id-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#delete-classes-class-id-members-class-member-id-401-schema) |
| [403](#delete-classes-class-id-members-class-member-id-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#delete-classes-class-id-members-class-member-id-403-schema) |
| [404](#delete-classes-class-id-members-class-member-id-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#delete-classes-class-id-members-class-member-id-404-schema) |
| [405](#delete-classes-class-id-members-class-member-id-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#delete-classes-class-id-members-class-member-id-405-schema) |
| [500](#delete-classes-class-id-members-class-member-id-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#delete-classes-class-id-members-class-member-id-500-schema) |
| [default](#delete-classes-class-id-members-class-member-id-default) | | Other error. See error details. |  | [schema](#delete-classes-class-id-members-class-member-id-default-schema) |

#### Responses


##### <span id="delete-classes-class-id-members-class-member-id-204"></span> 204 - Class members successfully deleted.
Status: No Content

###### <span id="delete-classes-class-id-members-class-member-id-204-schema"></span> Schema

##### <span id="delete-classes-class-id-members-class-member-id-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="delete-classes-class-id-members-class-member-id-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="delete-classes-class-id-members-class-member-id-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="delete-classes-class-id-members-class-member-id-401-schema"></span> Schema

##### <span id="delete-classes-class-id-members-class-member-id-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="delete-classes-class-id-members-class-member-id-403-schema"></span> Schema

##### <span id="delete-classes-class-id-members-class-member-id-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="delete-classes-class-id-members-class-member-id-404-schema"></span> Schema

##### <span id="delete-classes-class-id-members-class-member-id-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="delete-classes-class-id-members-class-member-id-405-schema"></span> Schema

##### <span id="delete-classes-class-id-members-class-member-id-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="delete-classes-class-id-members-class-member-id-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="delete-classes-class-id-members-class-member-id-default"></span> Default Response
Other error. See error details.

###### <span id="delete-classes-class-id-members-class-member-id-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="post-classes-class-id-members"></span> Add a new member in a nomenclature class (*PostClassesClassIDMembers*)

```
POST /admin/v1/classes/{classId}/members
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
| classMember | `body` | [Class](#class) | `models.Class` | | ✓ | | Class member metadata. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#post-classes-class-id-members-201) | Created | Class members added. | ✓ | [schema](#post-classes-class-id-members-201-schema) |
| [400](#post-classes-class-id-members-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#post-classes-class-id-members-400-schema) |
| [401](#post-classes-class-id-members-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#post-classes-class-id-members-401-schema) |
| [403](#post-classes-class-id-members-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#post-classes-class-id-members-403-schema) |
| [404](#post-classes-class-id-members-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#post-classes-class-id-members-404-schema) |
| [405](#post-classes-class-id-members-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#post-classes-class-id-members-405-schema) |
| [409](#post-classes-class-id-members-409) | Conflict | Resource already exists. An object creation was requested, but this object was already existing. |  | [schema](#post-classes-class-id-members-409-schema) |
| [500](#post-classes-class-id-members-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#post-classes-class-id-members-500-schema) |
| [default](#post-classes-class-id-members-default) | | Other error. See error details. |  | [schema](#post-classes-class-id-members-default-schema) |

#### Responses


##### <span id="post-classes-class-id-members-201"></span> 201 - Class members added.
Status: Created

###### <span id="post-classes-class-id-members-201-schema"></span> Schema

###### Response headers

| Name | Type | Go type | Separator | Default | Description |
|------|------|---------|-----------|---------|-------------|
| Location | uri (formatted string) | `strfmt.URI` |  |  | The URI of the newly created resource. |
| X-ID | int64 (formatted integer) | `int64` |  |  | The ID of the newly created resource. |

##### <span id="post-classes-class-id-members-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="post-classes-class-id-members-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="post-classes-class-id-members-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="post-classes-class-id-members-401-schema"></span> Schema

##### <span id="post-classes-class-id-members-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="post-classes-class-id-members-403-schema"></span> Schema

##### <span id="post-classes-class-id-members-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="post-classes-class-id-members-404-schema"></span> Schema

##### <span id="post-classes-class-id-members-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="post-classes-class-id-members-405-schema"></span> Schema

##### <span id="post-classes-class-id-members-409"></span> 409 - Resource already exists. An object creation was requested, but this object was already existing.
Status: Conflict

###### <span id="post-classes-class-id-members-409-schema"></span> Schema

##### <span id="post-classes-class-id-members-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="post-classes-class-id-members-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="post-classes-class-id-members-default"></span> Default Response
Other error. See error details.

###### <span id="post-classes-class-id-members-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="post-conversion"></span> Creates add new conversion specification from unit to unit. (*PostConversion*)

```
POST /admin/v1/conversion
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| fromUnit | `query` | string | `string` |  | ✓ |  | Original unit to be converted. |
| toUnit | `query` | string | `string` |  | ✓ |  | Original unit to be converted. |
| conversionSpec | `body` | [ConversionSpec](#conversion-spec) | `models.ConversionSpec` | | ✓ | | Unit conversion specification. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#post-conversion-201) | Created | Unit conversion added |  | [schema](#post-conversion-201-schema) |
| [400](#post-conversion-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#post-conversion-400-schema) |
| [401](#post-conversion-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#post-conversion-401-schema) |
| [403](#post-conversion-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#post-conversion-403-schema) |
| [404](#post-conversion-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#post-conversion-404-schema) |
| [405](#post-conversion-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#post-conversion-405-schema) |
| [409](#post-conversion-409) | Conflict | Resource already exists. An object creation was requested, but this object was already existing. |  | [schema](#post-conversion-409-schema) |
| [500](#post-conversion-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#post-conversion-500-schema) |
| [default](#post-conversion-default) | | Other error. See error details. |  | [schema](#post-conversion-default-schema) |

#### Responses


##### <span id="post-conversion-201"></span> 201 - Unit conversion added
Status: Created

###### <span id="post-conversion-201-schema"></span> Schema

##### <span id="post-conversion-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="post-conversion-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="post-conversion-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="post-conversion-401-schema"></span> Schema

##### <span id="post-conversion-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="post-conversion-403-schema"></span> Schema

##### <span id="post-conversion-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="post-conversion-404-schema"></span> Schema

##### <span id="post-conversion-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="post-conversion-405-schema"></span> Schema

##### <span id="post-conversion-409"></span> 409 - Resource already exists. An object creation was requested, but this object was already existing.
Status: Conflict

###### <span id="post-conversion-409-schema"></span> Schema

##### <span id="post-conversion-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="post-conversion-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="post-conversion-default"></span> Default Response
Other error. See error details.

###### <span id="post-conversion-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="post-conversion-from-unit-to-unit"></span> Creates add new conversion specification from unit to unit. (*PostConversionFromUnitToUnit*)

```
POST /admin/v1/conversion/{fromUnit}/{toUnit}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| fromUnit | `path` | string | `string` |  | ✓ |  | Original unit to be converted. |
| toUnit | `path` | string | `string` |  | ✓ |  | Original unit to be converted. |
| conversionSpec | `body` | [ConversionSpec](#conversion-spec) | `models.ConversionSpec` | | ✓ | | Unit conversion specification. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#post-conversion-from-unit-to-unit-201) | Created | Unit conversion added |  | [schema](#post-conversion-from-unit-to-unit-201-schema) |
| [400](#post-conversion-from-unit-to-unit-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#post-conversion-from-unit-to-unit-400-schema) |
| [401](#post-conversion-from-unit-to-unit-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#post-conversion-from-unit-to-unit-401-schema) |
| [403](#post-conversion-from-unit-to-unit-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#post-conversion-from-unit-to-unit-403-schema) |
| [404](#post-conversion-from-unit-to-unit-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#post-conversion-from-unit-to-unit-404-schema) |
| [405](#post-conversion-from-unit-to-unit-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#post-conversion-from-unit-to-unit-405-schema) |
| [409](#post-conversion-from-unit-to-unit-409) | Conflict | Resource already exists. An object creation was requested, but this object was already existing. |  | [schema](#post-conversion-from-unit-to-unit-409-schema) |
| [500](#post-conversion-from-unit-to-unit-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#post-conversion-from-unit-to-unit-500-schema) |
| [default](#post-conversion-from-unit-to-unit-default) | | Other error. See error details. |  | [schema](#post-conversion-from-unit-to-unit-default-schema) |

#### Responses


##### <span id="post-conversion-from-unit-to-unit-201"></span> 201 - Unit conversion added
Status: Created

###### <span id="post-conversion-from-unit-to-unit-201-schema"></span> Schema

##### <span id="post-conversion-from-unit-to-unit-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="post-conversion-from-unit-to-unit-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="post-conversion-from-unit-to-unit-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="post-conversion-from-unit-to-unit-401-schema"></span> Schema

##### <span id="post-conversion-from-unit-to-unit-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="post-conversion-from-unit-to-unit-403-schema"></span> Schema

##### <span id="post-conversion-from-unit-to-unit-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="post-conversion-from-unit-to-unit-404-schema"></span> Schema

##### <span id="post-conversion-from-unit-to-unit-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="post-conversion-from-unit-to-unit-405-schema"></span> Schema

##### <span id="post-conversion-from-unit-to-unit-409"></span> 409 - Resource already exists. An object creation was requested, but this object was already existing.
Status: Conflict

###### <span id="post-conversion-from-unit-to-unit-409-schema"></span> Schema

##### <span id="post-conversion-from-unit-to-unit-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="post-conversion-from-unit-to-unit-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="post-conversion-from-unit-to-unit-default"></span> Default Response
Other error. See error details.

###### <span id="post-conversion-from-unit-to-unit-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="put-classes-class-id"></span> Update descriptive metadata about a nomenclature class (*PutClassesClassID*)

```
PUT /admin/v1/classes/{classId}
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
| classDescription | `body` | [ClassDescription](#class-description) | `models.ClassDescription` | | ✓ | | Class descriptive metadata. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#put-classes-class-id-204) | No Content | Metadata about nomenclature classes successfully updated. |  | [schema](#put-classes-class-id-204-schema) |
| [400](#put-classes-class-id-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#put-classes-class-id-400-schema) |
| [401](#put-classes-class-id-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#put-classes-class-id-401-schema) |
| [403](#put-classes-class-id-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#put-classes-class-id-403-schema) |
| [404](#put-classes-class-id-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#put-classes-class-id-404-schema) |
| [405](#put-classes-class-id-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#put-classes-class-id-405-schema) |
| [500](#put-classes-class-id-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#put-classes-class-id-500-schema) |
| [default](#put-classes-class-id-default) | | Other error. See error details. |  | [schema](#put-classes-class-id-default-schema) |

#### Responses


##### <span id="put-classes-class-id-204"></span> 204 - Metadata about nomenclature classes successfully updated.
Status: No Content

###### <span id="put-classes-class-id-204-schema"></span> Schema
   
  

[ClassDescription](#class-description)

##### <span id="put-classes-class-id-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="put-classes-class-id-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="put-classes-class-id-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="put-classes-class-id-401-schema"></span> Schema

##### <span id="put-classes-class-id-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="put-classes-class-id-403-schema"></span> Schema

##### <span id="put-classes-class-id-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="put-classes-class-id-404-schema"></span> Schema

##### <span id="put-classes-class-id-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="put-classes-class-id-405-schema"></span> Schema

##### <span id="put-classes-class-id-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="put-classes-class-id-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="put-classes-class-id-default"></span> Default Response
Other error. See error details.

###### <span id="put-classes-class-id-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="put-classes-class-id-members-class-member-id"></span> Update a member of a nomenclature class (*PutClassesClassIDMembersClassMemberID*)

```
PUT /admin/v1/classes/{classId}/members/{classMemberId}
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
| classMemberId | `path` | int64 (formatted integer) | `int64` |  | ✓ |  | The unique identifier of a class member. |
| classMember | `body` | [Class](#class) | `models.Class` | | ✓ | | Class member metadata. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#put-classes-class-id-members-class-member-id-204) | No Content | Class members updated. | ✓ | [schema](#put-classes-class-id-members-class-member-id-204-schema) |
| [400](#put-classes-class-id-members-class-member-id-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#put-classes-class-id-members-class-member-id-400-schema) |
| [401](#put-classes-class-id-members-class-member-id-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#put-classes-class-id-members-class-member-id-401-schema) |
| [403](#put-classes-class-id-members-class-member-id-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#put-classes-class-id-members-class-member-id-403-schema) |
| [404](#put-classes-class-id-members-class-member-id-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#put-classes-class-id-members-class-member-id-404-schema) |
| [405](#put-classes-class-id-members-class-member-id-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#put-classes-class-id-members-class-member-id-405-schema) |
| [500](#put-classes-class-id-members-class-member-id-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#put-classes-class-id-members-class-member-id-500-schema) |
| [default](#put-classes-class-id-members-class-member-id-default) | | Other error. See error details. |  | [schema](#put-classes-class-id-members-class-member-id-default-schema) |

#### Responses


##### <span id="put-classes-class-id-members-class-member-id-204"></span> 204 - Class members updated.
Status: No Content

###### <span id="put-classes-class-id-members-class-member-id-204-schema"></span> Schema

###### Response headers

| Name | Type | Go type | Separator | Default | Description |
|------|------|---------|-----------|---------|-------------|
| Location | uri (formatted string) | `strfmt.URI` |  |  | The URI of the updated resource. |
| X-ID | int64 (formatted integer) | `int64` |  |  | The ID of the updated resource. |

##### <span id="put-classes-class-id-members-class-member-id-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="put-classes-class-id-members-class-member-id-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="put-classes-class-id-members-class-member-id-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="put-classes-class-id-members-class-member-id-401-schema"></span> Schema

##### <span id="put-classes-class-id-members-class-member-id-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="put-classes-class-id-members-class-member-id-403-schema"></span> Schema

##### <span id="put-classes-class-id-members-class-member-id-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="put-classes-class-id-members-class-member-id-404-schema"></span> Schema

##### <span id="put-classes-class-id-members-class-member-id-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="put-classes-class-id-members-class-member-id-405-schema"></span> Schema

##### <span id="put-classes-class-id-members-class-member-id-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="put-classes-class-id-members-class-member-id-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="put-classes-class-id-members-class-member-id-default"></span> Default Response
Other error. See error details.

###### <span id="put-classes-class-id-members-class-member-id-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="put-conversion"></span> Update an existing conversion specification from unit to unit. (*PutConversion*)

```
PUT /admin/v1/conversion
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| fromUnit | `query` | string | `string` |  | ✓ |  | Original unit to be converted. |
| toUnit | `query` | string | `string` |  | ✓ |  | Original unit to be converted. |
| conversionSpec | `body` | [ConversionSpec](#conversion-spec) | `models.ConversionSpec` | | ✓ | | Unit conversion specification. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#put-conversion-204) | No Content | Unit conversion updated. |  | [schema](#put-conversion-204-schema) |
| [400](#put-conversion-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#put-conversion-400-schema) |
| [401](#put-conversion-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#put-conversion-401-schema) |
| [403](#put-conversion-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#put-conversion-403-schema) |
| [404](#put-conversion-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#put-conversion-404-schema) |
| [405](#put-conversion-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#put-conversion-405-schema) |
| [500](#put-conversion-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#put-conversion-500-schema) |
| [default](#put-conversion-default) | | Other error. See error details. |  | [schema](#put-conversion-default-schema) |

#### Responses


##### <span id="put-conversion-204"></span> 204 - Unit conversion updated.
Status: No Content

###### <span id="put-conversion-204-schema"></span> Schema

##### <span id="put-conversion-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="put-conversion-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="put-conversion-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="put-conversion-401-schema"></span> Schema

##### <span id="put-conversion-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="put-conversion-403-schema"></span> Schema

##### <span id="put-conversion-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="put-conversion-404-schema"></span> Schema

##### <span id="put-conversion-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="put-conversion-405-schema"></span> Schema

##### <span id="put-conversion-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="put-conversion-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="put-conversion-default"></span> Default Response
Other error. See error details.

###### <span id="put-conversion-default-schema"></span> Schema

  

[APIError](#api-error)

### <span id="put-conversion-from-unit-to-unit"></span> Update an existing conversion specification from unit to unit. (*PutConversionFromUnitToUnit*)

```
PUT /admin/v1/conversion/{fromUnit}/{toUnit}
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| fromUnit | `path` | string | `string` |  | ✓ |  | Original unit to be converted. |
| toUnit | `path` | string | `string` |  | ✓ |  | Original unit to be converted. |
| conversionSpec | `body` | [ConversionSpec](#conversion-spec) | `models.ConversionSpec` | | ✓ | | Unit conversion specification. |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#put-conversion-from-unit-to-unit-204) | No Content | Unit conversion updated. |  | [schema](#put-conversion-from-unit-to-unit-204-schema) |
| [400](#put-conversion-from-unit-to-unit-400) | Bad Request | Client error in request. Input did not pass validations. See error details. |  | [schema](#put-conversion-from-unit-to-unit-400-schema) |
| [401](#put-conversion-from-unit-to-unit-401) | Unauthorized | Unauthorized access for a lack of authentication |  | [schema](#put-conversion-from-unit-to-unit-401-schema) |
| [403](#put-conversion-from-unit-to-unit-403) | Forbidden | Forbidden access for a lack of sufficient privileges |  | [schema](#put-conversion-from-unit-to-unit-403-schema) |
| [404](#put-conversion-from-unit-to-unit-404) | Not Found | Resource not found. The object requested does not exist or is not visible to the user. |  | [schema](#put-conversion-from-unit-to-unit-404-schema) |
| [405](#put-conversion-from-unit-to-unit-405) | Method Not Allowed | MethodNotAllowed |  | [schema](#put-conversion-from-unit-to-unit-405-schema) |
| [500](#put-conversion-from-unit-to-unit-500) | Internal Server Error | An internal error has occured. See error details. |  | [schema](#put-conversion-from-unit-to-unit-500-schema) |
| [default](#put-conversion-from-unit-to-unit-default) | | Other error. See error details. |  | [schema](#put-conversion-from-unit-to-unit-default-schema) |

#### Responses


##### <span id="put-conversion-from-unit-to-unit-204"></span> 204 - Unit conversion updated.
Status: No Content

###### <span id="put-conversion-from-unit-to-unit-204-schema"></span> Schema

##### <span id="put-conversion-from-unit-to-unit-400"></span> 400 - Client error in request. Input did not pass validations. See error details.
Status: Bad Request

###### <span id="put-conversion-from-unit-to-unit-400-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="put-conversion-from-unit-to-unit-401"></span> 401 - Unauthorized access for a lack of authentication
Status: Unauthorized

###### <span id="put-conversion-from-unit-to-unit-401-schema"></span> Schema

##### <span id="put-conversion-from-unit-to-unit-403"></span> 403 - Forbidden access for a lack of sufficient privileges
Status: Forbidden

###### <span id="put-conversion-from-unit-to-unit-403-schema"></span> Schema

##### <span id="put-conversion-from-unit-to-unit-404"></span> 404 - Resource not found. The object requested does not exist or is not visible to the user.
Status: Not Found

###### <span id="put-conversion-from-unit-to-unit-404-schema"></span> Schema

##### <span id="put-conversion-from-unit-to-unit-405"></span> 405 - MethodNotAllowed
Status: Method Not Allowed

###### <span id="put-conversion-from-unit-to-unit-405-schema"></span> Schema

##### <span id="put-conversion-from-unit-to-unit-500"></span> 500 - An internal error has occured. See error details.
Status: Internal Server Error

###### <span id="put-conversion-from-unit-to-unit-500-schema"></span> Schema
   
  

[APIError](#api-error)

##### <span id="put-conversion-from-unit-to-unit-default"></span> Default Response
Other error. See error details.

###### <span id="put-conversion-from-unit-to-unit-default-schema"></span> Schema

  

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



### <span id="translation"></span> translation


> Translation is a multi-language string.

  



[Translation](#translation)
