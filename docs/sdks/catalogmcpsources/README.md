# CatalogMCPSources

## Overview

Manage what sources and generates MCP versions.

### Available Operations

* [GetMcpSource](#getmcpsource) - Get MCP source
* [CreateMcpSource](#createmcpsource) - Create MCP source
* [ReplaceMcpSource](#replacemcpsource) - Replace MCP source
* [UpdateMcpSource](#updatemcpsource) - Update MCP source
* [DeleteMcpSource](#deletemcpsource) - Delete MCP source
* [GetMcpSourcePreview](#getmcpsourcepreview) - Get MCP source preview

## GetMcpSource

Returns the MCP's source information.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-mcp-source" method="get" path="/v1/mcp-servers/{mcpId}/source" -->
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

    res, err := s.CatalogMCPSources.GetMcpSource(ctx, "a0119846-f179-4d9f-a168-d701facce7fb")
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogMCPSource != nil {
        switch res.CatalogMCPSource.Type {
            case components.CatalogMCPSourceTypeCatalogMCPGatewaySource:
                // res.CatalogMCPSource.CatalogMCPGatewaySource is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `mcpID`                                                  | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the MCP.                        | a0119846-f179-4d9f-a168-d701facce7fb                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetMcpSourceResponse](../../models/operations/getmcpsourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateMcpSource

Creates the MCP's source information.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-mcp-source" method="post" path="/v1/mcp-servers/{mcpId}/source" -->
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

    res, err := s.CatalogMCPSources.CreateMcpSource(ctx, "a0119846-f179-4d9f-a168-d701facce7fb", components.CreateCatalogMCPSource{
        Source: components.CreateCreateCatalogMCPSourceSourceCatalogMCPGatewaySourceBlock(
            components.CatalogMCPGatewaySourceBlock{
                Type: components.CatalogMCPGatewaySourceBlockTypeAiGateway,
                Config: components.CatalogMCPGatewaySourceConfig{
                    GatewayControlPlaneID: "223e4567-e89b-12d3-a456-426614174999",
                    GatewayMcpServerID: "79087145-0159-4811-9042-6a9365f234bc",
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogMCPSource != nil {
        switch res.CatalogMCPSource.Type {
            case components.CatalogMCPSourceTypeCatalogMCPGatewaySource:
                // res.CatalogMCPSource.CatalogMCPGatewaySource is populated
        }

    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `mcpID`                                                                                | `string`                                                                               | :heavy_check_mark:                                                                     | The unique identifier of the MCP.                                                      | a0119846-f179-4d9f-a168-d701facce7fb                                                   |
| `createCatalogMCPSource`                                                               | [components.CreateCatalogMCPSource](../../models/components/createcatalogmcpsource.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.CreateMcpSourceResponse](../../models/operations/createmcpsourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ReplaceMcpSource

Replaces the MCP's source information.


### Example Usage

<!-- UsageSnippet language="go" operationID="replace-mcp-source" method="put" path="/v1/mcp-servers/{mcpId}/source" -->
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

    res, err := s.CatalogMCPSources.ReplaceMcpSource(ctx, "a0119846-f179-4d9f-a168-d701facce7fb", components.CreateCatalogMCPSource{
        Source: components.CreateCreateCatalogMCPSourceSourceCatalogMCPGatewaySourceBlock(
            components.CatalogMCPGatewaySourceBlock{
                Type: components.CatalogMCPGatewaySourceBlockTypeAiGateway,
                Config: components.CatalogMCPGatewaySourceConfig{
                    GatewayControlPlaneID: "223e4567-e89b-12d3-a456-426614174999",
                    GatewayMcpServerID: "79087145-0159-4811-9042-6a9365f234bc",
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogMCPSource != nil {
        switch res.CatalogMCPSource.Type {
            case components.CatalogMCPSourceTypeCatalogMCPGatewaySource:
                // res.CatalogMCPSource.CatalogMCPGatewaySource is populated
        }

    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `mcpID`                                                                                | `string`                                                                               | :heavy_check_mark:                                                                     | The unique identifier of the MCP.                                                      | a0119846-f179-4d9f-a168-d701facce7fb                                                   |
| `createCatalogMCPSource`                                                               | [components.CreateCatalogMCPSource](../../models/components/createcatalogmcpsource.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.ReplaceMcpSourceResponse](../../models/operations/replacemcpsourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateMcpSource

Partially updates the MCP's source information.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-mcp-source" method="patch" path="/v1/mcp-servers/{mcpId}/source" -->
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

    res, err := s.CatalogMCPSources.UpdateMcpSource(ctx, "a0119846-f179-4d9f-a168-d701facce7fb", components.UpdateCatalogMCPSource{
        Source: sdkkonnectgo.Pointer(components.CreateUpdateCatalogMCPSourceSourceCatalogMCPGatewaySourceBlock(
            components.CatalogMCPGatewaySourceBlock{
                Type: components.CatalogMCPGatewaySourceBlockTypeAiGateway,
                Config: components.CatalogMCPGatewaySourceConfig{
                    GatewayControlPlaneID: "223e4567-e89b-12d3-a456-426614174999",
                    GatewayMcpServerID: "79087145-0159-4811-9042-6a9365f234bc",
                },
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogMCPSource != nil {
        switch res.CatalogMCPSource.Type {
            case components.CatalogMCPSourceTypeCatalogMCPGatewaySource:
                // res.CatalogMCPSource.CatalogMCPGatewaySource is populated
        }

    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `mcpID`                                                                                | `string`                                                                               | :heavy_check_mark:                                                                     | The unique identifier of the MCP.                                                      | a0119846-f179-4d9f-a168-d701facce7fb                                                   |
| `updateCatalogMCPSource`                                                               | [components.UpdateCatalogMCPSource](../../models/components/updatecatalogmcpsource.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.UpdateMcpSourceResponse](../../models/operations/updatemcpsourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteMcpSource

Deletes the MCP's source information.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-mcp-source" method="delete" path="/v1/mcp-servers/{mcpId}/source" -->
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

    res, err := s.CatalogMCPSources.DeleteMcpSource(ctx, "a0119846-f179-4d9f-a168-d701facce7fb")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `mcpID`                                                  | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the MCP.                        | a0119846-f179-4d9f-a168-d701facce7fb                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteMcpSourceResponse](../../models/operations/deletemcpsourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetMcpSourcePreview

Returns the current configuration retrieved from the MCP's source.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-mcp-source-preview" method="get" path="/v1/mcp-servers/{mcpId}/source/preview" -->
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

    res, err := s.CatalogMCPSources.GetMcpSourcePreview(ctx, "a0119846-f179-4d9f-a168-d701facce7fb")
    if err != nil {
        log.Fatal(err)
    }
    if res.CatalogMCPSourcePreview != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `mcpID`                                                  | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the MCP.                        | a0119846-f179-4d9f-a168-d701facce7fb                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetMcpSourcePreviewResponse](../../models/operations/getmcpsourcepreviewresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |