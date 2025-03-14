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

// BaseEquipmentEventAllOf4 struct for BaseEquipmentEventAllOf4
type BaseEquipmentEventAllOf4 struct {
	// Unique code for the different equipment size/type used for transporting commodities. The code is a concatenation of ISO Equipment Size Code and ISO Equipment Type Code A and follows the ISO 6346 standard.
	ISOEquipmentCode *string `json:"ISOEquipmentCode,omitempty"`
}

// NewBaseEquipmentEventAllOf4 instantiates a new BaseEquipmentEventAllOf4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEquipmentEventAllOf4() *BaseEquipmentEventAllOf4 {
	this := BaseEquipmentEventAllOf4{}
	return &this
}

// NewBaseEquipmentEventAllOf4WithDefaults instantiates a new BaseEquipmentEventAllOf4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEquipmentEventAllOf4WithDefaults() *BaseEquipmentEventAllOf4 {
	this := BaseEquipmentEventAllOf4{}
	return &this
}

// GetISOEquipmentCode returns the ISOEquipmentCode field value if set, zero value otherwise.
func (o *BaseEquipmentEventAllOf4) GetISOEquipmentCode() string {
	if o == nil || o.ISOEquipmentCode == nil {
		var ret string
		return ret
	}
	return *o.ISOEquipmentCode
}

// GetISOEquipmentCodeOk returns a tuple with the ISOEquipmentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEquipmentEventAllOf4) GetISOEquipmentCodeOk() (*string, bool) {
	if o == nil || o.ISOEquipmentCode == nil {
		return nil, false
	}
	return o.ISOEquipmentCode, true
}

// HasISOEquipmentCode returns a boolean if a field has been set.
func (o *BaseEquipmentEventAllOf4) HasISOEquipmentCode() bool {
	if o != nil && o.ISOEquipmentCode != nil {
		return true
	}

	return false
}

// SetISOEquipmentCode gets a reference to the given string and assigns it to the ISOEquipmentCode field.
func (o *BaseEquipmentEventAllOf4) SetISOEquipmentCode(v string) {
	o.ISOEquipmentCode = &v
}

func (o BaseEquipmentEventAllOf4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ISOEquipmentCode != nil {
		toSerialize["ISOEquipmentCode"] = o.ISOEquipmentCode
	}
	return json.Marshal(toSerialize)
}

type NullableBaseEquipmentEventAllOf4 struct {
	value *BaseEquipmentEventAllOf4
	isSet bool
}

func (v NullableBaseEquipmentEventAllOf4) Get() *BaseEquipmentEventAllOf4 {
	return v.value
}

func (v *NullableBaseEquipmentEventAllOf4) Set(val *BaseEquipmentEventAllOf4) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEquipmentEventAllOf4) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEquipmentEventAllOf4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEquipmentEventAllOf4(val *BaseEquipmentEventAllOf4) *NullableBaseEquipmentEventAllOf4 {
	return &NullableBaseEquipmentEventAllOf4{value: val, isSet: true}
}

func (v NullableBaseEquipmentEventAllOf4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEquipmentEventAllOf4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
