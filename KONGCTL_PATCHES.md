# AI Gateway 2.3 development SDK

This branch regenerates internal SDK v0.5.1 (`7af4b170`) from
Kong/platform-api `224e6d7cf228f803fbcef570e5d53cf5337401f0` and applies
kongctl's canonical compatibility patch bundle from
`f4b147145071f89bac44c1ab726a7b0727bf5721`.

The AI Gateway source specification is 2.2.24. Its public runtime cutoff
remains 2.2; the unsplit internal SDK input includes the currently merged
2.3 features without promoting their public visibility. This snapshot includes
wildcard formats/capabilities, which remain dev-only enum values.
Their defaults are still disabled upstream, so formats and capabilities
remain explicit. The proposed additional MCP CEL and policy expressions
fields are not included. SDK presence does not establish backend deployment
or data-plane support.

## Input preparation

Use separate, sibling worktrees named `platform-api` and
`sdk-konnect-go-internal`. In the platform-api worktree at the commit above:

```sh
npm ci
bash tools/gen_internal_go_sdk.sh
```

The pinned platform-api commit includes the merged wildcard changes from
PR #3827 and the auth-generation fixes from draft PR #3845. Those fixes
restore concrete key-auth/OIDC configs, retain `credential_claims`, and
apply Go-scoped annotations to the internal SDK. No local platform-api
preparation patch is needed. PR #3845 is not yet merged into platform-api
main; reproduce this snapshot from the pinned commit above.
All selected internal SDK specifications are refreshed from that platform-api
snapshot; the input is not limited to AI Gateway changes.

The internal input includes Baseten provider/target variants, Bedrock
`vpc_endpoint` and `endpoint_type` (`runtime` or `mantle`, without a default),
expanded failover enums, Token Vault, upstream certificate/CA references,
MCP rate limiting, OIDC credential claims, and cost-tag capture operations.
The datastore discriminator is `token-vault`; older Jira examples using
`token-vault-konnect` are not this specification's wire format.

Cost-tag capture is part of the AI Gateway source specification and is
regional organization configuration, not a gateway child resource. The
separate AI Cost Management specification is not added to this SDK.

## Generation

The canonical patch bundle is snapshotted in `kongctl-patches/`, including
its source commit and SHA-256 hashes in `provenance.json`. It removes
automatic defaults from sparse requests and preserves explicit empty
`identity_realms` arrays. The audit-log webhook PATCH overlay removes all
property defaults, including the internal `include_principal_name` default;
an empty PATCH stays empty. Its regression test is part of the shared bundle.
The team membership overlay also covers internal organization-team create and
update requests, preserving omitted `konnect_managed` fields while retaining
explicit boolean values. It leaves those fields absent from public SDK input.
The optional internal `CreatePortal.create_default_content` field is patched
after generation to preserve omission without changing other create defaults.
The replacement skips public SDKs where the field is absent and rejects
unexpected tags or types when present.
Internal API-version create, update, and fetch requests also preserve omitted
`spec_format` query parameters while retaining explicit JSON/YAML selections.
No kongctl build applies these patches at runtime.

The SDK Makefile's tool installation and Speakeasy generation targets are
backported from kongctl's public SDK baseline `2314d4a9ddd8`. The internal
SDK's generator configuration and module path are retained. Speakeasy is
pinned to the canonical bundle's version, 1.802.0.
Generation normalizes trailing whitespace in the OpenAPI input.
Generation removes stale DeepCopy helpers, interfaces, and mocks before
Speakeasy compiles the new models, then rebuilds those helpers afterward.
Interface cleanup includes groups removed from the refreshed API input.

From the SDK worktree:

```sh
mise trust
mise install
CGO_ENABLED=0 GOFLAGS=-buildvcs=false mise exec -- make generate.sdk
```

To replace the compatibility bundle, run its preparation command from the
clean kongctl source checkout before SDK generation:

```sh
python3 tools/sdk-patches/apply.py prepare /absolute/path/to/sdk-checkout
```

## Validation

```sh
CGO_ENABLED=0 GOFLAGS=-buildvcs=false go test ./...
CGO_ENABLED=1 GOFLAGS=-buildvcs=false mise exec -- make test.unit
```

Live integration tests and kongctl dependency or implementation changes are
outside this SDK preparation. Consume the branch by an immutable commit
pseudo-version using the existing module path `github.com/Kong/sdk-konnect-go`.

The full SDK test suite and internal unit tests with the race detector pass.
Generation also compiles the SDK and runs staticcheck. Focused tests cover
the 2.3 request fields, omission of an unset Bedrock endpoint type, and the
regional cost-tag GET/PUT paths and request payload.
Wildcard tests verify that `*` is a known format and capability enum value
and survives model/API create and update request serialization.
