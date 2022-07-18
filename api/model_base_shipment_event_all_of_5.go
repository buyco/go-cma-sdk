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

// BaseShipmentEventAllOf5 struct for BaseShipmentEventAllOf5
type BaseShipmentEventAllOf5 struct {
	// Reason field in a Shipment event. This field can be used to explain why a specific event has been sent.
	Reason *string `json:"reason,omitempty"`
}

// NewBaseShipmentEventAllOf5 instantiates a new BaseShipmentEventAllOf5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseShipmentEventAllOf5() *BaseShipmentEventAllOf5 {
	this := BaseShipmentEventAllOf5{}
	return &this
}

// NewBaseShipmentEventAllOf5WithDefaults instantiates a new BaseShipmentEventAllOf5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseShipmentEventAllOf5WithDefaults() *BaseShipmentEventAllOf5 {
	this := BaseShipmentEventAllOf5{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *BaseShipmentEventAllOf5) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseShipmentEventAllOf5) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *BaseShipmentEventAllOf5) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *BaseShipmentEventAllOf5) SetReason(v string) {
	o.Reason = &v
}

func (o BaseShipmentEventAllOf5) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableBaseShipmentEventAllOf5 struct {
	value *BaseShipmentEventAllOf5
	isSet bool
}

func (v NullableBaseShipmentEventAllOf5) Get() *BaseShipmentEventAllOf5 {
	return v.value
}

func (v *NullableBaseShipmentEventAllOf5) Set(val *BaseShipmentEventAllOf5) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseShipmentEventAllOf5) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseShipmentEventAllOf5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseShipmentEventAllOf5(val *BaseShipmentEventAllOf5) *NullableBaseShipmentEventAllOf5 {
	return &NullableBaseShipmentEventAllOf5{value: val, isSet: true}
}

func (v NullableBaseShipmentEventAllOf5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseShipmentEventAllOf5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

