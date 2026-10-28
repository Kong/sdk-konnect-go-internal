# MeshControlPlaneReadinessMode

How the audited control plane runs, and null when the scan could not tell.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MeshControlPlaneReadinessModeStandalone

// Open enum: custom values can be created with a direct type cast
custom := components.MeshControlPlaneReadinessMode("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `MeshControlPlaneReadinessModeStandalone` | standalone                                |
| `MeshControlPlaneReadinessModeZone`       | zone                                      |
| `MeshControlPlaneReadinessModeGlobal`     | global                                    |