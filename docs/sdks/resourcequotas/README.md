# ResourceQuotas
(*ResourceQuotas*)

## Overview

### Available Operations

* [UpdateResourceQuota](#updateresourcequota) - Update Resource Quota

## UpdateResourceQuota

Updates a resource quota by ID.

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

    res, err := s.ResourceQuotas.UpdateResourceQuota(ctx, "9678f205-49a1-47bb-82d9-d01cafa42a0d", components.PatchResourceQuotaRequest{
        Value: 2,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResourceQuota != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  | Example                                                                                      |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |                                                                                              |
| `resourceQuotaID`                                                                            | *string*                                                                                     | :heavy_check_mark:                                                                           | The ID of the resource quota to operate on.                                                  | 9678f205-49a1-47bb-82d9-d01cafa42a0d                                                         |
| `patchResourceQuotaRequest`                                                                  | [components.PatchResourceQuotaRequest](../../models/components/patchresourcequotarequest.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |                                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |                                                                                              |

### Response

**[*operations.UpdateResourceQuotaResponse](../../models/operations/updateresourcequotaresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |