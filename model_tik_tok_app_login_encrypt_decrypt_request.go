/*
 * TikHub Private API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tikhub

// TikTokAppLoginEncryptDecryptRequest TikTok_APP_Login_Encrypt_Decrypt_Request
type TikTokAppLoginEncryptDecryptRequest struct {
	// Username，Plaintext or encrypted username
	Username string `json:"username,omitempty" xml:"username"`
	// Password，Plaintext or encrypted password
	Password string `json:"password,omitempty" xml:"password"`
	// Encrypt or decrypt the input string
	Mode ModeEnum `json:"mode,omitempty" xml:"mode"`
}
