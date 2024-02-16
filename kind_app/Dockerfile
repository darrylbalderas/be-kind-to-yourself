# Stage 1: Build the Go binary
FROM golang:1.22 AS builder

WORKDIR /app

# Copy the Go module files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create a minimal runtime image
FROM gcr.io/distroless/static-debian11
# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"]
