# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder
# Set the working directory inside the container
WORKDIR /app
COPY . .
# Build the Go application
RUN go build -o upload-service main.go

# Stage 2: Create a small image to run the application
FROM alpine:3.18
# Install certificates (necessary if your app makes HTTPS requests)
RUN apk add --no-cache ca-certificates
# Set the working directory inside the container
WORKDIR /app
# Copy the binary from the builder stage
COPY --from=builder /app /app
# Ensure the uploads directory exists
RUN mkdir -p /app/uploads
# Set environment variables
ENV PORT 8080
# Run the Go binary
CMD ["./upload-service"]

#docker run -d -p ${PORT}:${PORT} -e PORT=9090 -v $(pwd)/uploads:/app/uploads upload-service