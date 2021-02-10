# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JourneyLegs** | Pointer to [**[]JourneyLeg**](JourneyLeg.md) |  | [optional] 
**Containers** | Pointer to [**[]Container**](Container.md) |  | [optional] 

## Methods

### NewRoute

`func NewRoute() *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJourneyLegs

`func (o *Route) GetJourneyLegs() []JourneyLeg`

GetJourneyLegs returns the JourneyLegs field if non-nil, zero value otherwise.

### GetJourneyLegsOk

`func (o *Route) GetJourneyLegsOk() (*[]JourneyLeg, bool)`

GetJourneyLegsOk returns a tuple with the JourneyLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJourneyLegs

`func (o *Route) SetJourneyLegs(v []JourneyLeg)`

SetJourneyLegs sets JourneyLegs field to given value.

### HasJourneyLegs

`func (o *Route) HasJourneyLegs() bool`

HasJourneyLegs returns a boolean if a field has been set.

### GetContainers

`func (o *Route) GetContainers() []Container`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *Route) GetContainersOk() (*[]Container, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *Route) SetContainers(v []Container)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *Route) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


