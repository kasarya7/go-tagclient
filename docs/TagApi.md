# \TagApi

All URIs are relative to *https://localhost/apis/tag-api/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TagsDelete**](TagApi.md#TagsDelete) | **Delete** /tags/ | 批量删除标签
[**TagsGet**](TagApi.md#TagsGet) | **Get** /tags/ | 获取标签列表
[**TagsPost**](TagApi.md#TagsPost) | **Post** /tags/ | 创建标签
[**TagsTagIdGet**](TagApi.md#TagsTagIdGet) | **Get** /tags/{tag_id} | 获取单个标签


# **TagsDelete**
> TagsDelete(ctx, tags)
批量删除标签

# 删除标签 ## Errors + ## DELETETAG_ERR       删除标签失败       HTTP Status Code: 400   #### eg:   ```   {    \"code\":\"DELETETAG_ERR\",         \"message\":\"Delete Tag Fail.\"    \"data\":[            {\"tag\":\"tagtest1\",             \"id\":\"tagid1\",             \"fail_reason\"：\"Tag Not Found.\"},            {\"tag\":\"tagtest2\",             \"id\":\"tagid2\",             \"fail_reason\"：\"This label is being manipulated by another user.\"},            {\"tag\":\"tagtest3\",             \"id\":\"tagid3\",             \"fail_reason\"：\"Delete Tag Fail.\"},           ]   }   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tags** | [**ModelTagDeleteReq**](ModelTagDeleteReq.md)| 一组标签信息 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagsGet**
> interface{} TagsGet(ctx, optional)
获取标签列表

获取标签列表 (根据条件过滤标签)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TagApiTagsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagApiTagsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **optional.String**| 标签所属部门id | 
 **projectId** | **optional.String**| 标签所属项目id | 
 **user** | **optional.String**| 创建标签用户id | 
 **color** | **optional.String**| 标签颜色 | 
 **resNum** | **optional.Bool**| 标签绑定资源数 | [default to false]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagsPost**
> interface{} TagsPost(ctx, tags)
创建标签

# 创建一个新的标签 ## Errors + ## TAGALREADYEXISTS_ERR       要创建的标签已经存在       HTTP Status Code: 400   #### eg:   ```   {    \"code\":\"TAGALREADYEXISTS_ERR\",         \"message\":\"Conflict TagTest already exists.\"   }   ``` + ### MORETHANPROJECTTAGMAXNUM_ERR       一个项目下创建标签最多200个，超出最大限制       HTTP Status Code: 400    #### eg:   ```   {    \"code\":\"MORETHANPROJECTTAGMAXNUM_ERR\",         \"data\":{},    \"message\":\"A maximum of 200 labels can be created for a project.\"   }   ``` + ### PROJECTANDDOMAINISNOTCORRESPONDENCE_ERR       参数错误，项目不属于该部门       HTTP Status Code: 400   #### eg:   ```   {    \"code\":\"PROJECTANDDOMAINISNOTCORRESPONDENCE_ERR\",         \"data\":{},    \"message\":\"The project passed in does not belong to the domain.\"   }   ``` + ## PARAM_ERR       参数错误，请求参数不符合要求       HTTP Status Code: 422   #### eg:   ```   {    \"code\":\"PARAM_ERR\",         \"data\":{},    \"message\":\"error message\"   }   ``` + ## CREATETAG_ERR       创建标签失败       HTTP Status Code: 400   #### eg:   ```   {    \"code\":\"CREATETAG_ERR\",         \"data\":{},    \"message\":\"Create Tag Fail.\"   }   ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tags** | [**ModelTagCreateReq**](ModelTagCreateReq.md)| 一个新的标签信息 | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagsTagIdGet**
> interface{} TagsTagIdGet(ctx, tagId, optional)
获取单个标签

获取单个标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tagId** | **string**| 标签id | 
 **optional** | ***TagApiTagsTagIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TagApiTagsTagIdGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resNum** | **optional.Bool**| 标签绑定资源数 | [default to false]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

