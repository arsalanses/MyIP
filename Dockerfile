FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go .

RUN go build -o main .

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 3000

CMD ["./main"]
