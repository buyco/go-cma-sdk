# \SubscriptionRulesApi

All URIs are relative to *https://apis.cma-cgm.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscriptionRule**](SubscriptionRulesApi.md#CreateSubscriptionRule) | **Post** /eventhub/subscription/v1/subscriptionRules | Create an event subscription rule
[**DeleteSubscriptionRule**](SubscriptionRulesApi.md#DeleteSubscriptionRule) | **Delete** /eventhub/subscription/v1/subscriptionRules | Delete a subscription rule
[**ListSubscriptionRules**](SubscriptionRulesApi.md#ListSubscriptionRules) | **Get** /eventhub/subscription/v1/subscriptionRules | List event subscription rules



## CreateSubscriptionRule

> string CreateSubscriptionRule(ctx).SubscriptionRule(subscriptionRule).Execute()

Create an event subscription rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subscriptionRule := *openapiclient.NewSubscriptionRule() // SubscriptionRule | The subscription rule object

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionRulesApi.CreateSubscriptionRule(context.Background()).SubscriptionRule(subscriptionRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionRulesApi.CreateSubscriptionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscriptionRule`: string
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionRulesApi.CreateSubscriptionRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionRule** | [**SubscriptionRule**](SubscriptionRule.md) | The subscription rule object | 

### Return type

**string**

### Authorization

[oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscriptionRule

> string DeleteSubscriptionRule(ctx).SubscriptionRuleToDelete(subscriptionRuleToDelete).Execute()

Delete a subscription rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subscriptionRuleToDelete := *openapiclient.NewSubscriptionRuleToDelete() // SubscriptionRuleToDelete | the subscription rule to delete for a subscriber

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionRulesApi.DeleteSubscriptionRule(context.Background()).SubscriptionRuleToDelete(subscriptionRuleToDelete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionRulesApi.DeleteSubscriptionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubscriptionRule`: string
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionRulesApi.DeleteSubscriptionRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionRuleToDelete** | [**SubscriptionRuleToDelete**](SubscriptionRuleToDelete.md) | the subscription rule to delete for a subscriber | 

### Return type

**string**

### Authorization

[oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptionRules

> []SubscriptionRule ListSubscriptionRules(ctx).Range_(range_).Execute()

List event subscription rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    range_ := "range__example" // string | Pagination parameter. default 0-49. Maximum rows to return 50 (optional) (default to "0-49")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionRulesApi.ListSubscriptionRules(context.Background()).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionRulesApi.ListSubscriptionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubscriptionRules`: []SubscriptionRule
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionRulesApi.ListSubscriptionRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubscriptionRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Pagination parameter. default 0-49. Maximum rows to return 50 | [default to &quot;0-49&quot;]

### Return type

[**[]SubscriptionRule**](SubscriptionRule.md)

### Authorization

[oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

