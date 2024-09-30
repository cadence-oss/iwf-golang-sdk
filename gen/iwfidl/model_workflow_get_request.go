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

// checks if the WorkflowGetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowGetRequest{}

// WorkflowGetRequest struct for WorkflowGetRequest
type WorkflowGetRequest struct {
	WorkflowId      string  `json:"workflowId"`
	WorkflowRunId   *string `json:"workflowRunId,omitempty"`
	NeedsResults    *bool   `json:"needsResults,omitempty"`
	WaitTimeSeconds *int32  `json:"waitTimeSeconds,omitempty"`
}

// NewWorkflowGetRequest instantiates a new WorkflowGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowGetRequest(workflowId string) *WorkflowGetRequest {
	this := WorkflowGetRequest{}
	this.WorkflowId = workflowId
	return &this
}

// NewWorkflowGetRequestWithDefaults instantiates a new WorkflowGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowGetRequestWithDefaults() *WorkflowGetRequest {
	this := WorkflowGetRequest{}
	return &this
}

// GetWorkflowId returns the WorkflowId field value
func (o *WorkflowGetRequest) GetWorkflowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value
// and a boolean to check if the value has been set.
func (o *WorkflowGetRequest) GetWorkflowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowId, true
}

// SetWorkflowId sets field value
func (o *WorkflowGetRequest) SetWorkflowId(v string) {
	o.WorkflowId = v
}

// GetWorkflowRunId returns the WorkflowRunId field value if set, zero value otherwise.
func (o *WorkflowGetRequest) GetWorkflowRunId() string {
	if o == nil || IsNil(o.WorkflowRunId) {
		var ret string
		return ret
	}
	return *o.WorkflowRunId
}

// GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowGetRequest) GetWorkflowRunIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorkflowRunId) {
		return nil, false
	}
	return o.WorkflowRunId, true
}

// HasWorkflowRunId returns a boolean if a field has been set.
func (o *WorkflowGetRequest) HasWorkflowRunId() bool {
	if o != nil && !IsNil(o.WorkflowRunId) {
		return true
	}

	return false
}

// SetWorkflowRunId gets a reference to the given string and assigns it to the WorkflowRunId field.
func (o *WorkflowGetRequest) SetWorkflowRunId(v string) {
	o.WorkflowRunId = &v
}

// GetNeedsResults returns the NeedsResults field value if set, zero value otherwise.
func (o *WorkflowGetRequest) GetNeedsResults() bool {
	if o == nil || IsNil(o.NeedsResults) {
		var ret bool
		return ret
	}
	return *o.NeedsResults
}

// GetNeedsResultsOk returns a tuple with the NeedsResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowGetRequest) GetNeedsResultsOk() (*bool, bool) {
	if o == nil || IsNil(o.NeedsResults) {
		return nil, false
	}
	return o.NeedsResults, true
}

// HasNeedsResults returns a boolean if a field has been set.
func (o *WorkflowGetRequest) HasNeedsResults() bool {
	if o != nil && !IsNil(o.NeedsResults) {
		return true
	}

	return false
}

// SetNeedsResults gets a reference to the given bool and assigns it to the NeedsResults field.
func (o *WorkflowGetRequest) SetNeedsResults(v bool) {
	o.NeedsResults = &v
}

// GetWaitTimeSeconds returns the WaitTimeSeconds field value if set, zero value otherwise.
func (o *WorkflowGetRequest) GetWaitTimeSeconds() int32 {
	if o == nil || IsNil(o.WaitTimeSeconds) {
		var ret int32
		return ret
	}
	return *o.WaitTimeSeconds
}

// GetWaitTimeSecondsOk returns a tuple with the WaitTimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowGetRequest) GetWaitTimeSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.WaitTimeSeconds) {
		return nil, false
	}
	return o.WaitTimeSeconds, true
}

// HasWaitTimeSeconds returns a boolean if a field has been set.
func (o *WorkflowGetRequest) HasWaitTimeSeconds() bool {
	if o != nil && !IsNil(o.WaitTimeSeconds) {
		return true
	}

	return false
}

// SetWaitTimeSeconds gets a reference to the given int32 and assigns it to the WaitTimeSeconds field.
func (o *WorkflowGetRequest) SetWaitTimeSeconds(v int32) {
	o.WaitTimeSeconds = &v
}

func (o WorkflowGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowGetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["workflowId"] = o.WorkflowId
	if !IsNil(o.WorkflowRunId) {
		toSerialize["workflowRunId"] = o.WorkflowRunId
	}
	if !IsNil(o.NeedsResults) {
		toSerialize["needsResults"] = o.NeedsResults
	}
	if !IsNil(o.WaitTimeSeconds) {
		toSerialize["waitTimeSeconds"] = o.WaitTimeSeconds
	}
	return toSerialize, nil
}

type NullableWorkflowGetRequest struct {
	value *WorkflowGetRequest
	isSet bool
}

func (v NullableWorkflowGetRequest) Get() *WorkflowGetRequest {
	return v.value
}

func (v *NullableWorkflowGetRequest) Set(val *WorkflowGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowGetRequest(val *WorkflowGetRequest) *NullableWorkflowGetRequest {
	return &NullableWorkflowGetRequest{value: val, isSet: true}
}

func (v NullableWorkflowGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
