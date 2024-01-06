# Start from a base image with Go installed
FROM golang:1.21
LABEL authors="krishnansriram"
# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Compile the application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080