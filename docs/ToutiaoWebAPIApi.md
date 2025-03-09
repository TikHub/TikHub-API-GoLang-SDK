# \ToutiaoWebAPIApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetArticleInfoApiV1ToutiaoWebGetArticleInfoGet**](ToutiaoWebAPIApi.md#GetArticleInfoApiV1ToutiaoWebGetArticleInfoGet) | **Get** /api/v1/toutiao/web/get_article_info | 获取指定文章的信息/Get information of specified article
[**GetArticleInfoApiV1ToutiaoWebGetArticleInfoGet_0**](ToutiaoWebAPIApi.md#GetArticleInfoApiV1ToutiaoWebGetArticleInfoGet_0) | **Get** /api/v1/toutiao/web/get_article_info | 获取指定文章的信息/Get information of specified article
[**GetVideoInfoApiV1ToutiaoWebGetVideoInfoGet**](ToutiaoWebAPIApi.md#GetVideoInfoApiV1ToutiaoWebGetVideoInfoGet) | **Get** /api/v1/toutiao/web/get_video_info | 获取指定视频的信息/Get information of specified video
[**GetVideoInfoApiV1ToutiaoWebGetVideoInfoGet_0**](ToutiaoWebAPIApi.md#GetVideoInfoApiV1ToutiaoWebGetVideoInfoGet_0) | **Get** /api/v1/toutiao/web/get_video_info | 获取指定视频的信息/Get information of specified video



## GetArticleInfoApiV1ToutiaoWebGetArticleInfoGet

> ResponseModel GetArticleInfoApiV1ToutiaoWebGetArticleInfoGet(ctx, awemeId)

获取指定文章的信息/Get information of specified article

# [中文] ### 用途: - 获取指定文章的信息 ### 参数: - aweme_id: 作品ID，可以从链接中获取     - 例如: https://www.toutiao.com/article/7450114952884503059/ ### 返回: - 作品信息  # [English] ### Purpose: - Get information of specified post ### Parameters: - item_id: Post ID, can be obtained from the link     - For example: https://www.toutiao.com/article/7450114952884503059/ ### Return: - Post information  # [示例/Example] aweme_id = \"7450114952884503059\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awemeId** | **string**| 作品ID/Post ID | 

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


## GetArticleInfoApiV1ToutiaoWebGetArticleInfoGet_0

> ResponseModel GetArticleInfoApiV1ToutiaoWebGetArticleInfoGet_0(ctx, awemeId)

获取指定文章的信息/Get information of specified article

# [中文] ### 用途: - 获取指定文章的信息 ### 参数: - aweme_id: 作品ID，可以从链接中获取     - 例如: https://www.toutiao.com/article/7450114952884503059/ ### 返回: - 作品信息  # [English] ### Purpose: - Get information of specified post ### Parameters: - item_id: Post ID, can be obtained from the link     - For example: https://www.toutiao.com/article/7450114952884503059/ ### Return: - Post information  # [示例/Example] aweme_id = \"7450114952884503059\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awemeId** | **string**| 作品ID/Post ID | 

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


## GetVideoInfoApiV1ToutiaoWebGetVideoInfoGet

> ResponseModel GetVideoInfoApiV1ToutiaoWebGetVideoInfoGet(ctx, awemeId)

获取指定视频的信息/Get information of specified video

# [中文] ### 用途: - 获取指定视频的信息 ### 参数: - aweme_id: 作品ID，可以从链接中获取     - 例如: https://www.toutiao.com/video/7431543350882206242/ ### 返回: - 作品信息  # [English] ### Purpose: - Get information of specified video ### Parameters: - item_id: Post ID, can be obtained from the link     - For example: https://www.toutiao.com/video/7431543350882206242/ ### Return: - Post information  # [示例/Example] aweme_id = \"7431543350882206242\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awemeId** | **string**| 作品ID/Post ID | 

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


## GetVideoInfoApiV1ToutiaoWebGetVideoInfoGet_0

> ResponseModel GetVideoInfoApiV1ToutiaoWebGetVideoInfoGet_0(ctx, awemeId)

获取指定视频的信息/Get information of specified video

# [中文] ### 用途: - 获取指定视频的信息 ### 参数: - aweme_id: 作品ID，可以从链接中获取     - 例如: https://www.toutiao.com/video/7431543350882206242/ ### 返回: - 作品信息  # [English] ### Purpose: - Get information of specified video ### Parameters: - item_id: Post ID, can be obtained from the link     - For example: https://www.toutiao.com/video/7431543350882206242/ ### Return: - Post information  # [示例/Example] aweme_id = \"7431543350882206242\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**awemeId** | **string**| 作品ID/Post ID | 

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

