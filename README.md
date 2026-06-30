# planetopia-protocol

[![CI](https://github.com/superbrobenji/planetopia-protocol/actions/workflows/ci.yml/badge.svg)](https://github.com/superbrobenji/planetopia-protocol/actions/workflows/ci.yml)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Shared protocol definitions for the Planetopia mesh network. Defines the opcode and adapter-type constants consumed by all Planetopia services and firmware.

Used by:
- **motionSensorServer** (Go) — imports as a Go module
- **Planetopia-nodes** (ESP32/C++) — includes via git submodule

## Packages

| Package | Description |
|---------|-------------|
| `opcodes/` | Serial command opcode constants (Go) |
| `adapter/` | Adapter type identifiers and helpers (Go) |
| `c/` | Generated C headers for firmware — do not edit directly |
| `cmd/gen-headers/` | Generator that writes `c/` from the Go constants |

## Usage

### Go (motionSensorServer)

```go
import (
    "github.com/superbrobenji/planetopia-protocol/opcodes"
    "github.com/superbrobenji/planetopia-protocol/adapter"
)

payload[0] = opcodes.OpLEDSolid
if adapter.IsOutput(node.AdapterType) { ... }
```

### C (Planetopia-nodes — via git submodule at `main/lib/planetopia-protocol`)

```c
#include "lib/planetopia-protocol/c/opcodes.h"
#include "lib/planetopia-protocol/c/adapter_types.h"

payload[0] = OP_LED_SOLID;
```

Set up the submodule:

```sh
git submodule update --init
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for the full protocol contribution workflow — adding opcodes or adapter types, regenerating C headers, and opening a PR.

## Versioning

This module follows semver. Consumers pin to a tag.

| Tag | Notes |
|-----|-------|
| v0.2.1 | Lower go directive to 1.21.0 |
| v0.2.0 | Generated C headers; submodule support |
| v0.1.0 | Initial release |

## License

Copyright (C) 2026 Planetopia Contributors

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

See [LICENSE](LICENSE) for the full terms.
