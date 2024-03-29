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

// VesselAllOf struct for VesselAllOf
type VesselAllOf struct {
	// The unique reference for a registered Vessel. The reference is the International Maritime Organisation (IMO) number, also sometimes known as the Lloyd&apos;s register code, which does not change during the lifetime of the vessel
	VesselIMONumber *string `json:"vesselIMONumber,omitempty"`
}

// NewVesselAllOf instantiates a new VesselAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVesselAllOf() *VesselAllOf {
	this := VesselAllOf{}
	return &this
}

// NewVesselAllOfWithDefaults instantiates a new VesselAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVesselAllOfWithDefaults() *VesselAllOf {
	this := VesselAllOf{}
	return &this
}

// GetVesselIMONumber returns the VesselIMONumber field value if set, zero value otherwise.
func (o *VesselAllOf) GetVesselIMONumber() string {
	if o == nil || o.VesselIMONumber == nil {
		var ret string
		return ret
	}
	return *o.VesselIMONumber
}

// GetVesselIMONumberOk returns a tuple with the VesselIMONumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VesselAllOf) GetVesselIMONumberOk() (*string, bool) {
	if o == nil || o.VesselIMONumber == nil {
		return nil, false
	}
	return o.VesselIMONumber, true
}

// HasVesselIMONumber returns a boolean if a field has been set.
func (o *VesselAllOf) HasVesselIMONumber() bool {
	if o != nil && o.VesselIMONumber != nil {
		return true
	}

	return false
}

// SetVesselIMONumber gets a reference to the given string and assigns it to the VesselIMONumber field.
func (o *VesselAllOf) SetVesselIMONumber(v string) {
	o.VesselIMONumber = &v
}

func (o VesselAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VesselIMONumber != nil {
		toSerialize["vesselIMONumber"] = o.VesselIMONumber
	}
	return json.Marshal(toSerialize)
}

type NullableVesselAllOf struct {
	value *VesselAllOf
	isSet bool
}

func (v NullableVesselAllOf) Get() *VesselAllOf {
	return v.value
}

func (v *NullableVesselAllOf) Set(val *VesselAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVesselAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVesselAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVesselAllOf(val *VesselAllOf) *NullableVesselAllOf {
	return &NullableVesselAllOf{value: val, isSet: true}
}

func (v NullableVesselAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVesselAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
