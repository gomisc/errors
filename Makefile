gomod:
	go mod tidy
	go mod download

test:
	go test -cover ./...