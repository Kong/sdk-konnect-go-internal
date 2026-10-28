# SkillSync

## Overview

### Available Operations

* [SyncContextInterfaceSkill](#synccontextinterfaceskill) - Sync Skill

## SyncContextInterfaceSkill

Schedules a skill to be synced by its defined source

### Example Usage

<!-- UsageSnippet language="go" operationID="sync-context-interface-skill" method="put" path="/v1/context-interfaces/{interfaceId}/skills/{skillId}/sync" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.SkillSync.SyncContextInterfaceSkill(ctx, operations.SyncContextInterfaceSkillRequest{
        InterfaceID: "2d63b1e7-01b0-42bb-99ae-ba5681b871d1",
        SkillID: "212069d3-f109-48ee-84e3-49195314b74d",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SkillSync != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.SyncContextInterfaceSkillRequest](../../models/operations/synccontextinterfaceskillrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.SyncContextInterfaceSkillResponse](../../models/operations/synccontextinterfaceskillresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |