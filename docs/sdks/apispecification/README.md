# APISpecification
(*APISpecification*)

## Overview

### Available Operations

* [CreateAPISpec](#createapispec) - Create API Specification
* [ListAPISpecs](#listapispecs) - List API Specifications
* [FetchAPISpec](#fetchapispec) - Fetch API Specification
* [UpdateAPISpec](#updateapispec) - Update API Specification
* [DeleteAPISpec](#deleteapispec) - Delete API Specification

## CreateAPISpec

Creates a specification for an API.
**Note:** You can only have one specification for an API.


### Example Usage

```go
package main

import(
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

    res, err := s.APISpecification.CreateAPISpec(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", components.CreateAPISpecRequest{
        Content: "{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}",
        Type: components.CreateAPISpecRequestTypeOas3.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `apiID`                                                                            | *string*                                                                           | :heavy_check_mark:                                                                 | The UUID API identifier                                                            | 9f5061ce-78f6-4452-9108-ad7c02821fd5                                               |
| `createAPISpecRequest`                                                             | [components.CreateAPISpecRequest](../../models/components/createapispecrequest.md) | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.CreateAPISpecResponse](../../models/operations/createapispecresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.BadRequestError           | 400                                 | application/problem+json            |
| sdkerrors.UnauthorizedError         | 401                                 | application/problem+json            |
| sdkerrors.ForbiddenError            | 403                                 | application/problem+json            |
| sdkerrors.NotFoundError             | 404                                 | application/problem+json            |
| sdkerrors.ConflictError             | 409                                 | application/problem+json            |
| sdkerrors.UnsupportedMediaTypeError | 415                                 | application/problem+json            |
| sdkerrors.SDKError                  | 4XX, 5XX                            | \*/\*                               |

## ListAPISpecs

Returns a list of specifications for an API.
**Note:** You can only have one specification for an API.


### Example Usage

```go
package main

import(
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

    res, err := s.APISpecification.ListAPISpecs(ctx, operations.ListAPISpecsRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        PageSize: sdkkonnectgointernal.Int64(10),
        PageNumber: sdkkonnectgointernal.Int64(1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListAPISpecResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListAPISpecsRequest](../../models/operations/listapispecsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.ListAPISpecsResponse](../../models/operations/listapispecsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## FetchAPISpec

Fetches the specification of an API.

### Example Usage

```go
package main

import(
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

    res, err := s.APISpecification.FetchAPISpec(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "d32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `apiID`                                                  | *string*                                                 | :heavy_check_mark:                                       | The UUID API identifier                                  | 9f5061ce-78f6-4452-9108-ad7c02821fd5                     |
| `specID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The API specification identifier                         | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.FetchAPISpecResponse](../../models/operations/fetchapispecresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateAPISpec

Updates the specification of an API.

### Example Usage

```go
package main

import(
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

    res, err := s.APISpecification.UpdateAPISpec(ctx, operations.UpdateAPISpecRequest{
        APIID: "9f5061ce-78f6-4452-9108-ad7c02821fd5",
        SpecID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        APISpec: components.APISpec{
            Content: sdkkonnectgointernal.String("{\"openapi\":\"3.0.3\",\"info\":{\"title\":\"Example API\",\"version\":\"1.0.0\"},\"paths\":{\"/example\":{\"get\":{\"summary\":\"Example endpoint\",\"responses\":{\"200\":{\"description\":\"Successful response\"}}}}}}"),
            Type: components.APISpecTypeOas3.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.APISpecResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateAPISpecRequest](../../models/operations/updateapispecrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.UpdateAPISpecResponse](../../models/operations/updateapispecresponse.md), error**

### Errors

| Error Type                          | Status Code                         | Content Type                        |
| ----------------------------------- | ----------------------------------- | ----------------------------------- |
| sdkerrors.BadRequestError           | 400                                 | application/problem+json            |
| sdkerrors.UnauthorizedError         | 401                                 | application/problem+json            |
| sdkerrors.ForbiddenError            | 403                                 | application/problem+json            |
| sdkerrors.NotFoundError             | 404                                 | application/problem+json            |
| sdkerrors.ConflictError             | 409                                 | application/problem+json            |
| sdkerrors.UnsupportedMediaTypeError | 415                                 | application/problem+json            |
| sdkerrors.SDKError                  | 4XX, 5XX                            | \*/\*                               |

## DeleteAPISpec

Deletes the specification of an API.

### Example Usage

```go
package main

import(
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

    res, err := s.APISpecification.DeleteAPISpec(ctx, "9f5061ce-78f6-4452-9108-ad7c02821fd5", "d32d905a-ed33-46a3-a093-d8f536af9a8a")
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
| `apiID`                                                  | *string*                                                 | :heavy_check_mark:                                       | The UUID API identifier                                  | 9f5061ce-78f6-4452-9108-ad7c02821fd5                     |
| `specID`                                                 | *string*                                                 | :heavy_check_mark:                                       | The API specification identifier                         | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAPISpecResponse](../../models/operations/deleteapispecresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |