# SkillArtifacts

## Overview

### Available Operations

* [ListContextInterfaceSkillArtifacts](#listcontextinterfaceskillartifacts) - List Skill Artifacts

## ListContextInterfaceSkillArtifacts

Retrieve list of skill artifacts

### Example Usage

<!-- UsageSnippet language="go" operationID="list-context-interface-skill-artifacts" method="get" path="/v1/context-interfaces/{interfaceId}/skills/{skillId}/artifacts" -->
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

    res, err := s.SkillArtifacts.ListContextInterfaceSkillArtifacts(ctx, operations.ListContextInterfaceSkillArtifactsRequest{
        InterfaceID: "0c822f44-914e-4d5f-b94a-c7ca41ead4ea",
        SkillID: "2c5d08df-2019-4215-af4c-ea2a74beeb75",
        Page: &components.CursorPageParameters{
            Size: sdkkonnectgo.Pointer[int64](10),
            After: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
            Before: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListSkillArtifactsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.ListContextInterfaceSkillArtifactsRequest](../../models/operations/listcontextinterfaceskillartifactsrequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.ListContextInterfaceSkillArtifactsResponse](../../models/operations/listcontextinterfaceskillartifactsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |