/*
Logistic Tracking service API - DCSA OpenAPI specification for Track & Trace v2.2.0

Managing and sending Shipment-, Transport- and Equipment-events and subscriptions for Track &amp; Trace (T&amp;T). API specification issued by DCSA.org.  <i>Please note that shipment events and subscription management are not covered yet by CMA CGM API.</i>  <br> This API is accessible through <ul> <li> <b> Public </b> connection (authentication with API Key): this allows to retrieve standard equipment moves (e.g. ready to be loaded, discharged), Transhipment moves (e.g. discharged and re-loaded onboard) and planned vessel departure and arrival dates. </li> <li> <b> Private </b> connection (authentication with Oauth2 client credentials): this allows to retrieve additional private events such as actual rail and ramp moves, planned departure and arrival dates for inland moves. These private events are available only if you (or your end customer) are identified on the booking as: Booking Party, Deciding Party, Consignee, Notifier or Shipper. </li> </ul> </br> For explanation to specific values or objects please refer to the <a href='https://dcsa.org/wp-content/uploads/2021/10/202108_DCSA_P1_Information-Model-v3.3_TNT22_Final.pdf'>Information Model v3.3</a> Polling can be done on the <b>GET /events</b> endPoint. DCSA version 2.2.0

API version: 1.2.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// BaseEventAllOf1 struct for BaseEventAllOf1
type BaseEventAllOf1 struct {
	// The timestamp of when the event was created. <b>NB</b>&#58; This field should be considered Metadata
	EventCreatedDateTime *time.Time `json:"eventCreatedDateTime,omitempty"`
}

// NewBaseEventAllOf1 instantiates a new BaseEventAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEventAllOf1() *BaseEventAllOf1 {
	this := BaseEventAllOf1{}
	return &this
}

// NewBaseEventAllOf1WithDefaults instantiates a new BaseEventAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEventAllOf1WithDefaults() *BaseEventAllOf1 {
	this := BaseEventAllOf1{}
	return &this
}

// GetEventCreatedDateTime returns the EventCreatedDateTime field value if set, zero value otherwise.
func (o *BaseEventAllOf1) GetEventCreatedDateTime() time.Time {
	if o == nil || o.EventCreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EventCreatedDateTime
}

// GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEventAllOf1) GetEventCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.EventCreatedDateTime == nil {
		return nil, false
	}
	return o.EventCreatedDateTime, true
}

// HasEventCreatedDateTime returns a boolean if a field has been set.
func (o *BaseEventAllOf1) HasEventCreatedDateTime() bool {
	if o != nil && o.EventCreatedDateTime != nil {
		return true
	}

	return false
}

// SetEventCreatedDateTime gets a reference to the given time.Time and assigns it to the EventCreatedDateTime field.
func (o *BaseEventAllOf1) SetEventCreatedDateTime(v time.Time) {
	o.EventCreatedDateTime = &v
}

func (o BaseEventAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventCreatedDateTime != nil {
		toSerialize["eventCreatedDateTime"] = o.EventCreatedDateTime
	}
	return json.Marshal(toSerialize)
}

type NullableBaseEventAllOf1 struct {
	value *BaseEventAllOf1
	isSet bool
}

func (v NullableBaseEventAllOf1) Get() *BaseEventAllOf1 {
	return v.value
}

func (v *NullableBaseEventAllOf1) Set(val *BaseEventAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEventAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEventAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEventAllOf1(val *BaseEventAllOf1) *NullableBaseEventAllOf1 {
	return &NullableBaseEventAllOf1{value: val, isSet: true}
}

func (v NullableBaseEventAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEventAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
