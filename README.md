# planetopia-protocol

Shared protocol definitions for the Planetopia mesh network.

Used by:
- **motionSensorServer** (Go) — imports as a Go module
- **Planetopia-nodes** (ESP32/C++) — includes via git submodule

## Packages

- **`opcodes/`** — serial command opcode constants (Go)
- **`adapter/`** — adapter type identifiers and helpers (Go)
- **`c/`** — generated C headers for firmware; **do not edit directly**
- **`cmd/gen-headers/`** — generator that writes `c/` from the Go constants

## Usage

### Go (server)

```go
import (
    "github.com/superbrobenji/planetopia-protocol/opcodes"
    "github.com/superbrobenji/planetopia-protocol/adapter"
)

payload[0] = opcodes.OpLEDSolid
if adapter.IsOutput(node.AdapterType) { ... }
```

### C (firmware — via git submodule at `main/lib/planetopia-protocol`)

```c
#include "lib/planetopia-protocol/c/opcodes.h"
#include "lib/planetopia-protocol/c/adapter_types.h"

payload[0] = OP_LED_SOLID;
```

Set up the submodule:

```sh
git submodule update --init
```

## Changing constants

1. Edit the Go source (`opcodes/opcodes.go` or `adapter/types.go`)
2. Regenerate C headers:
   ```sh
   make generate
   # or: go generate ./...
   ```
3. Commit Go source and regenerated `c/` files together
4. Tag a new version and push

The `c/` files carry a `// Code generated; DO NOT EDIT` banner and are
derived entirely from the Go constants — they are never edited by hand.

## CI verification

```sh
make check   # go generate ./... && git diff --exit-code c/
```

Fails if C headers are out of sync with the Go constants.

## Versioning

This module follows semver. Consumers pin to a tag:

| Tag    | Notes |
|--------|-------|
| v0.2.1 | Lower go directive to 1.21.0 (no language features above 1.21 used) |
| v0.2.0 | Generated C headers; submodule support |
| v0.1.0 | Initial release |

<!-- TODO: open-source hardening needed after Planetopia spec is fully implemented.
     Items to address before making this repo public or accepting external contributions:
     - CONTRIBUTING.md + code of conduct
     - LICENSE file (choose OSS licence)
     - Security policy (SECURITY.md, responsible disclosure contact)
     - CI workflow (GitHub Actions: go test, make check, go vet)
     - Branch protection on main
     - Signed tags / release process
     - Review whether opcode/type values should be stabilised before v1.0
-->
