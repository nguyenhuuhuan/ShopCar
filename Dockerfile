# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app

COPY . .

RUN cd src/cmd && \
    CGO_ENABLED=0 go get -u github.com/pressly/goose/cmd/goose && \
    go build -o main main.go

# Run stage
FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/src/cmd/main .
COPY wait-for.sh .
EXPOSE 8080

CMD ["/app/main", "/app/migration.sh"]