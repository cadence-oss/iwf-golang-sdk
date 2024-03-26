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

// checks if the WorkflowConditionalClose type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowConditionalClose{}

// WorkflowConditionalClose struct for WorkflowConditionalClose
type WorkflowConditionalClose struct {
	ConditionalCloseType *WorkflowConditionalCloseType `json:"conditionalCloseType,omitempty"`
	ChannelName *string `json:"channelName,omitempty"`
	CloseInput *EncodedObject `json:"closeInput,omitempty"`
}

// NewWorkflowConditionalClose instantiates a new WorkflowConditionalClose object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowConditionalClose() *WorkflowConditionalClose {
	this := WorkflowConditionalClose{}
	return &this
}

// NewWorkflowConditionalCloseWithDefaults instantiates a new WorkflowConditionalClose object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowConditionalCloseWithDefaults() *WorkflowConditionalClose {
	this := WorkflowConditionalClose{}
	return &this
}

// GetConditionalCloseType returns the ConditionalCloseType field value if set, zero value otherwise.
func (o *WorkflowConditionalClose) GetConditionalCloseType() WorkflowConditionalCloseType {
	if o == nil || IsNil(o.ConditionalCloseType) {
		var ret WorkflowConditionalCloseType
		return ret
	}
	return *o.ConditionalCloseType
}

// GetConditionalCloseTypeOk returns a tuple with the ConditionalCloseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowConditionalClose) GetConditionalCloseTypeOk() (*WorkflowConditionalCloseType, bool) {
	if o == nil || IsNil(o.ConditionalCloseType) {
		return nil, false
	}
	return o.ConditionalCloseType, true
}

// HasConditionalCloseType returns a boolean if a field has been set.
func (o *WorkflowConditionalClose) HasConditionalCloseType() bool {
	if o != nil && !IsNil(o.ConditionalCloseType) {
		return true
	}

	return false
}

// SetConditionalCloseType gets a reference to the given WorkflowConditionalCloseType and assigns it to the ConditionalCloseType field.
func (o *WorkflowConditionalClose) SetConditionalCloseType(v WorkflowConditionalCloseType) {
	o.ConditionalCloseType = &v
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise.
func (o *WorkflowConditionalClose) GetChannelName() string {
	if o == nil || IsNil(o.ChannelName) {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowConditionalClose) GetChannelNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelName) {
		return nil, false
	}
	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *WorkflowConditionalClose) HasChannelName() bool {
	if o != nil && !IsNil(o.ChannelName) {
		return true
	}

	return false
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *WorkflowConditionalClose) SetChannelName(v string) {
	o.ChannelName = &v
}

// GetCloseInput returns the CloseInput field value if set, zero value otherwise.
func (o *WorkflowConditionalClose) GetCloseInput() EncodedObject {
	if o == nil || IsNil(o.CloseInput) {
		var ret EncodedObject
		return ret
	}
	return *o.CloseInput
}

// GetCloseInputOk returns a tuple with the CloseInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowConditionalClose) GetCloseInputOk() (*EncodedObject, bool) {
	if o == nil || IsNil(o.CloseInput) {
		return nil, false
	}
	return o.CloseInput, true
}

// HasCloseInput returns a boolean if a field has been set.
func (o *WorkflowConditionalClose) HasCloseInput() bool {
	if o != nil && !IsNil(o.CloseInput) {
		return true
	}

	return false
}

// SetCloseInput gets a reference to the given EncodedObject and assigns it to the CloseInput field.
func (o *WorkflowConditionalClose) SetCloseInput(v EncodedObject) {
	o.CloseInput = &v
}

func (o WorkflowConditionalClose) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowConditionalClose) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConditionalCloseType) {
		toSerialize["conditionalCloseType"] = o.ConditionalCloseType
	}
	if !IsNil(o.ChannelName) {
		toSerialize["channelName"] = o.ChannelName
	}
	if !IsNil(o.CloseInput) {
		toSerialize["closeInput"] = o.CloseInput
	}
	return toSerialize, nil
}

type NullableWorkflowConditionalClose struct {
	value *WorkflowConditionalClose
	isSet bool
}

func (v NullableWorkflowConditionalClose) Get() *WorkflowConditionalClose {
	return v.value
}

func (v *NullableWorkflowConditionalClose) Set(val *WorkflowConditionalClose) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowConditionalClose) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowConditionalClose) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowConditionalClose(val *WorkflowConditionalClose) *NullableWorkflowConditionalClose {
	return &NullableWorkflowConditionalClose{value: val, isSet: true}
}

func (v NullableWorkflowConditionalClose) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowConditionalClose) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


