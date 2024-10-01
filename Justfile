fmt:
    buf format -w --path proto

lint:
    buf lint --path proto

proto-gen:
    buf generate --path ./proto
