build:
	go get ./cmd/
	go build -o bin/main ./cmd/

run:
	go run ./cmd/main.go