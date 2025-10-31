BINARY_NAME=calendar
LOCAL_BIN=$(CURDIR)/bin
MAIN_FILE=cmd/main.go

install-deps:
	GOBIN=$(LOCAL_BIN) go install golang.org/x/lint/golint@latest

run: build
	$(LOCAL_BIN)/$(BINARY_NAME)

build:
	go build -o $(LOCAL_BIN)/$(BINARY_NAME) $(MAIN_FILE)


check: fmt vet lint sort_import
	go mod tidy

sort_import:
	goimports -w cmd/* internal/*

vet:
	go vet ./...

lint: install-deps
	$(LOCAL_BIN)/golint ./...

fmt:
	go fmt ./...

test:
	cd internal/ && go test -v ./...