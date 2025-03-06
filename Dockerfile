# Use a Go base image
FROM golang:1.23.5-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum
COPY ./go.mod ./go.sum ./

# Download Go dependencies (this will create vendor dir and download dependencies based on go.mod)
RUN go mod tidy

# Copy the rest of the Go source code
COPY . .

# Build the Go app (change the entry point if needed)
RUN go build -o /go-app ./cmd/main.go

# Final image to run the Go app
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the Go app from the build stage
COPY --from=build /go-app .

# Expose the port that your API is running on
EXPOSE 8080

# Start the app
CMD ["./go-app"]
