# BodyFetchHotTotalHighTopicListApiV1DouyinBillboardFetchHotTotalHighTopicListPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int32** | Page，页码 | [optional] [default to 1]
**PageSize** | **int32** | Page Size，每页数量 | [optional] [default to 10]
**DateWindow** | **int32** | Date Window，时间窗口，1 按小时 2 按天 | [optional] [default to 24]
**Tags** | **[]map[string]interface{}** | Tags，子级垂类标签，空则为全部，多个标签需传入{\&quot;value\&quot;: \&quot;{顶级垂类标签id}\&quot;, \&quot;children\&quot;: [{\&quot;value\&quot;: \&quot;{子级垂类标签id}\&quot;}, {\&quot;value\&quot;: \&quot;{子级垂类标签id}\&quot;}]} | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


