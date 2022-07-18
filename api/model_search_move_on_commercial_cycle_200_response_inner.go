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
	"fmt"
)

// SearchMoveOnCommercialCycle200ResponseInner struct for SearchMoveOnCommercialCycle200ResponseInner
type SearchMoveOnCommercialCycle200ResponseInner struct {
	EquipmentEvent *EquipmentEvent
	ShipmentEvent *ShipmentEvent
	TransportEvent *TransportEvent
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchMoveOnCommercialCycle200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EquipmentEvent
	err = json.Unmarshal(data, &dst.EquipmentEvent);
	if err == nil && dst.EquipmentEvent.GetEventType() == "EQUIPMENT" {
		jsonEquipmentEvent, _ := json.Marshal(dst.EquipmentEvent)
		if string(jsonEquipmentEvent) == "{}" { // empty struct
			dst.EquipmentEvent = nil
		} else {
			return nil // data stored in dst.EquipmentEvent, return on the first match
		}
	} else {
		dst.EquipmentEvent = nil
	}

	// try to unmarshal JSON data into ShipmentEvent
	err = json.Unmarshal(data, &dst.ShipmentEvent);
	if err == nil && dst.ShipmentEvent.GetEventType() == "SHIPMENT" {
		jsonShipmentEvent, _ := json.Marshal(dst.ShipmentEvent)
		if string(jsonShipmentEvent) == "{}" { // empty struct
			dst.ShipmentEvent = nil
		} else {
			return nil // data stored in dst.ShipmentEvent, return on the first match
		}
	} else {
		dst.ShipmentEvent = nil
	}

	// try to unmarshal JSON data into TransportEvent
	err = json.Unmarshal(data, &dst.TransportEvent);
	if err == nil && dst.TransportEvent.GetEventType() == "TRANSPORT" {
		jsonTransportEvent, _ := json.Marshal(dst.TransportEvent)
		if string(jsonTransportEvent) == "{}" { // empty struct
			dst.TransportEvent = nil
		} else {
			return nil // data stored in dst.TransportEvent, return on the first match
		}
	} else {
		dst.TransportEvent = nil
	}
	return fmt.Errorf("Data failed to match schemas in anyOf(SearchMoveOnCommercialCycle200ResponseInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchMoveOnCommercialCycle200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.EquipmentEvent != nil {
		return json.Marshal(&src.EquipmentEvent)
	}

	if src.ShipmentEvent != nil {
		return json.Marshal(&src.ShipmentEvent)
	}

	if src.TransportEvent != nil {
		return json.Marshal(&src.TransportEvent)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchMoveOnCommercialCycle200ResponseInner struct {
	value *SearchMoveOnCommercialCycle200ResponseInner
	isSet bool
}

func (v NullableSearchMoveOnCommercialCycle200ResponseInner) Get() *SearchMoveOnCommercialCycle200ResponseInner {
	return v.value
}

func (v *NullableSearchMoveOnCommercialCycle200ResponseInner) Set(val *SearchMoveOnCommercialCycle200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchMoveOnCommercialCycle200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchMoveOnCommercialCycle200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchMoveOnCommercialCycle200ResponseInner(val *SearchMoveOnCommercialCycle200ResponseInner) *NullableSearchMoveOnCommercialCycle200ResponseInner {
	return &NullableSearchMoveOnCommercialCycle200ResponseInner{value: val, isSet: true}
}

func (v NullableSearchMoveOnCommercialCycle200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchMoveOnCommercialCycle200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

