# \XiguaAppV2APIApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchOneVideoApiV1XiguaAppV2FetchOneVideoGet**](XiguaAppV2APIApi.md#FetchOneVideoApiV1XiguaAppV2FetchOneVideoGet) | **Get** /api/v1/xigua/app/v2/fetch_one_video | 获取单个作品数据/Get single video data
[**FetchOneVideoApiV1XiguaAppV2FetchOneVideoGet_0**](XiguaAppV2APIApi.md#FetchOneVideoApiV1XiguaAppV2FetchOneVideoGet_0) | **Get** /api/v1/xigua/app/v2/fetch_one_video | 获取单个作品数据/Get single video data
[**FetchOneVideoPlayUrlApiV1XiguaAppV2FetchOneVideoPlayUrlGet**](XiguaAppV2APIApi.md#FetchOneVideoPlayUrlApiV1XiguaAppV2FetchOneVideoPlayUrlGet) | **Get** /api/v1/xigua/app/v2/fetch_one_video_play_url | 获取单个作品的播放链接/Get single video play URL
[**FetchOneVideoPlayUrlApiV1XiguaAppV2FetchOneVideoPlayUrlGet_0**](XiguaAppV2APIApi.md#FetchOneVideoPlayUrlApiV1XiguaAppV2FetchOneVideoPlayUrlGet_0) | **Get** /api/v1/xigua/app/v2/fetch_one_video_play_url | 获取单个作品的播放链接/Get single video play URL
[**FetchOneVideoV2ApiV1XiguaAppV2FetchOneVideoV2Get**](XiguaAppV2APIApi.md#FetchOneVideoV2ApiV1XiguaAppV2FetchOneVideoV2Get) | **Get** /api/v1/xigua/app/v2/fetch_one_video_v2 | 获取单个作品数据 V2/Get single video data V2
[**FetchOneVideoV2ApiV1XiguaAppV2FetchOneVideoV2Get_0**](XiguaAppV2APIApi.md#FetchOneVideoV2ApiV1XiguaAppV2FetchOneVideoV2Get_0) | **Get** /api/v1/xigua/app/v2/fetch_one_video_v2 | 获取单个作品数据 V2/Get single video data V2
[**FetchUserInfoApiV1XiguaAppV2FetchUserInfoGet**](XiguaAppV2APIApi.md#FetchUserInfoApiV1XiguaAppV2FetchUserInfoGet) | **Get** /api/v1/xigua/app/v2/fetch_user_info | 个人信息/Personal information
[**FetchUserInfoApiV1XiguaAppV2FetchUserInfoGet_0**](XiguaAppV2APIApi.md#FetchUserInfoApiV1XiguaAppV2FetchUserInfoGet_0) | **Get** /api/v1/xigua/app/v2/fetch_user_info | 个人信息/Personal information
[**FetchUserPostListApiV1XiguaAppV2FetchUserPostListGet**](XiguaAppV2APIApi.md#FetchUserPostListApiV1XiguaAppV2FetchUserPostListGet) | **Get** /api/v1/xigua/app/v2/fetch_user_post_list | 获取个人作品列表/Get user post list
[**FetchUserPostListApiV1XiguaAppV2FetchUserPostListGet_0**](XiguaAppV2APIApi.md#FetchUserPostListApiV1XiguaAppV2FetchUserPostListGet_0) | **Get** /api/v1/xigua/app/v2/fetch_user_post_list | 获取个人作品列表/Get user post list
[**FetchVideoCommentListApiV1XiguaAppV2FetchVideoCommentListGet**](XiguaAppV2APIApi.md#FetchVideoCommentListApiV1XiguaAppV2FetchVideoCommentListGet) | **Get** /api/v1/xigua/app/v2/fetch_video_comment_list | 视频评论列表/Video comment list
[**FetchVideoCommentListApiV1XiguaAppV2FetchVideoCommentListGet_0**](XiguaAppV2APIApi.md#FetchVideoCommentListApiV1XiguaAppV2FetchVideoCommentListGet_0) | **Get** /api/v1/xigua/app/v2/fetch_video_comment_list | 视频评论列表/Video comment list
[**SearchVideoApiV1XiguaAppV2SearchVideoGet**](XiguaAppV2APIApi.md#SearchVideoApiV1XiguaAppV2SearchVideoGet) | **Get** /api/v1/xigua/app/v2/search_video | 搜索视频/Search video
[**SearchVideoApiV1XiguaAppV2SearchVideoGet_0**](XiguaAppV2APIApi.md#SearchVideoApiV1XiguaAppV2SearchVideoGet_0) | **Get** /api/v1/xigua/app/v2/search_video | 搜索视频/Search video



## FetchOneVideoApiV1XiguaAppV2FetchOneVideoGet

> ResponseModel FetchOneVideoApiV1XiguaAppV2FetchOneVideoGet(ctx, itemId)

获取单个作品数据/Get single video data

# [中文] ### 用途: - 获取单个作品数据（信息较少，不包含标题等信息，但是包含相关视频的信息） ### 参数: - item_id: 作品id ### 返回: - 作品数据，其中包含视频链接的Base64编码播放地址，需要前端解码后使用，或者使用 /fetch_one_video_play_url 获取播放链接。  # [English] ### Purpose: - Get single video data (less information, does not include title and other information, but includes information about related videos) ### Parameters: - item_id: Video id ### Return: - Video data, which contains the Base64 encoded playback address of the video link, which needs to be decoded and used by the front end, or use /fetch_one_video_play_url to get the playback link.  # [示例/Example] item_id: \"7354954305222377999\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string**| 作品id/Video id | 

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


## FetchOneVideoApiV1XiguaAppV2FetchOneVideoGet_0

> ResponseModel FetchOneVideoApiV1XiguaAppV2FetchOneVideoGet_0(ctx, itemId)

获取单个作品数据/Get single video data

# [中文] ### 用途: - 获取单个作品数据（信息较少，不包含标题等信息，但是包含相关视频的信息） ### 参数: - item_id: 作品id ### 返回: - 作品数据，其中包含视频链接的Base64编码播放地址，需要前端解码后使用，或者使用 /fetch_one_video_play_url 获取播放链接。  # [English] ### Purpose: - Get single video data (less information, does not include title and other information, but includes information about related videos) ### Parameters: - item_id: Video id ### Return: - Video data, which contains the Base64 encoded playback address of the video link, which needs to be decoded and used by the front end, or use /fetch_one_video_play_url to get the playback link.  # [示例/Example] item_id: \"7354954305222377999\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string**| 作品id/Video id | 

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


## FetchOneVideoPlayUrlApiV1XiguaAppV2FetchOneVideoPlayUrlGet

> ResponseModel FetchOneVideoPlayUrlApiV1XiguaAppV2FetchOneVideoPlayUrlGet(ctx, itemId)

获取单个作品的播放链接/Get single video play URL

# [中文] ### 用途: - 获取单个作品的播放链接，此接口返回的是已经解码后的播放链接，可以直接使用。 ### 参数: - item_id: 作品id ### 返回: - 作品的播放链接的明文链接。  # [English] ### Purpose: - Get single video play URL, the interface returns the decoded play URL, which can be used directly. ### Parameters: - item_id: Video id ### Return: - Play URL of the video, plaintext link.  # [示例/Example] item_id: \"7354954305222377999\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string**| 作品id/Video id | 

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


## FetchOneVideoPlayUrlApiV1XiguaAppV2FetchOneVideoPlayUrlGet_0

> ResponseModel FetchOneVideoPlayUrlApiV1XiguaAppV2FetchOneVideoPlayUrlGet_0(ctx, itemId)

获取单个作品的播放链接/Get single video play URL

# [中文] ### 用途: - 获取单个作品的播放链接，此接口返回的是已经解码后的播放链接，可以直接使用。 ### 参数: - item_id: 作品id ### 返回: - 作品的播放链接的明文链接。  # [English] ### Purpose: - Get single video play URL, the interface returns the decoded play URL, which can be used directly. ### Parameters: - item_id: Video id ### Return: - Play URL of the video, plaintext link.  # [示例/Example] item_id: \"7354954305222377999\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string**| 作品id/Video id | 

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


## FetchOneVideoV2ApiV1XiguaAppV2FetchOneVideoV2Get

> ResponseModel FetchOneVideoV2ApiV1XiguaAppV2FetchOneVideoV2Get(ctx, itemId)

获取单个作品数据 V2/Get single video data V2

# [中文] ### 用途: - 获取单个作品数据（信息全面，包含标题等信息，但是不包含相关视频推荐信息） ### 参数: - item_id: 作品id ### 返回: - 作品数据，其中包含视频链接的Base64编码播放地址，需要前端解码后使用，或者使用 /fetch_one_video_play_url 获取播放链接。  # [English] ### Purpose: - Get single video data (more comprehensive information, including title and other information, but not including information about related video recommendations) ### Parameters: - item_id: Video id ### Return: - Video data, which contains the Base64 encoded playback address of the video link, which needs to be decoded and used by the front end, or use /fetch_one_video_play_url to get the playback link.  # [示例/Example] item_id: \"7354954305222377999\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string**| 作品id/Video id | 

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


## FetchOneVideoV2ApiV1XiguaAppV2FetchOneVideoV2Get_0

> ResponseModel FetchOneVideoV2ApiV1XiguaAppV2FetchOneVideoV2Get_0(ctx, itemId)

获取单个作品数据 V2/Get single video data V2

# [中文] ### 用途: - 获取单个作品数据（信息全面，包含标题等信息，但是不包含相关视频推荐信息） ### 参数: - item_id: 作品id ### 返回: - 作品数据，其中包含视频链接的Base64编码播放地址，需要前端解码后使用，或者使用 /fetch_one_video_play_url 获取播放链接。  # [English] ### Purpose: - Get single video data (more comprehensive information, including title and other information, but not including information about related video recommendations) ### Parameters: - item_id: Video id ### Return: - Video data, which contains the Base64 encoded playback address of the video link, which needs to be decoded and used by the front end, or use /fetch_one_video_play_url to get the playback link.  # [示例/Example] item_id: \"7354954305222377999\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string**| 作品id/Video id | 

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


## FetchUserInfoApiV1XiguaAppV2FetchUserInfoGet

> ResponseModel FetchUserInfoApiV1XiguaAppV2FetchUserInfoGet(ctx, userId)

个人信息/Personal information

# [中文] ### 用途: - 个人信息 ### 参数: - user_id: 用户id ### 返回: - 个人信息  # [English] ### Purpose: - Personal information ### Parameters: - user_id: User id ### Return: - Personal information  # [示例/Example] user_id: \"52712347586\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| 用户id/User id | 

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


## FetchUserInfoApiV1XiguaAppV2FetchUserInfoGet_0

> ResponseModel FetchUserInfoApiV1XiguaAppV2FetchUserInfoGet_0(ctx, userId)

个人信息/Personal information

# [中文] ### 用途: - 个人信息 ### 参数: - user_id: 用户id ### 返回: - 个人信息  # [English] ### Purpose: - Personal information ### Parameters: - user_id: User id ### Return: - Personal information  # [示例/Example] user_id: \"52712347586\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| 用户id/User id | 

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


## FetchUserPostListApiV1XiguaAppV2FetchUserPostListGet

> ResponseModel FetchUserPostListApiV1XiguaAppV2FetchUserPostListGet(ctx, userId, optional)

获取个人作品列表/Get user post list

# [中文] ### 用途: - 获取个人作品列表 ### 参数: - user_id: 用户id - max_behot_time: 最大行为时间，默认空，第一次请求传空，后续请求传上一次请求返回数据中的JSON中的值。 - max_behot_time的值可以是JSON路径为：$.data.data.[-1].behot_time - 也就是data中的最后一个元素的cursor值 ### 返回: - 作品列表  # [English] ### Purpose: - Get user post list ### Parameters: - user_id: User id - max_behot_time: Maximum behavior time, default empty, pass empty for the first request, pass the max_behot_time returned by the previous request for subsequent requests - The value of max_behot_time can be the JSON path: $.data.data.[-1].behot_time - That is, the cursor value of the last element in data ### Return: - Post list  # [示例/Example] user_id = \"1922379661976311\" max_behot_time = \"\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| 用户id/User id | 
 **optional** | ***FetchUserPostListApiV1XiguaAppV2FetchUserPostListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchUserPostListApiV1XiguaAppV2FetchUserPostListGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxBehotTime** | **optional.String**| 最大行为时间/Maximum behavior time | 

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


## FetchUserPostListApiV1XiguaAppV2FetchUserPostListGet_0

> ResponseModel FetchUserPostListApiV1XiguaAppV2FetchUserPostListGet_0(ctx, userId, optional)

获取个人作品列表/Get user post list

# [中文] ### 用途: - 获取个人作品列表 ### 参数: - user_id: 用户id - max_behot_time: 最大行为时间，默认空，第一次请求传空，后续请求传上一次请求返回数据中的JSON中的值。 - max_behot_time的值可以是JSON路径为：$.data.data.[-1].behot_time - 也就是data中的最后一个元素的cursor值 ### 返回: - 作品列表  # [English] ### Purpose: - Get user post list ### Parameters: - user_id: User id - max_behot_time: Maximum behavior time, default empty, pass empty for the first request, pass the max_behot_time returned by the previous request for subsequent requests - The value of max_behot_time can be the JSON path: $.data.data.[-1].behot_time - That is, the cursor value of the last element in data ### Return: - Post list  # [示例/Example] user_id = \"1922379661976311\" max_behot_time = \"\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| 用户id/User id | 
 **optional** | ***FetchUserPostListApiV1XiguaAppV2FetchUserPostListGet_5Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchUserPostListApiV1XiguaAppV2FetchUserPostListGet_5Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxBehotTime** | **optional.String**| 最大行为时间/Maximum behavior time | 

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


## FetchVideoCommentListApiV1XiguaAppV2FetchVideoCommentListGet

> ResponseModel FetchVideoCommentListApiV1XiguaAppV2FetchVideoCommentListGet(ctx, itemId, optional)

视频评论列表/Video comment list

# [中文] ### 用途: - 视频评论列表 ### 参数: - item_id: 作品id - offset: 偏移量，第一次请求传0，后续请求传上一次请求返回的offset - count: 数量，默认20，建议保持默认。 ### 返回: - 评论列表  # [English] ### Purpose: - Video comment list ### Parameters: - item_id: Video id - offset: Offset, pass 0 for the first request, pass the offset returned by the previous request for subsequent requests - count: Quantity, default 20, it is recommended to keep the default. ### Return: - Comment list  # [示例/Example] item_id: \"7354954305222377999\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string**| 作品id/Video id | 
 **optional** | ***FetchVideoCommentListApiV1XiguaAppV2FetchVideoCommentListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchVideoCommentListApiV1XiguaAppV2FetchVideoCommentListGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| 偏移量/Offset | [default to 0]
 **count** | **optional.Int32**| 数量/Count | [default to 20]

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


## FetchVideoCommentListApiV1XiguaAppV2FetchVideoCommentListGet_0

> ResponseModel FetchVideoCommentListApiV1XiguaAppV2FetchVideoCommentListGet_0(ctx, itemId, optional)

视频评论列表/Video comment list

# [中文] ### 用途: - 视频评论列表 ### 参数: - item_id: 作品id - offset: 偏移量，第一次请求传0，后续请求传上一次请求返回的offset - count: 数量，默认20，建议保持默认。 ### 返回: - 评论列表  # [English] ### Purpose: - Video comment list ### Parameters: - item_id: Video id - offset: Offset, pass 0 for the first request, pass the offset returned by the previous request for subsequent requests - count: Quantity, default 20, it is recommended to keep the default. ### Return: - Comment list  # [示例/Example] item_id: \"7354954305222377999\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string**| 作品id/Video id | 
 **optional** | ***FetchVideoCommentListApiV1XiguaAppV2FetchVideoCommentListGet_6Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchVideoCommentListApiV1XiguaAppV2FetchVideoCommentListGet_6Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| 偏移量/Offset | [default to 0]
 **count** | **optional.Int32**| 数量/Count | [default to 20]

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


## SearchVideoApiV1XiguaAppV2SearchVideoGet

> ResponseModel SearchVideoApiV1XiguaAppV2SearchVideoGet(ctx, keyword, optional)

搜索视频/Search video

# [中文] ### 用途: - 搜索视频 ### 参数: - keyword: 关键词 - offset: 偏移量，第一次请求传0，后续请求传上一次请求返回的offset - order_type: 排序方式，为空时按照默认排序，以下为可选排序方式。     - 最新: publish_time     - 最热: play_count - min_duration: 最小时长，默认空，单位秒。 - max_duration: 最大时长，默认空，单位秒。 ### 返回: - 视频列表  # [English] ### Purpose: - Search video ### Parameters: - keyword: Keyword - offset: Offset, pass 0 for the first request, pass the offset returned by the previous request for subsequent requests - order_type: Order type, empty for default sorting, the following are optional sorting methods.     - Newest: publish_time     - Hottest: play_count - min_duration: Minimum duration, default empty, in seconds. - max_duration: Maximum duration, default empty, in seconds. ### Return: - Video list  # [示例/Example] > 搜索关键字为“抖音”的视频，按照播放量排序，时长1-180秒(1-3分钟) > Search for videos with the keyword \"抖音\", sorted by play count, duration 1-180 seconds (1-3 minutes) keyword: \"抖音\" order_type: \"play_count\" min_duration: 1 max_duration: 180

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyword** | **string**| 关键词/Keyword | 
 **optional** | ***SearchVideoApiV1XiguaAppV2SearchVideoGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchVideoApiV1XiguaAppV2SearchVideoGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| 偏移量/Offset | [default to 0]
 **orderType** | **optional.String**| 排序方式/Order type | 
 **minDuration** | **optional.Int32**| 最小时长/Minimum duration | 
 **maxDuration** | **optional.Int32**| 最大时长/Maximum duration | 

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


## SearchVideoApiV1XiguaAppV2SearchVideoGet_0

> ResponseModel SearchVideoApiV1XiguaAppV2SearchVideoGet_0(ctx, keyword, optional)

搜索视频/Search video

# [中文] ### 用途: - 搜索视频 ### 参数: - keyword: 关键词 - offset: 偏移量，第一次请求传0，后续请求传上一次请求返回的offset - order_type: 排序方式，为空时按照默认排序，以下为可选排序方式。     - 最新: publish_time     - 最热: play_count - min_duration: 最小时长，默认空，单位秒。 - max_duration: 最大时长，默认空，单位秒。 ### 返回: - 视频列表  # [English] ### Purpose: - Search video ### Parameters: - keyword: Keyword - offset: Offset, pass 0 for the first request, pass the offset returned by the previous request for subsequent requests - order_type: Order type, empty for default sorting, the following are optional sorting methods.     - Newest: publish_time     - Hottest: play_count - min_duration: Minimum duration, default empty, in seconds. - max_duration: Maximum duration, default empty, in seconds. ### Return: - Video list  # [示例/Example] > 搜索关键字为“抖音”的视频，按照播放量排序，时长1-180秒(1-3分钟) > Search for videos with the keyword \"抖音\", sorted by play count, duration 1-180 seconds (1-3 minutes) keyword: \"抖音\" order_type: \"play_count\" min_duration: 1 max_duration: 180

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyword** | **string**| 关键词/Keyword | 
 **optional** | ***SearchVideoApiV1XiguaAppV2SearchVideoGet_7Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchVideoApiV1XiguaAppV2SearchVideoGet_7Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| 偏移量/Offset | [default to 0]
 **orderType** | **optional.String**| 排序方式/Order type | 
 **minDuration** | **optional.Int32**| 最小时长/Minimum duration | 
 **maxDuration** | **optional.Int32**| 最大时长/Maximum duration | 

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

