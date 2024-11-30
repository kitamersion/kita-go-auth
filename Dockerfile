FROM docker.io/golang:1.23-alpine AS builder

# Install build tools and dependencies
RUN apk add --no-cache git

# Set the working directory
WORKDIR /usr/src/app

# Pre-cache Go dependencies (minimizing rebuilds)
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy only the application source code
COPY . .

# Build the Go application statically (no CGO dependencies)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app .

# Use a minimal image for runtime
FROM docker.io/alpine:latest

# Install SSL certificates (required for Go apps with HTTPS)
RUN apk --no-cache add ca-certificates

# Set the working directory for the runtime container
WORKDIR /root/

# Copy the built binary and .env from the builder stage
COPY --from=builder /app /usr/local/bin/app
COPY --from=builder /usr/src/app/env_template /root/.env

# Expose the port your app listens on
EXPOSE 3001

# Command to run the application
CMD ["/usr/local/bin/app"]

