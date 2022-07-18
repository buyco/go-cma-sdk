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

// ShipmentEvent struct for ShipmentEvent
type ShipmentEvent struct {
	// The unique identifier for the event (the message - not the source). <b>NB</b>&#58; This field should be considered Metadata
	EventID *string `json:"eventID,omitempty"`
	// The timestamp of when the event was created. <b>NB</b>&#58; This field should be considered Metadata
	EventCreatedDateTime time.Time `json:"eventCreatedDateTime"`
	EventType            string    `json:"eventType"`
	// Code for the event classifier can be - ACT (Actual) - PLN (Planned) - EST (Estimated)
	EventClassifierCode string `json:"eventClassifierCode"`
	// Value for eventDateTime must be the same value as eventCreatedDateTime
	EventDateTime         interface{}           `json:"eventDateTime"`
	CarrierSpecificData   *CarrierSpecificData  `json:"carrierSpecificData,omitempty"`
	ShipmentEventTypeCode ShipmentEventTypeCode `json:"shipmentEventTypeCode"`
	// The id of the object defined by the documentTypeCode.
	DocumentID       string           `json:"documentID"`
	DocumentTypeCode DocumentTypeCode `json:"documentTypeCode"`
	// Reason field in a Shipment event. This field can be used to explain why a specific event has been sent.
	Reason *string `json:"reason,omitempty"`
	// The identifier for a shipment
	ShipmentID *string     `json:"shipmentID,omitempty"`
	References []Reference `json:"references,omitempty"`
}

// NewShipmentEvent instantiates a new ShipmentEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentEvent(eventCreatedDateTime time.Time, eventType string, eventClassifierCode string, eventDateTime interface{}, shipmentEventTypeCode ShipmentEventTypeCode, documentID string, documentTypeCode DocumentTypeCode) *ShipmentEvent {
	this := ShipmentEvent{}
	this.EventCreatedDateTime = eventCreatedDateTime
	this.EventType = eventType
	this.EventClassifierCode = eventClassifierCode
	this.EventDateTime = eventDateTime
	this.ShipmentEventTypeCode = shipmentEventTypeCode
	this.DocumentID = documentID
	this.DocumentTypeCode = documentTypeCode
	return &this
}

// NewShipmentEventWithDefaults instantiates a new ShipmentEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentEventWithDefaults() *ShipmentEvent {
	this := ShipmentEvent{}
	return &this
}

// GetEventID returns the EventID field value if set, zero value otherwise.
func (o *ShipmentEvent) GetEventID() string {
	if o == nil || o.EventID == nil {
		var ret string
		return ret
	}
	return *o.EventID
}

// GetEventIDOk returns a tuple with the EventID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetEventIDOk() (*string, bool) {
	if o == nil || o.EventID == nil {
		return nil, false
	}
	return o.EventID, true
}

// HasEventID returns a boolean if a field has been set.
func (o *ShipmentEvent) HasEventID() bool {
	if o != nil && o.EventID != nil {
		return true
	}

	return false
}

// SetEventID gets a reference to the given string and assigns it to the EventID field.
func (o *ShipmentEvent) SetEventID(v string) {
	o.EventID = &v
}

// GetEventCreatedDateTime returns the EventCreatedDateTime field value
func (o *ShipmentEvent) GetEventCreatedDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventCreatedDateTime
}

// GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field value
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetEventCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCreatedDateTime, true
}

// SetEventCreatedDateTime sets field value
func (o *ShipmentEvent) SetEventCreatedDateTime(v time.Time) {
	o.EventCreatedDateTime = v
}

// GetEventType returns the EventType field value
func (o *ShipmentEvent) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *ShipmentEvent) SetEventType(v string) {
	o.EventType = v
}

// GetEventClassifierCode returns the EventClassifierCode field value
func (o *ShipmentEvent) GetEventClassifierCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventClassifierCode
}

// GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field value
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetEventClassifierCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventClassifierCode, true
}

// SetEventClassifierCode sets field value
func (o *ShipmentEvent) SetEventClassifierCode(v string) {
	o.EventClassifierCode = v
}

// GetEventDateTime returns the EventDateTime field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ShipmentEvent) GetEventDateTime() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.EventDateTime
}

// GetEventDateTimeOk returns a tuple with the EventDateTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipmentEvent) GetEventDateTimeOk() (*interface{}, bool) {
	if o == nil || o.EventDateTime == nil {
		return nil, false
	}
	return &o.EventDateTime, true
}

// SetEventDateTime sets field value
func (o *ShipmentEvent) SetEventDateTime(v interface{}) {
	o.EventDateTime = v
}

// GetCarrierSpecificData returns the CarrierSpecificData field value if set, zero value otherwise.
func (o *ShipmentEvent) GetCarrierSpecificData() CarrierSpecificData {
	if o == nil || o.CarrierSpecificData == nil {
		var ret CarrierSpecificData
		return ret
	}
	return *o.CarrierSpecificData
}

// GetCarrierSpecificDataOk returns a tuple with the CarrierSpecificData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetCarrierSpecificDataOk() (*CarrierSpecificData, bool) {
	if o == nil || o.CarrierSpecificData == nil {
		return nil, false
	}
	return o.CarrierSpecificData, true
}

// HasCarrierSpecificData returns a boolean if a field has been set.
func (o *ShipmentEvent) HasCarrierSpecificData() bool {
	if o != nil && o.CarrierSpecificData != nil {
		return true
	}

	return false
}

// SetCarrierSpecificData gets a reference to the given CarrierSpecificData and assigns it to the CarrierSpecificData field.
func (o *ShipmentEvent) SetCarrierSpecificData(v CarrierSpecificData) {
	o.CarrierSpecificData = &v
}

// GetShipmentEventTypeCode returns the ShipmentEventTypeCode field value
func (o *ShipmentEvent) GetShipmentEventTypeCode() ShipmentEventTypeCode {
	if o == nil {
		var ret ShipmentEventTypeCode
		return ret
	}

	return o.ShipmentEventTypeCode
}

// GetShipmentEventTypeCodeOk returns a tuple with the ShipmentEventTypeCode field value
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetShipmentEventTypeCodeOk() (*ShipmentEventTypeCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentEventTypeCode, true
}

// SetShipmentEventTypeCode sets field value
func (o *ShipmentEvent) SetShipmentEventTypeCode(v ShipmentEventTypeCode) {
	o.ShipmentEventTypeCode = v
}

// GetDocumentID returns the DocumentID field value
func (o *ShipmentEvent) GetDocumentID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentID
}

// GetDocumentIDOk returns a tuple with the DocumentID field value
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetDocumentIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentID, true
}

// SetDocumentID sets field value
func (o *ShipmentEvent) SetDocumentID(v string) {
	o.DocumentID = v
}

// GetDocumentTypeCode returns the DocumentTypeCode field value
func (o *ShipmentEvent) GetDocumentTypeCode() DocumentTypeCode {
	if o == nil {
		var ret DocumentTypeCode
		return ret
	}

	return o.DocumentTypeCode
}

// GetDocumentTypeCodeOk returns a tuple with the DocumentTypeCode field value
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetDocumentTypeCodeOk() (*DocumentTypeCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentTypeCode, true
}

// SetDocumentTypeCode sets field value
func (o *ShipmentEvent) SetDocumentTypeCode(v DocumentTypeCode) {
	o.DocumentTypeCode = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ShipmentEvent) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ShipmentEvent) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ShipmentEvent) SetReason(v string) {
	o.Reason = &v
}

// GetShipmentID returns the ShipmentID field value if set, zero value otherwise.
func (o *ShipmentEvent) GetShipmentID() string {
	if o == nil || o.ShipmentID == nil {
		var ret string
		return ret
	}
	return *o.ShipmentID
}

// GetShipmentIDOk returns a tuple with the ShipmentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetShipmentIDOk() (*string, bool) {
	if o == nil || o.ShipmentID == nil {
		return nil, false
	}
	return o.ShipmentID, true
}

// HasShipmentID returns a boolean if a field has been set.
func (o *ShipmentEvent) HasShipmentID() bool {
	if o != nil && o.ShipmentID != nil {
		return true
	}

	return false
}

// SetShipmentID gets a reference to the given string and assigns it to the ShipmentID field.
func (o *ShipmentEvent) SetShipmentID(v string) {
	o.ShipmentID = &v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *ShipmentEvent) GetReferences() []Reference {
	if o == nil || o.References == nil {
		var ret []Reference
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetReferencesOk() ([]Reference, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *ShipmentEvent) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []Reference and assigns it to the References field.
func (o *ShipmentEvent) SetReferences(v []Reference) {
	o.References = v
}

func (o ShipmentEvent) MarshalJSON() ([]byte, error) {
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
		toSerialize["shipmentEventTypeCode"] = o.ShipmentEventTypeCode
	}
	if true {
		toSerialize["documentID"] = o.DocumentID
	}
	if true {
		toSerialize["documentTypeCode"] = o.DocumentTypeCode
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.ShipmentID != nil {
		toSerialize["shipmentID"] = o.ShipmentID
	}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	return json.Marshal(toSerialize)
}

type NullableShipmentEvent struct {
	value *ShipmentEvent
	isSet bool
}

func (v NullableShipmentEvent) Get() *ShipmentEvent {
	return v.value
}

func (v *NullableShipmentEvent) Set(val *ShipmentEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentEvent(val *ShipmentEvent) *NullableShipmentEvent {
	return &NullableShipmentEvent{value: val, isSet: true}
}

func (v NullableShipmentEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
