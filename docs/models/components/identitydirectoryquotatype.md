# IdentityDirectoryQuotaType

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.IdentityDirectoryQuotaTypePrincipals

// Open enum: custom values can be created with a direct type cast
custom := components.IdentityDirectoryQuotaType("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `IdentityDirectoryQuotaTypePrincipals`         | principals                                     |
| `IdentityDirectoryQuotaTypeKeyAuths`           | key_auths                                      |
| `IdentityDirectoryQuotaTypeBasicAuths`         | basic_auths                                    |
| `IdentityDirectoryQuotaTypeBasicAuthPasswords` | basic_auth_passwords                           |
| `IdentityDirectoryQuotaTypeIdentities`         | identities                                     |