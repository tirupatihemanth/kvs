# Key-Value Store (KVS)
- Highly Scalable Distributed Key Value Store in Go

### Docker Compose
1. To scale key value store across multiple nodes.
2. docker-compose up => builds kv store and nginx images and runs them.
2. docker-compose down => terminates all the running docker containers.

## To RUN
1. python client.py #Req #Clients

## Docker Setup
1. If you have access to code build the image first and run as described below. 
2. Alpine is much lighter but the base isn't very stable according to official golang docker hub. But our code works fine on both. By default let's use the heavier one as recommended by official documentation.
3. If you just want to run/use the application then take the dockerhub route

### Build
1. docker build --tag docker-kvs:v1.0 . 
2. docker build --file DockerfileAlp --tag docker-kvs-alp:v1.0 .

### Run
1. docker run -dp 65432:65432 docker-kvs:v1.0
2. docker run -dp 65432:65432 docker-kvs-alp:v1.

### DockerHub
1. docker pull hemanthkumart007/docker-kvs:v1.0
2. docker run -dp 65432:65432 hemanthkumart007/docker-kvs:v1.0
3. docker pull hemanthkumart007/docker-kvs-alp:v1.0
4. docker run -dp 65432:65432 hemanthkumart007/docker-kvs-alp:v1.0




## Setup
To set up the project please follow the below steps
1. Create a .env file with ’PORT=”65432”’ inside it. Place the file in the root directory of the project.
2. In the root kvs directory, execute ’go build’ to compile and ’./kvs’ to run the API.
3. If any imports are missing and not automatically downloaded during build then do ’go
get’ for each missing package.
4. You can use `thunderclient` VSCODE extension or postman for HTTP requests
5. API will log to a file "kvsLog.txt" which will be created.
6. API will persist to "persist.json" every 5 seconds.

## What's done?
1. Server Implementation
2. Concurrent Requests
3. Persistence
4. Graceful Error Handling
5. Logging
