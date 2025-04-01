# Stage 1: Build
FROM golang:1.23 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /server ./cmd/ordersystem/.

# Stage 2: Run
FROM alpine:3.14

# Set working directory
WORKDIR /

# Copy the compiled binary from the builder stage
COPY --from=builder /server /server
COPY cmd/ordersystem/.env .env

# Expose necessary ports
EXPOSE 8000
EXPOSE 8080
EXPOSE 50052

# Set the entry point
CMD ["/server"]
