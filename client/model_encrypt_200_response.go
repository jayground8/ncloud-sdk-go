/*
ncloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Encrypt200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Encrypt200Response{}

// Encrypt200Response struct for Encrypt200Response
type Encrypt200Response struct {
	Code *string `json:"code,omitempty"`
	Msg *string `json:"msg,omitempty"`
	Data *Encrypt200ResponseData `json:"data,omitempty"`
}

// NewEncrypt200Response instantiates a new Encrypt200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncrypt200Response() *Encrypt200Response {
	this := Encrypt200Response{}
	return &this
}

// NewEncrypt200ResponseWithDefaults instantiates a new Encrypt200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncrypt200ResponseWithDefaults() *Encrypt200Response {
	this := Encrypt200Response{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Encrypt200Response) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encrypt200Response) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Encrypt200Response) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Encrypt200Response) SetCode(v string) {
	o.Code = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *Encrypt200Response) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encrypt200Response) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *Encrypt200Response) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *Encrypt200Response) SetMsg(v string) {
	o.Msg = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Encrypt200Response) GetData() Encrypt200ResponseData {
	if o == nil || IsNil(o.Data) {
		var ret Encrypt200ResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encrypt200Response) GetDataOk() (*Encrypt200ResponseData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Encrypt200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Encrypt200ResponseData and assigns it to the Data field.
func (o *Encrypt200Response) SetData(v Encrypt200ResponseData) {
	o.Data = &v
}

func (o Encrypt200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Encrypt200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableEncrypt200Response struct {
	value *Encrypt200Response
	isSet bool
}

func (v NullableEncrypt200Response) Get() *Encrypt200Response {
	return v.value
}

func (v *NullableEncrypt200Response) Set(val *Encrypt200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEncrypt200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEncrypt200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncrypt200Response(val *Encrypt200Response) *NullableEncrypt200Response {
	return &NullableEncrypt200Response{value: val, isSet: true}
}

func (v NullableEncrypt200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncrypt200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

