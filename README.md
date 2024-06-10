# Client 코드 생성

1. Docker로 CLI 실행

```bash
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v7.6.0 generate \
    -i /local/openapi.yml \
    -g go \
    --additional-properties=withGoMod=false \
    -o /local/client
```
2. 생성된 코드 중 `client.go` 파일에 Naver Cloud 인증을 위해서 아래의 코드를 추가

의존성 패키지 추가

```go
import (
	"crypto"
	...생략

	"github.com/jayground8/ncloud-sdk-go/hmac"
)
```

prepareRequest method에 아래 코드를 추가

```go
queryString := ""
if len(url.RawQuery) > 0 {
	queryString = "?" + url.RawQuery
}

timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
signer := hmac.NewSigner(c.cfg.Credentials.SecretKey(), crypto.SHA256)
signature, _ := signer.Sign(method, path+queryString, c.cfg.Credentials.AccessKey(), timestamp)

localVarRequest.Header.Add("x-ncp-apigw-timestamp", timestamp)
localVarRequest.Header.Add("x-ncp-iam-access-key", c.cfg.Credentials.AccessKey())
localVarRequest.Header.Add("x-ncp-apigw-signature-v1", signature)
```

3. 생성된 코드 중 `configuration.go` 파일에 NCLOUD_KMS_API_GW 환경변수 처리를 추가

네이버 공공 클라우드 아닌 경우에는 아래와 같이 환경변수를 셋팅.

```bash
export NCLOUD_KMS_API_GW=kms.apigw.ntruss.com
```

의존성 패키지 추가

```go
import (
	...생략
	ncloud "github.com/jayground8/ncloud-sdk-go/ncloud/credentials"
)
```

코드 추가 & 변경

```go
func getKmsUrl() string {
	u, err := url.Parse("https://kms.apigw.gov-ntruss.com/keys/v2")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed parse kms url\n")
		panic(err)
	}
	if overridedValue := os.Getenv("NCLOUD_KMS_API_GW"); overridedValue != "" {
		u.Host = overridedValue
	}
	return u.String()
}

// NewConfiguration returns a new Configuration object
func NewConfiguration(credentials *ncloud.Credentials) *Configuration {
	cfg := &Configuration{
		DefaultHeader:    make(map[string]string),
		UserAgent:        "OpenAPI-Generator/1.0.0/go",
		Debug:            false,
		Servers:          ServerConfigurations{
			{
				URL: "",
				Description: "No description provided",
			},
		},
		OperationServers: map[string]ServerConfigurations{
			"KmsAPIService.Decrypt": {
				{
					URL: getKmsUrl(),
					Description: "No description provided",
				},
			},
			"KmsAPIService.Encrypt": {
				{
					URL: getKmsUrl(),
					Description: "No description provided",
				},
			},
		},
		Credentials: credentials,
	}

	if !cfg.Credentials.Valid() {
		cfg.Credentials = ncloud.LoadCredentials(ncloud.DefaultCredentialsChain())
	}

	return cfg
}
```