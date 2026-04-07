# Stage 1: Build stage
# Use the official Golang alpine image as the build environment
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy dependency files first to leverage Docker cache
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application as a static binary for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Final stage (Runtime)
# Use a lightweight Alpine Linux image for the final container
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]