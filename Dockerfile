FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o echo-server ./src/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/echo-server .

EXPOSE 50000

CMD ["./echo-server"]
