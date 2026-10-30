# Domain

Domain for Connected Apps (konnect or portal)

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/operations"
)

value := operations.DomainKonnect

// Open enum: custom values can be created with a direct type cast
custom := operations.Domain("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `DomainKonnect` | konnect         |
| `DomainPortal`  | portal          |