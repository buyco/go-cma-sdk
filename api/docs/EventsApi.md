# \EventsApi

All URIs are relative to *https://apis.cma-cgm.net/operation/trackandtrace/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMoveOnCommercialCycle**](EventsApi.md#GetMoveOnCommercialCycle) | **Get** /events/{trackingReference} | Find Commercial events from unique tracking reference.
[**SearchMoveOnCommercialCycle**](EventsApi.md#SearchMoveOnCommercialCycle) | **Get** /events | Find Commercial events.



## GetMoveOnCommercialCycle

> []SearchMoveOnCommercialCycle200ResponseInner GetMoveOnCommercialCycle(ctx, trackingReference).BehalfOf(behalfOf).Limit(limit).Cursor(cursor).Execute()

Find Commercial events from unique tracking reference.



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
    trackingReference := "trackingReference_example" // string | Shipment reference or Equipment identifier
    behalfOf := "00002334567" // string | (Mandatory if you are a Third Party). This field specifies the end customer code you request a rate for. Use our referential API Partner to check if the end customer exists and to get its Partner ID code (optional)
    limit := int32(100) // int32 | Maximum number of items to return. (optional) (default to 100)
    cursor := "fE9mZnNldHw9MTAmbGltaXQ9MTA=" // string | A server generated value to specify a specific point in a collection result, used for pagination. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.GetMoveOnCommercialCycle(context.Background(), trackingReference).BehalfOf(behalfOf).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.GetMoveOnCommercialCycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMoveOnCommercialCycle`: []SearchMoveOnCommercialCycle200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.GetMoveOnCommercialCycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackingReference** | **string** | Shipment reference or Equipment identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMoveOnCommercialCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **behalfOf** | **string** | (Mandatory if you are a Third Party). This field specifies the end customer code you request a rate for. Use our referential API Partner to check if the end customer exists and to get its Partner ID code | 
 **limit** | **int32** | Maximum number of items to return. | [default to 100]
 **cursor** | **string** | A server generated value to specify a specific point in a collection result, used for pagination. | 

### Return type

[**[]SearchMoveOnCommercialCycle200ResponseInner**](SearchMoveOnCommercialCycle200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchMoveOnCommercialCycle

> []SearchMoveOnCommercialCycle200ResponseInner SearchMoveOnCommercialCycle(ctx).EventType(eventType).ShipmentEventTypeCode(shipmentEventTypeCode).DocumentTypeCode(documentTypeCode).CarrierBookingReference(carrierBookingReference).TransportDocumentReference(transportDocumentReference).TransportEventTypeCode(transportEventTypeCode).TransportCallID(transportCallID).VesselIMONumber(vesselIMONumber).ExportVoyageNumber(exportVoyageNumber).CarrierServiceCode(carrierServiceCode).UNLocationCode(uNLocationCode).EquipmentEventTypeCode(equipmentEventTypeCode).EquipmentReference(equipmentReference).EventCreatedDateTime(eventCreatedDateTime).EventDateTime(eventDateTime).BehalfOf(behalfOf).Limit(limit).Cursor(cursor).Execute()

Find Commercial events.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    eventType := []string{"EventType_example"} // []string | The type of event(s) to filter by. Possible values are - SHIPMENT (Shipment events) - TRANSPORT (Transport events) - EQUIPMENT (Equipment events) It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example eventType=SHIPMENT,EQUIPMENT matches both Shipment- and Equipment-events. Default value is all event types. (optional)
    shipmentEventTypeCode := []openapiclient.ShipmentEventTypeCode{openapiclient.shipmentEventTypeCode("RECE")} // []ShipmentEventTypeCode | The status of the document in the process to filter by. Possible values are - RECE (Received) - DRFT (Drafted) - PENA (Pending Approval) - PENU (Pending Update) - REJE (Rejected) - APPR (Approved) - ISSU (Issued) - SURR (Surrendered) - SUBM (Submitted) - VOID (Void) - CONF (Confirmed) - REQS (Requested) - CMPL (Completed) - HOLD (On Hold) - RELS (Released) It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example <i>shipmentEventTypeCode=RECE,DRFT</i>  Matches <b>both</b> Received (RECE) and Drafted (DRFT) shipment events. Default is all shipmentEventTypeCodes. This filter is only relevant when filtering on ShipmentEvents (optional) (default to ["RECE","DRFT","PENA","PENU","REJE","APPR","ISSU","SURR","SUBM","VOID","CONF","REQS","CMPL","HOLD","RELS"])
    documentTypeCode := []openapiclient.DocumentTypeCode{openapiclient.documentTypeCode("CBR")} // []DocumentTypeCode | The documentTypeCode to filter by. Possible values are - CBR (Carrier Booking Request Reference) - BKG (Booking) - SHI (Shipping Instruction) - SRM (Shipment Release Message) - TRD (Transport Document) - ARN (Arrival Notice) - VGM (Verified Gross Mass) - CAS (Cargo Survey) - CUS (Customs Inspection) - DGD (Dangerous Goods Declaration) - OOG (Out of Gauge) It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example <i>documentTypeCode=SHI,TRD</i> Matches <b>both</b> ShippingInstruction (SHI) and TransportDocument (TRD) shipment events. Default is all documentTypeCodes. This filter is only relevant when filtering on ShipmentEvents (optional) (default to ["CBR","BKG","SHI","SRM","TRD","ARN","VGM","CAS","CUS","DGD","OOG"])
    carrierBookingReference := "carrierBookingReference_example" // string | A set of unique characters provided by carrier to identify a booking. Specifying this filter will only return events related to this particular carrierBookingReference. (optional)
    transportDocumentReference := "transportDocumentReference_example" // string | A unique number reference allocated by the shipping line to the transport document and the main number used for the tracking of the status of the shipment. Specifying this filter will only return events related to this particular transportDocumentReference (optional)
    transportEventTypeCode := []openapiclient.TransportEventTypeCode{openapiclient.transportEventTypeCode("ARRI")} // []TransportEventTypeCode | Identifier for type of Transport event to filter by - ARRI (Arrived) - DEPA (Departed) It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example <i>transportEventTypeCode=ARRI,DEPA</i> matches <b>both</b> Arrived (ARRI) and Departed (DEPA) transport events. Default is all transportEventTypeCodes. This filter is only relevant when filtering on TransportEvents (optional) (default to ["ARRI","DEPA"])
    transportCallID := "transportCallID_example" // string | ID uniquely identifying a transport call, to filter events by. Specifying this filter will only return events related to this particular transportCallID (optional)
    vesselIMONumber := "vesselIMONumber_example" // string | The identifier of vessel for which schedule details are published. Depending on schedule type, this may not be available yet. Specifying this filter will only return events related to this particular vesselIMONumber. (optional)
    exportVoyageNumber := "exportVoyageNumber_example" // string | Filter on the vessel operator-specific identifier of the export Voyage. Specifying this filter will only return events related to this particular exportVoyageNumber. (optional)
    carrierServiceCode := "carrierServiceCode_example" // string | Filter on the carrier specific identifier of the service. Specifying this filter will only return events related to this particular carrierServiceCode. (optional)
    uNLocationCode := "uNLocationCode_example" // string | The UN Location code specifying where the place is located. Specifying this filter will only return events related to this particular UN Location code. (optional)
    equipmentEventTypeCode := []openapiclient.EquipmentEventTypeCode{openapiclient.equipmentEventTypeCode("LOAD")} // []EquipmentEventTypeCode | Unique identifier for equipmentEventTypeCode. - LOAD (Loaded) - DISC (Discharged) - GTIN (Gated in) - GTOT (Gated out) - STUF (Stuffed) - STRP (Stripped) - PICK (Pick-up) - DROP (Drop-off) - INSP (Inspected) - RSEA (Resealed) - RMVD (Removed) It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example <i>equipmentEventTypeCode=GTIN,GTOT</i> matches <b>both</b> Gated in (GTIN) and Gated out (GTOT) equipment events. Default is all equipmentEventTypeCodes. This filter is only relevant when filtering on EquipmentEvents (optional) (default to ["LOAD","DISC","GTIN","GTOT","STUF","STRP","PICK","DROP","INSP","RSEA","RMVD"])
    equipmentReference := "equipmentReference_example" // string | Will filter by the unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. Specifying this filter will only return events related to this particular equipmentReference (optional)
    eventCreatedDateTime := time.Now() // time.Time | Limit the result based on the creating date of the event. It is possible to use operators on this query parameter. This is done by adding a colon followed by an operator at the end of the queryParameterName (before the =) <i>eventCreatedDateTime<b>&#58;gte</b>=2021-04-01T14&#58;12&#58;56+01&#58;00</i> would result in all events created &#8805; 2021-04-01T14&#58;12&#58;56+01&#58;00 The following operators are supported - &#58;gte (&#8805; Greater than or equal) - &#58;gt (&#62; Greater than) - &#58;lte (&#8804; Less than or equal) - &#58;lt (&#60; Less than) - &#58;eq (&#61; Equal to) If no operator is provided, a <b>strictly equal</b> is used (this is equivalent to <b>&#58;eq</b> operator). (optional)
    eventDateTime := time.Now() // time.Time | not DCSA - when search is done by container reference , allow to search the last container cycle happening before the provided date (optional)
    behalfOf := "00002334567" // string | (Mandatory if you are a Third Party). This field specifies the end customer code you request a rate for. Use our referential API Partner to check if the end customer exists and to get its Partner ID code -- Not a standard DCSA attribute (optional)
    limit := int32(100) // int32 | Maximum number of items to return. (optional) (default to 100)
    cursor := "fE9mZnNldHw9MTAmbGltaXQ9MTA=" // string | A server generated value to specify a specific point in a collection result, used for pagination. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.SearchMoveOnCommercialCycle(context.Background()).EventType(eventType).ShipmentEventTypeCode(shipmentEventTypeCode).DocumentTypeCode(documentTypeCode).CarrierBookingReference(carrierBookingReference).TransportDocumentReference(transportDocumentReference).TransportEventTypeCode(transportEventTypeCode).TransportCallID(transportCallID).VesselIMONumber(vesselIMONumber).ExportVoyageNumber(exportVoyageNumber).CarrierServiceCode(carrierServiceCode).UNLocationCode(uNLocationCode).EquipmentEventTypeCode(equipmentEventTypeCode).EquipmentReference(equipmentReference).EventCreatedDateTime(eventCreatedDateTime).EventDateTime(eventDateTime).BehalfOf(behalfOf).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.SearchMoveOnCommercialCycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchMoveOnCommercialCycle`: []SearchMoveOnCommercialCycle200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.SearchMoveOnCommercialCycle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchMoveOnCommercialCycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventType** | **[]string** | The type of event(s) to filter by. Possible values are - SHIPMENT (Shipment events) - TRANSPORT (Transport events) - EQUIPMENT (Equipment events) It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example eventType&#x3D;SHIPMENT,EQUIPMENT matches both Shipment- and Equipment-events. Default value is all event types. | 
 **shipmentEventTypeCode** | [**[]ShipmentEventTypeCode**](ShipmentEventTypeCode.md) | The status of the document in the process to filter by. Possible values are - RECE (Received) - DRFT (Drafted) - PENA (Pending Approval) - PENU (Pending Update) - REJE (Rejected) - APPR (Approved) - ISSU (Issued) - SURR (Surrendered) - SUBM (Submitted) - VOID (Void) - CONF (Confirmed) - REQS (Requested) - CMPL (Completed) - HOLD (On Hold) - RELS (Released) It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example &lt;i&gt;shipmentEventTypeCode&#x3D;RECE,DRFT&lt;/i&gt;  Matches &lt;b&gt;both&lt;/b&gt; Received (RECE) and Drafted (DRFT) shipment events. Default is all shipmentEventTypeCodes. This filter is only relevant when filtering on ShipmentEvents | [default to [&quot;RECE&quot;,&quot;DRFT&quot;,&quot;PENA&quot;,&quot;PENU&quot;,&quot;REJE&quot;,&quot;APPR&quot;,&quot;ISSU&quot;,&quot;SURR&quot;,&quot;SUBM&quot;,&quot;VOID&quot;,&quot;CONF&quot;,&quot;REQS&quot;,&quot;CMPL&quot;,&quot;HOLD&quot;,&quot;RELS&quot;]]
 **documentTypeCode** | [**[]DocumentTypeCode**](DocumentTypeCode.md) | The documentTypeCode to filter by. Possible values are - CBR (Carrier Booking Request Reference) - BKG (Booking) - SHI (Shipping Instruction) - SRM (Shipment Release Message) - TRD (Transport Document) - ARN (Arrival Notice) - VGM (Verified Gross Mass) - CAS (Cargo Survey) - CUS (Customs Inspection) - DGD (Dangerous Goods Declaration) - OOG (Out of Gauge) It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example &lt;i&gt;documentTypeCode&#x3D;SHI,TRD&lt;/i&gt; Matches &lt;b&gt;both&lt;/b&gt; ShippingInstruction (SHI) and TransportDocument (TRD) shipment events. Default is all documentTypeCodes. This filter is only relevant when filtering on ShipmentEvents | [default to [&quot;CBR&quot;,&quot;BKG&quot;,&quot;SHI&quot;,&quot;SRM&quot;,&quot;TRD&quot;,&quot;ARN&quot;,&quot;VGM&quot;,&quot;CAS&quot;,&quot;CUS&quot;,&quot;DGD&quot;,&quot;OOG&quot;]]
 **carrierBookingReference** | **string** | A set of unique characters provided by carrier to identify a booking. Specifying this filter will only return events related to this particular carrierBookingReference. | 
 **transportDocumentReference** | **string** | A unique number reference allocated by the shipping line to the transport document and the main number used for the tracking of the status of the shipment. Specifying this filter will only return events related to this particular transportDocumentReference | 
 **transportEventTypeCode** | [**[]TransportEventTypeCode**](TransportEventTypeCode.md) | Identifier for type of Transport event to filter by - ARRI (Arrived) - DEPA (Departed) It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example &lt;i&gt;transportEventTypeCode&#x3D;ARRI,DEPA&lt;/i&gt; matches &lt;b&gt;both&lt;/b&gt; Arrived (ARRI) and Departed (DEPA) transport events. Default is all transportEventTypeCodes. This filter is only relevant when filtering on TransportEvents | [default to [&quot;ARRI&quot;,&quot;DEPA&quot;]]
 **transportCallID** | **string** | ID uniquely identifying a transport call, to filter events by. Specifying this filter will only return events related to this particular transportCallID | 
 **vesselIMONumber** | **string** | The identifier of vessel for which schedule details are published. Depending on schedule type, this may not be available yet. Specifying this filter will only return events related to this particular vesselIMONumber. | 
 **exportVoyageNumber** | **string** | Filter on the vessel operator-specific identifier of the export Voyage. Specifying this filter will only return events related to this particular exportVoyageNumber. | 
 **carrierServiceCode** | **string** | Filter on the carrier specific identifier of the service. Specifying this filter will only return events related to this particular carrierServiceCode. | 
 **uNLocationCode** | **string** | The UN Location code specifying where the place is located. Specifying this filter will only return events related to this particular UN Location code. | 
 **equipmentEventTypeCode** | [**[]EquipmentEventTypeCode**](EquipmentEventTypeCode.md) | Unique identifier for equipmentEventTypeCode. - LOAD (Loaded) - DISC (Discharged) - GTIN (Gated in) - GTOT (Gated out) - STUF (Stuffed) - STRP (Stripped) - PICK (Pick-up) - DROP (Drop-off) - INSP (Inspected) - RSEA (Resealed) - RMVD (Removed) It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example &lt;i&gt;equipmentEventTypeCode&#x3D;GTIN,GTOT&lt;/i&gt; matches &lt;b&gt;both&lt;/b&gt; Gated in (GTIN) and Gated out (GTOT) equipment events. Default is all equipmentEventTypeCodes. This filter is only relevant when filtering on EquipmentEvents | [default to [&quot;LOAD&quot;,&quot;DISC&quot;,&quot;GTIN&quot;,&quot;GTOT&quot;,&quot;STUF&quot;,&quot;STRP&quot;,&quot;PICK&quot;,&quot;DROP&quot;,&quot;INSP&quot;,&quot;RSEA&quot;,&quot;RMVD&quot;]]
 **equipmentReference** | **string** | Will filter by the unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. Specifying this filter will only return events related to this particular equipmentReference | 
 **eventCreatedDateTime** | **time.Time** | Limit the result based on the creating date of the event. It is possible to use operators on this query parameter. This is done by adding a colon followed by an operator at the end of the queryParameterName (before the &#x3D;) &lt;i&gt;eventCreatedDateTime&lt;b&gt;&amp;#58;gte&lt;/b&gt;&#x3D;2021-04-01T14&amp;#58;12&amp;#58;56+01&amp;#58;00&lt;/i&gt; would result in all events created &amp;#8805; 2021-04-01T14&amp;#58;12&amp;#58;56+01&amp;#58;00 The following operators are supported - &amp;#58;gte (&amp;#8805; Greater than or equal) - &amp;#58;gt (&amp;#62; Greater than) - &amp;#58;lte (&amp;#8804; Less than or equal) - &amp;#58;lt (&amp;#60; Less than) - &amp;#58;eq (&amp;#61; Equal to) If no operator is provided, a &lt;b&gt;strictly equal&lt;/b&gt; is used (this is equivalent to &lt;b&gt;&amp;#58;eq&lt;/b&gt; operator). | 
 **eventDateTime** | **time.Time** | not DCSA - when search is done by container reference , allow to search the last container cycle happening before the provided date | 
 **behalfOf** | **string** | (Mandatory if you are a Third Party). This field specifies the end customer code you request a rate for. Use our referential API Partner to check if the end customer exists and to get its Partner ID code -- Not a standard DCSA attribute | 
 **limit** | **int32** | Maximum number of items to return. | [default to 100]
 **cursor** | **string** | A server generated value to specify a specific point in a collection result, used for pagination. | 

### Return type

[**[]SearchMoveOnCommercialCycle200ResponseInner**](SearchMoveOnCommercialCycle200ResponseInner.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

