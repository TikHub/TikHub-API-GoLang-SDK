/*
 * TikHub Private API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tikhub

// BodyHcaptchaApiV1CaptchaHcaptchaPost Body_hcaptcha_api_v1_captcha_hcaptcha_post
type BodyHcaptchaApiV1CaptchaHcaptchaPost struct {
	// Sitekey，Sitekey
	Sitekey string `json:"sitekey" xml:"sitekey"`
	// Url，URL
	Url string `json:"url" xml:"url"`
	// Proxy，Proxy
	Proxy map[string]interface{} `json:"proxy,omitempty" xml:"proxy"`
}
