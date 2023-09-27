# Key-Value Store (KVS)
- Highly Scalable Distributed Key Value Store in Go

## Setup
- Create a `.env` file with `PORT="80"` inside it. Place the file in the root director of the project i.e `kvs/.env`
- In the root kvs directory execute `go build` to compile
- `./kvs` to execute
- If any imports are missings and not automatically download during build then do go get in the root of repo dir for each missing package. Ex: `go get github.com/go-chi/chi`

## Usage
Assuming the server is running on `localhost:80`. The API is at localhost:80/v1/kvs
- To get the Val for a Key do: HTTP GET http://localhost:80/v1/kvs with `Key` as a required Header. Returns a json with ```{"result":"Value", "ok":true}``` ok is false if Key is not present. `ok` is true if `Key` is present and `result` will give the value for the given key.
- To insert a new Key, Val pair do: HTTP PUT http://localhost:80/v1/kvs with `Key` as a required Header. Also add `Val` Header for a not default value. Returns a similar json with `ok` is true if insertion succeeds.
- To delete a Key do: HTTP DEL http://localhost:80/v1/kvs with `Key` as a required Header.  Returns a similar json with `ok` is true if the Key was present and deletion succeeds or the key was just absent. No error thrown if we delete a key that's not present (for Perf).
- You can use `thunderclient` VSCODE extension or postman for HTTP requests

## What's done?
1. Server Implementation
2. Concurrent Requests
4. Graceful Error Handling

## TODO
5. Logging with timestamps is being done only when some error occurs now. Add more logging for various ops like GET/PUT/DEL requests etc., as needed by `#5` in the Assignment description
3. Persistence for bonus points. That needs to be done.