# {{classname}}

All URIs are relative to *https://api.jumpcloud.com/insights/directory/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventCountGet**](DefaultApi.md#EventCountGet) | **Post** /events/count | Count Events
[**EventDistinctGet**](DefaultApi.md#EventDistinctGet) | **Post** /events/distinct | Query event distinct field values
[**EventGet**](DefaultApi.md#EventGet) | **Post** /events | Query Events
[**EventIntervalGet**](DefaultApi.md#EventIntervalGet) | **Post** /events/interval | Query event counts by bucketed by a time interval

# **EventCountGet**
> EventCount EventCountGet(ctx, body)
Count Events

Query the API for a count of matching events #### Sample Request ``` curl -X POST 'https://api.jumpcloud.com/insights/directory/v1/events/count' -H 'Content-Type: application/json' -H 'x-api-key: REPLACE_KEY_VALUE' --data '{\"service\": [\"all\"], \"start_time\": \"2021-07-14T23:00:00Z\", \"end_time\": \"2021-07-28T14:00:00Z\", \"sort\": \"DESC\", \"fields\": [\"timestamp\", \"event_type\", \"initiated_by\", \"success\", \"client_ip\", \"provider\", \"organization\"]}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventQuery**](EventQuery.md)| JSON event query body | 

### Return type

[**EventCount**](EventCount.md)

### Authorization

[x-api-key](../README.md#x-api-key), [x-org-id](../README.md#x-org-id)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventDistinctGet**
> InlineResponse200 EventDistinctGet(ctx, body)
Query event distinct field values

Query the API for a list of distinct values for a field #### Sample Request ``` curl -X POST 'https://api.jumpcloud.com/insights/directory/v1/events/distinct' -H 'Content-Type: application/json' -H 'x-api-key: REPLACE_KEY_VALUE' --data '{\"service\": [\"all\"], \"start_time\": \"2021-07-14T23:00:00Z\", \"end_time\": \"2021-07-28T14:00:00Z\", \"sort\": \"DESC\", \"field\": \"event_type\"}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventDistinctQuery**](EventDistinctQuery.md)| JSON event distinct query body | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[x-api-key](../README.md#x-api-key), [x-org-id](../README.md#x-org-id)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventGet**
> []ModelMap EventGet(ctx, body)
Query Events

Query the API for Directory Insights events #### Sample Request ``` curl -X POST 'https://api.jumpcloud.com/insights/directory/v1/events' -H 'Content-Type: application/json' -H 'x-api-key: REPLACE_KEY_VALUE' --data '{\"service\": [\"all\"], \"start_time\": \"2021-07-14T23:00:00Z\", \"end_time\": \"2021-07-28T14:00:00Z\", \"sort\": \"DESC\", \"fields\": [\"timestamp\", \"event_type\", \"initiated_by\", \"success\", \"client_ip\", \"provider\", \"organization\"]}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventQuery**](EventQuery.md)| JSON event query body | 

### Return type

[**[]ModelMap**](map.md)

### Authorization

[x-api-key](../README.md#x-api-key), [x-org-id](../README.md#x-org-id)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EventIntervalGet**
> InlineResponse2001 EventIntervalGet(ctx, body)
Query event counts by bucketed by a time interval

Query the API for a list of counts by time interval #### Sample Request ``` curl -X POST 'https://api.jumpcloud.com/insights/directory/v1/events/interval' -H 'Content-Type: application/json' -H 'x-api-key: REPLACE_KEY_VALUE' --data '{\"service\": [\"all\"], \"start_time\": \"2021-07-14T23:00:00Z\", \"end_time\": \"2021-07-28T14:00:00Z\", \"timezone\": \"-0500\", \"interval_unit\": \"h\", \"interval_value\": \"2\"}' ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventIntervalQuery**](EventIntervalQuery.md)| JSON event interval query body | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[x-api-key](../README.md#x-api-key), [x-org-id](../README.md#x-org-id)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

