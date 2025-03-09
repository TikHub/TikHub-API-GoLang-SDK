# \YouTubeWebAPIApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChannelIdApiV1YoutubeWebGetChannelIdGet**](YouTubeWebAPIApi.md#GetChannelIdApiV1YoutubeWebGetChannelIdGet) | **Get** /api/v1/youtube/web/get_channel_id | 获取频道ID/Get channel ID
[**GetChannelIdApiV1YoutubeWebGetChannelIdGet_0**](YouTubeWebAPIApi.md#GetChannelIdApiV1YoutubeWebGetChannelIdGet_0) | **Get** /api/v1/youtube/web/get_channel_id | 获取频道ID/Get channel ID
[**GetChannelInfoApiV1YoutubeWebGetChannelInfoGet**](YouTubeWebAPIApi.md#GetChannelInfoApiV1YoutubeWebGetChannelInfoGet) | **Get** /api/v1/youtube/web/get_channel_info | 获取频道信息/Get channel information
[**GetChannelInfoApiV1YoutubeWebGetChannelInfoGet_0**](YouTubeWebAPIApi.md#GetChannelInfoApiV1YoutubeWebGetChannelInfoGet_0) | **Get** /api/v1/youtube/web/get_channel_info | 获取频道信息/Get channel information
[**GetChannelShortVideosApiV1YoutubeWebGetChannelShortVideosGet**](YouTubeWebAPIApi.md#GetChannelShortVideosApiV1YoutubeWebGetChannelShortVideosGet) | **Get** /api/v1/youtube/web/get_channel_short_videos | 获取频道短视频/Get channel short videos
[**GetChannelShortVideosApiV1YoutubeWebGetChannelShortVideosGet_0**](YouTubeWebAPIApi.md#GetChannelShortVideosApiV1YoutubeWebGetChannelShortVideosGet_0) | **Get** /api/v1/youtube/web/get_channel_short_videos | 获取频道短视频/Get channel short videos
[**GetChannelVideosApiV1YoutubeWebGetChannelVideosGet**](YouTubeWebAPIApi.md#GetChannelVideosApiV1YoutubeWebGetChannelVideosGet) | **Get** /api/v1/youtube/web/get_channel_videos | 获取频道视频/Get channel videos
[**GetChannelVideosApiV1YoutubeWebGetChannelVideosGet_0**](YouTubeWebAPIApi.md#GetChannelVideosApiV1YoutubeWebGetChannelVideosGet_0) | **Get** /api/v1/youtube/web/get_channel_videos | 获取频道视频/Get channel videos
[**GetRelateVideoApiV1YoutubeWebGetRelateVideoGet**](YouTubeWebAPIApi.md#GetRelateVideoApiV1YoutubeWebGetRelateVideoGet) | **Get** /api/v1/youtube/web/get_relate_video | 获取推荐视频/Get related videos
[**GetRelateVideoApiV1YoutubeWebGetRelateVideoGet_0**](YouTubeWebAPIApi.md#GetRelateVideoApiV1YoutubeWebGetRelateVideoGet_0) | **Get** /api/v1/youtube/web/get_relate_video | 获取推荐视频/Get related videos
[**GetTrendingVideosApiV1YoutubeWebGetTrendingVideosGet**](YouTubeWebAPIApi.md#GetTrendingVideosApiV1YoutubeWebGetTrendingVideosGet) | **Get** /api/v1/youtube/web/get_trending_videos | 获取趋势视频/Get trending videos
[**GetTrendingVideosApiV1YoutubeWebGetTrendingVideosGet_0**](YouTubeWebAPIApi.md#GetTrendingVideosApiV1YoutubeWebGetTrendingVideosGet_0) | **Get** /api/v1/youtube/web/get_trending_videos | 获取趋势视频/Get trending videos
[**GetVideoCommentsApiV1YoutubeWebGetVideoCommentsGet**](YouTubeWebAPIApi.md#GetVideoCommentsApiV1YoutubeWebGetVideoCommentsGet) | **Get** /api/v1/youtube/web/get_video_comments | 获取视频评论/Get video comments
[**GetVideoCommentsApiV1YoutubeWebGetVideoCommentsGet_0**](YouTubeWebAPIApi.md#GetVideoCommentsApiV1YoutubeWebGetVideoCommentsGet_0) | **Get** /api/v1/youtube/web/get_video_comments | 获取视频评论/Get video comments
[**GetVideoInfoApiV1YoutubeWebGetVideoInfoGet**](YouTubeWebAPIApi.md#GetVideoInfoApiV1YoutubeWebGetVideoInfoGet) | **Get** /api/v1/youtube/web/get_video_info | 获取视频信息/Get video information
[**GetVideoInfoApiV1YoutubeWebGetVideoInfoGet_0**](YouTubeWebAPIApi.md#GetVideoInfoApiV1YoutubeWebGetVideoInfoGet_0) | **Get** /api/v1/youtube/web/get_video_info | 获取视频信息/Get video information
[**GetVideoSubtitlesApiV1YoutubeWebGetVideoSubtitlesGet**](YouTubeWebAPIApi.md#GetVideoSubtitlesApiV1YoutubeWebGetVideoSubtitlesGet) | **Get** /api/v1/youtube/web/get_video_subtitles | 获取视频字幕/Get video subtitles
[**GetVideoSubtitlesApiV1YoutubeWebGetVideoSubtitlesGet_0**](YouTubeWebAPIApi.md#GetVideoSubtitlesApiV1YoutubeWebGetVideoSubtitlesGet_0) | **Get** /api/v1/youtube/web/get_video_subtitles | 获取视频字幕/Get video subtitles
[**SearchChannelApiV1YoutubeWebSearchChannelGet**](YouTubeWebAPIApi.md#SearchChannelApiV1YoutubeWebSearchChannelGet) | **Get** /api/v1/youtube/web/search_channel | 搜索频道/Search channel
[**SearchChannelApiV1YoutubeWebSearchChannelGet_0**](YouTubeWebAPIApi.md#SearchChannelApiV1YoutubeWebSearchChannelGet_0) | **Get** /api/v1/youtube/web/search_channel | 搜索频道/Search channel
[**SearchVideoApiV1YoutubeWebSearchVideoGet**](YouTubeWebAPIApi.md#SearchVideoApiV1YoutubeWebSearchVideoGet) | **Get** /api/v1/youtube/web/search_video | 搜索视频/Search video
[**SearchVideoApiV1YoutubeWebSearchVideoGet_0**](YouTubeWebAPIApi.md#SearchVideoApiV1YoutubeWebSearchVideoGet_0) | **Get** /api/v1/youtube/web/search_video | 搜索视频/Search video



## GetChannelIdApiV1YoutubeWebGetChannelIdGet

> ResponseModel GetChannelIdApiV1YoutubeWebGetChannelIdGet(ctx, channelName)

获取频道ID/Get channel ID

# [中文] ### 用途: - 获取频道ID。 ### 参数: - channel_name: 频道名称。 ### 返回: - 频道ID。  # [English] ### Purpose: - Get channel ID. ### Parameters: - channel_name: Channel name. ### Returns: - Channel ID.  # [示例/Example] channel_name = \"LinusTechTips\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelName** | **string**| 频道名称/Channel name | 

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


## GetChannelIdApiV1YoutubeWebGetChannelIdGet_0

> ResponseModel GetChannelIdApiV1YoutubeWebGetChannelIdGet_0(ctx, channelName)

获取频道ID/Get channel ID

# [中文] ### 用途: - 获取频道ID。 ### 参数: - channel_name: 频道名称。 ### 返回: - 频道ID。  # [English] ### Purpose: - Get channel ID. ### Parameters: - channel_name: Channel name. ### Returns: - Channel ID.  # [示例/Example] channel_name = \"LinusTechTips\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelName** | **string**| 频道名称/Channel name | 

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


## GetChannelInfoApiV1YoutubeWebGetChannelInfoGet

> ResponseModel GetChannelInfoApiV1YoutubeWebGetChannelInfoGet(ctx, channelId)

获取频道信息/Get channel information

# [中文] ### 用途: - 获取频道信息。 ### 参数: - channel_id: 频道ID。 ### 返回: - 频道信息。  # [English] ### Purpose: - Get channel information. ### Parameters: - channel_id: Channel ID. ### Returns: - Channel information.  # [示例/Example] channel_id = \"UCXuqSBlHAE6Xw-yeJA0Tunw\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string**| 频道ID/Channel ID | 

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


## GetChannelInfoApiV1YoutubeWebGetChannelInfoGet_0

> ResponseModel GetChannelInfoApiV1YoutubeWebGetChannelInfoGet_0(ctx, channelId)

获取频道信息/Get channel information

# [中文] ### 用途: - 获取频道信息。 ### 参数: - channel_id: 频道ID。 ### 返回: - 频道信息。  # [English] ### Purpose: - Get channel information. ### Parameters: - channel_id: Channel ID. ### Returns: - Channel information.  # [示例/Example] channel_id = \"UCXuqSBlHAE6Xw-yeJA0Tunw\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string**| 频道ID/Channel ID | 

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


## GetChannelShortVideosApiV1YoutubeWebGetChannelShortVideosGet

> ResponseModel GetChannelShortVideosApiV1YoutubeWebGetChannelShortVideosGet(ctx, channelId, optional)

获取频道短视频/Get channel short videos

# [中文] ### 用途: - 获取频道短视频。 ### 参数: - channel_id: 频道ID。 - continuation_token: 用于继续获取频道短视频的令牌。默认为None。 ### 返回: - 频道短视频。  # [English] ### Purpose: - Get channel short videos. ### Parameters: - channel_id: Channel ID. - continuation_token: Token to continue fetching channel short videos. Default is None. ### Returns: - Channel short videos.  # [示例/Example] channel_id = \"UCXuqSBlHAE6Xw-yeJA0Tunw\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string**| 频道ID/Channel ID | 
 **optional** | ***GetChannelShortVideosApiV1YoutubeWebGetChannelShortVideosGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetChannelShortVideosApiV1YoutubeWebGetChannelShortVideosGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **continuationToken** | **optional.String**| 翻页令牌/Pagination token | 

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


## GetChannelShortVideosApiV1YoutubeWebGetChannelShortVideosGet_0

> ResponseModel GetChannelShortVideosApiV1YoutubeWebGetChannelShortVideosGet_0(ctx, channelId, optional)

获取频道短视频/Get channel short videos

# [中文] ### 用途: - 获取频道短视频。 ### 参数: - channel_id: 频道ID。 - continuation_token: 用于继续获取频道短视频的令牌。默认为None。 ### 返回: - 频道短视频。  # [English] ### Purpose: - Get channel short videos. ### Parameters: - channel_id: Channel ID. - continuation_token: Token to continue fetching channel short videos. Default is None. ### Returns: - Channel short videos.  # [示例/Example] channel_id = \"UCXuqSBlHAE6Xw-yeJA0Tunw\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string**| 频道ID/Channel ID | 
 **optional** | ***GetChannelShortVideosApiV1YoutubeWebGetChannelShortVideosGet_3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetChannelShortVideosApiV1YoutubeWebGetChannelShortVideosGet_3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **continuationToken** | **optional.String**| 翻页令牌/Pagination token | 

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


## GetChannelVideosApiV1YoutubeWebGetChannelVideosGet

> ResponseModel GetChannelVideosApiV1YoutubeWebGetChannelVideosGet(ctx, channelId, optional)

获取频道视频/Get channel videos

# [中文] ### 用途: - 获取频道视频。 ### 参数: - channel_id: 频道ID。 - continuation_token: 用于继续获取频道视频的令牌。默认为None。 ### 返回: - 频道视频。  # [English] ### Purpose: - Get channel videos. ### Parameters: - channel_id: Channel ID. - continuation_token: Token to continue fetching channel videos. Default is None. ### Returns: - Channel videos.  # [示例/Example] channel_id = \"UCXuqSBlHAE6Xw-yeJA0Tunw\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string**| 频道ID/Channel ID | 
 **optional** | ***GetChannelVideosApiV1YoutubeWebGetChannelVideosGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetChannelVideosApiV1YoutubeWebGetChannelVideosGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **continuationToken** | **optional.String**| 翻页令牌/Pagination token | 

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


## GetChannelVideosApiV1YoutubeWebGetChannelVideosGet_0

> ResponseModel GetChannelVideosApiV1YoutubeWebGetChannelVideosGet_0(ctx, channelId, optional)

获取频道视频/Get channel videos

# [中文] ### 用途: - 获取频道视频。 ### 参数: - channel_id: 频道ID。 - continuation_token: 用于继续获取频道视频的令牌。默认为None。 ### 返回: - 频道视频。  # [English] ### Purpose: - Get channel videos. ### Parameters: - channel_id: Channel ID. - continuation_token: Token to continue fetching channel videos. Default is None. ### Returns: - Channel videos.  # [示例/Example] channel_id = \"UCXuqSBlHAE6Xw-yeJA0Tunw\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string**| 频道ID/Channel ID | 
 **optional** | ***GetChannelVideosApiV1YoutubeWebGetChannelVideosGet_4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetChannelVideosApiV1YoutubeWebGetChannelVideosGet_4Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **continuationToken** | **optional.String**| 翻页令牌/Pagination token | 

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


## GetRelateVideoApiV1YoutubeWebGetRelateVideoGet

> ResponseModel GetRelateVideoApiV1YoutubeWebGetRelateVideoGet(ctx, videoId, optional)

获取推荐视频/Get related videos

# [中文] ### 用途: - 根据视频ID获取推荐视频数据。 ### 参数: - video_id: 视频ID，从URL中获取，例如：https://www.youtube.com/watch?v=LuIL5JATZsc，这里的video_id就是LuIL5JATZsc。 - continuation_token: 用于继续获取推荐视频的令牌。默认为None。 ### 返回: - 推荐视频数据。  # [English] ### Purpose: - Get related videos by video ID. ### Parameters: - video_id: Video ID, get it from the URL, for example: https://www.youtube.com/watch?v=LuIL5JATZsc, the id is LuIL5JATZsc. - continuation_token: Token to continue fetching related videos. Default is None. ### Returns: - Related videos.  # [示例/Example] video_id = \"LuIL5JATZsc\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string**| 视频ID/Video ID | 
 **optional** | ***GetRelateVideoApiV1YoutubeWebGetRelateVideoGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRelateVideoApiV1YoutubeWebGetRelateVideoGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **continuationToken** | **optional.String**| 翻页令牌/Pagination token | 

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


## GetRelateVideoApiV1YoutubeWebGetRelateVideoGet_0

> ResponseModel GetRelateVideoApiV1YoutubeWebGetRelateVideoGet_0(ctx, videoId, optional)

获取推荐视频/Get related videos

# [中文] ### 用途: - 根据视频ID获取推荐视频数据。 ### 参数: - video_id: 视频ID，从URL中获取，例如：https://www.youtube.com/watch?v=LuIL5JATZsc，这里的video_id就是LuIL5JATZsc。 - continuation_token: 用于继续获取推荐视频的令牌。默认为None。 ### 返回: - 推荐视频数据。  # [English] ### Purpose: - Get related videos by video ID. ### Parameters: - video_id: Video ID, get it from the URL, for example: https://www.youtube.com/watch?v=LuIL5JATZsc, the id is LuIL5JATZsc. - continuation_token: Token to continue fetching related videos. Default is None. ### Returns: - Related videos.  # [示例/Example] video_id = \"LuIL5JATZsc\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string**| 视频ID/Video ID | 
 **optional** | ***GetRelateVideoApiV1YoutubeWebGetRelateVideoGet_5Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetRelateVideoApiV1YoutubeWebGetRelateVideoGet_5Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **continuationToken** | **optional.String**| 翻页令牌/Pagination token | 

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


## GetTrendingVideosApiV1YoutubeWebGetTrendingVideosGet

> ResponseModel GetTrendingVideosApiV1YoutubeWebGetTrendingVideosGet(ctx, optional)

获取趋势视频/Get trending videos

# [中文] ### 用途: - 获取趋势视频。 ### 参数: - language_code: 语言代码，默认为en。 - country_code: 国家代码，默认为us。 - section: 类型，默认为Now，可选值为Music, Gaming, Movies。 ### 返回: - 趋势视频。  # [English] ### Purpose: - Get trending videos. ### Parameters: - language_code: Language code, default is en. - country_code: Country code, default is us. - section: Section, default is Now, optional values are Music, Gaming, Movies. ### Returns: - Trending videos.  # [示例/Example]

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTrendingVideosApiV1YoutubeWebGetTrendingVideosGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTrendingVideosApiV1YoutubeWebGetTrendingVideosGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **languageCode** | **optional.String**| 语言代码/Language code | [default to en]
 **countryCode** | **optional.String**| 国家代码/Country code | [default to us]
 **section** | **optional.String**| 类型/Section | [default to Now]

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


## GetTrendingVideosApiV1YoutubeWebGetTrendingVideosGet_0

> ResponseModel GetTrendingVideosApiV1YoutubeWebGetTrendingVideosGet_0(ctx, optional)

获取趋势视频/Get trending videos

# [中文] ### 用途: - 获取趋势视频。 ### 参数: - language_code: 语言代码，默认为en。 - country_code: 国家代码，默认为us。 - section: 类型，默认为Now，可选值为Music, Gaming, Movies。 ### 返回: - 趋势视频。  # [English] ### Purpose: - Get trending videos. ### Parameters: - language_code: Language code, default is en. - country_code: Country code, default is us. - section: Section, default is Now, optional values are Music, Gaming, Movies. ### Returns: - Trending videos.  # [示例/Example]

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTrendingVideosApiV1YoutubeWebGetTrendingVideosGet_6Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetTrendingVideosApiV1YoutubeWebGetTrendingVideosGet_6Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **languageCode** | **optional.String**| 语言代码/Language code | [default to en]
 **countryCode** | **optional.String**| 国家代码/Country code | [default to us]
 **section** | **optional.String**| 类型/Section | [default to Now]

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


## GetVideoCommentsApiV1YoutubeWebGetVideoCommentsGet

> ResponseModel GetVideoCommentsApiV1YoutubeWebGetVideoCommentsGet(ctx, videoId, optional)

获取视频评论/Get video comments

# [中文] ### 用途: - 获取单个视频的评论。 ### 参数: - video_id: 视频ID，从URL中获取，例如：https://www.youtube.com/watch?v=LuIL5JATZsc，这里的id就是LuIL5JATZsc。 - continuation_token: 用于继续获取评论的令牌。默认为None。 ### 返回: - 视频评论。  # [English] ### Purpose: - Get comments of a single video. ### Parameters: - id: Video ID, get it from the URL, for example: https://www.youtube.com/watch?v=LuIL5JATZsc, the id is LuIL5JATZsc. - continuation_token: Token to continue fetching comments. Default is None. ### Returns: - Video comments.  # [示例/Example] id = \"LuIL5JATZsc\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string**| 视频ID/Video ID | 
 **optional** | ***GetVideoCommentsApiV1YoutubeWebGetVideoCommentsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetVideoCommentsApiV1YoutubeWebGetVideoCommentsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **continuationToken** | **optional.String**| 翻页令牌/Pagination token | 

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


## GetVideoCommentsApiV1YoutubeWebGetVideoCommentsGet_0

> ResponseModel GetVideoCommentsApiV1YoutubeWebGetVideoCommentsGet_0(ctx, videoId, optional)

获取视频评论/Get video comments

# [中文] ### 用途: - 获取单个视频的评论。 ### 参数: - video_id: 视频ID，从URL中获取，例如：https://www.youtube.com/watch?v=LuIL5JATZsc，这里的id就是LuIL5JATZsc。 - continuation_token: 用于继续获取评论的令牌。默认为None。 ### 返回: - 视频评论。  # [English] ### Purpose: - Get comments of a single video. ### Parameters: - id: Video ID, get it from the URL, for example: https://www.youtube.com/watch?v=LuIL5JATZsc, the id is LuIL5JATZsc. - continuation_token: Token to continue fetching comments. Default is None. ### Returns: - Video comments.  # [示例/Example] id = \"LuIL5JATZsc\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string**| 视频ID/Video ID | 
 **optional** | ***GetVideoCommentsApiV1YoutubeWebGetVideoCommentsGet_7Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetVideoCommentsApiV1YoutubeWebGetVideoCommentsGet_7Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **continuationToken** | **optional.String**| 翻页令牌/Pagination token | 

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


## GetVideoInfoApiV1YoutubeWebGetVideoInfoGet

> ResponseModel GetVideoInfoApiV1YoutubeWebGetVideoInfoGet(ctx, videoId)

获取视频信息/Get video information

# [中文] ### 用途: - 获取单个视频的信息，包括视频下载链接、标题、作者、发布日期、观看次数、点赞次数、不喜欢次数、评论次数、时长、描述等。 ### 参数: - video_id: 视频ID，从URL中获取，例如：https://www.youtube.com/watch?v=LuIL5JATZsc，这里的video_id就是LuIL5JATZsc。 ### 返回: - 视频信息。  # [English] ### Purpose: - Get information of a single video, including video download link, title, author, publish date, view count, like count, dislike count, comment count, duration, description, etc. ### Parameters: - video_id: Video ID, get it from the URL, for example: https://www.youtube.com/watch?v=LuIL5JATZsc, the id is LuIL5JATZsc. ### Returns: - Video information.  # [示例/Example] video_id = \"LuIL5JATZsc\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string**| 视频ID/Video ID | 

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


## GetVideoInfoApiV1YoutubeWebGetVideoInfoGet_0

> ResponseModel GetVideoInfoApiV1YoutubeWebGetVideoInfoGet_0(ctx, videoId)

获取视频信息/Get video information

# [中文] ### 用途: - 获取单个视频的信息，包括视频下载链接、标题、作者、发布日期、观看次数、点赞次数、不喜欢次数、评论次数、时长、描述等。 ### 参数: - video_id: 视频ID，从URL中获取，例如：https://www.youtube.com/watch?v=LuIL5JATZsc，这里的video_id就是LuIL5JATZsc。 ### 返回: - 视频信息。  # [English] ### Purpose: - Get information of a single video, including video download link, title, author, publish date, view count, like count, dislike count, comment count, duration, description, etc. ### Parameters: - video_id: Video ID, get it from the URL, for example: https://www.youtube.com/watch?v=LuIL5JATZsc, the id is LuIL5JATZsc. ### Returns: - Video information.  # [示例/Example] video_id = \"LuIL5JATZsc\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string**| 视频ID/Video ID | 

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


## GetVideoSubtitlesApiV1YoutubeWebGetVideoSubtitlesGet

> ResponseModel GetVideoSubtitlesApiV1YoutubeWebGetVideoSubtitlesGet(ctx, videoId)

获取视频字幕/Get video subtitles

# [中文] ### 用途: - 获取单个视频的字幕。 ### 参数: - video_id: 视频ID，从URL中获取，例如：https://www.youtube.com/watch?v=LuIL5JATZsc，这里的video_id就是LuIL5JATZsc。 ### 返回: - 视频字幕。  # [English] ### Purpose: - Get subtitles of a single video. ### Parameters: - video_id: Video ID, get it from the URL, for example: https://www.youtube.com/watch?v=LuIL5JATZsc, the id is LuIL5JATZsc.  ### Returns: - Video subtitles.  # [示例/Example] video_id = \"LuIL5JATZsc\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string**| 视频ID/Video ID | 

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


## GetVideoSubtitlesApiV1YoutubeWebGetVideoSubtitlesGet_0

> ResponseModel GetVideoSubtitlesApiV1YoutubeWebGetVideoSubtitlesGet_0(ctx, videoId)

获取视频字幕/Get video subtitles

# [中文] ### 用途: - 获取单个视频的字幕。 ### 参数: - video_id: 视频ID，从URL中获取，例如：https://www.youtube.com/watch?v=LuIL5JATZsc，这里的video_id就是LuIL5JATZsc。 ### 返回: - 视频字幕。  # [English] ### Purpose: - Get subtitles of a single video. ### Parameters: - video_id: Video ID, get it from the URL, for example: https://www.youtube.com/watch?v=LuIL5JATZsc, the id is LuIL5JATZsc.  ### Returns: - Video subtitles.  # [示例/Example] video_id = \"LuIL5JATZsc\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string**| 视频ID/Video ID | 

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


## SearchChannelApiV1YoutubeWebSearchChannelGet

> ResponseModel SearchChannelApiV1YoutubeWebSearchChannelGet(ctx, channelId, searchQuery, optional)

搜索频道/Search channel

# [中文] ### 用途: - 搜索频道。 ### 参数: - search_query: 搜索关键字。 - language_code: 语言代码，默认为en。 - country_code: 国家代码，默认为us。 - continuation_token: 用于继续获取搜索结果的令牌。默认为None。 ### 返回: - 搜索结果。  # [English] ### Purpose: - Search channel. ### Parameters: - search_query: Search keyword. - language_code: Language code, default is en. - country_code: Country code, default is us. - continuation_token: Token to continue fetching search results. Default is None. ### Returns: - Search results.  # [示例/Example] channel_id = \"UCXuqSBlHAE6Xw-yeJA0Tunw\" search_query = \"AMD\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string**| 频道ID/Channel ID | 
**searchQuery** | **string**| 搜索关键字/Search keyword | 
 **optional** | ***SearchChannelApiV1YoutubeWebSearchChannelGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchChannelApiV1YoutubeWebSearchChannelGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **languageCode** | **optional.String**| 语言代码/Language code | [default to en]
 **countryCode** | **optional.String**| 国家代码/Country code | [default to us]
 **continuationToken** | **optional.String**| 翻页令牌/Pagination token | 

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


## SearchChannelApiV1YoutubeWebSearchChannelGet_0

> ResponseModel SearchChannelApiV1YoutubeWebSearchChannelGet_0(ctx, channelId, searchQuery, optional)

搜索频道/Search channel

# [中文] ### 用途: - 搜索频道。 ### 参数: - search_query: 搜索关键字。 - language_code: 语言代码，默认为en。 - country_code: 国家代码，默认为us。 - continuation_token: 用于继续获取搜索结果的令牌。默认为None。 ### 返回: - 搜索结果。  # [English] ### Purpose: - Search channel. ### Parameters: - search_query: Search keyword. - language_code: Language code, default is en. - country_code: Country code, default is us. - continuation_token: Token to continue fetching search results. Default is None. ### Returns: - Search results.  # [示例/Example] channel_id = \"UCXuqSBlHAE6Xw-yeJA0Tunw\" search_query = \"AMD\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string**| 频道ID/Channel ID | 
**searchQuery** | **string**| 搜索关键字/Search keyword | 
 **optional** | ***SearchChannelApiV1YoutubeWebSearchChannelGet_10Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchChannelApiV1YoutubeWebSearchChannelGet_10Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **languageCode** | **optional.String**| 语言代码/Language code | [default to en]
 **countryCode** | **optional.String**| 国家代码/Country code | [default to us]
 **continuationToken** | **optional.String**| 翻页令牌/Pagination token | 

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


## SearchVideoApiV1YoutubeWebSearchVideoGet

> ResponseModel SearchVideoApiV1YoutubeWebSearchVideoGet(ctx, searchQuery, optional)

搜索视频/Search video

# [中文] ### 用途: - 搜索视频。 ### 参数: - search_query: 搜索关键字。 - language_code: 语言代码，默认为en。 - order_by: 排序方式，默��为this_month。 - country_code: 国家代码，默认为us。 - continuation_token: 用于继续获取搜索结果的令牌。默认为None。 ### 返回: - 搜索结果。  # [English] ### Purpose: - Search video. ### Parameters: - search_query: Search keyword. - language_code: Language code, default is en. - order_by: Order by, default is this_month. - country_code: Country code, default is us. - continuation_token: Token to continue fetching search results. Default is None. ### Returns: - Search results.  # [示例/Example] search_query = \"Minecraft\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchQuery** | **string**| 搜索关键字/Search keyword | 
 **optional** | ***SearchVideoApiV1YoutubeWebSearchVideoGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchVideoApiV1YoutubeWebSearchVideoGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **languageCode** | **optional.String**| 语言代码/Language code | [default to en]
 **orderBy** | **optional.String**| 排序方式/Order by | [default to this_month]
 **countryCode** | **optional.String**| 国家代码/Country code | [default to us]
 **continuationToken** | **optional.String**| 翻页令牌/Pagination token | 

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


## SearchVideoApiV1YoutubeWebSearchVideoGet_0

> ResponseModel SearchVideoApiV1YoutubeWebSearchVideoGet_0(ctx, searchQuery, optional)

搜索视频/Search video

# [中文] ### 用途: - 搜索视频。 ### 参数: - search_query: 搜索关键字。 - language_code: 语言代码，默认为en。 - order_by: 排序方式，默��为this_month。 - country_code: 国家代码，默认为us。 - continuation_token: 用于继续获取搜索结果的令牌。默认为None。 ### 返回: - 搜索结果。  # [English] ### Purpose: - Search video. ### Parameters: - search_query: Search keyword. - language_code: Language code, default is en. - order_by: Order by, default is this_month. - country_code: Country code, default is us. - continuation_token: Token to continue fetching search results. Default is None. ### Returns: - Search results.  # [示例/Example] search_query = \"Minecraft\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**searchQuery** | **string**| 搜索关键字/Search keyword | 
 **optional** | ***SearchVideoApiV1YoutubeWebSearchVideoGet_11Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchVideoApiV1YoutubeWebSearchVideoGet_11Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **languageCode** | **optional.String**| 语言代码/Language code | [default to en]
 **orderBy** | **optional.String**| 排序方式/Order by | [default to this_month]
 **countryCode** | **optional.String**| 国家代码/Country code | [default to us]
 **continuationToken** | **optional.String**| 翻页令牌/Pagination token | 

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

