# SubscriptionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EquipmentReferences** | Pointer to **[]string** | List of equipment references | [optional] 
**Vas** | Pointer to **[]string** | List of vas | [optional] 
**ShipmentReferences** | Pointer to **[]string** | List of shipment references | [optional] 
**ContractReferences** | Pointer to **[]string** | List of contract references | [optional] 
**CvsServices** | Pointer to [**[]CvsService**](CvsService.md) | &lt;br&gt; List of (serviceCode only OR serviceCode + shippingCompany) OR carrierServiceCode to create.&lt;/br&gt; Either serviceCode OR carrierServiceCode should be set per cvs service rule.&lt;/br&gt; If serviceCode only is provided &#x3D;&gt; the subscription will be applied to the service code only. If serviceCode + shippingCompany is provided &#x3D;&gt; the subscription will be applied to the carrier service code of the specified shippingCompany. If carrierServiceCode is provided &#x3D;&gt; the subscription will be applied to the carrier service code.  | [optional] 
**ElectronicChannel** | Pointer to **string** | Dedicated channel where the booking request is submitted, allowed values -&gt; API, WEB, INTTRA_ACT, INTTRA_LINK | [optional] 
**StartDate** | Pointer to **string** | Start date of the subscription. By default, today | [optional] 
**EndDate** | Pointer to **string** | End date of the subscription. By default, today + 120 days / today + 365 days for electronic channel | [optional] 

## Methods

### NewSubscriptionRule

`func NewSubscriptionRule() *SubscriptionRule`

NewSubscriptionRule instantiates a new SubscriptionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionRuleWithDefaults

`func NewSubscriptionRuleWithDefaults() *SubscriptionRule`

NewSubscriptionRuleWithDefaults instantiates a new SubscriptionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEquipmentReferences

`func (o *SubscriptionRule) GetEquipmentReferences() []string`

GetEquipmentReferences returns the EquipmentReferences field if non-nil, zero value otherwise.

### GetEquipmentReferencesOk

`func (o *SubscriptionRule) GetEquipmentReferencesOk() (*[]string, bool)`

GetEquipmentReferencesOk returns a tuple with the EquipmentReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentReferences

`func (o *SubscriptionRule) SetEquipmentReferences(v []string)`

SetEquipmentReferences sets EquipmentReferences field to given value.

### HasEquipmentReferences

`func (o *SubscriptionRule) HasEquipmentReferences() bool`

HasEquipmentReferences returns a boolean if a field has been set.

### GetVas

`func (o *SubscriptionRule) GetVas() []string`

GetVas returns the Vas field if non-nil, zero value otherwise.

### GetVasOk

`func (o *SubscriptionRule) GetVasOk() (*[]string, bool)`

GetVasOk returns a tuple with the Vas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVas

`func (o *SubscriptionRule) SetVas(v []string)`

SetVas sets Vas field to given value.

### HasVas

`func (o *SubscriptionRule) HasVas() bool`

HasVas returns a boolean if a field has been set.

### GetShipmentReferences

`func (o *SubscriptionRule) GetShipmentReferences() []string`

GetShipmentReferences returns the ShipmentReferences field if non-nil, zero value otherwise.

### GetShipmentReferencesOk

`func (o *SubscriptionRule) GetShipmentReferencesOk() (*[]string, bool)`

GetShipmentReferencesOk returns a tuple with the ShipmentReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentReferences

`func (o *SubscriptionRule) SetShipmentReferences(v []string)`

SetShipmentReferences sets ShipmentReferences field to given value.

### HasShipmentReferences

`func (o *SubscriptionRule) HasShipmentReferences() bool`

HasShipmentReferences returns a boolean if a field has been set.

### GetContractReferences

`func (o *SubscriptionRule) GetContractReferences() []string`

GetContractReferences returns the ContractReferences field if non-nil, zero value otherwise.

### GetContractReferencesOk

`func (o *SubscriptionRule) GetContractReferencesOk() (*[]string, bool)`

GetContractReferencesOk returns a tuple with the ContractReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractReferences

`func (o *SubscriptionRule) SetContractReferences(v []string)`

SetContractReferences sets ContractReferences field to given value.

### HasContractReferences

`func (o *SubscriptionRule) HasContractReferences() bool`

HasContractReferences returns a boolean if a field has been set.

### GetCvsServices

`func (o *SubscriptionRule) GetCvsServices() []CvsService`

GetCvsServices returns the CvsServices field if non-nil, zero value otherwise.

### GetCvsServicesOk

`func (o *SubscriptionRule) GetCvsServicesOk() (*[]CvsService, bool)`

GetCvsServicesOk returns a tuple with the CvsServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvsServices

`func (o *SubscriptionRule) SetCvsServices(v []CvsService)`

SetCvsServices sets CvsServices field to given value.

### HasCvsServices

`func (o *SubscriptionRule) HasCvsServices() bool`

HasCvsServices returns a boolean if a field has been set.

### GetElectronicChannel

`func (o *SubscriptionRule) GetElectronicChannel() string`

GetElectronicChannel returns the ElectronicChannel field if non-nil, zero value otherwise.

### GetElectronicChannelOk

`func (o *SubscriptionRule) GetElectronicChannelOk() (*string, bool)`

GetElectronicChannelOk returns a tuple with the ElectronicChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectronicChannel

`func (o *SubscriptionRule) SetElectronicChannel(v string)`

SetElectronicChannel sets ElectronicChannel field to given value.

### HasElectronicChannel

`func (o *SubscriptionRule) HasElectronicChannel() bool`

HasElectronicChannel returns a boolean if a field has been set.

### GetStartDate

`func (o *SubscriptionRule) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SubscriptionRule) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SubscriptionRule) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SubscriptionRule) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *SubscriptionRule) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SubscriptionRule) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SubscriptionRule) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *SubscriptionRule) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


