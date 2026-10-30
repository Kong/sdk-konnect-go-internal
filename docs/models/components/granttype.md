# GrantType

OAuth 2.0 grant type

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.GrantTypeAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.GrantType("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `GrantTypeAuthorizationCode`                    | authorization_code                              |
| `GrantTypeImplicit`                             | implicit                                        |
| `GrantTypeClientCredentials`                    | client_credentials                              |
| `GrantTypeRefreshToken`                         | refresh_token                                   |
| `GrantTypeUrnIetfParamsOauthGrantTypeJwtBearer` | urn:ietf:params:oauth:grant-type:jwt-bearer     |