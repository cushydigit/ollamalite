.PHONY: run build
build:
	@echo "Building..."
	@go build -o ./bin/examples ./examples/main.go
run: build
	@echo "Running..."
	@./bin/examples
