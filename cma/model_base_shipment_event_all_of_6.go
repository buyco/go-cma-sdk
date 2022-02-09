/*
Logistic Tracking service API - DCSA OpenAPI specification for Track & Trace v2.2.0

Managing and sending Shipment-, Transport- and Equipment-events and subscriptions for Track &amp; Trace (T&amp;T). API specification issued by DCSA.org. For explanation to specific values or objects please refer to the <a href='https://dcsa.org/wp-content/uploads/2021/10/202108_DCSA_P1_Information-Model-v3.3_TNT22_Final.pdf'>Information Model v3.3</a> Polling can be done on the <b>GET /events</b> endPoint. DCSA version 2.2.0

API version: 1.0.3
Contact: ho.support-api@cma-cgm.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cma

import (
	"encoding/json"
)

// BaseShipmentEventAllOf6 struct for BaseShipmentEventAllOf6
type BaseShipmentEventAllOf6 struct {
	// The identifier for a shipment
	ShipmentID *string `json:"shipmentID,omitempty"`
}

// NewBaseShipmentEventAllOf6 instantiates a new BaseShipmentEventAllOf6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseShipmentEventAllOf6() *BaseShipmentEventAllOf6 {
	this := BaseShipmentEventAllOf6{}
	return &this
}

// NewBaseShipmentEventAllOf6WithDefaults instantiates a new BaseShipmentEventAllOf6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseShipmentEventAllOf6WithDefaults() *BaseShipmentEventAllOf6 {
	this := BaseShipmentEventAllOf6{}
	return &this
}

// GetShipmentID returns the ShipmentID field value if set, zero value otherwise.
func (o *BaseShipmentEventAllOf6) GetShipmentID() string {
	if o == nil || o.ShipmentID == nil {
		var ret string
		return ret
	}
	return *o.ShipmentID
}

// GetShipmentIDOk returns a tuple with the ShipmentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseShipmentEventAllOf6) GetShipmentIDOk() (*string, bool) {
	if o == nil || o.ShipmentID == nil {
		return nil, false
	}
	return o.ShipmentID, true
}

// HasShipmentID returns a boolean if a field has been set.
func (o *BaseShipmentEventAllOf6) HasShipmentID() bool {
	if o != nil && o.ShipmentID != nil {
		return true
	}

	return false
}

// SetShipmentID gets a reference to the given string and assigns it to the ShipmentID field.
func (o *BaseShipmentEventAllOf6) SetShipmentID(v string) {
	o.ShipmentID = &v
}

func (o BaseShipmentEventAllOf6) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ShipmentID != nil {
		toSerialize["shipmentID"] = o.ShipmentID
	}
	return json.Marshal(toSerialize)
}

type NullableBaseShipmentEventAllOf6 struct {
	value *BaseShipmentEventAllOf6
	isSet bool
}

func (v NullableBaseShipmentEventAllOf6) Get() *BaseShipmentEventAllOf6 {
	return v.value
}

func (v *NullableBaseShipmentEventAllOf6) Set(val *BaseShipmentEventAllOf6) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseShipmentEventAllOf6) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseShipmentEventAllOf6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseShipmentEventAllOf6(val *BaseShipmentEventAllOf6) *NullableBaseShipmentEventAllOf6 {
	return &NullableBaseShipmentEventAllOf6{value: val, isSet: true}
}

func (v NullableBaseShipmentEventAllOf6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseShipmentEventAllOf6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


