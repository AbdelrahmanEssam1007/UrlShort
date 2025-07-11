# Use Go 1.24.2 (or newer)
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# Optional: smaller runtime image (multi-stage build)
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 9808
CMD ["./main"]
