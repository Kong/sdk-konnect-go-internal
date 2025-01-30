# ResourceConfigurations
(*ResourceConfigurations*)

## Overview

### Available Operations

* [UpdateResourceConfiguration](#updateresourceconfiguration) - Update Resource Configuration

## UpdateResourceConfiguration

Updates a resource configuration by ID.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.ResourceConfigurations.UpdateResourceConfiguration(ctx, "9678f205-49a1-47bb-82d9-d01cafa42a0d", components.PatchResourceConfigurationRequest{
        Qualifier: components.ResourceConfigurationQualifierDataPlaneGroupIdleTimeoutMinutes,
        Value: 45,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResourceConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  | Example                                                                                                      |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |                                                                                                              |
| `resourceConfigurationID`                                                                                    | *string*                                                                                                     | :heavy_check_mark:                                                                                           | The ID of the resource configuration to operate on.                                                          | 9678f205-49a1-47bb-82d9-d01cafa42a0d                                                                         |
| `patchResourceConfigurationRequest`                                                                          | [components.PatchResourceConfigurationRequest](../../models/components/patchresourceconfigurationrequest.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |                                                                                                              |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |                                                                                                              |

### Response

**[*operations.UpdateResourceConfigurationResponse](../../models/operations/updateresourceconfigurationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |