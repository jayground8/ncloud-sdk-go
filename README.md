# Client 코드 생성

1. Docker로 CLI 실행

```bash
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.6.0 generate \
    -i /local/openapi.yml \
    -g go \
    --additional-properties=withGoMod=false \
    -o /local/client
```
2. 생성된 코드 중 `client.go` 파일에 Naver Cloud 인증을 위해서 아래의 코드를 추가한다.

의존성 패키지 추가

```go
import (
	"crypto"
	...생략

	"github.com/jayground8/ncloud-sdk-go/ncloud/credentials"
	"github.com/jayground8/ncloud-sdk-go/hmac"
)
```

prepareRequest method에 아래 코드를 추가

```go
queryString := ""
if len(url.RawQuery) > 0 {
  queryString = "?" + url.RawQuery
}

if auth := credentials.LoadCredentials(credentials.DefaultCredentialsChain()); auth != nil {
  timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
  signer := hmac.NewSigner(auth.SecretKey(), crypto.SHA256)
  signature, _ := signer.Sign(method, path+queryString, auth.AccessKey(), timestamp)

  localVarRequest.Header.Add("x-ncp-apigw-timestamp", timestamp)
  localVarRequest.Header.Add("x-ncp-iam-access-key", auth.AccessKey())
  localVarRequest.Header.Add("x-ncp-apigw-signature-v1", signature)
}
```