# CreateAIGatewayAgentRequestProtocol

The protocol used to reach the `targets`. Only valid with `targets`. If unset, the gateway uses `https`.

**Requires a minimum runtime version of `2.4`**.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.CreateAIGatewayAgentRequestProtocolHTTP

// Open enum: custom values can be created with a direct type cast
custom := components.CreateAIGatewayAgentRequestProtocol("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `CreateAIGatewayAgentRequestProtocolHTTP`  | http                                       |
| `CreateAIGatewayAgentRequestProtocolHTTPS` | https                                      |