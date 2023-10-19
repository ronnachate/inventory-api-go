# Base image
From golang:1.21.3

# Create app directory
WORKDIR /go/src/app

# Bundle app source
COPY . .

# Build the Go app binary
RUN go build -o main app/main.go

# Start the server using the binary
CMD ["./main"]