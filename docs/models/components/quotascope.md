# QuotaScope

The scope a quota limit applies to.

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.QuotaScopeOrg

// Open enum: custom values can be created with a direct type cast
custom := components.QuotaScope("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `QuotaScopeOrg`       | org                   |
| `QuotaScopeDirectory` | directory             |
| `QuotaScopeDefault`   | default               |