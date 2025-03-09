# \TikTokInteractionAPIApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyForScopeApiV1TiktokInteractionApplyGet**](TikTokInteractionAPIApi.md#ApplyForScopeApiV1TiktokInteractionApplyGet) | **Get** /api/v1/tiktok/interaction/apply | 申请使用TikTok交互API权限（Scope）/Apply for TikTok Interaction API permission (Scope)
[**ApplyForScopeApiV1TiktokInteractionApplyGet_0**](TikTokInteractionAPIApi.md#ApplyForScopeApiV1TiktokInteractionApplyGet_0) | **Get** /api/v1/tiktok/interaction/apply | 申请使用TikTok交互API权限（Scope）/Apply for TikTok Interaction API permission (Scope)
[**CollectApiV1TiktokInteractionCollectPost**](TikTokInteractionAPIApi.md#CollectApiV1TiktokInteractionCollectPost) | **Post** /api/v1/tiktok/interaction/collect | 收藏/Collect
[**CollectApiV1TiktokInteractionCollectPost_0**](TikTokInteractionAPIApi.md#CollectApiV1TiktokInteractionCollectPost_0) | **Post** /api/v1/tiktok/interaction/collect | 收藏/Collect
[**FollowApiV1TiktokInteractionFollowPost**](TikTokInteractionAPIApi.md#FollowApiV1TiktokInteractionFollowPost) | **Post** /api/v1/tiktok/interaction/follow | 关注/Follow
[**FollowApiV1TiktokInteractionFollowPost_0**](TikTokInteractionAPIApi.md#FollowApiV1TiktokInteractionFollowPost_0) | **Post** /api/v1/tiktok/interaction/follow | 关注/Follow
[**ForwardApiV1TiktokInteractionForwardPost**](TikTokInteractionAPIApi.md#ForwardApiV1TiktokInteractionForwardPost) | **Post** /api/v1/tiktok/interaction/forward | 转发/Forward
[**ForwardApiV1TiktokInteractionForwardPost_0**](TikTokInteractionAPIApi.md#ForwardApiV1TiktokInteractionForwardPost_0) | **Post** /api/v1/tiktok/interaction/forward | 转发/Forward
[**LikeApiV1TiktokInteractionLikePost**](TikTokInteractionAPIApi.md#LikeApiV1TiktokInteractionLikePost) | **Post** /api/v1/tiktok/interaction/like | 点赞/Like
[**LikeApiV1TiktokInteractionLikePost_0**](TikTokInteractionAPIApi.md#LikeApiV1TiktokInteractionLikePost_0) | **Post** /api/v1/tiktok/interaction/like | 点赞/Like
[**PostCommentApiV1TiktokInteractionPostCommentPost**](TikTokInteractionAPIApi.md#PostCommentApiV1TiktokInteractionPostCommentPost) | **Post** /api/v1/tiktok/interaction/post_comment | 发送评论/Post comment
[**PostCommentApiV1TiktokInteractionPostCommentPost_0**](TikTokInteractionAPIApi.md#PostCommentApiV1TiktokInteractionPostCommentPost_0) | **Post** /api/v1/tiktok/interaction/post_comment | 发送评论/Post comment
[**ReplyCommentApiV1TiktokInteractionReplyCommentPost**](TikTokInteractionAPIApi.md#ReplyCommentApiV1TiktokInteractionReplyCommentPost) | **Post** /api/v1/tiktok/interaction/reply_comment | 回复评论/Reply to comment
[**ReplyCommentApiV1TiktokInteractionReplyCommentPost_0**](TikTokInteractionAPIApi.md#ReplyCommentApiV1TiktokInteractionReplyCommentPost_0) | **Post** /api/v1/tiktok/interaction/reply_comment | 回复评论/Reply to comment



## ApplyForScopeApiV1TiktokInteractionApplyGet

> ResponseModel ApplyForScopeApiV1TiktokInteractionApplyGet(ctx, apiKey, inviteCode)

申请使用TikTok交互API权限（Scope）/Apply for TikTok Interaction API permission (Scope)

# [通知] - 此接口已经废弃，用户现在无需手动申请调用权限，只需要在用户后台更新API Key的对应权限即可，即API Key对应的的Scope。 # [中文] ### 接口用途: - 申请使用TikTok交互API的接口权限（Scope），请在使用交互类接口之前申请，否则将无法正常请求。 ### 申请流程: - 申请接口权限需要邀请码，如果你没有邀请码，可以在Discord服务器中联系管理员获取。 - Discord服务器链接: [TikHub Discord](https://discord.gg/aMEAS8Xsvz) ### 申请须知: - 此权限仅限于当前提交的API Key，不可跨API Key使用。 - 用户需要使用美国地区注册且有效的的TikTok账号进行登录，否则保证将无法正常请求。 - 用户需要使用美国地区的代理IP进行获取Cookie，否则将保证无法正常请求。 - 用户需要使用美国地区的代理IP进行请求，否则将无法保证正常请求。 ### 用户守则以及责任: - 请不要使用交互类接口对他人造成骚扰，或进行违法违规的操作，否则我们将会停止对你的服务，并且所有责任由你自己承担。 - 请不要将接口权限分享给他人，否则我们将会停止对你的服务。 - 接口请求目前暂时定为每秒5次请求。 ### 返回: - 申请结果以及申请的邀请码，请自行保留邀请码，以便后续使用。  # [Notice] - This interface has been deprecated, users no longer need to apply for permission to call the API, just update the corresponding permission of the API Key in the user background, that is, the Scope corresponding to the API Key. # [English] ### Purpose: - Apply for the interface permission (Scope) of TikTok Interaction API, please apply before using the interactive interface, otherwise the request will not be made normally. ### Application process: - Applying for interface permissions requires an invitation code, if you do not have an invitation code, you can contact the administrator on the Discord server. - Discord server link: [TikHub Discord](https://discord.gg/aMEAS8Xsvz) ### Application notes: - This permission is limited to the API Key submitted, and cannot be used across API Keys. - Users need to log in with a registered and valid TikTok account in the United States, otherwise the request will not be made normally. - Users need to use a proxy IP in the United States to obtain cookies, otherwise the request will not be made normally. - Users need to use a proxy IP in the United States for requests, otherwise the request will not be made normally. ### User guidelines and responsibilities: - Please do not use interactive interfaces to harass others, or engage in illegal or irregular operations, otherwise we will stop providing services to you, and all responsibilities will be borne by you. - Please do not share the interface permission with others, otherwise we will stop providing services to you. - The interface request is currently set to 5 requests per second. ### Return: - Application results and the invitation code applied for, please keep the invitation code for subsequent use.  # [示例/Example] ```python # Python Code invite_code = \"Your_Invite_Code\" ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKey** | **string**|  | 
**inviteCode** | **string**|  | 

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


## ApplyForScopeApiV1TiktokInteractionApplyGet_0

> ResponseModel ApplyForScopeApiV1TiktokInteractionApplyGet_0(ctx, apiKey, inviteCode)

申请使用TikTok交互API权限（Scope）/Apply for TikTok Interaction API permission (Scope)

# [通知] - 此接口已经废弃，用户现在无需手动申请调用权限，只需要在用户后台更新API Key的对应权限即可，即API Key对应的的Scope。 # [中文] ### 接口用途: - 申请使用TikTok交互API的接口权限（Scope），请在使用交互类接口之前申请，否则将无法正常请求。 ### 申请流程: - 申请接口权限需要邀请码，如果你没有邀请码，可以在Discord服务器中联系管理员获取。 - Discord服务器链接: [TikHub Discord](https://discord.gg/aMEAS8Xsvz) ### 申请须知: - 此权限仅限于当前提交的API Key，不可跨API Key使用。 - 用户需要使用美国地区注册且有效的的TikTok账号进行登录，否则保证将无法正常请求。 - 用户需要使用美国地区的代理IP进行获取Cookie，否则将保证无法正常请求。 - 用户需要使用美国地区的代理IP进行请求，否则将无法保证正常请求。 ### 用户守则以及责任: - 请不要使用交互类接口对他人造成骚扰，或进行违法违规的操作，否则我们将会停止对你的服务，并且所有责任由你自己承担。 - 请不要将接口权限分享给他人，否则我们将会停止对你的服务。 - 接口请求目前暂时定为每秒5次请求。 ### 返回: - 申请结果以及申请的邀请码，请自行保留邀请码，以便后续使用。  # [Notice] - This interface has been deprecated, users no longer need to apply for permission to call the API, just update the corresponding permission of the API Key in the user background, that is, the Scope corresponding to the API Key. # [English] ### Purpose: - Apply for the interface permission (Scope) of TikTok Interaction API, please apply before using the interactive interface, otherwise the request will not be made normally. ### Application process: - Applying for interface permissions requires an invitation code, if you do not have an invitation code, you can contact the administrator on the Discord server. - Discord server link: [TikHub Discord](https://discord.gg/aMEAS8Xsvz) ### Application notes: - This permission is limited to the API Key submitted, and cannot be used across API Keys. - Users need to log in with a registered and valid TikTok account in the United States, otherwise the request will not be made normally. - Users need to use a proxy IP in the United States to obtain cookies, otherwise the request will not be made normally. - Users need to use a proxy IP in the United States for requests, otherwise the request will not be made normally. ### User guidelines and responsibilities: - Please do not use interactive interfaces to harass others, or engage in illegal or irregular operations, otherwise we will stop providing services to you, and all responsibilities will be borne by you. - Please do not share the interface permission with others, otherwise we will stop providing services to you. - The interface request is currently set to 5 requests per second. ### Return: - Application results and the invitation code applied for, please keep the invitation code for subsequent use.  # [示例/Example] ```python # Python Code invite_code = \"Your_Invite_Code\" ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKey** | **string**|  | 
**inviteCode** | **string**|  | 

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


## CollectApiV1TiktokInteractionCollectPost

> ResponseModel CollectApiV1TiktokInteractionCollectPost(ctx, optional)

收藏/Collect

# [中文] ### 用途: - 使用用户Cookie收藏指定视频，当前请尽可能使用美国地区的TikTok账号，并且在获取Cookie时请使用美国地区的代理IP。 ### 注意: - 交互类接口都需要使用用户的Cookie，因此请确保你的Cookie是有效的，否则将无法正常请求。 - 交互类的接口可能会导致账号异常或封禁，因此请谨慎使用，建议使用代理IP进行请求。 - 交互类接口的最终结果可能会受到TikTok风控系统的影响，大多数情况跟你所使用的账号有关，比如新注册的账号可能无法关注用户或点赞视频，我们无法处理基于账号的风控问题。 - 请不要使用交互类接口对他人造成骚扰，或进行违法违规的操作，否则我们将会停止对你的服务，并且所有责任由你自己承担。 ### 参数: - aweme_id: 视频id，可以从分享链接中获取，例如：https://www.tiktok.com/@username/video/7419966340443819295 - cookie: 用户Cookie，可以从浏览器中登录自己的TikTok账号，然后复制Cookie信息，提交时请使用URL编码Cookie字符串。 - device_id: 设备id，可选，如果不填写则会自动生成，如果需要自定义设备id，请使用设备信息接口获取设备id。 - iid: 设备安装id，可选，如果不填写则会自动生成，如果需要自定义设备iid，请使用设备信息接口获取设备iid。 - proxy: 代理IP，可选，如果不填写则会使用服务器IP进行请求（不推荐），建议使用代理IP进行请求防止账号异常获被封禁，支持格式如下：     - 代理IP:端口     - 用户名:密码@代理IP:端口 ### 返回: - 点赞结果以及评论数据和设备信息，请自行保留设备信息，如请求时使用的`device_id`以及`iid`，以便后续调用接口时使用，频繁更换设备信息可能会导致账号异常或封禁。  # [English] ### Purpose: - Collect a specified video using user cookies, please try to use TikTok accounts in the United States as much as possible, and use proxy IPs in the United States when obtaining cookies. ### Note: - Interactive interfaces all need to use the user's Cookie, so please make sure that your Cookie is valid, otherwise the request will not be made normally. - Interactive interfaces may cause account exceptions or bans, so please use them with caution, and it is recommended to use proxy IPs for requests. - The final result of the interactive interface may be affected by the TikTok risk control system, and in most cases it is related to the account you are using, for example, a newly registered account may not be able to follow users or like videos, and we cannot handle risk control issues based on accounts. - Please do not use interactive interfaces to harass others, or engage in illegal or irregular operations, otherwise we will stop providing services to you, and all responsibilities will be borne by you. ### Parameters: - aweme_id: Video id, which can be obtained from the sharing link, for example: https://www.tiktok.com/@username/video/7419966340443819295 - cookie: User Cookie, you can log in to your TikTok account in the browser and then copy the Cookie information, please use URL-encoded Cookie string when submitting. - device_id: Device id, optional, if not filled in, it will be automatically generated, if you need to customize the device id, please use the device information interface to get the device id. - iid: Device install id, optional, if not filled in, it will be automatically generated, if you need to customize the device iid, please use the device information interface to get the device iid. - proxy: Proxy IP, optional, if not filled in, the server IP will be used for the request (not recommended), it is recommended to use a proxy IP for the request to prevent account exceptions or bans, support the following formats:     - Proxy IP:Port     - Username:Password@Proxy IP:Port ### Return: - Like results, comment data and device information, please keep the device information, such as the `device_id` and `iid` used when requesting, for subsequent calls to the interface, frequent replacement of device information may cause account exceptions or bans.  # [示例/Example] ```python # Python Code cookie = urllib.parse.quote(\"Your_Cookie_From_Browser\") payload = {     \"aweme_id\": \"7419966340443819295\",     \"cookie\": cookie,     \"proxy\": \"\", } ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectApiV1TiktokInteractionCollectPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CollectApiV1TiktokInteractionCollectPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collectRequest** | [**optional.Interface of CollectRequest**](CollectRequest.md)|  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CollectApiV1TiktokInteractionCollectPost_0

> ResponseModel CollectApiV1TiktokInteractionCollectPost_0(ctx, optional)

收藏/Collect

# [中文] ### 用途: - 使用用户Cookie收藏指定视频，当前请尽可能使用美国地区的TikTok账号，并且在获取Cookie时请使用美国地区的代理IP。 ### 注意: - 交互类接口都需要使用用户的Cookie，因此请确保你的Cookie是有效的，否则将无法正常请求。 - 交互类的接口可能会导致账号异常或封禁，因此请谨慎使用，建议使用代理IP进行请求。 - 交互类接口的最终结果可能会受到TikTok风控系统的影响，大多数情况跟你所使用的账号有关，比如新注册的账号可能无法关注用户或点赞视频，我们无法处理基于账号的风控问题。 - 请不要使用交互类接口对他人造成骚扰，或进行违法违规的操作，否则我们将会停止对你的服务，并且所有责任由你自己承担。 ### 参数: - aweme_id: 视频id，可以从分享链接中获取，例如：https://www.tiktok.com/@username/video/7419966340443819295 - cookie: 用户Cookie，可以从浏览器中登录自己的TikTok账号，然后复制Cookie信息，提交时请使用URL编码Cookie字符串。 - device_id: 设备id，可选，如果不填写则会自动生成，如果需要自定义设备id，请使用设备信息接口获取设备id。 - iid: 设备安装id，可选，如果不填写则会自动生成，如果需要自定义设备iid，请使用设备信息接口获取设备iid。 - proxy: 代理IP，可选，如果不填写则会使用服务器IP进行请求（不推荐），建议使用代理IP进行请求防止账号异常获被封禁，支持格式如下：     - 代理IP:端口     - 用户名:密码@代理IP:端口 ### 返回: - 点赞结果以及评论数据和设备信息，请自行保留设备信息，如请求时使用的`device_id`以及`iid`，以便后续调用接口时使用，频繁更换设备信息可能会导致账号异常或封禁。  # [English] ### Purpose: - Collect a specified video using user cookies, please try to use TikTok accounts in the United States as much as possible, and use proxy IPs in the United States when obtaining cookies. ### Note: - Interactive interfaces all need to use the user's Cookie, so please make sure that your Cookie is valid, otherwise the request will not be made normally. - Interactive interfaces may cause account exceptions or bans, so please use them with caution, and it is recommended to use proxy IPs for requests. - The final result of the interactive interface may be affected by the TikTok risk control system, and in most cases it is related to the account you are using, for example, a newly registered account may not be able to follow users or like videos, and we cannot handle risk control issues based on accounts. - Please do not use interactive interfaces to harass others, or engage in illegal or irregular operations, otherwise we will stop providing services to you, and all responsibilities will be borne by you. ### Parameters: - aweme_id: Video id, which can be obtained from the sharing link, for example: https://www.tiktok.com/@username/video/7419966340443819295 - cookie: User Cookie, you can log in to your TikTok account in the browser and then copy the Cookie information, please use URL-encoded Cookie string when submitting. - device_id: Device id, optional, if not filled in, it will be automatically generated, if you need to customize the device id, please use the device information interface to get the device id. - iid: Device install id, optional, if not filled in, it will be automatically generated, if you need to customize the device iid, please use the device information interface to get the device iid. - proxy: Proxy IP, optional, if not filled in, the server IP will be used for the request (not recommended), it is recommended to use a proxy IP for the request to prevent account exceptions or bans, support the following formats:     - Proxy IP:Port     - Username:Password@Proxy IP:Port ### Return: - Like results, comment data and device information, please keep the device information, such as the `device_id` and `iid` used when requesting, for subsequent calls to the interface, frequent replacement of device information may cause account exceptions or bans.  # [示例/Example] ```python # Python Code cookie = urllib.parse.quote(\"Your_Cookie_From_Browser\") payload = {     \"aweme_id\": \"7419966340443819295\",     \"cookie\": cookie,     \"proxy\": \"\", } ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CollectApiV1TiktokInteractionCollectPost_2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CollectApiV1TiktokInteractionCollectPost_2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collectRequest** | [**optional.Interface of CollectRequest**](CollectRequest.md)|  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FollowApiV1TiktokInteractionFollowPost

> ResponseModel FollowApiV1TiktokInteractionFollowPost(ctx, optional)

关注/Follow

# [中文] ### 用途: - 使用用户Cookie关注指定用户，当前请尽可能使用美国地区的TikTok账号，并且在获取Cookie时请使用美国地区的代理IP。 ### 注意: - 交互类接口都需要使用用户的Cookie，因此请确保你的Cookie是有效的，否则将无法正常请求。 - 交互类的接口可能会导致账号异常或封禁，因此请谨慎使用，建议使用代理IP进行请求。 - 交互类接口的最终结果可能会受到TikTok风控系统的影响，大多数情况跟你所使用的账号有关，比如新注册的账号可能无法关注用户或点赞视频，我们无法处理基于账号的风控问题。 - 请不要使用交互类接口对他人造成骚扰，或进行违法违规的操作，否则我们将会停止对你的服务，并且所有责任由你自己承担。 ### 参数: - user_id: 用户id，可以从接口`/api/v1/tiktok/app/v3/handler_user_profile`获取。 - sec_user_id: 用户sec_id，可以从分接口`/api/v1/tiktok/web/get_sec_user_id`获取。 - cookie: 用户Cookie，可以从浏览器中登录自己的TikTok账号，然后复制Cookie信息，提交时请使用URL编码Cookie字符串。 - device_id: 设备id，可选，如果不填写则会自动生成，如果需要自定义设备id，请使用设备信息接口获取设备id。 - iid: 设备安装id，可选，如果不填写则会自动生成，如果需要自定义设备iid，请使用设备信息接口获取设备iid。 - proxy: 代理IP，可选，如果不填写则会使用服务器IP进行请求（不推荐），建议使用代理IP进行请求防止账号异常获被封禁，支持格式如下：     - 代理IP:端口     - 用户名:密码@代理IP:端口 ### 返回: - 关注结果以及评论数据和设备信息，请自行保留设备信息，如请求时使用的`device_id`以及`iid`，以便后续调用接口时使用，频繁更换设备信息可能会导致账号异常或封禁。  # [English] ### Purpose: - Follow a specified user using user cookies, please try to use TikTok accounts in the United States as much as possible, and use proxy IPs in the United States when obtaining cookies. ### Note: - Interactive interfaces all need to use the user's Cookie, so please make sure that your Cookie is valid, otherwise the request will not be made normally. - Interactive interfaces may cause account exceptions or bans, so please use them with caution, and it is recommended to use proxy IPs for requests. - The final result of the interactive interface may be affected by the TikTok risk control system, and in most cases it is related to the account you are using, for example, a newly registered account may not be able to follow users or like videos, and we cannot handle risk control issues based on accounts. - Please do not use interactive interfaces to harass others, or engage in illegal or irregular operations, otherwise we will stop providing services to you, and all responsibilities will be borne by you. ### Parameters: - user_id: User id, which can be obtained from the sub-interface `/api/v1/tiktok/app/v3/handler_user_profile`. - sec_user_id: User sec_id, which can be obtained from the sub-interface `/api/v1/tiktok/web/get_sec_user_id`. - cookie: User Cookie, you can log in to your TikTok account in the browser and then copy the Cookie information, please use URL-encoded Cookie string when submitting. - device_id: Device id, optional, if not filled in, it will be automatically generated, if you need to customize the device id, please use the device information interface to get the device id. - iid: Device install id, optional, if not filled in, it will be automatically generated, if you need to customize the device iid, please use the device information interface to get the device iid. - proxy: Proxy IP, optional, if not filled in, the server IP will be used for the request (not recommended), it is recommended to use a proxy IP for the request to prevent account exceptions or bans, support the following formats:     - Proxy IP:Port     - Username:Password@Proxy IP:Port ### Return: - Follow results, comment data and device information, please keep the device information, such as the `device_id` and `iid` used when requesting, for subsequent calls to the interface, frequent replacement of device information may cause account exceptions or bans.  # [示例/Example] ```python # Python Code cookie = urllib.parse.quote(\"Your_Cookie_From_Browser\") payload = {     \"user_id\": \"6881290705605477381\",     \"sec_user_id\": \"MS4wLjABAAAAqB08cUbXaDWqbD6MCga2RbGTuhfO2EsHayBYx08NDrN7IE3jQuRDNNN6YwyfH6_6\",     \"cookie\": cookie,     \"proxy\": \"\", } ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FollowApiV1TiktokInteractionFollowPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FollowApiV1TiktokInteractionFollowPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **followRequest** | [**optional.Interface of FollowRequest**](FollowRequest.md)|  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FollowApiV1TiktokInteractionFollowPost_0

> ResponseModel FollowApiV1TiktokInteractionFollowPost_0(ctx, optional)

关注/Follow

# [中文] ### 用途: - 使用用户Cookie关注指定用户，当前请尽可能使用美国地区的TikTok账号，并且在获取Cookie时请使用美国地区的代理IP。 ### 注意: - 交互类接口都需要使用用户的Cookie，因此请确保你的Cookie是有效的，否则将无法正常请求。 - 交互类的接口可能会导致账号异常或封禁，因此请谨慎使用，建议使用代理IP进行请求。 - 交互类接口的最终结果可能会受到TikTok风控系统的影响，大多数情况跟你所使用的账号有关，比如新注册的账号可能无法关注用户或点赞视频，我们无法处理基于账号的风控问题。 - 请不要使用交互类接口对他人造成骚扰，或进行违法违规的操作，否则我们将会停止对你的服务，并且所有责任由你自己承担。 ### 参数: - user_id: 用户id，可以从接口`/api/v1/tiktok/app/v3/handler_user_profile`获取。 - sec_user_id: 用户sec_id，可以从分接口`/api/v1/tiktok/web/get_sec_user_id`获取。 - cookie: 用户Cookie，可以从浏览器中登录自己的TikTok账号，然后复制Cookie信息，提交时请使用URL编码Cookie字符串。 - device_id: 设备id，可选，如果不填写则会自动生成，如果需要自定义设备id，请使用设备信息接口获取设备id。 - iid: 设备安装id，可选，如果不填写则会自动生成，如果需要自定义设备iid，请使用设备信息接口获取设备iid。 - proxy: 代理IP，可选，如果不填写则会使用服务器IP进行请求（不推荐），建议使用代理IP进行请求防止账号异常获被封禁，支持格式如下：     - 代理IP:端口     - 用户名:密码@代理IP:端口 ### 返回: - 关注结果以及评论数据和设备信息，请自行保留设备信息，如请求时使用的`device_id`以及`iid`，以便后续调用接口时使用，频繁更换设备信息可能会导致账号异常或封禁。  # [English] ### Purpose: - Follow a specified user using user cookies, please try to use TikTok accounts in the United States as much as possible, and use proxy IPs in the United States when obtaining cookies. ### Note: - Interactive interfaces all need to use the user's Cookie, so please make sure that your Cookie is valid, otherwise the request will not be made normally. - Interactive interfaces may cause account exceptions or bans, so please use them with caution, and it is recommended to use proxy IPs for requests. - The final result of the interactive interface may be affected by the TikTok risk control system, and in most cases it is related to the account you are using, for example, a newly registered account may not be able to follow users or like videos, and we cannot handle risk control issues based on accounts. - Please do not use interactive interfaces to harass others, or engage in illegal or irregular operations, otherwise we will stop providing services to you, and all responsibilities will be borne by you. ### Parameters: - user_id: User id, which can be obtained from the sub-interface `/api/v1/tiktok/app/v3/handler_user_profile`. - sec_user_id: User sec_id, which can be obtained from the sub-interface `/api/v1/tiktok/web/get_sec_user_id`. - cookie: User Cookie, you can log in to your TikTok account in the browser and then copy the Cookie information, please use URL-encoded Cookie string when submitting. - device_id: Device id, optional, if not filled in, it will be automatically generated, if you need to customize the device id, please use the device information interface to get the device id. - iid: Device install id, optional, if not filled in, it will be automatically generated, if you need to customize the device iid, please use the device information interface to get the device iid. - proxy: Proxy IP, optional, if not filled in, the server IP will be used for the request (not recommended), it is recommended to use a proxy IP for the request to prevent account exceptions or bans, support the following formats:     - Proxy IP:Port     - Username:Password@Proxy IP:Port ### Return: - Follow results, comment data and device information, please keep the device information, such as the `device_id` and `iid` used when requesting, for subsequent calls to the interface, frequent replacement of device information may cause account exceptions or bans.  # [示例/Example] ```python # Python Code cookie = urllib.parse.quote(\"Your_Cookie_From_Browser\") payload = {     \"user_id\": \"6881290705605477381\",     \"sec_user_id\": \"MS4wLjABAAAAqB08cUbXaDWqbD6MCga2RbGTuhfO2EsHayBYx08NDrN7IE3jQuRDNNN6YwyfH6_6\",     \"cookie\": cookie,     \"proxy\": \"\", } ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FollowApiV1TiktokInteractionFollowPost_3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FollowApiV1TiktokInteractionFollowPost_3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **followRequest** | [**optional.Interface of FollowRequest**](FollowRequest.md)|  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForwardApiV1TiktokInteractionForwardPost

> ResponseModel ForwardApiV1TiktokInteractionForwardPost(ctx, optional)

转发/Forward

# [中文] ### 用途: - 使用用户Cookie转发指定作品，当前请尽可能使用美国地区的TikTok账号，并且在获取Cookie时请使用美国地区的代理IP。 ### 注意: - 交互类接口都需要使用用户的Cookie，因此请确保你的Cookie是有效的，否则将无法正常请求。 - 交互类的接口可能会导致账号异常或封禁，因此请谨慎使用，建议使用代理IP进行请求。 - 交互类接口的最终结果可能会受到TikTok风控系统的影响，大多数情况跟你所使用的账号有关，比如新注册的账号可能无法关注用户或点赞视频，我们无法处理基于账号的风控问题。 - 请不要使用交互类接口对他人造成骚扰，或进行违法违规的操作，否则我们将会停止对你的服务，并且所有责任由你自己承担。 ### 参数: - aweme_id: 视频id，可以从分享链接中获取，例如：https://www.tiktok.com/@username/video/7419966340443819295 - cookie: 用户Cookie，可以从浏览器中登录自己的TikTok账号，然后复制Cookie信息，提交时请使用URL编码Cookie字符串。 - device_id: 设备id，可选，如果不填写则会自动生成，如果需要自定义设备id，请使用设备信息接口获取设备id。 - iid: 设备安装id，可选，如果不填写则会自动生成，如果需要自定义设备iid，请使用设备信息接口获取设备iid。 - proxy: 代理IP，可选，如果不填写则会使用服务器IP进行请求（不推荐），建议使用代理IP进行请求防止账号异常获被封禁，支持格式如下：     - 代理IP:端口     - 用户名:密码@代理IP:端口 ### 返回: - 关注结果以及评论数据和设备信息，请自行保留设备信息，如请求时使用的`device_id`以及`iid`，以便后续调用接口时使用，频繁更换设备信息可能会导致账号异常或封禁。  # [English] ### Purpose: - Forward a specified post using user cookies, please try to use TikTok accounts in the United States as much as possible, and use proxy IPs in the United States when obtaining cookies. ### Note: - Interactive interfaces all need to use the user's Cookie, so please make sure that your Cookie is valid, otherwise the request will not be made normally. - Interactive interfaces may cause account exceptions or bans, so please use them with caution, and it is recommended to use proxy IPs for requests. - The final result of the interactive interface may be affected by the TikTok risk control system, and in most cases it is related to the account you are using, for example, a newly registered account may not be able to follow users or like videos, and we cannot handle risk control issues based on accounts. - Please do not use interactive interfaces to harass others, or engage in illegal or irregular operations, otherwise we will stop providing services to you, and all responsibilities will be borne by you. ### Parameters: - aweme_id: Video id, which can be obtained from the sharing link, for example: https://www.tiktok.com/@username/video/7419966340443819295 - sec_user_id: User sec_id, which can be obtained from the sub-interface `/api/v1/tiktok/web/get_sec_user_id`. - cookie: User Cookie, you can log in to your TikTok account in the browser and then copy the Cookie information, please use URL-encoded Cookie string when submitting. - device_id: Device id, optional, if not filled in, it will be automatically generated, if you need to customize the device id, please use the device information interface to get the device id. - iid: Device install id, optional, if not filled in, it will be automatically generated, if you need to customize the device iid, please use the device information interface to get the device iid. - proxy: Proxy IP, optional, if not filled in, the server IP will be used for the request (not recommended), it is recommended to use a proxy IP for the request to prevent account exceptions or bans, support the following formats:     - Proxy IP:Port     - Username:Password@Proxy IP:Port ### Return: - Follow results, comment data and device information, please keep the device information, such as the `device_id` and `iid` used when requesting, for subsequent calls to the interface, frequent replacement of device information may cause account exceptions or bans.  # [示例/Example] ```python # Python Code cookie = urllib.parse.quote(\"Your_Cookie_From_Browser\") payload = {     \"user_id\": \"6881290705605477381\",     \"sec_user_id\": \"MS4wLjABAAAAqB08cUbXaDWqbD6MCga2RbGTuhfO2EsHayBYx08NDrN7IE3jQuRDNNN6YwyfH6_6\",     \"cookie\": cookie,     \"proxy\": \"\", } ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ForwardApiV1TiktokInteractionForwardPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ForwardApiV1TiktokInteractionForwardPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forwardRequest** | [**optional.Interface of ForwardRequest**](ForwardRequest.md)|  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForwardApiV1TiktokInteractionForwardPost_0

> ResponseModel ForwardApiV1TiktokInteractionForwardPost_0(ctx, optional)

转发/Forward

# [中文] ### 用途: - 使用用户Cookie转发指定作品，当前请尽可能使用美国地区的TikTok账号，并且在获取Cookie时请使用美国地区的代理IP。 ### 注意: - 交互类接口都需要使用用户的Cookie，因此请确保你的Cookie是有效的，否则将无法正常请求。 - 交互类的接口可能会导致账号异常或封禁，因此请谨慎使用，建议使用代理IP进行请求。 - 交互类接口的最终结果可能会受到TikTok风控系统的影响，大多数情况跟你所使用的账号有关，比如新注册的账号可能无法关注用户或点赞视频，我们无法处理基于账号的风控问题。 - 请不要使用交互类接口对他人造成骚扰，或进行违法违规的操作，否则我们将会停止对你的服务，并且所有责任由你自己承担。 ### 参数: - aweme_id: 视频id，可以从分享链接中获取，例如：https://www.tiktok.com/@username/video/7419966340443819295 - cookie: 用户Cookie，可以从浏览器中登录自己的TikTok账号，然后复制Cookie信息，提交时请使用URL编码Cookie字符串。 - device_id: 设备id，可选，如果不填写则会自动生成，如果需要自定义设备id，请使用设备信息接口获取设备id。 - iid: 设备安装id，可选，如果不填写则会自动生成，如果需要自定义设备iid，请使用设备信息接口获取设备iid。 - proxy: 代理IP，可选，如果不填写则会使用服务器IP进行请求（不推荐），建议使用代理IP进行请求防止账号异常获被封禁，支持格式如下：     - 代理IP:端口     - 用户名:密码@代理IP:端口 ### 返回: - 关注结果以及评论数据和设备信息，请自行保留设备信息，如请求时使用的`device_id`以及`iid`，以便后续调用接口时使用，频繁更换设备信息可能会导致账号异常或封禁。  # [English] ### Purpose: - Forward a specified post using user cookies, please try to use TikTok accounts in the United States as much as possible, and use proxy IPs in the United States when obtaining cookies. ### Note: - Interactive interfaces all need to use the user's Cookie, so please make sure that your Cookie is valid, otherwise the request will not be made normally. - Interactive interfaces may cause account exceptions or bans, so please use them with caution, and it is recommended to use proxy IPs for requests. - The final result of the interactive interface may be affected by the TikTok risk control system, and in most cases it is related to the account you are using, for example, a newly registered account may not be able to follow users or like videos, and we cannot handle risk control issues based on accounts. - Please do not use interactive interfaces to harass others, or engage in illegal or irregular operations, otherwise we will stop providing services to you, and all responsibilities will be borne by you. ### Parameters: - aweme_id: Video id, which can be obtained from the sharing link, for example: https://www.tiktok.com/@username/video/7419966340443819295 - sec_user_id: User sec_id, which can be obtained from the sub-interface `/api/v1/tiktok/web/get_sec_user_id`. - cookie: User Cookie, you can log in to your TikTok account in the browser and then copy the Cookie information, please use URL-encoded Cookie string when submitting. - device_id: Device id, optional, if not filled in, it will be automatically generated, if you need to customize the device id, please use the device information interface to get the device id. - iid: Device install id, optional, if not filled in, it will be automatically generated, if you need to customize the device iid, please use the device information interface to get the device iid. - proxy: Proxy IP, optional, if not filled in, the server IP will be used for the request (not recommended), it is recommended to use a proxy IP for the request to prevent account exceptions or bans, support the following formats:     - Proxy IP:Port     - Username:Password@Proxy IP:Port ### Return: - Follow results, comment data and device information, please keep the device information, such as the `device_id` and `iid` used when requesting, for subsequent calls to the interface, frequent replacement of device information may cause account exceptions or bans.  # [示例/Example] ```python # Python Code cookie = urllib.parse.quote(\"Your_Cookie_From_Browser\") payload = {     \"user_id\": \"6881290705605477381\",     \"sec_user_id\": \"MS4wLjABAAAAqB08cUbXaDWqbD6MCga2RbGTuhfO2EsHayBYx08NDrN7IE3jQuRDNNN6YwyfH6_6\",     \"cookie\": cookie,     \"proxy\": \"\", } ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ForwardApiV1TiktokInteractionForwardPost_4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ForwardApiV1TiktokInteractionForwardPost_4Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forwardRequest** | [**optional.Interface of ForwardRequest**](ForwardRequest.md)|  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LikeApiV1TiktokInteractionLikePost

> ResponseModel LikeApiV1TiktokInteractionLikePost(ctx, optional)

点赞/Like

# [中文] ### 用途: - 使用用户Cookie点赞指定视频，当前请尽可能使用美国地区的TikTok账号，并且在获取Cookie时请使用美国地区的代理IP。 ### 注意: - 交互类接口都需要使用用户的Cookie，因此请确保你的Cookie是有效的，否则将无法正常请求。 - 交互类的接口可能会导致账号异常或封禁，因此请谨慎使用，建议使用代理IP进行请求。 - 交互类接口的最终结果可能会受到TikTok风控系统的影响，大多数情况跟你所使用的账号有关，比如新注册的账号可能无法关注用户或点赞视频，我们无法处理基于账号的风控问题。 - 请不要使用交互类接口对他人造成骚扰，或进行违法违规的操作，否则我们将会停止对你的服务，并且所有责任由你自己承担。 ### 参数: - aweme_id: 视频id，可以从分享链接中获取，例如：https://www.tiktok.com/@username/video/7419966340443819295 - cookie: 用户Cookie，可以从浏览器中登录自己的TikTok账号，然后复制Cookie信息，提交时请使用URL编码Cookie字符串。 - device_id: 设备id，可选，如果不填写则会自动生成，如果需要自定义设备id，请使用设备信息接口获取设备id。 - iid: 设备安装id，可选，如果不填写则会自动生成，如果需要自定义设备iid，请使用设备信息接口获取设备iid。 - proxy: 代理IP，可选，如果不填写则会使用服务器IP进行请求（不推荐），建议使用代理IP进行请求防止账号异常获被封禁，支持格式如下：     - 代理IP:端口     - 用户名:密码@代理IP:端口 ### 返回: - 点赞结果以及评论数据和设备信息，请自行保留设备信息，如请求时使用的`device_id`以及`iid`，以便后续调用接口时使用，频繁更换设备信息可能会导致账号异常或封禁。  # [English] ### Purpose: - Like a specified video using user cookies, please try to use TikTok accounts in the United States as much as possible, and use proxy IPs in the United States when obtaining cookies. ### Note: - Interactive interfaces all need to use the user's Cookie, so please make sure that your Cookie is valid, otherwise the request will not be made normally. - Interactive interfaces may cause account exceptions or bans, so please use them with caution, and it is recommended to use proxy IPs for requests. - The final result of the interactive interface may be affected by the TikTok risk control system, and in most cases it is related to the account you are using, for example, a newly registered account may not be able to follow users or like videos, and we cannot handle risk control issues based on accounts. - Please do not use interactive interfaces to harass others, or engage in illegal or irregular operations, otherwise we will stop providing services to you, and all responsibilities will be borne by you. ### Parameters: - aweme_id: Video id, which can be obtained from the sharing link, for example: https://www.tiktok.com/@username/video/7419966340443819295 - cookie: User Cookie, you can log in to your TikTok account in the browser and then copy the Cookie information, please use URL-encoded Cookie string when submitting. - device_id: Device id, optional, if not filled in, it will be automatically generated, if you need to customize the device id, please use the device information interface to get the device id. - iid: Device install id, optional, if not filled in, it will be automatically generated, if you need to customize the device iid, please use the device information interface to get the device iid. - proxy: Proxy IP, optional, if not filled in, the server IP will be used for the request (not recommended), it is recommended to use a proxy IP for the request to prevent account exceptions or bans, support the following formats:     - Proxy IP:Port     - Username:Password@Proxy IP:Port ### Return: - Like results, comment data and device information, please keep the device information, such as the `device_id` and `iid` used when requesting, for subsequent calls to the interface, frequent replacement of device information may cause account exceptions or bans.  # [示例/Example] ```python # Python Code cookie = urllib.parse.quote(\"Your_Cookie_From_Browser\") payload = {     \"aweme_id\": \"7419966340443819295\",     \"cookie\": cookie,     \"proxy\": \"\", } ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LikeApiV1TiktokInteractionLikePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LikeApiV1TiktokInteractionLikePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **likeRequest** | [**optional.Interface of LikeRequest**](LikeRequest.md)|  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LikeApiV1TiktokInteractionLikePost_0

> ResponseModel LikeApiV1TiktokInteractionLikePost_0(ctx, optional)

点赞/Like

# [中文] ### 用途: - 使用用户Cookie点赞指定视频，当前请尽可能使用美国地区的TikTok账号，并且在获取Cookie时请使用美国地区的代理IP。 ### 注意: - 交互类接口都需要使用用户的Cookie，因此请确保你的Cookie是有效的，否则将无法正常请求。 - 交互类的接口可能会导致账号异常或封禁，因此请谨慎使用，建议使用代理IP进行请求。 - 交互类接口的最终结果可能会受到TikTok风控系统的影响，大多数情况跟你所使用的账号有关，比如新注册的账号可能无法关注用户或点赞视频，我们无法处理基于账号的风控问题。 - 请不要使用交互类接口对他人造成骚扰，或进行违法违规的操作，否则我们将会停止对你的服务，并且所有责任由你自己承担。 ### 参数: - aweme_id: 视频id，可以从分享链接中获取，例如：https://www.tiktok.com/@username/video/7419966340443819295 - cookie: 用户Cookie，可以从浏览器中登录自己的TikTok账号，然后复制Cookie信息，提交时请使用URL编码Cookie字符串。 - device_id: 设备id，可选，如果不填写则会自动生成，如果需要自定义设备id，请使用设备信息接口获取设备id。 - iid: 设备安装id，可选，如果不填写则会自动生成，如果需要自定义设备iid，请使用设备信息接口获取设备iid。 - proxy: 代理IP，可选，如果不填写则会使用服务器IP进行请求（不推荐），建议使用代理IP进行请求防止账号异常获被封禁，支持格式如下：     - 代理IP:端口     - 用户名:密码@代理IP:端口 ### 返回: - 点赞结果以及评论数据和设备信息，请自行保留设备信息，如请求时使用的`device_id`以及`iid`，以便后续调用接口时使用，频繁更换设备信息可能会导致账号异常或封禁。  # [English] ### Purpose: - Like a specified video using user cookies, please try to use TikTok accounts in the United States as much as possible, and use proxy IPs in the United States when obtaining cookies. ### Note: - Interactive interfaces all need to use the user's Cookie, so please make sure that your Cookie is valid, otherwise the request will not be made normally. - Interactive interfaces may cause account exceptions or bans, so please use them with caution, and it is recommended to use proxy IPs for requests. - The final result of the interactive interface may be affected by the TikTok risk control system, and in most cases it is related to the account you are using, for example, a newly registered account may not be able to follow users or like videos, and we cannot handle risk control issues based on accounts. - Please do not use interactive interfaces to harass others, or engage in illegal or irregular operations, otherwise we will stop providing services to you, and all responsibilities will be borne by you. ### Parameters: - aweme_id: Video id, which can be obtained from the sharing link, for example: https://www.tiktok.com/@username/video/7419966340443819295 - cookie: User Cookie, you can log in to your TikTok account in the browser and then copy the Cookie information, please use URL-encoded Cookie string when submitting. - device_id: Device id, optional, if not filled in, it will be automatically generated, if you need to customize the device id, please use the device information interface to get the device id. - iid: Device install id, optional, if not filled in, it will be automatically generated, if you need to customize the device iid, please use the device information interface to get the device iid. - proxy: Proxy IP, optional, if not filled in, the server IP will be used for the request (not recommended), it is recommended to use a proxy IP for the request to prevent account exceptions or bans, support the following formats:     - Proxy IP:Port     - Username:Password@Proxy IP:Port ### Return: - Like results, comment data and device information, please keep the device information, such as the `device_id` and `iid` used when requesting, for subsequent calls to the interface, frequent replacement of device information may cause account exceptions or bans.  # [示例/Example] ```python # Python Code cookie = urllib.parse.quote(\"Your_Cookie_From_Browser\") payload = {     \"aweme_id\": \"7419966340443819295\",     \"cookie\": cookie,     \"proxy\": \"\", } ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LikeApiV1TiktokInteractionLikePost_5Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LikeApiV1TiktokInteractionLikePost_5Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **likeRequest** | [**optional.Interface of LikeRequest**](LikeRequest.md)|  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommentApiV1TiktokInteractionPostCommentPost

> ResponseModel PostCommentApiV1TiktokInteractionPostCommentPost(ctx, optional)

发送评论/Post comment

# [中文] ### 用途: - 使用用户Cookie发送评论到指定视频，当前请尽可能使用美国地区的TikTok账号，并且在获取Cookie时请使用美国地区的代理IP。 ### 注意: - 交互类接口都需要使用用户的Cookie，因此请确保你的Cookie是有效的，否则将无法正常请求。 - 交互类的接口可能会导致账号异常或封禁，因此请谨慎使用，建议使用代理IP进行请求。 - 交互类接口的最终结果可能会受到TikTok风控系统的影响，大多数情况跟你所使用的账号有关，比如新注册的账号可能无法关注用户或点赞视频，我们无法处理基于账号的风控问题。 - 请不要使用交互类接口对他人造成骚扰，或进行违法违规的操作，否则我们将会停止对你的服务，并且所有责任由你自己承担。 ### 参数: - aweme_id: 视频id，可以从分享链接中获取，例如：https://www.tiktok.com/@username/video/7419966340443819295 - text: 评论内容，TikTok评论内容需要符合规范，不要带有违规的关键词，否则即使请求成功也会被系统判定为垃圾评论从而不被展示，提交时请使用URL编码评论字符串。 - cookie: 用户Cookie，可以从浏览器中登录自己的TikTok账号，然后复制Cookie信息，提交时请使用URL编码Cookie字符串。 - device_id: 设备id，可选，如果不填写则会自动生成，如果需要自定义设备id，请使用设备信息接口获取设备id。 - iid: 设备安装id，可选，如果不填写则会自动生成，如果需要自定义设备iid，请使用设备信息接口获取设备iid。 - proxy: 代理IP，可选，如果不填写则会使用服务器IP进行请求（不推荐），建议使用代理IP进行请求防止账号异常获被封禁，支持格式如下：     - 代理IP:端口     - 用户名:密码@代理IP:端口 ### 返回: - 发送评论结果以及评论数据和设备信息，请自行保留设备信息，如请求时使用的`device_id`以及`iid`，以便后续调用接口时使用，频繁更换设备信息可能会导致账号异常或封禁。  # [English] ### Purpose: - Post comments to the specified video using user cookies, please try to use TikTok accounts in the United States as much as possible, and use proxy IPs in the United States when obtaining cookies. ### Note: - Interactive interfaces all need to use the user's Cookie, so please make sure that your Cookie is valid, otherwise the request will not be made normally. - Interactive interfaces may cause account exceptions or bans, so please use them with caution, and it is recommended to use proxy IPs for requests. - The final result of the interactive interface may be affected by the TikTok risk control system, and in most cases it is related to the account you are using, for example, a newly registered account may not be able to follow users or like videos, and we cannot handle risk control issues based on accounts. - Please do not use interactive interfaces to harass others, or engage in illegal or irregular operations, otherwise we will stop providing services to you, and all responsibilities will be borne by you. ### Parameters: - aweme_id: Video id, which can be obtained from the sharing link, for example: https://www.tiktok.com/@username/video/7419966340443819295 - text: Comment content, TikTok comment content needs to comply with the specifications, do not contain illegal keywords, otherwise, even if the request is successful, it will be judged as spam comments by the system and will not be displayed, please use URL-encoded comment string when submitting. - cookie: User Cookie, you can log in to your TikTok account in the browser and then copy the Cookie information, please use URL-encoded Cookie string when submitting. - device_id: Device id, optional, if not filled in, it will be automatically generated, if you need to customize the device id, please use the device information interface to get the device id. - iid: Device install id, optional, if not filled in, it will be automatically generated, if you need to customize the device iid, please use the device information interface to get the device iid. - proxy: Proxy IP, optional, if not filled in, the server IP will be used for the request (not recommended), it is recommended to use a proxy IP for the request to prevent account exceptions or bans, support the following formats:     - Proxy IP:Port     - Username:Password@Proxy IP:Port ### Return: - Post comment results, comment data and device information, please keep the device information, such as the `device_id` and `iid` used when requesting, for subsequent calls to the interface, frequent replacement of device information may cause account exceptions or bans.  # [示例/Example] ```python # Python Code text = urllib.parse.quote(\"Hello, TikTok!\") cookie = urllib.parse.quote(\"Your_Cookie_From_Browser\") payload = {     \"aweme_id\": \"7419966340443819295\",     \"text\": text,     \"cookie\": cookie,     \"proxy\": \"\", } ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostCommentApiV1TiktokInteractionPostCommentPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostCommentApiV1TiktokInteractionPostCommentPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postCommentRequest** | [**optional.Interface of PostCommentRequest**](PostCommentRequest.md)|  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommentApiV1TiktokInteractionPostCommentPost_0

> ResponseModel PostCommentApiV1TiktokInteractionPostCommentPost_0(ctx, optional)

发送评论/Post comment

# [中文] ### 用途: - 使用用户Cookie发送评论到指定视频，当前请尽可能使用美国地区的TikTok账号，并且在获取Cookie时请使用美国地区的代理IP。 ### 注意: - 交互类接口都需要使用用户的Cookie，因此请确保你的Cookie是有效的，否则将无法正常请求。 - 交互类的接口可能会导致账号异常或封禁，因此请谨慎使用，建议使用代理IP进行请求。 - 交互类接口的最终结果可能会受到TikTok风控系统的影响，大多数情况跟你所使用的账号有关，比如新注册的账号可能无法关注用户或点赞视频，我们无法处理基于账号的风控问题。 - 请不要使用交互类接口对他人造成骚扰，或进行违法违规的操作，否则我们将会停止对你的服务，并且所有责任由你自己承担。 ### 参数: - aweme_id: 视频id，可以从分享链接中获取，例如：https://www.tiktok.com/@username/video/7419966340443819295 - text: 评论内容，TikTok评论内容需要符合规范，不要带有违规的关键词，否则即使请求成功也会被系统判定为垃圾评论从而不被展示，提交时请使用URL编码评论字符串。 - cookie: 用户Cookie，可以从浏览器中登录自己的TikTok账号，然后复制Cookie信息，提交时请使用URL编码Cookie字符串。 - device_id: 设备id，可选，如果不填写则会自动生成，如果需要自定义设备id，请使用设备信息接口获取设备id。 - iid: 设备安装id，可选，如果不填写则会自动生成，如果需要自定义设备iid，请使用设备信息接口获取设备iid。 - proxy: 代理IP，可选，如果不填写则会使用服务器IP进行请求（不推荐），建议使用代理IP进行请求防止账号异常获被封禁，支持格式如下：     - 代理IP:端口     - 用户名:密码@代理IP:端口 ### 返回: - 发送评论结果以及评论数据和设备信息，请自行保留设备信息，如请求时使用的`device_id`以及`iid`，以便后续调用接口时使用，频繁更换设备信息可能会导致账号异常或封禁。  # [English] ### Purpose: - Post comments to the specified video using user cookies, please try to use TikTok accounts in the United States as much as possible, and use proxy IPs in the United States when obtaining cookies. ### Note: - Interactive interfaces all need to use the user's Cookie, so please make sure that your Cookie is valid, otherwise the request will not be made normally. - Interactive interfaces may cause account exceptions or bans, so please use them with caution, and it is recommended to use proxy IPs for requests. - The final result of the interactive interface may be affected by the TikTok risk control system, and in most cases it is related to the account you are using, for example, a newly registered account may not be able to follow users or like videos, and we cannot handle risk control issues based on accounts. - Please do not use interactive interfaces to harass others, or engage in illegal or irregular operations, otherwise we will stop providing services to you, and all responsibilities will be borne by you. ### Parameters: - aweme_id: Video id, which can be obtained from the sharing link, for example: https://www.tiktok.com/@username/video/7419966340443819295 - text: Comment content, TikTok comment content needs to comply with the specifications, do not contain illegal keywords, otherwise, even if the request is successful, it will be judged as spam comments by the system and will not be displayed, please use URL-encoded comment string when submitting. - cookie: User Cookie, you can log in to your TikTok account in the browser and then copy the Cookie information, please use URL-encoded Cookie string when submitting. - device_id: Device id, optional, if not filled in, it will be automatically generated, if you need to customize the device id, please use the device information interface to get the device id. - iid: Device install id, optional, if not filled in, it will be automatically generated, if you need to customize the device iid, please use the device information interface to get the device iid. - proxy: Proxy IP, optional, if not filled in, the server IP will be used for the request (not recommended), it is recommended to use a proxy IP for the request to prevent account exceptions or bans, support the following formats:     - Proxy IP:Port     - Username:Password@Proxy IP:Port ### Return: - Post comment results, comment data and device information, please keep the device information, such as the `device_id` and `iid` used when requesting, for subsequent calls to the interface, frequent replacement of device information may cause account exceptions or bans.  # [示例/Example] ```python # Python Code text = urllib.parse.quote(\"Hello, TikTok!\") cookie = urllib.parse.quote(\"Your_Cookie_From_Browser\") payload = {     \"aweme_id\": \"7419966340443819295\",     \"text\": text,     \"cookie\": cookie,     \"proxy\": \"\", } ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostCommentApiV1TiktokInteractionPostCommentPost_6Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostCommentApiV1TiktokInteractionPostCommentPost_6Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postCommentRequest** | [**optional.Interface of PostCommentRequest**](PostCommentRequest.md)|  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplyCommentApiV1TiktokInteractionReplyCommentPost

> ResponseModel ReplyCommentApiV1TiktokInteractionReplyCommentPost(ctx, optional)

回复评论/Reply to comment

# [中文] ### 用途: - 使用用户Cookie回复指定视频的一个评论，当前请尽可能使用美国地区的TikTok账号，并且在获取Cookie时请使用美国地区的代理IP。 ### 注意: - 交互类接口都需要使用用户的Cookie，因此请确保你的Cookie是有效的，否则将无法正常请求。 - 交互类的接口可能会导致账号异常或封禁，因此请谨慎使用，建议使用代理IP进行请求。 - 交互类接口的最终结果可能会受到TikTok风控系统的影响，大多数情况跟你所使用的账号有关，比如新注册的账号可能无法关注用户或点赞视频，我们无法处理基于账号的风控问题。 - 请不要使用交互类接口对他人造成骚扰，或进行违法违规的操作，否则我们将会停止对你的服务，并且所有责任由你自己承担。 ### 参数: - aweme_id: 视频id，可以从分享链接中获取，例如：https://www.tiktok.com/@username/video/7419966340443819295 - reply_id: 要回复的目标评论ID，可以从指定视频的评论数据接口中获取，通常关键字为`cid`或`comment_id`或`reply_id`。 - text: 评论内容，TikTok评论内容需要符合规范，不要带有违规的关键词，否则即使请求成功也会被系统判定为垃圾评论从而不被展示，提交时请使用URL编码评论字符串。 - cookie: 用户Cookie，可以从浏览器中登录自己的TikTok账号，然后复制Cookie信息，提交时请使用URL编码Cookie字符串。 - device_id: 设备id，可选，如果不填写则会自动生成，如果需要自定义设备id，请使用设备信息接口获取设备id。 - iid: 设备安装id，可选，如果不填写则会自动生成，如果需要自定义设备iid，请使用设备信息接口获取设备iid。 - proxy: 代理IP，可选，如果不填写则会使用服务器IP进行请求（不推荐），建议使用代理IP进行请求防止账号异常获被封禁，支持格式如下：     - 代理IP:端口     - 用户名:密码@代理IP:端口 ### 返回: - 回复评论结果以及评论数据和设备信息，请自行保留设备信息，如请求时使用的`device_id`以及`iid`，以便后续调用接口时使用，频繁更换设备信息可能会导致账号异常或封禁。  # [English] ### Purpose: - Reply to a comment on a specified video using user cookies, please try to use TikTok accounts in the United States as much as possible, and use proxy IPs in the United States when obtaining cookies. ### Note: - Interactive interfaces all need to use the user's Cookie, so please make sure that your Cookie is valid, otherwise the request will not be made normally. - Interactive interfaces may cause account exceptions or bans, so please use them with caution, and it is recommended to use proxy IPs for requests. - The final result of the interactive interface may be affected by the TikTok risk control system, and in most cases it is related to the account you are using, for example, a newly registered account may not be able to follow users or like videos, and we cannot handle risk control issues based on accounts. - Please do not use interactive interfaces to harass others, or engage in illegal or irregular operations, otherwise we will stop providing services to you, and all responsibilities will be borne by you. ### Parameters: - aweme_id: Video id, which can be obtained from the sharing link, for example: https://www.tiktok.com/@username/video/7419966340443819295 - reply_id: The target comment ID to reply to, which can be obtained from the comment data interface of the specified video, usually the keyword is `cid` or `comment_id` or `reply_id`. - text: Comment content, TikTok comment content needs to comply with the specifications, do not contain illegal keywords, otherwise, even if the request is successful, it will be judged as spam comments by the system and will not be displayed, please use URL-encoded comment string when submitting. - cookie: User Cookie, you can log in to your TikTok account in the browser and then copy the Cookie information, please use URL-encoded Cookie string when submitting. - device_id: Device id, optional, if not filled in, it will be automatically generated, if you need to customize the device id, please use the device information interface to get the device id. - iid: Device install id, optional, if not filled in, it will be automatically generated, if you need to customize the device iid, please use the device information interface to get the device iid. - proxy: Proxy IP, optional, if not filled in, the server IP will be used for the request (not recommended), it is recommended to use a proxy IP for the request to prevent account exceptions or bans, support the following formats:     - Proxy IP:Port     - Username:Password@Proxy IP:Port ### Return: - Reply comment results, comment data and device information, please keep the device information, such as the `device_id` and `iid` used when requesting, for subsequent calls to the interface, frequent replacement of device information may cause account exceptions or bans.  # [示例/Example] ```python # Python Code text = urllib.parse.quote(\"Hello, TikTok!\") cookie = urllib.parse.quote(\"Your_Cookie_From_Browser\") payload = {     \"aweme_id\": \"7419966340443819295\",     \"reply_id\": \"7420673787547419435\",     \"text\": text,     \"cookie\": cookie,     \"proxy\": \"\", } ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplyCommentApiV1TiktokInteractionReplyCommentPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReplyCommentApiV1TiktokInteractionReplyCommentPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replyCommentRequest** | [**optional.Interface of ReplyCommentRequest**](ReplyCommentRequest.md)|  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplyCommentApiV1TiktokInteractionReplyCommentPost_0

> ResponseModel ReplyCommentApiV1TiktokInteractionReplyCommentPost_0(ctx, optional)

回复评论/Reply to comment

# [中文] ### 用途: - 使用用户Cookie回复指定视频的一个评论，当前请尽可能使用美国地区的TikTok账号，并且在获取Cookie时请使用美国地区的代理IP。 ### 注意: - 交互类接口都需要使用用户的Cookie，因此请确保你的Cookie是有效的，否则将无法正常请求。 - 交互类的接口可能会导致账号异常或封禁，因此请谨慎使用，建议使用代理IP进行请求。 - 交互类接口的最终结果可能会受到TikTok风控系统的影响，大多数情况跟你所使用的账号有关，比如新注册的账号可能无法关注用户或点赞视频，我们无法处理基于账号的风控问题。 - 请不要使用交互类接口对他人造成骚扰，或进行违法违规的操作，否则我们将会停止对你的服务，并且所有责任由你自己承担。 ### 参数: - aweme_id: 视频id，可以从分享链接中获取，例如：https://www.tiktok.com/@username/video/7419966340443819295 - reply_id: 要回复的目标评论ID，可以从指定视频的评论数据接口中获取，通常关键字为`cid`或`comment_id`或`reply_id`。 - text: 评论内容，TikTok评论内容需要符合规范，不要带有违规的关键词，否则即使请求成功也会被系统判定为垃圾评论从而不被展示，提交时请使用URL编码评论字符串。 - cookie: 用户Cookie，可以从浏览器中登录自己的TikTok账号，然后复制Cookie信息，提交时请使用URL编码Cookie字符串。 - device_id: 设备id，可选，如果不填写则会自动生成，如果需要自定义设备id，请使用设备信息接口获取设备id。 - iid: 设备安装id，可选，如果不填写则会自动生成，如果需要自定义设备iid，请使用设备信息接口获取设备iid。 - proxy: 代理IP，可选，如果不填写则会使用服务器IP进行请求（不推荐），建议使用代理IP进行请求防止账号异常获被封禁，支持格式如下：     - 代理IP:端口     - 用户名:密码@代理IP:端口 ### 返回: - 回复评论结果以及评论数据和设备信息，请自行保留设备信息，如请求时使用的`device_id`以及`iid`，以便后续调用接口时使用，频繁更换设备信息可能会导致账号异常或封禁。  # [English] ### Purpose: - Reply to a comment on a specified video using user cookies, please try to use TikTok accounts in the United States as much as possible, and use proxy IPs in the United States when obtaining cookies. ### Note: - Interactive interfaces all need to use the user's Cookie, so please make sure that your Cookie is valid, otherwise the request will not be made normally. - Interactive interfaces may cause account exceptions or bans, so please use them with caution, and it is recommended to use proxy IPs for requests. - The final result of the interactive interface may be affected by the TikTok risk control system, and in most cases it is related to the account you are using, for example, a newly registered account may not be able to follow users or like videos, and we cannot handle risk control issues based on accounts. - Please do not use interactive interfaces to harass others, or engage in illegal or irregular operations, otherwise we will stop providing services to you, and all responsibilities will be borne by you. ### Parameters: - aweme_id: Video id, which can be obtained from the sharing link, for example: https://www.tiktok.com/@username/video/7419966340443819295 - reply_id: The target comment ID to reply to, which can be obtained from the comment data interface of the specified video, usually the keyword is `cid` or `comment_id` or `reply_id`. - text: Comment content, TikTok comment content needs to comply with the specifications, do not contain illegal keywords, otherwise, even if the request is successful, it will be judged as spam comments by the system and will not be displayed, please use URL-encoded comment string when submitting. - cookie: User Cookie, you can log in to your TikTok account in the browser and then copy the Cookie information, please use URL-encoded Cookie string when submitting. - device_id: Device id, optional, if not filled in, it will be automatically generated, if you need to customize the device id, please use the device information interface to get the device id. - iid: Device install id, optional, if not filled in, it will be automatically generated, if you need to customize the device iid, please use the device information interface to get the device iid. - proxy: Proxy IP, optional, if not filled in, the server IP will be used for the request (not recommended), it is recommended to use a proxy IP for the request to prevent account exceptions or bans, support the following formats:     - Proxy IP:Port     - Username:Password@Proxy IP:Port ### Return: - Reply comment results, comment data and device information, please keep the device information, such as the `device_id` and `iid` used when requesting, for subsequent calls to the interface, frequent replacement of device information may cause account exceptions or bans.  # [示例/Example] ```python # Python Code text = urllib.parse.quote(\"Hello, TikTok!\") cookie = urllib.parse.quote(\"Your_Cookie_From_Browser\") payload = {     \"aweme_id\": \"7419966340443819295\",     \"reply_id\": \"7420673787547419435\",     \"text\": text,     \"cookie\": cookie,     \"proxy\": \"\", } ```

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplyCommentApiV1TiktokInteractionReplyCommentPost_7Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReplyCommentApiV1TiktokInteractionReplyCommentPost_7Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replyCommentRequest** | [**optional.Interface of ReplyCommentRequest**](ReplyCommentRequest.md)|  | 

### Return type

[**ResponseModel**](ResponseModel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

