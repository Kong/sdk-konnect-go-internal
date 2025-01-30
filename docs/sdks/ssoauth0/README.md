# SSOAuth0
(*SSOAuth0*)

## Overview

### Available Operations

* [PostAuth0RegisterInternal](#postauth0registerinternal) - Register a new organization with an Auth0 authenticated user.
* [GetAuth0OrganizationsInternal](#getauth0organizationsinternal) - Retrieves the list of organizations available to the given user.

## PostAuth0RegisterInternal

Create a new organization with Auth0 authentication. The Auth0 user will be the organization owner.

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

    res, err := s.SSOAuth0.PostAuth0RegisterInternal(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Auth0Registration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.Auth0Registration](../../models/components/auth0registration.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.PostAuth0RegisterInternalResponse](../../models/operations/postauth0registerinternalresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetAuth0OrganizationsInternal

Retrieves the list of organizations available to the given user.

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

    res, err := s.SSOAuth0.GetAuth0OrganizationsInternal(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Auth0OrganizationsCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAuth0OrganizationsInternalResponse](../../models/operations/getauth0organizationsinternalresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |