# CustomsReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquipmentReference** | Pointer to **string** | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation | [optional] 
**CustomsReference** | Pointer to **string** | the customs reference for the container | [optional] 
**ReferenceDate** | Pointer to **string** | the customs reference date | [optional] 

## Methods

### NewCustomsReferences

`func NewCustomsReferences() *CustomsReferences`

NewCustomsReferences instantiates a new CustomsReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomsReferencesWithDefaults

`func NewCustomsReferencesWithDefaults() *CustomsReferences`

NewCustomsReferencesWithDefaults instantiates a new CustomsReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquipmentReference

`func (o *CustomsReferences) GetEquipmentReference() string`

GetEquipmentReference returns the EquipmentReference field if non-nil, zero value otherwise.

### GetEquipmentReferenceOk

`func (o *CustomsReferences) GetEquipmentReferenceOk() (*string, bool)`

GetEquipmentReferenceOk returns a tuple with the EquipmentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentReference

`func (o *CustomsReferences) SetEquipmentReference(v string)`

SetEquipmentReference sets EquipmentReference field to given value.

### HasEquipmentReference

`func (o *CustomsReferences) HasEquipmentReference() bool`

HasEquipmentReference returns a boolean if a field has been set.

### GetCustomsReference

`func (o *CustomsReferences) GetCustomsReference() string`

GetCustomsReference returns the CustomsReference field if non-nil, zero value otherwise.

### GetCustomsReferenceOk

`func (o *CustomsReferences) GetCustomsReferenceOk() (*string, bool)`

GetCustomsReferenceOk returns a tuple with the CustomsReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsReference

`func (o *CustomsReferences) SetCustomsReference(v string)`

SetCustomsReference sets CustomsReference field to given value.

### HasCustomsReference

`func (o *CustomsReferences) HasCustomsReference() bool`

HasCustomsReference returns a boolean if a field has been set.

### GetReferenceDate

`func (o *CustomsReferences) GetReferenceDate() string`

GetReferenceDate returns the ReferenceDate field if non-nil, zero value otherwise.

### GetReferenceDateOk

`func (o *CustomsReferences) GetReferenceDateOk() (*string, bool)`

GetReferenceDateOk returns a tuple with the ReferenceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceDate

`func (o *CustomsReferences) SetReferenceDate(v string)`

SetReferenceDate sets ReferenceDate field to given value.

### HasReferenceDate

`func (o *CustomsReferences) HasReferenceDate() bool`

HasReferenceDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


