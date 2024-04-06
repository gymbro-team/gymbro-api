# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOCLEAN=$(GOCMD) clean
BINARY_NAME=gymbro-api
BINARY_DIR=bin
BUILD_FLAGS=-v

# Targets
all: test build

build: clean
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BINARY_DIR)
	@$(GOBUILD) $(BUILD_FLAGS) -o $(BINARY_DIR)/ ./...

test:
	@echo "Running tests..."
	@$(GOTEST) $(BUILD_FLAGS) ./...

clean:
	@echo "Cleaning..."
	@$(GOCLEAN)
	@rm -rf $(BINARY_DIR)

run: build
	@echo "Running $(BINARY_NAME)..."
	@$(BINARY_DIR)/$(BINARY_NAME)

.PHONY: all build test clean run
