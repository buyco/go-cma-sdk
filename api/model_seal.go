/*
Logistic Tracking service API - DCSA OpenAPI specification for Track & Trace v2.2.0

Managing and sending Shipment-, Transport- and Equipment-events and subscriptions for Track &amp; Trace (T&amp;T). API specification issued by DCSA.org. For explanation to specific values or objects please refer to the <a href='https://dcsa.org/wp-content/uploads/2021/10/202108_DCSA_P1_Information-Model-v3.3_TNT22_Final.pdf'>Information Model v3.3</a> Polling can be done on the <b>GET /events</b> endPoint. DCSA version 2.2.0

API version: 1.0.3
Contact: ho.support-api@cma-cgm.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Seal addresses the seal-related information associated with the shipment equipment. A seal is put on a shipment equipment once it is loaded. This seal is meant to stay on until the shipment equipment reaches its final destination.
type Seal struct {
	// Identifies a seal affixed to the container.
	SealNumber string      `json:"sealNumber"`
	SealSource *SealSource `json:"sealSource,omitempty"`
	SealType   SealType    `json:"sealType"`
}

// NewSeal instantiates a new Seal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeal(sealNumber string, sealType SealType) *Seal {
	this := Seal{}
	this.SealNumber = sealNumber
	this.SealType = sealType
	return &this
}

// NewSealWithDefaults instantiates a new Seal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSealWithDefaults() *Seal {
	this := Seal{}
	return &this
}

// GetSealNumber returns the SealNumber field value
func (o *Seal) GetSealNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SealNumber
}

// GetSealNumberOk returns a tuple with the SealNumber field value
// and a boolean to check if the value has been set.
func (o *Seal) GetSealNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SealNumber, true
}

// SetSealNumber sets field value
func (o *Seal) SetSealNumber(v string) {
	o.SealNumber = v
}

// GetSealSource returns the SealSource field value if set, zero value otherwise.
func (o *Seal) GetSealSource() SealSource {
	if o == nil || o.SealSource == nil {
		var ret SealSource
		return ret
	}
	return *o.SealSource
}

// GetSealSourceOk returns a tuple with the SealSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Seal) GetSealSourceOk() (*SealSource, bool) {
	if o == nil || o.SealSource == nil {
		return nil, false
	}
	return o.SealSource, true
}

// HasSealSource returns a boolean if a field has been set.
func (o *Seal) HasSealSource() bool {
	if o != nil && o.SealSource != nil {
		return true
	}

	return false
}

// SetSealSource gets a reference to the given SealSource and assigns it to the SealSource field.
func (o *Seal) SetSealSource(v SealSource) {
	o.SealSource = &v
}

// GetSealType returns the SealType field value
func (o *Seal) GetSealType() SealType {
	if o == nil {
		var ret SealType
		return ret
	}

	return o.SealType
}

// GetSealTypeOk returns a tuple with the SealType field value
// and a boolean to check if the value has been set.
func (o *Seal) GetSealTypeOk() (*SealType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SealType, true
}

// SetSealType sets field value
func (o *Seal) SetSealType(v SealType) {
	o.SealType = v
}

func (o Seal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sealNumber"] = o.SealNumber
	}
	if o.SealSource != nil {
		toSerialize["sealSource"] = o.SealSource
	}
	if true {
		toSerialize["sealType"] = o.SealType
	}
	return json.Marshal(toSerialize)
}

type NullableSeal struct {
	value *Seal
	isSet bool
}

func (v NullableSeal) Get() *Seal {
	return v.value
}

func (v *NullableSeal) Set(val *Seal) {
	v.value = val
	v.isSet = true
}

func (v NullableSeal) IsSet() bool {
	return v.isSet
}

func (v *NullableSeal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeal(val *Seal) *NullableSeal {
	return &NullableSeal{value: val, isSet: true}
}

func (v NullableSeal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
