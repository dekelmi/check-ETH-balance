FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o eth-check .
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/eth-check .
COPY .env .
EXPOSE 5090
CMD ["./eth-check"]