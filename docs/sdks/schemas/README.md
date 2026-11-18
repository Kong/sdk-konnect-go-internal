# Schemas

## Overview

### Available Operations

* [ValidateExpression](#validateexpression) - Validate a CEL expression
* [ValidateEntitySchema](#validateentityschema) - Validate entity schema
* [FetchPartialSchema](#fetchpartialschema) - Get partial schema

## ValidateExpression

Validates a CEL expression for the `expressions` field or the `condition` field of a plugin. Send exactly one of `expression` or `condition`.

### Example Usage

<!-- UsageSnippet language="go" operationID="validate-expression" method="post" path="/v2/control-planes/{controlPlaneId}/expressions/validate" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.Schemas.ValidateExpression(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.CreateValidateExpressionBodyRequestValidateExpressionRequest(
        components.ValidateExpressionRequest{
            Expression: "http.method == \"GET\"",
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidateEntityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `controlPlaneID`                                                                                     | `string`                                                                                             | :heavy_check_mark:                                                                                   | The UUID of your control plane. This variable is available in the Konnect manager.                   | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                                 |
| `validateExpressionBodyRequest`                                                                      | [components.ValidateExpressionBodyRequest](../../models/components/validateexpressionbodyrequest.md) | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.ValidateExpressionResponse](../../models/operations/validateexpressionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ValidateEntitySchema

Validate schema for an entity

### Example Usage

<!-- UsageSnippet language="go" operationID="validate-entity-schema" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/schemas/{entityName}/validate" -->
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

    res, err := s.Schemas.ValidateEntitySchema(ctx, operations.ValidateEntitySchemaRequest{
        EntityName: "<value>",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidateEntityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.ValidateEntitySchemaRequest](../../models/operations/validateentityschemarequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.ValidateEntitySchemaResponse](../../models/operations/validateentityschemaresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## FetchPartialSchema

Get the schema for a partial

### Example Usage

<!-- UsageSnippet language="go" operationID="fetch-partial-schema" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/schemas/partials/{partialType}" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.Pointer("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.Schemas.FetchPartialSchema(ctx, "<value>", "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetPartialSchemaResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `partialType`                                                                      | `string`                                                                           | :heavy_check_mark:                                                                 | The type of a partial                                                              |                                                                                    |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.FetchPartialSchemaResponse](../../models/operations/fetchpartialschemaresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |