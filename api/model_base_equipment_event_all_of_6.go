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

// BaseEquipmentEventAllOf6 struct for BaseEquipmentEventAllOf6
type BaseEquipmentEventAllOf6 struct {
	EventLocation *Location `json:"eventLocation,omitempty"`
}

// NewBaseEquipmentEventAllOf6 instantiates a new BaseEquipmentEventAllOf6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEquipmentEventAllOf6() *BaseEquipmentEventAllOf6 {
	this := BaseEquipmentEventAllOf6{}
	return &this
}

// NewBaseEquipmentEventAllOf6WithDefaults instantiates a new BaseEquipmentEventAllOf6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEquipmentEventAllOf6WithDefaults() *BaseEquipmentEventAllOf6 {
	this := BaseEquipmentEventAllOf6{}
	return &this
}

// GetEventLocation returns the EventLocation field value if set, zero value otherwise.
func (o *BaseEquipmentEventAllOf6) GetEventLocation() Location {
	if o == nil || o.EventLocation == nil {
		var ret Location
		return ret
	}
	return *o.EventLocation
}

// GetEventLocationOk returns a tuple with the EventLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEquipmentEventAllOf6) GetEventLocationOk() (*Location, bool) {
	if o == nil || o.EventLocation == nil {
		return nil, false
	}
	return o.EventLocation, true
}

// HasEventLocation returns a boolean if a field has been set.
func (o *BaseEquipmentEventAllOf6) HasEventLocation() bool {
	if o != nil && o.EventLocation != nil {
		return true
	}

	return false
}

// SetEventLocation gets a reference to the given Location and assigns it to the EventLocation field.
func (o *BaseEquipmentEventAllOf6) SetEventLocation(v Location) {
	o.EventLocation = &v
}

func (o BaseEquipmentEventAllOf6) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventLocation != nil {
		toSerialize["eventLocation"] = o.EventLocation
	}
	return json.Marshal(toSerialize)
}

type NullableBaseEquipmentEventAllOf6 struct {
	value *BaseEquipmentEventAllOf6
	isSet bool
}

func (v NullableBaseEquipmentEventAllOf6) Get() *BaseEquipmentEventAllOf6 {
	return v.value
}

func (v *NullableBaseEquipmentEventAllOf6) Set(val *BaseEquipmentEventAllOf6) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEquipmentEventAllOf6) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEquipmentEventAllOf6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEquipmentEventAllOf6(val *BaseEquipmentEventAllOf6) *NullableBaseEquipmentEventAllOf6 {
	return &NullableBaseEquipmentEventAllOf6{value: val, isSet: true}
}

func (v NullableBaseEquipmentEventAllOf6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEquipmentEventAllOf6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
