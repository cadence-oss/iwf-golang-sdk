/*
Workflow APIs

This APIs for iwf SDKs to operate workflows

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iwfidl

import (
	"encoding/json"
)

// checks if the InterStateChannelCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterStateChannelCommand{}

// InterStateChannelCommand struct for InterStateChannelCommand
type InterStateChannelCommand struct {
	CommandId   *string `json:"commandId,omitempty"`
	ChannelName string  `json:"channelName"`
	AtLeast     *int32  `json:"atLeast,omitempty"`
	AtMost      *int32  `json:"atMost,omitempty"`
}

// NewInterStateChannelCommand instantiates a new InterStateChannelCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterStateChannelCommand(channelName string) *InterStateChannelCommand {
	this := InterStateChannelCommand{}
	this.ChannelName = channelName
	return &this
}

// NewInterStateChannelCommandWithDefaults instantiates a new InterStateChannelCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterStateChannelCommandWithDefaults() *InterStateChannelCommand {
	this := InterStateChannelCommand{}
	return &this
}

// GetCommandId returns the CommandId field value if set, zero value otherwise.
func (o *InterStateChannelCommand) GetCommandId() string {
	if o == nil || IsNil(o.CommandId) {
		var ret string
		return ret
	}
	return *o.CommandId
}

// GetCommandIdOk returns a tuple with the CommandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterStateChannelCommand) GetCommandIdOk() (*string, bool) {
	if o == nil || IsNil(o.CommandId) {
		return nil, false
	}
	return o.CommandId, true
}

// HasCommandId returns a boolean if a field has been set.
func (o *InterStateChannelCommand) HasCommandId() bool {
	if o != nil && !IsNil(o.CommandId) {
		return true
	}

	return false
}

// SetCommandId gets a reference to the given string and assigns it to the CommandId field.
func (o *InterStateChannelCommand) SetCommandId(v string) {
	o.CommandId = &v
}

// GetChannelName returns the ChannelName field value
func (o *InterStateChannelCommand) GetChannelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value
// and a boolean to check if the value has been set.
func (o *InterStateChannelCommand) GetChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelName, true
}

// SetChannelName sets field value
func (o *InterStateChannelCommand) SetChannelName(v string) {
	o.ChannelName = v
}

// GetAtLeast returns the AtLeast field value if set, zero value otherwise.
func (o *InterStateChannelCommand) GetAtLeast() int32 {
	if o == nil || IsNil(o.AtLeast) {
		var ret int32
		return ret
	}
	return *o.AtLeast
}

// GetAtLeastOk returns a tuple with the AtLeast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterStateChannelCommand) GetAtLeastOk() (*int32, bool) {
	if o == nil || IsNil(o.AtLeast) {
		return nil, false
	}
	return o.AtLeast, true
}

// HasAtLeast returns a boolean if a field has been set.
func (o *InterStateChannelCommand) HasAtLeast() bool {
	if o != nil && !IsNil(o.AtLeast) {
		return true
	}

	return false
}

// SetAtLeast gets a reference to the given int32 and assigns it to the AtLeast field.
func (o *InterStateChannelCommand) SetAtLeast(v int32) {
	o.AtLeast = &v
}

// GetAtMost returns the AtMost field value if set, zero value otherwise.
func (o *InterStateChannelCommand) GetAtMost() int32 {
	if o == nil || IsNil(o.AtMost) {
		var ret int32
		return ret
	}
	return *o.AtMost
}

// GetAtMostOk returns a tuple with the AtMost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterStateChannelCommand) GetAtMostOk() (*int32, bool) {
	if o == nil || IsNil(o.AtMost) {
		return nil, false
	}
	return o.AtMost, true
}

// HasAtMost returns a boolean if a field has been set.
func (o *InterStateChannelCommand) HasAtMost() bool {
	if o != nil && !IsNil(o.AtMost) {
		return true
	}

	return false
}

// SetAtMost gets a reference to the given int32 and assigns it to the AtMost field.
func (o *InterStateChannelCommand) SetAtMost(v int32) {
	o.AtMost = &v
}

func (o InterStateChannelCommand) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterStateChannelCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommandId) {
		toSerialize["commandId"] = o.CommandId
	}
	toSerialize["channelName"] = o.ChannelName
	if !IsNil(o.AtLeast) {
		toSerialize["atLeast"] = o.AtLeast
	}
	if !IsNil(o.AtMost) {
		toSerialize["atMost"] = o.AtMost
	}
	return toSerialize, nil
}

type NullableInterStateChannelCommand struct {
	value *InterStateChannelCommand
	isSet bool
}

func (v NullableInterStateChannelCommand) Get() *InterStateChannelCommand {
	return v.value
}

func (v *NullableInterStateChannelCommand) Set(val *InterStateChannelCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableInterStateChannelCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableInterStateChannelCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterStateChannelCommand(val *InterStateChannelCommand) *NullableInterStateChannelCommand {
	return &NullableInterStateChannelCommand{value: val, isSet: true}
}

func (v NullableInterStateChannelCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterStateChannelCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
