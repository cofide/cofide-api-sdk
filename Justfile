PROTOC_GEN_GO_VERSION := "v1.36.5"
GOBIN := `go env GOPATH`+ "/bin"
PROTOC_GEN_GO := GOBIN + "/protoc-gen-go"
PROTOC_GEN_CONNECT_GO := GOBIN + "/protoc-gen-connect-go"
PROTOC_GEN_DOC := GOBIN + "/protoc-gen-doc"

[private]
ensure-protoc-gen-go:
    #!/usr/bin/env bash
    if [ ! -f "{{PROTOC_GEN_GO}}" ] || [ ! "$({{PROTOC_GEN_GO}} --version)" = "protoc-gen-go {{PROTOC_GEN_GO_VERSION}}" ]; then
        echo "Please install protoc-gen-go {{PROTOC_GEN_GO_VERSION}}:"
        echo
        echo "  go install google.golang.org/protobuf/cmd/protoc-gen-go@{{PROTOC_GEN_GO_VERSION}}"
        echo
        exit 1
    fi

[private]
ensure-protoc-gen-connect-go:
    #!/usr/bin/env bash
    if [ ! -f "{{PROTOC_GEN_CONNECT_GO}}" ]; then
        echo "Please install protoc-gen-connect-go"
        echo
        echo " go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest"
        echo
        exit 1
    fi

fmt:
    buf format -w

buf-lint:
    buf lint

# Lint the new APIs against Google's AIPs. Only files whose path matches the
# pattern are linted, which by default is all of the new APIs in proto/cofide.
api-lint pattern="." *args:
    #!/usr/bin/env bash
    set -euo pipefail
    # api-linter is pinned as a tool in tools/go.mod, which keeps its
    # dependencies out of the SDK's own go.mod. `go tool -n` builds it if
    # needed and prints its path, so it can run from the repo root.
    api_linter=$(go -C tools tool -n api-linter)
    # buf reads the image's format from its extension, and BSD mktemp can't
    # add one, so the image goes in a temporary directory instead.
    dir=$(mktemp -d)
    trap 'rm -rf "$dir"' EXIT
    image="$dir/image.binpb"
    # Build a descriptor set to ensure api-linter can find the proto
    # dependencies.
    buf build --path proto/cofide -o "$image"
    # Listing the image gives the names api-linter expects, such as
    # cofide/federation/v1beta1/service.proto, without the imports.
    files=$(buf ls-files "$image" | grep -e '{{pattern}}' || true)
    if [ -z "$files" ]; then
        echo "No files under proto/cofide match the pattern: {{pattern}}" >&2
        exit 1
    fi
    "$api_linter" --config api-linter.yaml --descriptor-set-in "$image" --set-exit-status {{args}} $files

proto-gen: ensure-protoc-gen-go ensure-protoc-gen-connect-go
    buf generate

[private]
ensure-protoc-gen-doc:
    #!/usr/bin/env bash
    if [ ! -f "{{PROTOC_GEN_DOC}}" ]; then
        echo "Please install protoc-gen-doc:"
        echo
        echo "  go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest"
        echo
        exit 1
    fi

proto-docs out="docs": ensure-protoc-gen-doc
    mkdir -p {{out}}
    buf generate --template buf.gen.docs.yaml --output {{out}}

test *args:
    go run gotest.tools/gotestsum@latest --format github-actions ./... {{args}}

test-race: (test "--" "-race")

go-lint *args:
    golangci-lint run --show-stats {{args}}

lint *args: buf-lint api-lint (go-lint args)
