/*
 * TikHub Private API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tikhub

// ValidationError ValidationError
type ValidationError struct {
	// Location
	Loc []AnyOfstringinteger `json:"loc" xml:"loc"`
	// Message
	Msg string `json:"msg" xml:"msg"`
	// Error Type
	Type string `json:"type" xml:"type"`
}
