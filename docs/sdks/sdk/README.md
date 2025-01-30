# SDK

## Overview

Konnect API - Go Internal SDK: The Konnect platform API

Documentation for Kong Gateway and its APIs
<https://docs.konghq.com>

### Available Operations

* [GetUsersInternal](#getusersinternal) - Retrieves users matching a provided auth0 subject and email.

## GetUsersInternal

The response contains user ids and (non-deleted, non-deleting) organization ids for users associated with provided auth0 user subject.

### Example Usage

```go
package main

import(
	"context"
	sdkkonnectgo "github.com/Kong/sdk-konnect-go"
	"github.com/Kong/sdk-konnect-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkkonnectgo.New(
        sdkkonnectgo.WithSecurity(components.Security{
            PersonalAccessToken: sdkkonnectgo.String("<YOUR_BEARER_TOKEN_HERE>"),
        }),
    )

    res, err := s.GetUsersInternal(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.InternalUsers != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `aid`                                                    | **string*                                                | :heavy_minus_sign:                                       | Auth0 User ID                                            |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetUsersInternalResponse](../../models/operations/getusersinternalresponse.md), error**

### Errors

| Error Type                | Status Code               | Content Type              |
| ------------------------- | ------------------------- | ------------------------- |
| sdkerrors.BadRequestError | 400                       | application/problem+json  |
| sdkerrors.SDKError        | 4XX, 5XX                  | \*/\*                     |