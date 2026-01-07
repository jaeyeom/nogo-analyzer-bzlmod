# nogo-analyzer-bzlmod

Staticcheck analyzers for Bazel's nogo, packaged as a bzlmod module.

This module provides individual [staticcheck](https://staticcheck.dev/) analyzers that can be used with [nogo](https://github.com/bazelbuild/rules_go/blob/master/go/nogo.rst) in Bazel projects using bzlmod.

## Installation

Add the module to your `MODULE.bazel`:

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

## License

MIT
