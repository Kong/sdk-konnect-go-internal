# KongKeyHolderType

Classifies the kind of principal holding the API key. `agent` is a
deployed AI agent, `user` is a human in the organization, `group` is
a group principal. When omitted or null, the holder is an untyped
principal.


## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.KongKeyHolderTypeAgent

// Open enum: custom values can be created with a direct type cast
custom := components.KongKeyHolderType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `KongKeyHolderTypeAgent` | agent                    |
| `KongKeyHolderTypeUser`  | user                     |
| `KongKeyHolderTypeGroup` | group                    |