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

// checks if the EncodedObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EncodedObject{}

// EncodedObject struct for EncodedObject
type EncodedObject struct {
	Encoding *string `json:"encoding,omitempty"`
	Data *string `json:"data,omitempty"`
}

// NewEncodedObject instantiates a new EncodedObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncodedObject() *EncodedObject {
	this := EncodedObject{}
	return &this
}

// NewEncodedObjectWithDefaults instantiates a new EncodedObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncodedObjectWithDefaults() *EncodedObject {
	this := EncodedObject{}
	return &this
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *EncodedObject) GetEncoding() string {
	if o == nil || IsNil(o.Encoding) {
		var ret string
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncodedObject) GetEncodingOk() (*string, bool) {
	if o == nil || IsNil(o.Encoding) {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *EncodedObject) HasEncoding() bool {
	if o != nil && !IsNil(o.Encoding) {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given string and assigns it to the Encoding field.
func (o *EncodedObject) SetEncoding(v string) {
	o.Encoding = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EncodedObject) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncodedObject) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EncodedObject) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *EncodedObject) SetData(v string) {
	o.Data = &v
}

func (o EncodedObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EncodedObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Encoding) {
		toSerialize["encoding"] = o.Encoding
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableEncodedObject struct {
	value *EncodedObject
	isSet bool
}

func (v NullableEncodedObject) Get() *EncodedObject {
	return v.value
}

func (v *NullableEncodedObject) Set(val *EncodedObject) {
	v.value = val
	v.isSet = true
}

func (v NullableEncodedObject) IsSet() bool {
	return v.isSet
}

func (v *NullableEncodedObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncodedObject(val *EncodedObject) *NullableEncodedObject {
	return &NullableEncodedObject{value: val, isSet: true}
}

func (v NullableEncodedObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncodedObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


