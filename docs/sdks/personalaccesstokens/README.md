# PersonalAccessTokens
(*PersonalAccessTokens*)

## Overview

### Available Operations

* [ListUsersPersonalAccessTokens](#listuserspersonalaccesstokens) - List a user's personal access tokens for a user.
* [CreatePersonalAccessToken](#createpersonalaccesstoken) - Create a new personal access token.
* [GetPersonalAccessTokenDetails](#getpersonalaccesstokendetails) - Get Personal Access Token details.
* [UpdatePersonalAccessTokenDetails](#updatepersonalaccesstokendetails) - Update personal access token details.
* [RevokePersonalAccessToken](#revokepersonalaccesstoken) - Revoke Personal Access Token.

## ListUsersPersonalAccessTokens

List personal access tokens for a user.

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

    res, err := s.PersonalAccessTokens.ListUsersPersonalAccessTokens(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
    if err != nil {
        log.Fatal(err)
    }
    if res.PersonalAccessTokenListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | ID of the user.                                          | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ListUsersPersonalAccessTokensResponse](../../models/operations/listuserspersonalaccesstokensresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## CreatePersonalAccessToken

Create a new personal access token. A maximum of 10 personal access tokens can be created.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgointernal "github.com/Kong/sdk-konnect-go-internal"
	"github.com/Kong/sdk-konnect-go-internal/models/components"
	"github.com/Kong/sdk-konnect-go-internal/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkkonnectgointernal.New(
        sdkkonnectgointernal.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgointernal.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.PersonalAccessTokens.CreatePersonalAccessToken(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", &components.PersonalAccessTokenCreateRequest{
        Name: "<value>",
        ExpiresAt: types.MustTimeFromString("2022-11-04T20:10:06.927Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PersonalAccessTokenCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 | Example                                                                                                     |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                       | :heavy_check_mark:                                                                                          | The context to use for the request.                                                                         |                                                                                                             |
| `userID`                                                                                                    | *string*                                                                                                    | :heavy_check_mark:                                                                                          | ID of the user.                                                                                             | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                                                                        |
| `personalAccessTokenCreateRequest`                                                                          | [*components.PersonalAccessTokenCreateRequest](../../models/components/personalaccesstokencreaterequest.md) | :heavy_minus_sign:                                                                                          | Request body schema for creating personal access tokens.                                                    |                                                                                                             |
| `opts`                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                    | :heavy_minus_sign:                                                                                          | The options for this request.                                                                               |                                                                                                             |

### Response

**[*operations.CreatePersonalAccessTokenResponse](../../models/operations/createpersonalaccesstokenresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetPersonalAccessTokenDetails

Get Personal Access Token details.

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

    res, err := s.PersonalAccessTokens.GetPersonalAccessTokenDetails(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
    if err != nil {
        log.Fatal(err)
    }
    if res.PersonalAccessToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | ID of the user.                                          | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `tokenID`                                                | *string*                                                 | :heavy_check_mark:                                       | ID of the personal access token.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetPersonalAccessTokenDetailsResponse](../../models/operations/getpersonalaccesstokendetailsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## UpdatePersonalAccessTokenDetails

Update personal access token details.

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

    res, err := s.PersonalAccessTokens.UpdatePersonalAccessTokenDetails(ctx, operations.UpdatePersonalAccessTokenDetailsRequest{
        UserID: "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7",
        TokenID: "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PersonalAccessToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                                    | :heavy_check_mark:                                                                                                       | The context to use for the request.                                                                                      |
| `request`                                                                                                                | [operations.UpdatePersonalAccessTokenDetailsRequest](../../models/operations/updatepersonalaccesstokendetailsrequest.md) | :heavy_check_mark:                                                                                                       | The request object to use for the request.                                                                               |
| `opts`                                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                                 | :heavy_minus_sign:                                                                                                       | The options for this request.                                                                                            |

### Response

**[*operations.UpdatePersonalAccessTokenDetailsResponse](../../models/operations/updatepersonalaccesstokendetailsresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## RevokePersonalAccessToken

Revoke Personal Access Token.

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

    res, err := s.PersonalAccessTokens.RevokePersonalAccessToken(ctx, "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7", "5f9fd312-a987-4628-b4c5-bb4f4fddd5f7")
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
| `userID`                                                 | *string*                                                 | :heavy_check_mark:                                       | ID of the user.                                          | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `tokenID`                                                | *string*                                                 | :heavy_check_mark:                                       | ID of the personal access token.                         | 5f9fd312-a987-4628-b4c5-bb4f4fddd5f7                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.RevokePersonalAccessTokenResponse](../../models/operations/revokepersonalaccesstokenresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |