# AIGatewayAuthStrategyOpenIDConnectResponseIntrospectionEndpointAuthMethod

The introspection endpoint authentication method: : `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectResponseIntrospectionEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectResponseIntrospectionEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                                               | Value                                                                                              |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectResponseIntrospectionEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                                                |
| `AIGatewayAuthStrategyOpenIDConnectResponseIntrospectionEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                                                  |
| `AIGatewayAuthStrategyOpenIDConnectResponseIntrospectionEndpointAuthMethodClientSecretPost`        | client_secret_post                                                                                 |
| `AIGatewayAuthStrategyOpenIDConnectResponseIntrospectionEndpointAuthMethodNone`                    | none                                                                                               |
| `AIGatewayAuthStrategyOpenIDConnectResponseIntrospectionEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                                                    |
| `AIGatewayAuthStrategyOpenIDConnectResponseIntrospectionEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                                        |
| `AIGatewayAuthStrategyOpenIDConnectResponseIntrospectionEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                                                    |