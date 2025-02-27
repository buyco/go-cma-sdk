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
**ShipmentLocationType** | Pointer to **string** | The location type for the shipment | [optional] 
**TransportCallSequenceTotal** | Pointer to **int32** | The total number of sequence provided in transportation plan | [optional] 
**NumberOfUnits** | Pointer to **int32** | The total number equipment units concerned by the message | [optional] 
**CustomsReferences** | Pointer to [**[]CustomsReferences**](CustomsReferences.md) |  | [optional] 
**OriginBookingReference** | Pointer to **string** | The identifier for a shipment | [optional] 
**SplitToBookingReferences** | Pointer to **[]string** |  | [optional] 
**MergeToBookingReference** | Pointer to **string** | The identifier for a shipment | [optional] 
**MasterPartLoadBookingReference** | Pointer to **string** | The identifier for a shipment | [optional] 
**PartLoadGroupBookingReferences** | Pointer to **[]string** |  | [optional] 

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

### GetShipmentLocationType

`func (o *CarrierSpecificData) GetShipmentLocationType() string`

GetShipmentLocationType returns the ShipmentLocationType field if non-nil, zero value otherwise.

### GetShipmentLocationTypeOk

`func (o *CarrierSpecificData) GetShipmentLocationTypeOk() (*string, bool)`

GetShipmentLocationTypeOk returns a tuple with the ShipmentLocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentLocationType

`func (o *CarrierSpecificData) SetShipmentLocationType(v string)`

SetShipmentLocationType sets ShipmentLocationType field to given value.

### HasShipmentLocationType

`func (o *CarrierSpecificData) HasShipmentLocationType() bool`

HasShipmentLocationType returns a boolean if a field has been set.

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

### GetCustomsReferences

`func (o *CarrierSpecificData) GetCustomsReferences() []CustomsReferences`

GetCustomsReferences returns the CustomsReferences field if non-nil, zero value otherwise.

### GetCustomsReferencesOk

`func (o *CarrierSpecificData) GetCustomsReferencesOk() (*[]CustomsReferences, bool)`

GetCustomsReferencesOk returns a tuple with the CustomsReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsReferences

`func (o *CarrierSpecificData) SetCustomsReferences(v []CustomsReferences)`

SetCustomsReferences sets CustomsReferences field to given value.

### HasCustomsReferences

`func (o *CarrierSpecificData) HasCustomsReferences() bool`

HasCustomsReferences returns a boolean if a field has been set.

### GetOriginBookingReference

`func (o *CarrierSpecificData) GetOriginBookingReference() string`

GetOriginBookingReference returns the OriginBookingReference field if non-nil, zero value otherwise.

### GetOriginBookingReferenceOk

`func (o *CarrierSpecificData) GetOriginBookingReferenceOk() (*string, bool)`

GetOriginBookingReferenceOk returns a tuple with the OriginBookingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginBookingReference

`func (o *CarrierSpecificData) SetOriginBookingReference(v string)`

SetOriginBookingReference sets OriginBookingReference field to given value.

### HasOriginBookingReference

`func (o *CarrierSpecificData) HasOriginBookingReference() bool`

HasOriginBookingReference returns a boolean if a field has been set.

### GetSplitToBookingReferences

`func (o *CarrierSpecificData) GetSplitToBookingReferences() []string`

GetSplitToBookingReferences returns the SplitToBookingReferences field if non-nil, zero value otherwise.

### GetSplitToBookingReferencesOk

`func (o *CarrierSpecificData) GetSplitToBookingReferencesOk() (*[]string, bool)`

GetSplitToBookingReferencesOk returns a tuple with the SplitToBookingReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitToBookingReferences

`func (o *CarrierSpecificData) SetSplitToBookingReferences(v []string)`

SetSplitToBookingReferences sets SplitToBookingReferences field to given value.

### HasSplitToBookingReferences

`func (o *CarrierSpecificData) HasSplitToBookingReferences() bool`

HasSplitToBookingReferences returns a boolean if a field has been set.

### GetMergeToBookingReference

`func (o *CarrierSpecificData) GetMergeToBookingReference() string`

GetMergeToBookingReference returns the MergeToBookingReference field if non-nil, zero value otherwise.

### GetMergeToBookingReferenceOk

`func (o *CarrierSpecificData) GetMergeToBookingReferenceOk() (*string, bool)`

GetMergeToBookingReferenceOk returns a tuple with the MergeToBookingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeToBookingReference

`func (o *CarrierSpecificData) SetMergeToBookingReference(v string)`

SetMergeToBookingReference sets MergeToBookingReference field to given value.

### HasMergeToBookingReference

`func (o *CarrierSpecificData) HasMergeToBookingReference() bool`

HasMergeToBookingReference returns a boolean if a field has been set.

### GetMasterPartLoadBookingReference

`func (o *CarrierSpecificData) GetMasterPartLoadBookingReference() string`

GetMasterPartLoadBookingReference returns the MasterPartLoadBookingReference field if non-nil, zero value otherwise.

### GetMasterPartLoadBookingReferenceOk

`func (o *CarrierSpecificData) GetMasterPartLoadBookingReferenceOk() (*string, bool)`

GetMasterPartLoadBookingReferenceOk returns a tuple with the MasterPartLoadBookingReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterPartLoadBookingReference

`func (o *CarrierSpecificData) SetMasterPartLoadBookingReference(v string)`

SetMasterPartLoadBookingReference sets MasterPartLoadBookingReference field to given value.

### HasMasterPartLoadBookingReference

`func (o *CarrierSpecificData) HasMasterPartLoadBookingReference() bool`

HasMasterPartLoadBookingReference returns a boolean if a field has been set.

### GetPartLoadGroupBookingReferences

`func (o *CarrierSpecificData) GetPartLoadGroupBookingReferences() []string`

GetPartLoadGroupBookingReferences returns the PartLoadGroupBookingReferences field if non-nil, zero value otherwise.

### GetPartLoadGroupBookingReferencesOk

`func (o *CarrierSpecificData) GetPartLoadGroupBookingReferencesOk() (*[]string, bool)`

GetPartLoadGroupBookingReferencesOk returns a tuple with the PartLoadGroupBookingReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartLoadGroupBookingReferences

`func (o *CarrierSpecificData) SetPartLoadGroupBookingReferences(v []string)`

SetPartLoadGroupBookingReferences sets PartLoadGroupBookingReferences field to given value.

### HasPartLoadGroupBookingReferences

`func (o *CarrierSpecificData) HasPartLoadGroupBookingReferences() bool`

HasPartLoadGroupBookingReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


