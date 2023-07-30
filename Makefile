server:
	go build ./...

proto:
	./scripts/compile_proto.pl proto

dev:
	@nodemon --exec go run cmd/main.go --signal SIGTERM

test:
	go vet ./...
	go test  -v -coverpkg ./internal/... -coverprofile coverage.out -covermode count ./internal/...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

compose-up:
	docker-compose up -d

compose-down:
	docker-compose down

seed:
	go run ./cmd/. seed

.PHONY: proto test server compose-up compose-down seed