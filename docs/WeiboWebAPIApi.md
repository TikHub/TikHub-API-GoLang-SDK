# \WeiboWebAPIApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPostDetailApiV1WeiboWebFetchPostDetailGet**](WeiboWebAPIApi.md#FetchPostDetailApiV1WeiboWebFetchPostDetailGet) | **Get** /api/v1/weibo/web/fetch_post_detail | 获取单个作品数据/Get single video data
[**FetchPostDetailApiV1WeiboWebFetchPostDetailGet_0**](WeiboWebAPIApi.md#FetchPostDetailApiV1WeiboWebFetchPostDetailGet_0) | **Get** /api/v1/weibo/web/fetch_post_detail | 获取单个作品数据/Get single video data
[**FetchSearchDataApiV1WeiboWebFetchSearchDataGet**](WeiboWebAPIApi.md#FetchSearchDataApiV1WeiboWebFetchSearchDataGet) | **Get** /api/v1/weibo/web/fetch_search_data | 获取搜索数据/Get search data
[**FetchSearchDataApiV1WeiboWebFetchSearchDataGet_0**](WeiboWebAPIApi.md#FetchSearchDataApiV1WeiboWebFetchSearchDataGet_0) | **Get** /api/v1/weibo/web/fetch_search_data | 获取搜索数据/Get search data
[**FetchShortVideoDataApiV1WeiboWebFetchShortVideoDataGet**](WeiboWebAPIApi.md#FetchShortVideoDataApiV1WeiboWebFetchShortVideoDataGet) | **Get** /api/v1/weibo/web/fetch_short_video_data | 获取短视频数据/Get short video data
[**FetchShortVideoDataApiV1WeiboWebFetchShortVideoDataGet_0**](WeiboWebAPIApi.md#FetchShortVideoDataApiV1WeiboWebFetchShortVideoDataGet_0) | **Get** /api/v1/weibo/web/fetch_short_video_data | 获取短视频数据/Get short video data
[**FetchTopicDetailApiV1WeiboWebFetchTopicDetailGet**](WeiboWebAPIApi.md#FetchTopicDetailApiV1WeiboWebFetchTopicDetailGet) | **Get** /api/v1/weibo/web/fetch_topic_detail | 获取话题详情/Get topic details
[**FetchTopicDetailApiV1WeiboWebFetchTopicDetailGet_0**](WeiboWebAPIApi.md#FetchTopicDetailApiV1WeiboWebFetchTopicDetailGet_0) | **Get** /api/v1/weibo/web/fetch_topic_detail | 获取话题详情/Get topic details
[**FetchTopicStatsApiV1WeiboWebFetchTopicStatsGet**](WeiboWebAPIApi.md#FetchTopicStatsApiV1WeiboWebFetchTopicStatsGet) | **Get** /api/v1/weibo/web/fetch_topic_stats | 获取话题统计数据/Get topic statistics
[**FetchTopicStatsApiV1WeiboWebFetchTopicStatsGet_0**](WeiboWebAPIApi.md#FetchTopicStatsApiV1WeiboWebFetchTopicStatsGet_0) | **Get** /api/v1/weibo/web/fetch_topic_stats | 获取话题统计数据/Get topic statistics
[**FetchUserInfoApiV1WeiboWebFetchUserInfoGet**](WeiboWebAPIApi.md#FetchUserInfoApiV1WeiboWebFetchUserInfoGet) | **Get** /api/v1/weibo/web/fetch_user_info | 获取用户信息/Get user information
[**FetchUserInfoApiV1WeiboWebFetchUserInfoGet_0**](WeiboWebAPIApi.md#FetchUserInfoApiV1WeiboWebFetchUserInfoGet_0) | **Get** /api/v1/weibo/web/fetch_user_info | 获取用户信息/Get user information
[**FetchUserInfoV2ApiV1WeiboWebFetchUserInfoV2Get**](WeiboWebAPIApi.md#FetchUserInfoV2ApiV1WeiboWebFetchUserInfoV2Get) | **Get** /api/v1/weibo/web/fetch_user_info_v2 | 获取用户信息V2/Get user information V2
[**FetchUserInfoV2ApiV1WeiboWebFetchUserInfoV2Get_0**](WeiboWebAPIApi.md#FetchUserInfoV2ApiV1WeiboWebFetchUserInfoV2Get_0) | **Get** /api/v1/weibo/web/fetch_user_info_v2 | 获取用户信息V2/Get user information V2
[**FetchUserPostsApiV1WeiboWebFetchUserPostsGet**](WeiboWebAPIApi.md#FetchUserPostsApiV1WeiboWebFetchUserPostsGet) | **Get** /api/v1/weibo/web/fetch_user_posts | 获取微博用户文章数据/Get Weibo user article data
[**FetchUserPostsApiV1WeiboWebFetchUserPostsGet_0**](WeiboWebAPIApi.md#FetchUserPostsApiV1WeiboWebFetchUserPostsGet_0) | **Get** /api/v1/weibo/web/fetch_user_posts | 获取微博用户文章数据/Get Weibo user article data



## FetchPostDetailApiV1WeiboWebFetchPostDetailGet

> ResponseModel FetchPostDetailApiV1WeiboWebFetchPostDetailGet(ctx, id)

获取单个作品数据/Get single video data

# [中文] ### 用途: - 获取单个作品数据 ### 参数: - id: 作品id，从分享链接中获取  - https://weibo.com/5819018196/5092682368025584  - 作品id为：5092682368025584 ### 返回: - 作品数据  # [English] ### Purpose: - Get single video data ### Parameters: - id: Post id, obtained from the sharing link     - https://weibo.com/5819018196/5092682368025584     - The post id is: 5092682368025584 ### Return: - Post data  # [示例/Example] id = \"5092682368025584\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| 作品id/Post id | 

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


## FetchPostDetailApiV1WeiboWebFetchPostDetailGet_0

> ResponseModel FetchPostDetailApiV1WeiboWebFetchPostDetailGet_0(ctx, id)

获取单个作品数据/Get single video data

# [中文] ### 用途: - 获取单个作品数据 ### 参数: - id: 作品id，从分享链接中获取  - https://weibo.com/5819018196/5092682368025584  - 作品id为：5092682368025584 ### 返回: - 作品数据  # [English] ### Purpose: - Get single video data ### Parameters: - id: Post id, obtained from the sharing link     - https://weibo.com/5819018196/5092682368025584     - The post id is: 5092682368025584 ### Return: - Post data  # [示例/Example] id = \"5092682368025584\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| 作品id/Post id | 

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


## FetchSearchDataApiV1WeiboWebFetchSearchDataGet

> ResponseModel FetchSearchDataApiV1WeiboWebFetchSearchDataGet(ctx, keyword, optional)

获取搜索数据/Get search data

# [中文] ### 用途: - 获取搜索数据 ### 参数: - keyword: 关键词 - page: 页数 - search_type: 搜索类型     - **1**: 综合     - **61**: 实时     - **3**: 用户     - **60**: 热门     - **64**: 视频     - **63**: 图片     - **21**: 文章     - **38**: 话题     - **98**: 超话 ### 返回: - 搜索数据  # [English] ### Purpose: - Get search data ### Parameters: - keyword: Keyword - page: Page number - search_type: Search type     - **1**: Comprehensive     - **61**: Real-time     - **3**: User     - **60**: Hot     - **64**: Video     - **63**: Picture     - **21**: Article     - **38**: Topic     - **98**: Super topic ### Return: - Search data  # [示例/Example] keyword = \"游戏\" page = \"1\" search_type = \"1\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyword** | **string**| 关键词/Keyword | 
 **optional** | ***FetchSearchDataApiV1WeiboWebFetchSearchDataGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchSearchDataApiV1WeiboWebFetchSearchDataGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.String**| 页数/Page number | [default to 1]
 **searchType** | **optional.String**| 搜索类型/Search type | [default to 1]

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


## FetchSearchDataApiV1WeiboWebFetchSearchDataGet_0

> ResponseModel FetchSearchDataApiV1WeiboWebFetchSearchDataGet_0(ctx, keyword, optional)

获取搜索数据/Get search data

# [中文] ### 用途: - 获取搜索数据 ### 参数: - keyword: 关键词 - page: 页数 - search_type: 搜索类型     - **1**: 综合     - **61**: 实时     - **3**: 用户     - **60**: 热门     - **64**: 视频     - **63**: 图片     - **21**: 文章     - **38**: 话题     - **98**: 超话 ### 返回: - 搜索数据  # [English] ### Purpose: - Get search data ### Parameters: - keyword: Keyword - page: Page number - search_type: Search type     - **1**: Comprehensive     - **61**: Real-time     - **3**: User     - **60**: Hot     - **64**: Video     - **63**: Picture     - **21**: Article     - **38**: Topic     - **98**: Super topic ### Return: - Search data  # [示例/Example] keyword = \"游戏\" page = \"1\" search_type = \"1\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyword** | **string**| 关键词/Keyword | 
 **optional** | ***FetchSearchDataApiV1WeiboWebFetchSearchDataGet_2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchSearchDataApiV1WeiboWebFetchSearchDataGet_2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.String**| 页数/Page number | [default to 1]
 **searchType** | **optional.String**| 搜索类型/Search type | [default to 1]

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


## FetchShortVideoDataApiV1WeiboWebFetchShortVideoDataGet

> ResponseModel FetchShortVideoDataApiV1WeiboWebFetchShortVideoDataGet(ctx, shareText)

获取短视频数据/Get short video data

# [中文] ### 用途: - 获取短视频数据 ### 参数: - share_text: 视频分享链接 ### 返回: - 短视频数据  # [English] ### Purpose: - Get short video data ### Parameters: - share_text: Video sharing link ### Return: - Short video data  # [示例/Example] share_text = \"https://video.weibo.com/show?fid=1034:4986069612167172\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareText** | **string**| 视频分享链接/Video sharing link | 

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


## FetchShortVideoDataApiV1WeiboWebFetchShortVideoDataGet_0

> ResponseModel FetchShortVideoDataApiV1WeiboWebFetchShortVideoDataGet_0(ctx, shareText)

获取短视频数据/Get short video data

# [中文] ### 用途: - 获取短视频数据 ### 参数: - share_text: 视频分享链接 ### 返回: - 短视频数据  # [English] ### Purpose: - Get short video data ### Parameters: - share_text: Video sharing link ### Return: - Short video data  # [示例/Example] share_text = \"https://video.weibo.com/show?fid=1034:4986069612167172\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareText** | **string**| 视频分享链接/Video sharing link | 

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


## FetchTopicDetailApiV1WeiboWebFetchTopicDetailGet

> ResponseModel FetchTopicDetailApiV1WeiboWebFetchTopicDetailGet(ctx, topicName, optional)

获取话题详情/Get topic details

# [中文] ### 用途: - 获取话题详情 ### 参数: - topic_name: 话题名称 - page: 页数 ### 返回: - 话题详情  # [English] ### Purpose: - Get topic details ### Parameters: - topic_name: Topic name - page: Page number ### Return: - Topic details  # [示例/Example] topic_name = \"游戏\" page = \"\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicName** | **string**| 话题名称/Topic name | 
 **optional** | ***FetchTopicDetailApiV1WeiboWebFetchTopicDetailGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchTopicDetailApiV1WeiboWebFetchTopicDetailGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.String**| 页数/Page number | [default to ]

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


## FetchTopicDetailApiV1WeiboWebFetchTopicDetailGet_0

> ResponseModel FetchTopicDetailApiV1WeiboWebFetchTopicDetailGet_0(ctx, topicName, optional)

获取话题详情/Get topic details

# [中文] ### 用途: - 获取话题详情 ### 参数: - topic_name: 话题名称 - page: 页数 ### 返回: - 话题详情  # [English] ### Purpose: - Get topic details ### Parameters: - topic_name: Topic name - page: Page number ### Return: - Topic details  # [示例/Example] topic_name = \"游戏\" page = \"\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicName** | **string**| 话题名称/Topic name | 
 **optional** | ***FetchTopicDetailApiV1WeiboWebFetchTopicDetailGet_4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchTopicDetailApiV1WeiboWebFetchTopicDetailGet_4Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.String**| 页数/Page number | [default to ]

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


## FetchTopicStatsApiV1WeiboWebFetchTopicStatsGet

> ResponseModel FetchTopicStatsApiV1WeiboWebFetchTopicStatsGet(ctx, topicName)

获取话题统计数据/Get topic statistics

# [中文] ### 用途: - 获取话题统计数据 ### 参数: - topic_name: 话题名称 ### 返回: - 话题统计数据  # [English] ### Purpose: - Get topic statistics ### Parameters: - topic_name: Topic name ### Return: - Topic statistics  # [示例/Example] topic_name = \"游戏\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicName** | **string**| 话题名称/Topic name | 

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


## FetchTopicStatsApiV1WeiboWebFetchTopicStatsGet_0

> ResponseModel FetchTopicStatsApiV1WeiboWebFetchTopicStatsGet_0(ctx, topicName)

获取话题统计数据/Get topic statistics

# [中文] ### 用途: - 获取话题统计数据 ### 参数: - topic_name: 话题名称 ### 返回: - 话题统计数据  # [English] ### Purpose: - Get topic statistics ### Parameters: - topic_name: Topic name ### Return: - Topic statistics  # [示例/Example] topic_name = \"游戏\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicName** | **string**| 话题名称/Topic name | 

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


## FetchUserInfoApiV1WeiboWebFetchUserInfoGet

> ResponseModel FetchUserInfoApiV1WeiboWebFetchUserInfoGet(ctx, uid)

获取用户信息/Get user information

# [中文] ### 用途: - 获取用户信息 ### 参数: - uid: 用户id ### 返回: - 用户信息  # [English] ### Purpose: - Get user information ### Parameters: - uid: User id ### Return: - User information  # [示例/Example] uid = \"7277477906\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string**| 用户id/User id | 

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


## FetchUserInfoApiV1WeiboWebFetchUserInfoGet_0

> ResponseModel FetchUserInfoApiV1WeiboWebFetchUserInfoGet_0(ctx, uid)

获取用户信息/Get user information

# [中文] ### 用途: - 获取用户信息 ### 参数: - uid: 用户id ### 返回: - 用户信息  # [English] ### Purpose: - Get user information ### Parameters: - uid: User id ### Return: - User information  # [示例/Example] uid = \"7277477906\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string**| 用户id/User id | 

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


## FetchUserInfoV2ApiV1WeiboWebFetchUserInfoV2Get

> ResponseModel FetchUserInfoV2ApiV1WeiboWebFetchUserInfoV2Get(ctx, uid)

获取用户信息V2/Get user information V2

# [中文] ### 用途: - 获取用户信息V2 ### 参数: - uid: 用户id ### 返回: - 用户信息  # [English] ### Purpose: - Get user information V2 ### Parameters: - uid: User id ### Return: - User information  # [示例/Example] uid = \"7277477906\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string**| 用户id/User id | 

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


## FetchUserInfoV2ApiV1WeiboWebFetchUserInfoV2Get_0

> ResponseModel FetchUserInfoV2ApiV1WeiboWebFetchUserInfoV2Get_0(ctx, uid)

获取用户信息V2/Get user information V2

# [中文] ### 用途: - 获取用户信息V2 ### 参数: - uid: 用户id ### 返回: - 用户信息  # [English] ### Purpose: - Get user information V2 ### Parameters: - uid: User id ### Return: - User information  # [示例/Example] uid = \"7277477906\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string**| 用户id/User id | 

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


## FetchUserPostsApiV1WeiboWebFetchUserPostsGet

> ResponseModel FetchUserPostsApiV1WeiboWebFetchUserPostsGet(ctx, uid, optional)

获取微博用户文章数据/Get Weibo user article data

# [中文] ### 用途: - 获取微博用户文章数据 ### 参数: - uid: 用户id - page: 页数 - feature: 特征 ### 返回: - 用户文章数据  # [English] ### Purpose: - Get Weibo user article data ### Parameters: - uid: User id - page: Page number - feature: Feature ### Return: - User article data  # [示例/Example] uid = \"7277477906\" page = 1 feature = 0

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string**| 用户id/User id | 
 **optional** | ***FetchUserPostsApiV1WeiboWebFetchUserPostsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchUserPostsApiV1WeiboWebFetchUserPostsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| 页数/Page number | [default to 1]
 **feature** | **optional.Int32**| 特征/Feature | [default to 0]

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


## FetchUserPostsApiV1WeiboWebFetchUserPostsGet_0

> ResponseModel FetchUserPostsApiV1WeiboWebFetchUserPostsGet_0(ctx, uid, optional)

获取微博用户文章数据/Get Weibo user article data

# [中文] ### 用途: - 获取微博用户文章数据 ### 参数: - uid: 用户id - page: 页数 - feature: 特征 ### 返回: - 用户文章数据  # [English] ### Purpose: - Get Weibo user article data ### Parameters: - uid: User id - page: Page number - feature: Feature ### Return: - User article data  # [示例/Example] uid = \"7277477906\" page = 1 feature = 0

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string**| 用户id/User id | 
 **optional** | ***FetchUserPostsApiV1WeiboWebFetchUserPostsGet_8Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchUserPostsApiV1WeiboWebFetchUserPostsGet_8Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| 页数/Page number | [default to 1]
 **feature** | **optional.Int32**| 特征/Feature | [default to 0]

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

