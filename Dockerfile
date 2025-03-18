# Use the official Golang image as a build stage
FROM golang:1.18 AS builder

# Set the working directory inside the container
WORKDIR /app

# required if you have issues with proxy.golang.org being blocked
ENV GOPROXY=direct  

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main ./cmd

# Use a minimal image for the final stage
FROM alpine:latest

# Install necessary CA certificates
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Set environment variables
ENV PORT=4000
ENV ENV=production
ENV DB_DATA_SOURCE=postgres://admin@db/games_shelf?sslmode=disable
ENV APP_SECRET=games-shelf-api-secret
ENV RAWG_API_KEY=your_rawg_api_key

# Expose the port the app runs on
EXPOSE 4000

# Command to run the executable
CMD ["./main"]
