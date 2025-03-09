# \TempMailAPIApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEmailByIdApiV1TempMailV1GetEmailByIdGet**](TempMailAPIApi.md#GetEmailByIdApiV1TempMailV1GetEmailByIdGet) | **Get** /api/v1/temp_mail/v1/get_email_by_id | Get Email By Id
[**GetEmailByIdApiV1TempMailV1GetEmailByIdGet_0**](TempMailAPIApi.md#GetEmailByIdApiV1TempMailV1GetEmailByIdGet_0) | **Get** /api/v1/temp_mail/v1/get_email_by_id | Get Email By Id
[**GetEmailsApiV1TempMailV1GetEmailsInboxGet**](TempMailAPIApi.md#GetEmailsApiV1TempMailV1GetEmailsInboxGet) | **Get** /api/v1/temp_mail/v1/get_emails_inbox | Get Emails
[**GetEmailsApiV1TempMailV1GetEmailsInboxGet_0**](TempMailAPIApi.md#GetEmailsApiV1TempMailV1GetEmailsInboxGet_0) | **Get** /api/v1/temp_mail/v1/get_emails_inbox | Get Emails
[**GetTempEmailApiV1TempMailV1GetTempEmailAddressGet**](TempMailAPIApi.md#GetTempEmailApiV1TempMailV1GetTempEmailAddressGet) | **Get** /api/v1/temp_mail/v1/get_temp_email_address | Get Temp Email
[**GetTempEmailApiV1TempMailV1GetTempEmailAddressGet_0**](TempMailAPIApi.md#GetTempEmailApiV1TempMailV1GetTempEmailAddressGet_0) | **Get** /api/v1/temp_mail/v1/get_temp_email_address | Get Temp Email



## GetEmailByIdApiV1TempMailV1GetEmailByIdGet

> ResponseModel GetEmailByIdApiV1TempMailV1GetEmailByIdGet(ctx, token, messageId)

Get Email By Id

# [中文] ### 用途: - 通过邮件ID获取邮件数据 ### 参数: - token: 邮箱Bearer Token - message_id: 邮件ID ### 返回: - 邮件数据  # [English] ### Purpose: - Get email data by email ID ### Parameters: - token: Email Bearer Token - message_id: Email ID ### Returns: - Email data

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string**| Bearer Token | 
**messageId** | **string**| Message ID | 

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


## GetEmailByIdApiV1TempMailV1GetEmailByIdGet_0

> ResponseModel GetEmailByIdApiV1TempMailV1GetEmailByIdGet_0(ctx, token, messageId)

Get Email By Id

# [中文] ### 用途: - 通过邮件ID获取邮件数据 ### 参数: - token: 邮箱Bearer Token - message_id: 邮件ID ### 返回: - 邮件数据  # [English] ### Purpose: - Get email data by email ID ### Parameters: - token: Email Bearer Token - message_id: Email ID ### Returns: - Email data

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string**| Bearer Token | 
**messageId** | **string**| Message ID | 

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


## GetEmailsApiV1TempMailV1GetEmailsInboxGet

> ResponseModel GetEmailsApiV1TempMailV1GetEmailsInboxGet(ctx, token)

Get Emails

# [中文] ### 用途: - 获取邮件列表 ### 参数: - token: 邮箱Bearer Token ### 返回: - emails: 邮件列表  # [English] ### Purpose: - Get a list of emails ### Parameters: - token: Email Bearer Token ### Returns: - emails: List of emails

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string**| Bearer Token | 

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


## GetEmailsApiV1TempMailV1GetEmailsInboxGet_0

> ResponseModel GetEmailsApiV1TempMailV1GetEmailsInboxGet_0(ctx, token)

Get Emails

# [中文] ### 用途: - 获取邮件列表 ### 参数: - token: 邮箱Bearer Token ### 返回: - emails: 邮件列表  # [English] ### Purpose: - Get a list of emails ### Parameters: - token: Email Bearer Token ### Returns: - emails: List of emails

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string**| Bearer Token | 

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


## GetTempEmailApiV1TempMailV1GetTempEmailAddressGet

> ResponseModel GetTempEmailApiV1TempMailV1GetTempEmailAddressGet(ctx, )

Get Temp Email

# [中文] ### 用途: - 获取一个临时邮箱地址 - 用于注册或者接收邮件，该邮箱地址不会被删除，也不会被其他人使用。 - 该邮箱无法发送邮件，只能接收邮件。 - 请自行保存邮箱地址、用户名、密码、Bearer Token，我们无法帮助您找回这些关键信息。 ### 参数: - 无 ### 返回: - domain: 邮箱域名 - name: 邮箱用户名 - password: 邮箱密码 - email_address: 邮箱地址 - token: 邮箱Bearer Token  # [English] ### Purpose: - Get a temporary email address - Used for registration or receiving emails, this email address will not be deleted or used by others. - This email cannot send emails, only receive emails. - Please save the email address, username, password, and Bearer Token yourself, we cannot help you retrieve this critical information. ### Parameters: - None ### Returns: - domain: Email domain - name: Email username - password: Email password - email_address: Email address - token: Email Bearer Token

### Required Parameters

This endpoint does not need any parameter.

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


## GetTempEmailApiV1TempMailV1GetTempEmailAddressGet_0

> ResponseModel GetTempEmailApiV1TempMailV1GetTempEmailAddressGet_0(ctx, )

Get Temp Email

# [中文] ### 用途: - 获取一个临时邮箱地址 - 用于注册或者接收邮件，该邮箱地址不会被删除，也不会被其他人使用。 - 该邮箱无法发送邮件，只能接收邮件。 - 请自行保存邮箱地址、用户名、密码、Bearer Token，我们无法帮助您找回这些关键信息。 ### 参数: - 无 ### 返回: - domain: 邮箱域名 - name: 邮箱用户名 - password: 邮箱密码 - email_address: 邮箱地址 - token: 邮箱Bearer Token  # [English] ### Purpose: - Get a temporary email address - Used for registration or receiving emails, this email address will not be deleted or used by others. - This email cannot send emails, only receive emails. - Please save the email address, username, password, and Bearer Token yourself, we cannot help you retrieve this critical information. ### Parameters: - None ### Returns: - domain: Email domain - name: Email username - password: Email password - email_address: Email address - token: Email Bearer Token

### Required Parameters

This endpoint does not need any parameter.

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

