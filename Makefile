BINARY_NAME = Todoapp

# Default target to run the application
.PHONY: run
run:
	@echo "Running the application..."
	go run cmd/main.go
