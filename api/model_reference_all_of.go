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

// ReferenceAllOf struct for ReferenceAllOf
type ReferenceAllOf struct {
	ReferenceType *ReferenceType `json:"referenceType,omitempty"`
}

// NewReferenceAllOf instantiates a new ReferenceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferenceAllOf() *ReferenceAllOf {
	this := ReferenceAllOf{}
	return &this
}

// NewReferenceAllOfWithDefaults instantiates a new ReferenceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferenceAllOfWithDefaults() *ReferenceAllOf {
	this := ReferenceAllOf{}
	return &this
}

// GetReferenceType returns the ReferenceType field value if set, zero value otherwise.
func (o *ReferenceAllOf) GetReferenceType() ReferenceType {
	if o == nil || o.ReferenceType == nil {
		var ret ReferenceType
		return ret
	}
	return *o.ReferenceType
}

// GetReferenceTypeOk returns a tuple with the ReferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferenceAllOf) GetReferenceTypeOk() (*ReferenceType, bool) {
	if o == nil || o.ReferenceType == nil {
		return nil, false
	}
	return o.ReferenceType, true
}

// HasReferenceType returns a boolean if a field has been set.
func (o *ReferenceAllOf) HasReferenceType() bool {
	if o != nil && o.ReferenceType != nil {
		return true
	}

	return false
}

// SetReferenceType gets a reference to the given ReferenceType and assigns it to the ReferenceType field.
func (o *ReferenceAllOf) SetReferenceType(v ReferenceType) {
	o.ReferenceType = &v
}

func (o ReferenceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReferenceType != nil {
		toSerialize["referenceType"] = o.ReferenceType
	}
	return json.Marshal(toSerialize)
}

type NullableReferenceAllOf struct {
	value *ReferenceAllOf
	isSet bool
}

func (v NullableReferenceAllOf) Get() *ReferenceAllOf {
	return v.value
}

func (v *NullableReferenceAllOf) Set(val *ReferenceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReferenceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReferenceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferenceAllOf(val *ReferenceAllOf) *NullableReferenceAllOf {
	return &NullableReferenceAllOf{value: val, isSet: true}
}

func (v NullableReferenceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferenceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
