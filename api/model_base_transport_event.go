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

// BaseTransportEvent The transport event entity is a specialization of the event entity to support specification of data that only applies to a transport event.
type BaseTransportEvent struct {
	EventType *string `json:"eventType,omitempty"`
	// Code for the event classifier can be - ACT (Actual) - PLN (Planned) - EST (Estimated)
	EventClassifierCode    *string                `json:"eventClassifierCode,omitempty"`
	TransportEventTypeCode TransportEventTypeCode `json:"transportEventTypeCode"`
	// Reason code for the delay. The SMDG-Delay-Reason-Codes are used for this attribute. The code list can be found at http://www.smdg.org/smdg-code-lists/
	DelayReasonCode *string `json:"delayReasonCode,omitempty"`
	// Free text information provided by the vessel operator regarding the reasons for the change in schedule and/or plans to mitigate schedule slippage.
	ChangeRemark  *string       `json:"changeRemark,omitempty"`
	TransportCall TransportCall `json:"transportCall"`
}

// NewBaseTransportEvent instantiates a new BaseTransportEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseTransportEvent(transportEventTypeCode TransportEventTypeCode, transportCall TransportCall) *BaseTransportEvent {
	this := BaseTransportEvent{}
	this.TransportEventTypeCode = transportEventTypeCode
	this.TransportCall = transportCall
	return &this
}

// NewBaseTransportEventWithDefaults instantiates a new BaseTransportEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseTransportEventWithDefaults() *BaseTransportEvent {
	this := BaseTransportEvent{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *BaseTransportEvent) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTransportEvent) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *BaseTransportEvent) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *BaseTransportEvent) SetEventType(v string) {
	o.EventType = &v
}

// GetEventClassifierCode returns the EventClassifierCode field value if set, zero value otherwise.
func (o *BaseTransportEvent) GetEventClassifierCode() string {
	if o == nil || o.EventClassifierCode == nil {
		var ret string
		return ret
	}
	return *o.EventClassifierCode
}

// GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTransportEvent) GetEventClassifierCodeOk() (*string, bool) {
	if o == nil || o.EventClassifierCode == nil {
		return nil, false
	}
	return o.EventClassifierCode, true
}

// HasEventClassifierCode returns a boolean if a field has been set.
func (o *BaseTransportEvent) HasEventClassifierCode() bool {
	if o != nil && o.EventClassifierCode != nil {
		return true
	}

	return false
}

// SetEventClassifierCode gets a reference to the given string and assigns it to the EventClassifierCode field.
func (o *BaseTransportEvent) SetEventClassifierCode(v string) {
	o.EventClassifierCode = &v
}

// GetTransportEventTypeCode returns the TransportEventTypeCode field value
func (o *BaseTransportEvent) GetTransportEventTypeCode() TransportEventTypeCode {
	if o == nil {
		var ret TransportEventTypeCode
		return ret
	}

	return o.TransportEventTypeCode
}

// GetTransportEventTypeCodeOk returns a tuple with the TransportEventTypeCode field value
// and a boolean to check if the value has been set.
func (o *BaseTransportEvent) GetTransportEventTypeCodeOk() (*TransportEventTypeCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransportEventTypeCode, true
}

// SetTransportEventTypeCode sets field value
func (o *BaseTransportEvent) SetTransportEventTypeCode(v TransportEventTypeCode) {
	o.TransportEventTypeCode = v
}

// GetDelayReasonCode returns the DelayReasonCode field value if set, zero value otherwise.
func (o *BaseTransportEvent) GetDelayReasonCode() string {
	if o == nil || o.DelayReasonCode == nil {
		var ret string
		return ret
	}
	return *o.DelayReasonCode
}

// GetDelayReasonCodeOk returns a tuple with the DelayReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTransportEvent) GetDelayReasonCodeOk() (*string, bool) {
	if o == nil || o.DelayReasonCode == nil {
		return nil, false
	}
	return o.DelayReasonCode, true
}

// HasDelayReasonCode returns a boolean if a field has been set.
func (o *BaseTransportEvent) HasDelayReasonCode() bool {
	if o != nil && o.DelayReasonCode != nil {
		return true
	}

	return false
}

// SetDelayReasonCode gets a reference to the given string and assigns it to the DelayReasonCode field.
func (o *BaseTransportEvent) SetDelayReasonCode(v string) {
	o.DelayReasonCode = &v
}

// GetChangeRemark returns the ChangeRemark field value if set, zero value otherwise.
func (o *BaseTransportEvent) GetChangeRemark() string {
	if o == nil || o.ChangeRemark == nil {
		var ret string
		return ret
	}
	return *o.ChangeRemark
}

// GetChangeRemarkOk returns a tuple with the ChangeRemark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTransportEvent) GetChangeRemarkOk() (*string, bool) {
	if o == nil || o.ChangeRemark == nil {
		return nil, false
	}
	return o.ChangeRemark, true
}

// HasChangeRemark returns a boolean if a field has been set.
func (o *BaseTransportEvent) HasChangeRemark() bool {
	if o != nil && o.ChangeRemark != nil {
		return true
	}

	return false
}

// SetChangeRemark gets a reference to the given string and assigns it to the ChangeRemark field.
func (o *BaseTransportEvent) SetChangeRemark(v string) {
	o.ChangeRemark = &v
}

// GetTransportCall returns the TransportCall field value
func (o *BaseTransportEvent) GetTransportCall() TransportCall {
	if o == nil {
		var ret TransportCall
		return ret
	}

	return o.TransportCall
}

// GetTransportCallOk returns a tuple with the TransportCall field value
// and a boolean to check if the value has been set.
func (o *BaseTransportEvent) GetTransportCallOk() (*TransportCall, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransportCall, true
}

// SetTransportCall sets field value
func (o *BaseTransportEvent) SetTransportCall(v TransportCall) {
	o.TransportCall = v
}

func (o BaseTransportEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.EventClassifierCode != nil {
		toSerialize["eventClassifierCode"] = o.EventClassifierCode
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

type NullableBaseTransportEvent struct {
	value *BaseTransportEvent
	isSet bool
}

func (v NullableBaseTransportEvent) Get() *BaseTransportEvent {
	return v.value
}

func (v *NullableBaseTransportEvent) Set(val *BaseTransportEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseTransportEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseTransportEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseTransportEvent(val *BaseTransportEvent) *NullableBaseTransportEvent {
	return &NullableBaseTransportEvent{value: val, isSet: true}
}

func (v NullableBaseTransportEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseTransportEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
