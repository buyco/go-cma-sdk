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

// BaseEquipmentEvent The equipment event entity is a specialization of the event entity to support specification of data that only applies to an equipment event.
type BaseEquipmentEvent struct {
	EventType *string `json:"eventType,omitempty"`
	// Code for the event classifier can be - PLN (Planned) - ACT (Actual) - EST (Estimated)
	EventClassifierCode    *string                `json:"eventClassifierCode,omitempty"`
	EquipmentEventTypeCode EquipmentEventTypeCode `json:"equipmentEventTypeCode"`
	// The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation
	EquipmentReference *string `json:"equipmentReference,omitempty"`
	// Unique code for the different equipment size/type used for transporting commodities. The code is a concatenation of ISO Equipment Size Code and ISO Equipment Type Code A and follows the ISO 6346 standard.
	ISOEquipmentCode   *string            `json:"ISOEquipmentCode,omitempty"`
	EmptyIndicatorCode EmptyIndicatorCode `json:"emptyIndicatorCode"`
	EventLocation      *Location          `json:"eventLocation,omitempty"`
	TransportCall      *TransportCall     `json:"transportCall,omitempty"`
	// An optional list of key-value (documentReferenceType-documentReferenceValue) pairs representing links to objects relevant to the event. The <b>documentReferenceType</b>-field is used to describe where the <b>documentReferenceValue</b>-field is pointing to.
	DocumentReferences []DocumentReferencesInner `json:"documentReferences,omitempty"`
	References         []Reference               `json:"references,omitempty"`
	Seals              []Seal                    `json:"seals,omitempty"`
}

// NewBaseEquipmentEvent instantiates a new BaseEquipmentEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEquipmentEvent(equipmentEventTypeCode EquipmentEventTypeCode, emptyIndicatorCode EmptyIndicatorCode) *BaseEquipmentEvent {
	this := BaseEquipmentEvent{}
	this.EquipmentEventTypeCode = equipmentEventTypeCode
	this.EmptyIndicatorCode = emptyIndicatorCode
	return &this
}

// NewBaseEquipmentEventWithDefaults instantiates a new BaseEquipmentEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEquipmentEventWithDefaults() *BaseEquipmentEvent {
	this := BaseEquipmentEvent{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *BaseEquipmentEvent) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEquipmentEvent) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *BaseEquipmentEvent) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *BaseEquipmentEvent) SetEventType(v string) {
	o.EventType = &v
}

// GetEventClassifierCode returns the EventClassifierCode field value if set, zero value otherwise.
func (o *BaseEquipmentEvent) GetEventClassifierCode() string {
	if o == nil || o.EventClassifierCode == nil {
		var ret string
		return ret
	}
	return *o.EventClassifierCode
}

// GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEquipmentEvent) GetEventClassifierCodeOk() (*string, bool) {
	if o == nil || o.EventClassifierCode == nil {
		return nil, false
	}
	return o.EventClassifierCode, true
}

// HasEventClassifierCode returns a boolean if a field has been set.
func (o *BaseEquipmentEvent) HasEventClassifierCode() bool {
	if o != nil && o.EventClassifierCode != nil {
		return true
	}

	return false
}

// SetEventClassifierCode gets a reference to the given string and assigns it to the EventClassifierCode field.
func (o *BaseEquipmentEvent) SetEventClassifierCode(v string) {
	o.EventClassifierCode = &v
}

// GetEquipmentEventTypeCode returns the EquipmentEventTypeCode field value
func (o *BaseEquipmentEvent) GetEquipmentEventTypeCode() EquipmentEventTypeCode {
	if o == nil {
		var ret EquipmentEventTypeCode
		return ret
	}

	return o.EquipmentEventTypeCode
}

// GetEquipmentEventTypeCodeOk returns a tuple with the EquipmentEventTypeCode field value
// and a boolean to check if the value has been set.
func (o *BaseEquipmentEvent) GetEquipmentEventTypeCodeOk() (*EquipmentEventTypeCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EquipmentEventTypeCode, true
}

// SetEquipmentEventTypeCode sets field value
func (o *BaseEquipmentEvent) SetEquipmentEventTypeCode(v EquipmentEventTypeCode) {
	o.EquipmentEventTypeCode = v
}

// GetEquipmentReference returns the EquipmentReference field value if set, zero value otherwise.
func (o *BaseEquipmentEvent) GetEquipmentReference() string {
	if o == nil || o.EquipmentReference == nil {
		var ret string
		return ret
	}
	return *o.EquipmentReference
}

// GetEquipmentReferenceOk returns a tuple with the EquipmentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEquipmentEvent) GetEquipmentReferenceOk() (*string, bool) {
	if o == nil || o.EquipmentReference == nil {
		return nil, false
	}
	return o.EquipmentReference, true
}

// HasEquipmentReference returns a boolean if a field has been set.
func (o *BaseEquipmentEvent) HasEquipmentReference() bool {
	if o != nil && o.EquipmentReference != nil {
		return true
	}

	return false
}

// SetEquipmentReference gets a reference to the given string and assigns it to the EquipmentReference field.
func (o *BaseEquipmentEvent) SetEquipmentReference(v string) {
	o.EquipmentReference = &v
}

// GetISOEquipmentCode returns the ISOEquipmentCode field value if set, zero value otherwise.
func (o *BaseEquipmentEvent) GetISOEquipmentCode() string {
	if o == nil || o.ISOEquipmentCode == nil {
		var ret string
		return ret
	}
	return *o.ISOEquipmentCode
}

// GetISOEquipmentCodeOk returns a tuple with the ISOEquipmentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEquipmentEvent) GetISOEquipmentCodeOk() (*string, bool) {
	if o == nil || o.ISOEquipmentCode == nil {
		return nil, false
	}
	return o.ISOEquipmentCode, true
}

// HasISOEquipmentCode returns a boolean if a field has been set.
func (o *BaseEquipmentEvent) HasISOEquipmentCode() bool {
	if o != nil && o.ISOEquipmentCode != nil {
		return true
	}

	return false
}

// SetISOEquipmentCode gets a reference to the given string and assigns it to the ISOEquipmentCode field.
func (o *BaseEquipmentEvent) SetISOEquipmentCode(v string) {
	o.ISOEquipmentCode = &v
}

// GetEmptyIndicatorCode returns the EmptyIndicatorCode field value
func (o *BaseEquipmentEvent) GetEmptyIndicatorCode() EmptyIndicatorCode {
	if o == nil {
		var ret EmptyIndicatorCode
		return ret
	}

	return o.EmptyIndicatorCode
}

// GetEmptyIndicatorCodeOk returns a tuple with the EmptyIndicatorCode field value
// and a boolean to check if the value has been set.
func (o *BaseEquipmentEvent) GetEmptyIndicatorCodeOk() (*EmptyIndicatorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmptyIndicatorCode, true
}

// SetEmptyIndicatorCode sets field value
func (o *BaseEquipmentEvent) SetEmptyIndicatorCode(v EmptyIndicatorCode) {
	o.EmptyIndicatorCode = v
}

// GetEventLocation returns the EventLocation field value if set, zero value otherwise.
func (o *BaseEquipmentEvent) GetEventLocation() Location {
	if o == nil || o.EventLocation == nil {
		var ret Location
		return ret
	}
	return *o.EventLocation
}

// GetEventLocationOk returns a tuple with the EventLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEquipmentEvent) GetEventLocationOk() (*Location, bool) {
	if o == nil || o.EventLocation == nil {
		return nil, false
	}
	return o.EventLocation, true
}

// HasEventLocation returns a boolean if a field has been set.
func (o *BaseEquipmentEvent) HasEventLocation() bool {
	if o != nil && o.EventLocation != nil {
		return true
	}

	return false
}

// SetEventLocation gets a reference to the given Location and assigns it to the EventLocation field.
func (o *BaseEquipmentEvent) SetEventLocation(v Location) {
	o.EventLocation = &v
}

// GetTransportCall returns the TransportCall field value if set, zero value otherwise.
func (o *BaseEquipmentEvent) GetTransportCall() TransportCall {
	if o == nil || o.TransportCall == nil {
		var ret TransportCall
		return ret
	}
	return *o.TransportCall
}

// GetTransportCallOk returns a tuple with the TransportCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEquipmentEvent) GetTransportCallOk() (*TransportCall, bool) {
	if o == nil || o.TransportCall == nil {
		return nil, false
	}
	return o.TransportCall, true
}

// HasTransportCall returns a boolean if a field has been set.
func (o *BaseEquipmentEvent) HasTransportCall() bool {
	if o != nil && o.TransportCall != nil {
		return true
	}

	return false
}

// SetTransportCall gets a reference to the given TransportCall and assigns it to the TransportCall field.
func (o *BaseEquipmentEvent) SetTransportCall(v TransportCall) {
	o.TransportCall = &v
}

// GetDocumentReferences returns the DocumentReferences field value if set, zero value otherwise.
func (o *BaseEquipmentEvent) GetDocumentReferences() []DocumentReferencesInner {
	if o == nil || o.DocumentReferences == nil {
		var ret []DocumentReferencesInner
		return ret
	}
	return o.DocumentReferences
}

// GetDocumentReferencesOk returns a tuple with the DocumentReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEquipmentEvent) GetDocumentReferencesOk() ([]DocumentReferencesInner, bool) {
	if o == nil || o.DocumentReferences == nil {
		return nil, false
	}
	return o.DocumentReferences, true
}

// HasDocumentReferences returns a boolean if a field has been set.
func (o *BaseEquipmentEvent) HasDocumentReferences() bool {
	if o != nil && o.DocumentReferences != nil {
		return true
	}

	return false
}

// SetDocumentReferences gets a reference to the given []DocumentReferencesInner and assigns it to the DocumentReferences field.
func (o *BaseEquipmentEvent) SetDocumentReferences(v []DocumentReferencesInner) {
	o.DocumentReferences = v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *BaseEquipmentEvent) GetReferences() []Reference {
	if o == nil || o.References == nil {
		var ret []Reference
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEquipmentEvent) GetReferencesOk() ([]Reference, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *BaseEquipmentEvent) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []Reference and assigns it to the References field.
func (o *BaseEquipmentEvent) SetReferences(v []Reference) {
	o.References = v
}

// GetSeals returns the Seals field value if set, zero value otherwise.
func (o *BaseEquipmentEvent) GetSeals() []Seal {
	if o == nil || o.Seals == nil {
		var ret []Seal
		return ret
	}
	return o.Seals
}

// GetSealsOk returns a tuple with the Seals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEquipmentEvent) GetSealsOk() ([]Seal, bool) {
	if o == nil || o.Seals == nil {
		return nil, false
	}
	return o.Seals, true
}

// HasSeals returns a boolean if a field has been set.
func (o *BaseEquipmentEvent) HasSeals() bool {
	if o != nil && o.Seals != nil {
		return true
	}

	return false
}

// SetSeals gets a reference to the given []Seal and assigns it to the Seals field.
func (o *BaseEquipmentEvent) SetSeals(v []Seal) {
	o.Seals = v
}

func (o BaseEquipmentEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.EventClassifierCode != nil {
		toSerialize["eventClassifierCode"] = o.EventClassifierCode
	}
	if true {
		toSerialize["equipmentEventTypeCode"] = o.EquipmentEventTypeCode
	}
	if o.EquipmentReference != nil {
		toSerialize["equipmentReference"] = o.EquipmentReference
	}
	if o.ISOEquipmentCode != nil {
		toSerialize["ISOEquipmentCode"] = o.ISOEquipmentCode
	}
	if true {
		toSerialize["emptyIndicatorCode"] = o.EmptyIndicatorCode
	}
	if o.EventLocation != nil {
		toSerialize["eventLocation"] = o.EventLocation
	}
	if o.TransportCall != nil {
		toSerialize["transportCall"] = o.TransportCall
	}
	if o.DocumentReferences != nil {
		toSerialize["documentReferences"] = o.DocumentReferences
	}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	if o.Seals != nil {
		toSerialize["seals"] = o.Seals
	}
	return json.Marshal(toSerialize)
}

type NullableBaseEquipmentEvent struct {
	value *BaseEquipmentEvent
	isSet bool
}

func (v NullableBaseEquipmentEvent) Get() *BaseEquipmentEvent {
	return v.value
}

func (v *NullableBaseEquipmentEvent) Set(val *BaseEquipmentEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEquipmentEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEquipmentEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEquipmentEvent(val *BaseEquipmentEvent) *NullableBaseEquipmentEvent {
	return &NullableBaseEquipmentEvent{value: val, isSet: true}
}

func (v NullableBaseEquipmentEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEquipmentEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
