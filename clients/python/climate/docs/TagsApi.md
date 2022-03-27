# swagger_client.TagsApi

All URIs are relative to *https://api.example.com/v1/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**search_tags_get**](TagsApi.md#search_tags_get) | **GET** /search/tags | List all known tags
[**search_tags_tag_get**](TagsApi.md#search_tags_tag_get) | **GET** /search/tags/{tag} | Search all entities with some tag


# **search_tags_get**
> list[str] search_tags_get()

List all known tags

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.TagsApi()

try:
    # List all known tags
    api_response = api_instance.search_tags_get()
    pprint(api_response)
except ApiException as e:
    print("Exception when calling TagsApi->search_tags_get: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

**list[str]**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **search_tags_tag_get**
> Entities search_tags_tag_get(tag, status=status, deep=deep, brief=brief, audit=audit, owner_id=owner_id, email=email)

Search all entities with some tag

This endpoint returns mulitple entities, such as classes or series. 

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.TagsApi()
tag = 'tag_example' # str | Filter the result by search for a given keyword, whenever tag search is applicable.  This parameter has no effect on objects where tag search is not applicable. 
status = 'status_example' # str | Filter the result according to a given status.  Only series in the \"PUBLISHED\" status are returned to the public.  Series owner may consult their series in any status.  (optional)
deep = true # bool | When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  (optional)
brief = true # bool | When brief is specified, the response will only contain essential data and strip long descriptions.  (optional)
audit = true # bool | When audit is specified, the response will also contain the audit trail field.  (optional)
owner_id = 'owner_id_example' # str | The ID of the contributor.  (optional)
email = 'email_example' # str | Refers to some data owner by email.  (optional)

try:
    # Search all entities with some tag
    api_response = api_instance.search_tags_tag_get(tag, status=status, deep=deep, brief=brief, audit=audit, owner_id=owner_id, email=email)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling TagsApi->search_tags_tag_get: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **str**| Filter the result by search for a given keyword, whenever tag search is applicable.  This parameter has no effect on objects where tag search is not applicable.  | 
 **status** | **str**| Filter the result according to a given status.  Only series in the \&quot;PUBLISHED\&quot; status are returned to the public.  Series owner may consult their series in any status.  | [optional] 
 **deep** | **bool**| When deep is specified, the response will contain a deep representation of the object, rather than just a shallow description.  | [optional] 
 **brief** | **bool**| When brief is specified, the response will only contain essential data and strip long descriptions.  | [optional] 
 **audit** | **bool**| When audit is specified, the response will also contain the audit trail field.  | [optional] 
 **owner_id** | [**str**](.md)| The ID of the contributor.  | [optional] 
 **email** | **str**| Refers to some data owner by email.  | [optional] 

### Return type

[**Entities**](Entities.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

