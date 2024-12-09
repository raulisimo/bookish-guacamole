# Backend Dockerfile
FROM golang:1.22 AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy and build the application
COPY . .
RUN go build -o main ./cmd/main.go

# Final stage
FROM gcr.io/distroless/base
WORKDIR /root/

COPY --from=builder /app/main .

# Expose the backend port
EXPOSE 8080

# Run the backend
CMD ["./main"]
