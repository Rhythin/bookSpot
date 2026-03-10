# Use the official Golang image to build the service
FROM golang:1.26-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the workspace configuration and shared library
COPY go.work go.work.sum ./
COPY services/shared/ ./services/shared/

# Copy the books-service source code
COPY services/books-service/ ./services/books-service/

# Build the application
RUN go build -o /app/books-service ./services/books-service/cmd/server.go

# Use a minimal alpine image for the runtime
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/books-service .

# Expose the service port
EXPOSE 8081

# Run the binary
CMD ["./books-service"]
