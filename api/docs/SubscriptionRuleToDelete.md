# SubscriptionRuleToDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CvsServicesToDelete** | Pointer to [**[]CvsServiceToDelete**](CvsServiceToDelete.md) | &lt;br&gt; List of (serviceCode only OR serviceCode + shippingCompany) OR carrierServiceCode to delete.&lt;/br&gt; Either serviceCode OR carrierServiceCode should be set per cvs service rule.&lt;/br&gt; If serviceCode only is provided &#x3D;&gt; the unsubscription will be applied to the service code only. If serviceCode + shippingCompany is provided &#x3D;&gt; the unsubscription will be applied to the carrier service code of the specified shippingCompany. If carrierServiceCode is provided &#x3D;&gt; the unsubscription will be applied to the carrier service code.  | [optional] 
**EquipmentReferences** | Pointer to **[]string** | List of equipment references | [optional] 
**ShipmentReferences** | Pointer to **[]string** | List of shipment references | [optional] 
**ContractReferences** | Pointer to **[]string** | List of contract references | [optional] 
**ElectronicChannels** | Pointer to **[]string** | List of Dedicated channels | [optional] 

## Methods

### NewSubscriptionRuleToDelete

`func NewSubscriptionRuleToDelete() *SubscriptionRuleToDelete`

NewSubscriptionRuleToDelete instantiates a new SubscriptionRuleToDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionRuleToDeleteWithDefaults

`func NewSubscriptionRuleToDeleteWithDefaults() *SubscriptionRuleToDelete`

NewSubscriptionRuleToDeleteWithDefaults instantiates a new SubscriptionRuleToDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvsServicesToDelete

`func (o *SubscriptionRuleToDelete) GetCvsServicesToDelete() []CvsServiceToDelete`

GetCvsServicesToDelete returns the CvsServicesToDelete field if non-nil, zero value otherwise.

### GetCvsServicesToDeleteOk

`func (o *SubscriptionRuleToDelete) GetCvsServicesToDeleteOk() (*[]CvsServiceToDelete, bool)`

GetCvsServicesToDeleteOk returns a tuple with the CvsServicesToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvsServicesToDelete

`func (o *SubscriptionRuleToDelete) SetCvsServicesToDelete(v []CvsServiceToDelete)`

SetCvsServicesToDelete sets CvsServicesToDelete field to given value.

### HasCvsServicesToDelete

`func (o *SubscriptionRuleToDelete) HasCvsServicesToDelete() bool`

HasCvsServicesToDelete returns a boolean if a field has been set.

### GetEquipmentReferences

`func (o *SubscriptionRuleToDelete) GetEquipmentReferences() []string`

GetEquipmentReferences returns the EquipmentReferences field if non-nil, zero value otherwise.

### GetEquipmentReferencesOk

`func (o *SubscriptionRuleToDelete) GetEquipmentReferencesOk() (*[]string, bool)`

GetEquipmentReferencesOk returns a tuple with the EquipmentReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentReferences

`func (o *SubscriptionRuleToDelete) SetEquipmentReferences(v []string)`

SetEquipmentReferences sets EquipmentReferences field to given value.

### HasEquipmentReferences

`func (o *SubscriptionRuleToDelete) HasEquipmentReferences() bool`

HasEquipmentReferences returns a boolean if a field has been set.

### GetShipmentReferences

`func (o *SubscriptionRuleToDelete) GetShipmentReferences() []string`

GetShipmentReferences returns the ShipmentReferences field if non-nil, zero value otherwise.

### GetShipmentReferencesOk

`func (o *SubscriptionRuleToDelete) GetShipmentReferencesOk() (*[]string, bool)`

GetShipmentReferencesOk returns a tuple with the ShipmentReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentReferences

`func (o *SubscriptionRuleToDelete) SetShipmentReferences(v []string)`

SetShipmentReferences sets ShipmentReferences field to given value.

### HasShipmentReferences

`func (o *SubscriptionRuleToDelete) HasShipmentReferences() bool`

HasShipmentReferences returns a boolean if a field has been set.

### GetContractReferences

`func (o *SubscriptionRuleToDelete) GetContractReferences() []string`

GetContractReferences returns the ContractReferences field if non-nil, zero value otherwise.

### GetContractReferencesOk

`func (o *SubscriptionRuleToDelete) GetContractReferencesOk() (*[]string, bool)`

GetContractReferencesOk returns a tuple with the ContractReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractReferences

`func (o *SubscriptionRuleToDelete) SetContractReferences(v []string)`

SetContractReferences sets ContractReferences field to given value.

### HasContractReferences

`func (o *SubscriptionRuleToDelete) HasContractReferences() bool`

HasContractReferences returns a boolean if a field has been set.

### GetElectronicChannels

`func (o *SubscriptionRuleToDelete) GetElectronicChannels() []string`

GetElectronicChannels returns the ElectronicChannels field if non-nil, zero value otherwise.

### GetElectronicChannelsOk

`func (o *SubscriptionRuleToDelete) GetElectronicChannelsOk() (*[]string, bool)`

GetElectronicChannelsOk returns a tuple with the ElectronicChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectronicChannels

`func (o *SubscriptionRuleToDelete) SetElectronicChannels(v []string)`

SetElectronicChannels sets ElectronicChannels field to given value.

### HasElectronicChannels

`func (o *SubscriptionRuleToDelete) HasElectronicChannels() bool`

HasElectronicChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


