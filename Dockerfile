FROM golang:1.18 AS build

WORKDIR /go/src/hack
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go mod download && go build -o ./dist/hackapp ./cmd/hack/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=build /go/src/hack/dist/hackapp /usr/bin/hackapp
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/hackapp"]
