# Platform Rename: Lattice → Lattice

**Date:** 2026-07-01  
**Status:** Approved

## Overview

Rename the Lattice platform (formerly Lattice) across all 3 repos. Every instance of `lattice`/`Lattice` becomes `lattice`/`Lattice`. Repo names updated to clearly communicate each repo's role.

## Repos

| Old name | New name | Role |
|---|---|---|
| `motionSensorServer` | `lattice-hub` | Go orchestrator + Kafka + React dashboards |
| `lattice-nodes` | `lattice-nodes` | ESP32 C++ mesh firmware |
| `lattice-protocol` | `lattice-protocol` | Shared opcodes/adapter types (Go + generated C headers) |

## Execution Order

Dependency order matters: `lattice-hub` and `lattice-nodes` both depend on `lattice-protocol`.

1. Rename all 3 repos on GitHub (Settings → Rename) — redirects activate immediately
2. PR on `lattice-protocol` — rename internals, merge first
3. PR on `lattice-hub` + PR on `lattice-nodes` — parallel, both consume updated protocol after step 2

## Per-Repo Change Details

### lattice-protocol

| What | Old | New |
|---|---|---|
| `go.mod` module path | `github.com/superbrobenji/lattice-protocol` | `github.com/superbrobenji/lattice-protocol` |
| C header guard (opcodes) | `LATTICE_OPCODES_H` | `LATTICE_OPCODES_H` |
| C header guard (adapter types) | `LATTICE_ADAPTER_TYPES_H` | `LATTICE_ADAPTER_TYPES_H` |
| Copyright strings | `Lattice Contributors` | `Lattice Contributors` |
| All comments/strings | `lattice`/`Lattice` | `lattice`/`Lattice` |
| CI workflow files | any `lattice` refs | `lattice` |

Note: Go packages (`opcodes`, `adapter`) need no rename — already non-branded.

### lattice-nodes

| What | Old | New |
|---|---|---|
| All C++ namespaces | `lattice::mesh`, `lattice::adapter`, `lattice::error`, `lattice::utils`, `lattice::hardware`, `lattice::persistence`, `lattice::config` | `lattice::*` |
| Namespace declarations | `namespace lattice {` | `namespace lattice {` |
| Submodule path | `lib/lattice-protocol` | `lib/lattice-protocol` |
| `.gitmodules` URL | `...lattice-protocol.git` | `...lattice-protocol.git` |
| `#include` paths referencing submodule | `"lattice-protocol/..."` | `"lattice-protocol/..."` |
| C/C++ header include guards | `LATTICE_*` | `LATTICE_*` |
| All comments/strings | `lattice`/`Lattice` | `lattice`/`Lattice` |
| CMakeLists / build config | any `lattice` refs | `lattice` |

### lattice-hub

| What | Old | New |
|---|---|---|
| `go.mod` module path | `github.com/superbrobenji/motionServer` | `github.com/superbrobenji/lattice-hub` |
| All internal Go imports | `github.com/superbrobenji/motionServer/...` | `github.com/superbrobenji/lattice-hub/...` |
| Protocol dependency import | `github.com/superbrobenji/lattice-protocol` | `github.com/superbrobenji/lattice-protocol` |
| `go.sum` | stale entries | regenerated via `go mod tidy` |
| All comments/strings | `lattice`/`Lattice` | `lattice`/`Lattice` |
| Docker/compose service names | any `lattice` refs | `lattice` |
| CI workflow files | any `lattice` refs | `lattice` |

## Out of Scope

- Internal Go package names (`mesh`, `adapter`, `eventStore`, etc.) — not branded, no change needed
- React component names and UI copy — not Lattice-branded
- Kafka topic names — runtime config, not source code branding
- `.env` / secrets files
