# Start with a base image containing the Go runtime
FROM golang:1.22-alpine3.19 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage to the new stage
COPY --from=builder /app/main .

# Define environment variables
# Environment: dev or prod
ENV ENVIRONMENT=dev \
    SERVER_PORT=9090 \
    REDIS_HOST=localhost \
    REDIS_PORT=6379 \
    REDIS_PASSWORD= \
    REDIS_EXPIRATION_TIME_KEY=60

# Expose port 9090 to the outside world
EXPOSE 9090

# Command to run the executable
CMD ["./main"]