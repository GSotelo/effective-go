.PHONY: run build test test-variadic test-interfaces clean

# Run all snippets
run:
	go run .

# Build the binary
build:
	go build -o bin/snippets .

# Test all packages
test:
	go test ./... -v

# Test individual packages
test-variadic:
	go test ./variadic -v

test-interfaces:
	go test ./interfaces -v

# Clean build artifacts
clean:
	rm -rf bin/