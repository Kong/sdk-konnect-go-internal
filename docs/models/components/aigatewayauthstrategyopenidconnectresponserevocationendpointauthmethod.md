# AIGatewayAuthStrategyOpenIDConnectResponseRevocationEndpointAuthMethod

The revocation endpoint authentication method: : `client_secret_basic`, `client_secret_post`, `client_secret_jwt`, `private_key_jwt`, `tls_client_auth`, `self_signed_tls_client_auth`, or `none`: do not authenticate

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayAuthStrategyOpenIDConnectResponseRevocationEndpointAuthMethodClientSecretBasic

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayAuthStrategyOpenIDConnectResponseRevocationEndpointAuthMethod("custom_value")
```


## Values

| Name                                                                                            | Value                                                                                           |
| ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `AIGatewayAuthStrategyOpenIDConnectResponseRevocationEndpointAuthMethodClientSecretBasic`       | client_secret_basic                                                                             |
| `AIGatewayAuthStrategyOpenIDConnectResponseRevocationEndpointAuthMethodClientSecretJwt`         | client_secret_jwt                                                                               |
| `AIGatewayAuthStrategyOpenIDConnectResponseRevocationEndpointAuthMethodClientSecretPost`        | client_secret_post                                                                              |
| `AIGatewayAuthStrategyOpenIDConnectResponseRevocationEndpointAuthMethodNone`                    | none                                                                                            |
| `AIGatewayAuthStrategyOpenIDConnectResponseRevocationEndpointAuthMethodPrivateKeyJwt`           | private_key_jwt                                                                                 |
| `AIGatewayAuthStrategyOpenIDConnectResponseRevocationEndpointAuthMethodSelfSignedTLSClientAuth` | self_signed_tls_client_auth                                                                     |
| `AIGatewayAuthStrategyOpenIDConnectResponseRevocationEndpointAuthMethodTLSClientAuth`           | tls_client_auth                                                                                 |