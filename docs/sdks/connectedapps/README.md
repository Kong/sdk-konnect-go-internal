# ConnectedApps

## Overview

Connected Apps enable third-party applications to access Konnect APIs on behalf of an organization using OAuth 2.0. When enabled, a reserved auth server is created with predefined scopes.

### Available Operations

* [GetConnectedApps](#getconnectedapps) - Get Connected Apps configuration
* [UpdateConnectedApps](#updateconnectedapps) - Enable or disable Connected Apps
* [GetConnectedAppsJwks](#getconnectedappsjwks) - Get Connected Apps JWKS
* [PatchKonnectConnectedApps](#patchkonnectconnectedapps) - Patch Konnect Connected Apps configuration
* [ListKonnectConnectedAppSessionsAdmin](#listkonnectconnectedappsessionsadmin) - List Konnect Connected App sessions
* [RevokeKonnectConnectedAppSessionAdmin](#revokekonnectconnectedappsessionadmin) - Revoke a Konnect Connected App session
* [PatchPortalConnectedApps](#patchportalconnectedapps) - Patch Portal Connected Apps configuration
* [ListPortalConnectedAppSessionsAdmin](#listportalconnectedappsessionsadmin) - List Portal Connected App sessions
* [RevokePortalConnectedAppSessionAdmin](#revokeportalconnectedappsessionadmin) - Revoke a Portal Connected App session
* [ListKonnectConnectedAppSessions](#listkonnectconnectedappsessions) - List caller's Konnect Connected App sessions
* [RevokeKonnectConnectedAppSession](#revokekonnectconnectedappsession) - Revoke one of caller's Konnect App sessions
* [ListPortalConnectedAppSessions](#listportalconnectedappsessions) - List caller's Portal Connected App sessions
* [RevokePortalConnectedAppSession](#revokeportalconnectedappsession) - Revoke one of caller's Portal App sessions

## GetConnectedApps

Returns the reserved auth server for Connected Apps if enabled. Requires domain parameter. When domain is 'portal', portal_id is also required.

### Example Usage

<!-- UsageSnippet language="go" operationID="getConnectedApps" method="get" path="/v1/connected-apps" -->
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

    res, err := s.ConnectedApps.GetConnectedApps(ctx, operations.DomainPortal, sdkkonnectgo.Pointer("d32d905a-ed33-46a3-a093-d8f536af9a8a"))
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectedAppsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `domain`                                                 | [operations.Domain](../../models/operations/domain.md)   | :heavy_check_mark:                                       | Domain for Connected Apps (konnect or portal)            |                                                          |
| `portalID`                                               | `*string`                                                | :heavy_minus_sign:                                       | The ID of the portal. Required when domain is 'portal'.  | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetConnectedAppsResponse](../../models/operations/getconnectedappsresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## UpdateConnectedApps

Enable or disable Connected Apps. When enabled, creates a reserved auth server with konnect:read and konnect:write scopes. When disabled, deletes the reserved auth server. This auth server is managed by Konnect and cannot be modified via standard auth server endpoints.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateConnectedApps" method="post" path="/v1/connected-apps" -->
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

    res, err := s.ConnectedApps.UpdateConnectedApps(ctx, components.CreateEnableConnectedAppsPayloadKonnect(
        components.EnableConnectedAppsKonnect{
            Enabled: false,
            Domain: components.EnableConnectedAppsKonnectDomainKonnect,
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectedAppsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.EnableConnectedAppsPayload](../../models/components/enableconnectedappspayload.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpdateConnectedAppsResponse](../../models/operations/updateconnectedappsresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.ConflictError   | 409                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## GetConnectedAppsJwks

Returns the JSON Web Key Set (JWKS) containing public keys for all configured Connected Apps domains (konnect and portal). Keys are used to verify JWT signatures issued by Connected Apps auth servers.

### Example Usage

<!-- UsageSnippet language="go" operationID="getConnectedAppsJwks" method="get" path="/v1/connected-apps/.well-known/jwks.json" -->
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

    res, err := s.ConnectedApps.GetConnectedAppsJwks(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.RsaJwks != nil {
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

**[*operations.GetConnectedAppsJwksResponse](../../models/operations/getconnectedappsjwksresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.NotFoundError  | 404                      | application/problem+json |
| sdkerrors.SDKError       | 4XX, 5XX                 | \*/\*                    |

## PatchKonnectConnectedApps

Partially updates the Konnect-domain Connected Apps configuration. Only fields present in the request body are changed. Setting `enabled` enables or disables Connected Apps (same effect as POST /connected-apps); omitting it patches dcr on the already-enabled auth server, which returns 404 if Connected Apps has not been enabled yet.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchKonnectConnectedApps" method="patch" path="/v1/connected-apps/konnect" -->
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

    res, err := s.ConnectedApps.PatchKonnectConnectedApps(ctx, components.PatchKonnectConnectedApps{
        Dcr: &components.PatchDcrSettings{
            RedirectURIAllowlist: []string{
                "https://example.com/callback",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectedAppsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.PatchKonnectConnectedApps](../../models/components/patchkonnectconnectedapps.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.PatchKonnectConnectedAppsResponse](../../models/operations/patchkonnectconnectedappsresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListKonnectConnectedAppSessionsAdmin

Returns all active Konnect Connected App sessions (active refresh tokens) in the caller's organization, across all users. The organization is taken from the caller's token. Optionally filter to a single user with filter[user_id].

### Example Usage

<!-- UsageSnippet language="go" operationID="listKonnectConnectedAppSessionsAdmin" method="get" path="/v1/connected-apps/konnect/sessions" -->
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

    res, err := s.ConnectedApps.ListKonnectConnectedAppSessionsAdmin(ctx, operations.ListKonnectConnectedAppSessionsAdminRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListKonnectConnectedAppSessionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.ListKonnectConnectedAppSessionsAdminRequest](../../models/operations/listkonnectconnectedappsessionsadminrequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.ListKonnectConnectedAppSessionsAdminResponse](../../models/operations/listkonnectconnectedappsessionsadminresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## RevokeKonnectConnectedAppSessionAdmin

Revokes a single active Konnect Connected App session (refresh token) in the caller's organization, regardless of which user owns it. The organization is taken from the caller's token.

### Example Usage

<!-- UsageSnippet language="go" operationID="revokeKonnectConnectedAppSessionAdmin" method="delete" path="/v1/connected-apps/konnect/sessions/{sessionId}" -->
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

    res, err := s.ConnectedApps.RevokeKonnectConnectedAppSessionAdmin(ctx, "3f1e5c2a-7b8d-4e6f-9a0b-1c2d3e4f5a6b")
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
| `sessionID`                                              | `string`                                                 | :heavy_check_mark:                                       | The session ID                                           | 3f1e5c2a-7b8d-4e6f-9a0b-1c2d3e4f5a6b                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RevokeKonnectConnectedAppSessionAdminResponse](../../models/operations/revokekonnectconnectedappsessionadminresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## PatchPortalConnectedApps

Partially updates the given portal's Connected Apps configuration. Only fields present in the request body are changed. Setting `enabled` enables or disables Connected Apps for this portal (same effect as POST /connected-apps; enabling for the first time requires default_url to also be set); omitting it patches dcr and/or default_url on the already-enabled auth server, which returns 404 if Connected Apps has not been enabled yet for this portal. The portal id is immutable and always taken from the path.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchPortalConnectedApps" method="patch" path="/v1/connected-apps/portals/{portalId}" -->
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

    res, err := s.ConnectedApps.PatchPortalConnectedApps(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", components.PatchPortalConnectedApps{
        Dcr: &components.PatchDcrSettings{
            RedirectURIAllowlist: []string{
                "https://example.com/callback",
            },
        },
        DefaultURL: sdkkonnectgo.Pointer("https://portal.example.com"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConnectedAppsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `portalID`                                                                                 | `string`                                                                                   | :heavy_check_mark:                                                                         | The portal ID                                                                              | d32d905a-ed33-46a3-a093-d8f536af9a8a                                                       |
| `patchPortalConnectedApps`                                                                 | [components.PatchPortalConnectedApps](../../models/components/patchportalconnectedapps.md) | :heavy_check_mark:                                                                         | Portal Connected Apps configuration fields to patch                                        |                                                                                            |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |                                                                                            |

### Response

**[*operations.PatchPortalConnectedAppsResponse](../../models/operations/patchportalconnectedappsresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListPortalConnectedAppSessionsAdmin

Returns all active Portal Connected App sessions (active refresh tokens) for the given portal, across all developers. Optionally filter to a single developer with filter[developer_id].

### Example Usage

<!-- UsageSnippet language="go" operationID="listPortalConnectedAppSessionsAdmin" method="get" path="/v1/connected-apps/portals/{portalId}/sessions" -->
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

    res, err := s.ConnectedApps.ListPortalConnectedAppSessionsAdmin(ctx, operations.ListPortalConnectedAppSessionsAdminRequest{
        PortalID: "d32d905a-ed33-46a3-a093-d8f536af9a8a",
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListPortalConnectedAppSessionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.ListPortalConnectedAppSessionsAdminRequest](../../models/operations/listportalconnectedappsessionsadminrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |
| `opts`                                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                                       | :heavy_minus_sign:                                                                                                             | The options for this request.                                                                                                  |

### Response

**[*operations.ListPortalConnectedAppSessionsAdminResponse](../../models/operations/listportalconnectedappsessionsadminresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## RevokePortalConnectedAppSessionAdmin

Revokes a single active Portal Connected App session (refresh token) for the given portal, regardless of which developer owns it.

### Example Usage

<!-- UsageSnippet language="go" operationID="revokePortalConnectedAppSessionAdmin" method="delete" path="/v1/connected-apps/portals/{portalId}/sessions/{sessionId}" -->
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

    res, err := s.ConnectedApps.RevokePortalConnectedAppSessionAdmin(ctx, "d32d905a-ed33-46a3-a093-d8f536af9a8a", "3f1e5c2a-7b8d-4e6f-9a0b-1c2d3e4f5a6b")
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
| `portalID`                                               | `string`                                                 | :heavy_check_mark:                                       | The portal ID                                            | d32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `sessionID`                                              | `string`                                                 | :heavy_check_mark:                                       | The session ID                                           | 3f1e5c2a-7b8d-4e6f-9a0b-1c2d3e4f5a6b                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RevokePortalConnectedAppSessionAdminResponse](../../models/operations/revokeportalconnectedappsessionadminresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.NotFoundError   | 404                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |

## ListKonnectConnectedAppSessions

Returns the authenticated caller's active Konnect Connected App sessions (active refresh tokens). Restricted to user principals. Authenticated but not authorized.

### Example Usage

<!-- UsageSnippet language="go" operationID="listKonnectConnectedAppSessions" method="get" path="/v1/connected-apps/konnect/me/sessions" -->
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

    res, err := s.ConnectedApps.ListKonnectConnectedAppSessions(ctx, operations.ListKonnectConnectedAppSessionsRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListConnectedAppSessionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.ListKonnectConnectedAppSessionsRequest](../../models/operations/listkonnectconnectedappsessionsrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.ListKonnectConnectedAppSessionsResponse](../../models/operations/listkonnectconnectedappsessionsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## RevokeKonnectConnectedAppSession

Revokes a single active Konnect Connected App session (refresh token) owned by the authenticated caller. Restricted to user principals. Authenticated but not authorized.

### Example Usage

<!-- UsageSnippet language="go" operationID="revokeKonnectConnectedAppSession" method="delete" path="/v1/connected-apps/konnect/me/sessions/{sessionId}" -->
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

    res, err := s.ConnectedApps.RevokeKonnectConnectedAppSession(ctx, "3f1e5c2a-7b8d-4e6f-9a0b-1c2d3e4f5a6b")
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
| `sessionID`                                              | `string`                                                 | :heavy_check_mark:                                       | The session ID                                           | 3f1e5c2a-7b8d-4e6f-9a0b-1c2d3e4f5a6b                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RevokeKonnectConnectedAppSessionResponse](../../models/operations/revokekonnectconnectedappsessionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListPortalConnectedAppSessions

Returns the authenticated caller's active Portal Connected App sessions (active refresh tokens). Restricted to developer principals carrying a portal_id claim. Authenticated but not authorized.

### Example Usage

<!-- UsageSnippet language="go" operationID="listPortalConnectedAppSessions" method="get" path="/v1/connected-apps/portal/me/sessions" -->
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

    res, err := s.ConnectedApps.ListPortalConnectedAppSessions(ctx, operations.ListPortalConnectedAppSessionsRequest{
        PageSize: sdkkonnectgo.Pointer[int64](10),
        PageNumber: sdkkonnectgo.Pointer[int64](1),
        Sort: sdkkonnectgo.Pointer("created_at desc"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ListConnectedAppSessionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.ListPortalConnectedAppSessionsRequest](../../models/operations/listportalconnectedappsessionsrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.ListPortalConnectedAppSessionsResponse](../../models/operations/listportalconnectedappsessionsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## RevokePortalConnectedAppSession

Revokes a single active Portal Connected App session (refresh token) owned by the authenticated caller. Restricted to developer principals carrying a portal_id claim. Authenticated but not authorized.

### Example Usage

<!-- UsageSnippet language="go" operationID="revokePortalConnectedAppSession" method="delete" path="/v1/connected-apps/portal/me/sessions/{sessionId}" -->
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

    res, err := s.ConnectedApps.RevokePortalConnectedAppSession(ctx, "3f1e5c2a-7b8d-4e6f-9a0b-1c2d3e4f5a6b")
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
| `sessionID`                                              | `string`                                                 | :heavy_check_mark:                                       | The session ID                                           | 3f1e5c2a-7b8d-4e6f-9a0b-1c2d3e4f5a6b                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RevokePortalConnectedAppSessionResponse](../../models/operations/revokeportalconnectedappsessionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |