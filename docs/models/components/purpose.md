# Purpose

The model feature that uses the datastore. `semantic-balancer` supplies the connection of the
model vector store.

**`session-affinity` requires a minimum runtime version of `2.4`**.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.PurposeSemanticBalancer

// Open enum: custom values can be created with a direct type cast
custom := components.Purpose("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `PurposeSemanticBalancer` | semantic-balancer         |
| `PurposeSessionAffinity`  | session-affinity          |