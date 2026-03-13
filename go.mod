module github.com/cofide/cofide-api-sdk

go 1.25.8

// NOTE: Take care to avoid forcing unnecessary upgrades on consumers when
// updating dependencies.
require (
	connectrpc.com/connect v1.19.1
	github.com/google/uuid v1.6.0
	github.com/spiffe/spire-api-sdk v1.12.4
	github.com/stretchr/testify v1.10.0
	google.golang.org/genproto/googleapis/api v0.0.0-20260114163908-3f89685c29c3
	google.golang.org/grpc v1.73.1
	google.golang.org/protobuf v1.36.11
)

require github.com/spiffe/go-spiffe/v2 v2.5.0

require (
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-jose/go-jose/v4 v4.0.5 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/zeebo/errs v1.4.0 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251222181119-0a764e51fe1b // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
