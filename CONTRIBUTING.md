# Contributing to planetopia-protocol

## What lives here

This repo contains shared protocol definitions for the Planetopia mesh network:

- **`opcodes/opcodes.go`** — serial command opcode constants (Go)
- **`adapter/types.go`** — adapter type identifiers and helpers (Go)
- **`c/`** — C headers generated from the Go constants; never edit these by hand

Changes here affect all consumers: `motionSensorServer` (imports as Go module) and `Planetopia-nodes` (includes as git submodule). Treat every change as a protocol change.

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
