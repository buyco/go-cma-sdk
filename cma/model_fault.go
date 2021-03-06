/*
 * logistic.tracking.v1
 *
 * Retrieve tracking data about your cargo using container number, booking number or bill of lading.
 *
 * API version: 1.4.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cma

import (
	"encoding/json"
)

// Fault struct for Fault
type Fault struct {
	// HTTP error or Free text such as  Database Not Available , Invalid Credential, Mandatory, Invalid Format,  Invalid Length, ForbiddenValue, Invalid Data,...
	Reason string `json:"reason"`
	// Error Id,  Server Application error ID, Oracle error Id, ....
	Code *string `json:"code,omitempty"`
	// Human-readable error description including Data and Value
	Description *string `json:"description,omitempty"`
}

// NewFault instantiates a new Fault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFault(reason string, ) *Fault {
	this := Fault{}
	this.Reason = reason
	return &this
}

// NewFaultWithDefaults instantiates a new Fault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFaultWithDefaults() *Fault {
	this := Fault{}
	return &this
}

// GetReason returns the Reason field value
func (o *Fault) GetReason() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *Fault) GetReasonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *Fault) SetReason(v string) {
	o.Reason = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Fault) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fault) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Fault) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Fault) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Fault) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Fault) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Fault) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Fault) SetDescription(v string) {
	o.Description = &v
}

func (o Fault) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reason"] = o.Reason
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableFault struct {
	value *Fault
	isSet bool
}

func (v NullableFault) Get() *Fault {
	return v.value
}

func (v *NullableFault) Set(val *Fault) {
	v.value = val
	v.isSet = true
}

func (v NullableFault) IsSet() bool {
	return v.isSet
}

func (v *NullableFault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFault(val *Fault) *NullableFault {
	return &NullableFault{value: val, isSet: true}
}

func (v NullableFault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


