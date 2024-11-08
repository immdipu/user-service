# Use the Go image
FROM golang:1.23.3-alpine3.20

# Set the working directory inside the container
WORKDIR /app

RUN go install github.com/githubnemo/CompileDaemon@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the Go application and name the output binary 'user-service'
RUN go build -o user-service .


# Expose the port your service will run on
EXPOSE 8080


ENTRYPOINT ["CompileDaemon", "--build=go build -o user-service .", "--command=./user-service"]