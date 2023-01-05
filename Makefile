all: update module test example

update:
	go mod vendor
	go mod tidy

module:
	go build

test:
	go test -tags ledger_mock

.PHONY: example
example:
	go build -o example-ledger-go example/example.go

clean:
	rm -rf example-ledger-go
	rm -rf vendor
