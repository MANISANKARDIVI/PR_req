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
FROM golang@sha256:ad5c126b5cf501a8caef751a243bb717ec204ab1aa56dc41dc11be089fafcb4f AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files first to leverage caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the rest of the application code
COPY . .

# Build the Go application with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp main.go

# Final Stage
FROM scratch

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myapp /myapp

# Command to run the application
CMD ["/myapp"]
