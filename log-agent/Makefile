lint:
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...

test:
	go test -v -cover ./...
