# Key-Value Store (KVS)
- Highly Scalable Distributed Key Value Store in Go


## Setup
To set up the project please follow the below steps
1. Create a .env file with ’PORT=”65432”’ inside it. Place the file in the root directory
of the project.
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
