# AGENTS.md

This file provides guidance to AI coding agents when working with code in this repository.

## Overview

This repository contains:
1. **Protocol Buffer definitions** (`proto/`) - service and message definitions for Cofide's APIs
2. **Generated code** (`gen/`) - auto-generated Go and TypeScript stubs from the proto files
3. **Go client library** (`pkg/connect/client/`) - hand-written wrapper clients for the Cofide Connect API

APIs are versioned as `v1alpha1` or `v1beta1` (backward compatibility not guaranteed).

## Commands

Commands use [Just](https://github.com/casey/just) as the task runner.

| Command | Purpose |
|---|---|
| `just fmt` | Format proto files (`buf format -w`) |
| `just lint` | Run all of the linters below |
| `just buf-lint` | Lint proto files with buf |
| `just api-lint` | Lint the new APIs against Google's AIPs with api-linter (see [API linting](#api-linting)) |
| `just go-lint` | Run golangci-lint only |
| `just proto-gen` | Regenerate Go and TypeScript stubs from proto files |
| `just proto-docs` | Generate Markdown API docs (to `docs/`, git-ignored) |
| `just test` | Run Go tests |
| `just test-race` | Run Go tests with race detector |

To run a single test package: `go run gotest.tools/gotestsum@latest -- ./pkg/connect/client/...`

To run with `-run`: `go run gotest.tools/gotestsum@latest -- -run TestName ./...`

## Architecture

### Proto → Generated Code → Client Wrapper

The project follows a strict layered pattern:

1. **`proto/`** - Source of truth, split into two buf modules (see `buf.yaml`).
   - The old APIs: 13 gRPC services defined under `proto/connect/`, plus domain message types across 20+ packages. Their module is rooted at the repo, so packages are named `proto.<name>.<version>` and imports are written `proto/<name>/<version>/<file>.proto`.
   - The new APIs, with each resource's message and service in one package under `proto/cofide/<name>/<version>/`. Their module is rooted at `proto/` and includes only `proto/cofide/`, so packages are named `cofide.<name>.<version>` and imports are written `cofide/<name>/<version>/<file>.proto`, without the `proto/` prefix. An import written with the prefix fails with `file does not exist`.

2. **`gen/go/`** - Auto-generated, do not edit manually. Regenerate with `just proto-gen`. Mirrors the file paths within each module, so code for the old APIs is under `gen/go/proto/` and for the new APIs under `gen/go/cofide/`.

3. **`pkg/connect/client/`** - Hand-written Go clients. Each service gets its own package (e.g., `pkg/connect/client/cluster/v1alpha1/`). The `ClientSet` interface in `clientset.go` aggregates all service clients into a single accessor. `New(conn grpc.ClientConnInterface)` constructs a `ClientSet`.

### Fake clients

`pkg/connect/client/fake/` contains fake implementations of every service client for use in tests. These implement the same interfaces as the real clients. Test utilities live in `pkg/connect/client/test/`.

### Key files

- `pkg/connect/client/clientset.go` - `ClientSet` interface + `New()` constructor
- `pkg/connect/client/config.go` - `Config` struct (ConnectURL + ConnectTrustDomain)
- `buf.gen.yaml` - Code generation config (4 plugins: protoc-gen-go, protoc-gen-connect-go, grpc/go, bufbuild/es)
- `buf.yaml` - Buf linting rules (STANDARD)
- `api-linter.yaml` - api-linter config, with the AIP rules the new APIs don't follow and why

### Adding a new service

When adding a new service, the typical flow is:
1. Add proto definition under `proto/cofide/<name>/<version>/` (new services belong with the new APIs, not under `proto/connect/` with the old ones)
2. Run `just proto-gen` to generate stubs in `gen/`
3. Create a wrapper client package at `pkg/connect/client/<service_name>/v1alpha1/`
4. Add fake implementation in `pkg/connect/client/fake/`
5. Register the new client in `ClientSet` interface and `clientSet` struct in `clientset.go`

## API linting

The new APIs under `proto/cofide/` follow some of Google's [AIPs](https://google.aip.dev), checked by [api-linter](https://linter.aip.dev). The old APIs under `proto/connect/` aren't checked. CI runs `just api-lint` and annotates the PR with any problems.

`just api-lint` lints all of `proto/cofide/`. Its first argument is a pattern, matched with grep against file names such as `cofide/cloud_account/v1alpha1/service.proto`, to lint only some of them, and further arguments are passed to api-linter:

```sh
just api-lint cloud_account
just api-lint . --output-format=json
```

When a rule fails, fix the proto if the rule fits. If it doesn't, either:
- disable it in `api-linter.yaml`, with a comment saying why, if it conflicts with a convention all Cofide APIs follow, or
- disable it with an AIP-192 internal comment beside the element, if the rule doesn't fit that one element but should still apply everywhere else. Our generators don't strip internal comments as Google's tooling does, so they appear in the generated code and docs like any other comment:
  ```proto
  // (-- api-linter: core::0158::response-plural-first-field=disabled
  //     aip.dev/not-precedent: The linter thinks this is a list response
  //     because it has a next_page_token field. --)
  ```

Read `api-linter.yaml` before designing a new resource, since it lists where Cofide APIs differ from the AIPs, such as identifying resources by `id` rather than by resource name.

### Resource IDs

A resource's `id` field is marked `OUTPUT_ONLY` and `IMMUTABLE`, not `IDENTIFIER`:

```proto
string id = 1 [
  (google.api.field_behavior) = OUTPUT_ONLY,
  (google.api.field_behavior) = IMMUTABLE
];
```

AIP-203 reserves `IDENTIFIER` for a resource's `name` field, which in the AIPs holds a resource name such as `organizations/{org}/trustZones/{trust_zone}`. Cofide resources are identified by a flat `id` instead, which Connect generates and which never changes, and `OUTPUT_ONLY` and `IMMUTABLE` say exactly that. Marking `id` as `IDENTIFIER` would break `core::0203::resource-identifier-only` and `core::0203::field-behavior-required`, and both would need disabling.

## Code Generation Prerequisites

`just proto-gen` needs Go and the [Buf CLI](https://buf.build/docs/installation). The Go plugins are pinned as tools in `tools/go.mod` and run with `go tool`, so they don't need installing; bump them there with `go -C tools get -tool <module>@<version>`. The TypeScript plugin is a remote plugin, pinned in `buf.gen.yaml`.

CI regenerates the code from scratch and fails if `gen/` differs from what is committed, so commit regenerated code alongside proto changes.
