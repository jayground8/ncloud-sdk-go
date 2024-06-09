# EncryptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plaintext** | **string** |  | 
**Context** | Pointer to **string** |  | [optional] 

## Methods

### NewEncryptRequest

`func NewEncryptRequest(plaintext string, ) *EncryptRequest`

NewEncryptRequest instantiates a new EncryptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptRequestWithDefaults

`func NewEncryptRequestWithDefaults() *EncryptRequest`

NewEncryptRequestWithDefaults instantiates a new EncryptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaintext

`func (o *EncryptRequest) GetPlaintext() string`

GetPlaintext returns the Plaintext field if non-nil, zero value otherwise.

### GetPlaintextOk

`func (o *EncryptRequest) GetPlaintextOk() (*string, bool)`

GetPlaintextOk returns a tuple with the Plaintext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaintext

`func (o *EncryptRequest) SetPlaintext(v string)`

SetPlaintext sets Plaintext field to given value.


### GetContext

`func (o *EncryptRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *EncryptRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *EncryptRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *EncryptRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


