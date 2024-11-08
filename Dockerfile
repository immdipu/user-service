# Use the Go image
FROM golang:1.23.3-alpine3.20

# Set the working directory inside the container
WORKDIR /app


COPY . .

# Build the Go application and name the output binary 'user-service'
RUN go build -o user-service .

# Expose the port your service will run on
EXPOSE 8000


ENTRYPOINT ["./user-service"]
