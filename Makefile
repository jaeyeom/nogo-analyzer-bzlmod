.PHONY: all check build format check-format gazelle update-repos clean

# Local dev: format then build
all: format build

# CI-friendly: check without mutation
check: check-format build

# Build
build:
	bazel build //...

# Format
format:
	gofumpt -w .

check-format:
	@test -z "$$(gofumpt -l .)" || { gofumpt -l .; echo "Run 'make format' to fix."; exit 1; }

# Gazelle
gazelle:
	bazel run //:gazelle

update-repos:
	bazel run //:gazelle -- update-repos -from_file=go.mod

# Clean
clean:
	bazel clean
