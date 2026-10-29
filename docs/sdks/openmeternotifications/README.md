# OpenMeterNotifications

## Overview

Notification rules generate notification events, such as entitlement balance threshold crossings, and notification channels deliver them to external systems.

### Available Operations

* [ListBillingNotificationChannels](#listbillingnotificationchannels) - List notification channels
* [CreateBillingNotificationChannel](#createbillingnotificationchannel) - Create notification channel
* [GetBillingNotificationChannel](#getbillingnotificationchannel) - Get notification channel
* [UpdateBillingNotificationChannel](#updatebillingnotificationchannel) - Update notification channel
* [DeleteBillingNotificationChannel](#deletebillingnotificationchannel) - Delete notification channel
* [ListBillingNotificationEvents](#listbillingnotificationevents) - List notification events
* [GetBillingNotificationEvent](#getbillingnotificationevent) - Get notification event
* [ResendBillingNotificationEvent](#resendbillingnotificationevent) - Resend notification event
* [ListBillingNotificationRules](#listbillingnotificationrules) - List notification rules
* [CreateBillingNotificationRule](#createbillingnotificationrule) - Create notification rule
* [GetBillingNotificationRule](#getbillingnotificationrule) - Get notification rule
* [UpdateBillingNotificationRule](#updatebillingnotificationrule) - Update notification rule
* [DeleteBillingNotificationRule](#deletebillingnotificationrule) - Delete notification rule
* [TestBillingNotificationRule](#testbillingnotificationrule) - Test notification rule

## ListBillingNotificationChannels

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

List all notification channels.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-billing-notification-channels" method="get" path="/v3/openmeter/notification/channels" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/types"
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

    res, err := s.OpenMeterNotifications.ListBillingNotificationChannels(ctx, operations.ListBillingNotificationChannelsRequest{
        Sort: sdkkonnectgo.Pointer("created_at desc"),
        Filter: &components.ListNotificationChannelsParamsFilter{
            ID: sdkkonnectgo.Pointer(components.CreateULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
            CreatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldGTEFilter(
                components.DateTimeFieldGTEFilter{
                    Gte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            UpdatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                components.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.NotificationChannelPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.ListBillingNotificationChannelsRequest](../../models/operations/listbillingnotificationchannelsrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.ListBillingNotificationChannelsResponse](../../models/operations/listbillingnotificationchannelsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateBillingNotificationChannel

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Create a notification channel.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-billing-notification-channel" method="post" path="/v3/openmeter/notification/channels" -->
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

    res, err := s.OpenMeterNotifications.CreateBillingNotificationChannel(ctx, components.CreateNotificationChannelRequest{
        Name: "<value>",
        Labels: map[string]string{
            "env": "test",
        },
        Type: components.CreateNotificationChannelRequestTypeWebhook,
        URL: "https://grubby-gift.name/",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingNotificationChannel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [components.CreateNotificationChannelRequest](../../models/components/createnotificationchannelrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.CreateBillingNotificationChannelResponse](../../models/operations/createbillingnotificationchannelresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetBillingNotificationChannel

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Get a notification channel by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-billing-notification-channel" method="get" path="/v3/openmeter/notification/channels/{notificationChannelId}" -->
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

    res, err := s.OpenMeterNotifications.GetBillingNotificationChannel(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingNotificationChannel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `notificationChannelID`                                  | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetBillingNotificationChannelResponse](../../models/operations/getbillingnotificationchannelresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateBillingNotificationChannel

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Update a notification channel by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-billing-notification-channel" method="put" path="/v3/openmeter/notification/channels/{notificationChannelId}" -->
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

    res, err := s.OpenMeterNotifications.UpdateBillingNotificationChannel(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.UpdateBillingNotificationChannelRequest{
        Name: "<value>",
        Type: components.UpdateBillingNotificationChannelRequestTypeWebhook,
        URL: "https://ironclad-taxicab.com/",
        Labels: map[string]string{
            "env": "test",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingNotificationChannel != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              | Example                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |                                                                                                                          |
| `notificationChannelID`                                                                                                  | `string`                                                                                                                 | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                                                                                               |
| `updateBillingNotificationChannelRequest`                                                                                | [components.UpdateBillingNotificationChannelRequest](../../models/components/updatebillingnotificationchannelrequest.md) | :heavy_check_mark:                                                                                                       | N/A                                                                                                                      |                                                                                                                          |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |                                                                                                                          |

### Response

**[*operations.UpdateBillingNotificationChannelResponse](../../models/operations/updatebillingnotificationchannelresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteBillingNotificationChannel

**Pre-release Endpoint**
This endpoint is currently in beta and is subject to change.

Delete a notification channel by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-billing-notification-channel" method="delete" path="/v3/openmeter/notification/channels/{notificationChannelId}" -->
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

    res, err := s.OpenMeterNotifications.DeleteBillingNotificationChannel(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
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
| `notificationChannelID`                                  | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteBillingNotificationChannelResponse](../../models/operations/deletebillingnotificationchannelresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListBillingNotificationEvents

List all notification events.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-billing-notification-events" method="get" path="/v3/openmeter/notification/events" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/types"
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

    res, err := s.OpenMeterNotifications.ListBillingNotificationEvents(ctx, operations.ListBillingNotificationEventsRequest{
        Sort: sdkkonnectgo.Pointer("created_at desc"),
        Filter: &components.ListNotificationEventsParamsFilter{
            ID: sdkkonnectgo.Pointer(components.CreateULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
            CreatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                components.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            RuleID: sdkkonnectgo.Pointer(components.CreateULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
            ChannelID: sdkkonnectgo.Pointer(components.CreateListNotificationEventsParamsFilterULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
            SubjectID: sdkkonnectgo.Pointer(components.CreateListNotificationEventsParamsFilterSubjectIDULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
            FeatureID: sdkkonnectgo.Pointer(components.CreateListNotificationEventsParamsFilterFeatureIDULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.NotificationEventPagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListBillingNotificationEventsRequest](../../models/operations/listbillingnotificationeventsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListBillingNotificationEventsResponse](../../models/operations/listbillingnotificationeventsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetBillingNotificationEvent

Get a notification event by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-billing-notification-event" method="get" path="/v3/openmeter/notification/events/{notificationEventId}" -->
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

    res, err := s.OpenMeterNotifications.GetBillingNotificationEvent(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingNotificationEvent != nil {
        switch res.BillingNotificationEvent.Payload.Type {
            case components.PayloadUnionTypeEntitlementsBalanceThreshold:
                // res.BillingNotificationEvent.Payload.BillingNotificationEventBalanceThresholdPayload is populated
            case components.PayloadUnionTypeEntitlementsReset:
                // res.BillingNotificationEvent.Payload.BillingNotificationEventResetPayload is populated
            case components.PayloadUnionTypeInvoiceCreated:
                // res.BillingNotificationEvent.Payload.BillingNotificationEventInvoiceCreatedPayload is populated
            case components.PayloadUnionTypeInvoiceUpdated:
                // res.BillingNotificationEvent.Payload.BillingNotificationEventInvoiceUpdatedPayload is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `notificationEventID`                                    | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetBillingNotificationEventResponse](../../models/operations/getbillingnotificationeventresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ResendBillingNotificationEvent

Resend a notification event to the channels of the rule that generated it.

Delivery is asynchronous: the request marks the selected channels for redelivery
and returns immediately. Channels whose delivery is still pending or already
being resent are left untouched.

### Example Usage

<!-- UsageSnippet language="go" operationID="resend-billing-notification-event" method="post" path="/v3/openmeter/notification/events/{notificationEventId}/resend" -->
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

    res, err := s.OpenMeterNotifications.ResendBillingNotificationEvent(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.ResendBillingNotificationEventRequest{
        Channels: []components.NotificationChannelReference{
            components.NotificationChannelReference{
                ID: "01G65Z755AFWAKHE12NY0CQ9FH",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          | Example                                                                                                              |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |                                                                                                                      |
| `notificationEventID`                                                                                                | `string`                                                                                                             | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  | 01G65Z755AFWAKHE12NY0CQ9FH                                                                                           |
| `resendBillingNotificationEventRequest`                                                                              | [components.ResendBillingNotificationEventRequest](../../models/components/resendbillingnotificationeventrequest.md) | :heavy_check_mark:                                                                                                   | N/A                                                                                                                  |                                                                                                                      |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |                                                                                                                      |

### Response

**[*operations.ResendBillingNotificationEventResponse](../../models/operations/resendbillingnotificationeventresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListBillingNotificationRules

List all notification rules.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-billing-notification-rules" method="get" path="/v3/openmeter/notification/rules" -->
```go
package main

import(
	"context"
	"github.com/Kong/sdk-konnect-go/models/components"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/types"
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

    res, err := s.OpenMeterNotifications.ListBillingNotificationRules(ctx, operations.ListBillingNotificationRulesRequest{
        Sort: sdkkonnectgo.Pointer("created_at desc"),
        Filter: &components.ListNotificationRulesParamsFilter{
            ID: sdkkonnectgo.Pointer(components.CreateULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
            CreatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTFilter(
                components.DateTimeFieldLTFilter{
                    Lt: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            UpdatedAt: sdkkonnectgo.Pointer(components.CreateDateTimeFieldFilterDateTimeFieldLTEFilter(
                components.DateTimeFieldLTEFilter{
                    Lte: types.MustTimeFromString("2022-03-30T07:20:50Z"),
                },
            )),
            ChannelID: sdkkonnectgo.Pointer(components.CreateListNotificationRulesParamsFilterULIDFieldFilterStr(
                "01G65Z755AFWAKHE12NY0CQ9FH",
            )),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.NotificationRulePagePaginatedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.ListBillingNotificationRulesRequest](../../models/operations/listbillingnotificationrulesrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.ListBillingNotificationRulesResponse](../../models/operations/listbillingnotificationrulesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateBillingNotificationRule

Create a notification rule.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-billing-notification-rule" method="post" path="/v3/openmeter/notification/rules" -->
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

    res, err := s.OpenMeterNotifications.CreateBillingNotificationRule(ctx, components.CreateBillingNotificationRuleRequestInvoiceCreated(
        components.BillingNotificationRuleInvoiceCreatedRequest{
            Name: "<value>",
            Type: components.BillingNotificationRuleInvoiceCreatedRequestTypeInvoiceCreated,
            Channels: []components.NotificationChannelReference{},
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingNotificationRule != nil {
        switch res.BillingNotificationRule.Type {
            case components.BillingNotificationRuleTypeEntitlementsBalanceThreshold:
                // res.BillingNotificationRule.BillingNotificationRuleBalanceThreshold is populated
            case components.BillingNotificationRuleTypeEntitlementsReset:
                // res.BillingNotificationRule.BillingNotificationRuleEntitlementReset is populated
            case components.BillingNotificationRuleTypeInvoiceCreated:
                // res.BillingNotificationRule.BillingNotificationRuleInvoiceCreated is populated
            case components.BillingNotificationRuleTypeInvoiceUpdated:
                // res.BillingNotificationRule.BillingNotificationRuleInvoiceUpdated is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [components.BillingNotificationRuleRequest](../../models/components/billingnotificationrulerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.CreateBillingNotificationRuleResponse](../../models/operations/createbillingnotificationruleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetBillingNotificationRule

Get a notification rule by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-billing-notification-rule" method="get" path="/v3/openmeter/notification/rules/{notificationRuleId}" -->
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

    res, err := s.OpenMeterNotifications.GetBillingNotificationRule(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingNotificationRule != nil {
        switch res.BillingNotificationRule.Type {
            case components.BillingNotificationRuleTypeEntitlementsBalanceThreshold:
                // res.BillingNotificationRule.BillingNotificationRuleBalanceThreshold is populated
            case components.BillingNotificationRuleTypeEntitlementsReset:
                // res.BillingNotificationRule.BillingNotificationRuleEntitlementReset is populated
            case components.BillingNotificationRuleTypeInvoiceCreated:
                // res.BillingNotificationRule.BillingNotificationRuleInvoiceCreated is populated
            case components.BillingNotificationRuleTypeInvoiceUpdated:
                // res.BillingNotificationRule.BillingNotificationRuleInvoiceUpdated is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `notificationRuleID`                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetBillingNotificationRuleResponse](../../models/operations/getbillingnotificationruleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateBillingNotificationRule

Update a notification rule by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-billing-notification-rule" method="put" path="/v3/openmeter/notification/rules/{notificationRuleId}" -->
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

    res, err := s.OpenMeterNotifications.UpdateBillingNotificationRule(ctx, "01G65Z755AFWAKHE12NY0CQ9FH", components.CreateBillingNotificationRuleRequestEntitlementsReset(
        components.BillingNotificationRuleEntitlementResetRequest{
            Name: "<value>",
            Type: components.BillingNotificationRuleEntitlementResetRequestTypeEntitlementsReset,
            Channels: []components.NotificationChannelReference{},
        },
    ))
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingNotificationRule != nil {
        switch res.BillingNotificationRule.Type {
            case components.BillingNotificationRuleTypeEntitlementsBalanceThreshold:
                // res.BillingNotificationRule.BillingNotificationRuleBalanceThreshold is populated
            case components.BillingNotificationRuleTypeEntitlementsReset:
                // res.BillingNotificationRule.BillingNotificationRuleEntitlementReset is populated
            case components.BillingNotificationRuleTypeInvoiceCreated:
                // res.BillingNotificationRule.BillingNotificationRuleInvoiceCreated is populated
            case components.BillingNotificationRuleTypeInvoiceUpdated:
                // res.BillingNotificationRule.BillingNotificationRuleInvoiceUpdated is populated
        }

    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            | Example                                                                                                |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |                                                                                                        |
| `notificationRuleID`                                                                                   | `string`                                                                                               | :heavy_check_mark:                                                                                     | N/A                                                                                                    | 01G65Z755AFWAKHE12NY0CQ9FH                                                                             |
| `billingNotificationRuleRequest`                                                                       | [components.BillingNotificationRuleRequest](../../models/components/billingnotificationrulerequest.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |                                                                                                        |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |                                                                                                        |

### Response

**[*operations.UpdateBillingNotificationRuleResponse](../../models/operations/updatebillingnotificationruleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteBillingNotificationRule

Delete a notification rule by id.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-billing-notification-rule" method="delete" path="/v3/openmeter/notification/rules/{notificationRuleId}" -->
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

    res, err := s.OpenMeterNotifications.DeleteBillingNotificationRule(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
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
| `notificationRuleID`                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.DeleteBillingNotificationRuleResponse](../../models/operations/deletebillingnotificationruleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## TestBillingNotificationRule

Test a notification rule by generating an event with sample data and delivering
it to the rule's channels. The test event is persisted and listed like any other
event.

### Example Usage

<!-- UsageSnippet language="go" operationID="test-billing-notification-rule" method="post" path="/v3/openmeter/notification/rules/{notificationRuleId}/test" -->
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

    res, err := s.OpenMeterNotifications.TestBillingNotificationRule(ctx, "01G65Z755AFWAKHE12NY0CQ9FH")
    if err != nil {
        log.Fatal(err)
    }
    if res.BillingNotificationEvent != nil {
        switch res.BillingNotificationEvent.Payload.Type {
            case components.PayloadUnionTypeEntitlementsBalanceThreshold:
                // res.BillingNotificationEvent.Payload.BillingNotificationEventBalanceThresholdPayload is populated
            case components.PayloadUnionTypeEntitlementsReset:
                // res.BillingNotificationEvent.Payload.BillingNotificationEventResetPayload is populated
            case components.PayloadUnionTypeInvoiceCreated:
                // res.BillingNotificationEvent.Payload.BillingNotificationEventInvoiceCreatedPayload is populated
            case components.PayloadUnionTypeInvoiceUpdated:
                // res.BillingNotificationEvent.Payload.BillingNotificationEventInvoiceUpdatedPayload is populated
        }

    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `notificationRuleID`                                     | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      | 01G65Z755AFWAKHE12NY0CQ9FH                               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.TestBillingNotificationRuleResponse](../../models/operations/testbillingnotificationruleresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |