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

// checks if the WorkflowWaitForStateCompletionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowWaitForStateCompletionRequest{}

// WorkflowWaitForStateCompletionRequest struct for WorkflowWaitForStateCompletionRequest
type WorkflowWaitForStateCompletionRequest struct {
	WorkflowId string `json:"workflowId"`
	StateExecutionId string `json:"stateExecutionId"`
	WaitTimeSeconds *int32 `json:"waitTimeSeconds,omitempty"`
}

// NewWorkflowWaitForStateCompletionRequest instantiates a new WorkflowWaitForStateCompletionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWaitForStateCompletionRequest(workflowId string, stateExecutionId string) *WorkflowWaitForStateCompletionRequest {
	this := WorkflowWaitForStateCompletionRequest{}
	this.WorkflowId = workflowId
	this.StateExecutionId = stateExecutionId
	return &this
}

// NewWorkflowWaitForStateCompletionRequestWithDefaults instantiates a new WorkflowWaitForStateCompletionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWaitForStateCompletionRequestWithDefaults() *WorkflowWaitForStateCompletionRequest {
	this := WorkflowWaitForStateCompletionRequest{}
	return &this
}

// GetWorkflowId returns the WorkflowId field value
func (o *WorkflowWaitForStateCompletionRequest) GetWorkflowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value
// and a boolean to check if the value has been set.
func (o *WorkflowWaitForStateCompletionRequest) GetWorkflowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowId, true
}

// SetWorkflowId sets field value
func (o *WorkflowWaitForStateCompletionRequest) SetWorkflowId(v string) {
	o.WorkflowId = v
}

// GetStateExecutionId returns the StateExecutionId field value
func (o *WorkflowWaitForStateCompletionRequest) GetStateExecutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateExecutionId
}

// GetStateExecutionIdOk returns a tuple with the StateExecutionId field value
// and a boolean to check if the value has been set.
func (o *WorkflowWaitForStateCompletionRequest) GetStateExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateExecutionId, true
}

// SetStateExecutionId sets field value
func (o *WorkflowWaitForStateCompletionRequest) SetStateExecutionId(v string) {
	o.StateExecutionId = v
}

// GetWaitTimeSeconds returns the WaitTimeSeconds field value if set, zero value otherwise.
func (o *WorkflowWaitForStateCompletionRequest) GetWaitTimeSeconds() int32 {
	if o == nil || IsNil(o.WaitTimeSeconds) {
		var ret int32
		return ret
	}
	return *o.WaitTimeSeconds
}

// GetWaitTimeSecondsOk returns a tuple with the WaitTimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWaitForStateCompletionRequest) GetWaitTimeSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitTimeSeconds) {
		return nil, false
	}
	return o.WaitTimeSeconds, true
}

// HasWaitTimeSeconds returns a boolean if a field has been set.
func (o *WorkflowWaitForStateCompletionRequest) HasWaitTimeSeconds() bool {
	if o != nil && !IsNil(o.WaitTimeSeconds) {
		return true
	}

	return false
}

// SetWaitTimeSeconds gets a reference to the given int32 and assigns it to the WaitTimeSeconds field.
func (o *WorkflowWaitForStateCompletionRequest) SetWaitTimeSeconds(v int32) {
	o.WaitTimeSeconds = &v
}

func (o WorkflowWaitForStateCompletionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowWaitForStateCompletionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["workflowId"] = o.WorkflowId
	toSerialize["stateExecutionId"] = o.StateExecutionId
	if !IsNil(o.WaitTimeSeconds) {
		toSerialize["waitTimeSeconds"] = o.WaitTimeSeconds
	}
	return toSerialize, nil
}

type NullableWorkflowWaitForStateCompletionRequest struct {
	value *WorkflowWaitForStateCompletionRequest
	isSet bool
}

func (v NullableWorkflowWaitForStateCompletionRequest) Get() *WorkflowWaitForStateCompletionRequest {
	return v.value
}

func (v *NullableWorkflowWaitForStateCompletionRequest) Set(val *WorkflowWaitForStateCompletionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWaitForStateCompletionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWaitForStateCompletionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWaitForStateCompletionRequest(val *WorkflowWaitForStateCompletionRequest) *NullableWorkflowWaitForStateCompletionRequest {
	return &NullableWorkflowWaitForStateCompletionRequest{value: val, isSet: true}
}

func (v NullableWorkflowWaitForStateCompletionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWaitForStateCompletionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


