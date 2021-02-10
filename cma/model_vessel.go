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

// Vessel struct for Vessel
type Vessel struct {
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
	Imo *string `json:"imo,omitempty"`
}

// NewVessel instantiates a new Vessel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVessel() *Vessel {
	this := Vessel{}
	return &this
}

// NewVesselWithDefaults instantiates a new Vessel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVesselWithDefaults() *Vessel {
	this := Vessel{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Vessel) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vessel) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Vessel) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Vessel) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Vessel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vessel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Vessel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Vessel) SetName(v string) {
	o.Name = &v
}

// GetImo returns the Imo field value if set, zero value otherwise.
func (o *Vessel) GetImo() string {
	if o == nil || o.Imo == nil {
		var ret string
		return ret
	}
	return *o.Imo
}

// GetImoOk returns a tuple with the Imo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vessel) GetImoOk() (*string, bool) {
	if o == nil || o.Imo == nil {
		return nil, false
	}
	return o.Imo, true
}

// HasImo returns a boolean if a field has been set.
func (o *Vessel) HasImo() bool {
	if o != nil && o.Imo != nil {
		return true
	}

	return false
}

// SetImo gets a reference to the given string and assigns it to the Imo field.
func (o *Vessel) SetImo(v string) {
	o.Imo = &v
}

func (o Vessel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Imo != nil {
		toSerialize["imo"] = o.Imo
	}
	return json.Marshal(toSerialize)
}

type NullableVessel struct {
	value *Vessel
	isSet bool
}

func (v NullableVessel) Get() *Vessel {
	return v.value
}

func (v *NullableVessel) Set(val *Vessel) {
	v.value = val
	v.isSet = true
}

func (v NullableVessel) IsSet() bool {
	return v.isSet
}

func (v *NullableVessel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVessel(val *Vessel) *NullableVessel {
	return &NullableVessel{value: val, isSet: true}
}

func (v NullableVessel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVessel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


