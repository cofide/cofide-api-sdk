# cofide-api-sdk
This repository contains the service definitions and code generated stubs for [Cofide's](https://www.cofide.io/) APIs.

## Getting started
This repository uses the [Buf CLI](https://buf.build/docs/ecosystem/cli-overview) to help generate and manage Cofide's Protocol buffer (protobuf) definitions and schemas. To install the CLI, use the following [instructions](https://buf.build/docs/installation). For convenience, a set of useful commands have been added to the Justfile in the project root. If you haven't got [just](https://github.com/casey/just), you can follow these [instructions](https://github.com/casey/just) to do so.

Some of the key commands include:
- `just fmt` - Formats the protobuf definitions
- `just lint` - Lints the protobuf definitions in accordance with best practices
- `just proto-gen` - Generates the code stubs from the protobuf definitions using the defined protoc plugins
