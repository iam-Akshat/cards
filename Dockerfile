# Use the official Go image as the base stage
FROM golang:1.18-alpine as builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY src/go.mod src/go.sum .env ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY src .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
# Expose the API port
EXPOSE 3000

# Run the binary
CMD ["./main"]