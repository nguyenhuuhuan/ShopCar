# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app

COPY . .

RUN cd ${PWD}/src/cmd && \
    CGO_ENABLED=0 go get -u github.com/pressly/goose/cmd/goose && \
    go build -o main main.go

RUN apk add curl && \
    apk add make
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz && \
    go install github.com/pressly/goose/v3/cmd/goose@latest



# Run stage
FROM alpine:3.16

WORKDIR /app

COPY --from=builder /app/src/cmd/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY --from=builder /app/goose ./goose
COPY src/migrations/goose /app/src/migrations/goose
COPY scripts/migration.sh .
COPY wait-for.sh .
EXPOSE 8080

CMD ["/app/main"]
ENTRYPOINT ["/app/migration.sh"]