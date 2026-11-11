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

proto-docs out="docs":
    mkdir -p {{out}}
    buf generate --template buf.gen.docs.yaml --output {{out}}

# tools/ is a separate module, which ./... leaves out, so its tests are run on
# their own. The args, such as package patterns, are for the SDK's tests, so
# they aren't passed on.
#
# Run the Go tests, including those in tools/.
test *args:
    go run gotest.tools/gotestsum@latest --format github-actions ./... {{args}}
    go -C tools run gotest.tools/gotestsum@latest --format github-actions ./...

test-race: (test "--" "-race")

# tools/ is a separate module, which golangci-lint leaves out, so it's linted
# on its own. The args are for the SDK's lint, so they aren't passed on.
go-lint *args:
    golangci-lint run --show-stats {{args}}
    cd tools && golangci-lint run --show-stats

lint *args: buf-lint api-lint (go-lint args)
