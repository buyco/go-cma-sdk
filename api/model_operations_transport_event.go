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

// OperationsTransportEvent struct for OperationsTransportEvent
type OperationsTransportEvent struct {
	// The unique identifier for the event (the message - not the source). <b>NB</b>&#58; This field should be considered Metadata
	EventID *string `json:"eventID,omitempty"`
	// The timestamp of when the event was created. <b>NB</b>&#58; This field should be considered Metadata
	EventCreatedDateTime time.Time `json:"eventCreatedDateTime"`
	EventType            string    `json:"eventType"`
	// Code for the event classifier can be - ACT (Actual) - PLN (Planned) - EST (Estimated)
	EventClassifierCode string `json:"eventClassifierCode"`
	// The local date and time, where the event took place or when the event will take place, in ISO 8601 format.
	EventDateTime          *time.Time             `json:"eventDateTime,omitempty"`
	CarrierSpecificData    *CarrierSpecificData   `json:"carrierSpecificData,omitempty"`
	TransportEventTypeCode TransportEventTypeCode `json:"transportEventTypeCode"`
	// Reason code for the delay. The SMDG-Delay-Reason-Codes are used for this attribute. The code list can be found at http://www.smdg.org/smdg-code-lists/
	DelayReasonCode *string `json:"delayReasonCode,omitempty"`
	// Free text information provided by the vessel operator regarding the reasons for the change in schedule and/or plans to mitigate schedule slippage.
	ChangeRemark  *string       `json:"changeRemark,omitempty"`
	TransportCall TransportCall `json:"transportCall"`
}

// NewOperationsTransportEvent instantiates a new OperationsTransportEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationsTransportEvent(eventCreatedDateTime time.Time, eventType string, eventClassifierCode string, transportEventTypeCode TransportEventTypeCode, transportCall TransportCall) *OperationsTransportEvent {
	this := OperationsTransportEvent{}
	this.EventCreatedDateTime = eventCreatedDateTime
	this.EventType = eventType
	this.EventClassifierCode = eventClassifierCode
	this.TransportEventTypeCode = transportEventTypeCode
	this.TransportCall = transportCall
	return &this
}

// NewOperationsTransportEventWithDefaults instantiates a new OperationsTransportEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationsTransportEventWithDefaults() *OperationsTransportEvent {
	this := OperationsTransportEvent{}
	return &this
}

// GetEventID returns the EventID field value if set, zero value otherwise.
func (o *OperationsTransportEvent) GetEventID() string {
	if o == nil || o.EventID == nil {
		var ret string
		return ret
	}
	return *o.EventID
}

// GetEventIDOk returns a tuple with the EventID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetEventIDOk() (*string, bool) {
	if o == nil || o.EventID == nil {
		return nil, false
	}
	return o.EventID, true
}

// HasEventID returns a boolean if a field has been set.
func (o *OperationsTransportEvent) HasEventID() bool {
	if o != nil && o.EventID != nil {
		return true
	}

	return false
}

// SetEventID gets a reference to the given string and assigns it to the EventID field.
func (o *OperationsTransportEvent) SetEventID(v string) {
	o.EventID = &v
}

// GetEventCreatedDateTime returns the EventCreatedDateTime field value
func (o *OperationsTransportEvent) GetEventCreatedDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventCreatedDateTime
}

// GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field value
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetEventCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCreatedDateTime, true
}

// SetEventCreatedDateTime sets field value
func (o *OperationsTransportEvent) SetEventCreatedDateTime(v time.Time) {
	o.EventCreatedDateTime = v
}

// GetEventType returns the EventType field value
func (o *OperationsTransportEvent) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *OperationsTransportEvent) SetEventType(v string) {
	o.EventType = v
}

// GetEventClassifierCode returns the EventClassifierCode field value
func (o *OperationsTransportEvent) GetEventClassifierCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventClassifierCode
}

// GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field value
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetEventClassifierCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventClassifierCode, true
}

// SetEventClassifierCode sets field value
func (o *OperationsTransportEvent) SetEventClassifierCode(v string) {
	o.EventClassifierCode = v
}

// GetEventDateTime returns the EventDateTime field value if set, zero value otherwise.
func (o *OperationsTransportEvent) GetEventDateTime() time.Time {
	if o == nil || o.EventDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EventDateTime
}

// GetEventDateTimeOk returns a tuple with the EventDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetEventDateTimeOk() (*time.Time, bool) {
	if o == nil || o.EventDateTime == nil {
		return nil, false
	}
	return o.EventDateTime, true
}

// HasEventDateTime returns a boolean if a field has been set.
func (o *OperationsTransportEvent) HasEventDateTime() bool {
	if o != nil && o.EventDateTime != nil {
		return true
	}

	return false
}

// SetEventDateTime gets a reference to the given time.Time and assigns it to the EventDateTime field.
func (o *OperationsTransportEvent) SetEventDateTime(v time.Time) {
	o.EventDateTime = &v
}

// GetCarrierSpecificData returns the CarrierSpecificData field value if set, zero value otherwise.
func (o *OperationsTransportEvent) GetCarrierSpecificData() CarrierSpecificData {
	if o == nil || o.CarrierSpecificData == nil {
		var ret CarrierSpecificData
		return ret
	}
	return *o.CarrierSpecificData
}

// GetCarrierSpecificDataOk returns a tuple with the CarrierSpecificData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetCarrierSpecificDataOk() (*CarrierSpecificData, bool) {
	if o == nil || o.CarrierSpecificData == nil {
		return nil, false
	}
	return o.CarrierSpecificData, true
}

// HasCarrierSpecificData returns a boolean if a field has been set.
func (o *OperationsTransportEvent) HasCarrierSpecificData() bool {
	if o != nil && o.CarrierSpecificData != nil {
		return true
	}

	return false
}

// SetCarrierSpecificData gets a reference to the given CarrierSpecificData and assigns it to the CarrierSpecificData field.
func (o *OperationsTransportEvent) SetCarrierSpecificData(v CarrierSpecificData) {
	o.CarrierSpecificData = &v
}

// GetTransportEventTypeCode returns the TransportEventTypeCode field value
func (o *OperationsTransportEvent) GetTransportEventTypeCode() TransportEventTypeCode {
	if o == nil {
		var ret TransportEventTypeCode
		return ret
	}

	return o.TransportEventTypeCode
}

// GetTransportEventTypeCodeOk returns a tuple with the TransportEventTypeCode field value
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetTransportEventTypeCodeOk() (*TransportEventTypeCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransportEventTypeCode, true
}

// SetTransportEventTypeCode sets field value
func (o *OperationsTransportEvent) SetTransportEventTypeCode(v TransportEventTypeCode) {
	o.TransportEventTypeCode = v
}

// GetDelayReasonCode returns the DelayReasonCode field value if set, zero value otherwise.
func (o *OperationsTransportEvent) GetDelayReasonCode() string {
	if o == nil || o.DelayReasonCode == nil {
		var ret string
		return ret
	}
	return *o.DelayReasonCode
}

// GetDelayReasonCodeOk returns a tuple with the DelayReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetDelayReasonCodeOk() (*string, bool) {
	if o == nil || o.DelayReasonCode == nil {
		return nil, false
	}
	return o.DelayReasonCode, true
}

// HasDelayReasonCode returns a boolean if a field has been set.
func (o *OperationsTransportEvent) HasDelayReasonCode() bool {
	if o != nil && o.DelayReasonCode != nil {
		return true
	}

	return false
}

// SetDelayReasonCode gets a reference to the given string and assigns it to the DelayReasonCode field.
func (o *OperationsTransportEvent) SetDelayReasonCode(v string) {
	o.DelayReasonCode = &v
}

// GetChangeRemark returns the ChangeRemark field value if set, zero value otherwise.
func (o *OperationsTransportEvent) GetChangeRemark() string {
	if o == nil || o.ChangeRemark == nil {
		var ret string
		return ret
	}
	return *o.ChangeRemark
}

// GetChangeRemarkOk returns a tuple with the ChangeRemark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetChangeRemarkOk() (*string, bool) {
	if o == nil || o.ChangeRemark == nil {
		return nil, false
	}
	return o.ChangeRemark, true
}

// HasChangeRemark returns a boolean if a field has been set.
func (o *OperationsTransportEvent) HasChangeRemark() bool {
	if o != nil && o.ChangeRemark != nil {
		return true
	}

	return false
}

// SetChangeRemark gets a reference to the given string and assigns it to the ChangeRemark field.
func (o *OperationsTransportEvent) SetChangeRemark(v string) {
	o.ChangeRemark = &v
}

// GetTransportCall returns the TransportCall field value
func (o *OperationsTransportEvent) GetTransportCall() TransportCall {
	if o == nil {
		var ret TransportCall
		return ret
	}

	return o.TransportCall
}

// GetTransportCallOk returns a tuple with the TransportCall field value
// and a boolean to check if the value has been set.
func (o *OperationsTransportEvent) GetTransportCallOk() (*TransportCall, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransportCall, true
}

// SetTransportCall sets field value
func (o *OperationsTransportEvent) SetTransportCall(v TransportCall) {
	o.TransportCall = v
}

func (o OperationsTransportEvent) MarshalJSON() ([]byte, error) {
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
	if o.EventDateTime != nil {
		toSerialize["eventDateTime"] = o.EventDateTime
	}
	if o.CarrierSpecificData != nil {
		toSerialize["carrierSpecificData"] = o.CarrierSpecificData
	}
	if true {
		toSerialize["transportEventTypeCode"] = o.TransportEventTypeCode
	}
	if o.DelayReasonCode != nil {
		toSerialize["delayReasonCode"] = o.DelayReasonCode
	}
	if o.ChangeRemark != nil {
		toSerialize["changeRemark"] = o.ChangeRemark
	}
	if true {
		toSerialize["transportCall"] = o.TransportCall
	}
	return json.Marshal(toSerialize)
}

type NullableOperationsTransportEvent struct {
	value *OperationsTransportEvent
	isSet bool
}

func (v NullableOperationsTransportEvent) Get() *OperationsTransportEvent {
	return v.value
}

func (v *NullableOperationsTransportEvent) Set(val *OperationsTransportEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationsTransportEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationsTransportEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationsTransportEvent(val *OperationsTransportEvent) *NullableOperationsTransportEvent {
	return &NullableOperationsTransportEvent{value: val, isSet: true}
}

func (v NullableOperationsTransportEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationsTransportEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
