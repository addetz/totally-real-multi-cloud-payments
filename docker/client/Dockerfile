# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Copy source from current folder and handlers
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o client ./cmd/client/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/client .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./client"] 