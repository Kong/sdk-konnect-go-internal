# AIGatewayAuthStrategyOpenIDConnectResponseLoginMethods

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectResponseLoginMethodsAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectResponseLoginMethods("custom_value")
```


## Values

| Name                                                                      | Value                                                                     |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectResponseLoginMethodsAuthorizationCode` | authorization_code                                                        |
| `AIGatewayAuthStrategyOpenIDConnectResponseLoginMethodsBearer`            | bearer                                                                    |
| `AIGatewayAuthStrategyOpenIDConnectResponseLoginMethodsClientCredentials` | client_credentials                                                        |
| `AIGatewayAuthStrategyOpenIDConnectResponseLoginMethodsIntrospection`     | introspection                                                             |
| `AIGatewayAuthStrategyOpenIDConnectResponseLoginMethodsKongOauth2`        | kong_oauth2                                                               |
| `AIGatewayAuthStrategyOpenIDConnectResponseLoginMethodsPassword`          | password                                                                  |
| `AIGatewayAuthStrategyOpenIDConnectResponseLoginMethodsRefreshToken`      | refresh_token                                                             |
| `AIGatewayAuthStrategyOpenIDConnectResponseLoginMethodsSession`           | session                                                                   |
| `AIGatewayAuthStrategyOpenIDConnectResponseLoginMethodsUserinfo`          | userinfo                                                                  |