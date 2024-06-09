/*
ncloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DecryptRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DecryptRequest{}

// DecryptRequest struct for DecryptRequest
type DecryptRequest struct {
	Ciphertext string `json:"ciphertext"`
	Context *string `json:"context,omitempty"`
}

type _DecryptRequest DecryptRequest

// NewDecryptRequest instantiates a new DecryptRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecryptRequest(ciphertext string) *DecryptRequest {
	this := DecryptRequest{}
	this.Ciphertext = ciphertext
	return &this
}

// NewDecryptRequestWithDefaults instantiates a new DecryptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecryptRequestWithDefaults() *DecryptRequest {
	this := DecryptRequest{}
	return &this
}

// GetCiphertext returns the Ciphertext field value
func (o *DecryptRequest) GetCiphertext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ciphertext
}

// GetCiphertextOk returns a tuple with the Ciphertext field value
// and a boolean to check if the value has been set.
func (o *DecryptRequest) GetCiphertextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ciphertext, true
}

// SetCiphertext sets field value
func (o *DecryptRequest) SetCiphertext(v string) {
	o.Ciphertext = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *DecryptRequest) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecryptRequest) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *DecryptRequest) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *DecryptRequest) SetContext(v string) {
	o.Context = &v
}

func (o DecryptRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DecryptRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ciphertext"] = o.Ciphertext
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	return toSerialize, nil
}

func (o *DecryptRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ciphertext",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDecryptRequest := _DecryptRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDecryptRequest)

	if err != nil {
		return err
	}

	*o = DecryptRequest(varDecryptRequest)

	return err
}

type NullableDecryptRequest struct {
	value *DecryptRequest
	isSet bool
}

func (v NullableDecryptRequest) Get() *DecryptRequest {
	return v.value
}

func (v *NullableDecryptRequest) Set(val *DecryptRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDecryptRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDecryptRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecryptRequest(val *DecryptRequest) *NullableDecryptRequest {
	return &NullableDecryptRequest{value: val, isSet: true}
}

func (v NullableDecryptRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecryptRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


