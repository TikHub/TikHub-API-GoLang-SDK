# BodyFetchHotAccountListApiV1DouyinBillboardFetchHotAccountListPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateWindow** | **string** | Date Window，时间窗口，格式 小时，默认24小时 | [optional] [default to 24]
**PageNum** | **string** | Page Num，页码，默认1 | [optional] [default to 1]
**PageSize** | **string** | Page Size，每页数量，默认10 | [optional] [default to 10]
**QueryTag** | [**map[string]interface{}**](.md) | Query Tag，子级垂类标签，空则为全部，多个标签需传入{\&quot;value\&quot;: \&quot;{顶级垂类标签id}\&quot;, \&quot;children\&quot;: [{\&quot;value\&quot;: \&quot;{子级垂类标签id}\&quot;}, {\&quot;value\&quot;: \&quot;{子级垂类标签id}\&quot;}]} | [optional] [default to {}]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


