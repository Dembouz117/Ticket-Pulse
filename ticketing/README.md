# How to start

```bash
go mod download
```

```bash
cd ticketing/cmd
go run main.go
```

## Clear data

```bash
go run main.go cleardata
```

## Seed data

```bash
go run math.go seeddata
```

# Swagger

To generate swagger docs,

1. Navigate to parent directory `/ticketing`
1. run

```bash
swag init --parseInternal --parseDependency --dir cmd,internal --output docs
```

# At the start

```bash
go run -mod=mod entgo.io/ent/cmd/ent new [ent_type]
brew install protobuf
go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
export PATH="$PATH:$(go env GOPATH)/bin"
go get entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc@master

go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema
go generate ./ent
```

---

### Queue Endpoints

#### Join Queue

- Endpoint: /join
- Method: POST
- Description: Allows a user to join the waiting room or queue for an event.
- Request Body: NIL

- Response
  - Status 200 OK: Successfully joined the waiting room or queue.
  - Status 500 Internal Server Error: An error occurred while processing the request.

#### Queue

- Endpoint: /queue
- Method: GET
- Description: Retrieves a user's position in the event queue
- Request Body: NIL

- Response
  - Status 200 OK: Returns the user's position in the queue as a response message. If the API returns "Redirect to buy tickets", please redirect the user to the purchasing page.
  - Status 400 Bad Request: The user is not in the queue or an error occurred.
  - Status 500 Internal Server Error: An error occurred while processing the request.

##### WebSocket Support

The /queue endpoint supports WebSocket connections for real-time updates on the user's position in the queue. After making a GET request to /queue, you can establish a WebSocket connection to receive continuous updates. The WebSocket connection will provide updates on the user's position in the queue.
