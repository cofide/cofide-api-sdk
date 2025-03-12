PROTOC_GEN_GO_VERSION := "v1.36.5"
GOBIN := `go env GOPATH`+ "/bin"
PROTOC_GEN_GO := GOBIN + "/protoc-gen-go"

[private]
ensure-protoc-gen-go:
    #!/usr/bin/env bash
    if [ ! -f "{{PROTOC_GEN_GO}}" ] || [ ! "$({{PROTOC_GEN_GO}} --version)" = "protoc-gen-go {{PROTOC_GEN_GO_VERSION}}" ]; then
        echo "Please install protoc-gen-go {{PROTOC_GEN_GO_VERSION}}:"
        echo
        echo "  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest"
        # See https://github.com/golang/protobuf/issues/1090
        echo
        echo "Note that it does not appear to be possible to install a specific version using go install."
        echo "If the pinned version is no longer the latest you may need to build from source or update the pinned version."
        exit 1
    fi

fmt:
    buf format -w --path proto

buf-lint:
    buf lint --path proto

proto-gen: ensure-protoc-gen-go
    buf generate --path ./proto

test *args:
    go run gotest.tools/gotestsum@latest --format github-actions ./... {{args}}

test-race: (test "--" "-race")

go-lint *args:
    golangci-lint run --show-stats {{args}}

lint *args: buf-lint (go-lint args)
