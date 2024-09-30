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

// checks if the WorkflowStateExecuteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowStateExecuteRequest{}

// WorkflowStateExecuteRequest struct for WorkflowStateExecuteRequest
type WorkflowStateExecuteRequest struct {
	Context Context `json:"context"`
	WorkflowType string `json:"workflowType"`
	WorkflowStateId string `json:"workflowStateId"`
	StateInput *EncodedObject `json:"stateInput,omitempty"`
	SearchAttributes []SearchAttribute `json:"searchAttributes,omitempty"`
	DataObjects []KeyValue `json:"DataObjects,omitempty"`
	StateLocals []KeyValue `json:"stateLocals,omitempty"`
	CommandResults *CommandResults `json:"commandResults,omitempty"`
}

// NewWorkflowStateExecuteRequest instantiates a new WorkflowStateExecuteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStateExecuteRequest(context Context, workflowType string, workflowStateId string) *WorkflowStateExecuteRequest {
	this := WorkflowStateExecuteRequest{}
	this.Context = context
	this.WorkflowType = workflowType
	this.WorkflowStateId = workflowStateId
	return &this
}

// NewWorkflowStateExecuteRequestWithDefaults instantiates a new WorkflowStateExecuteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStateExecuteRequestWithDefaults() *WorkflowStateExecuteRequest {
	this := WorkflowStateExecuteRequest{}
	return &this
}

// GetContext returns the Context field value
func (o *WorkflowStateExecuteRequest) GetContext() Context {
	if o == nil {
		var ret Context
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *WorkflowStateExecuteRequest) GetContextOk() (*Context, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *WorkflowStateExecuteRequest) SetContext(v Context) {
	o.Context = v
}

// GetWorkflowType returns the WorkflowType field value
func (o *WorkflowStateExecuteRequest) GetWorkflowType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowType
}

// GetWorkflowTypeOk returns a tuple with the WorkflowType field value
// and a boolean to check if the value has been set.
func (o *WorkflowStateExecuteRequest) GetWorkflowTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowType, true
}

// SetWorkflowType sets field value
func (o *WorkflowStateExecuteRequest) SetWorkflowType(v string) {
	o.WorkflowType = v
}

// GetWorkflowStateId returns the WorkflowStateId field value
func (o *WorkflowStateExecuteRequest) GetWorkflowStateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowStateId
}

// GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field value
// and a boolean to check if the value has been set.
func (o *WorkflowStateExecuteRequest) GetWorkflowStateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowStateId, true
}

// SetWorkflowStateId sets field value
func (o *WorkflowStateExecuteRequest) SetWorkflowStateId(v string) {
	o.WorkflowStateId = v
}

// GetStateInput returns the StateInput field value if set, zero value otherwise.
func (o *WorkflowStateExecuteRequest) GetStateInput() EncodedObject {
	if o == nil || IsNil(o.StateInput) {
		var ret EncodedObject
		return ret
	}
	return *o.StateInput
}

// GetStateInputOk returns a tuple with the StateInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateExecuteRequest) GetStateInputOk() (*EncodedObject, bool) {
	if o == nil || IsNil(o.StateInput) {
		return nil, false
	}
	return o.StateInput, true
}

// HasStateInput returns a boolean if a field has been set.
func (o *WorkflowStateExecuteRequest) HasStateInput() bool {
	if o != nil && !IsNil(o.StateInput) {
		return true
	}

	return false
}

// SetStateInput gets a reference to the given EncodedObject and assigns it to the StateInput field.
func (o *WorkflowStateExecuteRequest) SetStateInput(v EncodedObject) {
	o.StateInput = &v
}

// GetSearchAttributes returns the SearchAttributes field value if set, zero value otherwise.
func (o *WorkflowStateExecuteRequest) GetSearchAttributes() []SearchAttribute {
	if o == nil || IsNil(o.SearchAttributes) {
		var ret []SearchAttribute
		return ret
	}
	return o.SearchAttributes
}

// GetSearchAttributesOk returns a tuple with the SearchAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateExecuteRequest) GetSearchAttributesOk() ([]SearchAttribute, bool) {
	if o == nil || IsNil(o.SearchAttributes) {
		return nil, false
	}
	return o.SearchAttributes, true
}

// HasSearchAttributes returns a boolean if a field has been set.
func (o *WorkflowStateExecuteRequest) HasSearchAttributes() bool {
	if o != nil && !IsNil(o.SearchAttributes) {
		return true
	}

	return false
}

// SetSearchAttributes gets a reference to the given []SearchAttribute and assigns it to the SearchAttributes field.
func (o *WorkflowStateExecuteRequest) SetSearchAttributes(v []SearchAttribute) {
	o.SearchAttributes = v
}

// GetDataObjects returns the DataObjects field value if set, zero value otherwise.
func (o *WorkflowStateExecuteRequest) GetDataObjects() []KeyValue {
	if o == nil || IsNil(o.DataObjects) {
		var ret []KeyValue
		return ret
	}
	return o.DataObjects
}

// GetDataObjectsOk returns a tuple with the DataObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateExecuteRequest) GetDataObjectsOk() ([]KeyValue, bool) {
	if o == nil || IsNil(o.DataObjects) {
		return nil, false
	}
	return o.DataObjects, true
}

// HasDataObjects returns a boolean if a field has been set.
func (o *WorkflowStateExecuteRequest) HasDataObjects() bool {
	if o != nil && !IsNil(o.DataObjects) {
		return true
	}

	return false
}

// SetDataObjects gets a reference to the given []KeyValue and assigns it to the DataObjects field.
func (o *WorkflowStateExecuteRequest) SetDataObjects(v []KeyValue) {
	o.DataObjects = v
}

// GetStateLocals returns the StateLocals field value if set, zero value otherwise.
func (o *WorkflowStateExecuteRequest) GetStateLocals() []KeyValue {
	if o == nil || IsNil(o.StateLocals) {
		var ret []KeyValue
		return ret
	}
	return o.StateLocals
}

// GetStateLocalsOk returns a tuple with the StateLocals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateExecuteRequest) GetStateLocalsOk() ([]KeyValue, bool) {
	if o == nil || IsNil(o.StateLocals) {
		return nil, false
	}
	return o.StateLocals, true
}

// HasStateLocals returns a boolean if a field has been set.
func (o *WorkflowStateExecuteRequest) HasStateLocals() bool {
	if o != nil && !IsNil(o.StateLocals) {
		return true
	}

	return false
}

// SetStateLocals gets a reference to the given []KeyValue and assigns it to the StateLocals field.
func (o *WorkflowStateExecuteRequest) SetStateLocals(v []KeyValue) {
	o.StateLocals = v
}

// GetCommandResults returns the CommandResults field value if set, zero value otherwise.
func (o *WorkflowStateExecuteRequest) GetCommandResults() CommandResults {
	if o == nil || IsNil(o.CommandResults) {
		var ret CommandResults
		return ret
	}
	return *o.CommandResults
}

// GetCommandResultsOk returns a tuple with the CommandResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateExecuteRequest) GetCommandResultsOk() (*CommandResults, bool) {
	if o == nil || IsNil(o.CommandResults) {
		return nil, false
	}
	return o.CommandResults, true
}

// HasCommandResults returns a boolean if a field has been set.
func (o *WorkflowStateExecuteRequest) HasCommandResults() bool {
	if o != nil && !IsNil(o.CommandResults) {
		return true
	}

	return false
}

// SetCommandResults gets a reference to the given CommandResults and assigns it to the CommandResults field.
func (o *WorkflowStateExecuteRequest) SetCommandResults(v CommandResults) {
	o.CommandResults = &v
}

func (o WorkflowStateExecuteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowStateExecuteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["context"] = o.Context
	toSerialize["workflowType"] = o.WorkflowType
	toSerialize["workflowStateId"] = o.WorkflowStateId
	if !IsNil(o.StateInput) {
		toSerialize["stateInput"] = o.StateInput
	}
	if !IsNil(o.SearchAttributes) {
		toSerialize["searchAttributes"] = o.SearchAttributes
	}
	if !IsNil(o.DataObjects) {
		toSerialize["DataObjects"] = o.DataObjects
	}
	if !IsNil(o.StateLocals) {
		toSerialize["stateLocals"] = o.StateLocals
	}
	if !IsNil(o.CommandResults) {
		toSerialize["commandResults"] = o.CommandResults
	}
	return toSerialize, nil
}

type NullableWorkflowStateExecuteRequest struct {
	value *WorkflowStateExecuteRequest
	isSet bool
}

func (v NullableWorkflowStateExecuteRequest) Get() *WorkflowStateExecuteRequest {
	return v.value
}

func (v *NullableWorkflowStateExecuteRequest) Set(val *WorkflowStateExecuteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStateExecuteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStateExecuteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStateExecuteRequest(val *WorkflowStateExecuteRequest) *NullableWorkflowStateExecuteRequest {
	return &NullableWorkflowStateExecuteRequest{value: val, isSet: true}
}

func (v NullableWorkflowStateExecuteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStateExecuteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


