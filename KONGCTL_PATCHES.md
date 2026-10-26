# AI Gateway 2.3 development SDK

This branch regenerates internal SDK v0.5.1 (`7af4b170`) from
Kong/platform-api `c3595a0837f61a9e06f7856101420c0ad42244f6` and applies
kongctl's canonical compatibility patch bundle from
`fdded07792e5a08f6214c0e040692242fe3a8b93`.

The AI Gateway source specification is 2.2.23. Its public runtime cutoff
remains 2.2; the unsplit internal SDK input includes the currently merged
2.3 features without promoting their public visibility. This snapshot does
not include unmerged wildcard formats/capabilities or the proposed additional
MCP CEL and policy expressions fields. SDK presence does not establish
backend deployment or data-plane support.

## Input preparation

Use separate, sibling worktrees named `platform-api` and
`sdk-konnect-go-internal`. In the platform-api worktree at the commit above:

```sh
git apply ../sdk-konnect-go-internal/.speakeasy/platform-api-aigw-2.3.patch
npm ci
bash tools/gen_internal_go_sdk.sh
```

The recorded patch makes the internal input use the public SDK's concrete
generated key-auth and OpenID Connect configurations. It also retains the
new nested-array `credential_claims` field that is missing from the older
generated OIDC configuration, applies Go-scoped annotations to the internal
SDK, and makes the internal generation entry point fail on errors.
The generated input also has trailing whitespace removed.
These changes are local preparation patches, not an upstream API release.
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
`identity_realms` arrays. No kongctl build applies these patches at runtime.

The SDK Makefile's tool installation and Speakeasy generation targets are
backported from kongctl's public SDK baseline `2314d4a9ddd8`. The internal
SDK's generator configuration and module path are retained. Speakeasy is
pinned to the canonical bundle's version, 1.802.0.
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
