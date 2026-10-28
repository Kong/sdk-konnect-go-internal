# RootPermissions

Whether the caller has root permissions on the organization. When `root`, the caller may perform every action on every resource, so `data` is empty rather than enumerating them, and the requested keys, filters and pagination are ignored. `root_read_only` is similar but used for impersonation.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.RootPermissionsRoot

// Open enum: custom values can be created with a direct type cast
custom := components.RootPermissions("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `RootPermissionsRoot`         | root                          |
| `RootPermissionsRootReadOnly` | root_read_only                |
| `RootPermissionsNone`         | none                          |