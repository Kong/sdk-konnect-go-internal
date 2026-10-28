# DestinationUpdatePayloadState

Administrative state of the destination. Only enabled and disabled are settable through this field. Enabled takes a destination that the system disabled (disabled_system) back into service, with a clean failure streak.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.DestinationUpdatePayloadStateEnabled

// Open enum: custom values can be created with a direct type cast
custom := components.DestinationUpdatePayloadState("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `DestinationUpdatePayloadStateEnabled`  | enabled                                 |
| `DestinationUpdatePayloadStateDisabled` | disabled                                |