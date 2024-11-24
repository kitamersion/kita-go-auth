# Use the official Go image as the base image for building the app
FROM docker.io/golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /usr/src/app

# Pre-cache Go dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the application source code
COPY . .

# Copy env_template as .env (if the application needs it)
COPY env_template .env

# Build the Go application statically (no CGO dependencies)
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /usr/local/bin/app .

# Use a minimal image for runtime
FROM docker.io/alpine:latest

# Install SSL certificates (required for Go apps with HTTPS)
RUN apk --no-cache add ca-certificates

# Set the working directory for the runtime container
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /usr/local/bin/app /usr/local/bin/app

# Copy the .env file from the builder stage to runtime container
COPY --from=builder /usr/src/app/.env /root/.env

# Expose the port your app listens on
EXPOSE 3001

# Command to run the application
CMD ["/usr/local/bin/app"]

