# Kong Identity internal SDK development build

This branch generates Kong Identity management APIs for KongCTL development.
It is consumed by commit; no SDK version bump, tag, release, or new pull
request is part of this build.

## Inputs

- Platform API: `Kong/platform-api@9adee27ed98a74eacc269dbc7e2ba5db96384a61`,
  branch `build/ki-internal-sdk`. This includes `sliver/v1` and
  `kong-identity/v2` in both Go SDK products.
- Internal SDK base: `8108bd7e`, the existing patched v0.7.0 branch based
  on internal SDK main at `cc8967c3`.
- KongCTL patch source:
  `Kong/kongctl@2beef7beeab6b3abd1e0d39978a1c48c203d83dd`,
  `tools/sdk-patches`. File hashes are in `kongctl-patches/provenance.json`.
- Speakeasy: `1.805.0`, pinned by the patch manifest and mise configuration.
- Go: `1.26.10`, inherited from the internal SDK base and pinned by its
  generation Makefile.

The source specification contains the gated directory APIs. Internal SDK
generation uses the complete bundle, so it can include them before the
public visibility changes in `Kong/platform-api#3877` merge. This branch does
not change their public visibility or establish deployed API availability.

## Generate

Use sibling `platform-api` and `sdk-konnect-go-internal` worktrees. The
platform script copies the generated specification into that SDK directory.

From the KongCTL source checkout:

```sh
python3 tools/sdk-patches/apply.py prepare /path/to/sdk-konnect-go-internal
```

From the platform-api worktree at the input commit:

```sh
npm ci
./tools/gen_internal_go_sdk.sh
```

From the prepared internal SDK worktree:

```sh
mise trust
mise install
CGO_ENABLED=0 GOFLAGS=-buildvcs=false mise exec -- make generate.sdk
CGO_ENABLED=0 go test . ./models/... ./internal/... ./pkg/...
```

The generation target preserves the SDK version and skips spec upload. It
applies the KongCTL overlays, regenerates interfaces/mocks/deep-copy helpers,
and applies the targeted post-generation fixes.

The new principal filter introduces an `Active` type. Speakeasy consequently
names the upstream health-check type `UpstreamActive`. This branch's helper
recipe marks that renamed type for deep-copy generation too; otherwise the
generated `Healthchecks` helper does not compile. The existing SDK model
compile/tests exercise this compatibility fix.

## Coverage

The generated input includes authorization servers, clients, scopes, claims,
and client-secret management; directories, principals and identities;
basic-auth/password and API-key credentials; groups and memberships; Token
Vault providers, templates, public keys and credentials; directory Trusted
IdPs; and organization-chart import.

These are typed SDK operations, not KongCTL resource implementations. The
matching, ownership, secret-output, enrollment, and importer boundaries
tracked in `Kong/kongctl#1132` and `Kong/kongctl#1133` still apply.

The module retains the SDK's `github.com/Kong/sdk-konnect-go` package path.
KongCTL can select this private repository through a Go module replacement
or a local replacement while developing resources.

## Validation and consumer compatibility

- Platform API internal spec generation completed successfully.
- All 133 operations under the auth-server, directory, and provider-template
  management routes have generated Go clients.
- SDK root, model, operation, internal, and metadata package tests pass,
  including the installed KongCTL patch regression tests.
- Every SDK package compiles, including generated interfaces and mocks.
- The SDK version remains `0.7.0`; this branch is consumed by commit.

Current KongCTL at the patch-source commit is not yet compatible with the
full internal SDK. A build using a separate local module-replacement file
reports these existing-API shape differences; the same KongCTL source builds
successfully with its original public SDK:

- API implementation service-reference variants and their generated names.
- API-version create/fetch request objects, which include internal query
  parameters absent from the public method signatures.
- Update API force-query enum naming.
- Control-plane name-filter types.
- Portal OIDC identity-provider configuration variants.
- Service list output types.

These require a follow-up KongCTL compatibility branch before changing its
normal SDK dependency. This build does not modify KongCTL source or its
checked-in module files, and it does not implement Kong Identity resources.
No live Konnect resource lifecycle was exercised.
