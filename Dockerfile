FROM golang:1.21.1
WORKDIR /kvs
# Keep all dependencies Ready in this layer without your src code
COPY go.mod go.sum ./
RUN go mod download

# Get all socurce code in this layer.
COPY *.go client.py ./
RUN CGO_ENABLED=0 GOOS=linux go build -o kvs

# Command that will be executed on docker run
CMD ["./kvs"]