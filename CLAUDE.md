# CLAUDE.md

## Project Overview

Bzlmod module providing individual staticcheck analyzers for Bazel's nogo. One
Go source file (`staticcheck/analyzer.go`) serves 167 analyzer targets via
Bazel `x_defs` injection at build time.

## Build and Validate

```bash
# Build all targets (requires Bazel 7.0+ with bzlmod)
bazel build //...

# Regenerate BUILD files after changing Go imports
bazel run //:gazelle

# Update go.mod / go.sum after dependency changes
bazel run //:gazelle -- update-repos -from_file=go.mod
```

## Code Formatting

All Go code must be formatted with **gofumpt** (enforced in CI):

```bash
gofumpt -w .
```

## Project Structure

- `def.bzl` — Public API: `ANALYZERS` list and `staticcheck_analyzers()` macro
- `staticcheck/analyzer.go` — Single source compiled per-analyzer via `x_defs`
- `staticcheck/util/` — Analyzer registry (`util.go`) and ignore directives (`directive.go`)
- `staticcheck/BUILD.bazel` — Generates 167 `go_library` targets in a loop
- `MODULE.bazel` — Bzlmod module definition (Go 1.24.2, rules_go 0.53.0)

## Key Constraints

- The `ANALYZERS` list in `def.bzl` must stay in sync with the analyzer
  packages registered in `staticcheck/util/util.go`.
- Analyzer targets share a single source file; do not add per-analyzer Go files.
- CI runs `gofumpt -l .` and `bazel build //...` — both must pass.
- Releases are created by pushing a `v*` tag; `.gitattributes` controls archive
  exclusions.
