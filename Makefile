test:
	go test -tags ledger_mock

.PHONY: example
example:
	go build -o example-ledger-go example/example.go
