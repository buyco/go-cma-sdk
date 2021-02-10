# Container

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Container ID | [optional] 
**Size** | Pointer to **int32** | Size in TEU | [optional] 
**EmptyReturnDepot** | Pointer to **string** | Empty return depot | [optional] 
**Type** | Pointer to **string** | type of the container such as ST, .. | [optional] 
**Movements** | Pointer to [**[]Movement**](Movement.md) |  | [optional] 

## Methods

### NewContainer

`func NewContainer() *Container`

NewContainer instantiates a new Container object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerWithDefaults

`func NewContainerWithDefaults() *Container`

NewContainerWithDefaults instantiates a new Container object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Container) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Container) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Container) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Container) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSize

`func (o *Container) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Container) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Container) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Container) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetEmptyReturnDepot

`func (o *Container) GetEmptyReturnDepot() string`

GetEmptyReturnDepot returns the EmptyReturnDepot field if non-nil, zero value otherwise.

### GetEmptyReturnDepotOk

`func (o *Container) GetEmptyReturnDepotOk() (*string, bool)`

GetEmptyReturnDepotOk returns a tuple with the EmptyReturnDepot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyReturnDepot

`func (o *Container) SetEmptyReturnDepot(v string)`

SetEmptyReturnDepot sets EmptyReturnDepot field to given value.

### HasEmptyReturnDepot

`func (o *Container) HasEmptyReturnDepot() bool`

HasEmptyReturnDepot returns a boolean if a field has been set.

### GetType

`func (o *Container) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Container) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Container) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Container) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMovements

`func (o *Container) GetMovements() []Movement`

GetMovements returns the Movements field if non-nil, zero value otherwise.

### GetMovementsOk

`func (o *Container) GetMovementsOk() (*[]Movement, bool)`

GetMovementsOk returns a tuple with the Movements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovements

`func (o *Container) SetMovements(v []Movement)`

SetMovements sets Movements field to given value.

### HasMovements

`func (o *Container) HasMovements() bool`

HasMovements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


