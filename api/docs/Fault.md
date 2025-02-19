# Fault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | **string** | HTTP error or Free text such as  Database Not Available , Invalid Credential, Mandatory, Invalid Format,  Invalid Length, ForbiddenValue, Invalid Data,... | 
**Code** | Pointer to **string** | Error Id,  Server Application error ID, Oracle error Id, .... | [optional] 
**Description** | Pointer to **string** | Human-readable error description including Data and Value | [optional] 

## Methods

### NewFault

`func NewFault(reason string, ) *Fault`

NewFault instantiates a new Fault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaultWithDefaults

`func NewFaultWithDefaults() *Fault`

NewFaultWithDefaults instantiates a new Fault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *Fault) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Fault) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Fault) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCode

`func (o *Fault) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Fault) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Fault) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Fault) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *Fault) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Fault) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Fault) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Fault) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


