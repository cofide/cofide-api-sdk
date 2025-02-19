PROTOC_GEN_GO_VERSION := "v1.35.1"
GOBIN := `go env GOPATH`+ "/bin"
PROTOC_GEN_GO := GOBIN + "/protoc-gen-go"

[private]
ensure-protoc-gen-go:
    #!/usr/bin/env bash
    if [ ! -f "{{PROTOC_GEN_GO}}" ] || [ ! "$({{PROTOC_GEN_GO}} --version)" = "protoc-gen-go {{PROTOC_GEN_GO_VERSION}}" ]; then
        echo "Please install protoc-gen-go {{PROTOC_GEN_GO_VERSION}}:"
        echo
        echo "  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"
        # See https://github.com/golang/protobuf/issues/1090
        echo
        echo "Note that it does not appear to be possible to install a specific version using go install."
        echo "If the pinned version is no longer the latest you may need to build from source or update the pinned version."
        exit 1
    fi

fmt:
    buf format -w --path proto

lint:
    buf lint --path proto

proto-gen: ensure-protoc-gen-go
    buf generate --path ./proto
