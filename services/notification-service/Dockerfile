# Use the official Golang image to build the service
FROM golang:1.26-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the workspace configuration and shared library
COPY go.work go.work.sum ./
COPY services/shared/ ./services/shared/

# Copy the notification-service source code
COPY services/notification-service/ ./services/notification-service/

# Build the application
RUN go build -o /app/notification-service ./services/notification-service/cmd/server.go

# Use a minimal alpine image for the runtime
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/notification-service .

# Expose the service port
EXPOSE 8082

# Run the binary
CMD ["./notification-service"]
