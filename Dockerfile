# Start from a lightweight base image
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY ./go.mod ./go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o go-api ./cmd/server/main.go

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./go-api"]
