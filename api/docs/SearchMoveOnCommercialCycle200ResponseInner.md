# SearchMoveOnCommercialCycle200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | Pointer to **string** | The unique identifier for the event (the message - not the source). &lt;b&gt;NB&lt;/b&gt;&amp;#58; This field should be considered Metadata | [optional] 
**EventCreatedDateTime** | **time.Time** | The timestamp of when the event was created. &lt;b&gt;NB&lt;/b&gt;&amp;#58; This field should be considered Metadata | 
**EventType** | **string** |  | 
**EventClassifierCode** | **string** | Code for the event classifier can be - PLN (Planned) - ACT (Actual) - EST (Estimated) | 
**EventDateTime** | **time.Time** | The local date and time, where the event took place or when the event will take place, in ISO 8601 format. | 
**CarrierSpecificData** | Pointer to [**CarrierSpecificData**](CarrierSpecificData.md) |  | [optional] 
**TransportEventTypeCode** | [**TransportEventTypeCode**](TransportEventTypeCode.md) |  | 
**DelayReasonCode** | Pointer to **string** | Reason code for the delay. The SMDG-Delay-Reason-Codes are used for this attribute. The code list can be found at http://www.smdg.org/smdg-code-lists/ | [optional] 
**ChangeRemark** | Pointer to **string** | Free text information provided by the vessel operator regarding the reasons for the change in schedule and/or plans to mitigate schedule slippage. | [optional] 
**TransportCall** | [**TransportCall**](TransportCall.md) |  | 
**DocumentReferences** | Pointer to [**[]DocumentReferencesInner**](DocumentReferencesInner.md) | An optional list of key-value (documentReferenceType-documentReferenceValue) pairs representing links to objects relevant to the event. The &lt;b&gt;documentReferenceType&lt;/b&gt;-field is used to describe where the &lt;b&gt;documentReferenceValue&lt;/b&gt;-field is pointing to. | [optional] 
**References** | Pointer to [**[]Reference**](Reference.md) |  | [optional] 
**ShipmentEventTypeCode** | [**ShipmentEventTypeCode**](ShipmentEventTypeCode.md) |  | 
**DocumentID** | **string** | The id of the object defined by the documentTypeCode. | 
**DocumentTypeCode** | [**DocumentTypeCode**](DocumentTypeCode.md) |  | 
**Reason** | Pointer to **string** | Reason field in a Shipment event. This field can be used to explain why a specific event has been sent. | [optional] 
**ShipmentID** | Pointer to **string** | The identifier for a shipment | [optional] 
**EquipmentEventTypeCode** | [**EquipmentEventTypeCode**](EquipmentEventTypeCode.md) |  | 
**EquipmentReference** | Pointer to **string** | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation | [optional] 
**ISOEquipmentCode** | Pointer to **string** | Unique code for the different equipment size/type used for transporting commodities. The code is a concatenation of ISO Equipment Size Code and ISO Equipment Type Code A and follows the ISO 6346 standard. | [optional] 
**EmptyIndicatorCode** | [**EmptyIndicatorCode**](EmptyIndicatorCode.md) |  | 
**EventLocation** | Pointer to [**Location**](Location.md) |  | [optional] 
**Seals** | Pointer to [**[]Seal**](Seal.md) |  | [optional] 

## Methods

### NewSearchMoveOnCommercialCycle200ResponseInner

`func NewSearchMoveOnCommercialCycle200ResponseInner(eventCreatedDateTime time.Time, eventType string, eventClassifierCode string, eventDateTime time.Time, transportEventTypeCode TransportEventTypeCode, transportCall TransportCall, shipmentEventTypeCode ShipmentEventTypeCode, documentID string, documentTypeCode DocumentTypeCode, equipmentEventTypeCode EquipmentEventTypeCode, emptyIndicatorCode EmptyIndicatorCode, ) *SearchMoveOnCommercialCycle200ResponseInner`

NewSearchMoveOnCommercialCycle200ResponseInner instantiates a new SearchMoveOnCommercialCycle200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchMoveOnCommercialCycle200ResponseInnerWithDefaults

`func NewSearchMoveOnCommercialCycle200ResponseInnerWithDefaults() *SearchMoveOnCommercialCycle200ResponseInner`

NewSearchMoveOnCommercialCycle200ResponseInnerWithDefaults instantiates a new SearchMoveOnCommercialCycle200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventID

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetEventID(v string)`

SetEventID sets EventID field to given value.

### HasEventID

`func (o *SearchMoveOnCommercialCycle200ResponseInner) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### GetEventCreatedDateTime

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEventCreatedDateTime() time.Time`

GetEventCreatedDateTime returns the EventCreatedDateTime field if non-nil, zero value otherwise.

### GetEventCreatedDateTimeOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEventCreatedDateTimeOk() (*time.Time, bool)`

GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCreatedDateTime

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetEventCreatedDateTime(v time.Time)`

SetEventCreatedDateTime sets EventCreatedDateTime field to given value.


### GetEventType

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventClassifierCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEventClassifierCode() string`

GetEventClassifierCode returns the EventClassifierCode field if non-nil, zero value otherwise.

### GetEventClassifierCodeOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEventClassifierCodeOk() (*string, bool)`

GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClassifierCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetEventClassifierCode(v string)`

SetEventClassifierCode sets EventClassifierCode field to given value.


### GetEventDateTime

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.


### GetCarrierSpecificData

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetCarrierSpecificData() CarrierSpecificData`

GetCarrierSpecificData returns the CarrierSpecificData field if non-nil, zero value otherwise.

### GetCarrierSpecificDataOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetCarrierSpecificDataOk() (*CarrierSpecificData, bool)`

GetCarrierSpecificDataOk returns a tuple with the CarrierSpecificData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierSpecificData

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetCarrierSpecificData(v CarrierSpecificData)`

SetCarrierSpecificData sets CarrierSpecificData field to given value.

### HasCarrierSpecificData

`func (o *SearchMoveOnCommercialCycle200ResponseInner) HasCarrierSpecificData() bool`

HasCarrierSpecificData returns a boolean if a field has been set.

### GetTransportEventTypeCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetTransportEventTypeCode() TransportEventTypeCode`

GetTransportEventTypeCode returns the TransportEventTypeCode field if non-nil, zero value otherwise.

### GetTransportEventTypeCodeOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetTransportEventTypeCodeOk() (*TransportEventTypeCode, bool)`

GetTransportEventTypeCodeOk returns a tuple with the TransportEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportEventTypeCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetTransportEventTypeCode(v TransportEventTypeCode)`

SetTransportEventTypeCode sets TransportEventTypeCode field to given value.


### GetDelayReasonCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetDelayReasonCode() string`

GetDelayReasonCode returns the DelayReasonCode field if non-nil, zero value otherwise.

### GetDelayReasonCodeOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetDelayReasonCodeOk() (*string, bool)`

GetDelayReasonCodeOk returns a tuple with the DelayReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayReasonCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetDelayReasonCode(v string)`

SetDelayReasonCode sets DelayReasonCode field to given value.

### HasDelayReasonCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) HasDelayReasonCode() bool`

HasDelayReasonCode returns a boolean if a field has been set.

### GetChangeRemark

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetChangeRemark() string`

GetChangeRemark returns the ChangeRemark field if non-nil, zero value otherwise.

### GetChangeRemarkOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetChangeRemarkOk() (*string, bool)`

GetChangeRemarkOk returns a tuple with the ChangeRemark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRemark

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetChangeRemark(v string)`

SetChangeRemark sets ChangeRemark field to given value.

### HasChangeRemark

`func (o *SearchMoveOnCommercialCycle200ResponseInner) HasChangeRemark() bool`

HasChangeRemark returns a boolean if a field has been set.

### GetTransportCall

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetTransportCall() TransportCall`

GetTransportCall returns the TransportCall field if non-nil, zero value otherwise.

### GetTransportCallOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetTransportCallOk() (*TransportCall, bool)`

GetTransportCallOk returns a tuple with the TransportCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCall

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetTransportCall(v TransportCall)`

SetTransportCall sets TransportCall field to given value.


### GetDocumentReferences

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetDocumentReferences() []DocumentReferencesInner`

GetDocumentReferences returns the DocumentReferences field if non-nil, zero value otherwise.

### GetDocumentReferencesOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetDocumentReferencesOk() (*[]DocumentReferencesInner, bool)`

GetDocumentReferencesOk returns a tuple with the DocumentReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentReferences

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetDocumentReferences(v []DocumentReferencesInner)`

SetDocumentReferences sets DocumentReferences field to given value.

### HasDocumentReferences

`func (o *SearchMoveOnCommercialCycle200ResponseInner) HasDocumentReferences() bool`

HasDocumentReferences returns a boolean if a field has been set.

### GetReferences

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetReferences() []Reference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetReferencesOk() (*[]Reference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetReferences(v []Reference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *SearchMoveOnCommercialCycle200ResponseInner) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetShipmentEventTypeCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetShipmentEventTypeCode() ShipmentEventTypeCode`

GetShipmentEventTypeCode returns the ShipmentEventTypeCode field if non-nil, zero value otherwise.

### GetShipmentEventTypeCodeOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetShipmentEventTypeCodeOk() (*ShipmentEventTypeCode, bool)`

GetShipmentEventTypeCodeOk returns a tuple with the ShipmentEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentEventTypeCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetShipmentEventTypeCode(v ShipmentEventTypeCode)`

SetShipmentEventTypeCode sets ShipmentEventTypeCode field to given value.


### GetDocumentID

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetDocumentID() string`

GetDocumentID returns the DocumentID field if non-nil, zero value otherwise.

### GetDocumentIDOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetDocumentIDOk() (*string, bool)`

GetDocumentIDOk returns a tuple with the DocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentID

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetDocumentID(v string)`

SetDocumentID sets DocumentID field to given value.


### GetDocumentTypeCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetDocumentTypeCode() DocumentTypeCode`

GetDocumentTypeCode returns the DocumentTypeCode field if non-nil, zero value otherwise.

### GetDocumentTypeCodeOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetDocumentTypeCodeOk() (*DocumentTypeCode, bool)`

GetDocumentTypeCodeOk returns a tuple with the DocumentTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetDocumentTypeCode(v DocumentTypeCode)`

SetDocumentTypeCode sets DocumentTypeCode field to given value.


### GetReason

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SearchMoveOnCommercialCycle200ResponseInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetShipmentID

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetShipmentID() string`

GetShipmentID returns the ShipmentID field if non-nil, zero value otherwise.

### GetShipmentIDOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetShipmentIDOk() (*string, bool)`

GetShipmentIDOk returns a tuple with the ShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentID

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetShipmentID(v string)`

SetShipmentID sets ShipmentID field to given value.

### HasShipmentID

`func (o *SearchMoveOnCommercialCycle200ResponseInner) HasShipmentID() bool`

HasShipmentID returns a boolean if a field has been set.

### GetEquipmentEventTypeCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEquipmentEventTypeCode() EquipmentEventTypeCode`

GetEquipmentEventTypeCode returns the EquipmentEventTypeCode field if non-nil, zero value otherwise.

### GetEquipmentEventTypeCodeOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEquipmentEventTypeCodeOk() (*EquipmentEventTypeCode, bool)`

GetEquipmentEventTypeCodeOk returns a tuple with the EquipmentEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentEventTypeCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetEquipmentEventTypeCode(v EquipmentEventTypeCode)`

SetEquipmentEventTypeCode sets EquipmentEventTypeCode field to given value.


### GetEquipmentReference

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEquipmentReference() string`

GetEquipmentReference returns the EquipmentReference field if non-nil, zero value otherwise.

### GetEquipmentReferenceOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEquipmentReferenceOk() (*string, bool)`

GetEquipmentReferenceOk returns a tuple with the EquipmentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentReference

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetEquipmentReference(v string)`

SetEquipmentReference sets EquipmentReference field to given value.

### HasEquipmentReference

`func (o *SearchMoveOnCommercialCycle200ResponseInner) HasEquipmentReference() bool`

HasEquipmentReference returns a boolean if a field has been set.

### GetISOEquipmentCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetISOEquipmentCode() string`

GetISOEquipmentCode returns the ISOEquipmentCode field if non-nil, zero value otherwise.

### GetISOEquipmentCodeOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetISOEquipmentCodeOk() (*string, bool)`

GetISOEquipmentCodeOk returns a tuple with the ISOEquipmentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISOEquipmentCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetISOEquipmentCode(v string)`

SetISOEquipmentCode sets ISOEquipmentCode field to given value.

### HasISOEquipmentCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) HasISOEquipmentCode() bool`

HasISOEquipmentCode returns a boolean if a field has been set.

### GetEmptyIndicatorCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEmptyIndicatorCode() EmptyIndicatorCode`

GetEmptyIndicatorCode returns the EmptyIndicatorCode field if non-nil, zero value otherwise.

### GetEmptyIndicatorCodeOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEmptyIndicatorCodeOk() (*EmptyIndicatorCode, bool)`

GetEmptyIndicatorCodeOk returns a tuple with the EmptyIndicatorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyIndicatorCode

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetEmptyIndicatorCode(v EmptyIndicatorCode)`

SetEmptyIndicatorCode sets EmptyIndicatorCode field to given value.


### GetEventLocation

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEventLocation() Location`

GetEventLocation returns the EventLocation field if non-nil, zero value otherwise.

### GetEventLocationOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetEventLocationOk() (*Location, bool)`

GetEventLocationOk returns a tuple with the EventLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLocation

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetEventLocation(v Location)`

SetEventLocation sets EventLocation field to given value.

### HasEventLocation

`func (o *SearchMoveOnCommercialCycle200ResponseInner) HasEventLocation() bool`

HasEventLocation returns a boolean if a field has been set.

### GetSeals

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetSeals() []Seal`

GetSeals returns the Seals field if non-nil, zero value otherwise.

### GetSealsOk

`func (o *SearchMoveOnCommercialCycle200ResponseInner) GetSealsOk() (*[]Seal, bool)`

GetSealsOk returns a tuple with the Seals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeals

`func (o *SearchMoveOnCommercialCycle200ResponseInner) SetSeals(v []Seal)`

SetSeals sets Seals field to given value.

### HasSeals

`func (o *SearchMoveOnCommercialCycle200ResponseInner) HasSeals() bool`

HasSeals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


