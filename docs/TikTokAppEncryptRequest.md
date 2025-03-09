# TikTokAppEncryptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Url，需要加密的URL/URL to be encrypted | [optional] [default to https://api16-normal-useast5.tiktokv.us/tiktok/v1/upvote/item/list?user_id=6726034365602628610&offset=0&count=21&scene=0&iid=7425045478163400491&device_id=7349721034012280362&ac=WIFI&channel=googleplay&aid=1233&app_name=musical_ly&version_code=360704&version_name=36.7.4&device_platform=android&os=android&ab_version=36.7.4&ssmix=a&device_type=Pixel+6+Pro&device_brand=google&language=zh&os_api=33&os_version=13&openudid=711192517a8bbf03&manifest_version_code=2023607040&resolution=1440*2891&dpi=560&update_version_code=2023607040&_rticket=1728977220468&is_pad=0&app_type=normal&sys_region=CN&last_install_time=1728977141&timezone_name=America%2FLos_Angeles&app_language=zh-Hans&ac2=wifi5g&uoo=0&op_region=CN&timezone_offset=-28800&build_number=36.7.4&host_abi=arm64-v8a&locale=zh-Hans&region=CN&content_language=en%2C&ts=1728977220&cdid=aa21524b-8633-49ca-8e6e-3275fe1672db]
**Data** | **string** | Data，如果有POST请求，请填写POST请求的数据参与加密计算/If there is a POST request, please fill in the data of the POST request to participate in the encryption calculation | [optional] [default to ]
**DeviceInfo** | [**map[string]interface{}**](.md) | Device Info，设备信息，可选参数，如果不填写则使用默认设备信息/Device information, optional parameter, if not filled in, the default device information is used | [optional] [default to {"aid":"1233","cdid":"b820f79c-c74a-47b0-912f-ee3002ce60dc","channel":"googleplay","cookies":{},"device_brand":"HONOR","device_id":"7423364899755607598","device_manufacturer":"HUAWEI","device_model":"HONOR V30","device_platform":"android","device_type":"OXF-AN00","dpi":480,"host_abi":"arm64-v8a","iid":"7423365134775469866","lanusk":"","manifest_version_code":"2023604040","mc":"7E:EE:BA:BC:5E:40","mssdk_token":"","openudid":"63401ab5140125d1","os_api":29,"os_version":"10","resolution":"2400*1080","rom":"11.0.0.185C00","rom_version":"OXF-AN00-user 11.0.0 HUAWEIOXF-AN00 185-CHN-LGRP3 release-keys","server_time":1728386909,"ua":"com.zhiliaoapp.musically/2023604040 (Linux; U; Android 10; zh_CN; OXF-AN00; Build/185-CHN-LGRP3;tt-ok/3.12.13.4-tiktok)","update_version_code":"2023604040","uuid":"350244698061054","version_code":"360404","version_name":"36.4.4"}]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


