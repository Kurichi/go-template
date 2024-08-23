
.PHONY: build
build: BIN_DIR ?= ./bin
build:
	go build -o $(BIN_DIR)/main ./cmd/server/main.go