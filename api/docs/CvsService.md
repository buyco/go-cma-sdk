# CvsService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceCode** | Pointer to **string** | Service code of voyage | [optional] 
**CarrierServiceCode** | Pointer to **string** | Service code for a shipping company | [optional] 
**ShippingCompany** | Pointer to **string** | Shipping company code | [optional] 
**ServiceStartDate** | Pointer to **string** | Start date of the service subscription. By default, today | [optional] 
**ServiceEndDate** | Pointer to **string** | End date of the service subscription. By default, today + 1 year | [optional] 

## Methods

### NewCvsService

`func NewCvsService() *CvsService`

NewCvsService instantiates a new CvsService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCvsServiceWithDefaults

`func NewCvsServiceWithDefaults() *CvsService`

NewCvsServiceWithDefaults instantiates a new CvsService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceCode

`func (o *CvsService) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *CvsService) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *CvsService) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.

### HasServiceCode

`func (o *CvsService) HasServiceCode() bool`

HasServiceCode returns a boolean if a field has been set.

### GetCarrierServiceCode

`func (o *CvsService) GetCarrierServiceCode() string`

GetCarrierServiceCode returns the CarrierServiceCode field if non-nil, zero value otherwise.

### GetCarrierServiceCodeOk

`func (o *CvsService) GetCarrierServiceCodeOk() (*string, bool)`

GetCarrierServiceCodeOk returns a tuple with the CarrierServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierServiceCode

`func (o *CvsService) SetCarrierServiceCode(v string)`

SetCarrierServiceCode sets CarrierServiceCode field to given value.

### HasCarrierServiceCode

`func (o *CvsService) HasCarrierServiceCode() bool`

HasCarrierServiceCode returns a boolean if a field has been set.

### GetShippingCompany

`func (o *CvsService) GetShippingCompany() string`

GetShippingCompany returns the ShippingCompany field if non-nil, zero value otherwise.

### GetShippingCompanyOk

`func (o *CvsService) GetShippingCompanyOk() (*string, bool)`

GetShippingCompanyOk returns a tuple with the ShippingCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCompany

`func (o *CvsService) SetShippingCompany(v string)`

SetShippingCompany sets ShippingCompany field to given value.

### HasShippingCompany

`func (o *CvsService) HasShippingCompany() bool`

HasShippingCompany returns a boolean if a field has been set.

### GetServiceStartDate

`func (o *CvsService) GetServiceStartDate() string`

GetServiceStartDate returns the ServiceStartDate field if non-nil, zero value otherwise.

### GetServiceStartDateOk

`func (o *CvsService) GetServiceStartDateOk() (*string, bool)`

GetServiceStartDateOk returns a tuple with the ServiceStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStartDate

`func (o *CvsService) SetServiceStartDate(v string)`

SetServiceStartDate sets ServiceStartDate field to given value.

### HasServiceStartDate

`func (o *CvsService) HasServiceStartDate() bool`

HasServiceStartDate returns a boolean if a field has been set.

### GetServiceEndDate

`func (o *CvsService) GetServiceEndDate() string`

GetServiceEndDate returns the ServiceEndDate field if non-nil, zero value otherwise.

### GetServiceEndDateOk

`func (o *CvsService) GetServiceEndDateOk() (*string, bool)`

GetServiceEndDateOk returns a tuple with the ServiceEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndDate

`func (o *CvsService) SetServiceEndDate(v string)`

SetServiceEndDate sets ServiceEndDate field to given value.

### HasServiceEndDate

`func (o *CvsService) HasServiceEndDate() bool`

HasServiceEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


