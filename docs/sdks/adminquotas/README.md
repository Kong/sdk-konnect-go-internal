# AdminQuotas

## Overview

Quotas cap the number of admin-managed resources of a given type that a Konnect organization is allowed to create. Use this endpoint to create or update the configured limit for a quota type.

### Available Operations

* [UpdateSliverOrgQuota](#updatesliverorgquota) - Create or update a quota
* [PatchSliverOrgQuota](#patchsliverorgquota) - Partially update a quota
* [ListSliverOrgQuotas](#listsliverorgquotas) - List quotas
* [UpdateIdentityDirectoryQuota](#updateidentitydirectoryquota) - Create or update a directory-scoped quota
* [PatchIdentityDirectoryQuota](#patchidentitydirectoryquota) - Partially update a directory-scoped quota
* [DeleteIdentityDirectoryQuota](#deleteidentitydirectoryquota) - Delete a directory-scoped quota override
* [ListIdentityDirectoryQuotas](#listidentitydirectoryquotas) - List directory-scoped quotas
* [UpdateIdentityQuota](#updateidentityquota) - Create or update a quota
* [PatchIdentityQuota](#patchidentityquota) - Partially update a quota
* [DeleteIdentityQuota](#deleteidentityquota) - Delete a quota override
* [ListIdentityQuotas](#listidentityquotas) - List quotas

## UpdateSliverOrgQuota

Creates or updates the quota limit for the given quota type, scoped to either an organization or a specific directory within it.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateSliverOrgQuota" method="put" path="/v1/quotas/{quotaType}" -->
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

    res, err := s.AdminQuotas.UpdateSliverOrgQuota(ctx, components.SliverQuotaTypeCimdClients, components.QuotaUpdateBody{
        QuotaValue: 50,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Quota != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `quotaType`                                                              | [components.SliverQuotaType](../../models/components/sliverquotatype.md) | :heavy_check_mark:                                                       | The type of resource the quota applies to                                |
| `quotaUpdateBody`                                                        | [components.QuotaUpdateBody](../../models/components/quotaupdatebody.md) | :heavy_check_mark:                                                       | Quota configuration to create or update                                  |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.UpdateSliverOrgQuotaResponse](../../models/operations/updatesliverorgquotaresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchSliverOrgQuota

Partially updates the quota limit for the given quota type. Unlike PUT, this never creates a quota override — it 404s if none exists yet for this quota type.

### Example Usage

<!-- UsageSnippet language="go" operationID="patchSliverOrgQuota" method="patch" path="/v1/quotas/{quotaType}" -->
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

    res, err := s.AdminQuotas.PatchSliverOrgQuota(ctx, components.SliverQuotaTypeCimdClients, components.QuotaPatchBody{
        QuotaValue: sdkkonnectgo.Pointer[int64](50),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Quota != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `quotaType`                                                              | [components.SliverQuotaType](../../models/components/sliverquotatype.md) | :heavy_check_mark:                                                       | The type of resource the quota applies to                                |
| `quotaPatchBody`                                                         | [components.QuotaPatchBody](../../models/components/quotapatchbody.md)   | :heavy_check_mark:                                                       | Quota fields to partially update                                         |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.PatchSliverOrgQuotaResponse](../../models/operations/patchsliverorgquotaresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListSliverOrgQuotas

Lists the platform default and any organization-level override for every quota type this API manages.

### Example Usage

<!-- UsageSnippet language="go" operationID="listSliverOrgQuotas" method="get" path="/v1/quotas" -->
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

    res, err := s.AdminQuotas.ListSliverOrgQuotas(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.QuotaList != nil {
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

**[*operations.ListSliverOrgQuotasResponse](../../models/operations/listsliverorgquotasresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateIdentityDirectoryQuota

Creates or updates the quota limit for the given quota type, scoped to this directory.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateIdentityDirectoryQuota" method="put" path="/v2/directories/{directoryId}/quotas/{quotaType}" -->
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

    res, err := s.AdminQuotas.UpdateIdentityDirectoryQuota(ctx, operations.UpdateIdentityDirectoryQuotaRequest{
        DirectoryID: "5581988f-6201-4d29-b382-8e8a9e76b092",
        QuotaType: components.IdentityDirectoryQuotaTypeKeyAuths,
        QuotaUpdateBody: components.QuotaUpdateBody{
            QuotaValue: 50,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Quota != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.UpdateIdentityDirectoryQuotaRequest](../../models/operations/updateidentitydirectoryquotarequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.UpdateIdentityDirectoryQuotaResponse](../../models/operations/updateidentitydirectoryquotaresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchIdentityDirectoryQuota

Partially updates the quota limit for the given quota type, scoped to this directory.
Unlike PUT, this never creates a quota override — it 404s if none exists yet at this scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="patchIdentityDirectoryQuota" method="patch" path="/v2/directories/{directoryId}/quotas/{quotaType}" -->
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

    res, err := s.AdminQuotas.PatchIdentityDirectoryQuota(ctx, operations.PatchIdentityDirectoryQuotaRequest{
        DirectoryID: "2523fda6-340f-4fc6-81ea-5bf16eeb636f",
        QuotaType: components.IdentityDirectoryQuotaTypePrincipals,
        QuotaPatchBody: components.QuotaPatchBody{
            QuotaValue: sdkkonnectgo.Pointer[int64](50),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Quota != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PatchIdentityDirectoryQuotaRequest](../../models/operations/patchidentitydirectoryquotarequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.PatchIdentityDirectoryQuotaResponse](../../models/operations/patchidentitydirectoryquotaresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteIdentityDirectoryQuota

Deletes the directory-scoped quota override for the given quota type. This does not disable
enforcement — quota checks fall back to the org override (or platform default) once the
directory override is removed.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteIdentityDirectoryQuota" method="delete" path="/v2/directories/{directoryId}/quotas/{quotaType}" -->
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

    res, err := s.AdminQuotas.DeleteIdentityDirectoryQuota(ctx, "1a24e485-7115-4112-9da1-d42e3c98a0df", components.IdentityDirectoryQuotaTypeBasicAuthPasswords)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `directoryID`                                                                                  | `string`                                                                                       | :heavy_check_mark:                                                                             | ID of the directory                                                                            |
| `quotaType`                                                                                    | [components.IdentityDirectoryQuotaType](../../models/components/identitydirectoryquotatype.md) | :heavy_check_mark:                                                                             | The type of directory-scoped resource the quota applies to                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.DeleteIdentityDirectoryQuotaResponse](../../models/operations/deleteidentitydirectoryquotaresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListIdentityDirectoryQuotas

Lists every directory-level quota override set for this directory. This does not include
the organization-level override or platform default that apply when no directory override
is set.


### Example Usage

<!-- UsageSnippet language="go" operationID="listIdentityDirectoryQuotas" method="get" path="/v2/directories/{directoryId}/quotas" -->
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

    res, err := s.AdminQuotas.ListIdentityDirectoryQuotas(ctx, "92c5be2c-2372-4248-a49a-6d53b84297ce")
    if err != nil {
        log.Fatal(err)
    }
    if res.QuotaList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `directoryID`                                            | `string`                                                 | :heavy_check_mark:                                       | ID of the directory                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListIdentityDirectoryQuotasResponse](../../models/operations/listidentitydirectoryquotasresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateIdentityQuota

Creates or updates the organization-scoped quota limit for the given quota type.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateIdentityQuota" method="put" path="/v2/quotas/{quotaType}" -->
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

    res, err := s.AdminQuotas.UpdateIdentityQuota(ctx, components.IdentityQuotaTypeBasicAuthPasswords, components.QuotaUpdateBody{
        QuotaValue: 50,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Quota != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `quotaType`                                                                  | [components.IdentityQuotaType](../../models/components/identityquotatype.md) | :heavy_check_mark:                                                           | The type of resource the quota applies to                                    |
| `quotaUpdateBody`                                                            | [components.QuotaUpdateBody](../../models/components/quotaupdatebody.md)     | :heavy_check_mark:                                                           | Quota configuration to create or update                                      |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.UpdateIdentityQuotaResponse](../../models/operations/updateidentityquotaresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## PatchIdentityQuota

Partially updates the organization-scoped quota limit for the given quota type. Unlike PUT,
this never creates a quota override — it 404s if none exists yet for this quota type.


### Example Usage

<!-- UsageSnippet language="go" operationID="patchIdentityQuota" method="patch" path="/v2/quotas/{quotaType}" -->
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

    res, err := s.AdminQuotas.PatchIdentityQuota(ctx, components.IdentityQuotaTypeIdentities, components.QuotaPatchBody{
        QuotaValue: sdkkonnectgo.Pointer[int64](50),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Quota != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `quotaType`                                                                  | [components.IdentityQuotaType](../../models/components/identityquotatype.md) | :heavy_check_mark:                                                           | The type of resource the quota applies to                                    |
| `quotaPatchBody`                                                             | [components.QuotaPatchBody](../../models/components/quotapatchbody.md)       | :heavy_check_mark:                                                           | Quota fields to partially update                                             |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.PatchIdentityQuotaResponse](../../models/operations/patchidentityquotaresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteIdentityQuota

Deletes the organization-scoped quota override for the given quota type. This does not
disable enforcement — quota checks fall back to the platform default (or a directory-scoped
override, where applicable) once the org override is removed.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteIdentityQuota" method="delete" path="/v2/quotas/{quotaType}" -->
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

    res, err := s.AdminQuotas.DeleteIdentityQuota(ctx, components.IdentityQuotaTypeIdentities)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `quotaType`                                                                  | [components.IdentityQuotaType](../../models/components/identityquotatype.md) | :heavy_check_mark:                                                           | The type of resource the quota applies to                                    |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.DeleteIdentityQuotaResponse](../../models/operations/deleteidentityquotaresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListIdentityQuotas

Lists the platform default and any organization-level override for every quota type this
API manages.


### Example Usage

<!-- UsageSnippet language="go" operationID="listIdentityQuotas" method="get" path="/v2/quotas" -->
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

    res, err := s.AdminQuotas.ListIdentityQuotas(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.QuotaList != nil {
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

**[*operations.ListIdentityQuotasResponse](../../models/operations/listidentityquotasresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |