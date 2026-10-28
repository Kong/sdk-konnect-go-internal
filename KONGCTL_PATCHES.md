# kongctl compatibility patches for internal SDK v0.6.0

Base: `Kong/sdk-konnect-go-internal` v0.6.0
(`20adb91c19b5ac31b1df0a29a2bbd6550b065e1f`). The release's OpenAPI input
includes AI Gateway runtime 2.3, wildcard formats/capabilities, Token Vault,
MCP rate limiting, outbound TLS, and OIDC credential claims. No replacement
API snapshot is used.

The canonical patch bundle comes from `Kong/kongctl` commit
`60676c3acd19a8d4331949d65e171d181f2108cc`. Its files and SHA-256 hashes are
recorded in `kongctl-patches/provenance.json`.

The patches preserve omitted request defaults, explicit boolean values,
explicit JSON/YAML format selections, and explicit empty identity realms.
The release's flexible auth configs preserve advanced plugin properties
through `additionalProperties`, so the legacy generated-config tag patch is
not needed. The identity-realms wire regression still runs.

Generation uses Speakeasy 1.805.0, matching the release's workflow lock. The
Makefile uses the existing modern tool-installation and patch hooks from the
previous kongctl SDK branch, retains the release's Go toolchain pin, and
rebuilds generated interfaces, mocks, and DeepCopy helpers. Tool versions from
the release are retained. A single `.mise/config.toml` pins Speakeasy.

From this prepared checkout:

```sh
mise trust
CGO_ENABLED=0 GOFLAGS=-buildvcs=false mise exec -- make generate.sdk
CGO_ENABLED=0 GOFLAGS=-buildvcs=false go test . ./models/... ./internal/... ./pkg/...
CGO_ENABLED=1 GOFLAGS=-buildvcs=false mise exec -- make test.unit
```

To refresh the patch bundle, commit it in kongctl first, then run its
`tools/sdk-patches/apply.py prepare` command against this checkout before
generation. Consume the resulting SDK commit through an immutable Go
pseudo-version; ordinary kongctl builds do not apply patches.
