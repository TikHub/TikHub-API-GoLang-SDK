/*
 * TikHub Private API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tikhub

// BodyFetchHotTotalHighSearchListApiV1DouyinBillboardFetchHotTotalHighSearchListPost Body_fetch_hot_total_high_search_list_api_v1_douyin_billboard_fetch_hot_total_high_search_list_post
type BodyFetchHotTotalHighSearchListApiV1DouyinBillboardFetchHotTotalHighSearchListPost struct {
	// Page Num，页码
	PageNum int32 `json:"page_num,omitempty" xml:"page_num"`
	// Page Size，每页数量
	PageSize int32 `json:"page_size,omitempty" xml:"page_size"`
	// Date Window，时间窗口，1 按小时 2 按天
	DateWindow int32 `json:"date_window,omitempty" xml:"date_window"`
	// Keyword，搜索关键字
	Keyword string `json:"keyword,omitempty" xml:"keyword"`
}
