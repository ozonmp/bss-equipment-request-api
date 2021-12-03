GO_VERSION_SHORT:=$(shell echo `go version` | sed -E 's/.* go(.*) .*/\1/g')
ifneq ("1.16","$(shell printf "$(GO_VERSION_SHORT)\n1.16" | sort -V | head -1)")
$(error NEED GO VERSION >= 1.16. Found: $(GO_VERSION_SHORT))
endif

export GO111MODULE=on

SERVICE_NAME=bss-equipment-request-api
SERVICE_PATH=ozonmp/bss-equipment-request-api

PGV_VERSION:="v0.6.1"
BUF_VERSION:="v1.0.0-rc3"

OS_NAME=$(shell uname -s)
OS_ARCH=$(shell uname -m)
GO_BIN=$(shell go env GOPATH)/bin
BUF_EXE=$(GO_BIN)/buf$(shell go env GOEXE)

ifeq ("NT", "$(findstring NT,$(OS_NAME))")
OS_NAME=Windows
endif

.PHONY: run
run:
	go run cmd/grpc-server/main.go

.PHONY: lint
lint:
	golangci-lint run ./...


.PHONY: test
test:
	go test -v -race -timeout 30s -coverprofile cover.out ./...
	go tool cover -func cover.out | grep total | awk '{print $$3}'

.PHONY: migrate
migrate:
	go run cmd/migration/main.go
# ----------------------------------------------------------------

.PHONY: generate
generate: .generate-install-buf .generate-go .generate-python .generate-finalize-go .generate-finalize-python

.PHONY: generate
generate-go: .generate-install-buf .generate-go .generate-finalize-go

.generate-install-buf:
	@ command -v buf 2>&1 > /dev/null || (echo "Install buf" && \
    		mkdir -p "$(GO_BIN)" && \
    		curl -sSL0 https://github.com/bufbuild/buf/releases/download/$(BUF_VERSION)/buf-$(OS_NAME)-$(OS_ARCH)$(shell go env GOEXE) -o "$(BUF_EXE)" && \
    		chmod +x "$(BUF_EXE)")

.generate-go:
	$(BUF_EXE) generate

.generate-python:
	$(BUF_EXE) generate --template buf.gen.python.yaml

.generate-finalize-go:
	mv pkg/$(SERVICE_NAME)/github.com/$(SERVICE_PATH)/pkg/$(SERVICE_NAME)/* pkg/$(SERVICE_NAME)
	rm -rf pkg/$(SERVICE_NAME)/github.com/
	cd pkg/$(SERVICE_NAME) && ls go.mod || (go mod init github.com/$(SERVICE_PATH)/pkg/$(SERVICE_NAME) && go mod tidy)

.generate-finalize-python:
	find pypkg/bss-equipment-request-api -type d -exec touch {}/__init__.py \;

# ----------------------------------------------------------------

.PHONY: deps
deps: deps-go .deps-python

.PHONY: deps-go
deps-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
	go install github.com/envoyproxy/protoc-gen-validate@$(PGV_VERSION)
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest

.deps-python:
	python3 -m pip install grpcio-tools grpclib protobuf --user

.PHONY: build
build: generate .build .build-migration .build-retranslator

.PHONY: build-go
build-go: generate-go .build

.build:
	go mod download && CGO_ENABLED=0  go build \
		-tags='no_mysql no_sqlite3' \
		-ldflags=" \
			-X 'github.com/$(SERVICE_PATH)/internal/config.version=$(VERSION)' \
			-X 'github.com/$(SERVICE_PATH)/internal/config.commitHash=$(COMMIT_HASH)' \
		" \
		-o ./bin/grpc-server$(shell go env GOEXE) ./cmd/grpc-server/main.go

.PHONY: build-migration-go
build-migration-go: .build-migration

.build-migration:
	go mod download && CGO_ENABLED=0  go build \
    		-tags='no_mysql no_sqlite3' \
    		-ldflags=" \
    			-X 'github.com/$(SERVICE_PATH)/internal/config.version=$(VERSION)' \
    			-X 'github.com/$(SERVICE_PATH)/internal/config.commitHash=$(COMMIT_HASH)' \
    		" \
    		-o ./bin/migration$(shell go env GOEXE) ./cmd/migration/main.go

.PHONY: build-retranslator-go
build-retranslator-go: .build-retranslator

.build-retranslator:
	go mod download && CGO_ENABLED=0  go build \
    		-tags='no_mysql no_sqlite3' \
    		-ldflags=" \
    			-X 'github.com/$(SERVICE_PATH)/internal/retranslator/config.version=$(VERSION)' \
    			-X 'github.com/$(SERVICE_PATH)/internal/retranslator/config.commitHash=$(COMMIT_HASH)' \
    		" \
    		-o ./bin/retranslator$(shell go env GOEXE) ./cmd/retranslator/main.go
