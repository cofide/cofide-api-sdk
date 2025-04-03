# Cofide API SDK

[![Buf CI](https://github.com/cofide/cofide-api-sdk/workflows/buf-ci/badge.svg)](https://github.com/cofide/cofide-api-sdk/actions?query=workflow%3Abuf-ci+branch%3Amain)

This repository contains the service definitions and code generated stubs for [Cofide's](https://www.cofide.io/) APIs, used for open source projects (e.g. [`cofidectl`](https://github.com/cofide/cofidectl)) and the Cofide Connect product.
Services and messages are defined using Protocol Buffers (protobuf).
This repository also contains a Go client implementation for the Cofide Connect API.

## Prerequisites

This repository uses the [Buf CLI](https://buf.build/docs/ecosystem/cli-overview) to help generate and manage Cofide's protobuf definitions and schemas.
The following tools must be available in order to generate code stubs.

- [Go 1.22.6 toolchain](https://golang.org/doc/install)
- [protobuf version 3 compiler](https://grpc.io/docs/protoc-installation/)
- [Go gRPC and protobuf](https://grpc.io/docs/languages/go/quickstart/#prerequisites) build tools.
- [Connect RPC tools](https://connectrpc.com/docs/go/getting-started#prerequisites)
- [Buf CLI](https://buf.build/docs/installation)
- [Just](https://github.com/casey/just)

## Getting started

For convenience, a set of useful commands have been added to the *Justfile* in the project root.
Some of the key commands include:

- `just fmt` - Formats the protobuf definitions
- `just lint` - Lints the protobuf definitions and Go source code in accordance with best practices
- `just proto-gen` - Generates the code stubs from the protobuf definitions using the defined plugins (specified in *buf.gen.yaml*)
- `just test` - Runs Go unit tests

The `.proto` files are in the `proto` directory, and the generated stubs for Go are in the `gen/proto/go` directory.
The Cofide Connect API client is in the `pkg/connect/client` directory.

## Stability

The service definitions in this repository are not currently guaranteed to be backward compatible over time, and have been versioned as `v1alpha1` to indicate this.
As Cofide projects and product mature, we will move to `v1` and guarantee backward compatibility from that point.
