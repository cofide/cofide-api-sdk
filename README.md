# Cofide API SDK

[![Buf CI](https://github.com/cofide/cofide-api-sdk/workflows/buf-ci/badge.svg)](https://github.com/cofide/cofide-api-sdk/actions?workflow=buf-ci+branch%3Amain)

This repository contains the service definitions and code generated stubs for [Cofide's](https://www.cofide.io/) APIs.

## Getting started
This repository uses the [Buf CLI](https://buf.build/docs/ecosystem/cli-overview) to help generate and manage Cofide's Protocol buffer (protobuf) definitions and schemas. To install the CLI, use the following [instructions](https://buf.build/docs/installation). For convenience, a set of useful commands have been added to the *Justfile* in the project root. If you haven't got [just](https://github.com/casey/just), you can follow these [instructions](https://github.com/casey/just) to install it.

Some of the key commands include:
- `just fmt` - Formats the protobuf definitions
- `just lint` - Lints the protobuf definitions in accordance with best practices
- `just proto-gen` - Generates the code stubs from the protobuf definitions using the defined plugins (specified in *buf.gen.yaml*)

The `.proto` files are in the `proto` directory, and the generated stubs are in the `gen` directory.

## Stability
The service definitions in this repository are not currently guaranteed to be backward compatible over time, and have been versioned as `v1alhpa1` to indicate this.
As the Cofide product matures, we will move to `v1` and guarantee backward compatibility from that point.
