# Destinations

## Overview

Operations related to http destinations and their deliveries. Only an org admin can use these operations. The destination routes exist only when http destinations are enabled.

### Available Operations

* [ListDestinations](#listdestinations) - List destinations
* [CreateDestination](#createdestination) - Create a destination
* [GetDestination](#getdestination) - Get a destination
* [UpdateDestination](#updatedestination) - Update a destination
* [DeleteDestination](#deletedestination) - Delete a destination
* [TestDestination](#testdestination) - Test a destination
* [RotateDestinationSecret](#rotatedestinationsecret) - Rotate the destination signing secret
* [GetDestinationDelivery](#getdestinationdelivery) - Get a delivery and its attempts
* [ResendDestinationDelivery](#resenddestinationdelivery) - Resend a delivery
* [ListDestinationSubscriptions](#listdestinationsubscriptions) - List the org subscriptions of a destination
* [ListDeliveries](#listdeliveries) - List deliveries of the organization

## ListDestinations

List the http destinations of an organization, newest first. filter[event_type] returns only the destinations with an enabled org subscription to that event type.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-destinations" method="get" path="/v1/notifications/destinations" -->
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

    res, err := s.Destinations.ListDestinations(ctx, operations.ListDestinationsRequest{
        PageBefore: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DestinationListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListDestinationsRequest](../../models/operations/listdestinationsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ListDestinationsResponse](../../models/operations/listdestinationsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreateDestination

Create an http destination for the organization. The new destination is pending verification: the request does not call the endpoint. A successful delivery, or a successful test, enables the destination. The answer holds the signing secret once. An organization holds at most 50 destinations, in every state; creating one beyond the limit answers 409.

### Example Usage

<!-- UsageSnippet language="go" operationID="create-destination" method="post" path="/v1/notifications/destinations" -->
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

    res, err := s.Destinations.CreateDestination(ctx, components.DestinationCreatePayload{
        Name: "Production webhooks",
        Type: components.DestinationTypeHTTP,
        URL: "https://example.com/webhooks",
        CustomHeaders: map[string]string{
            "X-Environment": "production",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DestinationWithSecret != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [components.DestinationCreatePayload](../../models/components/destinationcreatepayload.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.CreateDestinationResponse](../../models/operations/createdestinationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetDestination

Get one http destination.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-destination" method="get" path="/v1/notifications/destinations/{destinationId}" -->
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

    res, err := s.Destinations.GetDestination(ctx, "d26034f1-0cdd-4b6f-9158-13c03f4cd006")
    if err != nil {
        log.Fatal(err)
    }
    if res.Destination != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `destinationID`                                          | `string`                                                 | :heavy_check_mark:                                       | ID of the http destination, a UUID v7.                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetDestinationResponse](../../models/operations/getdestinationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdateDestination

Update one http destination. A change of the url follows the rules of the create request. It does not verify the destination again, and it does not change the state or the signing key. A destination in any state can change to enabled or disabled; enabling a destination that is pending verification skips the verification. A concurrent change of the destination answers 409.

### Example Usage

<!-- UsageSnippet language="go" operationID="update-destination" method="patch" path="/v1/notifications/destinations/{destinationId}" -->
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

    res, err := s.Destinations.UpdateDestination(ctx, "310711be-def8-442b-89d4-1677c4fb3473", components.DestinationUpdatePayload{
        Type: components.DestinationTypeHTTP.ToPointer(),
        Name: sdkkonnectgo.Pointer("Production webhooks"),
        CustomHeaders: map[string]string{
            "X-Environment": "staging",
        },
        State: components.DestinationUpdatePayloadStateEnabled.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Destination != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                               | Type                                                                                                                    | Required                                                                                                                | Description                                                                                                             | Example                                                                                                                 |
| ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                   | :heavy_check_mark:                                                                                                      | The context to use for the request.                                                                                     |                                                                                                                         |
| `destinationID`                                                                                                         | `string`                                                                                                                | :heavy_check_mark:                                                                                                      | ID of the http destination, a UUID v7.                                                                                  |                                                                                                                         |
| `destinationUpdatePayload`                                                                                              | [components.DestinationUpdatePayload](../../models/components/destinationupdatepayload.md)                              | :heavy_check_mark:                                                                                                      | Request body schema for updating a destination.                                                                         | {<br/>"type": "http",<br/>"name": "Production webhooks",<br/>"custom_headers": {<br/>"X-Environment": "staging"<br/>},<br/>"state": "enabled"<br/>} |
| `opts`                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                | :heavy_minus_sign:                                                                                                      | The options for this request.                                                                                           |                                                                                                                         |

### Response

**[*operations.UpdateDestinationResponse](../../models/operations/updatedestinationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## DeleteDestination

Delete one http destination. The request removes the destination from every org subscription that holds it, and deletes an org subscription that holds no other destination. It also deletes the signing keys and all the deliveries of the destination. It cannot be undone.

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-destination" method="delete" path="/v1/notifications/destinations/{destinationId}" -->
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

    res, err := s.Destinations.DeleteDestination(ctx, "f535bc11-b5c3-4614-b8f4-a9d563f17078")
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
| `destinationID`                                          | `string`                                                 | :heavy_check_mark:                                       | ID of the http destination, a UUID v7.                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteDestinationResponse](../../models/operations/deletedestinationresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## TestDestination

Execute a destination connection test. The test is one delivery with the mode test. It gets one attempt and no retry, and its outcome is visible in the delivery log. The answer comes before the attempt, so the delivery in the answer is still scheduled. A test that succeeds enables a destination that is pending verification, after the attempt and not in the answer. A disabled destination answers 409. A test within 5 seconds of another test of the same destination answers 429.

### Example Usage

<!-- UsageSnippet language="go" operationID="test-destination" method="post" path="/v1/notifications/destinations/{destinationId}/test" -->
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

    res, err := s.Destinations.TestDestination(ctx, "92133a48-f45f-426c-86c6-a72eae55a903")
    if err != nil {
        log.Fatal(err)
    }
    if res.Delivery != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `destinationID`                                          | `string`                                                 | :heavy_check_mark:                                       | ID of the http destination, a UUID v7.                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.TestDestinationResponse](../../models/operations/testdestinationresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.ConflictError        | 409                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## RotateDestinationSecret

Rotate the signing secret of an http destination. The new secret is returned once, in this response. The replaced secret stays valid for an overlap of 24 hours, so the consumer can switch without losing a delivery. A rotation during that overlap answers 409.

### Example Usage

<!-- UsageSnippet language="go" operationID="rotate-destination-secret" method="post" path="/v1/notifications/destinations/{destinationId}/secret/rotate" -->
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

    res, err := s.Destinations.RotateDestinationSecret(ctx, "3cf186ca-b834-4e41-a770-136cdef73322")
    if err != nil {
        log.Fatal(err)
    }
    if res.DestinationSecret != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `destinationID`                                          | `string`                                                 | :heavy_check_mark:                                       | ID of the http destination, a UUID v7.                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.RotateDestinationSecretResponse](../../models/operations/rotatedestinationsecretresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.ConflictError     | 409                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetDestinationDelivery

Get one delivery of a destination and all its attempts.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-destination-delivery" method="get" path="/v1/notifications/destinations/{destinationId}/deliveries/{deliveryId}" -->
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

    res, err := s.Destinations.GetDestinationDelivery(ctx, "4e6f10e2-b5b7-4642-a338-bfeec5a0f74d", "70c31e56-5ade-409c-bac3-95884d81ed86")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeliveryDetailResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `destinationID`                                          | `string`                                                 | :heavy_check_mark:                                       | ID of the http destination, a UUID v7.                   |
| `deliveryID`                                             | `string`                                                 | :heavy_check_mark:                                       | ID of the delivery, a UUID v7.                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetDestinationDeliveryResponse](../../models/operations/getdestinationdeliveryresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ResendDestinationDelivery

Trigger a manual redelivery of one delivery. The resend is a new delivery with its own delivery ID, so its webhook-id header differs from the one of the original delivery. The payload is the same bytes as the payload of the original delivery, and the new delivery keeps its event_id. The resend keeps the mode of the original delivery. A delivery without a payload, as a delivery that failed to render, answers 400. A disabled destination answers 409, and so does a delivery that is still scheduled and not expired, which includes a delivery that waits for a retry. A resend within 5 seconds of another delivery of the same event to the same destination answers 429.

### Example Usage

<!-- UsageSnippet language="go" operationID="resend-destination-delivery" method="post" path="/v1/notifications/destinations/{destinationId}/deliveries/{deliveryId}/resend" -->
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

    res, err := s.Destinations.ResendDestinationDelivery(ctx, "0c0bdaad-5a5e-4f8a-9f99-503d665b1ee6", "c5c24211-1c53-478c-9f22-9be0d9e20a22")
    if err != nil {
        log.Fatal(err)
    }
    if res.Delivery != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `destinationID`                                          | `string`                                                 | :heavy_check_mark:                                       | ID of the http destination, a UUID v7.                   |
| `deliveryID`                                             | `string`                                                 | :heavy_check_mark:                                       | ID of the delivery, a UUID v7.                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ResendDestinationDeliveryResponse](../../models/operations/resenddestinationdeliveryresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.BadRequestError      | 400                            | application/problem+json       |
| sdkerrors.UnauthorizedError    | 401                            | application/problem+json       |
| sdkerrors.ForbiddenError       | 403                            | application/problem+json       |
| sdkerrors.NotFoundError        | 404                            | application/problem+json       |
| sdkerrors.ConflictError        | 409                            | application/problem+json       |
| sdkerrors.TooManyRequestsError | 429                            | application/problem+json       |
| sdkerrors.SDKError             | 4XX, 5XX                       | \*/\*                          |

## ListDestinationSubscriptions

List the org subscriptions that hold one destination, for every event type, enabled or paused, oldest first. Each subscription lists all its destinations. The list is not paginated.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-destination-subscriptions" method="get" path="/v1/notifications/destinations/{destinationId}/subscriptions" -->
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

    res, err := s.Destinations.ListDestinationSubscriptions(ctx, "e2be6e2d-f3ab-4205-aa48-7a0052bfec36")
    if err != nil {
        log.Fatal(err)
    }
    if res.OrgSubscriptionListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `destinationID`                                          | `string`                                                 | :heavy_check_mark:                                       | ID of the http destination, a UUID v7.                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListDestinationSubscriptionsResponse](../../models/operations/listdestinationsubscriptionsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ListDeliveries

List the deliveries of all the destinations of the organization, newest first. Without a filter, the list holds every delivery of the organization. Each delivery names its destination. filter[destination_id] keeps the deliveries of one destination. A destination that does not exist, or that belongs to another organization, gives an empty list, not a 404. The payload is present on the detail view only.

### Example Usage

<!-- UsageSnippet language="go" operationID="list-deliveries" method="get" path="/v1/notifications/deliveries" -->
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

    res, err := s.Destinations.ListDeliveries(ctx, operations.ListDeliveriesRequest{
        PageBefore: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        PageAfter: sdkkonnectgo.Pointer("ewogICJpZCI6ICJoZWxsbyB3b3JsZCIKfQ"),
        Filter: &components.OrgDeliveryFilterParameters{
            Status: &components.StatusFieldFilter{
                Eq: sdkkonnectgo.Pointer("?filter[status][eq]=failed"),
                Neq: sdkkonnectgo.Pointer("?filter[status][neq]=failed"),
                Contains: sdkkonnectgo.Pointer("?filter[status][contains]=fail"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeliveryListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListDeliveriesRequest](../../models/operations/listdeliveriesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ListDeliveriesResponse](../../models/operations/listdeliveriesresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |