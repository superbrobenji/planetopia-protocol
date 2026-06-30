# Open Source Hardening — Design Spec

**Date:** 2026-06-30
**Status:** Approved

## Goal

Prepare the `planetopia-protocol` repo for public open-source release.

## Decisions Made

| Decision | Choice | Reason |
|----------|--------|--------|
| License | GPL v3 | Copyleft — matches sibling repos (planetopia-nodes, motionSensorServer) |
| Track structure | 3 tracks, 3 PRs | Small repo; CI and Docs are coupled via CI badge in README |
| CI header check | `make check` (existing) | Already implemented — just needs a CI wrapper |
| dependency-review | Skip | No external Go dependencies |
| CodeQL | Add | Protocol dictates behaviour of all other repos; security-sensitive |
| README | Full rewrite | Remove TODO comment, add badges and license section, tighten structure |

## Scope

Three tracks, executed in order. Each track = one independent PR.

### Track A — Legal & Security *(blocks public release)*

1. **`LICENSE`** — GPL v3 full text. Copyright: `Copyright (C) 2026 Planetopia Contributors`
2. **`SECURITY.md`** — responsible disclosure policy:
   - Supported: latest `main` only
   - Report via GitHub Security Advisory (private) or maintainer email
   - 72h acknowledgement, credited fix
   - Out of scope: upstream Go stdlib vulnerabilities, physical hardware access to ESP32 nodes

### Track B — Community Health

3. **`CODE_OF_CONDUCT.md`** — Contributor Covenant 2.1
4. **`CONTRIBUTING.md`** — full protocol contribution workflow:
   - Adding an opcode: edit `opcodes/opcodes.go`
   - Adding an adapter type: edit `adapter/types.go`
   - Regenerate C headers: `make generate`
   - Verify sync: `make check`
   - Run tests: `go test ./...` and `go vet ./...`
   - Semver rules: patch for additive changes, minor for breaking opcode/type changes
   - Commit Go source and regenerated `c/` files together in one commit
   - Tag format (`vX.Y.Z`), push tag, update versioning table in `README.md`
   - PR process and what blocks merge
5. **`.github/pull_request_template.md`**:
   - Summary
   - Type of change checkboxes: opcode add / adapter type add / bug fix / docs / CI
   - Checklist: C headers regenerated (`make generate`), `make check` passes, `go test ./...` passes, version table updated if constants changed
6. **`.github/ISSUE_TEMPLATE/bug_report.md`** — standard bug report template
7. **`.github/ISSUE_TEMPLATE/feature_request.md`** — covers both feature requests and new opcode/adapter type proposals

### Track C — CI + Docs

8. **`.github/workflows/ci.yml`** — triggers on push to `main` and PR to `main`:

   | Job | Steps |
   |-----|-------|
   | `go-test` | `go test ./...` + `go vet ./...` in repo root |
   | `header-sync` | `make check` — runs `go generate ./...` then `git diff --exit-code c/`; fails if C headers drift from Go constants |

   Both jobs must pass before merge.

9. **`.github/workflows/codeql.yml`** — language: `go`, triggers: push to `main`, PRs to `main`, weekly schedule (standard CodeQL query suite)

10. **Branch protection (manual step in GitHub settings)**:
    - Require both CI jobs green before merge
    - Require 1 approving review
    - Dismiss stale reviews on new commits
    - No direct push to `main`

11. **`README.md`** — full rewrite:
    - Badges: CI status, license (GPL v3)
    - One-paragraph project description
    - Packages overview (`opcodes/`, `adapter/`, `c/`, `cmd/gen-headers/`)
    - Usage examples (Go module import, C submodule include)
    - Changing constants workflow (condensed — full detail in CONTRIBUTING.md)
    - Versioning table (retain existing, update going forward)
    - License section
    - Remove `<!-- TODO: open-source hardening needed -->` comment

## Out of Scope

- No source code changes (Go constants, C header generator logic)
- No dependency-review workflow (no external Go dependencies)
- No CHANGELOG.md (no prior versioned releases to document)
- No Dependabot (no package manager dependencies to update)

## Success Criteria

- `LICENSE` present, GPL v3
- `SECURITY.md` present with private disclosure path (GitHub Security Advisory)
- All CI jobs green on first run: `go-test`, `header-sync`, `codeql`
- `README.md` has CI and license badges, no `<!-- TODO: open-source hardening -->` comment
- `CONTRIBUTING.md` covers full protocol contribution workflow: edit Go source, `make generate`, `make check`, `go test ./...`, semver bump, versioning table update
- Branch protection rules documented for manual configuration
