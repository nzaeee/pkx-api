.PHONY: build clean build-auth build-gateway run-all

# Build all services
build: build-auth build-gateway

# Build auth service
build-auth:
	@echo "Building auth service..."
	@cd auth && go build -o bin/auth ./cmd/api

# Build gateway service
build-gateway:
	@echo "Building API gateway..."
	@cd gateway && go build -o bin/gateway ./cmd/api

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf auth/bin gateway/bin

# Run auth service
run-auth:
	@cd auth && go run cmd/api/main.go

# Run gateway service
run-gateway:
	@cd gateway && go run cmd/api/main.go

# Run all services concurrently
run-all:
	@echo "Starting all services..."
	@cd auth && go run cmd/api/main.go & \
	cd gateway && go run cmd/api/main.go & \
	wait

# Install dependencies for all services
deps:
	@echo "Installing dependencies..."
	@cd auth && go mod tidy
	@cd gateway && go mod tidy 
