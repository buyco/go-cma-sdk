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
	"time"
)

// BaseEvent struct for BaseEvent
type BaseEvent struct {
	// The unique identifier for the event (the message - not the source). <b>NB</b>&#58; This field should be considered Metadata
	EventID *string `json:"eventID,omitempty"`
	// The timestamp of when the event was created. <b>NB</b>&#58; This field should be considered Metadata
	EventCreatedDateTime time.Time `json:"eventCreatedDateTime"`
	// The Event Type of the object - to be used as a discriminator. <b>NB</b>&#58; This field should be considered Metadata
	EventType string `json:"eventType"`
	// Code for the event classifier. Values can vary depending on eventType
	EventClassifierCode string `json:"eventClassifierCode"`
	// The local date and time, where the event took place or when the event will take place, in ISO 8601 format.
	EventDateTime       time.Time            `json:"eventDateTime"`
	CarrierSpecificData *CarrierSpecificData `json:"carrierSpecificData,omitempty"`
}

// NewBaseEvent instantiates a new BaseEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEvent(eventCreatedDateTime time.Time, eventType string, eventClassifierCode string, eventDateTime time.Time) *BaseEvent {
	this := BaseEvent{}
	this.EventCreatedDateTime = eventCreatedDateTime
	this.EventType = eventType
	this.EventClassifierCode = eventClassifierCode
	this.EventDateTime = eventDateTime
	return &this
}

// NewBaseEventWithDefaults instantiates a new BaseEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEventWithDefaults() *BaseEvent {
	this := BaseEvent{}
	return &this
}

// GetEventID returns the EventID field value if set, zero value otherwise.
func (o *BaseEvent) GetEventID() string {
	if o == nil || o.EventID == nil {
		var ret string
		return ret
	}
	return *o.EventID
}

// GetEventIDOk returns a tuple with the EventID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEvent) GetEventIDOk() (*string, bool) {
	if o == nil || o.EventID == nil {
		return nil, false
	}
	return o.EventID, true
}

// HasEventID returns a boolean if a field has been set.
func (o *BaseEvent) HasEventID() bool {
	if o != nil && o.EventID != nil {
		return true
	}

	return false
}

// SetEventID gets a reference to the given string and assigns it to the EventID field.
func (o *BaseEvent) SetEventID(v string) {
	o.EventID = &v
}

// GetEventCreatedDateTime returns the EventCreatedDateTime field value
func (o *BaseEvent) GetEventCreatedDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventCreatedDateTime
}

// GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field value
// and a boolean to check if the value has been set.
func (o *BaseEvent) GetEventCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCreatedDateTime, true
}

// SetEventCreatedDateTime sets field value
func (o *BaseEvent) SetEventCreatedDateTime(v time.Time) {
	o.EventCreatedDateTime = v
}

// GetEventType returns the EventType field value
func (o *BaseEvent) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *BaseEvent) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *BaseEvent) SetEventType(v string) {
	o.EventType = v
}

// GetEventClassifierCode returns the EventClassifierCode field value
func (o *BaseEvent) GetEventClassifierCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventClassifierCode
}

// GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field value
// and a boolean to check if the value has been set.
func (o *BaseEvent) GetEventClassifierCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventClassifierCode, true
}

// SetEventClassifierCode sets field value
func (o *BaseEvent) SetEventClassifierCode(v string) {
	o.EventClassifierCode = v
}

// GetEventDateTime returns the EventDateTime field value
func (o *BaseEvent) GetEventDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventDateTime
}

// GetEventDateTimeOk returns a tuple with the EventDateTime field value
// and a boolean to check if the value has been set.
func (o *BaseEvent) GetEventDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventDateTime, true
}

// SetEventDateTime sets field value
func (o *BaseEvent) SetEventDateTime(v time.Time) {
	o.EventDateTime = v
}

// GetCarrierSpecificData returns the CarrierSpecificData field value if set, zero value otherwise.
func (o *BaseEvent) GetCarrierSpecificData() CarrierSpecificData {
	if o == nil || o.CarrierSpecificData == nil {
		var ret CarrierSpecificData
		return ret
	}
	return *o.CarrierSpecificData
}

// GetCarrierSpecificDataOk returns a tuple with the CarrierSpecificData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEvent) GetCarrierSpecificDataOk() (*CarrierSpecificData, bool) {
	if o == nil || o.CarrierSpecificData == nil {
		return nil, false
	}
	return o.CarrierSpecificData, true
}

// HasCarrierSpecificData returns a boolean if a field has been set.
func (o *BaseEvent) HasCarrierSpecificData() bool {
	if o != nil && o.CarrierSpecificData != nil {
		return true
	}

	return false
}

// SetCarrierSpecificData gets a reference to the given CarrierSpecificData and assigns it to the CarrierSpecificData field.
func (o *BaseEvent) SetCarrierSpecificData(v CarrierSpecificData) {
	o.CarrierSpecificData = &v
}

func (o BaseEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventID != nil {
		toSerialize["eventID"] = o.EventID
	}
	if true {
		toSerialize["eventCreatedDateTime"] = o.EventCreatedDateTime
	}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["eventClassifierCode"] = o.EventClassifierCode
	}
	if true {
		toSerialize["eventDateTime"] = o.EventDateTime
	}
	if o.CarrierSpecificData != nil {
		toSerialize["carrierSpecificData"] = o.CarrierSpecificData
	}
	return json.Marshal(toSerialize)
}

type NullableBaseEvent struct {
	value *BaseEvent
	isSet bool
}

func (v NullableBaseEvent) Get() *BaseEvent {
	return v.value
}

func (v *NullableBaseEvent) Set(val *BaseEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEvent(val *BaseEvent) *NullableBaseEvent {
	return &NullableBaseEvent{value: val, isSet: true}
}

func (v NullableBaseEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
