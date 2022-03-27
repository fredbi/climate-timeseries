# swagger_client.VersioningApi

All URIs are relative to *https://api.example.com/v1/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**series_versions_versioned_series_id_delete**](VersioningApi.md#series_versions_versioned_series_id_delete) | **DELETE** /series/versions/{versionedSeriesId} | Deletes a version of a time series
[**series_versions_versioned_series_id_get**](VersioningApi.md#series_versions_versioned_series_id_get) | **GET** /series/versions/{versionedSeriesId} | Get a version of a time series
[**series_versions_versioned_series_id_post**](VersioningApi.md#series_versions_versioned_series_id_post) | **POST** /series/versions/{versionedSeriesId} | creates a new version of a time series
[**series_versions_versioned_series_id_put**](VersioningApi.md#series_versions_versioned_series_id_put) | **PUT** /series/versions/{versionedSeriesId} | Updates the metadata of version of a time series
[**series_versions_versioned_series_id_values_delete**](VersioningApi.md#series_versions_versioned_series_id_values_delete) | **DELETE** /series/versions/{versionedSeriesId}/values | Deletes the values attached to a version of a time series
[**series_versions_versioned_series_id_values_get**](VersioningApi.md#series_versions_versioned_series_id_values_get) | **GET** /series/versions/{versionedSeriesId}/values | Get the values of version of a time series
[**series_versions_versioned_series_id_values_post**](VersioningApi.md#series_versions_versioned_series_id_values_post) | **POST** /series/versions/{versionedSeriesId}/values | Uploads values to the version of a time series
[**series_versions_versioned_series_id_values_put**](VersioningApi.md#series_versions_versioned_series_id_values_put) | **PUT** /series/versions/{versionedSeriesId}/values | Replaces the values of version of a time series


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
api_instance = swagger_client.VersioningApi()
versioned_series_id = 'versioned_series_id_example' # str | The UUID of a version of the time series.

try:
    # Deletes a version of a time series
    api_instance.series_versions_versioned_series_id_delete(versioned_series_id)
except ApiException as e:
    print("Exception when calling VersioningApi->series_versions_versioned_series_id_delete: %s\n" % e)
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
api_instance = swagger_client.VersioningApi()
versioned_series_id = 'versioned_series_id_example' # str | The UUID of a version of the time series.
deep = true # bool | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  (optional)
brief = true # bool | When brief is specified, the response will only contain essential data and strip long descriptions.  (optional)
audit = true # bool | When audit is specified, the response will also contain the audit trail field.  (optional)

try:
    # Get a version of a time series
    api_response = api_instance.series_versions_versioned_series_id_get(versioned_series_id, deep=deep, brief=brief, audit=audit)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling VersioningApi->series_versions_versioned_series_id_get: %s\n" % e)
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
api_instance = swagger_client.VersioningApi()
versioned_series_id = 'versioned_series_id_example' # str | The UUID of a version of the time series.
version_series = swagger_client.VersionedSeries() # VersionedSeries | The description of a version of a time series. 

try:
    # creates a new version of a time series
    api_instance.series_versions_versioned_series_id_post(versioned_series_id, version_series)
except ApiException as e:
    print("Exception when calling VersioningApi->series_versions_versioned_series_id_post: %s\n" % e)
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
api_instance = swagger_client.VersioningApi()
versioned_series_id = 'versioned_series_id_example' # str | The UUID of a version of the time series.
version_series = swagger_client.VersionedSeries() # VersionedSeries | The description of a version of a time series. 

try:
    # Updates the metadata of version of a time series
    api_instance.series_versions_versioned_series_id_put(versioned_series_id, version_series)
except ApiException as e:
    print("Exception when calling VersioningApi->series_versions_versioned_series_id_put: %s\n" % e)
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
api_instance = swagger_client.VersioningApi()
versioned_series_id = 'versioned_series_id_example' # str | The UUID of a version of the time series.

try:
    # Deletes the values attached to a version of a time series
    api_instance.series_versions_versioned_series_id_values_delete(versioned_series_id)
except ApiException as e:
    print("Exception when calling VersioningApi->series_versions_versioned_series_id_values_delete: %s\n" % e)
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
api_instance = swagger_client.VersioningApi()
versioned_series_id = 'versioned_series_id_example' # str | The UUID of a version of the time series.
_from = '2013-10-20T19:20:30+01:00' # datetime | Truncates values starting from this date (boundary is included).  (optional)
to = '2013-10-20T19:20:30+01:00' # datetime | Truncates values up to this date (boundary is included).  (optional)
convert = 'convert_example' # str | Converts values in some other compatible measurement unit, possibly specified as a composite measurement unit.  (optional)

try:
    # Get the values of version of a time series
    api_response = api_instance.series_versions_versioned_series_id_values_get(versioned_series_id, _from=_from, to=to, convert=convert)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling VersioningApi->series_versions_versioned_series_id_values_get: %s\n" % e)
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
api_instance = swagger_client.VersioningApi()
versioned_series_id = 'versioned_series_id_example' # str | The UUID of a version of the time series.
values = swagger_client.TsValues() # TsValues | The values of a time series 

try:
    # Uploads values to the version of a time series
    api_instance.series_versions_versioned_series_id_values_post(versioned_series_id, values)
except ApiException as e:
    print("Exception when calling VersioningApi->series_versions_versioned_series_id_values_post: %s\n" % e)
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
api_instance = swagger_client.VersioningApi()
versioned_series_id = 'versioned_series_id_example' # str | The UUID of a version of the time series.
values = swagger_client.TsValues() # TsValues | The values of a time series 

try:
    # Replaces the values of version of a time series
    api_instance.series_versions_versioned_series_id_values_put(versioned_series_id, values)
except ApiException as e:
    print("Exception when calling VersioningApi->series_versions_versioned_series_id_values_put: %s\n" % e)
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

