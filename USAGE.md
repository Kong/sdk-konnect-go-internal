<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	sdkkonnectgointernal "github.com/Kong/sdk-konnect-go-internal"
	"github.com/Kong/sdk-konnect-go-internal/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkkonnectgointernal.New(
		sdkkonnectgointernal.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgointernal.String("<YOUR_BEARER_TOKEN_HERE>"),
		}),
	)

	res, err := s.GetUsersInternal(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.InternalUsers != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->