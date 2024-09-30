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

// checks if the WorkflowStartRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowStartRequest{}

// WorkflowStartRequest struct for WorkflowStartRequest
type WorkflowStartRequest struct {
	WorkflowId string `json:"workflowId"`
	IwfWorkflowType string `json:"iwfWorkflowType"`
	WorkflowTimeoutSeconds int32 `json:"workflowTimeoutSeconds"`
	IwfWorkerUrl string `json:"iwfWorkerUrl"`
	StartStateId *string `json:"startStateId,omitempty"`
	WaitForCompletionStateIds []string `json:"waitForCompletionStateIds,omitempty"`
	WaitForCompletionStateExecutionIds []string `json:"waitForCompletionStateExecutionIds,omitempty"`
	StateInput *EncodedObject `json:"stateInput,omitempty"`
	StateOptions *WorkflowStateOptions `json:"stateOptions,omitempty"`
	WorkflowStartOptions *WorkflowStartOptions `json:"workflowStartOptions,omitempty"`
}

// NewWorkflowStartRequest instantiates a new WorkflowStartRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStartRequest(workflowId string, iwfWorkflowType string, workflowTimeoutSeconds int32, iwfWorkerUrl string) *WorkflowStartRequest {
	this := WorkflowStartRequest{}
	this.WorkflowId = workflowId
	this.IwfWorkflowType = iwfWorkflowType
	this.WorkflowTimeoutSeconds = workflowTimeoutSeconds
	this.IwfWorkerUrl = iwfWorkerUrl
	return &this
}

// NewWorkflowStartRequestWithDefaults instantiates a new WorkflowStartRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStartRequestWithDefaults() *WorkflowStartRequest {
	this := WorkflowStartRequest{}
	return &this
}

// GetWorkflowId returns the WorkflowId field value
func (o *WorkflowStartRequest) GetWorkflowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value
// and a boolean to check if the value has been set.
func (o *WorkflowStartRequest) GetWorkflowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowId, true
}

// SetWorkflowId sets field value
func (o *WorkflowStartRequest) SetWorkflowId(v string) {
	o.WorkflowId = v
}

// GetIwfWorkflowType returns the IwfWorkflowType field value
func (o *WorkflowStartRequest) GetIwfWorkflowType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IwfWorkflowType
}

// GetIwfWorkflowTypeOk returns a tuple with the IwfWorkflowType field value
// and a boolean to check if the value has been set.
func (o *WorkflowStartRequest) GetIwfWorkflowTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IwfWorkflowType, true
}

// SetIwfWorkflowType sets field value
func (o *WorkflowStartRequest) SetIwfWorkflowType(v string) {
	o.IwfWorkflowType = v
}

// GetWorkflowTimeoutSeconds returns the WorkflowTimeoutSeconds field value
func (o *WorkflowStartRequest) GetWorkflowTimeoutSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WorkflowTimeoutSeconds
}

// GetWorkflowTimeoutSecondsOk returns a tuple with the WorkflowTimeoutSeconds field value
// and a boolean to check if the value has been set.
func (o *WorkflowStartRequest) GetWorkflowTimeoutSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowTimeoutSeconds, true
}

// SetWorkflowTimeoutSeconds sets field value
func (o *WorkflowStartRequest) SetWorkflowTimeoutSeconds(v int32) {
	o.WorkflowTimeoutSeconds = v
}

// GetIwfWorkerUrl returns the IwfWorkerUrl field value
func (o *WorkflowStartRequest) GetIwfWorkerUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IwfWorkerUrl
}

// GetIwfWorkerUrlOk returns a tuple with the IwfWorkerUrl field value
// and a boolean to check if the value has been set.
func (o *WorkflowStartRequest) GetIwfWorkerUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IwfWorkerUrl, true
}

// SetIwfWorkerUrl sets field value
func (o *WorkflowStartRequest) SetIwfWorkerUrl(v string) {
	o.IwfWorkerUrl = v
}

// GetStartStateId returns the StartStateId field value if set, zero value otherwise.
func (o *WorkflowStartRequest) GetStartStateId() string {
	if o == nil || IsNil(o.StartStateId) {
		var ret string
		return ret
	}
	return *o.StartStateId
}

// GetStartStateIdOk returns a tuple with the StartStateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStartRequest) GetStartStateIdOk() (*string, bool) {
	if o == nil || IsNil(o.StartStateId) {
		return nil, false
	}
	return o.StartStateId, true
}

// HasStartStateId returns a boolean if a field has been set.
func (o *WorkflowStartRequest) HasStartStateId() bool {
	if o != nil && !IsNil(o.StartStateId) {
		return true
	}

	return false
}

// SetStartStateId gets a reference to the given string and assigns it to the StartStateId field.
func (o *WorkflowStartRequest) SetStartStateId(v string) {
	o.StartStateId = &v
}

// GetWaitForCompletionStateIds returns the WaitForCompletionStateIds field value if set, zero value otherwise.
func (o *WorkflowStartRequest) GetWaitForCompletionStateIds() []string {
	if o == nil || IsNil(o.WaitForCompletionStateIds) {
		var ret []string
		return ret
	}
	return o.WaitForCompletionStateIds
}

// GetWaitForCompletionStateIdsOk returns a tuple with the WaitForCompletionStateIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStartRequest) GetWaitForCompletionStateIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.WaitForCompletionStateIds) {
		return nil, false
	}
	return o.WaitForCompletionStateIds, true
}

// HasWaitForCompletionStateIds returns a boolean if a field has been set.
func (o *WorkflowStartRequest) HasWaitForCompletionStateIds() bool {
	if o != nil && !IsNil(o.WaitForCompletionStateIds) {
		return true
	}

	return false
}

// SetWaitForCompletionStateIds gets a reference to the given []string and assigns it to the WaitForCompletionStateIds field.
func (o *WorkflowStartRequest) SetWaitForCompletionStateIds(v []string) {
	o.WaitForCompletionStateIds = v
}

// GetWaitForCompletionStateExecutionIds returns the WaitForCompletionStateExecutionIds field value if set, zero value otherwise.
func (o *WorkflowStartRequest) GetWaitForCompletionStateExecutionIds() []string {
	if o == nil || IsNil(o.WaitForCompletionStateExecutionIds) {
		var ret []string
		return ret
	}
	return o.WaitForCompletionStateExecutionIds
}

// GetWaitForCompletionStateExecutionIdsOk returns a tuple with the WaitForCompletionStateExecutionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStartRequest) GetWaitForCompletionStateExecutionIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.WaitForCompletionStateExecutionIds) {
		return nil, false
	}
	return o.WaitForCompletionStateExecutionIds, true
}

// HasWaitForCompletionStateExecutionIds returns a boolean if a field has been set.
func (o *WorkflowStartRequest) HasWaitForCompletionStateExecutionIds() bool {
	if o != nil && !IsNil(o.WaitForCompletionStateExecutionIds) {
		return true
	}

	return false
}

// SetWaitForCompletionStateExecutionIds gets a reference to the given []string and assigns it to the WaitForCompletionStateExecutionIds field.
func (o *WorkflowStartRequest) SetWaitForCompletionStateExecutionIds(v []string) {
	o.WaitForCompletionStateExecutionIds = v
}

// GetStateInput returns the StateInput field value if set, zero value otherwise.
func (o *WorkflowStartRequest) GetStateInput() EncodedObject {
	if o == nil || IsNil(o.StateInput) {
		var ret EncodedObject
		return ret
	}
	return *o.StateInput
}

// GetStateInputOk returns a tuple with the StateInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStartRequest) GetStateInputOk() (*EncodedObject, bool) {
	if o == nil || IsNil(o.StateInput) {
		return nil, false
	}
	return o.StateInput, true
}

// HasStateInput returns a boolean if a field has been set.
func (o *WorkflowStartRequest) HasStateInput() bool {
	if o != nil && !IsNil(o.StateInput) {
		return true
	}

	return false
}

// SetStateInput gets a reference to the given EncodedObject and assigns it to the StateInput field.
func (o *WorkflowStartRequest) SetStateInput(v EncodedObject) {
	o.StateInput = &v
}

// GetStateOptions returns the StateOptions field value if set, zero value otherwise.
func (o *WorkflowStartRequest) GetStateOptions() WorkflowStateOptions {
	if o == nil || IsNil(o.StateOptions) {
		var ret WorkflowStateOptions
		return ret
	}
	return *o.StateOptions
}

// GetStateOptionsOk returns a tuple with the StateOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStartRequest) GetStateOptionsOk() (*WorkflowStateOptions, bool) {
	if o == nil || IsNil(o.StateOptions) {
		return nil, false
	}
	return o.StateOptions, true
}

// HasStateOptions returns a boolean if a field has been set.
func (o *WorkflowStartRequest) HasStateOptions() bool {
	if o != nil && !IsNil(o.StateOptions) {
		return true
	}

	return false
}

// SetStateOptions gets a reference to the given WorkflowStateOptions and assigns it to the StateOptions field.
func (o *WorkflowStartRequest) SetStateOptions(v WorkflowStateOptions) {
	o.StateOptions = &v
}

// GetWorkflowStartOptions returns the WorkflowStartOptions field value if set, zero value otherwise.
func (o *WorkflowStartRequest) GetWorkflowStartOptions() WorkflowStartOptions {
	if o == nil || IsNil(o.WorkflowStartOptions) {
		var ret WorkflowStartOptions
		return ret
	}
	return *o.WorkflowStartOptions
}

// GetWorkflowStartOptionsOk returns a tuple with the WorkflowStartOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStartRequest) GetWorkflowStartOptionsOk() (*WorkflowStartOptions, bool) {
	if o == nil || IsNil(o.WorkflowStartOptions) {
		return nil, false
	}
	return o.WorkflowStartOptions, true
}

// HasWorkflowStartOptions returns a boolean if a field has been set.
func (o *WorkflowStartRequest) HasWorkflowStartOptions() bool {
	if o != nil && !IsNil(o.WorkflowStartOptions) {
		return true
	}

	return false
}

// SetWorkflowStartOptions gets a reference to the given WorkflowStartOptions and assigns it to the WorkflowStartOptions field.
func (o *WorkflowStartRequest) SetWorkflowStartOptions(v WorkflowStartOptions) {
	o.WorkflowStartOptions = &v
}

func (o WorkflowStartRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowStartRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["workflowId"] = o.WorkflowId
	toSerialize["iwfWorkflowType"] = o.IwfWorkflowType
	toSerialize["workflowTimeoutSeconds"] = o.WorkflowTimeoutSeconds
	toSerialize["iwfWorkerUrl"] = o.IwfWorkerUrl
	if !IsNil(o.StartStateId) {
		toSerialize["startStateId"] = o.StartStateId
	}
	if !IsNil(o.WaitForCompletionStateIds) {
		toSerialize["waitForCompletionStateIds"] = o.WaitForCompletionStateIds
	}
	if !IsNil(o.WaitForCompletionStateExecutionIds) {
		toSerialize["waitForCompletionStateExecutionIds"] = o.WaitForCompletionStateExecutionIds
	}
	if !IsNil(o.StateInput) {
		toSerialize["stateInput"] = o.StateInput
	}
	if !IsNil(o.StateOptions) {
		toSerialize["stateOptions"] = o.StateOptions
	}
	if !IsNil(o.WorkflowStartOptions) {
		toSerialize["workflowStartOptions"] = o.WorkflowStartOptions
	}
	return toSerialize, nil
}

type NullableWorkflowStartRequest struct {
	value *WorkflowStartRequest
	isSet bool
}

func (v NullableWorkflowStartRequest) Get() *WorkflowStartRequest {
	return v.value
}

func (v *NullableWorkflowStartRequest) Set(val *WorkflowStartRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStartRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStartRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStartRequest(val *WorkflowStartRequest) *NullableWorkflowStartRequest {
	return &NullableWorkflowStartRequest{value: val, isSet: true}
}

func (v NullableWorkflowStartRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStartRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


