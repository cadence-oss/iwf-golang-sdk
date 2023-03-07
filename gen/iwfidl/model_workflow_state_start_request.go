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

// checks if the WorkflowStateStartRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowStateStartRequest{}

// WorkflowStateStartRequest struct for WorkflowStateStartRequest
type WorkflowStateStartRequest struct {
	Context Context `json:"context"`
	WorkflowType string `json:"workflowType"`
	WorkflowStateId string `json:"workflowStateId"`
	StateInput *EncodedObject `json:"stateInput,omitempty"`
	SearchAttributes []SearchAttribute `json:"searchAttributes,omitempty"`
	DataObjects []KeyValue `json:"dataObjects,omitempty"`
}

// NewWorkflowStateStartRequest instantiates a new WorkflowStateStartRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStateStartRequest(context Context, workflowType string, workflowStateId string) *WorkflowStateStartRequest {
	this := WorkflowStateStartRequest{}
	this.Context = context
	this.WorkflowType = workflowType
	this.WorkflowStateId = workflowStateId
	return &this
}

// NewWorkflowStateStartRequestWithDefaults instantiates a new WorkflowStateStartRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStateStartRequestWithDefaults() *WorkflowStateStartRequest {
	this := WorkflowStateStartRequest{}
	return &this
}

// GetContext returns the Context field value
func (o *WorkflowStateStartRequest) GetContext() Context {
	if o == nil {
		var ret Context
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartRequest) GetContextOk() (*Context, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *WorkflowStateStartRequest) SetContext(v Context) {
	o.Context = v
}

// GetWorkflowType returns the WorkflowType field value
func (o *WorkflowStateStartRequest) GetWorkflowType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowType
}

// GetWorkflowTypeOk returns a tuple with the WorkflowType field value
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartRequest) GetWorkflowTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowType, true
}

// SetWorkflowType sets field value
func (o *WorkflowStateStartRequest) SetWorkflowType(v string) {
	o.WorkflowType = v
}

// GetWorkflowStateId returns the WorkflowStateId field value
func (o *WorkflowStateStartRequest) GetWorkflowStateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowStateId
}

// GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field value
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartRequest) GetWorkflowStateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowStateId, true
}

// SetWorkflowStateId sets field value
func (o *WorkflowStateStartRequest) SetWorkflowStateId(v string) {
	o.WorkflowStateId = v
}

// GetStateInput returns the StateInput field value if set, zero value otherwise.
func (o *WorkflowStateStartRequest) GetStateInput() EncodedObject {
	if o == nil || IsNil(o.StateInput) {
		var ret EncodedObject
		return ret
	}
	return *o.StateInput
}

// GetStateInputOk returns a tuple with the StateInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartRequest) GetStateInputOk() (*EncodedObject, bool) {
	if o == nil || IsNil(o.StateInput) {
		return nil, false
	}
	return o.StateInput, true
}

// HasStateInput returns a boolean if a field has been set.
func (o *WorkflowStateStartRequest) HasStateInput() bool {
	if o != nil && !IsNil(o.StateInput) {
		return true
	}

	return false
}

// SetStateInput gets a reference to the given EncodedObject and assigns it to the StateInput field.
func (o *WorkflowStateStartRequest) SetStateInput(v EncodedObject) {
	o.StateInput = &v
}

// GetSearchAttributes returns the SearchAttributes field value if set, zero value otherwise.
func (o *WorkflowStateStartRequest) GetSearchAttributes() []SearchAttribute {
	if o == nil || IsNil(o.SearchAttributes) {
		var ret []SearchAttribute
		return ret
	}
	return o.SearchAttributes
}

// GetSearchAttributesOk returns a tuple with the SearchAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartRequest) GetSearchAttributesOk() ([]SearchAttribute, bool) {
	if o == nil || IsNil(o.SearchAttributes) {
		return nil, false
	}
	return o.SearchAttributes, true
}

// HasSearchAttributes returns a boolean if a field has been set.
func (o *WorkflowStateStartRequest) HasSearchAttributes() bool {
	if o != nil && !IsNil(o.SearchAttributes) {
		return true
	}

	return false
}

// SetSearchAttributes gets a reference to the given []SearchAttribute and assigns it to the SearchAttributes field.
func (o *WorkflowStateStartRequest) SetSearchAttributes(v []SearchAttribute) {
	o.SearchAttributes = v
}

// GetDataObjects returns the DataObjects field value if set, zero value otherwise.
func (o *WorkflowStateStartRequest) GetDataObjects() []KeyValue {
	if o == nil || IsNil(o.DataObjects) {
		var ret []KeyValue
		return ret
	}
	return o.DataObjects
}

// GetDataObjectsOk returns a tuple with the DataObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStateStartRequest) GetDataObjectsOk() ([]KeyValue, bool) {
	if o == nil || IsNil(o.DataObjects) {
		return nil, false
	}
	return o.DataObjects, true
}

// HasDataObjects returns a boolean if a field has been set.
func (o *WorkflowStateStartRequest) HasDataObjects() bool {
	if o != nil && !IsNil(o.DataObjects) {
		return true
	}

	return false
}

// SetDataObjects gets a reference to the given []KeyValue and assigns it to the DataObjects field.
func (o *WorkflowStateStartRequest) SetDataObjects(v []KeyValue) {
	o.DataObjects = v
}

func (o WorkflowStateStartRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowStateStartRequest) ToMap() (map[string]interface{}, error) {
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
		toSerialize["dataObjects"] = o.DataObjects
	}
	return toSerialize, nil
}

type NullableWorkflowStateStartRequest struct {
	value *WorkflowStateStartRequest
	isSet bool
}

func (v NullableWorkflowStateStartRequest) Get() *WorkflowStateStartRequest {
	return v.value
}

func (v *NullableWorkflowStateStartRequest) Set(val *WorkflowStateStartRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStateStartRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStateStartRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStateStartRequest(val *WorkflowStateStartRequest) *NullableWorkflowStateStartRequest {
	return &NullableWorkflowStateStartRequest{value: val, isSet: true}
}

func (v NullableWorkflowStateStartRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStateStartRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


