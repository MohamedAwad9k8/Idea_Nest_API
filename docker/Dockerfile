# Build Stage
FROM golang:latest AS builder
WORKDIR /app

# # Copy only the necessary files for dependency resolution
# COPY go.mod go.sum ./

# Copy the entire project
COPY . .

# Download dependencies
RUN go mod download

# # Copy the entire project
# COPY . .

# Build the Golang application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Final Stage
FROM alpine:latest
WORKDIR /app

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/main /app/main

# Expose the port on which your application will run
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

## docker build -t org-management-app-image -f docker/Dockerfile .
## docker run -it --rm -p 8080:8080 org-management-app-image