# \XiaohongshuWebV2APIApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchFeedNotesApiV1XiaohongshuWebV2FetchFeedNotesGet**](XiaohongshuWebV2APIApi.md#FetchFeedNotesApiV1XiaohongshuWebV2FetchFeedNotesGet) | **Get** /api/v1/xiaohongshu/web_v2/fetch_feed_notes | 获取推荐笔记/Fetch feed notes
[**FetchFeedNotesApiV1XiaohongshuWebV2FetchFeedNotesGet_0**](XiaohongshuWebV2APIApi.md#FetchFeedNotesApiV1XiaohongshuWebV2FetchFeedNotesGet_0) | **Get** /api/v1/xiaohongshu/web_v2/fetch_feed_notes | 获取推荐笔记/Fetch feed notes
[**FetchFollowerListApiV1XiaohongshuWebV2FetchFollowerListGet**](XiaohongshuWebV2APIApi.md#FetchFollowerListApiV1XiaohongshuWebV2FetchFollowerListGet) | **Get** /api/v1/xiaohongshu/web_v2/fetch_follower_list | 获取用户粉丝列表/Fetch follower list
[**FetchFollowerListApiV1XiaohongshuWebV2FetchFollowerListGet_0**](XiaohongshuWebV2APIApi.md#FetchFollowerListApiV1XiaohongshuWebV2FetchFollowerListGet_0) | **Get** /api/v1/xiaohongshu/web_v2/fetch_follower_list | 获取用户粉丝列表/Fetch follower list
[**FetchFollowingListApiV1XiaohongshuWebV2FetchFollowingListGet**](XiaohongshuWebV2APIApi.md#FetchFollowingListApiV1XiaohongshuWebV2FetchFollowingListGet) | **Get** /api/v1/xiaohongshu/web_v2/fetch_following_list | 获取用户关注列表/Fetch following list
[**FetchFollowingListApiV1XiaohongshuWebV2FetchFollowingListGet_0**](XiaohongshuWebV2APIApi.md#FetchFollowingListApiV1XiaohongshuWebV2FetchFollowingListGet_0) | **Get** /api/v1/xiaohongshu/web_v2/fetch_following_list | 获取用户关注列表/Fetch following list
[**FetchHomeNotesApiV1XiaohongshuWebV2FetchHomeNotesGet**](XiaohongshuWebV2APIApi.md#FetchHomeNotesApiV1XiaohongshuWebV2FetchHomeNotesGet) | **Get** /api/v1/xiaohongshu/web_v2/fetch_home_notes | 获取主页笔记/Fetch home notes
[**FetchHomeNotesApiV1XiaohongshuWebV2FetchHomeNotesGet_0**](XiaohongshuWebV2APIApi.md#FetchHomeNotesApiV1XiaohongshuWebV2FetchHomeNotesGet_0) | **Get** /api/v1/xiaohongshu/web_v2/fetch_home_notes | 获取主页笔记/Fetch home notes
[**FetchNoteCommentsApiV1XiaohongshuWebV2FetchNoteCommentsGet**](XiaohongshuWebV2APIApi.md#FetchNoteCommentsApiV1XiaohongshuWebV2FetchNoteCommentsGet) | **Get** /api/v1/xiaohongshu/web_v2/fetch_note_comments | 获取笔记评论/Fetch note comments
[**FetchNoteCommentsApiV1XiaohongshuWebV2FetchNoteCommentsGet_0**](XiaohongshuWebV2APIApi.md#FetchNoteCommentsApiV1XiaohongshuWebV2FetchNoteCommentsGet_0) | **Get** /api/v1/xiaohongshu/web_v2/fetch_note_comments | 获取笔记评论/Fetch note comments
[**FetchSearchNotesApiV1XiaohongshuWebV2FetchSearchNotesGet**](XiaohongshuWebV2APIApi.md#FetchSearchNotesApiV1XiaohongshuWebV2FetchSearchNotesGet) | **Get** /api/v1/xiaohongshu/web_v2/fetch_search_notes | 获取搜索笔记/Fetch search notes
[**FetchSearchNotesApiV1XiaohongshuWebV2FetchSearchNotesGet_0**](XiaohongshuWebV2APIApi.md#FetchSearchNotesApiV1XiaohongshuWebV2FetchSearchNotesGet_0) | **Get** /api/v1/xiaohongshu/web_v2/fetch_search_notes | 获取搜索笔记/Fetch search notes
[**FetchSubCommentsApiV1XiaohongshuWebV2FetchSubCommentsGet**](XiaohongshuWebV2APIApi.md#FetchSubCommentsApiV1XiaohongshuWebV2FetchSubCommentsGet) | **Get** /api/v1/xiaohongshu/web_v2/fetch_sub_comments | 获取子评论/Fetch sub comments
[**FetchSubCommentsApiV1XiaohongshuWebV2FetchSubCommentsGet_0**](XiaohongshuWebV2APIApi.md#FetchSubCommentsApiV1XiaohongshuWebV2FetchSubCommentsGet_0) | **Get** /api/v1/xiaohongshu/web_v2/fetch_sub_comments | 获取子评论/Fetch sub comments
[**FetchUserInfoApiV1XiaohongshuWebV2FetchUserInfoGet**](XiaohongshuWebV2APIApi.md#FetchUserInfoApiV1XiaohongshuWebV2FetchUserInfoGet) | **Get** /api/v1/xiaohongshu/web_v2/fetch_user_info | 获取用户信息/Fetch user info
[**FetchUserInfoApiV1XiaohongshuWebV2FetchUserInfoGet_0**](XiaohongshuWebV2APIApi.md#FetchUserInfoApiV1XiaohongshuWebV2FetchUserInfoGet_0) | **Get** /api/v1/xiaohongshu/web_v2/fetch_user_info | 获取用户信息/Fetch user info



## FetchFeedNotesApiV1XiaohongshuWebV2FetchFeedNotesGet

> ResponseModel FetchFeedNotesApiV1XiaohongshuWebV2FetchFeedNotesGet(ctx, noteId)

获取推荐笔记/Fetch feed notes

# [中文] ### 用途: - 获取推荐笔记 ### 参数: - note_id: 笔记ID，可以从小红书的分享链接中获取 ### 返回: - 推荐笔记  # [English] ### Purpose: - Get feed notes ### Parameters: - note_id: Note ID, can be obtained from the sharing link of Xiaohongshu website. ### Return: - Feed notes  # [示例/Example] note_id = \"66c9cc31000000001f03a4bc\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string**| 笔记ID/Note ID | 

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


## FetchFeedNotesApiV1XiaohongshuWebV2FetchFeedNotesGet_0

> ResponseModel FetchFeedNotesApiV1XiaohongshuWebV2FetchFeedNotesGet_0(ctx, noteId)

获取推荐笔记/Fetch feed notes

# [中文] ### 用途: - 获取推荐笔记 ### 参数: - note_id: 笔记ID，可以从小红书的分享链接中获取 ### 返回: - 推荐笔记  # [English] ### Purpose: - Get feed notes ### Parameters: - note_id: Note ID, can be obtained from the sharing link of Xiaohongshu website. ### Return: - Feed notes  # [示例/Example] note_id = \"66c9cc31000000001f03a4bc\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string**| 笔记ID/Note ID | 

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


## FetchFollowerListApiV1XiaohongshuWebV2FetchFollowerListGet

> ResponseModel FetchFollowerListApiV1XiaohongshuWebV2FetchFollowerListGet(ctx, userId, optional)

获取用户粉丝列表/Fetch follower list

# [中文] ### 用途: - 获取用户粉丝列表 ### 参数: - user_id: 用户ID - cursor: 游标 ### 返回: - 用户粉丝列表  # [English] ### Purpose: - Get follower list ### Parameters: - user_id: User ID - cursor: Cursor ### Return: - Follower list  # [示例/Example] user_id = \"604a28420000000001005211\" cursor = \"\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| 用户ID/User ID | 
 **optional** | ***FetchFollowerListApiV1XiaohongshuWebV2FetchFollowerListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchFollowerListApiV1XiaohongshuWebV2FetchFollowerListGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| 游标/Cursor | [default to ]

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


## FetchFollowerListApiV1XiaohongshuWebV2FetchFollowerListGet_0

> ResponseModel FetchFollowerListApiV1XiaohongshuWebV2FetchFollowerListGet_0(ctx, userId, optional)

获取用户粉丝列表/Fetch follower list

# [中文] ### 用途: - 获取用户粉丝列表 ### 参数: - user_id: 用户ID - cursor: 游标 ### 返回: - 用户粉丝列表  # [English] ### Purpose: - Get follower list ### Parameters: - user_id: User ID - cursor: Cursor ### Return: - Follower list  # [示例/Example] user_id = \"604a28420000000001005211\" cursor = \"\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| 用户ID/User ID | 
 **optional** | ***FetchFollowerListApiV1XiaohongshuWebV2FetchFollowerListGet_2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchFollowerListApiV1XiaohongshuWebV2FetchFollowerListGet_2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| 游标/Cursor | [default to ]

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


## FetchFollowingListApiV1XiaohongshuWebV2FetchFollowingListGet

> ResponseModel FetchFollowingListApiV1XiaohongshuWebV2FetchFollowingListGet(ctx, userId, optional)

获取用户关注列表/Fetch following list

# [中文] ### 用途: - 获取用户关注列表 ### 参数: - user_id: 用户ID - cursor: 游标 ### 返回: - 用户关注列表  # [English] ### Purpose: - Get following list ### Parameters: - user_id: User ID - cursor: Cursor ### Return: - Following list  # [示例/Example] user_id = \"604a28420000000001005211\" cursor = \"\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| 用户ID/User ID | 
 **optional** | ***FetchFollowingListApiV1XiaohongshuWebV2FetchFollowingListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchFollowingListApiV1XiaohongshuWebV2FetchFollowingListGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| 游标/Cursor | [default to ]

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


## FetchFollowingListApiV1XiaohongshuWebV2FetchFollowingListGet_0

> ResponseModel FetchFollowingListApiV1XiaohongshuWebV2FetchFollowingListGet_0(ctx, userId, optional)

获取用户关注列表/Fetch following list

# [中文] ### 用途: - 获取用户关注列表 ### 参数: - user_id: 用户ID - cursor: 游标 ### 返回: - 用户关注列表  # [English] ### Purpose: - Get following list ### Parameters: - user_id: User ID - cursor: Cursor ### Return: - Following list  # [示例/Example] user_id = \"604a28420000000001005211\" cursor = \"\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| 用户ID/User ID | 
 **optional** | ***FetchFollowingListApiV1XiaohongshuWebV2FetchFollowingListGet_3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchFollowingListApiV1XiaohongshuWebV2FetchFollowingListGet_3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| 游标/Cursor | [default to ]

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


## FetchHomeNotesApiV1XiaohongshuWebV2FetchHomeNotesGet

> ResponseModel FetchHomeNotesApiV1XiaohongshuWebV2FetchHomeNotesGet(ctx, userId, optional)

获取主页笔记/Fetch home notes

# [中文] ### 用途: - 获取主页笔记 ### 参数: - user_id: 用户ID - cursor: 游标 ### 返回: - 主页笔记  # [English] ### Purpose: - Get home notes ### Parameters: - user_id: User ID - cursor: Cursor ### Return: - Home notes  # [示例/Example] user_id = \"5f3e0d00000000001f03a4bc\" cursor = \"\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| 用户ID/User ID | 
 **optional** | ***FetchHomeNotesApiV1XiaohongshuWebV2FetchHomeNotesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchHomeNotesApiV1XiaohongshuWebV2FetchHomeNotesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| 游标/Cursor | [default to ]

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


## FetchHomeNotesApiV1XiaohongshuWebV2FetchHomeNotesGet_0

> ResponseModel FetchHomeNotesApiV1XiaohongshuWebV2FetchHomeNotesGet_0(ctx, userId, optional)

获取主页笔记/Fetch home notes

# [中文] ### 用途: - 获取主页笔记 ### 参数: - user_id: 用户ID - cursor: 游标 ### 返回: - 主页笔记  # [English] ### Purpose: - Get home notes ### Parameters: - user_id: User ID - cursor: Cursor ### Return: - Home notes  # [示例/Example] user_id = \"5f3e0d00000000001f03a4bc\" cursor = \"\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| 用户ID/User ID | 
 **optional** | ***FetchHomeNotesApiV1XiaohongshuWebV2FetchHomeNotesGet_4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchHomeNotesApiV1XiaohongshuWebV2FetchHomeNotesGet_4Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| 游标/Cursor | [default to ]

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


## FetchNoteCommentsApiV1XiaohongshuWebV2FetchNoteCommentsGet

> ResponseModel FetchNoteCommentsApiV1XiaohongshuWebV2FetchNoteCommentsGet(ctx, noteId, optional)

获取笔记评论/Fetch note comments

# [中文] ### 用途: - 获取笔记评论 ### 参数: - note_id: 笔记ID - cursor: 游标 ### 返回: - 笔记评论  # [English] ### Purpose: - Get note comments ### Parameters: - note_id: Note ID - cursor: Cursor ### Return: - Note comments  # [示例/Example] note_id = \"651ccaa9000000001f03d7f7\" cursor = \"\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string**| 笔记ID/Note ID | 
 **optional** | ***FetchNoteCommentsApiV1XiaohongshuWebV2FetchNoteCommentsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchNoteCommentsApiV1XiaohongshuWebV2FetchNoteCommentsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| 游标/Cursor | [default to ]

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


## FetchNoteCommentsApiV1XiaohongshuWebV2FetchNoteCommentsGet_0

> ResponseModel FetchNoteCommentsApiV1XiaohongshuWebV2FetchNoteCommentsGet_0(ctx, noteId, optional)

获取笔记评论/Fetch note comments

# [中文] ### 用途: - 获取笔记评论 ### 参数: - note_id: 笔记ID - cursor: 游标 ### 返回: - 笔记评论  # [English] ### Purpose: - Get note comments ### Parameters: - note_id: Note ID - cursor: Cursor ### Return: - Note comments  # [示例/Example] note_id = \"651ccaa9000000001f03d7f7\" cursor = \"\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string**| 笔记ID/Note ID | 
 **optional** | ***FetchNoteCommentsApiV1XiaohongshuWebV2FetchNoteCommentsGet_5Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchNoteCommentsApiV1XiaohongshuWebV2FetchNoteCommentsGet_5Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| 游标/Cursor | [default to ]

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


## FetchSearchNotesApiV1XiaohongshuWebV2FetchSearchNotesGet

> ResponseModel FetchSearchNotesApiV1XiaohongshuWebV2FetchSearchNotesGet(ctx, keywords, optional)

获取搜索笔记/Fetch search notes

# [中文] ### 用途: - 获取搜索笔记 ### 参数: - keywords：搜索关键词 - sort_type：排序方式     - general：综合     - time_descending：最新     - popularity_descending：最热 - note_type: 笔记类型     - 0：全部     - 1：视频     - 2：图文 ### 返回: - 搜索笔记  # [English] ### Purpose: - Get search notes ### Parameters: - keywords: Search keywords - sort_type: Sort type     - general: General     - time_descending: Latest     - popularity_descending: Popular - note_type: Note type     - 0: All     - 1: Video     - 2: Note ### Return: - Search notes  # [示例/Example] keywords = \"口红\" page = 1 sort_type = \"general\" note_type = \"1\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keywords** | **string**| 搜索关键词/Search keywords | 
 **optional** | ***FetchSearchNotesApiV1XiaohongshuWebV2FetchSearchNotesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchSearchNotesApiV1XiaohongshuWebV2FetchSearchNotesGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| 页码/Page number | [default to 1]
 **sortType** | **optional.String**| 排序方式/Sort type | [default to general]
 **noteType** | **optional.String**| 笔记类型/Note type | [default to 0]

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


## FetchSearchNotesApiV1XiaohongshuWebV2FetchSearchNotesGet_0

> ResponseModel FetchSearchNotesApiV1XiaohongshuWebV2FetchSearchNotesGet_0(ctx, keywords, optional)

获取搜索笔记/Fetch search notes

# [中文] ### 用途: - 获取搜索笔记 ### 参数: - keywords：搜索关键词 - sort_type：排序方式     - general：综合     - time_descending：最新     - popularity_descending：最热 - note_type: 笔记类型     - 0：全部     - 1：视频     - 2：图文 ### 返回: - 搜索笔记  # [English] ### Purpose: - Get search notes ### Parameters: - keywords: Search keywords - sort_type: Sort type     - general: General     - time_descending: Latest     - popularity_descending: Popular - note_type: Note type     - 0: All     - 1: Video     - 2: Note ### Return: - Search notes  # [示例/Example] keywords = \"口红\" page = 1 sort_type = \"general\" note_type = \"1\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keywords** | **string**| 搜索关键词/Search keywords | 
 **optional** | ***FetchSearchNotesApiV1XiaohongshuWebV2FetchSearchNotesGet_6Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchSearchNotesApiV1XiaohongshuWebV2FetchSearchNotesGet_6Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| 页码/Page number | [default to 1]
 **sortType** | **optional.String**| 排序方式/Sort type | [default to general]
 **noteType** | **optional.String**| 笔记类型/Note type | [default to 0]

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


## FetchSubCommentsApiV1XiaohongshuWebV2FetchSubCommentsGet

> ResponseModel FetchSubCommentsApiV1XiaohongshuWebV2FetchSubCommentsGet(ctx, noteId, commentId, optional)

获取子评论/Fetch sub comments

# [中文] ### 用途: - 获取子评论 ### 参数: - note_id: 笔记ID - comment_id: 评论ID - cursor: 游标 ### 返回: - 子评论  # [English] ### Purpose: - Get sub comments ### Parameters: - note_id: Note ID - comment_id: Comment ID - cursor: Cursor ### Return: - Sub comments  # [示例/Example] note_id = \"673c894c0000000007033f92\" comment_id = \"673ecdfc000000001503bf8b\" cursor = \"\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string**| 笔记ID/Note ID | 
**commentId** | **string**| 评论ID/Comment ID | 
 **optional** | ***FetchSubCommentsApiV1XiaohongshuWebV2FetchSubCommentsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchSubCommentsApiV1XiaohongshuWebV2FetchSubCommentsGetOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| 游标/Cursor | [default to ]

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


## FetchSubCommentsApiV1XiaohongshuWebV2FetchSubCommentsGet_0

> ResponseModel FetchSubCommentsApiV1XiaohongshuWebV2FetchSubCommentsGet_0(ctx, noteId, commentId, optional)

获取子评论/Fetch sub comments

# [中文] ### 用途: - 获取子评论 ### 参数: - note_id: 笔记ID - comment_id: 评论ID - cursor: 游标 ### 返回: - 子评论  # [English] ### Purpose: - Get sub comments ### Parameters: - note_id: Note ID - comment_id: Comment ID - cursor: Cursor ### Return: - Sub comments  # [示例/Example] note_id = \"673c894c0000000007033f92\" comment_id = \"673ecdfc000000001503bf8b\" cursor = \"\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string**| 笔记ID/Note ID | 
**commentId** | **string**| 评论ID/Comment ID | 
 **optional** | ***FetchSubCommentsApiV1XiaohongshuWebV2FetchSubCommentsGet_7Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FetchSubCommentsApiV1XiaohongshuWebV2FetchSubCommentsGet_7Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| 游标/Cursor | [default to ]

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


## FetchUserInfoApiV1XiaohongshuWebV2FetchUserInfoGet

> ResponseModel FetchUserInfoApiV1XiaohongshuWebV2FetchUserInfoGet(ctx, userId)

获取用户信息/Fetch user info

# [中文] ### 用途: - 获取用户信息 ### 参数: - user_id: 用户ID ### 返回: - 用户信息  # [English] ### Purpose: - Get user info ### Parameters: - user_id: User ID ### Return: - User info  # [示例/Example] user_id = \"5e3a8ee700000000010070c6\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| 用户ID/User ID | 

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


## FetchUserInfoApiV1XiaohongshuWebV2FetchUserInfoGet_0

> ResponseModel FetchUserInfoApiV1XiaohongshuWebV2FetchUserInfoGet_0(ctx, userId)

获取用户信息/Fetch user info

# [中文] ### 用途: - 获取用户信息 ### 参数: - user_id: 用户ID ### 返回: - 用户信息  # [English] ### Purpose: - Get user info ### Parameters: - user_id: User ID ### Return: - User info  # [示例/Example] user_id = \"5e3a8ee700000000010070c6\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| 用户ID/User ID | 

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

