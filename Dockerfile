#We used the defacto go image as our base image. The docker hub page for Go lang 
# recommends (https://hub.docker.com/_/golang)using the defacto image instead of the alpine #image due to unexpected behavior observed in the alpine image for Go. However, we have the #docker file that uses alpine too present in the git repository.
FROM golang:1.21.1

# Set the working directory inside the container
WORKDIR /kvs

# Install the dependencies 
COPY go.mod go.sum ./
RUN go mod download

# Add source files to docker image and build the API
COPY *.go client.py ./
RUN CGO_ENABLED=0 GOOS=linux go build -o kvs

#Set env variable PORT to be used by API code and expose the PORT 65432
ENV PORT 8080
EXPOSE 8080

# Entry point command to execute to run the API after launching the container
CMD ["./kvs"]