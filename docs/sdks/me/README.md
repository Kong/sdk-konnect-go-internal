# Me
(*Me*)

## Overview

### Available Operations

* [GetUsersMe](#getusersme) - Retrieve My User Account
* [DeleteUsersMe](#deleteusersme) - Delete My User Account
* [PatchUsersMe](#patchusersme) - Update My User Account
* [GetUsersMePermissions](#getusersmepermissions) - Get My Permissions
* [GetOrganizationsMe](#getorganizationsme) - Retrieve My Organization
* [UpdateOrganizationsMe](#updateorganizationsme) - Update My Organization

## GetUsersMe

Returns the user account for the user identified in the token of the request.

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

    res, err := s.Me.GetUsersMe(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
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

**[*operations.GetUsersMeResponse](../../models/operations/getusersmeresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteUsersMe

Deletes the user account for the user identified in the token of the request.

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

    res, err := s.Me.DeleteUsersMe(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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

**[*operations.DeleteUsersMeResponse](../../models/operations/deleteusersmeresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchUsersMe

Updates the user account for the user identified in the token of the request.

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

    res, err := s.Me.PatchUsersMe(ctx, &components.UpdateUser{
        FullName: sdkkonnectgointernal.String("James C. Woods"),
        PreferredName: sdkkonnectgointernal.String("Jimmy"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [components.UpdateUser](../../models/components/updateuser.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.PatchUsersMeResponse](../../models/operations/patchusersmeresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetUsersMePermissions

Returns the permissions for the current user

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

    res, err := s.Me.GetUsersMePermissions(ctx, &operations.GetUsersMePermissionsQueryParamFilter{
        TopLevel: sdkkonnectgointernal.Bool(true),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserPermissions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                             | Type                                                                                                                  | Required                                                                                                              | Description                                                                                                           |
| --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                                                 | :heavy_check_mark:                                                                                                    | The context to use for the request.                                                                                   |
| `filter`                                                                                                              | [*operations.GetUsersMePermissionsQueryParamFilter](../../models/operations/getusersmepermissionsqueryparamfilter.md) | :heavy_minus_sign:                                                                                                    | Filter permissions returned in the response.                                                                          |
| `opts`                                                                                                                | [][operations.Option](../../models/operations/option.md)                                                              | :heavy_minus_sign:                                                                                                    | The options for this request.                                                                                         |

### Response

**[*operations.GetUsersMePermissionsResponse](../../models/operations/getusersmepermissionsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetOrganizationsMe

Returns the organization of the user identified in the token of the request.

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

    res, err := s.Me.GetOrganizationsMe(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.MeOrganization != nil {
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

**[*operations.GetOrganizationsMeResponse](../../models/operations/getorganizationsmeresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateOrganizationsMe

Updates the current user's organization. When updating the owner, the new owner 
must be an organization admin.


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

    res, err := s.Me.UpdateOrganizationsMe(ctx, &components.UpdateOrganization{
        OwnerID: sdkkonnectgointernal.String("d9c56d92-0339-4dfb-be0c-2253162edba9"),
        Name: sdkkonnectgointernal.String("Kong Inc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.MeOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.UpdateOrganization](../../models/components/updateorganization.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.UpdateOrganizationsMeResponse](../../models/operations/updateorganizationsmeresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |