# Use Go image as image base
FROM golang:1.20-alpine

# Set environment variables
ENV GO111MODULE=on

# Set working directory inner container
WORKDIR /app

# Copy all files from local project into working directory in cointainer
COPY . .

# Download application dependency
RUN go mod download

# Build application
RUN go build -o main .

# Execute application when container running
CMD ["./main"]

# Expose usage port of application
EXPOSE 8080
