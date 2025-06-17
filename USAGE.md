<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	sdkkonnectgointernal "github.com/Kong/sdk-konnect-go-internal"
	"github.com/Kong/sdk-konnect-go-internal/models/components"
	"github.com/Kong/sdk-konnect-go-internal/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkkonnectgointernal.New(
		sdkkonnectgointernal.WithSecurity(components.Security{
			PersonalAccessToken: sdkkonnectgointernal.String("<YOUR_BEARER_TOKEN_HERE>"),
		}),
	)

	res, err := s.ControlPlanes.ListControlPlanes(ctx, operations.ListControlPlanesRequest{
		PageSize:   sdkkonnectgointernal.Int64(10),
		PageNumber: sdkkonnectgointernal.Int64(1),
		Filter: &components.ControlPlaneFilterParameters{
			CloudGateway: sdkkonnectgointernal.Bool(true),
		},
		Labels: sdkkonnectgointernal.String("key:value,existCheck"),
		Sort:   sdkkonnectgointernal.String("created_at desc"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ListControlPlanesResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->