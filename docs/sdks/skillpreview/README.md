# SkillPreview

## Overview

### Available Operations

* [PreviewContextInterfaceSkill](#previewcontextinterfaceskill) - Preview a Skill

## PreviewContextInterfaceSkill

Fetches the SKILL.md content from the given source, parses and validates its frontmatter, and returns the result without persisting anything.


### Example Usage

<!-- UsageSnippet language="go" operationID="preview-context-interface-skill" method="post" path="/v1/context-interfaces/{interfaceId}/skills/preview" -->
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

    res, err := s.SkillPreview.PreviewContextInterfaceSkill(ctx, "b281c961-b0f6-4c45-b08f-14cc7823aa46", operations.PreviewContextInterfaceSkillRequestBody{
        Source: components.CreateSkillSourcePayloadRaw(
            components.RawSkillSourcePayload{
                Type: components.RawSkillSourcePayloadTypeRaw,
                Config: components.RawSkillSourceConfigPayload{
                    Contents: "---\nname: pdf-processing\ndescription: Extract and summarize content from PDF documents\n---\n# PDF Processing\n...",
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SkillPreview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `interfaceID`                                                                                                            | `string`                                                                                                                 | :heavy_check_mark:                                                                                                       | The ID of the Context Interface.                                                                                         |
| `requestBody`                                                                                                            | [operations.PreviewContextInterfaceSkillRequestBody](../../models/operations/previewcontextinterfaceskillrequestbody.md) | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.PreviewContextInterfaceSkillResponse](../../models/operations/previewcontextinterfaceskillresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |