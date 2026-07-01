# OSS Hardening Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Prepare `lattice-protocol` for public open-source release across three independently-mergeable tracks.

**Architecture:** Three sequential tracks (Legal → Community → CI+Docs), each merged as its own PR. Track A blocks public release; B and C build on top.

**Tech Stack:** GitHub Actions (Go 1.23 toolchain), GPL v3, Contributor Covenant 2.1, CodeQL (Go)

## Global Constraints

- License: GPL v3. Copyright line: `Copyright (C) 2026 Lattice Contributors`
- Repo slug: `github.com/superbrobenji/lattice-protocol`
- Go minimum: `1.21.0` (go.mod); use `go-version: '1.23'` in CI toolchain
- All CI jobs must be green on first run
- No external Go dependencies — keep go.mod clean
- Each track ends with its own PR

---

## Track A — Legal & Security

### Task 1: LICENSE

**Files:**
- Create: `LICENSE`

- [ ] **Step 1: Fetch GPL v3 text**

```bash
curl -o LICENSE https://www.gnu.org/licenses/gpl-3.0.txt
```

- [ ] **Step 2: Verify**

```bash
head -3 LICENSE
```

Expected output:
```
                    GNU GENERAL PUBLIC LICENSE
                       Version 3, 29 June 2007
```

- [ ] **Step 3: Commit**

```bash
git add LICENSE
git commit -m "docs: add GPL v3 license"
```

---

### Task 2: SECURITY.md

**Files:**
- Create: `SECURITY.md`

- [ ] **Step 1: Create SECURITY.md**

Replace `[maintainer-email]` with the actual contact address before committing.

```markdown
# Security Policy

## Supported Versions

Only the latest commit on `main` is supported. No backports are provided.

## Reporting a Vulnerability

**Do not open a public GitHub issue for security vulnerabilities.**

Report vulnerabilities via one of:

1. **GitHub Security Advisory (preferred):** Use the [private vulnerability reporting](../../security/advisories/new) link in this repository.
2. **Email:** Contact the maintainer directly at `[maintainer-email]`.

You will receive an acknowledgement within 72 hours. Confirmed vulnerabilities will be credited in the fix.

## Out of Scope

- Vulnerabilities in the Go standard library (report those to the Go team)
- Physical hardware access to ESP32 nodes consuming this protocol
```

- [ ] **Step 2: Commit**

```bash
git add SECURITY.md
git commit -m "docs: add security policy"
```

- [ ] **Step 3: Open PR for Track A**

Open a PR titled `chore: OSS hardening — Track A (legal & security)` targeting `main`. Two commits: LICENSE and SECURITY.md.

---

## Track B — Community Health

### Task 3: CODE_OF_CONDUCT.md

**Files:**
- Create: `CODE_OF_CONDUCT.md`

- [ ] **Step 1: Create CODE_OF_CONDUCT.md**

```markdown
# Contributor Covenant Code of Conduct

## Our Pledge

We as members, contributors, and leaders pledge to make participation in our community a harassment-free experience for everyone, regardless of age, body size, visible or invisible disability, ethnicity, sex characteristics, gender identity and expression, level of experience, education, socio-economic status, nationality, personal appearance, race, caste, color, religion, or sexual identity and orientation.

We pledge to act and interact in ways that contribute to an open, welcoming, diverse, inclusive, and healthy community.

## Our Standards

Examples of behavior that contributes to a positive environment for our community include:

* Demonstrating empathy and kindness toward other people
* Being respectful of differing opinions, viewpoints, and experiences
* Giving and gracefully accepting constructive feedback
* Accepting responsibility and apologizing to those affected by our mistakes, and learning from the experience
* Focusing on what is best not just for us as individuals, but for the overall community

Examples of unacceptable behavior include:

* The use of sexualized language or imagery, and sexual attention or advances of any kind
* Trolling, insulting or derogatory comments, and personal or political attacks
* Public or private harassment
* Publishing others' private information, such as a physical or electronic address, without their explicit permission
* Other conduct which could reasonably be considered inappropriate in a professional setting

## Enforcement Responsibilities

Community leaders are responsible for clarifying and enforcing our standards of acceptable behavior and will take appropriate and fair corrective action in response to any behavior that they find inappropriate, threatening, offensive, or harmful.

Community leaders have the right and responsibility to remove, edit, or reject comments, commits, code, wiki edits, issues, and other contributions that are not aligned to this Code of Conduct, and will communicate reasons for moderation decisions when appropriate.

## Scope

This Code of Conduct applies within all community spaces, and also applies when an individual is officially representing the community in public spaces.

## Enforcement

Instances of abusive, harassing, or otherwise unacceptable behavior may be reported to the community leaders responsible for enforcement. All complaints will be reviewed and investigated promptly and fairly.

All community leaders are obligated to respect the privacy and security of the reporter of any incident.

## Enforcement Guidelines

### 1. Correction

**Community Impact**: Use of inappropriate language or other behavior deemed unprofessional or unwelcome in the community.

**Consequence**: A private, written warning from community leaders, providing clarity around the nature of the violation and an explanation of why the behavior was inappropriate. A public apology may be requested.

### 2. Warning

**Community Impact**: A violation through a single incident or series of actions.

**Consequence**: A warning with consequences for continued behavior. No interaction with the people involved, including unsolicited interaction with those enforcing the Code of Conduct, for a specified period of time. Violating these terms may lead to a temporary or permanent ban.

### 3. Temporary Ban

**Community Impact**: A serious violation of community standards, including sustained inappropriate behavior.

**Consequence**: A temporary ban from any sort of interaction or public communication with the community for a specified period of time.

### 4. Permanent Ban

**Community Impact**: Demonstrating a pattern of violation of community standards, including sustained inappropriate behavior, harassment of an individual, or aggression toward or disparagement of classes of individuals.

**Consequence**: A permanent ban from any sort of public interaction within the community.

## Attribution

This Code of Conduct is adapted from the [Contributor Covenant][homepage], version 2.1, available at [https://www.contributor-covenant.org/version/2/1/code_of_conduct.html][v2.1].

[homepage]: https://www.contributor-covenant.org
[v2.1]: https://www.contributor-covenant.org/version/2/1/code_of_conduct.html
```

- [ ] **Step 2: Commit**

```bash
git add CODE_OF_CONDUCT.md
git commit -m "docs: add Contributor Covenant 2.1 code of conduct"
```

---

### Task 4: CONTRIBUTING.md

**Files:**
- Create: `CONTRIBUTING.md`

- [ ] **Step 1: Create CONTRIBUTING.md**

```markdown
# Contributing to lattice-protocol

## What lives here

This repo contains shared protocol definitions for the Lattice mesh network:

- **`opcodes/opcodes.go`** — serial command opcode constants (Go)
- **`adapter/types.go`** — adapter type identifiers and helpers (Go)
- **`c/`** — C headers generated from the Go constants; never edit these by hand

Changes here affect all consumers: `motionSensorServer` (imports as Go module) and `Lattice-nodes` (includes as git submodule). Treat every change as a protocol change.

## Prerequisites

- Go 1.21 or later (`go version`)
- Git

## Adding an opcode

1. Edit `opcodes/opcodes.go`. Add your constant in the appropriate group, following the existing naming convention (`OpXxx` in Go, which becomes `OP_XXX` in the generated C header).

   ```go
   const (
       OpYourNewOpcode Opcode = 0xXX
   )
   ```

2. Regenerate C headers:

   ```sh
   make generate
   # equivalent: go generate ./...
   ```

3. Verify the headers are in sync with the Go constants:

   ```sh
   make check
   # Runs: go generate ./... && git diff --exit-code c/
   # Must exit 0 before you commit.
   ```

4. Run tests:

   ```sh
   go test ./...
   go vet ./...
   ```

5. Commit Go source and generated `c/` files together in one commit:

   ```sh
   git add opcodes/opcodes.go c/opcodes.h
   git commit -m "feat(opcodes): add OpYourNewOpcode (0xXX)"
   ```

## Adding an adapter type

Same flow as adding an opcode, but edit `adapter/types.go` instead and include `c/adapter_types.h` in the commit.

## Semver rules

| Change | Bump |
|--------|------|
| Add new opcode or adapter type | Patch (`v0.Y.Z+1`) |
| Rename or remove an existing constant | Minor (`v0.Y+1.0`) |
| Breaking protocol redesign | Open an issue first |

After merging, create and push a semver tag:

```sh
git tag v0.Y.Z
git push origin v0.Y.Z
```

Update the versioning table in `README.md` with the new tag and a one-line description of what changed.

## Code style

- Run `gofmt` before committing (enforced by `go vet ./...`)
- Constant names: `OpXxx` for opcodes, `AdapterTypeXxx` for adapter types
- Group related constants together, separated by a blank line from unrelated groups

## Pull request process

1. Fork the repo and create a branch: `feature/your-opcode-name` or `fix/your-fix`
2. Follow the steps above for adding an opcode or adapter type
3. Ensure `make check`, `go test ./...`, and `go vet ./...` all pass locally
4. Open a PR against `main` and fill in the PR template
5. All CI jobs (`go-test`, `header-sync`, CodeQL) must be green; 1 approving review required
```

- [ ] **Step 2: Commit**

```bash
git add CONTRIBUTING.md
git commit -m "docs: add contributing guide with full protocol workflow"
```

---

### Task 5: GitHub Templates

**Files:**
- Create: `.github/pull_request_template.md`
- Create: `.github/ISSUE_TEMPLATE/bug_report.md`
- Create: `.github/ISSUE_TEMPLATE/feature_request.md`

- [ ] **Step 1: Create directory structure**

```bash
mkdir -p .github/ISSUE_TEMPLATE
```

- [ ] **Step 2: Create PR template**

Create `.github/pull_request_template.md`:

```markdown
## Summary

<!-- What does this change and why? -->

## Type of change

- [ ] Opcode addition
- [ ] Adapter type addition
- [ ] Bug fix
- [ ] Documentation
- [ ] CI

## Checklist

- [ ] C headers regenerated (`make generate`)
- [ ] `make check` passes (headers in sync with Go constants)
- [ ] `go test ./...` passes
- [ ] `go vet ./...` passes
- [ ] Versioning table in `README.md` updated (if constants changed)
- [ ] Go source and `c/` files committed together in one commit

## Testing done

<!-- How did you verify this change? -->
```

- [ ] **Step 3: Create bug report template**

Create `.github/ISSUE_TEMPLATE/bug_report.md`:

```markdown
---
name: Bug report
about: Report incorrect protocol behaviour or a defect in the header generator
title: ''
labels: bug
assignees: ''
---

## Describe the bug

<!-- What went wrong? -->

## Expected behaviour

<!-- What should have happened? -->

## Steps to reproduce

1.
2.
3.

## Environment

- Go version:
- Consumer repo: motionSensorServer / Lattice-nodes / other
- Protocol version (tag):
```

- [ ] **Step 4: Create feature request template**

Create `.github/ISSUE_TEMPLATE/feature_request.md`:

```markdown
---
name: Feature / opcode request
about: Propose a new opcode, adapter type, or protocol change
title: ''
labels: enhancement
assignees: ''
---

## What do you need?

<!-- Describe the new opcode, adapter type, or behaviour you need -->

## Why is it needed?

<!-- What problem does this solve? Which consumer repo needs it? -->

## Proposed constant value(s)

<!-- Suggest an opcode byte or adapter type ID, or leave blank if unsure -->

## Backwards compatibility

<!-- Patch bump (new addition) or minor bump (rename/removal)? See CONTRIBUTING.md for semver rules -->
```

- [ ] **Step 5: Commit**

```bash
git add .github/
git commit -m "chore: add GitHub PR and issue templates"
```

- [ ] **Step 6: Open PR for Track B**

Open a PR titled `chore: OSS hardening — Track B (community health)` targeting `main`. Three commits: CODE_OF_CONDUCT.md, CONTRIBUTING.md, GitHub templates.

---

## Track C — CI + Docs

### Task 6: CI Workflow

**Files:**
- Create: `.github/workflows/ci.yml`

- [ ] **Step 1: Create ci.yml**

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
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.23'
      - run: go vet ./...
      - run: go test ./...

  header-sync:
    name: Header sync
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.23'
      - name: Check C headers are in sync with Go constants
        run: make check
```

- [ ] **Step 2: Verify YAML is valid**

```bash
python3 -c "import yaml; yaml.safe_load(open('.github/workflows/ci.yml'))" && echo "YAML valid"
```

Expected: `YAML valid`

- [ ] **Step 3: Commit**

```bash
git add .github/workflows/ci.yml
git commit -m "ci: add go-test and header-sync jobs"
```

---

### Task 7: CodeQL Workflow

**Files:**
- Create: `.github/workflows/codeql.yml`

- [ ] **Step 1: Create codeql.yml**

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
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.23'
      - uses: github/codeql-action/init@v3
        with:
          languages: go
          queries: security-and-quality
      - uses: github/codeql-action/autobuild@v3
      - uses: github/codeql-action/analyze@v3
        with:
          category: '/language:go'
```

- [ ] **Step 2: Verify YAML is valid**

```bash
python3 -c "import yaml; yaml.safe_load(open('.github/workflows/codeql.yml'))" && echo "YAML valid"
```

Expected: `YAML valid`

- [ ] **Step 3: Commit**

```bash
git add .github/workflows/codeql.yml
git commit -m "ci: add CodeQL static analysis for Go"
```

---

### Task 8: README Rewrite + Branch Protection

**Files:**
- Modify: `README.md`

- [ ] **Step 1: Rewrite README.md**

Replace the entire file contents:

```markdown
# lattice-protocol

[![CI](https://github.com/superbrobenji/lattice-protocol/actions/workflows/ci.yml/badge.svg)](https://github.com/superbrobenji/lattice-protocol/actions/workflows/ci.yml)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Shared protocol definitions for the Lattice mesh network. Defines the opcode and adapter-type constants consumed by all Lattice services and firmware.

Used by:
- **motionSensorServer** (Go) — imports as a Go module
- **Lattice-nodes** (ESP32/C++) — includes via git submodule

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
    "github.com/superbrobenji/lattice-protocol/opcodes"
    "github.com/superbrobenji/lattice-protocol/adapter"
)

payload[0] = opcodes.OpLEDSolid
if adapter.IsOutput(node.AdapterType) { ... }
```

### C (Lattice-nodes — via git submodule at `main/lib/lattice-protocol`)

```c
#include "lib/lattice-protocol/c/opcodes.h"
#include "lib/lattice-protocol/c/adapter_types.h"

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

Copyright (C) 2026 Lattice Contributors

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

See [LICENSE](LICENSE) for the full terms.
```

- [ ] **Step 2: Verify TODO comment is removed**

```bash
grep -c "TODO: open-source hardening" README.md
```

Expected: `0`

- [ ] **Step 3: Commit**

```bash
git add README.md
git commit -m "docs: rewrite README — add badges, license section, remove hardening TODO"
```

- [ ] **Step 4: Open PR for Track C**

Open a PR titled `chore: OSS hardening — Track C (CI + docs)` targeting `main`. Three commits: ci.yml, codeql.yml, README rewrite.

- [ ] **Step 5: Configure branch protection (manual step in GitHub settings)**

Navigate to: `https://github.com/superbrobenji/lattice-protocol/settings/branches` → Add rule for `main`:

- Require status checks to pass before merging
  - Add required checks: `Test`, `Header sync`
- Require approvals: 1
- Dismiss stale pull request approvals when new commits are pushed
- Do not allow bypassing the above settings
