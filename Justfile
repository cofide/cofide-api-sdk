GOBIN := `go env GOPATH`+ "/bin"
PROTOC_GEN_DOC := GOBIN + "/protoc-gen-doc"

fmt:
    buf format -w

buf-lint:
    buf lint

# Lint the new APIs against Google's AIPs. Only files whose path matches the
# pattern are linted, which by default is all of the new APIs in proto/cofide.
api-lint pattern="." *args:
    #!/usr/bin/env bash
    set -euo pipefail
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
    # Only grep may fail without stopping the script, so that a failure from
    # buf still does.
    files=$(buf ls-files "$image" | { grep -e '{{pattern}}' || true; })
    if [ -z "$files" ]; then
        echo "No files under proto/cofide match the pattern: {{pattern}}" >&2
        exit 1
    fi
    # api-linter is pinned as a tool in tools/go.mod, which keeps its
    # dependencies out of the SDK's own go.mod. It runs in tools/, so the
    # config needs an absolute path. The files are names in the image, not
    # paths, so they don't.
    go -C tools tool api-linter --config "$PWD/api-linter.yaml" --descriptor-set-in "$image" --set-exit-status {{args}} $files

proto-gen:
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
