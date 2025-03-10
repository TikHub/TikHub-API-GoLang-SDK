/*
 * TikHub Private API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tikhub

// XhsWebSignRequestModel XhsWebSignRequestModel
type XhsWebSignRequestModel struct {
	// Path，请求接口的路径/Request API path
	Path string `json:"path,omitempty" xml:"path"`
	// Data，请求API的荷载数据/Payload data of request API
	Data map[string]interface{} `json:"data,omitempty" xml:"data"`
	// Cookie，请求接口的Cookie/Request API cookie
	Cookie string `json:"cookie,omitempty" xml:"cookie"`
}
