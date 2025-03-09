/*
 * TikHub Private API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tikhub

// BodyFetchHotTotalHighPlayListApiV1DouyinBillboardFetchHotTotalHighPlayListPost Body_fetch_hot_total_high_play_list_api_v1_douyin_billboard_fetch_hot_total_high_play_list_post
type BodyFetchHotTotalHighPlayListApiV1DouyinBillboardFetchHotTotalHighPlayListPost struct {
	// Page，页码
	Page int32 `json:"page,omitempty" xml:"page"`
	// Page Size，每页数量
	PageSize int32 `json:"page_size,omitempty" xml:"page_size"`
	// Date Window，时间窗口，1 按小时 2 按天
	DateWindow int32 `json:"date_window,omitempty" xml:"date_window"`
	// Tags，子级垂类标签，空则为全部，多个标签需传入{\"value\": \"{顶级垂类标签id}\", \"children\": [{\"value\": \"{子级垂类标签id}\"}, {\"value\": \"{子级垂类标签id}\"}]}
	Tags []map[string]interface{} `json:"tags,omitempty" xml:"tags"`
}
