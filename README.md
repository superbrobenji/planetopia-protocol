# planetopia-protocol

Shared protocol definitions for the Planetopia mesh network.

## Packages

- **`opcodes`** — serial command opcode constants (Go)
- **`adapter`** — adapter type identifiers and helpers (Go)
- **`c/`** — equivalent C headers for firmware (ESP32/Arduino)

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

### C (firmware)

```c
#include "planetopia-protocol/opcodes.h"
#include "planetopia-protocol/adapter_types.h"

payload[0] = OP_LED_SOLID;
```

## Sync rule

Go constants and C `#define`s must always match. A change to one requires a change to the other in the same commit.
