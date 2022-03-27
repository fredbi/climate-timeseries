# swagger_client.ClassesApi

All URIs are relative to *https://api.example.com/v1/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**classes_class_id_get**](ClassesApi.md#classes_class_id_get) | **GET** /classes/{classId} | Get metadata about a nomenclature class
[**classes_class_id_members_class_member_id_delete**](ClassesApi.md#classes_class_id_members_class_member_id_delete) | **DELETE** /classes/{classId}/members/{classMemberId} | Removes a member from a nomenclature class
[**classes_class_id_members_class_member_id_put**](ClassesApi.md#classes_class_id_members_class_member_id_put) | **PUT** /classes/{classId}/members/{classMemberId} | Update a member of a nomenclature class
[**classes_class_id_members_get**](ClassesApi.md#classes_class_id_members_get) | **GET** /classes/{classId}/members | Get all the members of a nomenclature class
[**classes_class_id_members_post**](ClassesApi.md#classes_class_id_members_post) | **POST** /classes/{classId}/members | Add a new member in a nomenclature class
[**classes_class_id_put**](ClassesApi.md#classes_class_id_put) | **PUT** /classes/{classId} | Update descriptive metadata about a nomenclature class
[**classes_get**](ClassesApi.md#classes_get) | **GET** /classes | Get all valid nomenclature classes


# **classes_class_id_get**
> ClassDescription classes_class_id_get(class_id, deep=deep, brief=brief, audit=audit)

Get metadata about a nomenclature class

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.ClassesApi()
class_id = 'class_id_example' # str | The internal name of a nomenclature class.  Valid classes are:   * constant: mathematical and physical constants   * mdimension: base measured dimensions   * mdomain: domains that pertain to measurements   * measurement: physical and economic measurements   * multiplier: unit multipliers aka prefixes (e.g. k,M,G...)   * munit: measurement units   * musystem: systems of measurement units   * ostatus: owner statuses   * owner: series owner   * period: time series periods (e.g. monthly, yearly...)   * role: series owner role   * source: data sources   * status: series and versions statuses   * theme: climate themes   * zone: geographical zones   * ztype: zone types 
deep = true # bool | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  (optional)
brief = true # bool | When brief is specified, the response will only contain essential data and strip long descriptions.  (optional)
audit = true # bool | When audit is specified, the response will also contain the audit trail field.  (optional)

try:
    # Get metadata about a nomenclature class
    api_response = api_instance.classes_class_id_get(class_id, deep=deep, brief=brief, audit=audit)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ClassesApi->classes_class_id_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **class_id** | **str**| The internal name of a nomenclature class.  Valid classes are:   * constant: mathematical and physical constants   * mdimension: base measured dimensions   * mdomain: domains that pertain to measurements   * measurement: physical and economic measurements   * multiplier: unit multipliers aka prefixes (e.g. k,M,G...)   * munit: measurement units   * musystem: systems of measurement units   * ostatus: owner statuses   * owner: series owner   * period: time series periods (e.g. monthly, yearly...)   * role: series owner role   * source: data sources   * status: series and versions statuses   * theme: climate themes   * zone: geographical zones   * ztype: zone types  | 
 **deep** | **bool**| When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  | [optional] 
 **brief** | **bool**| When brief is specified, the response will only contain essential data and strip long descriptions.  | [optional] 
 **audit** | **bool**| When audit is specified, the response will also contain the audit trail field.  | [optional] 

### Return type

[**ClassDescription**](ClassDescription.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **classes_class_id_members_class_member_id_delete**
> classes_class_id_members_class_member_id_delete(class_id, class_member_id, class_member)

Removes a member from a nomenclature class

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# Configure API key authorization: apiKey
configuration = swagger_client.Configuration()
configuration.api_key['X-API-Key'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-API-Key'] = 'Bearer'
# Configure OAuth2 access token for authorization: bearerToken
configuration = swagger_client.Configuration()
configuration.access_token = 'YOUR_ACCESS_TOKEN'

# create an instance of the API class
api_instance = swagger_client.ClassesApi(swagger_client.ApiClient(configuration))
class_id = 'class_id_example' # str | The internal name of a nomenclature class.  Valid classes are:   * constant: mathematical and physical constants   * mdimension: base measured dimensions   * mdomain: domains that pertain to measurements   * measurement: physical and economic measurements   * multiplier: unit multipliers aka prefixes (e.g. k,M,G...)   * munit: measurement units   * musystem: systems of measurement units   * ostatus: owner statuses   * owner: series owner   * period: time series periods (e.g. monthly, yearly...)   * role: series owner role   * source: data sources   * status: series and versions statuses   * theme: climate themes   * zone: geographical zones   * ztype: zone types 
class_member_id = 789 # int | The unique identifier of a class member. 
class_member = swagger_client.ClassNomenclature() # ClassNomenclature | Class member metadata. 

try:
    # Removes a member from a nomenclature class
    api_instance.classes_class_id_members_class_member_id_delete(class_id, class_member_id, class_member)
except ApiException as e:
    print("Exception when calling ClassesApi->classes_class_id_members_class_member_id_delete: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **class_id** | **str**| The internal name of a nomenclature class.  Valid classes are:   * constant: mathematical and physical constants   * mdimension: base measured dimensions   * mdomain: domains that pertain to measurements   * measurement: physical and economic measurements   * multiplier: unit multipliers aka prefixes (e.g. k,M,G...)   * munit: measurement units   * musystem: systems of measurement units   * ostatus: owner statuses   * owner: series owner   * period: time series periods (e.g. monthly, yearly...)   * role: series owner role   * source: data sources   * status: series and versions statuses   * theme: climate themes   * zone: geographical zones   * ztype: zone types  | 
 **class_member_id** | **int**| The unique identifier of a class member.  | 
 **class_member** | [**ClassNomenclature**](ClassNomenclature.md)| Class member metadata.  | 

### Return type

void (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **classes_class_id_members_class_member_id_put**
> classes_class_id_members_class_member_id_put(class_id, class_member_id, class_member)

Update a member of a nomenclature class

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# Configure API key authorization: apiKey
configuration = swagger_client.Configuration()
configuration.api_key['X-API-Key'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-API-Key'] = 'Bearer'
# Configure OAuth2 access token for authorization: bearerToken
configuration = swagger_client.Configuration()
configuration.access_token = 'YOUR_ACCESS_TOKEN'

# create an instance of the API class
api_instance = swagger_client.ClassesApi(swagger_client.ApiClient(configuration))
class_id = 'class_id_example' # str | The internal name of a nomenclature class.  Valid classes are:   * constant: mathematical and physical constants   * mdimension: base measured dimensions   * mdomain: domains that pertain to measurements   * measurement: physical and economic measurements   * multiplier: unit multipliers aka prefixes (e.g. k,M,G...)   * munit: measurement units   * musystem: systems of measurement units   * ostatus: owner statuses   * owner: series owner   * period: time series periods (e.g. monthly, yearly...)   * role: series owner role   * source: data sources   * status: series and versions statuses   * theme: climate themes   * zone: geographical zones   * ztype: zone types 
class_member_id = 789 # int | The unique identifier of a class member. 
class_member = swagger_client.ClassNomenclature() # ClassNomenclature | Class member metadata. 

try:
    # Update a member of a nomenclature class
    api_instance.classes_class_id_members_class_member_id_put(class_id, class_member_id, class_member)
except ApiException as e:
    print("Exception when calling ClassesApi->classes_class_id_members_class_member_id_put: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **class_id** | **str**| The internal name of a nomenclature class.  Valid classes are:   * constant: mathematical and physical constants   * mdimension: base measured dimensions   * mdomain: domains that pertain to measurements   * measurement: physical and economic measurements   * multiplier: unit multipliers aka prefixes (e.g. k,M,G...)   * munit: measurement units   * musystem: systems of measurement units   * ostatus: owner statuses   * owner: series owner   * period: time series periods (e.g. monthly, yearly...)   * role: series owner role   * source: data sources   * status: series and versions statuses   * theme: climate themes   * zone: geographical zones   * ztype: zone types  | 
 **class_member_id** | **int**| The unique identifier of a class member.  | 
 **class_member** | [**ClassNomenclature**](ClassNomenclature.md)| Class member metadata.  | 

### Return type

void (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **classes_class_id_members_get**
> list[ClassNomenclature] classes_class_id_members_get(class_id, glob=glob, deep=deep, brief=brief, audit=audit, tag=tag, tags=tags)

Get all the members of a nomenclature class

Returns a a collection of nomenclature class members.  Example:   GET /classes/munit/members returns all measurement units.  If the request negotiates a response MIME type with text/csv (with the Accept header), this endpoint produces a CSV file. 

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.ClassesApi()
class_id = 'class_id_example' # str | The internal name of a nomenclature class.  Valid classes are:   * constant: mathematical and physical constants   * mdimension: base measured dimensions   * mdomain: domains that pertain to measurements   * measurement: physical and economic measurements   * multiplier: unit multipliers aka prefixes (e.g. k,M,G...)   * munit: measurement units   * musystem: systems of measurement units   * ostatus: owner statuses   * owner: series owner   * period: time series periods (e.g. monthly, yearly...)   * role: series owner role   * source: data sources   * status: series and versions statuses   * theme: climate themes   * zone: geographical zones   * ztype: zone types 
glob = 'glob_example' # str | Filters the result according to a glob pattern applied on the short name of the requested object.  (optional)
deep = true # bool | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  (optional)
brief = true # bool | When brief is specified, the response will only contain essential data and strip long descriptions.  (optional)
audit = true # bool | When audit is specified, the response will also contain the audit trail field.  (optional)
tag = 'tag_example' # str | Filter the result by search for a given keyword, whenever tag search is applicable.  This parameter has no effect on objects where tag search is not applicable.  (optional)
tags = ['tags_example'] # list[str] | Filter the result by search for one tag in the given list of keywords, whenever tag search is applicable.  Up to 5 tags can be searched.  This parameter has no effect on objects where tag search is not applicable.  (optional)

try:
    # Get all the members of a nomenclature class
    api_response = api_instance.classes_class_id_members_get(class_id, glob=glob, deep=deep, brief=brief, audit=audit, tag=tag, tags=tags)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ClassesApi->classes_class_id_members_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **class_id** | **str**| The internal name of a nomenclature class.  Valid classes are:   * constant: mathematical and physical constants   * mdimension: base measured dimensions   * mdomain: domains that pertain to measurements   * measurement: physical and economic measurements   * multiplier: unit multipliers aka prefixes (e.g. k,M,G...)   * munit: measurement units   * musystem: systems of measurement units   * ostatus: owner statuses   * owner: series owner   * period: time series periods (e.g. monthly, yearly...)   * role: series owner role   * source: data sources   * status: series and versions statuses   * theme: climate themes   * zone: geographical zones   * ztype: zone types  | 
 **glob** | **str**| Filters the result according to a glob pattern applied on the short name of the requested object.  | [optional] 
 **deep** | **bool**| When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  | [optional] 
 **brief** | **bool**| When brief is specified, the response will only contain essential data and strip long descriptions.  | [optional] 
 **audit** | **bool**| When audit is specified, the response will also contain the audit trail field.  | [optional] 
 **tag** | **str**| Filter the result by search for a given keyword, whenever tag search is applicable.  This parameter has no effect on objects where tag search is not applicable.  | [optional] 
 **tags** | [**list[str]**](str.md)| Filter the result by search for one tag in the given list of keywords, whenever tag search is applicable.  Up to 5 tags can be searched.  This parameter has no effect on objects where tag search is not applicable.  | [optional] 

### Return type

[**list[ClassNomenclature]**](ClassNomenclature.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **classes_class_id_members_post**
> classes_class_id_members_post(class_id, class_member)

Add a new member in a nomenclature class

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# Configure API key authorization: apiKey
configuration = swagger_client.Configuration()
configuration.api_key['X-API-Key'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-API-Key'] = 'Bearer'
# Configure OAuth2 access token for authorization: bearerToken
configuration = swagger_client.Configuration()
configuration.access_token = 'YOUR_ACCESS_TOKEN'

# create an instance of the API class
api_instance = swagger_client.ClassesApi(swagger_client.ApiClient(configuration))
class_id = 'class_id_example' # str | The internal name of a nomenclature class.  Valid classes are:   * constant: mathematical and physical constants   * mdimension: base measured dimensions   * mdomain: domains that pertain to measurements   * measurement: physical and economic measurements   * multiplier: unit multipliers aka prefixes (e.g. k,M,G...)   * munit: measurement units   * musystem: systems of measurement units   * ostatus: owner statuses   * owner: series owner   * period: time series periods (e.g. monthly, yearly...)   * role: series owner role   * source: data sources   * status: series and versions statuses   * theme: climate themes   * zone: geographical zones   * ztype: zone types 
class_member = swagger_client.ClassNomenclature() # ClassNomenclature | Class member metadata. 

try:
    # Add a new member in a nomenclature class
    api_instance.classes_class_id_members_post(class_id, class_member)
except ApiException as e:
    print("Exception when calling ClassesApi->classes_class_id_members_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **class_id** | **str**| The internal name of a nomenclature class.  Valid classes are:   * constant: mathematical and physical constants   * mdimension: base measured dimensions   * mdomain: domains that pertain to measurements   * measurement: physical and economic measurements   * multiplier: unit multipliers aka prefixes (e.g. k,M,G...)   * munit: measurement units   * musystem: systems of measurement units   * ostatus: owner statuses   * owner: series owner   * period: time series periods (e.g. monthly, yearly...)   * role: series owner role   * source: data sources   * status: series and versions statuses   * theme: climate themes   * zone: geographical zones   * ztype: zone types  | 
 **class_member** | [**ClassNomenclature**](ClassNomenclature.md)| Class member metadata.  | 

### Return type

void (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **classes_class_id_put**
> ClassDescription classes_class_id_put(class_id, class_description)

Update descriptive metadata about a nomenclature class

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# Configure API key authorization: apiKey
configuration = swagger_client.Configuration()
configuration.api_key['X-API-Key'] = 'YOUR_API_KEY'
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-API-Key'] = 'Bearer'
# Configure OAuth2 access token for authorization: bearerToken
configuration = swagger_client.Configuration()
configuration.access_token = 'YOUR_ACCESS_TOKEN'

# create an instance of the API class
api_instance = swagger_client.ClassesApi(swagger_client.ApiClient(configuration))
class_id = 'class_id_example' # str | The internal name of a nomenclature class.  Valid classes are:   * constant: mathematical and physical constants   * mdimension: base measured dimensions   * mdomain: domains that pertain to measurements   * measurement: physical and economic measurements   * multiplier: unit multipliers aka prefixes (e.g. k,M,G...)   * munit: measurement units   * musystem: systems of measurement units   * ostatus: owner statuses   * owner: series owner   * period: time series periods (e.g. monthly, yearly...)   * role: series owner role   * source: data sources   * status: series and versions statuses   * theme: climate themes   * zone: geographical zones   * ztype: zone types 
class_description = swagger_client.ClassDescription() # ClassDescription | Class descriptive metadata. 

try:
    # Update descriptive metadata about a nomenclature class
    api_response = api_instance.classes_class_id_put(class_id, class_description)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ClassesApi->classes_class_id_put: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **class_id** | **str**| The internal name of a nomenclature class.  Valid classes are:   * constant: mathematical and physical constants   * mdimension: base measured dimensions   * mdomain: domains that pertain to measurements   * measurement: physical and economic measurements   * multiplier: unit multipliers aka prefixes (e.g. k,M,G...)   * munit: measurement units   * musystem: systems of measurement units   * ostatus: owner statuses   * owner: series owner   * period: time series periods (e.g. monthly, yearly...)   * role: series owner role   * source: data sources   * status: series and versions statuses   * theme: climate themes   * zone: geographical zones   * ztype: zone types  | 
 **class_description** | [**ClassDescription**](ClassDescription.md)| Class descriptive metadata.  | 

### Return type

[**ClassDescription**](ClassDescription.md)

### Authorization

[apiKey](../README.md#apiKey), [bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **classes_get**
> list[ClassNomenclature] classes_get(glob=glob, deep=deep, brief=brief, audit=audit)

Get all valid nomenclature classes

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.ClassesApi()
glob = 'glob_example' # str | Filters the result according to a glob pattern applied on the short name of the requested object.  (optional)
deep = true # bool | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  (optional)
brief = true # bool | When brief is specified, the response will only contain essential data and strip long descriptions.  (optional)
audit = true # bool | When audit is specified, the response will also contain the audit trail field.  (optional)

try:
    # Get all valid nomenclature classes
    api_response = api_instance.classes_get(glob=glob, deep=deep, brief=brief, audit=audit)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ClassesApi->classes_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glob** | **str**| Filters the result according to a glob pattern applied on the short name of the requested object.  | [optional] 
 **deep** | **bool**| When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  | [optional] 
 **brief** | **bool**| When brief is specified, the response will only contain essential data and strip long descriptions.  | [optional] 
 **audit** | **bool**| When audit is specified, the response will also contain the audit trail field.  | [optional] 

### Return type

[**list[ClassNomenclature]**](ClassNomenclature.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

