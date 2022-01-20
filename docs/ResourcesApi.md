# \ResourcesApi

All URIs are relative to *https://localhost/apis/tag-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResourcesActionPost**](ResourcesApi.md#ResourcesActionPost) | **Post** /resources/action/ | 绑定或者解绑标签
[**ResourcesPost**](ResourcesApi.md#ResourcesPost) | **Post** /resources/ | 根据标签、部门、项目信息等检索资源信息
[**ResourcesTagsPost**](ResourcesApi.md#ResourcesTagsPost) | **Post** /resources/tags/ | 根据资源检索资源绑定的标签信息


# **ResourcesActionPost**
> ResourcesActionPost(ctx, resource)
绑定或者解绑标签

# 绑定或者解绑标签 ## Errors + ## BOUNDORUNBOUNDTAG_ERR       绑定标签失败       HTTP Status Code: 400   #### eg:   ```   {    \"code\":\"BOUNDORUNBOUNDTAG_ERR\",         \"message\":\"Bound Or UnBound Tag Fail.\",    \"data\":[             {\"error_resource\": \"tagtest1\",              \"error_resource_type\":\"OS_Tag\",              \"error_tags\":[tagtest1,tagtest2],              \"error_reason\":\"Bound Or UnBound Fail.\"             },             {\"error_resource\": \"tagtest2\",              \"error_resource_type\":\"OS_Tag\",              \"error_tags\":[tagtest3,tagtest4],              \"error_reason\":\"The resource binding labels exceed 20.\"             },              {\"error_resource\": \"tagtest3\",              \"error_resource_type\":\"OS_Tag\",              \"error_tags\":[tagtest5,tagtest6],              \"error_reason\":\"This resource is being operated by another user and cannot be bound. Please try again.\"             }            ]   }   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resource** | [**ModelResourceBoundOrUnBoundTagReq**](ModelResourceBoundOrUnBoundTagReq.md)| 需要绑定的标签信息和资源信息 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResourcesPost**
> interface{} ResourcesPost(ctx, data)
根据标签、部门、项目信息等检索资源信息

根据标签、部门、项目信息等检索资源信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ModelResourceSearchReq**](ModelResourceSearchReq.md)| 标签信息 | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResourcesTagsPost**
> interface{} ResourcesTagsPost(ctx, data)
根据资源检索资源绑定的标签信息

根据资源检索资源绑定的标签信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ModelResourceGetBoundTagReq**](ModelResourceGetBoundTagReq.md)| 资源信息 | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

