# EventIntervalQuery

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | [**time.Time**](time.Time.md) | optional query end time, UTC in RFC3339 format | [optional] [default to null]
**Service** | **[]string** | service name to query. | [default to null]
**StartTime** | [**time.Time**](time.Time.md) | query start time, UTC in RFC3339 format | [default to null]
**IntervalUnit** | **string** |  | [default to null]
**IntervalValue** | **string** | Interval Value. This specifies how many units you want to bucket the event counts by | [optional] [default to null]
**Q** | **string** | optional string for specifying a full text query | [optional] [default to null]
**SearchTerm** | [***SearchTerm**](SearchTerm.md) |  | [optional] [default to null]
**Timezone** | **string** | TimeZone. Specify the timezone in which the user is in optional | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

