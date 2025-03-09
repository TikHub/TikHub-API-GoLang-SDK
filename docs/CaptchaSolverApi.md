# \CaptchaSolverApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AmazonCaptchaApiV1CaptchaAmazonCaptchaPost**](CaptchaSolverApi.md#AmazonCaptchaApiV1CaptchaAmazonCaptchaPost) | **Post** /api/v1/captcha/amazon_captcha | Amazon Captcha Solver/Amazon验证码解决器
[**AmazonCaptchaApiV1CaptchaAmazonCaptchaPost_0**](CaptchaSolverApi.md#AmazonCaptchaApiV1CaptchaAmazonCaptchaPost_0) | **Post** /api/v1/captcha/amazon_captcha | Amazon Captcha Solver/Amazon验证码解决器
[**CloudflareTurnstileApiV1CaptchaCloudflareTurnstilePost**](CaptchaSolverApi.md#CloudflareTurnstileApiV1CaptchaCloudflareTurnstilePost) | **Post** /api/v1/captcha/cloudflare_turnstile | Cloudflare Turnstile Solver/Cloudflare Turnstile解决器
[**CloudflareTurnstileApiV1CaptchaCloudflareTurnstilePost_0**](CaptchaSolverApi.md#CloudflareTurnstileApiV1CaptchaCloudflareTurnstilePost_0) | **Post** /api/v1/captcha/cloudflare_turnstile | Cloudflare Turnstile Solver/Cloudflare Turnstile解决器
[**HcaptchaApiV1CaptchaHcaptchaPost**](CaptchaSolverApi.md#HcaptchaApiV1CaptchaHcaptchaPost) | **Post** /api/v1/captcha/hcaptcha | hCaptcha Solver/hCaptcha解决器
[**HcaptchaApiV1CaptchaHcaptchaPost_0**](CaptchaSolverApi.md#HcaptchaApiV1CaptchaHcaptchaPost_0) | **Post** /api/v1/captcha/hcaptcha | hCaptcha Solver/hCaptcha解决器
[**RecaptchaV2ApiV1CaptchaRecaptchaV2Post**](CaptchaSolverApi.md#RecaptchaV2ApiV1CaptchaRecaptchaV2Post) | **Post** /api/v1/captcha/recaptcha_v2 | Recaptcha V2 Solver/Recaptcha V2解决器
[**RecaptchaV2ApiV1CaptchaRecaptchaV2Post_0**](CaptchaSolverApi.md#RecaptchaV2ApiV1CaptchaRecaptchaV2Post_0) | **Post** /api/v1/captcha/recaptcha_v2 | Recaptcha V2 Solver/Recaptcha V2解决器
[**RecaptchaV3ApiV1CaptchaRecaptchaV3Post**](CaptchaSolverApi.md#RecaptchaV3ApiV1CaptchaRecaptchaV3Post) | **Post** /api/v1/captcha/recaptcha_v3 | Recaptcha V3 Solver/Recaptcha V3解决器
[**RecaptchaV3ApiV1CaptchaRecaptchaV3Post_0**](CaptchaSolverApi.md#RecaptchaV3ApiV1CaptchaRecaptchaV3Post_0) | **Post** /api/v1/captcha/recaptcha_v3 | Recaptcha V3 Solver/Recaptcha V3解决器
[**TencentCaptchaApiV1CaptchaTencentCaptchaPost**](CaptchaSolverApi.md#TencentCaptchaApiV1CaptchaTencentCaptchaPost) | **Post** /api/v1/captcha/tencent_captcha | Tencent Captcha Solver/Tencent验证码解决器
[**TencentCaptchaApiV1CaptchaTencentCaptchaPost_0**](CaptchaSolverApi.md#TencentCaptchaApiV1CaptchaTencentCaptchaPost_0) | **Post** /api/v1/captcha/tencent_captcha | Tencent Captcha Solver/Tencent验证码解决器



## AmazonCaptchaApiV1CaptchaAmazonCaptchaPost

> ResponseModel AmazonCaptchaApiV1CaptchaAmazonCaptchaPost(ctx, optional)

Amazon Captcha Solver/Amazon验证码解决器

# [中文] ### 用途: - Amazon Captcha验证码解决器 ### 参数: - app_id: 在HTML中可以找到网站对应的app_id - url: 需要解决验证码的URL - proxy: 默认为None ### 返回: - 返回验证码解决结果  # [English] ### Purpose: - Amazon Captcha solver ### Parameters: - app_id: The app_id corresponding to the website can be found in the HTML - url: URL that needs to solve the captcha - proxy: Default is None ### Return: - Return the captcha solution result  # [Example/示例] app_id = \"10000000\" url = \"https://www.amazon.com/\" proxy = None

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AmazonCaptchaApiV1CaptchaAmazonCaptchaPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AmazonCaptchaApiV1CaptchaAmazonCaptchaPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyAmazonCaptchaApiV1CaptchaAmazonCaptchaPost** | [**optional.Interface of BodyAmazonCaptchaApiV1CaptchaAmazonCaptchaPost**](BodyAmazonCaptchaApiV1CaptchaAmazonCaptchaPost.md)|  | 

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


## AmazonCaptchaApiV1CaptchaAmazonCaptchaPost_0

> ResponseModel AmazonCaptchaApiV1CaptchaAmazonCaptchaPost_0(ctx, optional)

Amazon Captcha Solver/Amazon验证码解决器

# [中文] ### 用途: - Amazon Captcha验证码解决器 ### 参数: - app_id: 在HTML中可以找到网站对应的app_id - url: 需要解决验证码的URL - proxy: 默认为None ### 返回: - 返回验证码解决结果  # [English] ### Purpose: - Amazon Captcha solver ### Parameters: - app_id: The app_id corresponding to the website can be found in the HTML - url: URL that needs to solve the captcha - proxy: Default is None ### Return: - Return the captcha solution result  # [Example/示例] app_id = \"10000000\" url = \"https://www.amazon.com/\" proxy = None

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AmazonCaptchaApiV1CaptchaAmazonCaptchaPost_1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AmazonCaptchaApiV1CaptchaAmazonCaptchaPost_1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyAmazonCaptchaApiV1CaptchaAmazonCaptchaPost** | [**optional.Interface of BodyAmazonCaptchaApiV1CaptchaAmazonCaptchaPost**](BodyAmazonCaptchaApiV1CaptchaAmazonCaptchaPost.md)|  | 

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


## CloudflareTurnstileApiV1CaptchaCloudflareTurnstilePost

> ResponseModel CloudflareTurnstileApiV1CaptchaCloudflareTurnstilePost(ctx, optional)

Cloudflare Turnstile Solver/Cloudflare Turnstile解决器

# [中文] ### 用途: - Cloudflare Turnstile验证码解决器 ### 参数: - sitekey: 在HTML中可以找到网站对应的sitekey - url: 需要解决验证码的URL - proxy: 默认为None ### 返回: - 返回验证码解决结果  # [English] ### Purpose: - Cloudflare Turnstile captcha solver ### Parameters: - sitekey: The sitekey corresponding to the website can be found in the HTML - url: URL that needs to solve the captcha - action: Default is None - data: Default is None - proxy: Default is None ### Return: - Return the captcha solution result  # [Example/示例] sitekey = \"1x00000000000000000000AA\" url = \"https://demo.turnstile.workers.dev\" proxy = None

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CloudflareTurnstileApiV1CaptchaCloudflareTurnstilePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CloudflareTurnstileApiV1CaptchaCloudflareTurnstilePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyCloudflareTurnstileApiV1CaptchaCloudflareTurnstilePost** | [**optional.Interface of BodyCloudflareTurnstileApiV1CaptchaCloudflareTurnstilePost**](BodyCloudflareTurnstileApiV1CaptchaCloudflareTurnstilePost.md)|  | 

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


## CloudflareTurnstileApiV1CaptchaCloudflareTurnstilePost_0

> ResponseModel CloudflareTurnstileApiV1CaptchaCloudflareTurnstilePost_0(ctx, optional)

Cloudflare Turnstile Solver/Cloudflare Turnstile解决器

# [中文] ### 用途: - Cloudflare Turnstile验证码解决器 ### 参数: - sitekey: 在HTML中可以找到网站对应的sitekey - url: 需要解决验证码的URL - proxy: 默认为None ### 返回: - 返回验证码解决结果  # [English] ### Purpose: - Cloudflare Turnstile captcha solver ### Parameters: - sitekey: The sitekey corresponding to the website can be found in the HTML - url: URL that needs to solve the captcha - action: Default is None - data: Default is None - proxy: Default is None ### Return: - Return the captcha solution result  # [Example/示例] sitekey = \"1x00000000000000000000AA\" url = \"https://demo.turnstile.workers.dev\" proxy = None

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CloudflareTurnstileApiV1CaptchaCloudflareTurnstilePost_2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CloudflareTurnstileApiV1CaptchaCloudflareTurnstilePost_2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyCloudflareTurnstileApiV1CaptchaCloudflareTurnstilePost** | [**optional.Interface of BodyCloudflareTurnstileApiV1CaptchaCloudflareTurnstilePost**](BodyCloudflareTurnstileApiV1CaptchaCloudflareTurnstilePost.md)|  | 

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


## HcaptchaApiV1CaptchaHcaptchaPost

> ResponseModel HcaptchaApiV1CaptchaHcaptchaPost(ctx, optional)

hCaptcha Solver/hCaptcha解决器

# [中文] ### 用途: - hCaptcha验证码解决器 ### 参数: - sitekey: 在HTML中可以找到网站对应的sitekey - url: 需要解决验证码的URL - proxy: 默认为None ### 返回: - 返回验证码解决结果  # [English] ### Purpose: - hCaptcha captcha solver ### Parameters: - sitekey: The sitekey corresponding to the website can be found in the HTML - url: URL that needs to solve the captcha - proxy: Default is None ### Return: - Return the captcha solution result  # [Example/示例] sitekey = \"10000000-ffff-ffff-ffff-000000000001\" url = \"https://www.hcaptcha.com/\" proxy = None

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HcaptchaApiV1CaptchaHcaptchaPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HcaptchaApiV1CaptchaHcaptchaPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyHcaptchaApiV1CaptchaHcaptchaPost** | [**optional.Interface of BodyHcaptchaApiV1CaptchaHcaptchaPost**](BodyHcaptchaApiV1CaptchaHcaptchaPost.md)|  | 

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


## HcaptchaApiV1CaptchaHcaptchaPost_0

> ResponseModel HcaptchaApiV1CaptchaHcaptchaPost_0(ctx, optional)

hCaptcha Solver/hCaptcha解决器

# [中文] ### 用途: - hCaptcha验证码解决器 ### 参数: - sitekey: 在HTML中可以找到网站对应的sitekey - url: 需要解决验证码的URL - proxy: 默认为None ### 返回: - 返回验证码解决结果  # [English] ### Purpose: - hCaptcha captcha solver ### Parameters: - sitekey: The sitekey corresponding to the website can be found in the HTML - url: URL that needs to solve the captcha - proxy: Default is None ### Return: - Return the captcha solution result  # [Example/示例] sitekey = \"10000000-ffff-ffff-ffff-000000000001\" url = \"https://www.hcaptcha.com/\" proxy = None

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HcaptchaApiV1CaptchaHcaptchaPost_3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HcaptchaApiV1CaptchaHcaptchaPost_3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyHcaptchaApiV1CaptchaHcaptchaPost** | [**optional.Interface of BodyHcaptchaApiV1CaptchaHcaptchaPost**](BodyHcaptchaApiV1CaptchaHcaptchaPost.md)|  | 

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


## RecaptchaV2ApiV1CaptchaRecaptchaV2Post

> ResponseModel RecaptchaV2ApiV1CaptchaRecaptchaV2Post(ctx, optional)

Recaptcha V2 Solver/Recaptcha V2解决器

# [中文] ### 用途: - Recaptcha V2验证码解决器 ### 参数: - sitekey: 在HTML中可以找到网站对应的sitekey - url: 需要解决验证码的URL - proxy: 默认为None ### 返回: - 返回验证码解决结果  # [English] ### Purpose: - Recaptcha V2 captcha solver ### Parameters: - sitekey: The sitekey corresponding to the website can be found in the HTML - url: URL that needs to solve the captcha - proxy: Default is None ### Return: - Return the captcha solution result  # [Example/示例] sitekey = \"6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-\" url = \"https://www.google.com/recaptcha/api2/demo\" proxy = None

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecaptchaV2ApiV1CaptchaRecaptchaV2PostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RecaptchaV2ApiV1CaptchaRecaptchaV2PostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyRecaptchaV2ApiV1CaptchaRecaptchaV2Post** | [**optional.Interface of BodyRecaptchaV2ApiV1CaptchaRecaptchaV2Post**](BodyRecaptchaV2ApiV1CaptchaRecaptchaV2Post.md)|  | 

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


## RecaptchaV2ApiV1CaptchaRecaptchaV2Post_0

> ResponseModel RecaptchaV2ApiV1CaptchaRecaptchaV2Post_0(ctx, optional)

Recaptcha V2 Solver/Recaptcha V2解决器

# [中文] ### 用途: - Recaptcha V2验证码解决器 ### 参数: - sitekey: 在HTML中可以找到网站对应的sitekey - url: 需要解决验证码的URL - proxy: 默认为None ### 返回: - 返回验证码解决结果  # [English] ### Purpose: - Recaptcha V2 captcha solver ### Parameters: - sitekey: The sitekey corresponding to the website can be found in the HTML - url: URL that needs to solve the captcha - proxy: Default is None ### Return: - Return the captcha solution result  # [Example/示例] sitekey = \"6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-\" url = \"https://www.google.com/recaptcha/api2/demo\" proxy = None

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecaptchaV2ApiV1CaptchaRecaptchaV2Post_4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RecaptchaV2ApiV1CaptchaRecaptchaV2Post_4Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyRecaptchaV2ApiV1CaptchaRecaptchaV2Post** | [**optional.Interface of BodyRecaptchaV2ApiV1CaptchaRecaptchaV2Post**](BodyRecaptchaV2ApiV1CaptchaRecaptchaV2Post.md)|  | 

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


## RecaptchaV3ApiV1CaptchaRecaptchaV3Post

> ResponseModel RecaptchaV3ApiV1CaptchaRecaptchaV3Post(ctx, optional)

Recaptcha V3 Solver/Recaptcha V3解决器

# [中文] ### 用途: - Recaptcha V3验证码解决器 ### 参数: - sitekey: 在HTML中可以找到网站对应的sitekey - url: 需要解决验证码的URL - proxy: 默认为None ### 返回: - 返回验证码解决结果  # [English] ### Purpose: - Recaptcha V3 captcha solver ### Parameters: - sitekey: The sitekey corresponding to the website can be found in the HTML - url: URL that needs to solve the captcha - proxy: Default is None ### Return: - Return the captcha solution result  # [Example/示例] sitekey = \"6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-\" url = \"https://www.google.com/recaptcha/api2/demo\" proxy = None

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecaptchaV3ApiV1CaptchaRecaptchaV3PostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RecaptchaV3ApiV1CaptchaRecaptchaV3PostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyRecaptchaV3ApiV1CaptchaRecaptchaV3Post** | [**optional.Interface of BodyRecaptchaV3ApiV1CaptchaRecaptchaV3Post**](BodyRecaptchaV3ApiV1CaptchaRecaptchaV3Post.md)|  | 

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


## RecaptchaV3ApiV1CaptchaRecaptchaV3Post_0

> ResponseModel RecaptchaV3ApiV1CaptchaRecaptchaV3Post_0(ctx, optional)

Recaptcha V3 Solver/Recaptcha V3解决器

# [中文] ### 用途: - Recaptcha V3验证码解决器 ### 参数: - sitekey: 在HTML中可以找到网站对应的sitekey - url: 需要解决验证码的URL - proxy: 默认为None ### 返回: - 返回验证码解决结果  # [English] ### Purpose: - Recaptcha V3 captcha solver ### Parameters: - sitekey: The sitekey corresponding to the website can be found in the HTML - url: URL that needs to solve the captcha - proxy: Default is None ### Return: - Return the captcha solution result  # [Example/示例] sitekey = \"6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-\" url = \"https://www.google.com/recaptcha/api2/demo\" proxy = None

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RecaptchaV3ApiV1CaptchaRecaptchaV3Post_5Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RecaptchaV3ApiV1CaptchaRecaptchaV3Post_5Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyRecaptchaV3ApiV1CaptchaRecaptchaV3Post** | [**optional.Interface of BodyRecaptchaV3ApiV1CaptchaRecaptchaV3Post**](BodyRecaptchaV3ApiV1CaptchaRecaptchaV3Post.md)|  | 

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


## TencentCaptchaApiV1CaptchaTencentCaptchaPost

> ResponseModel TencentCaptchaApiV1CaptchaTencentCaptchaPost(ctx, optional)

Tencent Captcha Solver/Tencent验证码解决器

# [中文] ### 用途: - Tencent Captcha验证码解决器 ### 参数: - app_id: 在HTML中可以找到网站对应的app_id - url: 需要解决验证码的URL - proxy: 默认为None ### 返回: - 返回验证码解决结果  # [English] ### Purpose: - Tencent Captcha solver ### Parameters: - app_id: The app_id corresponding to the website can be found in the HTML - url: URL that needs to solve the captcha - proxy: Default is None ### Return: - Return the captcha solution result  # [Example/示例] app_id = \"10000000\" url = \"https://www.tencent.com/\" proxy = None

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TencentCaptchaApiV1CaptchaTencentCaptchaPostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TencentCaptchaApiV1CaptchaTencentCaptchaPostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyTencentCaptchaApiV1CaptchaTencentCaptchaPost** | [**optional.Interface of BodyTencentCaptchaApiV1CaptchaTencentCaptchaPost**](BodyTencentCaptchaApiV1CaptchaTencentCaptchaPost.md)|  | 

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


## TencentCaptchaApiV1CaptchaTencentCaptchaPost_0

> ResponseModel TencentCaptchaApiV1CaptchaTencentCaptchaPost_0(ctx, optional)

Tencent Captcha Solver/Tencent验证码解决器

# [中文] ### 用途: - Tencent Captcha验证码解决器 ### 参数: - app_id: 在HTML中可以找到网站对应的app_id - url: 需要解决验证码的URL - proxy: 默认为None ### 返回: - 返回验证码解决结果  # [English] ### Purpose: - Tencent Captcha solver ### Parameters: - app_id: The app_id corresponding to the website can be found in the HTML - url: URL that needs to solve the captcha - proxy: Default is None ### Return: - Return the captcha solution result  # [Example/示例] app_id = \"10000000\" url = \"https://www.tencent.com/\" proxy = None

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TencentCaptchaApiV1CaptchaTencentCaptchaPost_6Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TencentCaptchaApiV1CaptchaTencentCaptchaPost_6Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyTencentCaptchaApiV1CaptchaTencentCaptchaPost** | [**optional.Interface of BodyTencentCaptchaApiV1CaptchaTencentCaptchaPost**](BodyTencentCaptchaApiV1CaptchaTencentCaptchaPost.md)|  | 

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

