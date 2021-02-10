# \DefaultApi

All URIs are relative to *https://digital-services-apis.cma-cgm.com/logistic/tracking/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLastCycleMoveFromEquipment**](DefaultApi.md#GetLastCycleMoveFromEquipment) | **Get** /equipments/{eqpId}/moves/lastCycle | all moves information related to a container of its last cycle
[**GetMoveOnCommercialCycleFromShipment**](DefaultApi.md#GetMoveOnCommercialCycleFromShipment) | **Get** /shipments/{shipmentRef}/equipments/moves/commercialCycle | all moves information related to all containers of a shipment
[**GetMoveOnCommercialCycleFromShipmentAndEquipmentId**](DefaultApi.md#GetMoveOnCommercialCycleFromShipmentAndEquipmentId) | **Get** /shipments/{shipmentRef}/equipments/{eqpId}/moves/commercialCycle | all moves information related to a container of a shipment



## GetLastCycleMoveFromEquipment

> Shipment GetLastCycleMoveFromEquipment(ctx, eqpId).Execute()

all moves information related to a container of its last cycle



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
    eqpId := "eqpId_example" // string | equipment id (container reference)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetLastCycleMoveFromEquipment(context.Background(), eqpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetLastCycleMoveFromEquipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLastCycleMoveFromEquipment`: Shipment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetLastCycleMoveFromEquipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eqpId** | **string** | equipment id (container reference) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastCycleMoveFromEquipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Shipment**](Shipment.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMoveOnCommercialCycleFromShipment

> Shipment GetMoveOnCommercialCycleFromShipment(ctx, shipmentRef).Execute()

all moves information related to all containers of a shipment



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
    shipmentRef := "shipmentRef_example" // string | shipment reference (booking reference or BL number)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetMoveOnCommercialCycleFromShipment(context.Background(), shipmentRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMoveOnCommercialCycleFromShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMoveOnCommercialCycleFromShipment`: Shipment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMoveOnCommercialCycleFromShipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentRef** | **string** | shipment reference (booking reference or BL number) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMoveOnCommercialCycleFromShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Shipment**](Shipment.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMoveOnCommercialCycleFromShipmentAndEquipmentId

> Shipment GetMoveOnCommercialCycleFromShipmentAndEquipmentId(ctx, shipmentRef, eqpId).Execute()

all moves information related to a container of a shipment



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
    shipmentRef := "shipmentRef_example" // string | shipment reference (booking reference or BL number)
    eqpId := "eqpId_example" // string | equipment id (container reference)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetMoveOnCommercialCycleFromShipmentAndEquipmentId(context.Background(), shipmentRef, eqpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMoveOnCommercialCycleFromShipmentAndEquipmentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMoveOnCommercialCycleFromShipmentAndEquipmentId`: Shipment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMoveOnCommercialCycleFromShipmentAndEquipmentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentRef** | **string** | shipment reference (booking reference or BL number) | 
**eqpId** | **string** | equipment id (container reference) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMoveOnCommercialCycleFromShipmentAndEquipmentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Shipment**](Shipment.md)

### Authorization

[API Key](../README.md#API Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

