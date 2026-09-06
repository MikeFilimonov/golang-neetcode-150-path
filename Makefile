.PHONY: test lint fmt
test:
	go test ./...
lint:
	go vet ./...
	test -z "$$(gofmt -l .)"
fmt:
	gofmt -w .