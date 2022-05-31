gomod:
	go mod tidy -compat=1.16 && go mod tidy -compat=1.17
	go mod download

test:
	go test -cover ./...