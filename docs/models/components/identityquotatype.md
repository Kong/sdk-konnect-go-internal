# IdentityQuotaType

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.IdentityQuotaTypeDirectories

// Open enum: custom values can be created with a direct type cast
custom := components.IdentityQuotaType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `IdentityQuotaTypeDirectories`        | directories                           |
| `IdentityQuotaTypePrincipals`         | principals                            |
| `IdentityQuotaTypeKeyAuths`           | key_auths                             |
| `IdentityQuotaTypeBasicAuths`         | basic_auths                           |
| `IdentityQuotaTypeBasicAuthPasswords` | basic_auth_passwords                  |
| `IdentityQuotaTypeIdentities`         | identities                            |