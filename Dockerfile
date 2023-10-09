# Use an official Golang runtime as a parent image
FROM golang:1.20

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Install any needed dependencies
RUN go get ./...

# Build the Go application
RUN go build -o main ./cmd/server

# Expose port 8080 to the outside world
EXPOSE 8080

CMD ["./main"]
