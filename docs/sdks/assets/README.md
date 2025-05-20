# Assets
(*Assets*)

## Overview

APIs for managing static assets for Konnect Developer Portals.

### Available Operations

* [GetPortalAssetLogo](#getportalassetlogo) - Get Logo
* [ReplacePortalAssetLogo](#replaceportalassetlogo) - Replace Logo
* [GetPortalAssetFavicon](#getportalassetfavicon) - Get Favicon
* [ReplacePortalAssetFavicon](#replaceportalassetfavicon) - Replace Favicon

## GetPortalAssetLogo

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns the logo of the portal.

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

    res, err := s.Assets.GetPortalAssetLogo(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredImagePngPortalImageAsset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `portalID`                                               | *string*                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPortalAssetLogoResponse](../../models/operations/getportalassetlogoresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ReplacePortalAssetLogo

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Replaces the logo of the portal.

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

    res, err := s.Assets.ReplacePortalAssetLogo(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.ReplacePortalImageAsset{
        Data: "data:image/jpeg;base64,bmljZV9sb29raW5nX3BpY3R1cmU=",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredImagePngPortalImageAsset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `portalID`                                                                               | *string*                                                                                 | :heavy_check_mark:                                                                       | ID of the portal.                                                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                     |
| `replacePortalImageAsset`                                                                | [components.ReplacePortalImageAsset](../../models/components/replaceportalimageasset.md) | :heavy_check_mark:                                                                       | Update an image asset for the portal.                                                    | {<br/>"data": "data:image/jpeg;base64,bmljZV9sb29raW5nX3BpY3R1cmU="<br/>}                |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.ReplacePortalAssetLogoResponse](../../models/operations/replaceportalassetlogoresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPortalAssetFavicon

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Returns the favicon of the portal.

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

    res, err := s.Assets.GetPortalAssetFavicon(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredImagePngPortalImageAsset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `portalID`                                               | *string*                                                 | :heavy_check_mark:                                       | ID of the portal.                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPortalAssetFaviconResponse](../../models/operations/getportalassetfaviconresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ReplacePortalAssetFavicon

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Replaces the favicon of the portal. The favicon is used in the browser tab of the portal.

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

    res, err := s.Assets.ReplacePortalAssetFavicon(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.ReplacePortalImageAsset{
        Data: "data:image/jpeg;base64,bmljZV9sb29raW5nX3BpY3R1cmU=",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredImagePngPortalImageAsset != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |                                                                                          |
| `portalID`                                                                               | *string*                                                                                 | :heavy_check_mark:                                                                       | ID of the portal.                                                                        | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                     |
| `replacePortalImageAsset`                                                                | [components.ReplacePortalImageAsset](../../models/components/replaceportalimageasset.md) | :heavy_check_mark:                                                                       | Update an image asset for the portal.                                                    | {<br/>"data": "data:image/jpeg;base64,bmljZV9sb29raW5nX3BpY3R1cmU="<br/>}                |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |                                                                                          |

### Response

**[*operations.ReplacePortalAssetFaviconResponse](../../models/operations/replaceportalassetfaviconresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |