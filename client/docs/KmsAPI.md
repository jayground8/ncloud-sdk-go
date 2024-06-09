# \KmsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Decrypt**](KmsAPI.md#Decrypt) | **Post** /{keyTag}/decrypt | KMS decrypt
[**Encrypt**](KmsAPI.md#Encrypt) | **Post** /{keyTag}/encrypt | KMS encrypt



## Decrypt

> Decrypt200Response Decrypt(ctx, keyTag).DecryptRequest(decryptRequest).Execute()

KMS decrypt

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	keyTag := "keyTag_example" // string | 
	decryptRequest := *openapiclient.NewDecryptRequest("Ciphertext_example") // DecryptRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAPI.Decrypt(context.Background(), keyTag).DecryptRequest(decryptRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAPI.Decrypt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Decrypt`: Decrypt200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsAPI.Decrypt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyTag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDecryptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **decryptRequest** | [**DecryptRequest**](DecryptRequest.md) |  | 

### Return type

[**Decrypt200Response**](Decrypt200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Encrypt

> Encrypt200Response Encrypt(ctx, keyTag).EncryptRequest(encryptRequest).Execute()

KMS encrypt

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	keyTag := "keyTag_example" // string | 
	encryptRequest := *openapiclient.NewEncryptRequest("Plaintext_example") // EncryptRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KmsAPI.Encrypt(context.Background(), keyTag).EncryptRequest(encryptRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KmsAPI.Encrypt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Encrypt`: Encrypt200Response
	fmt.Fprintf(os.Stdout, "Response from `KmsAPI.Encrypt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyTag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEncryptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **encryptRequest** | [**EncryptRequest**](EncryptRequest.md) |  | 

### Return type

[**Encrypt200Response**](Encrypt200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

