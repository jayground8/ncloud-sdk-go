# DecryptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ciphertext** | **string** |  | 
**Context** | Pointer to **string** |  | [optional] 

## Methods

### NewDecryptRequest

`func NewDecryptRequest(ciphertext string, ) *DecryptRequest`

NewDecryptRequest instantiates a new DecryptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecryptRequestWithDefaults

`func NewDecryptRequestWithDefaults() *DecryptRequest`

NewDecryptRequestWithDefaults instantiates a new DecryptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCiphertext

`func (o *DecryptRequest) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *DecryptRequest) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *DecryptRequest) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.


### GetContext

`func (o *DecryptRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DecryptRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DecryptRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *DecryptRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


