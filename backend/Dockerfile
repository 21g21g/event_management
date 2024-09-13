
# FROM golang:1.18-alpine AS build

# WORKDIR /app
# COPY . .

# RUN go mod download
# RUN go build -o /app/start

# FROM alpine:latest

# WORKDIR /app
# COPY --from=build /app/start .
# EXPOSE 4000
# CMD ["./start"]
# Use a base image
FROM golang:1.23

# Set the working directory
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download Go modules dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN go build -o main .

# Run the application
CMD ["./main"]
