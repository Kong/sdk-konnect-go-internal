# kongctl AI Gateway 2.2 development SDK

Baseline: internal SDK `v0.5.0` (`01e246f3`). Its bundled internal OpenAPI
document records AI Gateway `2.1.16`. The feature reference is `2.1.15`;
`2.1.16` only updates minimum-runtime-version documentation. The internal
OpenAPI input is unchanged, including its visibility-gated API surfaces.
The converter reference is `Kong/ai-deck-converter` `v0.19.0`.

## Patch provenance

kongctl's public SDK dependency was
`v0.68.2-0.20260922202403-4d68939eb6fd`, from
`patch/v0.68.1/targeted-patch-defaults`. The complete delta from `v0.68.1`
spans `16ace680`, `7e77ff86`, and `4d68939e`; the final commit alone does
not contain all patches.

All seven targeted overlays remain necessary on the internal baseline:

- Backend-cluster TLS `insecure_skip_verify` omission.
- Custom portal email-template PATCH `enabled` omission.
- Portal team create/update `konnect_managed` omission.
- HTTP DCR PATCH `allow_multiple_credentials` omission.
- Portal audit-webhook PATCH `enabled` omission.
- Organization and portal identity-provider PATCH `enabled` omission.
- Portal PATCH boolean omission.

Compatibility testing additionally requires removing the default from
audit-webhook PATCH `include_principal_name` and from organization-team
create/update `konnect_managed`. The latter also prevents an embedded SDK
JSON unmarshaller from consuming kongctl's sibling resource metadata.
These additions retain the fields and explicitly supplied values.

The public SDK's upstream `RouteJSON.strip_path` post-generation fix is
absent from internal `v0.5.0` and is ported through the Makefile. The
`GroupMembership.members` post-generation fix is already present.

The public branch's six handwritten JSON wrappers are unnecessary here:
Speakeasy `1.799.0` removes the default-aware methods when the overlays
remove the relevant defaults. Sparse-request regression tests cover this.
Keep the internal generator configuration, including error-model settings;
copying the public configuration would change unrelated internal surfaces.

## Reproduction and checks

```sh
mise trust
mise install
mise exec -- make generate.sdk
go test ./...
make test.unit
```

Generation is pinned to the baseline lockfile's Speakeasy `1.799.0`, not
the older `1.798.0` in the public patched SDK. The workflow strictly
validates overlays, regenerates SDK code and DeepCopy methods, and applies
the existing post-generation fixes. Generation also builds, compiles tests,
and runs staticcheck. Generated comment/documentation formatting changes
are unrelated to API behavior; no internal API types are removed.

The Makefile's `test.integration` target refers to `test/integration`,
which is absent from this internal baseline. It cannot provide backend
validation. The custom-policy CRUD regression uses a mock HTTP client;
endpoint presence in the SDK does not imply availability in dev.

## kongctl compatibility

The replacement retains module `github.com/Kong/sdk-konnect-go`. kongctl
needs mechanical adaptations for generated names in API implementations,
control-plane filters, auth-strategy filters, portal menu visibility,
Event Gateway enums, and test fixtures. API-version create/fetch calls
must use operation request structs. Existing wrappers can preserve their
kongctl-facing signatures. Explain schemas and the reviewed dump-default
inventory must also account for additional internal SDK fields.

Policy configuration remains `map[string]any`; prompt-compressor,
rate-limiting, and OpenTelemetry options need no generated typed fields.
Passthrough formats, skills, multiple selector values, Typesafe provider
and target configuration, decisions, and custom-policy CRUD are retained
and covered by focused regression tests.

For remote Go resolution, configure authenticated Git access to the private
repository and include `github.com/Kong/sdk-konnect-go-internal` in
`GOPRIVATE` and, if explicitly configured, `GONOSUMDB`/`GONOPROXY`.
kongctl's `scripts/setup-private-sdk.sh` accepts `GH_TOKEN_PRIVATE_READ`
(or `GH_PRIVATE_READ_TOKEN`), clones the pinned commit, and installs a
temporary local replacement. CI already supports this private SDK flow.
