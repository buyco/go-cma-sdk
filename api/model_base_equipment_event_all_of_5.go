/*
Logistic Tracking service API - DCSA OpenAPI specification for Track & Trace v2.2.0

Managing and sending Shipment-, Transport- and Equipment-events and subscriptions for Track &amp; Trace (T&amp;T). API specification issued by DCSA.org.  <i>Please note that shipment events and subscription management are not covered yet by CMA CGM API.</i>  <br> This API is accessible through <ul> <li> <b> Public </b> connection (authentication with API Key): this allows to retrieve standard equipment moves (e.g. ready to be loaded, discharged), Transhipment moves (e.g. discharged and re-loaded onboard) and planned vessel departure and arrival dates. </li> <li> <b> Private </b> connection (authentication with Oauth2 client credentials): this allows to retrieve additional private events such as actual rail and ramp moves, planned departure and arrival dates for inland moves. These private events are available only if you (or your end customer) are identified on the booking as: Booking Party, Deciding Party, Consignee, Notifier or Shipper. </li> </ul> </br> For explanation to specific values or objects please refer to the <a href='https://dcsa.org/wp-content/uploads/2021/10/202108_DCSA_P1_Information-Model-v3.3_TNT22_Final.pdf'>Information Model v3.3</a> Polling can be done on the <b>GET /events</b> endPoint. DCSA version 2.2.0

API version: 1.2.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// BaseEquipmentEventAllOf5 struct for BaseEquipmentEventAllOf5
type BaseEquipmentEventAllOf5 struct {
	EmptyIndicatorCode *EmptyIndicatorCode `json:"emptyIndicatorCode,omitempty"`
}

// NewBaseEquipmentEventAllOf5 instantiates a new BaseEquipmentEventAllOf5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEquipmentEventAllOf5() *BaseEquipmentEventAllOf5 {
	this := BaseEquipmentEventAllOf5{}
	return &this
}

// NewBaseEquipmentEventAllOf5WithDefaults instantiates a new BaseEquipmentEventAllOf5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEquipmentEventAllOf5WithDefaults() *BaseEquipmentEventAllOf5 {
	this := BaseEquipmentEventAllOf5{}
	return &this
}

// GetEmptyIndicatorCode returns the EmptyIndicatorCode field value if set, zero value otherwise.
func (o *BaseEquipmentEventAllOf5) GetEmptyIndicatorCode() EmptyIndicatorCode {
	if o == nil || o.EmptyIndicatorCode == nil {
		var ret EmptyIndicatorCode
		return ret
	}
	return *o.EmptyIndicatorCode
}

// GetEmptyIndicatorCodeOk returns a tuple with the EmptyIndicatorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEquipmentEventAllOf5) GetEmptyIndicatorCodeOk() (*EmptyIndicatorCode, bool) {
	if o == nil || o.EmptyIndicatorCode == nil {
		return nil, false
	}
	return o.EmptyIndicatorCode, true
}

// HasEmptyIndicatorCode returns a boolean if a field has been set.
func (o *BaseEquipmentEventAllOf5) HasEmptyIndicatorCode() bool {
	if o != nil && o.EmptyIndicatorCode != nil {
		return true
	}

	return false
}

// SetEmptyIndicatorCode gets a reference to the given EmptyIndicatorCode and assigns it to the EmptyIndicatorCode field.
func (o *BaseEquipmentEventAllOf5) SetEmptyIndicatorCode(v EmptyIndicatorCode) {
	o.EmptyIndicatorCode = &v
}

func (o BaseEquipmentEventAllOf5) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmptyIndicatorCode != nil {
		toSerialize["emptyIndicatorCode"] = o.EmptyIndicatorCode
	}
	return json.Marshal(toSerialize)
}

type NullableBaseEquipmentEventAllOf5 struct {
	value *BaseEquipmentEventAllOf5
	isSet bool
}

func (v NullableBaseEquipmentEventAllOf5) Get() *BaseEquipmentEventAllOf5 {
	return v.value
}

func (v *NullableBaseEquipmentEventAllOf5) Set(val *BaseEquipmentEventAllOf5) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEquipmentEventAllOf5) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEquipmentEventAllOf5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEquipmentEventAllOf5(val *BaseEquipmentEventAllOf5) *NullableBaseEquipmentEventAllOf5 {
	return &NullableBaseEquipmentEventAllOf5{value: val, isSet: true}
}

func (v NullableBaseEquipmentEventAllOf5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEquipmentEventAllOf5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
