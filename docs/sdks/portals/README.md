# Portals
(*Portals*)

## Overview

APIs related to configuration of Konnect Developer Portals.

### Available Operations

* [ListPortals](#listportals) - List Portals
* [CreatePortal](#createportal) - Create Portal
* [GetPortal](#getportal) - Fetch Portal
* [UpdatePortal](#updateportal) - Update Portal
* [DeletePortal](#deleteportal) - Delete Portal

## ListPortals

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Lists developer portals defined in this region for this organization. Each developer portal is available at a unique address and has isolated configuration, customization, developers, and applications.

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

    res, err := s.Portals.ListPortals(ctx, operations.ListPortalsRequest{
        PageSize: sdkkonnectgointernal.Int64(10),
        PageNumber: sdkkonnectgointernal.Int64(1),
        Filter: &components.PortalFilterParameters{
            AuthenticationEnabled: sdkkonnectgointernal.Bool(true),
            RbacEnabled: sdkkonnectgointernal.Bool(true),
            AutoApproveDevelopers: sdkkonnectgointernal.Bool(true),
            AutoApproveApplications: sdkkonnectgointernal.Bool(true),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPortalsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListPortalsRequest](../../models/operations/listportalsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.ListPortalsResponse](../../models/operations/listportalsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreatePortal

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Creates a new developer portal scoped in this region for this organization.

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

    res, err := s.Portals.CreatePortal(ctx, components.CreatePortal{
        Name: "<value>",
        Labels: map[string]*string{
            "env": sdkkonnectgointernal.String("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [components.CreatePortal](../../models/components/createportal.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.CreatePortalResponse](../../models/operations/createportalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPortal

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns the configuration for a single developer portal, including the current visibility, access, and domain settings.

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

    res, err := s.Portals.GetPortal(ctx, "1672e313-9f8e-4485-a62b-716aad351889")
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `portalID`                                               | *string*                                                 | :heavy_check_mark:                                       | ID of the portal.                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPortalResponse](../../models/operations/getportalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePortal

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Updates the configuration for a single portal including the visibility, access, and custom domain settings.

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

    res, err := s.Portals.UpdatePortal(ctx, "efe853b6-a55c-4929-9b80-c39a22b01c7a", components.UpdatePortal{
        Labels: map[string]*string{
            "env": sdkkonnectgointernal.String("test"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `portalID`                                                         | *string*                                                           | :heavy_check_mark:                                                 | ID of the portal.                                                  |
| `updatePortal`                                                     | [components.UpdatePortal](../../models/components/updateportal.md) | :heavy_check_mark:                                                 | Update a portal's settings.                                        |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.UpdatePortalResponse](../../models/operations/updateportalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeletePortal

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Deletes a single portal, along with all related entities.

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

    res, err := s.Portals.DeletePortal(ctx, "304c8bf0-1c10-414d-af00-7cb8c3204da0", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                 | Type                                                                                                                                                                                      | Required                                                                                                                                                                                  | Description                                                                                                                                                                               |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                        | The context to use for the request.                                                                                                                                                       |
| `portalID`                                                                                                                                                                                | *string*                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                        | ID of the portal.                                                                                                                                                                         |
| `force`                                                                                                                                                                                   | [*operations.Force](../../models/operations/force.md)                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                        | If true, the portal will be deleted, automatically deleting all API publications. If the force param is not set, the deletion will only succeed if there are no APIs currently published. |
| `opts`                                                                                                                                                                                    | [][operations.Option](../../models/operations/option.md)                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                        | The options for this request.                                                                                                                                                             |

### Response

**[*operations.DeletePortalResponse](../../models/operations/deleteportalresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |