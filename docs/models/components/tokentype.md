# TokenType

The type of the token

## Example Usage

```go
import (
	"github.com/Kong/sdk-konnect-go/models/components"
)

value := components.TokenTypeAccessToken

// Open enum: custom values can be created with a direct type cast
custom := components.TokenType("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `TokenTypeAccessToken`  | access_token            |
| `TokenTypeRefreshToken` | refresh_token           |