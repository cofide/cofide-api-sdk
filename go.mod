module github.com/cofide/cofide-api-sdk

go 1.26.8

// NOTE: Take care to avoid forcing unnecessary upgrades on consumers when
// updating dependencies.
require (
	connectrpc.com/connect v1.19.2
	github.com/google/uuid v1.6.0
	github.com/spiffe/spire-api-sdk v1.12.4
	github.com/stretchr/testify v1.11.1
	google.golang.org/genproto/googleapis/api v0.0.0-20260526163538-3dc84a4a5aaa
	google.golang.org/grpc v1.83.2
	google.golang.org/protobuf v1.36.12
)

require github.com/spiffe/go-spiffe/v2 v2.7.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.58.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260526163538-3dc84a4a5aaa // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
