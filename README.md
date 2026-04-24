# nogo-analyzer-bzlmod

Staticcheck analyzers for Bazel's nogo, packaged as a bzlmod module.

This module provides individual [staticcheck](https://staticcheck.dev/) analyzers that can be used with [nogo](https://github.com/bazelbuild/rules_go/blob/master/go/nogo.rst) in Bazel projects using bzlmod.

## Installation

Add the module to your `MODULE.bazel`:

```starlark
bazel_dep(name = "nogo_analyzer_bzlmod", version = "0.1.0")
```

### Using Unreleased Versions

To use an unreleased version from the repository, you can use `git_override`:

```starlark
bazel_dep(name = "nogo_analyzer_bzlmod", version = "0.1.0")
git_override(
    module_name = "nogo_analyzer_bzlmod",
    remote = "https://github.com/jaeyeom/nogo-analyzer-bzlmod.git",
    commit = "<commit-sha>",
)
```

## Usage

In your `BUILD.bazel` file where you define your nogo target:

```starlark
load("@rules_go//go:def.bzl", "nogo")
load("@nogo_analyzer_bzlmod//:def.bzl", "ANALYZERS", "staticcheck_analyzers")

nogo(
    name = "nogo",
    deps = staticcheck_analyzers(ANALYZERS),
    visibility = ["//visibility:public"],
)
```

### Excluding Specific Analyzers

To exclude specific analyzers, use the `-` prefix:

```starlark
nogo(
    name = "nogo",
    deps = staticcheck_analyzers(ANALYZERS + ["-U1000", "-ST1000"]),
    visibility = ["//visibility:public"],
)
```

### Using a Subset of Analyzers

You can also specify only the analyzers you want:

```starlark
nogo(
    name = "nogo",
    deps = staticcheck_analyzers(["SA1000", "SA1001", "SA1002"]),
    visibility = ["//visibility:public"],
)
```

## Available Analyzers

The module includes all staticcheck analyzer categories:

| Category | Description                        |
|----------|------------------------------------|
| **QF**   | Quickfix suggestions               |
| **S**    | Simple code simplifications        |
| **SA**   | Staticcheck (bugs and correctness) |
| **ST**   | Stylecheck (code style)            |
| **U**    | Unused code detection              |

See [staticcheck.dev/docs/checks](https://staticcheck.dev/docs/checks/) for detailed documentation of each analyzer.

## Requirements

- Bazel 7.0+ with bzlmod enabled
- rules_go 0.53.0+

## Verifying releases

Releases are signed with [GitHub artifact attestations](https://docs.github.com/en/actions/security-guides/using-artifact-attestations-to-establish-provenance-for-builds),
which produce [SLSA](https://slsa.dev/) build provenance for the source
archive. Each release asset `nogo-analyzer-bzlmod-vX.Y.Z.tar.gz` is
published alongside `nogo-analyzer-bzlmod-vX.Y.Z.tar.gz.intoto.jsonl`
(the Sigstore bundle, following the naming convention expected by the
[Bazel Central Registry](https://github.com/bazelbuild/bazel-central-registry/discussions/2721)).

### Online verification

From a machine with internet access, the GitHub CLI fetches the
attestation and the trusted roots automatically:

```sh
gh attestation verify nogo-analyzer-bzlmod-vX.Y.Z.tar.gz \
  --repo jaeyeom/nogo-analyzer-bzlmod
```

### Bundle-based verification

If you already have the `.intoto.jsonl` bundle downloaded alongside the
archive, you can skip the attestation lookup and point `gh` at the
bundle directly (still requires network access to fetch Sigstore
trusted roots):

```sh
gh attestation verify nogo-analyzer-bzlmod-vX.Y.Z.tar.gz \
  --repo jaeyeom/nogo-analyzer-bzlmod \
  --bundle nogo-analyzer-bzlmod-vX.Y.Z.tar.gz.intoto.jsonl
```

### Offline verification

Fully offline verification additionally requires a pre-exported
`trusted_root.jsonl`. Generate it on an online machine, then import
everything into your air-gapped environment. See
[Verifying artifact attestations offline](https://docs.github.com/en/actions/how-tos/secure-your-work/use-artifact-attestations/verify-attestations-offline)
for the full procedure.

On an online machine:

```sh
gh attestation trusted-root > trusted_root.jsonl
```

Transfer `trusted_root.jsonl`, the archive, and the `.intoto.jsonl`
bundle into the offline environment (along with the `gh` CLI), then
run:

```sh
gh attestation verify nogo-analyzer-bzlmod-vX.Y.Z.tar.gz \
  --repo jaeyeom/nogo-analyzer-bzlmod \
  --bundle nogo-analyzer-bzlmod-vX.Y.Z.tar.gz.intoto.jsonl \
  --custom-trusted-root trusted_root.jsonl
```

## License

MIT
