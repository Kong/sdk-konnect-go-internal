# AIGatewayAuthStrategyOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethod

The pushed authorization request endpoint authentication method: `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                                                            | Value                                                                                                           |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                                                             |
| `AIGatewayAuthStrategyOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                                                               |
| `AIGatewayAuthStrategyOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethodClientSecretPost`        | client_secret_post                                                                                              |
| `AIGatewayAuthStrategyOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethodNone`                    | none                                                                                                            |
| `AIGatewayAuthStrategyOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                                                                 |
| `AIGatewayAuthStrategyOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                                                     |
| `AIGatewayAuthStrategyOpenIDConnectResponsePushedAuthorizationRequestEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                                                                 |