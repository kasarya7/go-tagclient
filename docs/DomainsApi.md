# \DomainsApi

All URIs are relative to *https://localhost/apis/tag-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsGet**](DomainsApi.md#DomainsGet) | **Get** /domains/ | 获取部门信息


# **DomainsGet**
> interface{} DomainsGet(ctx, )
获取部门信息

获取部门信息

### Required Parameters
This endpoint does not need any parameter.

### Return type

**200**

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**400/401**

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | [optional] [default to null]
**Message** | **string** |  | [optional] [default to null]

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

