# AuthServerTrustedIdPs

## Overview

Trusted IdPs allow an Auth Server to accept identity assertions from an external identity provider and exchange them for Konnect-issued tokens. The management API will give you the ability to configure the trusted IdP for an Auth Server.

### Available Operations

* [ListAuthServerTrustedIdps](#listauthservertrustedidps) - List all auth server trusted IdPs
* [CreateAuthServerTrustedIdp](#createauthservertrustedidp) - Create a new auth server trusted IdP
* [GetAuthServerTrustedIdp](#getauthservertrustedidp) - Get an auth server trusted IdP
* [UpdateAuthServerTrustedIdp](#updateauthservertrustedidp) - Update an auth server trusted IdP
* [DeleteAuthServerTrustedIdp](#deleteauthservertrustedidp) - Delete an auth server trusted IdP

## ListAuthServerTrustedIdps

List all trusted IdPs for a given auth server.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAuthServerTrustedIdps" method="get" path="/v1/auth-servers/{authServerId}/trusted-idps" -->
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

    res, err := s.AuthServerTrustedIDPs.ListAuthServerTrustedIdps(ctx, operations.ListAuthServerTrustedIdpsRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthServerTrustedIdpList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.ListAuthServerTrustedIdpsRequest](../../models/operations/listauthservertrustedidpsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.ListAuthServerTrustedIdpsResponse](../../models/operations/listauthservertrustedidpsresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## CreateAuthServerTrustedIdp

Create a trusted IdP for a given auth server. The `issuer_url` must be unique within the auth server. Only one trusted IdP per auth server is allowed.

### Example Usage

<!-- UsageSnippet language="go" operationID="createAuthServerTrustedIdp" method="post" path="/v1/auth-servers/{authServerId}/trusted-idps" -->
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

    res, err := s.AuthServerTrustedIDPs.CreateAuthServerTrustedIdp(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.CreateAuthServerTrustedIdpPayload{
        IssuerURL: "https://acme.okta.com/oauth2/default",
        DirectoryID: "2a0268d8-1432-4941-a3a1-5613ccb89441",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthServerTrustedIdp != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  | Example                                                                                                      |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |                                                                                                              |
| `authServerID`                                                                                               | `string`                                                                                                     | :heavy_check_mark:                                                                                           | The auth server ID                                                                                           | d32d905a-ed33-46a3-a093-d8f536af9a8a                                                                         |
| `createAuthServerTrustedIdpPayload`                                                                          | [components.CreateAuthServerTrustedIdpPayload](../../models/components/createauthservertrustedidppayload.md) | :heavy_check_mark:                                                                                           | Trusted IdP to be created                                                                                    |                                                                                                              |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |                                                                                                              |

### Response

**[*operations.CreateAuthServerTrustedIdpResponse](../../models/operations/createauthservertrustedidpresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.ConflictError   | 409                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetAuthServerTrustedIdp

Get a trusted IdP for a given auth server.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAuthServerTrustedIdp" method="get" path="/v1/auth-servers/{authServerId}/trusted-idps/{authServerTrustedIdpId}" -->
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

    res, err := s.AuthServerTrustedIDPs.GetAuthServerTrustedIdp(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", "7b5c1c2e-9f3a-4b6d-8e21-0a1b2c3d4e5f")
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthServerTrustedIdp != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `authServerID`                                           | `string`                                                 | :heavy_check_mark:                                       | The auth server ID                                       | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `authServerTrustedIdpID`                                 | `string`                                                 | :heavy_check_mark:                                       | The trusted IdP ID                                       | 7b5c1c2e-9f3a-4b6d-8e21-0a1b2c3d4e5f                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAuthServerTrustedIdpResponse](../../models/operations/getauthservertrustedidpresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## UpdateAuthServerTrustedIdp

Replace a trusted IdP for a given auth server. The request body is the full representation; omitted optional fields are reset to their defaults.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAuthServerTrustedIdp" method="put" path="/v1/auth-servers/{authServerId}/trusted-idps/{authServerTrustedIdpId}" -->
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

    res, err := s.AuthServerTrustedIDPs.UpdateAuthServerTrustedIdp(ctx, operations.UpdateAuthServerTrustedIdpRequest{
        AuthServerID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        AuthServerTrustedIdpID: "7b5c1c2e-9f3a-4b6d-8e21-0a1b2c3d4e5f",
        UpdateAuthServerTrustedIdpPayload: components.UpdateAuthServerTrustedIdpPayload{
            IssuerURL: "https://acme.okta.com/oauth2/default",
            DirectoryID: "36df4224-1e9e-4ded-8a38-93ea9f6a8778",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthServerTrustedIdp != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.UpdateAuthServerTrustedIdpRequest](../../models/operations/updateauthservertrustedidprequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.UpdateAuthServerTrustedIdpResponse](../../models/operations/updateauthservertrustedidpresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.ConflictError   | 409                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## DeleteAuthServerTrustedIdp

Delete a trusted IdP for a given auth server.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteAuthServerTrustedIdp" method="delete" path="/v1/auth-servers/{authServerId}/trusted-idps/{authServerTrustedIdpId}" -->
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

    res, err := s.AuthServerTrustedIDPs.DeleteAuthServerTrustedIdp(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", "7b5c1c2e-9f3a-4b6d-8e21-0a1b2c3d4e5f")
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
| `authServerID`                                           | `string`                                                 | :heavy_check_mark:                                       | The auth server ID                                       | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `authServerTrustedIdpID`                                 | `string`                                                 | :heavy_check_mark:                                       | The trusted IdP ID                                       | 7b5c1c2e-9f3a-4b6d-8e21-0a1b2c3d4e5f                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteAuthServerTrustedIdpResponse](../../models/operations/deleteauthservertrustedidpresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |