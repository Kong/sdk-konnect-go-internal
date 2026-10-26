# MeshControlPlaneReadinessSeverity

How much a finding matters to the upgrade. A blocker stops it, and everything else is context the reader may act on or ignore.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.MeshControlPlaneReadinessSeverityBlocker

// Open enum: custom values can be created with a direct type cast
custom := components.MeshControlPlaneReadinessSeverity("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `MeshControlPlaneReadinessSeverityBlocker` | blocker                                    |
| `MeshControlPlaneReadinessSeverityInfo`    | info                                       |