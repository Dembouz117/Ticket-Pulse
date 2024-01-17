# Use the official Go image as the base image
FROM golang:1.21 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files to the container
COPY go.mod go.sum ./

# Download and cache the Go modules
RUN go mod download

# Copy the rest of the application source code to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd

# Use a minimal base image for the final image
FROM alpine:3.14

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder image to the final image
COPY --from=builder /app/app .

# Expose the port your application listens on
EXPOSE 8081

EXPOSE 5003

# Run your application
CMD ["./app"]
