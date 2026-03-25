# AGENTS.md

This file provides guidance to AI coding agents when working with code in this repository.

## Overview

This repository contains:
1. **Protocol Buffer definitions** (`proto/`) - service and message definitions for Cofide's APIs
2. **Generated code** (`gen/`) - auto-generated Go and TypeScript stubs from the proto files
3. **Go client library** (`pkg/connect/client/`) - hand-written wrapper clients for the Cofide Connect API

All APIs are versioned as `v1alpha1` (backward compatibility not guaranteed).

## Commands

Commands use [Just](https://github.com/casey/just) as the task runner.

| Command | Purpose |
|---|---|
| `just fmt` | Format proto files (`buf format -w`) |
| `just lint` | Lint proto files and Go source |
| `just buf-lint` | Lint proto files only |
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

1. **`proto/`** - Source of truth. 13 gRPC services defined under `proto/connect/`, plus domain message types across 20+ packages. All in `v1alpha1`.

2. **`gen/go/`** - Auto-generated, do not edit manually. Regenerate with `just proto-gen`. Mirrors the `proto/` directory structure.

3. **`pkg/connect/client/`** - Hand-written Go clients. Each service gets its own package (e.g., `pkg/connect/client/cluster/v1alpha1/`). The `ClientSet` interface in `clientset.go` aggregates all service clients into a single accessor. `New(conn grpc.ClientConnInterface)` constructs a `ClientSet`.

### Fake clients

`pkg/connect/client/fake/` contains fake implementations of every service client for use in tests. These implement the same interfaces as the real clients. Test utilities live in `pkg/connect/client/test/`.

### Key files

- `pkg/connect/client/clientset.go` - `ClientSet` interface + `New()` constructor
- `pkg/connect/client/config.go` - `Config` struct (ConnectURL + ConnectTrustDomain)
- `buf.gen.yaml` - Code generation config (4 plugins: protoc-gen-go, protoc-gen-connect-go, grpc/go, bufbuild/es)
- `buf.yaml` - Buf linting rules (STANDARD minus FIELD_NOT_REQUIRED and PACKAGE_NO_IMPORT_CYCLE)

### Adding a new service

When adding a new service, the typical flow is:
1. Add proto definition under `proto/connect/<service_name>_service/`
2. Run `just proto-gen` to generate stubs in `gen/`
3. Create a wrapper client package at `pkg/connect/client/<service_name>/v1alpha1/`
4. Add fake implementation in `pkg/connect/client/fake/`
5. Register the new client in `ClientSet` interface and `clientSet` struct in `clientset.go`

## Code Generation Prerequisites

Required tools for `just proto-gen`:
- `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.5`
- `go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest`
- [Buf CLI](https://buf.build/docs/installation)
