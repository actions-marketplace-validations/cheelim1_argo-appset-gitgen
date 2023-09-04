# Use the official Golang image from the DockerHub
FROM golang:1.17-alpine

# Install git, required for fetching Go dependencies
RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files to the app directory
COPY go.mod go.sum ./

# Download all Go dependencies
RUN go mod download

# Copy the source code as the last Docker layer because it changes the most
COPY main.go .

# Build the Go app for a specific OS and architecture
RUN GOOS=linux GOARCH=amd64 go build -o /argo-appset-gitgen main.go

# Run the executable
ENTRYPOINT ["/argo-appset-gitgen"]
