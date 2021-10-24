.PHONY: build
build:
	go build cmd/bss-equipment-request-api/main.go

.PHONY: test
test:
	go test -v ./...