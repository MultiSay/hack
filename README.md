# hack
Moscow City Hack 2022

## Installation locally
### Prerequisites

- go 1.17
- golang-migrate cli (install on unix/macOs: brew install golang-migrate)

1. Clone this repository

2. Install dependencies
    ```bash
    $ go mod download
    $ migrate -path migrations -database "postgresql://postgres:user1user1@51.250.38.34:5432/db1?sslmode=disable" up
    $ cp .env.example .env
    ```

## Run app local

1. Start server locally
    ```bash
    $ make
    ```
2. Open in browser 
    ```
    http://51.250.38.34:8080/status
    ```

## Build

1. Build app
    ```bash
    $ make build
    ```
2. Для запуска на стенде
    ```bash
    $ GOOS=linux GOARCH=amd64 go build -o ./build/hack ./cmd/hack
    $ scp -i ~/.ssh/id_rsa ~/project/hack/build/hack mak@51.250.38.34:/home/mak
    $ scp -i ~/.ssh/id_rsa ~/project/hack/.env mak@51.250.38.34:/home/mak
    ```

## Testing

1. Start server tests
    ```bash
    $ make test
    ```