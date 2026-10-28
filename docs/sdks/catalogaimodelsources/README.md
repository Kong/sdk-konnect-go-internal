# CatalogAIModelSources

## Overview

Manage what sources and generates AI Model versions.

### Available Operations

* [GetAiModelSource](#getaimodelsource) - Get AI Model source
* [CreateAiModelSource](#createaimodelsource) - Create AI Model source
* [ReplaceAiModelSource](#replaceaimodelsource) - Replace AI Model source
* [UpdateAiModelSource](#updateaimodelsource) - Update AI Model source
* [DeleteAiModelSource](#deleteaimodelsource) - Delete AI Model source
* [GetAiModelSourcePreview](#getaimodelsourcepreview) - Get AI Model source preview

## GetAiModelSource

Returns the AI Model's source information.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-model-source" method="get" path="/v1/ai-models/{aiModelId}/source" -->
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

    res, err := s.CatalogAIModelSources.GetAiModelSource(ctx, "123e4567-e89b-12d3-a456-426614174000")
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModelSource != nil {
        switch res.AiModelSource.Type {
            case components.AiModelSourceTypeAiModelGatewaySource:
                // res.AiModelSource.AiModelGatewaySource is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `aiModelID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the AI Model.                   | 123e4567-e89b-12d3-a456-426614174000                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiModelSourceResponse](../../models/operations/getaimodelsourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateAiModelSource

Creates the AI Model's source information.


### Example Usage

<!-- UsageSnippet language="go" operationID="create-ai-model-source" method="post" path="/v1/ai-models/{aiModelId}/source" -->
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

    res, err := s.CatalogAIModelSources.CreateAiModelSource(ctx, "123e4567-e89b-12d3-a456-426614174000", components.CreateAiModelSource{
        Source: components.CreateCreateAiModelSourceSourceAiModelGatewaySourceBlock(
            components.AiModelGatewaySourceBlock{
                Type: components.AiModelGatewaySourceBlockTypeAiGateway,
                Config: components.AiModelGatewaySourceConfig{
                    GatewayControlPlaneID: "223e4567-e89b-12d3-a456-426614174999",
                    GatewayModelID: "2a0e84cc-b5bb-432a-bcc3-ab86463d57bb",
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModelSource != nil {
        switch res.AiModelSource.Type {
            case components.AiModelSourceTypeAiModelGatewaySource:
                // res.AiModelSource.AiModelGatewaySource is populated
        }

    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `aiModelID`                                                                      | `string`                                                                         | :heavy_check_mark:                                                               | The unique identifier of the AI Model.                                           | 123e4567-e89b-12d3-a456-426614174000                                             |
| `createAiModelSource`                                                            | [components.CreateAiModelSource](../../models/components/createaimodelsource.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.CreateAiModelSourceResponse](../../models/operations/createaimodelsourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ReplaceAiModelSource

Replaces the AI Model's source information.


### Example Usage

<!-- UsageSnippet language="go" operationID="replace-ai-model-source" method="put" path="/v1/ai-models/{aiModelId}/source" -->
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

    res, err := s.CatalogAIModelSources.ReplaceAiModelSource(ctx, "123e4567-e89b-12d3-a456-426614174000", components.CreateAiModelSource{
        Source: components.CreateCreateAiModelSourceSourceAiModelGatewaySourceBlock(
            components.AiModelGatewaySourceBlock{
                Type: components.AiModelGatewaySourceBlockTypeAiGateway,
                Config: components.AiModelGatewaySourceConfig{
                    GatewayControlPlaneID: "223e4567-e89b-12d3-a456-426614174999",
                    GatewayModelID: "2a0e84cc-b5bb-432a-bcc3-ab86463d57bb",
                },
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModelSource != nil {
        switch res.AiModelSource.Type {
            case components.AiModelSourceTypeAiModelGatewaySource:
                // res.AiModelSource.AiModelGatewaySource is populated
        }

    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `aiModelID`                                                                      | `string`                                                                         | :heavy_check_mark:                                                               | The unique identifier of the AI Model.                                           | 123e4567-e89b-12d3-a456-426614174000                                             |
| `createAiModelSource`                                                            | [components.CreateAiModelSource](../../models/components/createaimodelsource.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.ReplaceAiModelSourceResponse](../../models/operations/replaceaimodelsourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateAiModelSource

Partially updates an AI Model's source information.


### Example Usage

<!-- UsageSnippet language="go" operationID="update-ai-model-source" method="patch" path="/v1/ai-models/{aiModelId}/source" -->
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

    res, err := s.CatalogAIModelSources.UpdateAiModelSource(ctx, "123e4567-e89b-12d3-a456-426614174000", components.UpdateAiModelSource{
        Source: sdkkonnectgo.Pointer(components.CreateUpdateAiModelSourceSourceAiModelGatewaySourceBlock(
            components.AiModelGatewaySourceBlock{
                Type: components.AiModelGatewaySourceBlockTypeAiGateway,
                Config: components.AiModelGatewaySourceConfig{
                    GatewayControlPlaneID: "223e4567-e89b-12d3-a456-426614174999",
                    GatewayModelID: "2a0e84cc-b5bb-432a-bcc3-ab86463d57bb",
                },
            },
        )),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModelSource != nil {
        switch res.AiModelSource.Type {
            case components.AiModelSourceTypeAiModelGatewaySource:
                // res.AiModelSource.AiModelGatewaySource is populated
        }

    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `aiModelID`                                                                      | `string`                                                                         | :heavy_check_mark:                                                               | The unique identifier of the AI Model.                                           | 123e4567-e89b-12d3-a456-426614174000                                             |
| `updateAiModelSource`                                                            | [components.UpdateAiModelSource](../../models/components/updateaimodelsource.md) | :heavy_check_mark:                                                               | N/A                                                                              |                                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.UpdateAiModelSourceResponse](../../models/operations/updateaimodelsourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteAiModelSource

Deletes the AI Model's source information.


### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ai-model-source" method="delete" path="/v1/ai-models/{aiModelId}/source" -->
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

    res, err := s.CatalogAIModelSources.DeleteAiModelSource(ctx, "123e4567-e89b-12d3-a456-426614174000")
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
| `aiModelID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the AI Model.                   | 123e4567-e89b-12d3-a456-426614174000                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAiModelSourceResponse](../../models/operations/deleteaimodelsourceresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetAiModelSourcePreview

Returns the current configuration retrieved from the AI Model's source.


### Example Usage

<!-- UsageSnippet language="go" operationID="get-ai-model-source-preview" method="get" path="/v1/ai-models/{aiModelId}/source/preview" -->
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

    res, err := s.CatalogAIModelSources.GetAiModelSourcePreview(ctx, "123e4567-e89b-12d3-a456-426614174000")
    if err != nil {
        log.Fatal(err)
    }
    if res.AiModelVersionCreate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `aiModelID`                                              | `string`                                                 | :heavy_check_mark:                                       | The unique identifier of the AI Model.                   | 123e4567-e89b-12d3-a456-426614174000                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAiModelSourcePreviewResponse](../../models/operations/getaimodelsourcepreviewresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |