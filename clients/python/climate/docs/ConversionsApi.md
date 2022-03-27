# swagger_client.ConversionsApi

All URIs are relative to *https://api.example.com/v1/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**conversion_from_unit_to_unit_get**](ConversionsApi.md#conversion_from_unit_to_unit_get) | **GET** /conversion/{fromUnit}/{toUnit} | Get the conversion specification from unit to unit.
[**conversion_get**](ConversionsApi.md#conversion_get) | **GET** /conversion | Get the conversion specification from unit to unit.
[**conversions_from_unit_to_unit_delete**](ConversionsApi.md#conversions_from_unit_to_unit_delete) | **DELETE** /conversions/{fromUnit}/{toUnit} | Delete an existing conversion specification from unit to unit.
[**conversions_from_unit_to_unit_post**](ConversionsApi.md#conversions_from_unit_to_unit_post) | **POST** /conversions/{fromUnit}/{toUnit} | Creates add new conversion specification from unit to unit.
[**conversions_from_unit_to_unit_put**](ConversionsApi.md#conversions_from_unit_to_unit_put) | **PUT** /conversions/{fromUnit}/{toUnit} | Update an existing conversion specification from unit to unit.
[**conversions_get**](ConversionsApi.md#conversions_get) | **GET** /conversions | List available unit conversions, with some query filters


# **conversion_from_unit_to_unit_get**
> ConversionSpec conversion_from_unit_to_unit_get(from_unit, to_unit, deep=deep, brief=brief, audit=audit)

Get the conversion specification from unit to unit.

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.ConversionsApi()
from_unit = 'from_unit_example' # str | Original unit to be converted. 
to_unit = 'to_unit_example' # str | Original unit to be converted. 
deep = true # bool | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  (optional)
brief = true # bool | When brief is specified, the response will only contain essential data and strip long descriptions.  (optional)
audit = true # bool | When audit is specified, the response will also contain the audit trail field.  (optional)

try:
    # Get the conversion specification from unit to unit.
    api_response = api_instance.conversion_from_unit_to_unit_get(from_unit, to_unit, deep=deep, brief=brief, audit=audit)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ConversionsApi->conversion_from_unit_to_unit_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from_unit** | **str**| Original unit to be converted.  | 
 **to_unit** | **str**| Original unit to be converted.  | 
 **deep** | **bool**| When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  | [optional] 
 **brief** | **bool**| When brief is specified, the response will only contain essential data and strip long descriptions.  | [optional] 
 **audit** | **bool**| When audit is specified, the response will also contain the audit trail field.  | [optional] 

### Return type

[**ConversionSpec**](ConversionSpec.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **conversion_get**
> Conversion conversion_get(from_unit=from_unit, to_unit=to_unit, deep=deep, brief=brief, audit=audit)

Get the conversion specification from unit to unit.

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.ConversionsApi()
from_unit = 'from_unit_example' # str | Original unit to be converted.  (optional)
to_unit = 'to_unit_example' # str | Target unit result of the conversion.  (optional)
deep = true # bool | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  (optional)
brief = true # bool | When brief is specified, the response will only contain essential data and strip long descriptions.  (optional)
audit = true # bool | When audit is specified, the response will also contain the audit trail field.  (optional)

try:
    # Get the conversion specification from unit to unit.
    api_response = api_instance.conversion_get(from_unit=from_unit, to_unit=to_unit, deep=deep, brief=brief, audit=audit)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ConversionsApi->conversion_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from_unit** | **str**| Original unit to be converted.  | [optional] 
 **to_unit** | **str**| Target unit result of the conversion.  | [optional] 
 **deep** | **bool**| When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  | [optional] 
 **brief** | **bool**| When brief is specified, the response will only contain essential data and strip long descriptions.  | [optional] 
 **audit** | **bool**| When audit is specified, the response will also contain the audit trail field.  | [optional] 

### Return type

[**Conversion**](Conversion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **conversions_from_unit_to_unit_delete**
> conversions_from_unit_to_unit_delete(from_unit, to_unit)

Delete an existing conversion specification from unit to unit.

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
api_instance = swagger_client.ConversionsApi(swagger_client.ApiClient(configuration))
from_unit = 'from_unit_example' # str | Original unit to be converted. 
to_unit = 'to_unit_example' # str | Original unit to be converted. 

try:
    # Delete an existing conversion specification from unit to unit.
    api_instance.conversions_from_unit_to_unit_delete(from_unit, to_unit)
except ApiException as e:
    print("Exception when calling ConversionsApi->conversions_from_unit_to_unit_delete: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from_unit** | **str**| Original unit to be converted.  | 
 **to_unit** | **str**| Original unit to be converted.  | 

### Return type

void (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **conversions_from_unit_to_unit_post**
> conversions_from_unit_to_unit_post(from_unit, to_unit, conversion_spec)

Creates add new conversion specification from unit to unit.

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
api_instance = swagger_client.ConversionsApi(swagger_client.ApiClient(configuration))
from_unit = 'from_unit_example' # str | Original unit to be converted. 
to_unit = 'to_unit_example' # str | Original unit to be converted. 
conversion_spec = swagger_client.ConversionSpec() # ConversionSpec | Unit conversion specification. 

try:
    # Creates add new conversion specification from unit to unit.
    api_instance.conversions_from_unit_to_unit_post(from_unit, to_unit, conversion_spec)
except ApiException as e:
    print("Exception when calling ConversionsApi->conversions_from_unit_to_unit_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from_unit** | **str**| Original unit to be converted.  | 
 **to_unit** | **str**| Original unit to be converted.  | 
 **conversion_spec** | [**ConversionSpec**](ConversionSpec.md)| Unit conversion specification.  | 

### Return type

void (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **conversions_from_unit_to_unit_put**
> conversions_from_unit_to_unit_put(from_unit, to_unit, conversion_spec)

Update an existing conversion specification from unit to unit.

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
api_instance = swagger_client.ConversionsApi(swagger_client.ApiClient(configuration))
from_unit = 'from_unit_example' # str | Original unit to be converted. 
to_unit = 'to_unit_example' # str | Original unit to be converted. 
conversion_spec = swagger_client.ConversionSpec() # ConversionSpec | Unit conversion specification. 

try:
    # Update an existing conversion specification from unit to unit.
    api_instance.conversions_from_unit_to_unit_put(from_unit, to_unit, conversion_spec)
except ApiException as e:
    print("Exception when calling ConversionsApi->conversions_from_unit_to_unit_put: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from_unit** | **str**| Original unit to be converted.  | 
 **to_unit** | **str**| Original unit to be converted.  | 
 **conversion_spec** | [**ConversionSpec**](ConversionSpec.md)| Unit conversion specification.  | 

### Return type

void (empty response body)

### Authorization

[apiKey](../README.md#apiKey), [bearerToken](../README.md#bearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **conversions_get**
> list[Conversion] conversions_get(from_unit=from_unit, to_unit=to_unit)

List available unit conversions, with some query filters

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.ConversionsApi()
from_unit = 'from_unit_example' # str | Original unit to be converted.  (optional)
to_unit = 'to_unit_example' # str | Target unit result of the conversion.  (optional)

try:
    # List available unit conversions, with some query filters
    api_response = api_instance.conversions_get(from_unit=from_unit, to_unit=to_unit)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling ConversionsApi->conversions_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from_unit** | **str**| Original unit to be converted.  | [optional] 
 **to_unit** | **str**| Target unit result of the conversion.  | [optional] 

### Return type

[**list[Conversion]**](Conversion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

