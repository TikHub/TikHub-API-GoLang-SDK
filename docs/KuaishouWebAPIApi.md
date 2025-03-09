# \KuaishouWebAPIApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchOneVideoByUrlApiV1KuaishouWebFetchOneVideoByUrlGet**](KuaishouWebAPIApi.md#FetchOneVideoByUrlApiV1KuaishouWebFetchOneVideoByUrlGet) | **Get** /api/v1/kuaishou/web/fetch_one_video_by_url | 根据链接获取单个作品数据/Fetch single video by URL
[**FetchOneVideoByUrlApiV1KuaishouWebFetchOneVideoByUrlGet_0**](KuaishouWebAPIApi.md#FetchOneVideoByUrlApiV1KuaishouWebFetchOneVideoByUrlGet_0) | **Get** /api/v1/kuaishou/web/fetch_one_video_by_url | 根据链接获取单个作品数据/Fetch single video by URL
[**FetchOneVideoV2ApiV1KuaishouWebFetchOneVideoV2Get**](KuaishouWebAPIApi.md#FetchOneVideoV2ApiV1KuaishouWebFetchOneVideoV2Get) | **Get** /api/v1/kuaishou/web/fetch_one_video_v2 | 快手单一视频查询接口V2/Kuaishou single video query API V2
[**FetchOneVideoV2ApiV1KuaishouWebFetchOneVideoV2Get_0**](KuaishouWebAPIApi.md#FetchOneVideoV2ApiV1KuaishouWebFetchOneVideoV2Get_0) | **Get** /api/v1/kuaishou/web/fetch_one_video_v2 | 快手单一视频查询接口V2/Kuaishou single video query API V2



## FetchOneVideoByUrlApiV1KuaishouWebFetchOneVideoByUrlGet

> ResponseModel FetchOneVideoByUrlApiV1KuaishouWebFetchOneVideoByUrlGet(ctx, url)

根据链接获取单个作品数据/Fetch single video by URL

# [中文] ### 用途: - 根据链接获取单个作品数据 ### 参数: - url: 作品链接 ### 返回: - 视频数据  # [English] ### Purpose: - Fetch single video by URL ### Parameters: - url: Photo URL ### Returns: - Video data  # [示例/Example] url = \"https://v.kuaishou.com/GKTpYm\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**url** | **string**|  | 

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


## FetchOneVideoByUrlApiV1KuaishouWebFetchOneVideoByUrlGet_0

> ResponseModel FetchOneVideoByUrlApiV1KuaishouWebFetchOneVideoByUrlGet_0(ctx, url)

根据链接获取单个作品数据/Fetch single video by URL

# [中文] ### 用途: - 根据链接获取单个作品数据 ### 参数: - url: 作品链接 ### 返回: - 视频数据  # [English] ### Purpose: - Fetch single video by URL ### Parameters: - url: Photo URL ### Returns: - Video data  # [示例/Example] url = \"https://v.kuaishou.com/GKTpYm\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**url** | **string**|  | 

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


## FetchOneVideoV2ApiV1KuaishouWebFetchOneVideoV2Get

> ResponseModel FetchOneVideoV2ApiV1KuaishouWebFetchOneVideoV2Get(ctx, photoId)

快手单一视频查询接口V2/Kuaishou single video query API V2

# [中文] ### 用途: - 快手单一视频查询接口V2 ### 参数: - photo_id: 作品ID，作品ID可以从作品链接中提取 ### 返回: - 视频数据  # [English] ### Purpose: - Kuaishou single video query API V2 ### Parameters: - photo_id: Photo ID, the photo ID can be extracted from the photo link ### Returns: - Video data  # [示例/Example] photo_id = \"3xtdqvdnqd3psuc\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**photoId** | **string**|  | 

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


## FetchOneVideoV2ApiV1KuaishouWebFetchOneVideoV2Get_0

> ResponseModel FetchOneVideoV2ApiV1KuaishouWebFetchOneVideoV2Get_0(ctx, photoId)

快手单一视频查询接口V2/Kuaishou single video query API V2

# [中文] ### 用途: - 快手单一视频查询接口V2 ### 参数: - photo_id: 作品ID，作品ID可以从作品链接中提取 ### 返回: - 视频数据  # [English] ### Purpose: - Kuaishou single video query API V2 ### Parameters: - photo_id: Photo ID, the photo ID can be extracted from the photo link ### Returns: - Video data  # [示例/Example] photo_id = \"3xtdqvdnqd3psuc\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**photoId** | **string**|  | 

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

