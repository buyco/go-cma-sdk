# CarrierSpecificData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalEventCode** | Pointer to **string** | CMA internal event Code | [optional] 
**InternalEventLabel** | Pointer to **string** | CMA internal event Label | [optional] 
**InternalLocationCode** | Pointer to **string** | CMA internal location Code | [optional] 
**InternalFacilityCode** | Pointer to **string** | CMA internal facility Code | [optional] 
**BookingExportVoyageReference** | Pointer to **string** | Carrier export voyage reference defined on Booking | [optional] 
**TransportationPhase** | Pointer to **string** | The Transportation phase, Export, Transshipent or Import | [optional] 
**TransportCallSequenceTotal** | Pointer to **int32** | The total number of sequence provided in transportation plan | [optional] 
**NumberOfUnits** | Pointer to **int32** | The total number equipment units concerned by the message | [optional] 

## Methods

### NewCarrierSpecificData

`func NewCarrierSpecificData() *CarrierSpecificData`

NewCarrierSpecificData instantiates a new CarrierSpecificData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCarrierSpecificDataWithDefaults

`func NewCarrierSpecificDataWithDefaults() *CarrierSpecificData`

NewCarrierSpecificDataWithDefaults instantiates a new CarrierSpecificData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInternalEventCode

`func (o *CarrierSpecificData) GetInternalEventCode() string`

GetInternalEventCode returns the InternalEventCode field if non-nil, zero value otherwise.

### GetInternalEventCodeOk

`func (o *CarrierSpecificData) GetInternalEventCodeOk() (*string, bool)`

GetInternalEventCodeOk returns a tuple with the InternalEventCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalEventCode

`func (o *CarrierSpecificData) SetInternalEventCode(v string)`

SetInternalEventCode sets InternalEventCode field to given value.

### HasInternalEventCode

`func (o *CarrierSpecificData) HasInternalEventCode() bool`

HasInternalEventCode returns a boolean if a field has been set.

### GetInternalEventLabel

`func (o *CarrierSpecificData) GetInternalEventLabel() string`

GetInternalEventLabel returns the InternalEventLabel field if non-nil, zero value otherwise.

### GetInternalEventLabelOk

`func (o *CarrierSpecificData) GetInternalEventLabelOk() (*string, bool)`

GetInternalEventLabelOk returns a tuple with the InternalEventLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalEventLabel

`func (o *CarrierSpecificData) SetInternalEventLabel(v string)`

SetInternalEventLabel sets InternalEventLabel field to given value.

### HasInternalEventLabel

`func (o *CarrierSpecificData) HasInternalEventLabel() bool`

HasInternalEventLabel returns a boolean if a field has been set.

### GetInternalLocationCode

`func (o *CarrierSpecificData) GetInternalLocationCode() string`

GetInternalLocationCode returns the InternalLocationCode field if non-nil, zero value otherwise.

### GetInternalLocationCodeOk

`func (o *CarrierSpecificData) GetInternalLocationCodeOk() (*string, bool)`

GetInternalLocationCodeOk returns a tuple with the InternalLocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalLocationCode

`func (o *CarrierSpecificData) SetInternalLocationCode(v string)`

SetInternalLocationCode sets InternalLocationCode field to given value.

### HasInternalLocationCode

`func (o *CarrierSpecificData) HasInternalLocationCode() bool`

HasInternalLocationCode returns a boolean if a field has been set.

### GetInternalFacilityCode

`func (o *CarrierSpecificData) GetInternalFacilityCode() string`

GetInternalFacilityCode returns the InternalFacilityCode field if non-nil, zero value otherwise.

### GetInternalFacilityCodeOk

`func (o *CarrierSpecificData) GetInternalFacilityCodeOk() (*string, bool)`

GetInternalFacilityCodeOk returns a tuple with the InternalFacilityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalFacilityCode

`func (o *CarrierSpecificData) SetInternalFacilityCode(v string)`

SetInternalFacilityCode sets InternalFacilityCode field to given value.

### HasInternalFacilityCode

`func (o *CarrierSpecificData) HasInternalFacilityCode() bool`

HasInternalFacilityCode returns a boolean if a field has been set.

### GetBookingExportVoyageReference

`func (o *CarrierSpecificData) GetBookingExportVoyageReference() string`

GetBookingExportVoyageReference returns the BookingExportVoyageReference field if non-nil, zero value otherwise.

### GetBookingExportVoyageReferenceOk

`func (o *CarrierSpecificData) GetBookingExportVoyageReferenceOk() (*string, bool)`

GetBookingExportVoyageReferenceOk returns a tuple with the BookingExportVoyageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingExportVoyageReference

`func (o *CarrierSpecificData) SetBookingExportVoyageReference(v string)`

SetBookingExportVoyageReference sets BookingExportVoyageReference field to given value.

### HasBookingExportVoyageReference

`func (o *CarrierSpecificData) HasBookingExportVoyageReference() bool`

HasBookingExportVoyageReference returns a boolean if a field has been set.

### GetTransportationPhase

`func (o *CarrierSpecificData) GetTransportationPhase() string`

GetTransportationPhase returns the TransportationPhase field if non-nil, zero value otherwise.

### GetTransportationPhaseOk

`func (o *CarrierSpecificData) GetTransportationPhaseOk() (*string, bool)`

GetTransportationPhaseOk returns a tuple with the TransportationPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportationPhase

`func (o *CarrierSpecificData) SetTransportationPhase(v string)`

SetTransportationPhase sets TransportationPhase field to given value.

### HasTransportationPhase

`func (o *CarrierSpecificData) HasTransportationPhase() bool`

HasTransportationPhase returns a boolean if a field has been set.

### GetTransportCallSequenceTotal

`func (o *CarrierSpecificData) GetTransportCallSequenceTotal() int32`

GetTransportCallSequenceTotal returns the TransportCallSequenceTotal field if non-nil, zero value otherwise.

### GetTransportCallSequenceTotalOk

`func (o *CarrierSpecificData) GetTransportCallSequenceTotalOk() (*int32, bool)`

GetTransportCallSequenceTotalOk returns a tuple with the TransportCallSequenceTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCallSequenceTotal

`func (o *CarrierSpecificData) SetTransportCallSequenceTotal(v int32)`

SetTransportCallSequenceTotal sets TransportCallSequenceTotal field to given value.

### HasTransportCallSequenceTotal

`func (o *CarrierSpecificData) HasTransportCallSequenceTotal() bool`

HasTransportCallSequenceTotal returns a boolean if a field has been set.

### GetNumberOfUnits

`func (o *CarrierSpecificData) GetNumberOfUnits() int32`

GetNumberOfUnits returns the NumberOfUnits field if non-nil, zero value otherwise.

### GetNumberOfUnitsOk

`func (o *CarrierSpecificData) GetNumberOfUnitsOk() (*int32, bool)`

GetNumberOfUnitsOk returns a tuple with the NumberOfUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfUnits

`func (o *CarrierSpecificData) SetNumberOfUnits(v int32)`

SetNumberOfUnits sets NumberOfUnits field to given value.

### HasNumberOfUnits

`func (o *CarrierSpecificData) HasNumberOfUnits() bool`

HasNumberOfUnits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


