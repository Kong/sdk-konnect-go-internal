# DeclarativeConfiguration

## Overview

### Available Operations

* [UpsertDeclarativeConfig](#upsertdeclarativeconfig) - Create or Update the declarative config
* [GetDeclarativeConfig](#getdeclarativeconfig) - Get the declarative configuration
* [DeleteDeclarativeConfig](#deletedeclarativeconfig) - Delete the declarative config
* [ValidateControlPlaneConfig](#validatecontrolplaneconfig) - Validate a control plane configuration

## UpsertDeclarativeConfig

Create or Update the declarative config

### Example Usage

<!-- UsageSnippet language="go" operationID="upsert-declarative-config" method="put" path="/v2/control-planes/{controlPlaneId}/declarative-config" -->
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

    res, err := s.DeclarativeConfiguration.UpsertDeclarativeConfig(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.DeclarativeConfigInput{})
    if err != nil {
        log.Fatal(err)
    }
    if res.DeclarativeConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `controlPlaneID`                                                                       | `string`                                                                               | :heavy_check_mark:                                                                     | The UUID of your control plane. This variable is available in the Konnect manager.     | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                                   |
| `declarativeConfig`                                                                    | [components.DeclarativeConfigInput](../../models/components/declarativeconfiginput.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |                                                                                        |

### Response

**[*operations.UpsertDeclarativeConfigResponse](../../models/operations/upsertdeclarativeconfigresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.BadRequestError   | 400                         | application/problem+json    |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## GetDeclarativeConfig

Returns the declarative configuration set for the control plane.

### Example Usage

<!-- UsageSnippet language="go" operationID="get-declarative-config" method="get" path="/v2/control-planes/{controlPlaneId}/declarative-config" -->
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

    res, err := s.DeclarativeConfiguration.GetDeclarativeConfig(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res.DeclarativeConfig != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.GetDeclarativeConfigResponse](../../models/operations/getdeclarativeconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## DeleteDeclarativeConfig

Delete the declarative config

### Example Usage

<!-- UsageSnippet language="go" operationID="delete-declarative-config" method="delete" path="/v2/control-planes/{controlPlaneId}/declarative-config" -->
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

    res, err := s.DeclarativeConfiguration.DeleteDeclarativeConfig(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.DeleteDeclarativeConfigResponse](../../models/operations/deletedeclarativeconfigresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |

## ValidateControlPlaneConfig

Validates a set of core entities as one configuration for the control plane. Entities refer to each other by ID. The configuration is not stored.

### Example Usage

<!-- UsageSnippet language="go" operationID="validate-control-plane-config" method="post" path="/v2/control-planes/{controlPlaneId}/config/validate" -->
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

    res, err := s.DeclarativeConfiguration.ValidateControlPlaneConfig(ctx, "9524ec7d-36d9-465d-a8c5-83a3c9390458", components.ControlPlaneConfig{
        Services: []components.Service{
            components.Service{
                Host: "example.internal",
                ID: sdkkonnectgo.Pointer("49fd316e-c457-481c-9fc7-8079153e4f3c"),
                Name: sdkkonnectgo.Pointer("example-service"),
                Path: sdkkonnectgo.Pointer("/"),
            },
        },
        Routes: []components.Route{
            components.CreateRouteRouteJSON(
                components.RouteJSON{
                    Hosts: []string{
                        "foo.example.com",
                        "foo.example.us",
                    },
                    ID: sdkkonnectgo.Pointer("56c4566c-14cc-4132-9011-4139fcbbe50a"),
                    Name: sdkkonnectgo.Pointer("example-route"),
                    Paths: []string{
                        "/v1",
                        "/v2",
                    },
                    Service: &components.RouteJSONService{
                        ID: sdkkonnectgo.Pointer("bd380f99-659d-415e-b0e7-72ea05df3218"),
                    },
                },
            ),
        },
        Plugins: []components.Plugin{
            components.Plugin{
                Config: map[string]any{
                    "anonymous": "<value>",
                    "hide_credentials": false,
                    "key_in_body": false,
                    "key_in_header": true,
                    "key_in_query": true,
                    "key_names": []any{
                        "apikey",
                    },
                    "run_on_preflight": true,
                },
                ID: sdkkonnectgo.Pointer("3fd1eea1-885a-4011-b986-289943ff8177"),
                Name: "key-auth",
                Partials: []components.Partials{
                    components.Partials{
                        ID: sdkkonnectgo.Pointer("cff1230a-00f7-4ae8-b376-c370f0eb4dae"),
                        Name: sdkkonnectgo.Pointer("foo-partial"),
                        Path: sdkkonnectgo.Pointer("config.redis"),
                    },
                    components.Partials{
                        ID: sdkkonnectgo.Pointer("129ee345-cba8-4e55-9d6d-93c223ff91ae"),
                        Name: sdkkonnectgo.Pointer("bar-partial"),
                        Path: sdkkonnectgo.Pointer("config.redis"),
                    },
                },
                Protocols: []components.PluginProtocols{
                    components.PluginProtocolsGrpc,
                    components.PluginProtocolsGrpcs,
                    components.PluginProtocolsHTTP,
                    components.PluginProtocolsHTTPS,
                },
            },
        },
        Upstreams: []components.Upstream{
            components.Upstream{
                Healthchecks: &components.Healthchecks{
                    Active: &components.Active{
                        Healthy: &components.Healthy{
                            HTTPStatuses: []int64{
                                200,
                                302,
                            },
                        },
                        Unhealthy: &components.Unhealthy{
                            HTTPStatuses: []int64{
                                429,
                                404,
                                500,
                                501,
                                502,
                                503,
                                504,
                                505,
                            },
                        },
                    },
                    Passive: &components.Passive{
                        Healthy: &components.UpstreamHealthy{
                            HTTPStatuses: []int64{
                                200,
                                201,
                                202,
                                203,
                                204,
                                205,
                                206,
                                207,
                                208,
                                226,
                                300,
                                301,
                                302,
                                303,
                                304,
                                305,
                                306,
                                307,
                                308,
                            },
                        },
                        Unhealthy: &components.UpstreamUnhealthy{
                            HTTPStatuses: []int64{
                                429,
                                500,
                                503,
                            },
                        },
                    },
                },
                ID: sdkkonnectgo.Pointer("6eed5e9c-5398-4026-9a4c-d48f18a2431e"),
                Name: "api.example.internal",
            },
        },
        Targets: []components.Target{
            components.Target{
                ID: sdkkonnectgo.Pointer("089292a7-ba3d-4d88-acf0-97b4b2e2621a"),
                Target: sdkkonnectgo.Pointer("203.0.113.42"),
                Upstream: &components.TargetUpstream{
                    ID: sdkkonnectgo.Pointer("5f1d7e76-2fed-4806-a6af-869984f025cb"),
                },
            },
        },
        Consumers: []components.Consumer{
            components.Consumer{
                CustomID: sdkkonnectgo.Pointer("4200"),
                ID: sdkkonnectgo.Pointer("8a388226-80e8-4027-a486-25e4f7db5d21"),
                Tags: []string{
                    "silver-tier",
                },
                Username: sdkkonnectgo.Pointer("bob-the-builder"),
            },
        },
        Certificates: []components.Certificate{
            components.Certificate{
                Cert: "-----BEGIN CERTIFICATE-----\ncertificate-content\n-----END CERTIFICATE-----",
                ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f269"),
                Key: "-----BEGIN PRIVATE KEY-----\nprivate-key-content\n-----END PRIVATE KEY-----",
            },
        },
        CaCertificates: []components.CACertificate{
            components.CACertificate{
                Cert: "-----BEGIN CERTIFICATE-----\ncertificate-content\n-----END CERTIFICATE-----",
                CertDigest: sdkkonnectgo.Pointer("9b8aaf19a276885f6c8a6bc48a30700fdb3a351d8b05374d153bfb7b178e2a9f"),
                CreatedAt: sdkkonnectgo.Pointer[int64](1706598432),
                ID: sdkkonnectgo.Pointer("b2f34145-0343-41a4-9602-4c69dec2f260"),
                Tags: []string{
                    "trusted",
                    "api",
                },
            },
        },
        Snis: []components.Sni{
            components.Sni{
                Certificate: components.SNICertificate{
                    ID: sdkkonnectgo.Pointer("bd380f99-659d-415e-b0e7-72ea05df3218"),
                },
                ID: sdkkonnectgo.Pointer("36c4566c-14cc-4132-9011-4139fcbbe50a"),
                Name: "some.example.org",
            },
        },
        Vaults: []components.Vault{
            components.Vault{
                Config: map[string]any{
                    "prefix": "ENV_PREFIX",
                },
                Description: sdkkonnectgo.Pointer("environment variable based vault"),
                ID: sdkkonnectgo.Pointer("2747d1e5-8246-4f65-a939-b392f1ee17f8"),
                Name: "env",
                Prefix: "env",
                Tags: []string{
                    "foo",
                    "bar",
                },
            },
        },
        Keys: []components.Key{
            components.Key{
                ID: sdkkonnectgo.Pointer("d958f66b-8e99-44d2-b0b4-edd5bbf24658"),
                Jwk: sdkkonnectgo.Pointer("{\"alg\":\"RSA\",  \"kid\": \"42\",  ...}"),
                Kid: "42",
                Name: sdkkonnectgo.Pointer("a-key"),
                Pem: &components.Pem{
                    PrivateKey: sdkkonnectgo.Pointer("-----BEGIN"),
                    PublicKey: sdkkonnectgo.Pointer("-----BEGIN"),
                },
                Set: &components.Set{
                    ID: sdkkonnectgo.Pointer("b86b331c-dcd0-4b3e-97ce-47c5a9543031"),
                },
            },
        },
        KeySets: []components.KeySet{
            components.KeySet{
                ID: sdkkonnectgo.Pointer("b58c7d9d-e54f-444c-b24d-cdfc4159f61e"),
                Name: sdkkonnectgo.Pointer("example-key-set"),
                Tags: []string{
                    "idp-keys",
                },
            },
        },
        CustomPlugins: []components.CustomPlugin{
            components.CustomPlugin{
                Handler: "return { VERSION = '1.0,0', PRIORITY = 500, access = function(self, config) kong.service.request.set_header(config.name, config.value) end }",
                ID: sdkkonnectgo.Pointer("868346aa-1105-4b77-8346-aa1105fb77c4"),
                Name: "set-header",
                Schema: "return { name = 'set-header', fields = { { protocols = require('kong.db.schema.typedefs').protocols_http }, { config = { type = 'record', fields = { { name = { description = 'The name of the header to set.', type = 'string', required = true } }, { value = { description = 'The value for the header.', type = 'string', required = true } } } } } } }",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ControlPlaneConfigValidationResult != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `controlPlaneID`                                                                   | `string`                                                                           | :heavy_check_mark:                                                                 | The UUID of your control plane. This variable is available in the Konnect manager. | 9524ec7d-36d9-465d-a8c5-83a3c9390458                                               |
| `controlPlaneConfig`                                                               | [components.ControlPlaneConfig](../../models/components/controlplaneconfig.md)     | :heavy_check_mark:                                                                 | N/A                                                                                |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

### Response

**[*operations.ValidateControlPlaneConfigResponse](../../models/operations/validatecontrolplaneconfigresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.UnauthorizedError | 401                         | application/problem+json    |
| sdkerrors.ForbiddenError    | 403                         | application/problem+json    |
| sdkerrors.NotFoundError     | 404                         | application/problem+json    |
| sdkerrors.SDKError          | 4XX, 5XX                    | \*/\*                       |