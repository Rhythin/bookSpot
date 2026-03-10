# Use the official Golang image to build the service
FROM golang:1.26-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the workspace configuration and shared library
# We need the shared library because the service depends on it
COPY go.work go.work.sum ./
COPY services/shared/ ./services/shared/

# Copy the auth-service source code
COPY services/auth-service/ ./services/auth-service/

# Build the application
# We use the full path in the workspace for the main package
RUN go build -o /app/auth-service ./services/auth-service/cmd/server.go

# Use a minimal alpine image for the runtime
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/auth-service .

# Copy the sample env file as .env for the service to use (optional)
# COPY --from=builder /app/services/auth-service/.env.example .env

# Expose the service port
EXPOSE 8080

# Run the binary
CMD ["./auth-service"]
