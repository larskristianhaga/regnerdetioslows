FROM golang:1.16-alpine

# Set the destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go ./

# Build
RUN go build -o /docker-gs-ping

# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 8080

# Run
CMD [ "/docker-gs-ping" ]
