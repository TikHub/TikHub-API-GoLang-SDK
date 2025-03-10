/*
 * TikHub Private API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tikhub

// ChallengePostRequest ChallengePostRequest
type ChallengePostRequest struct {
	// Challenge Id，话题ID/Challenge ID
	ChallengeId string `json:"challenge_id,omitempty" xml:"challenge_id"`
	// Sort Type，排序类型/Sort type
	SortType int32 `json:"sort_type,omitempty" xml:"sort_type"`
	// Cursor，游标/Cursor
	Cursor int32 `json:"cursor,omitempty" xml:"cursor"`
	// Count，数量/Count
	Count int32 `json:"count,omitempty" xml:"count"`
	// Cookie，用户自行提供的Cookie/User provided Cookie
	Cookie string `json:"cookie,omitempty" xml:"cookie"`
}
