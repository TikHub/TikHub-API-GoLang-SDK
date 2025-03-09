# \HybridParsingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HybridParsingSingleVideoApiV1HybridVideoDataGet**](HybridParsingApi.md#HybridParsingSingleVideoApiV1HybridVideoDataGet) | **Get** /api/v1/hybrid/video_data | 混合解析单一视频接口/Hybrid parsing single video endpoint
[**HybridParsingSingleVideoApiV1HybridVideoDataGet_0**](HybridParsingApi.md#HybridParsingSingleVideoApiV1HybridVideoDataGet_0) | **Get** /api/v1/hybrid/video_data | 混合解析单一视频接口/Hybrid parsing single video endpoint
[**HybridParsingSingleVideoApiV1HybridVideoDataGet_1**](HybridParsingApi.md#HybridParsingSingleVideoApiV1HybridVideoDataGet_1) | **Get** /api/v1/hybrid/video_data | 混合解析单一视频接口/Hybrid parsing single video endpoint



## HybridParsingSingleVideoApiV1HybridVideoDataGet

> ResponseModel HybridParsingSingleVideoApiV1HybridVideoDataGet(ctx, url, optional)

混合解析单一视频接口/Hybrid parsing single video endpoint

# [中文] ### 用途: - 该接口用于解析抖音/TikTok单一视频的数据。 ### 参数: - `url`: 视频链接、分享链接、分享文本。 ### 返回: - `data`: 视频数据。  # [English] ### Purpose: - This endpoint is used to parse data of a single Douyin/TikTok video. ### Parameters: - `url`: Video link, share link, or share text. ### Returns: - `data`: Video data.  # [Example] url = \"https://v.douyin.com/L4FJNR3/\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**url** | **string**|  | 
 **optional** | ***HybridParsingSingleVideoApiV1HybridVideoDataGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HybridParsingSingleVideoApiV1HybridVideoDataGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **minimal** | **optional.Bool**| 是否返回最小数据/Whether to return minimal data | [default to false]
 **base64Url** | **optional.Bool**| 是否Base64编码提交URL/Base64 encoding URL | [default to false]

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HybridParsingSingleVideoApiV1HybridVideoDataGet_0

> ResponseModel HybridParsingSingleVideoApiV1HybridVideoDataGet_0(ctx, url, optional)

混合解析单一视频接口/Hybrid parsing single video endpoint

# [中文] ### 用途: - 该接口用于解析抖音/TikTok单一视频的数据。 ### 参数: - `url`: 视频链接、分享链接、分享文本。 ### 返回: - `data`: 视频数据。  # [English] ### Purpose: - This endpoint is used to parse data of a single Douyin/TikTok video. ### Parameters: - `url`: Video link, share link, or share text. ### Returns: - `data`: Video data.  # [Example] url = \"https://v.douyin.com/L4FJNR3/\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**url** | **string**|  | 
 **optional** | ***HybridParsingSingleVideoApiV1HybridVideoDataGet_1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HybridParsingSingleVideoApiV1HybridVideoDataGet_1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **minimal** | **optional.Bool**| 是否返回最小数据/Whether to return minimal data | [default to false]
 **base64Url** | **optional.Bool**| 是否Base64编码提交URL/Base64 encoding URL | [default to false]

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HybridParsingSingleVideoApiV1HybridVideoDataGet_1

> ResponseModel HybridParsingSingleVideoApiV1HybridVideoDataGet_1(ctx, url, optional)

混合解析单一视频接口/Hybrid parsing single video endpoint

# [中文] ### 用途: - 该接口用于解析抖音/TikTok单一视频的数据。 ### 参数: - `url`: 视频链接、分享链接、分享文本。 ### 返回: - `data`: 视频数据。  # [English] ### Purpose: - This endpoint is used to parse data of a single Douyin/TikTok video. ### Parameters: - `url`: Video link, share link, or share text. ### Returns: - `data`: Video data.  # [Example] url = \"https://v.douyin.com/L4FJNR3/\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**url** | **string**|  | 
 **optional** | ***HybridParsingSingleVideoApiV1HybridVideoDataGet_2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HybridParsingSingleVideoApiV1HybridVideoDataGet_2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **minimal** | **optional.Bool**| 是否返回最小数据/Whether to return minimal data | [default to false]
 **base64Url** | **optional.Bool**| 是否Base64编码提交URL/Base64 encoding URL | [default to false]

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

