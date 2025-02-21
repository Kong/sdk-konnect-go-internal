# PortalCustomization
(*PortalCustomization*)

## Overview

### Available Operations

* [GetPortalCustomizationV3](#getportalcustomizationv3) - Get Customization
* [ReplacePortalCustomizationV3](#replaceportalcustomizationv3) - Replace Customization
* [UpdatePortalCustomizationV3](#updateportalcustomizationv3) - Update Portal Customization

## GetPortalCustomizationV3

Returns the portal customization options.

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

    res, err := s.PortalCustomization.GetPortalCustomizationV3(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a")
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalCustomizationV3 != nil {
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

**[*operations.GetPortalCustomizationV3Response](../../models/operations/getportalcustomizationv3response.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ReplacePortalCustomizationV3

Replace the portal customization options.

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

    res, err := s.PortalCustomization.ReplacePortalCustomizationV3(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalCustomizationV3{
        Menu: &components.Menu{
            Main: []components.PortalMenuItemV3{
                components.PortalMenuItemV3{
                    Path: "/about/company",
                    Title: "My Page",
                    Visibility: components.VisibilityPublic,
                    External: true,
                },
                components.PortalMenuItemV3{
                    Path: "/about/company",
                    Title: "My Page",
                    Visibility: components.VisibilityPublic,
                    External: false,
                },
                components.PortalMenuItemV3{
                    Path: "/about/company",
                    Title: "My Page",
                    Visibility: components.VisibilityPublic,
                    External: false,
                },
            },
            FooterSections: []components.PortalFooterMenuSectionV3{
                components.PortalFooterMenuSectionV3{
                    Title: "<value>",
                    Items: []components.PortalMenuItemV3{
                        components.PortalMenuItemV3{
                            Path: "/about/company",
                            Title: "My Page",
                            Visibility: components.VisibilityPublic,
                            External: true,
                        },
                        components.PortalMenuItemV3{
                            Path: "/about/company",
                            Title: "My Page",
                            Visibility: components.VisibilityPublic,
                            External: true,
                        },
                    },
                },
                components.PortalFooterMenuSectionV3{
                    Title: "<value>",
                    Items: []components.PortalMenuItemV3{
                        components.PortalMenuItemV3{
                            Path: "/about/company",
                            Title: "My Page",
                            Visibility: components.VisibilityPublic,
                            External: false,
                        },
                        components.PortalMenuItemV3{
                            Path: "/about/company",
                            Title: "My Page",
                            Visibility: components.VisibilityPublic,
                            External: true,
                        },
                        components.PortalMenuItemV3{
                            Path: "/about/company",
                            Title: "My Page",
                            Visibility: components.VisibilityPublic,
                            External: true,
                        },
                    },
                },
            },
            FooterBottom: []components.PortalMenuItemV3{
                components.PortalMenuItemV3{
                    Path: "/about/company",
                    Title: "My Page",
                    Visibility: components.VisibilityPublic,
                    External: true,
                },
                components.PortalMenuItemV3{
                    Path: "/about/company",
                    Title: "My Page",
                    Visibility: components.VisibilityPublic,
                    External: true,
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalCustomizationV3 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           | Example                                                                               |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |                                                                                       |
| `portalID`                                                                            | *string*                                                                              | :heavy_check_mark:                                                                    | ID of the portal.                                                                     | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                  |
| `portalCustomizationV3`                                                               | [*components.PortalCustomizationV3](../../models/components/portalcustomizationv3.md) | :heavy_minus_sign:                                                                    | N/A                                                                                   |                                                                                       |
| `opts`                                                                                | [][operations.Option](../../models/operations/option.md)                              | :heavy_minus_sign:                                                                    | The options for this request.                                                         |                                                                                       |

### Response

**[*operations.ReplacePortalCustomizationV3Response](../../models/operations/replaceportalcustomizationv3response.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePortalCustomizationV3

Update the portal customization options, merging properties.

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

    res, err := s.PortalCustomization.UpdatePortalCustomizationV3(ctx, "f32d905a-ed33-46a3-a093-d8f536af9a8a", &components.PortalCustomizationV3{
        Menu: &components.Menu{
            Main: []components.PortalMenuItemV3{
                components.PortalMenuItemV3{
                    Path: "/about/company",
                    Title: "My Page",
                    Visibility: components.VisibilityPublic,
                    External: true,
                },
                components.PortalMenuItemV3{
                    Path: "/about/company",
                    Title: "My Page",
                    Visibility: components.VisibilityPublic,
                    External: true,
                },
            },
            FooterSections: []components.PortalFooterMenuSectionV3{
                components.PortalFooterMenuSectionV3{
                    Title: "<value>",
                    Items: []components.PortalMenuItemV3{
                        components.PortalMenuItemV3{
                            Path: "/about/company",
                            Title: "My Page",
                            Visibility: components.VisibilityPublic,
                            External: false,
                        },
                        components.PortalMenuItemV3{
                            Path: "/about/company",
                            Title: "My Page",
                            Visibility: components.VisibilityPublic,
                            External: false,
                        },
                        components.PortalMenuItemV3{
                            Path: "/about/company",
                            Title: "My Page",
                            Visibility: components.VisibilityPublic,
                            External: false,
                        },
                    },
                },
                components.PortalFooterMenuSectionV3{
                    Title: "<value>",
                    Items: []components.PortalMenuItemV3{

                    },
                },
            },
            FooterBottom: []components.PortalMenuItemV3{
                components.PortalMenuItemV3{
                    Path: "/about/company",
                    Title: "My Page",
                    Visibility: components.VisibilityPublic,
                    External: false,
                },
                components.PortalMenuItemV3{
                    Path: "/about/company",
                    Title: "My Page",
                    Visibility: components.VisibilityPublic,
                    External: false,
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PortalCustomizationV3 != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           | Example                                                                               |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |                                                                                       |
| `portalID`                                                                            | *string*                                                                              | :heavy_check_mark:                                                                    | ID of the portal.                                                                     | f32d905a-ed33-46a3-a093-d8f536af9a8a                                                  |
| `portalCustomizationV3`                                                               | [*components.PortalCustomizationV3](../../models/components/portalcustomizationv3.md) | :heavy_minus_sign:                                                                    | N/A                                                                                   |                                                                                       |
| `opts`                                                                                | [][operations.Option](../../models/operations/option.md)                              | :heavy_minus_sign:                                                                    | The options for this request.                                                         |                                                                                       |

### Response

**[*operations.UpdatePortalCustomizationV3Response](../../models/operations/updateportalcustomizationv3response.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |