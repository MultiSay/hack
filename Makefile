
.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o ./dist/hack ./cmd/hack

.PHONY: test
test:
	go fmt ./internal/app/...
	go vet -composites=false ./internal/app/...
	go test -cover -v -timeout 30s ./internal/app/...

.PHONY: run
run: 
	go run ./...

.PHONE: migrate-create
migrate-create:
	migrate create -ext sql -dir migrations/ -seq $(name)

.PHONY: swagger
swagger: 
	go get github.com/swaggo/swag && ~/go/bin/swag init -g cmd/hack/main.go --exclude dist/

.DEFAULT_GOAL := run
