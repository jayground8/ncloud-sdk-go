package main

import (
	"context"
	"fmt"
	"os"
	b64 "encoding/base64"
	openapi "github.com/jayground8/ncloud-sdk-go/client"
)

func encrypt(client *openapi.APIClient, keyTag string, plaintext string) string {
	body := openapi.NewEncryptRequest(plaintext)
	req := client.KmsAPI.Encrypt(context.Background(), keyTag).EncryptRequest(*body)
			
	value, res, err := client.KmsAPI.EncryptExecute(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
		panic(err)
	}
	if res.StatusCode != 200 {
		return ""
	}

	data := value.GetData()
	ciphertext := data.GetCiphertext()
	return ciphertext
}

func decrypt(client *openapi.APIClient, keyTag string, ciphertext string) string {
	body := openapi.NewDecryptRequest(ciphertext)
	req := client.KmsAPI.Decrypt(context.Background(), keyTag).DecryptRequest(*body)
			
	value, res, err := client.KmsAPI.DecryptExecute(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res)
		panic(err)
	}
	if res.StatusCode != 200 {
		return ""
	}

	data := value.GetData()
	plaintext := data.GetPlaintext()
	return plaintext
}

func main() {
	keyTag := os.Getenv("KMS_KEY_TAG")
	config := openapi.NewConfiguration()
	client := openapi.NewAPIClient(config)
	b64EncodedPlaintext := b64.StdEncoding.EncodeToString([]byte("test"))
	ciphertext := encrypt(client, keyTag, b64EncodedPlaintext)
	println(ciphertext)
	plaintext := decrypt(client, keyTag, ciphertext)
	b64DecodedPlaintext, _ := b64.StdEncoding.DecodeString(plaintext)
	println(string(b64DecodedPlaintext))
}
