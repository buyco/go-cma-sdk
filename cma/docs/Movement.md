# Movement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusOrder** | Pointer to **string** | movement display order | [optional] 
**Status** | Pointer to [**CodeAndName**](CodeAndName.md) |  | [optional] 
**Date** | Pointer to **time.Time** | date of the movement | [optional] 
**ReportedOn** | Pointer to **time.Time** | date of move integration | [optional] 
**PoolLocation** | Pointer to **string** | terminal of the movement | [optional] 
**PointLocation** | Pointer to [**CodeAndName**](CodeAndName.md) |  | [optional] 
**VoyageReference** | Pointer to **string** | Voyage Reference | [optional] 
**Vessel** | Pointer to [**Vessel**](Vessel.md) |  | [optional] 
**PointOfDischarge** | Pointer to **string** | point of discharge | [optional] 
**PortOfOrigin** | Pointer to **string** | point of origin | [optional] 
**PortOfLoading** | Pointer to **string** | point of loading | [optional] 
**FinalPod** | Pointer to **string** | final port of discharge | [optional] 
**FinalDest** | Pointer to **string** | final point of destination | [optional] 
**CountryCode** | Pointer to **string** | country code such as FR | [optional] 
**ShipCompCode** | Pointer to **string** | shipping companies codes are 0001, 0002... provided by CMA CGM referential | [optional] 
**VoyageShipCompCode** | Pointer to **string** | voyage number | [optional] 

## Methods

### NewMovement

`func NewMovement() *Movement`

NewMovement instantiates a new Movement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMovementWithDefaults

`func NewMovementWithDefaults() *Movement`

NewMovementWithDefaults instantiates a new Movement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusOrder

`func (o *Movement) GetStatusOrder() string`

GetStatusOrder returns the StatusOrder field if non-nil, zero value otherwise.

### GetStatusOrderOk

`func (o *Movement) GetStatusOrderOk() (*string, bool)`

GetStatusOrderOk returns a tuple with the StatusOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusOrder

`func (o *Movement) SetStatusOrder(v string)`

SetStatusOrder sets StatusOrder field to given value.

### HasStatusOrder

`func (o *Movement) HasStatusOrder() bool`

HasStatusOrder returns a boolean if a field has been set.

### GetStatus

`func (o *Movement) GetStatus() CodeAndName`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Movement) GetStatusOk() (*CodeAndName, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Movement) SetStatus(v CodeAndName)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Movement) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDate

`func (o *Movement) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Movement) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Movement) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *Movement) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetReportedOn

`func (o *Movement) GetReportedOn() time.Time`

GetReportedOn returns the ReportedOn field if non-nil, zero value otherwise.

### GetReportedOnOk

`func (o *Movement) GetReportedOnOk() (*time.Time, bool)`

GetReportedOnOk returns a tuple with the ReportedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedOn

`func (o *Movement) SetReportedOn(v time.Time)`

SetReportedOn sets ReportedOn field to given value.

### HasReportedOn

`func (o *Movement) HasReportedOn() bool`

HasReportedOn returns a boolean if a field has been set.

### GetPoolLocation

`func (o *Movement) GetPoolLocation() string`

GetPoolLocation returns the PoolLocation field if non-nil, zero value otherwise.

### GetPoolLocationOk

`func (o *Movement) GetPoolLocationOk() (*string, bool)`

GetPoolLocationOk returns a tuple with the PoolLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolLocation

`func (o *Movement) SetPoolLocation(v string)`

SetPoolLocation sets PoolLocation field to given value.

### HasPoolLocation

`func (o *Movement) HasPoolLocation() bool`

HasPoolLocation returns a boolean if a field has been set.

### GetPointLocation

`func (o *Movement) GetPointLocation() CodeAndName`

GetPointLocation returns the PointLocation field if non-nil, zero value otherwise.

### GetPointLocationOk

`func (o *Movement) GetPointLocationOk() (*CodeAndName, bool)`

GetPointLocationOk returns a tuple with the PointLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointLocation

`func (o *Movement) SetPointLocation(v CodeAndName)`

SetPointLocation sets PointLocation field to given value.

### HasPointLocation

`func (o *Movement) HasPointLocation() bool`

HasPointLocation returns a boolean if a field has been set.

### GetVoyageReference

`func (o *Movement) GetVoyageReference() string`

GetVoyageReference returns the VoyageReference field if non-nil, zero value otherwise.

### GetVoyageReferenceOk

`func (o *Movement) GetVoyageReferenceOk() (*string, bool)`

GetVoyageReferenceOk returns a tuple with the VoyageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageReference

`func (o *Movement) SetVoyageReference(v string)`

SetVoyageReference sets VoyageReference field to given value.

### HasVoyageReference

`func (o *Movement) HasVoyageReference() bool`

HasVoyageReference returns a boolean if a field has been set.

### GetVessel

`func (o *Movement) GetVessel() Vessel`

GetVessel returns the Vessel field if non-nil, zero value otherwise.

### GetVesselOk

`func (o *Movement) GetVesselOk() (*Vessel, bool)`

GetVesselOk returns a tuple with the Vessel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVessel

`func (o *Movement) SetVessel(v Vessel)`

SetVessel sets Vessel field to given value.

### HasVessel

`func (o *Movement) HasVessel() bool`

HasVessel returns a boolean if a field has been set.

### GetPointOfDischarge

`func (o *Movement) GetPointOfDischarge() string`

GetPointOfDischarge returns the PointOfDischarge field if non-nil, zero value otherwise.

### GetPointOfDischargeOk

`func (o *Movement) GetPointOfDischargeOk() (*string, bool)`

GetPointOfDischargeOk returns a tuple with the PointOfDischarge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointOfDischarge

`func (o *Movement) SetPointOfDischarge(v string)`

SetPointOfDischarge sets PointOfDischarge field to given value.

### HasPointOfDischarge

`func (o *Movement) HasPointOfDischarge() bool`

HasPointOfDischarge returns a boolean if a field has been set.

### GetPortOfOrigin

`func (o *Movement) GetPortOfOrigin() string`

GetPortOfOrigin returns the PortOfOrigin field if non-nil, zero value otherwise.

### GetPortOfOriginOk

`func (o *Movement) GetPortOfOriginOk() (*string, bool)`

GetPortOfOriginOk returns a tuple with the PortOfOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfOrigin

`func (o *Movement) SetPortOfOrigin(v string)`

SetPortOfOrigin sets PortOfOrigin field to given value.

### HasPortOfOrigin

`func (o *Movement) HasPortOfOrigin() bool`

HasPortOfOrigin returns a boolean if a field has been set.

### GetPortOfLoading

`func (o *Movement) GetPortOfLoading() string`

GetPortOfLoading returns the PortOfLoading field if non-nil, zero value otherwise.

### GetPortOfLoadingOk

`func (o *Movement) GetPortOfLoadingOk() (*string, bool)`

GetPortOfLoadingOk returns a tuple with the PortOfLoading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfLoading

`func (o *Movement) SetPortOfLoading(v string)`

SetPortOfLoading sets PortOfLoading field to given value.

### HasPortOfLoading

`func (o *Movement) HasPortOfLoading() bool`

HasPortOfLoading returns a boolean if a field has been set.

### GetFinalPod

`func (o *Movement) GetFinalPod() string`

GetFinalPod returns the FinalPod field if non-nil, zero value otherwise.

### GetFinalPodOk

`func (o *Movement) GetFinalPodOk() (*string, bool)`

GetFinalPodOk returns a tuple with the FinalPod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalPod

`func (o *Movement) SetFinalPod(v string)`

SetFinalPod sets FinalPod field to given value.

### HasFinalPod

`func (o *Movement) HasFinalPod() bool`

HasFinalPod returns a boolean if a field has been set.

### GetFinalDest

`func (o *Movement) GetFinalDest() string`

GetFinalDest returns the FinalDest field if non-nil, zero value otherwise.

### GetFinalDestOk

`func (o *Movement) GetFinalDestOk() (*string, bool)`

GetFinalDestOk returns a tuple with the FinalDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalDest

`func (o *Movement) SetFinalDest(v string)`

SetFinalDest sets FinalDest field to given value.

### HasFinalDest

`func (o *Movement) HasFinalDest() bool`

HasFinalDest returns a boolean if a field has been set.

### GetCountryCode

`func (o *Movement) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *Movement) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *Movement) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *Movement) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetShipCompCode

`func (o *Movement) GetShipCompCode() string`

GetShipCompCode returns the ShipCompCode field if non-nil, zero value otherwise.

### GetShipCompCodeOk

`func (o *Movement) GetShipCompCodeOk() (*string, bool)`

GetShipCompCodeOk returns a tuple with the ShipCompCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipCompCode

`func (o *Movement) SetShipCompCode(v string)`

SetShipCompCode sets ShipCompCode field to given value.

### HasShipCompCode

`func (o *Movement) HasShipCompCode() bool`

HasShipCompCode returns a boolean if a field has been set.

### GetVoyageShipCompCode

`func (o *Movement) GetVoyageShipCompCode() string`

GetVoyageShipCompCode returns the VoyageShipCompCode field if non-nil, zero value otherwise.

### GetVoyageShipCompCodeOk

`func (o *Movement) GetVoyageShipCompCodeOk() (*string, bool)`

GetVoyageShipCompCodeOk returns a tuple with the VoyageShipCompCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageShipCompCode

`func (o *Movement) SetVoyageShipCompCode(v string)`

SetVoyageShipCompCode sets VoyageShipCompCode field to given value.

### HasVoyageShipCompCode

`func (o *Movement) HasVoyageShipCompCode() bool`

HasVoyageShipCompCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


