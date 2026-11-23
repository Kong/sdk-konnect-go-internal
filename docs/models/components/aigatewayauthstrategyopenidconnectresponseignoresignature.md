# AIGatewayAuthStrategyOpenIDConnectResponseIgnoreSignature

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectResponseIgnoreSignatureAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectResponseIgnoreSignature("custom_value")
```


## Values

| Name                                                                         | Value                                                                        |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectResponseIgnoreSignatureAuthorizationCode` | authorization_code                                                           |
| `AIGatewayAuthStrategyOpenIDConnectResponseIgnoreSignatureClientCredentials` | client_credentials                                                           |
| `AIGatewayAuthStrategyOpenIDConnectResponseIgnoreSignatureIntrospection`     | introspection                                                                |
| `AIGatewayAuthStrategyOpenIDConnectResponseIgnoreSignaturePassword`          | password                                                                     |
| `AIGatewayAuthStrategyOpenIDConnectResponseIgnoreSignatureRefreshToken`      | refresh_token                                                                |
| `AIGatewayAuthStrategyOpenIDConnectResponseIgnoreSignatureSession`           | session                                                                      |
| `AIGatewayAuthStrategyOpenIDConnectResponseIgnoreSignatureUserinfo`          | userinfo                                                                     |