# AIGatewayAuthStrategyOpenIDConnectResponseDisableSession

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectResponseDisableSessionAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectResponseDisableSession("custom_value")
```


## Values

| Name                                                                        | Value                                                                       |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectResponseDisableSessionAuthorizationCode` | authorization_code                                                          |
| `AIGatewayAuthStrategyOpenIDConnectResponseDisableSessionBearer`            | bearer                                                                      |
| `AIGatewayAuthStrategyOpenIDConnectResponseDisableSessionClientCredentials` | client_credentials                                                          |
| `AIGatewayAuthStrategyOpenIDConnectResponseDisableSessionIntrospection`     | introspection                                                               |
| `AIGatewayAuthStrategyOpenIDConnectResponseDisableSessionKongOauth2`        | kong_oauth2                                                                 |
| `AIGatewayAuthStrategyOpenIDConnectResponseDisableSessionPassword`          | password                                                                    |
| `AIGatewayAuthStrategyOpenIDConnectResponseDisableSessionRefreshToken`      | refresh_token                                                               |
| `AIGatewayAuthStrategyOpenIDConnectResponseDisableSessionSession`           | session                                                                     |
| `AIGatewayAuthStrategyOpenIDConnectResponseDisableSessionUserinfo`          | userinfo                                                                    |