module github.com/cofide/cofide-api-sdk

go 1.24.11

// NOTE: Take care to avoid forcing unnecessary upgrades on consumers when
// updating dependencies.
require (
	connectrpc.com/connect v1.19.1
	github.com/google/uuid v1.6.0
	github.com/spiffe/spire-api-sdk v1.12.4
	github.com/stretchr/testify v1.10.0
	google.golang.org/grpc v1.73.1
	google.golang.org/protobuf v1.36.11
)

require google.golang.org/genproto/googleapis/api v0.0.0-20260114163908-3f89685c29c3

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251222181119-0a764e51fe1b // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
