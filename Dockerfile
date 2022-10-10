# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app

COPY . .

RUN cd ${PWD}/src/cmd && \
    go build -o main main.go

# Run stage
FROM alpine:3.16

WORKDIR /app
COPY --from=builder /app/src/cmd/main .
COPY .env .
EXPOSE 8080

CMD ["/app/main"]