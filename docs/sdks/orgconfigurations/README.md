# OrgConfigurations

## Overview

Operations related to org scoped notifications and their org subscriptions. The http channel of an org scoped notification is set up for the whole organization, and no user can switch it on or off. Only an org admin can use these operations. The routes exist only when http destinations are enabled.

### Available Operations

* [ListOrgConfigurations](#listorgconfigurations) - List available org configurations
* [ListOrgSubscriptions](#listorgsubscriptions) - List the org subscriptions of an event type
* [CreateOrgSubscription](#createorgsubscription) - Create an org subscription
* [GetOrgSubscription](#getorgsubscription) - Get an org subscription
* [UpdateOrgSubscription](#updateorgsubscription) - Update an org subscription
* [DeleteOrgSubscription](#deleteorgsubscription) - Delete an org subscription

## ListOrgConfigurations

List the org scoped events of the organization. An org scoped event sends its notification to the http destinations that an org admin subscribes to it, with the org subscription operations. No user can switch the http channel of an org scoped notification on or off. Only an org admin can use this operation.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-org-configurations" method="get" path="/v1/notifications/configurations/organization" -->
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

    res, err := s.OrgConfigurations.ListOrgConfigurations(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OrgConfigurationListResponse != nil {
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

**[*operations.ListOrgConfigurationsResponse](../../models/operations/listorgconfigurationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListOrgSubscriptions

List the org subscriptions of one org scoped event type, enabled or paused. Only an org admin can use this operation.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-org-subscriptions" method="get" path="/v1/notifications/configurations/organization/{eventId}/subscriptions" -->
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

    res, err := s.OrgConfigurations.ListOrgSubscriptions(ctx, "invoice-ready")
    if err != nil {
        log.Fatal(err)
    }
    if res.OrgSubscriptionListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `eventID`                                                | `string`                                                 | :heavy_check_mark:                                       | Formatted string ID of the notification event.           | invoice-ready                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListOrgSubscriptionsResponse](../../models/operations/listorgsubscriptionsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateOrgSubscription

Subscribe one or more http destinations to one org scoped event type for the regions and the entities that the request names. The request makes one subscription that holds all the destinations, so one update changes them together. When a destination does not exist, is being deleted or is not an http destination, the request answers 400 and makes no subscription. A destination can be in more than one subscription of the same event type: it still gets one delivery for each event. A subscription applies to the whole organization: no user can switch it on or off. The event type must be org scoped and must have an http template. Only an org admin can use this operation.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-org-subscription" method="post" path="/v1/notifications/configurations/organization/{eventId}/subscriptions" -->
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

    res, err := s.OrgConfigurations.CreateOrgSubscription(ctx, "invoice-ready", components.OrgSubscriptionCreatePayload{
        DestinationIds: []string{},
        Regions: []components.NotificationRegion{
            components.NotificationRegionIn,
        },
        Entities: []string{
            "<value 1>",
            "<value 2>",
        },
        Enabled: false,
        Name: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OrgSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        | Example                                                                                            |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |                                                                                                    |
| `eventID`                                                                                          | `string`                                                                                           | :heavy_check_mark:                                                                                 | Formatted string ID of the notification event.                                                     | invoice-ready                                                                                      |
| `orgSubscriptionCreatePayload`                                                                     | [components.OrgSubscriptionCreatePayload](../../models/components/orgsubscriptioncreatepayload.md) | :heavy_check_mark:                                                                                 | Request body schema for creating an org subscription.                                              |                                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |                                                                                                    |

### Response

**[*operations.CreateOrgSubscriptionResponse](../../models/operations/createorgsubscriptionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetOrgSubscription

Get one org subscription of an event type. Only an org admin can use this operation.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-org-subscription" method="get" path="/v1/notifications/configurations/organization/{eventId}/subscriptions/{subscriptionId}" -->
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

    res, err := s.OrgConfigurations.GetOrgSubscription(ctx, "invoice-ready", "2efea526-5094-430a-8577-7c0653ac52d0")
    if err != nil {
        log.Fatal(err)
    }
    if res.OrgSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `eventID`                                                | `string`                                                 | :heavy_check_mark:                                       | Formatted string ID of the notification event.           | invoice-ready                                            |
| `subscriptionID`                                         | `string`                                                 | :heavy_check_mark:                                       | ID of the org subscription, a UUID v7.                   |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetOrgSubscriptionResponse](../../models/operations/getorgsubscriptionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateOrgSubscription

Replace the name, the destinations, the regions, the entities and the enabled flag of an org subscription. The request names all five. The destinations of the request replace the stored destinations. When a destination does not exist, is being deleted or is not an http destination, the request answers 400 and changes nothing. The enabled flag applies to all the destinations of the subscription. Only an org admin can use this operation.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-org-subscription" method="patch" path="/v1/notifications/configurations/organization/{eventId}/subscriptions/{subscriptionId}" -->
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

    res, err := s.OrgConfigurations.UpdateOrgSubscription(ctx, operations.UpdateOrgSubscriptionRequest{
        EventID: "invoice-ready",
        SubscriptionID: "fa4719bc-b16a-4b07-a385-6f5fba92fbbf",
        OrgSubscriptionUpdatePayload: components.OrgSubscriptionUpdatePayload{
            DestinationIds: []string{
                "c1c16fb7-821b-47fc-8e98-d62169595727",
                "e77e5ebf-dc9b-4898-86a7-1df8224525dd",
            },
            Regions: []components.NotificationRegion{
                components.NotificationRegionUs,
            },
            Entities: []string{},
            Enabled: true,
            Name: "<value>",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OrgSubscription != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.UpdateOrgSubscriptionRequest](../../models/operations/updateorgsubscriptionrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateOrgSubscriptionResponse](../../models/operations/updateorgsubscriptionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteOrgSubscription

Delete one org subscription of an event type. Only an org admin can use this operation.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-org-subscription" method="delete" path="/v1/notifications/configurations/organization/{eventId}/subscriptions/{subscriptionId}" -->
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

    res, err := s.OrgConfigurations.DeleteOrgSubscription(ctx, "invoice-ready", "8e7f1c6a-cbea-4be6-a2a4-40f3e593e6d6")
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
| `eventID`                                                | `string`                                                 | :heavy_check_mark:                                       | Formatted string ID of the notification event.           | invoice-ready                                            |
| `subscriptionID`                                         | `string`                                                 | :heavy_check_mark:                                       | ID of the org subscription, a UUID v7.                   |                                                          |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteOrgSubscriptionResponse](../../models/operations/deleteorgsubscriptionresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |