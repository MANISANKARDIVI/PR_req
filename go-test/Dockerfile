# ## Build Stage
# FROM golang:latest AS builder

# # Set the working directory
# WORKDIR /app

# # Copy go.mod and go.sum files first to leverage caching
# # COPY go.mod go.sum ./
# COPY go-test/go.mod go-test/go.sum ./

# # Download dependencies
# RUN go mod download

# # Copy the rest of the application code
# COPY . .

# # Build the Go application with static linking
# RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o myapp main.go

# # Final Stage
# FROM scratch

# # Copy the compiled binary from the builder stage
# COPY --from=builder /app/myapp /myapp

# # Command to run the application
# CMD ["/myapp"]
# Build Stage
FROM golang:latest AS builder

# Install CGO dependencies if CGO_ENABLED=1 is required
RUN apt-get update && apt-get install -y gcc libc6-dev

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files first to leverage caching
COPY go-test/go.mod go-test/go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application (adjust build flags as needed)
RUN CGO_ENABLED=1 GOOS=linux go build -o myapp main.go

# Final Stage
FROM scratch

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myapp /myapp

# Command to run the application
CMD ["/myapp"]
