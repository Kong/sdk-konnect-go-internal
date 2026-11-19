# AIGatewayRuntimeDefaultDeploymentType

A deployment type where Konnect runs the AI Gateway data planes.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.AIGatewayRuntimeDefaultDeploymentTypeManaged

// Open enum: custom values can be created with a direct type cast
custom := components.AIGatewayRuntimeDefaultDeploymentType("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `AIGatewayRuntimeDefaultDeploymentTypeManaged`    | managed                                           |
| `AIGatewayRuntimeDefaultDeploymentTypeServerless` | serverless                                        |