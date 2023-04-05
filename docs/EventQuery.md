# EventQuery

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | [**time.Time**](time.Time.md) | optional query end time, UTC in RFC3339 format | [optional] [default to null]
**Service** | **[]string** | service name to query. | [default to null]
**StartTime** | [**time.Time**](time.Time.md) | query start time, UTC in RFC3339 format | [default to null]
**Fields** | **[]string** | optional list of fields to return from query | [optional] [default to null]
**Limit** | **int64** | Max number of rows to return | [optional] [default to null]
**Q** | **string** | optional string for specifying a full text query | [optional] [default to null]
**SearchAfter** | **[]string** | Specific query to search after, see x-* response headers for next values | [optional] [default to null]
**SearchTerm** | [***SearchTerm**](SearchTerm.md) |  | [optional] [default to null]
**Sort** | **string** | ASC or DESC order for timestamp | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

