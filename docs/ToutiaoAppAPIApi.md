# \ToutiaoAppAPIApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetArticleInfoApiV1ToutiaoAppGetArticleInfoGet**](ToutiaoAppAPIApi.md#GetArticleInfoApiV1ToutiaoAppGetArticleInfoGet) | **Get** /api/v1/toutiao/app/get_article_info | 获取指定文章的信息/Get information of specified article
[**GetArticleInfoApiV1ToutiaoAppGetArticleInfoGet_0**](ToutiaoAppAPIApi.md#GetArticleInfoApiV1ToutiaoAppGetArticleInfoGet_0) | **Get** /api/v1/toutiao/app/get_article_info | 获取指定文章的信息/Get information of specified article
[**GetCommentsApiV1ToutiaoAppGetCommentsGet**](ToutiaoAppAPIApi.md#GetCommentsApiV1ToutiaoAppGetCommentsGet) | **Get** /api/v1/toutiao/app/get_comments | 获取指定作品的评论/Get comments of specified post
[**GetCommentsApiV1ToutiaoAppGetCommentsGet_0**](ToutiaoAppAPIApi.md#GetCommentsApiV1ToutiaoAppGetCommentsGet_0) | **Get** /api/v1/toutiao/app/get_comments | 获取指定作品的评论/Get comments of specified post
[**GetUserIdApiV1ToutiaoAppGetUserIdGet**](ToutiaoAppAPIApi.md#GetUserIdApiV1ToutiaoAppGetUserIdGet) | **Get** /api/v1/toutiao/app/get_user_id | 从头条用户主页获取用户user_id/Get user_id from user profile
[**GetUserIdApiV1ToutiaoAppGetUserIdGet_0**](ToutiaoAppAPIApi.md#GetUserIdApiV1ToutiaoAppGetUserIdGet_0) | **Get** /api/v1/toutiao/app/get_user_id | 从头条用户主页获取用户user_id/Get user_id from user profile
[**GetUserInfoApiV1ToutiaoAppGetUserInfoGet**](ToutiaoAppAPIApi.md#GetUserInfoApiV1ToutiaoAppGetUserInfoGet) | **Get** /api/v1/toutiao/app/get_user_info | 获取指定用户的信息/Get information of specified user
[**GetUserInfoApiV1ToutiaoAppGetUserInfoGet_0**](ToutiaoAppAPIApi.md#GetUserInfoApiV1ToutiaoAppGetUserInfoGet_0) | **Get** /api/v1/toutiao/app/get_user_info | 获取指定用户的信息/Get information of specified user
[**GetVideoInfoApiV1ToutiaoAppGetVideoInfoGet**](ToutiaoAppAPIApi.md#GetVideoInfoApiV1ToutiaoAppGetVideoInfoGet) | **Get** /api/v1/toutiao/app/get_video_info | 获取指定视频的信息/Get information of specified video
[**GetVideoInfoApiV1ToutiaoAppGetVideoInfoGet_0**](ToutiaoAppAPIApi.md#GetVideoInfoApiV1ToutiaoAppGetVideoInfoGet_0) | **Get** /api/v1/toutiao/app/get_video_info | 获取指定视频的信息/Get information of specified video



## GetArticleInfoApiV1ToutiaoAppGetArticleInfoGet

> ResponseModel GetArticleInfoApiV1ToutiaoAppGetArticleInfoGet(ctx, groupId)

获取指定文章的信息/Get information of specified article

# [中文] ### 用途: - 获取指定文章的信息 ### 参数: - group_id: 作品ID，可以从链接中获取     - 例如: https://www.toutiao.com/article/7450114952884503059/ ### 返回: - 作品信息  # [English] ### Purpose: - Get information of specified post ### Parameters: - group_id: Post ID, can be obtained from the link     - For example: https://www.toutiao.com/article/7450114952884503059/ ### Return: - Post information  # [示例/Example] group_id = \"7450114952884503059\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| 作品ID/Post ID | 

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


## GetArticleInfoApiV1ToutiaoAppGetArticleInfoGet_0

> ResponseModel GetArticleInfoApiV1ToutiaoAppGetArticleInfoGet_0(ctx, groupId)

获取指定文章的信息/Get information of specified article

# [中文] ### 用途: - 获取指定文章的信息 ### 参数: - group_id: 作品ID，可以从链接中获取     - 例如: https://www.toutiao.com/article/7450114952884503059/ ### 返回: - 作品信息  # [English] ### Purpose: - Get information of specified post ### Parameters: - group_id: Post ID, can be obtained from the link     - For example: https://www.toutiao.com/article/7450114952884503059/ ### Return: - Post information  # [示例/Example] group_id = \"7450114952884503059\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| 作品ID/Post ID | 

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


## GetCommentsApiV1ToutiaoAppGetCommentsGet

> ResponseModel GetCommentsApiV1ToutiaoAppGetCommentsGet(ctx, groupId, offset)

获取指定作品的评论/Get comments of specified post

# [中文] ### 用途: - 获取指定作品的评论 ### 参数: - group_id: 作品ID，可以从链接中获取     - 例如: https://www.toutiao.com/i7453372680222523931/ - offset: 偏移量，用于分页，默认为0，然后每次加20 ### 返回: - 评论列表  # [English] ### Purpose: - Get comments of specified post ### Parameters: - group_id: Post ID, can be obtained from the link     - For example: https://www.toutiao.com/i7453372680222523931/ - offset: Offset, used for pagination, default is 0, then add 20 each time ### Return: - Comment list  # [示例/Example] group_id = \"7453372680222523931\" offset = \"0\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| 作品ID/Post ID | 
**offset** | **string**| 偏移量/Offset | 

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


## GetCommentsApiV1ToutiaoAppGetCommentsGet_0

> ResponseModel GetCommentsApiV1ToutiaoAppGetCommentsGet_0(ctx, groupId, offset)

获取指定作品的评论/Get comments of specified post

# [中文] ### 用途: - 获取指定作品的评论 ### 参数: - group_id: 作品ID，可以从链接中获取     - 例如: https://www.toutiao.com/i7453372680222523931/ - offset: 偏移量，用于分页，默认为0，然后每次加20 ### 返回: - 评论列表  # [English] ### Purpose: - Get comments of specified post ### Parameters: - group_id: Post ID, can be obtained from the link     - For example: https://www.toutiao.com/i7453372680222523931/ - offset: Offset, used for pagination, default is 0, then add 20 each time ### Return: - Comment list  # [示例/Example] group_id = \"7453372680222523931\" offset = \"0\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| 作品ID/Post ID | 
**offset** | **string**| 偏移量/Offset | 

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


## GetUserIdApiV1ToutiaoAppGetUserIdGet

> ResponseModel GetUserIdApiV1ToutiaoAppGetUserIdGet(ctx, userProfileUrl)

从头条用户主页获取用户user_id/Get user_id from user profile

# [中文] ### 用途: - 从头条用户主页获取用户user_id ### 参数: - user_profile_url: 用户主页链接     - 例如: https://www.toutiao.com/c/user/token/MS4wLjABAAAAwK6echNksY69R8l2vcZudupfhTItbGSGt-8ineO5UaB4L-djqkYDgB6TkAdMvrmW/ ### 返回: - 用户ID  # [English] ### Purpose: - Get user_id from user profile ### Parameters: - user_profile_url: User profile URL     - For example: https://www.toutiao.com/c/user/token/MS4wLjABAAAAwK6echNksY69R8l2vcZudupfhTItbGSGt-8ineO5UaB4L-djqkYDgB6TkAdMvrmW/ ### Return: - User ID  # [示例/Example] user_profile_url = \"https://www.toutiao.com/c/user/token/MS4wLjABAAAAwK6echNksY69R8l2vcZudupfhTItbGSGt-8ineO5UaB4L-djqkYDgB6TkAdMvrmW/\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userProfileUrl** | **string**| 用户主页链接/User profile URL | 

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


## GetUserIdApiV1ToutiaoAppGetUserIdGet_0

> ResponseModel GetUserIdApiV1ToutiaoAppGetUserIdGet_0(ctx, userProfileUrl)

从头条用户主页获取用户user_id/Get user_id from user profile

# [中文] ### 用途: - 从头条用户主页获取用户user_id ### 参数: - user_profile_url: 用户主页链接     - 例如: https://www.toutiao.com/c/user/token/MS4wLjABAAAAwK6echNksY69R8l2vcZudupfhTItbGSGt-8ineO5UaB4L-djqkYDgB6TkAdMvrmW/ ### 返回: - 用户ID  # [English] ### Purpose: - Get user_id from user profile ### Parameters: - user_profile_url: User profile URL     - For example: https://www.toutiao.com/c/user/token/MS4wLjABAAAAwK6echNksY69R8l2vcZudupfhTItbGSGt-8ineO5UaB4L-djqkYDgB6TkAdMvrmW/ ### Return: - User ID  # [示例/Example] user_profile_url = \"https://www.toutiao.com/c/user/token/MS4wLjABAAAAwK6echNksY69R8l2vcZudupfhTItbGSGt-8ineO5UaB4L-djqkYDgB6TkAdMvrmW/\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userProfileUrl** | **string**| 用户主页链接/User profile URL | 

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


## GetUserInfoApiV1ToutiaoAppGetUserInfoGet

> ResponseModel GetUserInfoApiV1ToutiaoAppGetUserInfoGet(ctx, userId)

获取指定用户的信息/Get information of specified user

# [中文] ### 用途: - 获取指定用户的信息 ### 参数: - user_id: 用户ID，可以从以下接口获取：     - `/api/v1/toutiao/app/get_user_id` ### 返回: - 用户信息  # [English] ### Purpose: - Get information of specified user ### Parameters: - user_id: User ID, can be obtained from the following API:     - `/api/v1/toutiao/app/get_user_id` ### Return: - User information  # [示例/Example] user_id = \"1352838578180211\"

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


## GetUserInfoApiV1ToutiaoAppGetUserInfoGet_0

> ResponseModel GetUserInfoApiV1ToutiaoAppGetUserInfoGet_0(ctx, userId)

获取指定用户的信息/Get information of specified user

# [中文] ### 用途: - 获取指定用户的信息 ### 参数: - user_id: 用户ID，可以从以下接口获取：     - `/api/v1/toutiao/app/get_user_id` ### 返回: - 用户信息  # [English] ### Purpose: - Get information of specified user ### Parameters: - user_id: User ID, can be obtained from the following API:     - `/api/v1/toutiao/app/get_user_id` ### Return: - User information  # [示例/Example] user_id = \"1352838578180211\"

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


## GetVideoInfoApiV1ToutiaoAppGetVideoInfoGet

> ResponseModel GetVideoInfoApiV1ToutiaoAppGetVideoInfoGet(ctx, groupId)

获取指定视频的信息/Get information of specified video

# [中文] ### 用途: - 获取指定视频的信息 ### 参数: - group_id: 作品ID，可以从链接中获取     - 例如: https://www.toutiao.com/video/7431543350882206242/ ### 返回: - 作品信息  # [English] ### Purpose: - Get information of specified video ### Parameters: - group_id: Post ID, can be obtained from the link     - For example: https://www.toutiao.com/video/7431543350882206242/ ### Return: - Post information  # [示例/Example] group_id = \"7431543350882206242\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| 作品ID/Post ID | 

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


## GetVideoInfoApiV1ToutiaoAppGetVideoInfoGet_0

> ResponseModel GetVideoInfoApiV1ToutiaoAppGetVideoInfoGet_0(ctx, groupId)

获取指定视频的信息/Get information of specified video

# [中文] ### 用途: - 获取指定视频的信息 ### 参数: - group_id: 作品ID，可以从链接中获取     - 例如: https://www.toutiao.com/video/7431543350882206242/ ### 返回: - 作品信息  # [English] ### Purpose: - Get information of specified video ### Parameters: - group_id: Post ID, can be obtained from the link     - For example: https://www.toutiao.com/video/7431543350882206242/ ### Return: - Post information  # [示例/Example] group_id = \"7431543350882206242\"

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| 作品ID/Post ID | 

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

