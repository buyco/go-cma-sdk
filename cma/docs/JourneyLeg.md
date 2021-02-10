# JourneyLeg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SequenceNumber** | Pointer to **int32** | sequence number of the journey leg | [optional] 
**PointFrom** | Pointer to [**CodeAndName**](CodeAndName.md) |  | [optional] 
**VesselFrom** | Pointer to [**Vessel**](Vessel.md) |  | [optional] 
**VesselTo** | Pointer to [**Vessel**](Vessel.md) |  | [optional] 
**PointTo** | Pointer to [**CodeAndName**](CodeAndName.md) |  | [optional] 
**PoolLocationFromCode** | Pointer to **string** | pool code of journey leg departure | [optional] 
**PoolLocationToCode** | Pointer to **string** | pool code of journey leg arrival | [optional] 
**CollectionDate** | Pointer to **time.Time** | collection date | [optional] 
**VoyageReference** | Pointer to **string** | voyage reference at the journey leg departure | [optional] 
**DischargeVoyageReference** | Pointer to **string** | voyage reference at the journey leg arrival | [optional] 
**DeliveryDate** | Pointer to **time.Time** | delivery date | [optional] 
**ShipCompCode** | Pointer to **string** | shipping companies codes are 0001, 0002... provided by CMA CGM | [optional] 

## Methods

### NewJourneyLeg

`func NewJourneyLeg() *JourneyLeg`

NewJourneyLeg instantiates a new JourneyLeg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJourneyLegWithDefaults

`func NewJourneyLegWithDefaults() *JourneyLeg`

NewJourneyLegWithDefaults instantiates a new JourneyLeg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSequenceNumber

`func (o *JourneyLeg) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *JourneyLeg) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *JourneyLeg) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *JourneyLeg) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetPointFrom

`func (o *JourneyLeg) GetPointFrom() CodeAndName`

GetPointFrom returns the PointFrom field if non-nil, zero value otherwise.

### GetPointFromOk

`func (o *JourneyLeg) GetPointFromOk() (*CodeAndName, bool)`

GetPointFromOk returns a tuple with the PointFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointFrom

`func (o *JourneyLeg) SetPointFrom(v CodeAndName)`

SetPointFrom sets PointFrom field to given value.

### HasPointFrom

`func (o *JourneyLeg) HasPointFrom() bool`

HasPointFrom returns a boolean if a field has been set.

### GetVesselFrom

`func (o *JourneyLeg) GetVesselFrom() Vessel`

GetVesselFrom returns the VesselFrom field if non-nil, zero value otherwise.

### GetVesselFromOk

`func (o *JourneyLeg) GetVesselFromOk() (*Vessel, bool)`

GetVesselFromOk returns a tuple with the VesselFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselFrom

`func (o *JourneyLeg) SetVesselFrom(v Vessel)`

SetVesselFrom sets VesselFrom field to given value.

### HasVesselFrom

`func (o *JourneyLeg) HasVesselFrom() bool`

HasVesselFrom returns a boolean if a field has been set.

### GetVesselTo

`func (o *JourneyLeg) GetVesselTo() Vessel`

GetVesselTo returns the VesselTo field if non-nil, zero value otherwise.

### GetVesselToOk

`func (o *JourneyLeg) GetVesselToOk() (*Vessel, bool)`

GetVesselToOk returns a tuple with the VesselTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVesselTo

`func (o *JourneyLeg) SetVesselTo(v Vessel)`

SetVesselTo sets VesselTo field to given value.

### HasVesselTo

`func (o *JourneyLeg) HasVesselTo() bool`

HasVesselTo returns a boolean if a field has been set.

### GetPointTo

`func (o *JourneyLeg) GetPointTo() CodeAndName`

GetPointTo returns the PointTo field if non-nil, zero value otherwise.

### GetPointToOk

`func (o *JourneyLeg) GetPointToOk() (*CodeAndName, bool)`

GetPointToOk returns a tuple with the PointTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointTo

`func (o *JourneyLeg) SetPointTo(v CodeAndName)`

SetPointTo sets PointTo field to given value.

### HasPointTo

`func (o *JourneyLeg) HasPointTo() bool`

HasPointTo returns a boolean if a field has been set.

### GetPoolLocationFromCode

`func (o *JourneyLeg) GetPoolLocationFromCode() string`

GetPoolLocationFromCode returns the PoolLocationFromCode field if non-nil, zero value otherwise.

### GetPoolLocationFromCodeOk

`func (o *JourneyLeg) GetPoolLocationFromCodeOk() (*string, bool)`

GetPoolLocationFromCodeOk returns a tuple with the PoolLocationFromCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolLocationFromCode

`func (o *JourneyLeg) SetPoolLocationFromCode(v string)`

SetPoolLocationFromCode sets PoolLocationFromCode field to given value.

### HasPoolLocationFromCode

`func (o *JourneyLeg) HasPoolLocationFromCode() bool`

HasPoolLocationFromCode returns a boolean if a field has been set.

### GetPoolLocationToCode

`func (o *JourneyLeg) GetPoolLocationToCode() string`

GetPoolLocationToCode returns the PoolLocationToCode field if non-nil, zero value otherwise.

### GetPoolLocationToCodeOk

`func (o *JourneyLeg) GetPoolLocationToCodeOk() (*string, bool)`

GetPoolLocationToCodeOk returns a tuple with the PoolLocationToCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolLocationToCode

`func (o *JourneyLeg) SetPoolLocationToCode(v string)`

SetPoolLocationToCode sets PoolLocationToCode field to given value.

### HasPoolLocationToCode

`func (o *JourneyLeg) HasPoolLocationToCode() bool`

HasPoolLocationToCode returns a boolean if a field has been set.

### GetCollectionDate

`func (o *JourneyLeg) GetCollectionDate() time.Time`

GetCollectionDate returns the CollectionDate field if non-nil, zero value otherwise.

### GetCollectionDateOk

`func (o *JourneyLeg) GetCollectionDateOk() (*time.Time, bool)`

GetCollectionDateOk returns a tuple with the CollectionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionDate

`func (o *JourneyLeg) SetCollectionDate(v time.Time)`

SetCollectionDate sets CollectionDate field to given value.

### HasCollectionDate

`func (o *JourneyLeg) HasCollectionDate() bool`

HasCollectionDate returns a boolean if a field has been set.

### GetVoyageReference

`func (o *JourneyLeg) GetVoyageReference() string`

GetVoyageReference returns the VoyageReference field if non-nil, zero value otherwise.

### GetVoyageReferenceOk

`func (o *JourneyLeg) GetVoyageReferenceOk() (*string, bool)`

GetVoyageReferenceOk returns a tuple with the VoyageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageReference

`func (o *JourneyLeg) SetVoyageReference(v string)`

SetVoyageReference sets VoyageReference field to given value.

### HasVoyageReference

`func (o *JourneyLeg) HasVoyageReference() bool`

HasVoyageReference returns a boolean if a field has been set.

### GetDischargeVoyageReference

`func (o *JourneyLeg) GetDischargeVoyageReference() string`

GetDischargeVoyageReference returns the DischargeVoyageReference field if non-nil, zero value otherwise.

### GetDischargeVoyageReferenceOk

`func (o *JourneyLeg) GetDischargeVoyageReferenceOk() (*string, bool)`

GetDischargeVoyageReferenceOk returns a tuple with the DischargeVoyageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDischargeVoyageReference

`func (o *JourneyLeg) SetDischargeVoyageReference(v string)`

SetDischargeVoyageReference sets DischargeVoyageReference field to given value.

### HasDischargeVoyageReference

`func (o *JourneyLeg) HasDischargeVoyageReference() bool`

HasDischargeVoyageReference returns a boolean if a field has been set.

### GetDeliveryDate

`func (o *JourneyLeg) GetDeliveryDate() time.Time`

GetDeliveryDate returns the DeliveryDate field if non-nil, zero value otherwise.

### GetDeliveryDateOk

`func (o *JourneyLeg) GetDeliveryDateOk() (*time.Time, bool)`

GetDeliveryDateOk returns a tuple with the DeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryDate

`func (o *JourneyLeg) SetDeliveryDate(v time.Time)`

SetDeliveryDate sets DeliveryDate field to given value.

### HasDeliveryDate

`func (o *JourneyLeg) HasDeliveryDate() bool`

HasDeliveryDate returns a boolean if a field has been set.

### GetShipCompCode

`func (o *JourneyLeg) GetShipCompCode() string`

GetShipCompCode returns the ShipCompCode field if non-nil, zero value otherwise.

### GetShipCompCodeOk

`func (o *JourneyLeg) GetShipCompCodeOk() (*string, bool)`

GetShipCompCodeOk returns a tuple with the ShipCompCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipCompCode

`func (o *JourneyLeg) SetShipCompCode(v string)`

SetShipCompCode sets ShipCompCode field to given value.

### HasShipCompCode

`func (o *JourneyLeg) HasShipCompCode() bool`

HasShipCompCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


