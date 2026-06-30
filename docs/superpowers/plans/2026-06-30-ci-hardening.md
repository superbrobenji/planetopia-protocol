# CI Hardening Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Harden `ci.yml` and `codeql.yml` to match sibling-repo security standards — SHA-pinned actions, per-job permissions, `go-version-file`, and a new `go-lint` job.

**Architecture:** Single PR touching only `.github/workflows/`. Two files modified; `dependency-review.yml` and `dependabot.yml` already hardened and untouched.

**Tech Stack:** GitHub Actions YAML, `gh` CLI

## Global Constraints

- All action references must use commit SHA + inline version comment: `owner/action@<sha> # vX.Y.Z`
- Every job must declare `permissions: contents: read`
- Use `go-version-file: go.mod` — never hardcode `go-version`
- No `version:` key in `golangci-lint-action`

### SHA Reference Table

| Action | SHA | Version |
|--------|-----|---------|
| `actions/checkout` | `34e114876b0b11c390a56381ad16ebd13914f8d5` | v4.3.1 |
| `actions/setup-go` | `40f1582b2485089dde7abd97c1529aa768e1baff` | v5.6.0 |
| `golangci/golangci-lint-action` | `55c2c1448f86e01eaae002a5a3a9624417608d84` | v6.5.2 |
| `github/codeql-action/init` | `dd903d2e4f5405488e5ef1422510ee31c8b32357` | v3.36.2 |
| `github/codeql-action/autobuild` | `dd903d2e4f5405488e5ef1422510ee31c8b32357` | v3.36.2 |
| `github/codeql-action/analyze` | `dd903d2e4f5405488e5ef1422510ee31c8b32357` | v3.36.2 |

---

## File Map

| File | Action | Responsibility |
|------|--------|---------------|
| `.github/workflows/ci.yml` | Modify | Permissions + SHA pins + `go-version-file` on existing jobs; new `go-lint` job |
| `.github/workflows/codeql.yml` | Modify | SHA pins + `go-version-file` |

---

### Task 1: Create feature branch

**Files:** none

- [ ] **Step 1: Create and switch to branch**

```bash
git checkout -b feat/ci-hardening
```

---

### Task 2: Harden `ci.yml`

**Files:**
- Modify: `.github/workflows/ci.yml`

Add `permissions: contents: read`, SHA-pin both existing jobs, switch `go-version: '1.21'` to `go-version-file: go.mod`, and add a new `go-lint` job. Replace the entire file.

- [ ] **Step 1: Write hardened ci.yml**

```yaml
name: CI

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  go-test:
    name: Test
    runs-on: ubuntu-latest
    permissions:
      contents: read
    steps:
      - uses: actions/checkout@34e114876b0b11c390a56381ad16ebd13914f8d5 # v4.3.1
      - uses: actions/setup-go@40f1582b2485089dde7abd97c1529aa768e1baff # v5.6.0
        with:
          go-version-file: go.mod
      - run: go vet ./...
      - run: go test ./...

  header-sync:
    name: Header sync
    runs-on: ubuntu-latest
    permissions:
      contents: read
    steps:
      - uses: actions/checkout@34e114876b0b11c390a56381ad16ebd13914f8d5 # v4.3.1
      - uses: actions/setup-go@40f1582b2485089dde7abd97c1529aa768e1baff # v5.6.0
        with:
          go-version-file: go.mod
      - name: Check C headers are in sync with Go constants
        run: make check

  go-lint:
    name: Lint
    runs-on: ubuntu-latest
    permissions:
      contents: read
    steps:
      - uses: actions/checkout@34e114876b0b11c390a56381ad16ebd13914f8d5 # v4.3.1
      - uses: actions/setup-go@40f1582b2485089dde7abd97c1529aa768e1baff # v5.6.0
        with:
          go-version-file: go.mod
      - uses: golangci/golangci-lint-action@55c2c1448f86e01eaae002a5a3a9624417608d84 # v6.5.2
```

- [ ] **Step 2: Validate YAML**

```bash
python3 -c "import yaml; yaml.safe_load(open('.github/workflows/ci.yml'))" && echo "VALID"
```

Expected: `VALID`

- [ ] **Step 3: Commit**

```bash
git add .github/workflows/ci.yml
git commit -m "ci: add permissions, SHA-pin actions, go-version-file, add lint job"
```

---

### Task 3: Harden `codeql.yml`

**Files:**
- Modify: `.github/workflows/codeql.yml`

SHA-pin all 5 action references and switch `go-version: '1.21'` to `go-version-file: go.mod`. Existing permissions block is already correct.

- [ ] **Step 1: Write hardened codeql.yml**

```yaml
name: CodeQL

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]
  schedule:
    - cron: '30 1 * * 0'

jobs:
  analyze:
    name: Analyze (go)
    runs-on: ubuntu-latest
    permissions:
      security-events: write
      packages: read
      actions: read
      contents: read
    steps:
      - uses: actions/checkout@34e114876b0b11c390a56381ad16ebd13914f8d5 # v4.3.1
      - uses: actions/setup-go@40f1582b2485089dde7abd97c1529aa768e1baff # v5.6.0
        with:
          go-version-file: go.mod
      - uses: github/codeql-action/init@dd903d2e4f5405488e5ef1422510ee31c8b32357 # v3.36.2
        with:
          languages: go
          queries: security-and-quality
      - uses: github/codeql-action/autobuild@dd903d2e4f5405488e5ef1422510ee31c8b32357 # v3.36.2
      - uses: github/codeql-action/analyze@dd903d2e4f5405488e5ef1422510ee31c8b32357 # v3.36.2
        with:
          category: '/language:go'
```

- [ ] **Step 2: Validate YAML**

```bash
python3 -c "import yaml; yaml.safe_load(open('.github/workflows/codeql.yml'))" && echo "VALID"
```

Expected: `VALID`

- [ ] **Step 3: Commit**

```bash
git add .github/workflows/codeql.yml
git commit -m "ci: SHA-pin codeql-action references and use go-version-file"
```

---

### Task 4: Open PR

**Files:** none

- [ ] **Step 1: Push branch**

```bash
git push -u origin feat/ci-hardening
```

- [ ] **Step 2: Open PR**

```bash
gh pr create \
  --title "ci: harden workflows — SHA pins, permissions, go-version-file, lint job" \
  --body "$(cat <<'EOF'
## Summary

- SHA-pin all action references in `ci.yml` and `codeql.yml`
- Add `permissions: contents: read` to `go-test` and `header-sync` jobs
- Replace `go-version: '1.21'` with `go-version-file: go.mod` in both files
- Add new `go-lint` job (`golangci-lint-action`, SHA-pinned, no version key)

`dependency-review.yml` and `dependabot.yml` already hardened — untouched.

## Test plan

- [ ] All 3 CI jobs pass on this PR (`go-test`, `header-sync`, `go-lint`)
- [ ] CodeQL analysis runs clean
EOF
)"
```

- [ ] **Step 3: Verify all CI jobs pass**

```bash
gh pr checks --watch
```

Expected: all checks pass (green).
