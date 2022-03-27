# swagger_client.SeriesApi

All URIs are relative to *https://api.example.com/v1/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**series_get**](SeriesApi.md#series_get) | **GET** /series | Get all available series, with search filters
[**series_post**](SeriesApi.md#series_post) | **POST** /series | Creates a new time series
[**series_series_id_delete**](SeriesApi.md#series_series_id_delete) | **DELETE** /series/{seriesId} | Deletes a time series
[**series_series_id_get**](SeriesApi.md#series_series_id_get) | **GET** /series/{seriesId} | Get all versions of a time series.
[**series_series_id_latest_get**](SeriesApi.md#series_series_id_latest_get) | **GET** /series/{seriesId}/latest | Get the latest version of the description of a time series
[**series_series_id_latest_values_get**](SeriesApi.md#series_series_id_latest_values_get) | **GET** /series/{seriesId}/latest/values | Get the values from the latest version of a time series
[**series_series_id_post**](SeriesApi.md#series_series_id_post) | **POST** /series/{seriesId} | Creates a new version of a time series
[**series_series_id_put**](SeriesApi.md#series_series_id_put) | **PUT** /series/{seriesId} | Updates metadata about a time series
[**series_series_id_semver_get**](SeriesApi.md#series_series_id_semver_get) | **GET** /series/{seriesId}/semver | Get all semver tags associated to a series
[**series_series_id_semver_semver_get**](SeriesApi.md#series_series_id_semver_semver_get) | **GET** /series/{seriesId}/semver/{semver} | Get a version of a time series with a semver tag
[**series_series_id_semver_semver_values_get**](SeriesApi.md#series_series_id_semver_semver_values_get) | **GET** /series/{seriesId}/semver/{semver}/values | Get the values of version of a time series with a semver tag
[**series_versions_versioned_series_id_delete**](SeriesApi.md#series_versions_versioned_series_id_delete) | **DELETE** /series/versions/{versionedSeriesId} | Deletes a version of a time series
[**series_versions_versioned_series_id_get**](SeriesApi.md#series_versions_versioned_series_id_get) | **GET** /series/versions/{versionedSeriesId} | Get a version of a time series
[**series_versions_versioned_series_id_post**](SeriesApi.md#series_versions_versioned_series_id_post) | **POST** /series/versions/{versionedSeriesId} | creates a new version of a time series
[**series_versions_versioned_series_id_put**](SeriesApi.md#series_versions_versioned_series_id_put) | **PUT** /series/versions/{versionedSeriesId} | Updates the metadata of version of a time series
[**series_versions_versioned_series_id_values_delete**](SeriesApi.md#series_versions_versioned_series_id_values_delete) | **DELETE** /series/versions/{versionedSeriesId}/values | Deletes the values attached to a version of a time series
[**series_versions_versioned_series_id_values_get**](SeriesApi.md#series_versions_versioned_series_id_values_get) | **GET** /series/versions/{versionedSeriesId}/values | Get the values of version of a time series
[**series_versions_versioned_series_id_values_post**](SeriesApi.md#series_versions_versioned_series_id_values_post) | **POST** /series/versions/{versionedSeriesId}/values | Uploads values to the version of a time series
[**series_versions_versioned_series_id_values_put**](SeriesApi.md#series_versions_versioned_series_id_values_put) | **PUT** /series/versions/{versionedSeriesId}/values | Replaces the values of version of a time series


# **series_get**
> list[Series] series_get(glob=glob, tag=tag, tags=tags, theme=theme, zone=zone, status=status, owner_id=owner_id, datasource=datasource, email=email)

Get all available series, with search filters

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
glob = 'glob_example' # str | Filters the result according to a glob pattern applied on the short name of the requested object.  (optional)
tag = 'tag_example' # str | Filter the result by search for a given keyword, whenever tag search is applicable.  This parameter has no effect on objects where tag search is not applicable.  (optional)
tags = ['tags_example'] # list[str] | Filter the result by search for one tag in the given list of keywords, whenever tag search is applicable.  Up to 5 tags can be searched.  This parameter has no effect on objects where tag search is not applicable.  (optional)
theme = 'theme_example' # str | A theme path (url-encoded) to be searched for.  (optional)
zone = 'zone_example' # str | A zone short code (url-encoded) to be searched for.  (optional)
status = 'status_example' # str | Filter the result according to a given status.  Only series in the \"PUBLISHED\" status are returned to the public.  Series owner may consult their series in any status.  (optional)
owner_id = 'owner_id_example' # str | The ID of the contributor.  (optional)
datasource = 'datasource_example' # str | A datasource short code to be searched for.  (optional)
email = 'email_example' # str | Refers to some data owner by email.  (optional)

try:
    # Get all available series, with search filters
    api_response = api_instance.series_get(glob=glob, tag=tag, tags=tags, theme=theme, zone=zone, status=status, owner_id=owner_id, datasource=datasource, email=email)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SeriesApi->series_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **glob** | **str**| Filters the result according to a glob pattern applied on the short name of the requested object.  | [optional] 
 **tag** | **str**| Filter the result by search for a given keyword, whenever tag search is applicable.  This parameter has no effect on objects where tag search is not applicable.  | [optional] 
 **tags** | [**list[str]**](str.md)| Filter the result by search for one tag in the given list of keywords, whenever tag search is applicable.  Up to 5 tags can be searched.  This parameter has no effect on objects where tag search is not applicable.  | [optional] 
 **theme** | **str**| A theme path (url-encoded) to be searched for.  | [optional] 
 **zone** | **str**| A zone short code (url-encoded) to be searched for.  | [optional] 
 **status** | **str**| Filter the result according to a given status.  Only series in the \&quot;PUBLISHED\&quot; status are returned to the public.  Series owner may consult their series in any status.  | [optional] 
 **owner_id** | [**str**](.md)| The ID of the contributor.  | [optional] 
 **datasource** | **str**| A datasource short code to be searched for.  | [optional] 
 **email** | **str**| Refers to some data owner by email.  | [optional] 

### Return type

[**list[Series]**](Series.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_post**
> series_post(series)

Creates a new time series

This creates a new time time series.  If some values are specified, the initial version for the new series is created with these values. 

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
series = swagger_client.Series() # Series | The description of a time series. 

try:
    # Creates a new time series
    api_instance.series_post(series)
except ApiException as e:
    print("Exception when calling SeriesApi->series_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **series** | [**Series**](Series.md)| The description of a time series.  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_series_id_delete**
> series_series_id_delete(series_id)

Deletes a time series

All versions and values associated to this series are deleted.  Only the original owner of a series may delete it. 

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
series_id = 789 # int | The unique ID of a time series.

try:
    # Deletes a time series
    api_instance.series_series_id_delete(series_id)
except ApiException as e:
    print("Exception when calling SeriesApi->series_series_id_delete: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **series_id** | **int**| The unique ID of a time series. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_series_id_get**
> list[VersionedSeries] series_series_id_get(series_id, glob=glob, deep=deep, brief=brief, audit=audit, status=status)

Get all versions of a time series.

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
series_id = 789 # int | The unique ID of a time series.
glob = 'glob_example' # str | Filters the result according to a glob pattern applied on the short name of the requested object.  (optional)
deep = true # bool | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  (optional)
brief = true # bool | When brief is specified, the response will only contain essential data and strip long descriptions.  (optional)
audit = true # bool | When audit is specified, the response will also contain the audit trail field.  (optional)
status = 'status_example' # str | Filter the result according to a given status.  Only series in the \"PUBLISHED\" status are returned to the public.  Series owner may consult their series in any status.  (optional)

try:
    # Get all versions of a time series.
    api_response = api_instance.series_series_id_get(series_id, glob=glob, deep=deep, brief=brief, audit=audit, status=status)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SeriesApi->series_series_id_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **series_id** | **int**| The unique ID of a time series. | 
 **glob** | **str**| Filters the result according to a glob pattern applied on the short name of the requested object.  | [optional] 
 **deep** | **bool**| When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  | [optional] 
 **brief** | **bool**| When brief is specified, the response will only contain essential data and strip long descriptions.  | [optional] 
 **audit** | **bool**| When audit is specified, the response will also contain the audit trail field.  | [optional] 
 **status** | **str**| Filter the result according to a given status.  Only series in the \&quot;PUBLISHED\&quot; status are returned to the public.  Series owner may consult their series in any status.  | [optional] 

### Return type

[**list[VersionedSeries]**](VersionedSeries.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_series_id_latest_get**
> VersionedSeries series_series_id_latest_get(series_id, deep=deep, brief=brief, audit=audit, status=status)

Get the latest version of the description of a time series

The latest version is determined according to semantic versioning rules (e.g. v1.2.3 < v1.2.4).  Unless requested by the query parameters, the time series values are not returned by default. 

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
series_id = 789 # int | The unique ID of a time series.
deep = true # bool | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  (optional)
brief = true # bool | When brief is specified, the response will only contain essential data and strip long descriptions.  (optional)
audit = true # bool | When audit is specified, the response will also contain the audit trail field.  (optional)
status = 'status_example' # str | Filter the result according to a given status.  Only series in the \"PUBLISHED\" status are returned to the public.  Series owner may consult their series in any status.  (optional)

try:
    # Get the latest version of the description of a time series
    api_response = api_instance.series_series_id_latest_get(series_id, deep=deep, brief=brief, audit=audit, status=status)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SeriesApi->series_series_id_latest_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **series_id** | **int**| The unique ID of a time series. | 
 **deep** | **bool**| When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  | [optional] 
 **brief** | **bool**| When brief is specified, the response will only contain essential data and strip long descriptions.  | [optional] 
 **audit** | **bool**| When audit is specified, the response will also contain the audit trail field.  | [optional] 
 **status** | **str**| Filter the result according to a given status.  Only series in the \&quot;PUBLISHED\&quot; status are returned to the public.  Series owner may consult their series in any status.  | [optional] 

### Return type

[**VersionedSeries**](VersionedSeries.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_series_id_latest_values_get**
> TsValues series_series_id_latest_values_get(series_id, _from=_from, to=to, convert=convert)

Get the values from the latest version of a time series

The latest version is determined according to semantic versioning rules (e.g. v1.2.3 < v1.2.4).  Values are returned in chronological order.  If the request negotiates a response MIME type with text/csv (with the Accept header), this endpoint produces a CSV file. 

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
series_id = 789 # int | The unique ID of a time series.
_from = '2013-10-20T19:20:30+01:00' # datetime | Truncates values starting from this date (boundary is included).  (optional)
to = '2013-10-20T19:20:30+01:00' # datetime | Truncates values up to this date (boundary is included).  (optional)
convert = 'convert_example' # str | Converts values in some other compatible measurement unit, possibly specified as a composite measurement unit.  (optional)

try:
    # Get the values from the latest version of a time series
    api_response = api_instance.series_series_id_latest_values_get(series_id, _from=_from, to=to, convert=convert)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SeriesApi->series_series_id_latest_values_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **series_id** | **int**| The unique ID of a time series. | 
 **_from** | **datetime**| Truncates values starting from this date (boundary is included).  | [optional] 
 **to** | **datetime**| Truncates values up to this date (boundary is included).  | [optional] 
 **convert** | **str**| Converts values in some other compatible measurement unit, possibly specified as a composite measurement unit.  | [optional] 

### Return type

[**TsValues**](TsValues.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_series_id_post**
> series_series_id_post(series_id, series)

Creates a new version of a time series

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
series_id = 789 # int | The unique ID of a time series.
series = swagger_client.Series() # Series | The description of a time series. 

try:
    # Creates a new version of a time series
    api_instance.series_series_id_post(series_id, series)
except ApiException as e:
    print("Exception when calling SeriesApi->series_series_id_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **series_id** | **int**| The unique ID of a time series. | 
 **series** | [**Series**](Series.md)| The description of a time series.  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_series_id_put**
> series_series_id_put(series_id, series)

Updates metadata about a time series

This action only updates metadata about the series (such as associated themes, tags, etc) and does not create a new version.  Only owners registered for a series may update the series.  Any time series values specified in the input are ignored. One must update a specific version to modify the values of a time series. 

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
series_id = 789 # int | The unique ID of a time series.
series = swagger_client.Series() # Series | The description of a time series. 

try:
    # Updates metadata about a time series
    api_instance.series_series_id_put(series_id, series)
except ApiException as e:
    print("Exception when calling SeriesApi->series_series_id_put: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **series_id** | **int**| The unique ID of a time series. | 
 **series** | [**Series**](Series.md)| The description of a time series.  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_series_id_semver_get**
> list[str] series_series_id_semver_get(series_id)

Get all semver tags associated to a series

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
series_id = 789 # int | The unique ID of a time series.

try:
    # Get all semver tags associated to a series
    api_response = api_instance.series_series_id_semver_get(series_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SeriesApi->series_series_id_semver_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **series_id** | **int**| The unique ID of a time series. | 

### Return type

**list[str]**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_series_id_semver_semver_get**
> VersionedSeries series_series_id_semver_semver_get(series_id, semver, deep=deep, brief=brief, audit=audit)

Get a version of a time series with a semver tag

Semver tags follow semantic versioning rules (e.g. v1.2.3, v1.2.3-rc1, v1.3.0 ...).  Unless requested by the query parameters, the time series values are not returned by default. 

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
series_id = 789 # int | The unique ID of a time series.
semver = 'semver_example' # str | The semver tag of a series' version.
deep = true # bool | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  (optional)
brief = true # bool | When brief is specified, the response will only contain essential data and strip long descriptions.  (optional)
audit = true # bool | When audit is specified, the response will also contain the audit trail field.  (optional)

try:
    # Get a version of a time series with a semver tag
    api_response = api_instance.series_series_id_semver_semver_get(series_id, semver, deep=deep, brief=brief, audit=audit)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SeriesApi->series_series_id_semver_semver_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **series_id** | **int**| The unique ID of a time series. | 
 **semver** | [**str**](.md)| The semver tag of a series&#39; version. | 
 **deep** | **bool**| When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  | [optional] 
 **brief** | **bool**| When brief is specified, the response will only contain essential data and strip long descriptions.  | [optional] 
 **audit** | **bool**| When audit is specified, the response will also contain the audit trail field.  | [optional] 

### Return type

[**VersionedSeries**](VersionedSeries.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_series_id_semver_semver_values_get**
> TsValues series_series_id_semver_semver_values_get(series_id, semver, _from=_from, to=to, convert=convert)

Get the values of version of a time series with a semver tag

 If the request negotiates a response MIME type with text/csv (with the Accept header), this endpoint produces a CSV file. 

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
series_id = 789 # int | The unique ID of a time series.
semver = 'semver_example' # str | The semver tag of a series' version.
_from = '2013-10-20T19:20:30+01:00' # datetime | Truncates values starting from this date (boundary is included).  (optional)
to = '2013-10-20T19:20:30+01:00' # datetime | Truncates values up to this date (boundary is included).  (optional)
convert = 'convert_example' # str | Converts values in some other compatible measurement unit, possibly specified as a composite measurement unit.  (optional)

try:
    # Get the values of version of a time series with a semver tag
    api_response = api_instance.series_series_id_semver_semver_values_get(series_id, semver, _from=_from, to=to, convert=convert)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SeriesApi->series_series_id_semver_semver_values_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **series_id** | **int**| The unique ID of a time series. | 
 **semver** | [**str**](.md)| The semver tag of a series&#39; version. | 
 **_from** | **datetime**| Truncates values starting from this date (boundary is included).  | [optional] 
 **to** | **datetime**| Truncates values up to this date (boundary is included).  | [optional] 
 **convert** | **str**| Converts values in some other compatible measurement unit, possibly specified as a composite measurement unit.  | [optional] 

### Return type

[**TsValues**](TsValues.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_versions_versioned_series_id_delete**
> series_versions_versioned_series_id_delete(versioned_series_id)

Deletes a version of a time series

Only the original owner of this version of the series may delete it.  All values associated to that version are deleted. 

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
versioned_series_id = 'versioned_series_id_example' # str | The UUID of a version of the time series.

try:
    # Deletes a version of a time series
    api_instance.series_versions_versioned_series_id_delete(versioned_series_id)
except ApiException as e:
    print("Exception when calling SeriesApi->series_versions_versioned_series_id_delete: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **versioned_series_id** | [**str**](.md)| The UUID of a version of the time series. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_versions_versioned_series_id_get**
> VersionedSeries series_versions_versioned_series_id_get(versioned_series_id, deep=deep, brief=brief, audit=audit)

Get a version of a time series

Unless requested by the query parameters, the time series values are not returned by default. 

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
versioned_series_id = 'versioned_series_id_example' # str | The UUID of a version of the time series.
deep = true # bool | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  (optional)
brief = true # bool | When brief is specified, the response will only contain essential data and strip long descriptions.  (optional)
audit = true # bool | When audit is specified, the response will also contain the audit trail field.  (optional)

try:
    # Get a version of a time series
    api_response = api_instance.series_versions_versioned_series_id_get(versioned_series_id, deep=deep, brief=brief, audit=audit)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SeriesApi->series_versions_versioned_series_id_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **versioned_series_id** | [**str**](.md)| The UUID of a version of the time series. | 
 **deep** | **bool**| When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  | [optional] 
 **brief** | **bool**| When brief is specified, the response will only contain essential data and strip long descriptions.  | [optional] 
 **audit** | **bool**| When audit is specified, the response will also contain the audit trail field.  | [optional] 

### Return type

[**VersionedSeries**](VersionedSeries.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_versions_versioned_series_id_post**
> series_versions_versioned_series_id_post(versioned_series_id, version_series)

creates a new version of a time series

A conflict is reported if the version was already attributed to a version of the same series. Use PUT to update existing values. 

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
versioned_series_id = 'versioned_series_id_example' # str | The UUID of a version of the time series.
version_series = swagger_client.VersionedSeries() # VersionedSeries | The description of a version of a time series. 

try:
    # creates a new version of a time series
    api_instance.series_versions_versioned_series_id_post(versioned_series_id, version_series)
except ApiException as e:
    print("Exception when calling SeriesApi->series_versions_versioned_series_id_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **versioned_series_id** | [**str**](.md)| The UUID of a version of the time series. | 
 **version_series** | [**VersionedSeries**](VersionedSeries.md)| The description of a version of a time series.  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_versions_versioned_series_id_put**
> series_versions_versioned_series_id_put(versioned_series_id, version_series)

Updates the metadata of version of a time series

This endpoint replaces the metadata for the requested version, without creating a new version.  Only the owners of the time series may modify values. 

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
versioned_series_id = 'versioned_series_id_example' # str | The UUID of a version of the time series.
version_series = swagger_client.VersionedSeries() # VersionedSeries | The description of a version of a time series. 

try:
    # Updates the metadata of version of a time series
    api_instance.series_versions_versioned_series_id_put(versioned_series_id, version_series)
except ApiException as e:
    print("Exception when calling SeriesApi->series_versions_versioned_series_id_put: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **versioned_series_id** | [**str**](.md)| The UUID of a version of the time series. | 
 **version_series** | [**VersionedSeries**](VersionedSeries.md)| The description of a version of a time series.  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_versions_versioned_series_id_values_delete**
> series_versions_versioned_series_id_values_delete(versioned_series_id)

Deletes the values attached to a version of a time series

Only the original owner of this version of the series may delete it.  The version is not deleted. Only the values associated to that version are deleted. 

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
versioned_series_id = 'versioned_series_id_example' # str | The UUID of a version of the time series.

try:
    # Deletes the values attached to a version of a time series
    api_instance.series_versions_versioned_series_id_values_delete(versioned_series_id)
except ApiException as e:
    print("Exception when calling SeriesApi->series_versions_versioned_series_id_values_delete: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **versioned_series_id** | [**str**](.md)| The UUID of a version of the time series. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_versions_versioned_series_id_values_get**
> TsValues series_versions_versioned_series_id_values_get(versioned_series_id, _from=_from, to=to, convert=convert)

Get the values of version of a time series

If the request negotiates a response MIME type with text/csv (with the Accept header), this endpoint produces a CSV file. 

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
versioned_series_id = 'versioned_series_id_example' # str | The UUID of a version of the time series.
_from = '2013-10-20T19:20:30+01:00' # datetime | Truncates values starting from this date (boundary is included).  (optional)
to = '2013-10-20T19:20:30+01:00' # datetime | Truncates values up to this date (boundary is included).  (optional)
convert = 'convert_example' # str | Converts values in some other compatible measurement unit, possibly specified as a composite measurement unit.  (optional)

try:
    # Get the values of version of a time series
    api_response = api_instance.series_versions_versioned_series_id_values_get(versioned_series_id, _from=_from, to=to, convert=convert)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SeriesApi->series_versions_versioned_series_id_values_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **versioned_series_id** | [**str**](.md)| The UUID of a version of the time series. | 
 **_from** | **datetime**| Truncates values starting from this date (boundary is included).  | [optional] 
 **to** | **datetime**| Truncates values up to this date (boundary is included).  | [optional] 
 **convert** | **str**| Converts values in some other compatible measurement unit, possibly specified as a composite measurement unit.  | [optional] 

### Return type

[**TsValues**](TsValues.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_versions_versioned_series_id_values_post**
> series_versions_versioned_series_id_values_post(versioned_series_id, values)

Uploads values to the version of a time series

This endpoint creates time series values for the latest version.  Only the owners of the time series may add values.  A conflict is reported if values were already attributed to this version. Use PUT to update existing values.  If the request negotiates a request MIME type with text/csv (with the Content-Type header), this endpoint consumes a CSV file. 

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
versioned_series_id = 'versioned_series_id_example' # str | The UUID of a version of the time series.
values = swagger_client.TsValues() # TsValues | The values of a time series 

try:
    # Uploads values to the version of a time series
    api_instance.series_versions_versioned_series_id_values_post(versioned_series_id, values)
except ApiException as e:
    print("Exception when calling SeriesApi->series_versions_versioned_series_id_values_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **versioned_series_id** | [**str**](.md)| The UUID of a version of the time series. | 
 **values** | [**TsValues**](TsValues.md)| The values of a time series  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **series_versions_versioned_series_id_values_put**
> series_versions_versioned_series_id_values_put(versioned_series_id, values)

Replaces the values of version of a time series

This endpoint replaces time series values for the requested version, without creating a new version.  Only the owners of the time series may modify values.  A conflict is reported if values were already attributed to this version. Use PUT to update existing values.  If the request negotiates a request MIME type with text/csv (with the Content-Type header), this endpoint consumes a CSV file. 

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.SeriesApi()
versioned_series_id = 'versioned_series_id_example' # str | The UUID of a version of the time series.
values = swagger_client.TsValues() # TsValues | The values of a time series 

try:
    # Replaces the values of version of a time series
    api_instance.series_versions_versioned_series_id_values_put(versioned_series_id, values)
except ApiException as e:
    print("Exception when calling SeriesApi->series_versions_versioned_series_id_values_put: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **versioned_series_id** | [**str**](.md)| The UUID of a version of the time series. | 
 **values** | [**TsValues**](TsValues.md)| The values of a time series  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/csv
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

