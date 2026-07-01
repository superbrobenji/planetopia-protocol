# Platform Rename: Planetopia → Lattice

**Date:** 2026-07-01  
**Status:** Approved

## Overview

Rename the Lattice platform (formerly Planetopia) across all 3 repos. Every instance of `planetopia`/`Planetopia` becomes `lattice`/`Lattice`. Repo names updated to clearly communicate each repo's role.

## Repos

| Old name | New name | Role |
|---|---|---|
| `motionSensorServer` | `lattice-hub` | Go orchestrator + Kafka + React dashboards |
| `planetopia-nodes` | `lattice-nodes` | ESP32 C++ mesh firmware |
| `planetopia-protocol` | `lattice-protocol` | Shared opcodes/adapter types (Go + generated C headers) |

## Execution Order

Dependency order matters: `lattice-hub` and `lattice-nodes` both depend on `lattice-protocol`.

1. Rename all 3 repos on GitHub (Settings → Rename) — redirects activate immediately
2. PR on `lattice-protocol` — rename internals, merge first
3. PR on `lattice-hub` + PR on `lattice-nodes` — parallel, both consume updated protocol after step 2

## Per-Repo Change Details

### lattice-protocol

| What | Old | New |
|---|---|---|
| `go.mod` module path | `github.com/superbrobenji/planetopia-protocol` | `github.com/superbrobenji/lattice-protocol` |
| C header guard (opcodes) | `PLANETOPIA_OPCODES_H` | `LATTICE_OPCODES_H` |
| C header guard (adapter types) | `PLANETOPIA_ADAPTER_TYPES_H` | `LATTICE_ADAPTER_TYPES_H` |
| Copyright strings | `Planetopia Contributors` | `Lattice Contributors` |
| All comments/strings | `planetopia`/`Planetopia` | `lattice`/`Lattice` |
| CI workflow files | any `planetopia` refs | `lattice` |

Note: Go packages (`opcodes`, `adapter`) need no rename — already non-branded.

### lattice-nodes

| What | Old | New |
|---|---|---|
| All C++ namespaces | `planetopia::mesh`, `planetopia::adapter`, `planetopia::error`, `planetopia::utils`, `planetopia::hardware`, `planetopia::persistence`, `planetopia::config` | `lattice::*` |
| Namespace declarations | `namespace planetopia {` | `namespace lattice {` |
| Submodule path | `lib/planetopia-protocol` | `lib/lattice-protocol` |
| `.gitmodules` URL | `...planetopia-protocol.git` | `...lattice-protocol.git` |
| `#include` paths referencing submodule | `"planetopia-protocol/..."` | `"lattice-protocol/..."` |
| C/C++ header include guards | `PLANETOPIA_*` | `LATTICE_*` |
| All comments/strings | `planetopia`/`Planetopia` | `lattice`/`Lattice` |
| CMakeLists / build config | any `planetopia` refs | `lattice` |

### lattice-hub

| What | Old | New |
|---|---|---|
| `go.mod` module path | `github.com/superbrobenji/motionServer` | `github.com/superbrobenji/lattice-hub` |
| All internal Go imports | `github.com/superbrobenji/motionServer/...` | `github.com/superbrobenji/lattice-hub/...` |
| Protocol dependency import | `github.com/superbrobenji/planetopia-protocol` | `github.com/superbrobenji/lattice-protocol` |
| `go.sum` | stale entries | regenerated via `go mod tidy` |
| All comments/strings | `planetopia`/`Planetopia` | `lattice`/`Lattice` |
| Docker/compose service names | any `planetopia` refs | `lattice` |
| CI workflow files | any `planetopia` refs | `lattice` |

## Out of Scope

- Internal Go package names (`mesh`, `adapter`, `eventStore`, etc.) — not branded, no change needed
- React component names and UI copy — not Planetopia-branded
- Kafka topic names — runtime config, not source code branding
- `.env` / secrets files
