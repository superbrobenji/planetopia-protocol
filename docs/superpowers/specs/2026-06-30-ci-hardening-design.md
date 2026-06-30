# CI Hardening Design

**Date:** 2026-06-30
**Status:** Approved

## Goal

Bring planetopia-protocol CI to the same security standard as the sibling repo `planetopia-nodes`. All changes are confined to `.github/workflows/`.

## Scope

Two files modified, none created:

| File | Change |
|------|--------|
| `.github/workflows/ci.yml` | Permissions + SHA pins + `go-version-file` + new golangci-lint job |
| `.github/workflows/codeql.yml` | SHA pins + `go-version-file` |
| `.github/workflows/dependency-review.yml` | Already hardened — no changes |
| `.github/dependabot.yml` | Already complete — no changes |

## 1. `ci.yml`

### Permissions

Add `permissions: contents: read` to both existing jobs (`go-test`, `header-sync`).

### SHA-pinned actions

Replace mutable version tags with commit SHAs + inline version comment in both jobs:

| Action | Current | Pinned |
|--------|---------|--------|
| `actions/checkout` | `@v4` | `34e114876b0b11c390a56381ad16ebd13914f8d5 # v4.3.1` |
| `actions/setup-go` | `@v5` | SHA resolved at implementation time via `gh api` |

### `go-version-file`

Replace `go-version: '1.21'` with `go-version-file: go.mod` in both jobs. The module declares `go 1.21.0` at root — this keeps the runner version in sync automatically as the module is updated.

### New `go-lint` job

```
name: Go lint
runs-on: ubuntu-latest
permissions:
  contents: read
steps:
  - actions/checkout (SHA-pinned)
  - actions/setup-go (SHA-pinned, go-version-file: go.mod)
  - golangci/golangci-lint-action (SHA-pinned, no version key)
```

Matches `go-lint` job in motionSensorServer `ci.yml` exactly. Working directory is root (go.mod at root). No `.golangci.yml` config introduced — action bundled defaults apply.

## 2. `codeql.yml`

### SHA-pinned actions

All 5 unpinned actions replaced:

| Action | Current | Pinned |
|--------|---------|--------|
| `actions/checkout` | `@v4` | `34e114876b0b11c390a56381ad16ebd13914f8d5 # v4.3.1` |
| `actions/setup-go` | `@v5` | SHA resolved at implementation time |
| `github/codeql-action/init` | `@v3` | `dd903d2e4f5405488e5ef1422510ee31c8b32357 # v3.36.2` |
| `github/codeql-action/autobuild` | `@v3` | `dd903d2e4f5405488e5ef1422510ee31c8b32357 # v3.36.2` |
| `github/codeql-action/analyze` | `@v3` | `dd903d2e4f5405488e5ef1422510ee31c8b32357 # v3.36.2` |

codeql-action SHAs sourced from planetopia-nodes `codeql.yml` (already verified at v3.36.2).

### `go-version-file`

Replace `go-version: '1.21'` with `go-version-file: go.mod`. Existing permissions block already correct — no change.

## Non-goals

- No changes to `dependency-review.yml` (already SHA-pinned at v4.3.1 + v5.0.0)
- No changes to `dependabot.yml` (already covers github-actions + gomod)
- No `.golangci.yml` config introduced
- No changes to job triggers or branch filters
- No application code changes

## Delivery

Single PR. All changes in `.github/workflows/` only.
