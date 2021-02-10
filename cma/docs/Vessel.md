# Vessel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Imo** | Pointer to **string** |  | [optional] 

## Methods

### NewVessel

`func NewVessel() *Vessel`

NewVessel instantiates a new Vessel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVesselWithDefaults

`func NewVesselWithDefaults() *Vessel`

NewVesselWithDefaults instantiates a new Vessel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Vessel) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Vessel) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Vessel) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Vessel) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *Vessel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vessel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vessel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Vessel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImo

`func (o *Vessel) GetImo() string`

GetImo returns the Imo field if non-nil, zero value otherwise.

### GetImoOk

`func (o *Vessel) GetImoOk() (*string, bool)`

GetImoOk returns a tuple with the Imo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImo

`func (o *Vessel) SetImo(v string)`

SetImo sets Imo field to given value.

### HasImo

`func (o *Vessel) HasImo() bool`

HasImo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


