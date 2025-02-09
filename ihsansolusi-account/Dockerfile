# ======= STAGE 1: BUILD THE APPLICATION =======
FROM golang:1.22 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app (disable CGO for better compatibility)
RUN CGO_ENABLED=0 go build -o app ./main.go

# ======= STAGE 2: RUN THE APPLICATION =======
FROM alpine:latest

# Install necessary dependencies (CA certificates, curl, and PostgreSQL client)
RUN apk --no-cache add ca-certificates curl postgresql-client

# Set working directory
WORKDIR /app

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/app /app/

# Copy the environment file (jika menggunakan env file)
COPY .env /app/.env

# Copy the migrations folder
COPY migrations /app/migrations

# 🔥 Download dan install migrate CLI
RUN curl -fsSL -o migrate.tar.gz https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz && \
    tar -xzf migrate.tar.gz -C /usr/local/bin/ && \
    rm migrate.tar.gz && \
    chmod +x /usr/local/bin/migrate

# Expose the port used by the application
EXPOSE 3001

# Command to run the application
CMD ["/app/app"]
