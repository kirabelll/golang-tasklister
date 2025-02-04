run:
	@echo "Running tests"
	@go test ./...
	@echo "Running app"
	@go run cmd/main.go

build:
	@go build cmd/main.go -o bin/tasklist

clean:
	@rm -rf bin