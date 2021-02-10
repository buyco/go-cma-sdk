/*
 * logistic.tracking.v1
 *
 * Retrieve tracking data about your cargo using container number, booking number or bill of lading.
 *
 * API version: 1.4.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cma

import (
	"encoding/json"
)

// Route struct for Route
type Route struct {
	JourneyLegs *[]JourneyLeg `json:"journeyLegs,omitempty"`
	Containers *[]Container `json:"containers,omitempty"`
}

// NewRoute instantiates a new Route object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoute() *Route {
	this := Route{}
	return &this
}

// NewRouteWithDefaults instantiates a new Route object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteWithDefaults() *Route {
	this := Route{}
	return &this
}

// GetJourneyLegs returns the JourneyLegs field value if set, zero value otherwise.
func (o *Route) GetJourneyLegs() []JourneyLeg {
	if o == nil || o.JourneyLegs == nil {
		var ret []JourneyLeg
		return ret
	}
	return *o.JourneyLegs
}

// GetJourneyLegsOk returns a tuple with the JourneyLegs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetJourneyLegsOk() (*[]JourneyLeg, bool) {
	if o == nil || o.JourneyLegs == nil {
		return nil, false
	}
	return o.JourneyLegs, true
}

// HasJourneyLegs returns a boolean if a field has been set.
func (o *Route) HasJourneyLegs() bool {
	if o != nil && o.JourneyLegs != nil {
		return true
	}

	return false
}

// SetJourneyLegs gets a reference to the given []JourneyLeg and assigns it to the JourneyLegs field.
func (o *Route) SetJourneyLegs(v []JourneyLeg) {
	o.JourneyLegs = &v
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *Route) GetContainers() []Container {
	if o == nil || o.Containers == nil {
		var ret []Container
		return ret
	}
	return *o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Route) GetContainersOk() (*[]Container, bool) {
	if o == nil || o.Containers == nil {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *Route) HasContainers() bool {
	if o != nil && o.Containers != nil {
		return true
	}

	return false
}

// SetContainers gets a reference to the given []Container and assigns it to the Containers field.
func (o *Route) SetContainers(v []Container) {
	o.Containers = &v
}

func (o Route) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JourneyLegs != nil {
		toSerialize["journeyLegs"] = o.JourneyLegs
	}
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}
	return json.Marshal(toSerialize)
}

type NullableRoute struct {
	value *Route
	isSet bool
}

func (v NullableRoute) Get() *Route {
	return v.value
}

func (v *NullableRoute) Set(val *Route) {
	v.value = val
	v.isSet = true
}

func (v NullableRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoute(val *Route) *NullableRoute {
	return &NullableRoute{value: val, isSet: true}
}

func (v NullableRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

