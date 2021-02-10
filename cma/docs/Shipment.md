# Shipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortOfLoading** | Pointer to [**CodeAndName**](CodeAndName.md) |  | [optional] 
**PortOfLoadingCountryCode** | Pointer to **string** | country code of the port of loading such as FR | [optional] 
**PortOfDischarge** | Pointer to [**CodeAndName**](CodeAndName.md) |  | [optional] 
**PortOfDischargeCountryCode** | Pointer to **string** | country code of the port of discharge such as FR | [optional] 
**VoyageReference** | Pointer to **string** | Voyage Reference | [optional] 
**NbUnits** | Pointer to **int32** | Number of units | [optional] 
**Routes** | Pointer to [**[]Route**](Route.md) |  | [optional] 

## Methods

### NewShipment

`func NewShipment() *Shipment`

NewShipment instantiates a new Shipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentWithDefaults

`func NewShipmentWithDefaults() *Shipment`

NewShipmentWithDefaults instantiates a new Shipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortOfLoading

`func (o *Shipment) GetPortOfLoading() CodeAndName`

GetPortOfLoading returns the PortOfLoading field if non-nil, zero value otherwise.

### GetPortOfLoadingOk

`func (o *Shipment) GetPortOfLoadingOk() (*CodeAndName, bool)`

GetPortOfLoadingOk returns a tuple with the PortOfLoading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfLoading

`func (o *Shipment) SetPortOfLoading(v CodeAndName)`

SetPortOfLoading sets PortOfLoading field to given value.

### HasPortOfLoading

`func (o *Shipment) HasPortOfLoading() bool`

HasPortOfLoading returns a boolean if a field has been set.

### GetPortOfLoadingCountryCode

`func (o *Shipment) GetPortOfLoadingCountryCode() string`

GetPortOfLoadingCountryCode returns the PortOfLoadingCountryCode field if non-nil, zero value otherwise.

### GetPortOfLoadingCountryCodeOk

`func (o *Shipment) GetPortOfLoadingCountryCodeOk() (*string, bool)`

GetPortOfLoadingCountryCodeOk returns a tuple with the PortOfLoadingCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfLoadingCountryCode

`func (o *Shipment) SetPortOfLoadingCountryCode(v string)`

SetPortOfLoadingCountryCode sets PortOfLoadingCountryCode field to given value.

### HasPortOfLoadingCountryCode

`func (o *Shipment) HasPortOfLoadingCountryCode() bool`

HasPortOfLoadingCountryCode returns a boolean if a field has been set.

### GetPortOfDischarge

`func (o *Shipment) GetPortOfDischarge() CodeAndName`

GetPortOfDischarge returns the PortOfDischarge field if non-nil, zero value otherwise.

### GetPortOfDischargeOk

`func (o *Shipment) GetPortOfDischargeOk() (*CodeAndName, bool)`

GetPortOfDischargeOk returns a tuple with the PortOfDischarge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfDischarge

`func (o *Shipment) SetPortOfDischarge(v CodeAndName)`

SetPortOfDischarge sets PortOfDischarge field to given value.

### HasPortOfDischarge

`func (o *Shipment) HasPortOfDischarge() bool`

HasPortOfDischarge returns a boolean if a field has been set.

### GetPortOfDischargeCountryCode

`func (o *Shipment) GetPortOfDischargeCountryCode() string`

GetPortOfDischargeCountryCode returns the PortOfDischargeCountryCode field if non-nil, zero value otherwise.

### GetPortOfDischargeCountryCodeOk

`func (o *Shipment) GetPortOfDischargeCountryCodeOk() (*string, bool)`

GetPortOfDischargeCountryCodeOk returns a tuple with the PortOfDischargeCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortOfDischargeCountryCode

`func (o *Shipment) SetPortOfDischargeCountryCode(v string)`

SetPortOfDischargeCountryCode sets PortOfDischargeCountryCode field to given value.

### HasPortOfDischargeCountryCode

`func (o *Shipment) HasPortOfDischargeCountryCode() bool`

HasPortOfDischargeCountryCode returns a boolean if a field has been set.

### GetVoyageReference

`func (o *Shipment) GetVoyageReference() string`

GetVoyageReference returns the VoyageReference field if non-nil, zero value otherwise.

### GetVoyageReferenceOk

`func (o *Shipment) GetVoyageReferenceOk() (*string, bool)`

GetVoyageReferenceOk returns a tuple with the VoyageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoyageReference

`func (o *Shipment) SetVoyageReference(v string)`

SetVoyageReference sets VoyageReference field to given value.

### HasVoyageReference

`func (o *Shipment) HasVoyageReference() bool`

HasVoyageReference returns a boolean if a field has been set.

### GetNbUnits

`func (o *Shipment) GetNbUnits() int32`

GetNbUnits returns the NbUnits field if non-nil, zero value otherwise.

### GetNbUnitsOk

`func (o *Shipment) GetNbUnitsOk() (*int32, bool)`

GetNbUnitsOk returns a tuple with the NbUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbUnits

`func (o *Shipment) SetNbUnits(v int32)`

SetNbUnits sets NbUnits field to given value.

### HasNbUnits

`func (o *Shipment) HasNbUnits() bool`

HasNbUnits returns a boolean if a field has been set.

### GetRoutes

`func (o *Shipment) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *Shipment) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *Shipment) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *Shipment) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


