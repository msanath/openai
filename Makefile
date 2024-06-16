BUILD_DIR := build
BIN_DIR := $(BUILD_DIR)/bin

.PHONY: clean-build
clean-build: clean build

.PHONY: build
build: gomodtidy
	mkdir -p $(BIN_DIR)
	@echo "Building the project..."
	@go build -o $(BIN_DIR)/openai cmd/main.go

.PHONY: gomodtidy
gomodtidy:
	@echo "Tidying go modules..."
	go mod tidy

.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

.PHONE: clean
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR) 2>/dev/null || true
	@rm go.sum 2>/dev/null || true
