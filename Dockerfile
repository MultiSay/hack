FROM golang:1.18 AS build

WORKDIR /go/src/hack
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN ./pkg/migrate.linux-amd64 -path ./migrations -database "postgresql://user:password@database:5432/hack?sslmode=disable" up
RUN go mod download && go build -o ./dist/hackapp ./cmd/hack/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=build /go/src/hack/dist/hackapp /usr/bin/hackapp
ENTRYPOINT ["/usr/bin/hackapp"]
