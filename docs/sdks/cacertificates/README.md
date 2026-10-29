# CACertificates

## Overview

A CA certificate object represents a trusted certificate authority.
These objects are used by Kong Gateway to verify the validity of a client or server certificate.

### Available Operations

* [ListCaCertificate](#listcacertificate) - List all CA Certificates in a workspace
* [CreateCaCertificate](#createcacertificate) - Create a new CA Certificate in a workspace
* [DeleteCaCertificate](#deletecacertificate) - Delete a CA Certificate in a workspace
* [GetCaCertificate](#getcacertificate) - Get a CA Certificate in a workspace
* [UpsertCaCertificate](#upsertcacertificate) - Upsert a CA Certificate in a workspace

## ListCaCertificate

List all CA Certificates in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="list-ca_certificate" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/ca_certificates" -->
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

    res, err := s.CACertificates.ListCaCertificate(ctx, operations.ListCaCertificateRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Tags: sdkkonnectgo.Pointer("tag1,tag2"),
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListCaCertificateRequest](../../models/operations/listcacertificaterequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ListCaCertificateResponse](../../models/operations/listcacertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## CreateCaCertificate

Create a new CA Certificate in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="create-ca_certificate" method="post" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/ca_certificates" -->
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

    res, err := s.CACertificates.CreateCaCertificate(ctx, operations.CreateCaCertificateRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        CACertificate: components.CACertificate{
            Cert: "-----BEGIN CERTIFICATE-----\ncertificate-content\n-----END CERTIFICATE-----",
            CertDigest: sdkkonnectgo.Pointer("9b8aaf19a276885f6c8a6bc48a30700fdb3a351d8b05374d153bfb7b178e2a9f"),
            CreatedAt: sdkkonnectgo.Pointer[int64](1706598432),
            ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f260"),
            Tags: []string{
                "trusted",
                "api",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CACertificate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateCaCertificateRequest](../../models/operations/createcacertificaterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateCaCertificateResponse](../../models/operations/createcacertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## DeleteCaCertificate

Delete a CA Certificate in a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-ca_certificate" method="delete" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/ca_certificates/{CACertificateId}" -->
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

    res, err := s.CACertificates.DeleteCaCertificate(ctx, operations.DeleteCaCertificateRequest{
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        CACertificateID: "3c31f18a-f27a-4f9b-8cd4-bf841554612f",
        Workspace: "team-payments",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteCaCertificateRequest](../../models/operations/deletecacertificaterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.DeleteCaCertificateResponse](../../models/operations/deletecacertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## GetCaCertificate

Get a CA Certificate using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-ca_certificate" method="get" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/ca_certificates/{CACertificateId}" -->
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

    res, err := s.CACertificates.GetCaCertificate(ctx, operations.GetCaCertificateRequest{
        CACertificateID: "3c31f18a-f27a-4f9b-8cd4-bf841554612f",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CACertificate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetCaCertificateRequest](../../models/operations/getcacertificaterequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetCaCertificateResponse](../../models/operations/getcacertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |

## UpsertCaCertificate

Create or Update CA Certificate using ID in a workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-ca_certificate" method="put" path="/v2/control-planes/{controlPlaneId}/core-entities/{workspace}/ca_certificates/{CACertificateId}" -->
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

    res, err := s.CACertificates.UpsertCaCertificate(ctx, operations.UpsertCaCertificateRequest{
        CACertificateID: "3c31f18a-f27a-4f9b-8cd4-bf841554612f",
        ControlPlaneID: "9524ec7d-36d9-465d-a8c5-83a3c9390458",
        Workspace: "team-payments",
        CACertificate: components.CACertificate{
            Cert: "-----BEGIN CERTIFICATE-----\ncertificate-content\n-----END CERTIFICATE-----",
            CertDigest: sdkkonnectgo.Pointer("9b8aaf19a276885f6c8a6bc48a30700fdb3a351d8b05374d153bfb7b178e2a9f"),
            CreatedAt: sdkkonnectgo.Pointer[int64](1706598432),
            ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f260"),
            Tags: []string{
                "trusted",
                "api",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CACertificate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpsertCaCertificateRequest](../../models/operations/upsertcacertificaterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.UpsertCaCertificateResponse](../../models/operations/upsertcacertificateresponse.md), error**

### Errors

| Error Type                         | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.GatewayUnauthorizedError | 401                                | application/json                   |
| sdkerrors.SDKError                 | 4XX, 5XX                           | \*/\*                              |