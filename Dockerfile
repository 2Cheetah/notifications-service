# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o notifications-service .

# Production stage
FROM scratch
WORKDIR /app
COPY --from=builder /app/notifications-service .

CMD ["./notifications-service"]
