# kongctl compatibility patches for internal SDK v0.7.0

Base: `Kong/sdk-konnect-go-internal` v0.7.0
(`cc8967c36d485c862e86475cd374ea1a0f182431`). Generation uses the release's
OpenAPI input directly, without substituting a development snapshot.

The canonical patch bundle comes from `Kong/kongctl` commit
`e545c18c5b2b4dd431ad4673ea1055d91bb2aeb5`. Its source files and SHA-256
hashes are recorded in `kongctl-patches/provenance.json`.

The patches preserve omitted request defaults, explicit boolean values,
explicit JSON/YAML format selections, and explicit empty identity realms.
This release restores generated auth-config types, so the targeted
`identity_realms` tag patch applies again. HTTP wire regression tests verify
omitted, empty, and populated arrays.

Speakeasy 1.805.0 matches the release's workflow lock. The Makefile and
`.mise/config.toml` retain the generation tooling and hooks from the previous
patched SDK branch; release Go and other tool-version pins are unchanged.
Generation rebuilds SDK code, interfaces, mocks, and DeepCopy helpers.

```sh
mise trust
CGO_ENABLED=0 GOFLAGS=-buildvcs=false mise exec -- make generate.sdk
CGO_ENABLED=0 GOFLAGS=-buildvcs=false go test . ./models/... ./internal/... ./pkg/...
CGO_ENABLED=1 GOFLAGS=-buildvcs=false mise exec -- make test.unit
```

To refresh the bundle, commit it in kongctl first and run
`python3 tools/sdk-patches/apply.py prepare /absolute/path/to/this/checkout`
from kongctl before generation. Consume the resulting SDK commit through an
immutable Go pseudo-version. Ordinary kongctl builds do not apply patches.
