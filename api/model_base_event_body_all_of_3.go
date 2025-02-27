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

// BaseEventBodyAllOf3 struct for BaseEventBodyAllOf3
type BaseEventBodyAllOf3 struct {
	CarrierSpecificData *CarrierSpecificData `json:"carrierSpecificData,omitempty"`
}

// NewBaseEventBodyAllOf3 instantiates a new BaseEventBodyAllOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEventBodyAllOf3() *BaseEventBodyAllOf3 {
	this := BaseEventBodyAllOf3{}
	return &this
}

// NewBaseEventBodyAllOf3WithDefaults instantiates a new BaseEventBodyAllOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEventBodyAllOf3WithDefaults() *BaseEventBodyAllOf3 {
	this := BaseEventBodyAllOf3{}
	return &this
}

// GetCarrierSpecificData returns the CarrierSpecificData field value if set, zero value otherwise.
func (o *BaseEventBodyAllOf3) GetCarrierSpecificData() CarrierSpecificData {
	if o == nil || o.CarrierSpecificData == nil {
		var ret CarrierSpecificData
		return ret
	}
	return *o.CarrierSpecificData
}

// GetCarrierSpecificDataOk returns a tuple with the CarrierSpecificData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEventBodyAllOf3) GetCarrierSpecificDataOk() (*CarrierSpecificData, bool) {
	if o == nil || o.CarrierSpecificData == nil {
		return nil, false
	}
	return o.CarrierSpecificData, true
}

// HasCarrierSpecificData returns a boolean if a field has been set.
func (o *BaseEventBodyAllOf3) HasCarrierSpecificData() bool {
	if o != nil && o.CarrierSpecificData != nil {
		return true
	}

	return false
}

// SetCarrierSpecificData gets a reference to the given CarrierSpecificData and assigns it to the CarrierSpecificData field.
func (o *BaseEventBodyAllOf3) SetCarrierSpecificData(v CarrierSpecificData) {
	o.CarrierSpecificData = &v
}

func (o BaseEventBodyAllOf3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CarrierSpecificData != nil {
		toSerialize["carrierSpecificData"] = o.CarrierSpecificData
	}
	return json.Marshal(toSerialize)
}

type NullableBaseEventBodyAllOf3 struct {
	value *BaseEventBodyAllOf3
	isSet bool
}

func (v NullableBaseEventBodyAllOf3) Get() *BaseEventBodyAllOf3 {
	return v.value
}

func (v *NullableBaseEventBodyAllOf3) Set(val *BaseEventBodyAllOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEventBodyAllOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEventBodyAllOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEventBodyAllOf3(val *BaseEventBodyAllOf3) *NullableBaseEventBodyAllOf3 {
	return &NullableBaseEventBodyAllOf3{value: val, isSet: true}
}

func (v NullableBaseEventBodyAllOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEventBodyAllOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
