module github.com/cofide/cofide-api-sdk/tools

go 1.26.11

tool (
	connectrpc.com/connect/cmd/protoc-gen-connect-go
	github.com/cofide/cofide-api-sdk/tools/cmd/protoc-gen-doc
	github.com/googleapis/api-linter/v2/cmd/api-linter
	google.golang.org/grpc/cmd/protoc-gen-go-grpc
	google.golang.org/protobuf/cmd/protoc-gen-go
)

require (
	github.com/pseudomuto/protoc-gen-doc v1.5.1
	github.com/pseudomuto/protokit v0.3.0
	google.golang.org/protobuf v1.36.12
)

require (
	bitbucket.org/creachadair/stringset v0.0.12 // indirect
	cloud.google.com/go/longrunning v1.4.1 // indirect
	connectrpc.com/connect v1.19.2 // indirect
	github.com/Masterminds/semver v1.4.2 // indirect
	github.com/Masterminds/sprig v2.15.0+incompatible // indirect
	github.com/aokoli/goutils v1.0.1 // indirect
	github.com/bmatcuk/doublestar/v4 v4.10.4 // indirect
	github.com/bufbuild/protocompile v0.14.1 // indirect
	github.com/gertd/go-pluralize v0.2.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/api-linter/v2 v2.4.1 // indirect
	github.com/huandu/xstrings v1.0.0 // indirect
	github.com/imdario/mergo v0.3.4 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/stoewer/go-strcase v1.3.1 // indirect
	golang.org/x/crypto v0.58.0 // indirect
	golang.org/x/net v0.61.0 // indirect
	golang.org/x/sync v0.24.0 // indirect
	golang.org/x/sys v0.49.0 // indirect
	golang.org/x/text v0.43.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20261026165753-2ffab32bce7e // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20261019180522-ce908d29299b // indirect
	google.golang.org/grpc v1.84.0 // indirect
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.5.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
