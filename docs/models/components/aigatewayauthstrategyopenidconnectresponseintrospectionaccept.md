# AIGatewayAuthStrategyOpenIDConnectResponseIntrospectionAccept

The value of `Accept` header for introspection requests: - `application/json`: introspection response as JSON - `application/token-introspection+jwt`: introspection response as JWT (from the current IETF draft document) - `application/jwt`: introspection response as JWT (from the obsolete IETF draft document).

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectResponseIntrospectionAcceptApplicationJSON

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectResponseIntrospectionAccept("custom_value")
```


## Values

| Name                                                                                                | Value                                                                                               |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectResponseIntrospectionAcceptApplicationJSON`                      | application/json                                                                                    |
| `AIGatewayAuthStrategyOpenIDConnectResponseIntrospectionAcceptApplicationJwt`                       | application/jwt                                                                                     |
| `AIGatewayAuthStrategyOpenIDConnectResponseIntrospectionAcceptApplicationTokenIntrospectionPlusJwt` | application/token-introspection+jwt                                                                 |