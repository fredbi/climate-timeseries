# swagger_client.SemverApi

All URIs are relative to *https://api.example.com/v1/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**series_series_id_semver_get**](SemverApi.md#series_series_id_semver_get) | **GET** /series/{seriesId}/semver | Get all semver tags associated to a series
[**series_series_id_semver_semver_get**](SemverApi.md#series_series_id_semver_semver_get) | **GET** /series/{seriesId}/semver/{semver} | Get a version of a time series with a semver tag
[**series_series_id_semver_semver_values_get**](SemverApi.md#series_series_id_semver_semver_values_get) | **GET** /series/{seriesId}/semver/{semver}/values | Get the values of version of a time series with a semver tag


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
api_instance = swagger_client.SemverApi()
series_id = 789 # int | The unique ID of a time series.

try:
    # Get all semver tags associated to a series
    api_response = api_instance.series_series_id_semver_get(series_id)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling SemverApi->series_series_id_semver_get: %s\n" % e)
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
api_instance = swagger_client.SemverApi()
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
    print("Exception when calling SemverApi->series_series_id_semver_semver_get: %s\n" % e)
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
api_instance = swagger_client.SemverApi()
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
    print("Exception when calling SemverApi->series_series_id_semver_semver_values_get: %s\n" % e)
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

