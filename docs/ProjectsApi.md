# \ProjectsApi

All URIs are relative to *https://localhost/apis/tag-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsGet**](ProjectsApi.md#ProjectsGet) | **Get** /projects/ | 获取项目信息


# **ProjectsGet**
> interface{} ProjectsGet(ctx, optional)
获取项目信息

获取项目信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectsApiProjectsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsApiProjectsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **optional.String**| 部门id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

