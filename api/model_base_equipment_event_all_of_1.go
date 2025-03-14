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

// BaseEquipmentEventAllOf1 struct for BaseEquipmentEventAllOf1
type BaseEquipmentEventAllOf1 struct {
	// Code for the event classifier can be - PLN (Planned) - ACT (Actual) - EST (Estimated)
	EventClassifierCode *string `json:"eventClassifierCode,omitempty"`
}

// NewBaseEquipmentEventAllOf1 instantiates a new BaseEquipmentEventAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEquipmentEventAllOf1() *BaseEquipmentEventAllOf1 {
	this := BaseEquipmentEventAllOf1{}
	return &this
}

// NewBaseEquipmentEventAllOf1WithDefaults instantiates a new BaseEquipmentEventAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEquipmentEventAllOf1WithDefaults() *BaseEquipmentEventAllOf1 {
	this := BaseEquipmentEventAllOf1{}
	return &this
}

// GetEventClassifierCode returns the EventClassifierCode field value if set, zero value otherwise.
func (o *BaseEquipmentEventAllOf1) GetEventClassifierCode() string {
	if o == nil || o.EventClassifierCode == nil {
		var ret string
		return ret
	}
	return *o.EventClassifierCode
}

// GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEquipmentEventAllOf1) GetEventClassifierCodeOk() (*string, bool) {
	if o == nil || o.EventClassifierCode == nil {
		return nil, false
	}
	return o.EventClassifierCode, true
}

// HasEventClassifierCode returns a boolean if a field has been set.
func (o *BaseEquipmentEventAllOf1) HasEventClassifierCode() bool {
	if o != nil && o.EventClassifierCode != nil {
		return true
	}

	return false
}

// SetEventClassifierCode gets a reference to the given string and assigns it to the EventClassifierCode field.
func (o *BaseEquipmentEventAllOf1) SetEventClassifierCode(v string) {
	o.EventClassifierCode = &v
}

func (o BaseEquipmentEventAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventClassifierCode != nil {
		toSerialize["eventClassifierCode"] = o.EventClassifierCode
	}
	return json.Marshal(toSerialize)
}

type NullableBaseEquipmentEventAllOf1 struct {
	value *BaseEquipmentEventAllOf1
	isSet bool
}

func (v NullableBaseEquipmentEventAllOf1) Get() *BaseEquipmentEventAllOf1 {
	return v.value
}

func (v *NullableBaseEquipmentEventAllOf1) Set(val *BaseEquipmentEventAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEquipmentEventAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEquipmentEventAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEquipmentEventAllOf1(val *BaseEquipmentEventAllOf1) *NullableBaseEquipmentEventAllOf1 {
	return &NullableBaseEquipmentEventAllOf1{value: val, isSet: true}
}

func (v NullableBaseEquipmentEventAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEquipmentEventAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
