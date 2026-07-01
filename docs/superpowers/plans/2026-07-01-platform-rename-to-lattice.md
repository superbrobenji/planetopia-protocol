# Platform Rename: Planetopia → Lattice Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Rename every occurrence of `planetopia`/`Planetopia`/`PLANETOPIA` across all 3 repos to `lattice`/`Lattice`/`LATTICE`, including repo names, Go module paths, C++ namespaces, C header guards, submodule paths, and strings.

**Architecture:** Dependency-ordered — `lattice-protocol` renamed and merged first (both other repos depend on it), then `lattice-hub` and `lattice-nodes` in parallel. GitHub repo renames happen before any code changes to activate redirects.

**Tech Stack:** Go 1.21+, C/C++ (Arduino/CMake/GoogleTest), bash `sed`/`find` for bulk renames, `gh` CLI for PRs.

## Global Constraints

- GitHub username: `superbrobenji` — all module paths and submodule URLs use this prefix
- `Lattice` — capitalized in prose/comments; `lattice` in code identifiers; `LATTICE` in C preprocessor macros
- Zero occurrences of `planetopia`, `Planetopia`, or `PLANETOPIA` anywhere after completion
- macOS `sed`: always `sed -i ''` (BSD sed requires empty string after `-i`)
- PR branch: `feat/rename-planetopia-to-lattice` in each repo
- No force pushes; no `--no-verify`
- Internal Go package names (`mesh`, `adapter`, `eventStore`, etc.) are out of scope — not branded
- React component names, UI copy, Kafka topic names, `.env` files are out of scope

---

### Task 1: Rename repos on GitHub

**Files:** None (GitHub UI)

**Prerequisite:** None

- [ ] **Step 1: Rename motionSensorServer → lattice-hub**

  GitHub → https://github.com/superbrobenji/motionSensorServer → Settings → General → Repository name → `lattice-hub` → Rename repository

- [ ] **Step 2: Rename planetopia-nodes → lattice-nodes**

  GitHub → https://github.com/superbrobenji/planetopia-nodes → Settings → General → Repository name → `lattice-nodes` → Rename repository

- [ ] **Step 3: Rename planetopia-protocol → lattice-protocol**

  GitHub → https://github.com/superbrobenji/planetopia-protocol → Settings → General → Repository name → `lattice-protocol` → Rename repository

- [ ] **Step 4: Verify GitHub redirects are live**

```bash
curl -sI https://github.com/superbrobenji/planetopia-protocol | grep location
# Expected: location: https://github.com/superbrobenji/lattice-protocol
curl -sI https://github.com/superbrobenji/planetopia-nodes | grep location
# Expected: location: https://github.com/superbrobenji/lattice-nodes
curl -sI https://github.com/superbrobenji/motionSensorServer | grep location
# Expected: location: https://github.com/superbrobenji/lattice-hub
```

---

### Task 2: lattice-protocol — Update Go module path and all Go source

**Files:**
- Modify: `go.mod`
- Modify: `opcodes/opcodes.go`
- Modify: `adapter/types.go`
- Modify: `cmd/gen-headers/main.go`
- Modify: `go.sum` (via `go mod tidy`)

**Prerequisite:** Task 1

- [ ] **Step 1: Verify tests pass before any changes**

```bash
cd /Users/benjamin.swanepoel/projects/personal/planetopia-protocol
go test ./...
# Expected: ok  github.com/superbrobenji/planetopia-protocol/...
```

- [ ] **Step 2: Create feature branch**

```bash
git checkout -b feat/rename-planetopia-to-lattice
```

- [ ] **Step 3: Update go.mod module declaration**

```bash
sed -i '' 's|module github.com/superbrobenji/planetopia-protocol|module github.com/superbrobenji/lattice-protocol|g' go.mod
```

- [ ] **Step 4: Update all planetopia occurrences in Go source files**

```bash
find . -name "*.go" -not -path "./.git/*" -exec sed -i '' \
  -e 's|github.com/superbrobenji/planetopia-protocol|github.com/superbrobenji/lattice-protocol|g' \
  -e 's/Planetopia/Lattice/g' \
  -e 's/planetopia/lattice/g' \
  -e 's/PLANETOPIA/LATTICE/g' {} +
```

- [ ] **Step 5: Verify no planetopia refs remain in Go files**

```bash
grep -rn "planetopia\|Planetopia\|PLANETOPIA" --include="*.go" .
# Expected: no output
```

- [ ] **Step 6: Run go mod tidy**

```bash
go mod tidy
```

- [ ] **Step 7: Verify tests still pass**

```bash
go test ./...
# Expected: ok  github.com/superbrobenji/lattice-protocol/...
```

- [ ] **Step 8: Commit**

```bash
git add go.mod go.sum
git add opcodes/opcodes.go adapter/types.go cmd/gen-headers/main.go
git commit -m "feat: rename Go module path and strings planetopia → lattice"
```

---

### Task 3: lattice-protocol — Regenerate C headers with LATTICE_ guards

The `cmd/gen-headers` tool generates `c/opcodes.h` and `c/adapter_types.h`. Task 2 updated the generator source; now regenerate the output.

**Files:**
- Modify: `c/opcodes.h` (generated)
- Modify: `c/adapter_types.h` (generated)

**Prerequisite:** Task 2

- [ ] **Step 1: Verify gen-headers outputs LATTICE_ guards (not PLANETOPIA_)**

```bash
grep -n "PLANETOPIA\|LATTICE" cmd/gen-headers/main.go
# Expected: only LATTICE_ lines remain; no PLANETOPIA_
```

  If any `PLANETOPIA_` remain (e.g. hardcoded template strings not caught by Task 2 sed):

```bash
sed -i '' 's/PLANETOPIA_/LATTICE_/g' cmd/gen-headers/main.go
```

- [ ] **Step 2: Regenerate C headers**

```bash
go generate ./...
```

- [ ] **Step 3: Verify header guards updated**

```bash
grep -E "ifndef|define|endif" c/opcodes.h
# Expected: lines containing LATTICE_OPCODES_H
grep -E "ifndef|define|endif" c/adapter_types.h
# Expected: lines containing LATTICE_ADAPTER_TYPES_H
```

- [ ] **Step 4: Verify no planetopia refs in generated headers**

```bash
grep -i "planetopia" c/opcodes.h c/adapter_types.h
# Expected: no output
```

- [ ] **Step 5: Commit**

```bash
git add c/opcodes.h c/adapter_types.h
git add cmd/gen-headers/main.go
git commit -m "feat: regenerate C headers with LATTICE_ include guards"
```

---

### Task 4: lattice-protocol — Update CI, docs, final scan, push PR

**Files:**
- Modify: `.github/workflows/*.yml`
- Modify: `Makefile`
- Modify: `*.md` files (if any)

**Prerequisite:** Task 3

- [ ] **Step 1: Update workflow files**

```bash
find .github -name "*.yml" -exec sed -i '' \
  -e 's/Planetopia/Lattice/g' \
  -e 's/planetopia/lattice/g' \
  -e 's/PLANETOPIA/LATTICE/g' {} +
```

- [ ] **Step 2: Update Makefile**

```bash
sed -i '' \
  -e 's/Planetopia/Lattice/g' \
  -e 's/planetopia/lattice/g' \
  -e 's/PLANETOPIA/LATTICE/g' Makefile
```

- [ ] **Step 3: Update any markdown files**

```bash
find . -name "*.md" -not -path "./.git/*" -exec sed -i '' \
  -e 's/Planetopia/Lattice/g' \
  -e 's/planetopia/lattice/g' \
  -e 's/PLANETOPIA/LATTICE/g' {} +
```

- [ ] **Step 4: Full repo scan — verify zero planetopia occurrences**

```bash
grep -rin "planetopia" --exclude-dir=.git .
# Expected: no output
```

  If any hits appear, fix them before continuing.

- [ ] **Step 5: Run all tests and checks**

```bash
go test ./...
make check
# Expected: all pass
```

- [ ] **Step 6: Commit remaining changes**

```bash
git add .github/ Makefile
git add $(git ls-files --modified "*.md")
git diff --cached --quiet || git commit -m "feat: update CI, Makefile, docs planetopia → lattice"
```

- [ ] **Step 7: Push and open PR**

```bash
git push -u origin feat/rename-planetopia-to-lattice
gh pr create \
  --title "feat: rename platform planetopia → lattice" \
  --body "$(cat <<'EOF'
Renames all internal references from Planetopia to Lattice.

- go.mod module path: github.com/superbrobenji/planetopia-protocol → lattice-protocol
- C header guards: PLANETOPIA_* → LATTICE_*
- Regenerated c/opcodes.h and c/adapter_types.h via go generate
- Updated CI workflows, Makefile, comments, copyright strings
EOF
)"
```

- [ ] **Step 8: Merge PR before starting Tasks 5 and 6**

  Wait for CI to pass, then merge. Tasks 5 and 6 depend on this module path being live.

---

### Task 5: lattice-hub — Update all Go modules, imports, strings, CI

**Files:**
- Modify: all `go.mod` files (find with `find . -name "go.mod"`)
- Modify: all `*.go` files
- Modify: all `go.sum` files (via `go mod tidy`)
- Modify: `docker-compose*.yml` / `Dockerfile*`
- Modify: `.github/workflows/*.yml`
- Modify: `*.md` files

**Prerequisite:** Task 4 merged

**Interfaces:**
- Consumes: `github.com/superbrobenji/lattice-protocol` (live after Task 4 merged)
- Produces: Go module at `github.com/superbrobenji/lattice-hub`; all imports updated

- [ ] **Step 1: Verify tests pass before changes**

```bash
cd /Users/benjamin.swanepoel/projects/personal/motionSensorServer
# Run from each Go module root found below
find . -name "go.mod" -not -path "./.git/*"
# For each go.mod directory, run: go test ./...
```

- [ ] **Step 2: Create feature branch**

```bash
git checkout -b feat/rename-planetopia-to-lattice
```

- [ ] **Step 3: Update all go.mod module declarations (motionServer → lattice-hub)**

```bash
find . -name "go.mod" -not -path "./.git/*" -exec sed -i '' \
  's|module github.com/superbrobenji/motionServer|module github.com/superbrobenji/lattice-hub|g' {} +
```

- [ ] **Step 4: Update protocol dependency in all go.mod files**

```bash
find . -name "go.mod" -not -path "./.git/*" -exec sed -i '' \
  's|github.com/superbrobenji/planetopia-protocol|github.com/superbrobenji/lattice-protocol|g' {} +
```

- [ ] **Step 5: Update all Go source imports and strings**

```bash
find . -name "*.go" -not -path "./.git/*" -exec sed -i '' \
  -e 's|github.com/superbrobenji/motionServer|github.com/superbrobenji/lattice-hub|g' \
  -e 's|github.com/superbrobenji/planetopia-protocol|github.com/superbrobenji/lattice-protocol|g' \
  -e 's/Planetopia/Lattice/g' \
  -e 's/planetopia/lattice/g' \
  -e 's/PLANETOPIA/LATTICE/g' {} +
```

- [ ] **Step 6: Run go mod tidy in each Go module**

```bash
find . -name "go.mod" -not -path "./.git/*" | while read f; do
  dir=$(dirname "$f")
  echo "Tidying $dir"
  (cd "$dir" && go mod tidy)
done
```

Expected: downloads `github.com/superbrobenji/lattice-protocol`, updates go.sum.

- [ ] **Step 7: Verify no old paths or planetopia refs remain in Go files**

```bash
grep -rn "superbrobenji/motionServer\|superbrobenji/planetopia-protocol\|[Pp]lanetopia\|PLANETOPIA" \
  --include="*.go" --exclude-dir=.git .
# Expected: no output
```

- [ ] **Step 8: Run tests**

```bash
find . -name "go.mod" -not -path "./.git/*" | while read f; do
  dir=$(dirname "$f")
  echo "Testing $dir"
  (cd "$dir" && go test ./...)
done
# Expected: all pass
```

- [ ] **Step 9: Update Docker/compose files**

```bash
find . \( -name "docker-compose*.yml" -o -name "Dockerfile*" \) -not -path "./.git/*" \
  -exec sed -i '' \
  -e 's/Planetopia/Lattice/g' \
  -e 's/planetopia/lattice/g' \
  -e 's/PLANETOPIA/LATTICE/g' {} +
```

- [ ] **Step 10: Update CI workflows and markdown**

```bash
find .github -name "*.yml" -exec sed -i '' \
  -e 's/Planetopia/Lattice/g' \
  -e 's/planetopia/lattice/g' \
  -e 's/PLANETOPIA/LATTICE/g' {} +

find . -name "*.md" -not -path "./.git/*" -exec sed -i '' \
  -e 's/Planetopia/Lattice/g' \
  -e 's/planetopia/lattice/g' \
  -e 's/PLANETOPIA/LATTICE/g' {} +
```

- [ ] **Step 11: Full repo scan — verify zero planetopia occurrences**

```bash
grep -rin "planetopia" --exclude-dir=.git --exclude-dir=node_modules .
# Expected: no output
```

- [ ] **Step 12: Commit and push PR**

```bash
git add -A
git commit -m "feat: rename module paths and all planetopia refs → lattice"
git push -u origin feat/rename-planetopia-to-lattice
gh pr create \
  --title "feat: rename platform planetopia → lattice" \
  --body "$(cat <<'EOF'
Renames all Planetopia references to Lattice in the hub service.

- go.mod module path: motionServer → lattice-hub
- Protocol dependency: planetopia-protocol → lattice-protocol
- All internal Go import paths updated
- go mod tidy run across all modules
- Docker/compose service names updated
- CI workflows and docs updated
EOF
)"
```

---

### Task 6: lattice-nodes — Update submodule path

Run this in parallel with Task 5 (both can start after Task 4 is merged).

**Files:**
- Modify: `.gitmodules`
- Rename: `main/lib/planetopia-protocol/` → `main/lib/lattice-protocol/`

**Prerequisite:** Task 4 merged

- [ ] **Step 1: Create feature branch**

```bash
cd /Users/benjamin.swanepoel/projects/personal/planetopia-nodes
git checkout -b feat/rename-planetopia-to-lattice
```

- [ ] **Step 2: Update .gitmodules path and URL**

```bash
sed -i '' \
  -e 's|path = main/lib/planetopia-protocol|path = main/lib/lattice-protocol|g' \
  -e 's|url = .*planetopia-protocol.*|url = https://github.com/superbrobenji/lattice-protocol|g' \
  .gitmodules
```

- [ ] **Step 3: Rename the submodule directory**

```bash
git mv main/lib/planetopia-protocol main/lib/lattice-protocol
```

- [ ] **Step 4: Sync git submodule config**

```bash
git submodule sync
git submodule update --init main/lib/lattice-protocol
```

- [ ] **Step 5: Verify submodule points to new location**

```bash
cat .gitmodules
# Expected: path = main/lib/lattice-protocol
#           url = https://github.com/superbrobenji/lattice-protocol
git submodule status
# Expected: hash main/lib/lattice-protocol (...)
```

- [ ] **Step 6: Commit**

```bash
git add .gitmodules main/lib/lattice-protocol
git commit -m "feat: rename submodule planetopia-protocol → lattice-protocol"
```

---

### Task 7: lattice-nodes — Rename C++ namespaces and header guards

**Files:**
- Modify: `main/main.ino`
- Modify: all `main/src/**/*.h` and `main/src/**/*.cpp`
- Modify: all `tests/**/*.h` and `tests/**/*.cpp`
- Modify: `main/project_config.h`

**Prerequisite:** Task 6

- [ ] **Step 1: Rename all C++ namespace declarations**

```bash
find main tests \( -name "*.h" -o -name "*.cpp" -o -name "*.ino" \) \
  -exec sed -i '' 's/namespace planetopia/namespace lattice/g' {} +
```

- [ ] **Step 2: Rename all namespace-qualified references**

```bash
find main tests \( -name "*.h" -o -name "*.cpp" -o -name "*.ino" \) \
  -exec sed -i '' 's/planetopia::/lattice::/g' {} +
```

- [ ] **Step 3: Update using namespace directives**

```bash
find main tests \( -name "*.h" -o -name "*.cpp" -o -name "*.ino" \) \
  -exec sed -i '' 's/using namespace planetopia/using namespace lattice/g' {} +
```

- [ ] **Step 4: Update all C/C++ include guard macros**

```bash
find main tests \( -name "*.h" -o -name "*.cpp" -o -name "*.ino" \) \
  -exec sed -i '' \
  -e 's/PLANETOPIA_/LATTICE_/g' \
  -e 's/_PLANETOPIA/_LATTICE/g' {} +
```

- [ ] **Step 5: Update include paths referencing the submodule**

```bash
find main tests \( -name "*.h" -o -name "*.cpp" -o -name "*.ino" \) \
  -exec sed -i '' 's|planetopia-protocol/|lattice-protocol/|g' {} +
```

- [ ] **Step 6: Update all remaining string/comment occurrences**

```bash
find main tests \( -name "*.h" -o -name "*.cpp" -o -name "*.ino" \) \
  -exec sed -i '' \
  -e 's/Planetopia/Lattice/g' \
  -e 's/planetopia/lattice/g' {} +
```

- [ ] **Step 7: Verify no planetopia refs remain in C++ source**

```bash
grep -rn "planetopia\|Planetopia\|PLANETOPIA" --include="*.h" --include="*.cpp" --include="*.ino" main/ tests/
# Expected: no output
```

- [ ] **Step 8: Commit**

```bash
git add main/ tests/
git commit -m "feat: rename C++ namespace planetopia:: → lattice:: and update header guards"
```

---

### Task 8: lattice-nodes — Update build config, run tests, final scan, push PR

**Files:**
- Modify: `tests/CMakeLists.txt`
- Modify: any `*.cmake` files
- Modify: `.github/workflows/*.yml`
- Modify: `*.md` files

**Prerequisite:** Task 7

- [ ] **Step 1: Update CMakeLists and build config**

```bash
find . \( -name "CMakeLists.txt" -o -name "*.cmake" \) -not -path "./.git/*" -exec sed -i '' \
  -e 's/Planetopia/Lattice/g' \
  -e 's/planetopia/lattice/g' \
  -e 's/PLANETOPIA/LATTICE/g' {} +
```

- [ ] **Step 2: Update CI workflows and markdown**

```bash
find .github -name "*.yml" -exec sed -i '' \
  -e 's/Planetopia/Lattice/g' \
  -e 's/planetopia/lattice/g' \
  -e 's/PLANETOPIA/LATTICE/g' {} +

find . -name "*.md" -not -path "./.git/*" -exec sed -i '' \
  -e 's/Planetopia/Lattice/g' \
  -e 's/planetopia/lattice/g' \
  -e 's/PLANETOPIA/LATTICE/g' {} +
```

- [ ] **Step 3: Build and run tests**

```bash
cmake -B tests/build tests/
cmake --build tests/build --parallel
cd tests/build && ctest --output-on-failure
# Expected: all tests pass
```

- [ ] **Step 4: Full repo scan — verify zero planetopia occurrences**

```bash
grep -rin "planetopia" --exclude-dir=.git --exclude-dir=tests/build .
# Expected: no output
```

  If any hits appear, fix them before continuing.

- [ ] **Step 5: Commit and push PR**

```bash
git add .github/ tests/CMakeLists.txt
git add $(git ls-files --modified "*.md" "*.cmake")
git diff --cached --quiet || git commit -m "feat: update build config and CI planetopia → lattice"
git push -u origin feat/rename-planetopia-to-lattice
gh pr create \
  --title "feat: rename platform planetopia → lattice" \
  --body "$(cat <<'EOF'
Renames all Planetopia references to Lattice in the nodes firmware.

- C++ namespaces: planetopia:: → lattice::
- Submodule: lib/planetopia-protocol → lib/lattice-protocol (URL updated)
- C header guards: PLANETOPIA_* → LATTICE_*
- #include paths updated for renamed submodule
- CMakeLists, CI workflows, docs updated
- All tests pass
EOF
)"
```
