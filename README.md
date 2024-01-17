# Introduction

TicketPulse enhances the ticketing experience by tackling common issues found in existing platforms: unclear queue positions, bot interference, and system lags. We've ensured each user can have only one account to prevent queue bloating and introduced features showing clear queue positioning which are updated at regular intervals. To deter bots and fraudulent activities, two-factor authentication and machine learning safeguards have been put in place. By understanding the pain points of users, we've created a system that prioritises fairness and transparency for ticket purchasers.

# System Architecture

Our system is designed for resilience, scalability and security. We are utilising microservices for our software development. The services are as follow:

1. Frontend
1. Ticketing Backend
1. Authentication Backend
1. Payment Backend
1. Machine Learning Backend
1. Queue Backend (WIP)

# How to start

These are the current port allocations:

| Port | Service        |
| ---- | -------------- |
| 3000 | Frontend       |
| 8080 | Authentication |
| 8081 | Ticketing      |
| 8082 | Payment        |
| 5003 | Ticketing gRPC |

## Docker

```bash
docker compose up
# if want logs
docker compose logs auth -f
```

If there are issues with Go related, eg `=> ERROR [auth builder 6/6] RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app . `, run `go mod tidy` in that folder.

## Installation

1. Ensure that you have Go installed. Refer to this [link](https://go.dev/doc/install). For each Go microservice, run the following command first:
    ```bash
    go mod download
    ```
1. Ensure that you have npm and pnpm installed. Refer to this [link](https://pnpm.io/installation). For Frontend and Payment, run the following command first:
    ```bash
    npm install
    # or
    yarn install
    # or
    pnpm install
    ```
1. Ensure that you have the stripe CLI installed. Refer to this [link](<[url](https://github.com/stripe/stripe-cli)>):
1. Before running Payment Microservice
    Run:
    ```bash
    npx prisma generate
    ```

## Run

Run the following services in this order from the root directory each time.

### Authentication microservice

```bash
cd auth
go run auth.go
```

### Ticketing microservice

```bash
cd ticketing/cmd
go run main.go
```

run the gRPC server:

```bash
cd ticketing/server
go run grpc_server.go
```

### Payment microservice

For payment microservice, there are prerequisites.

1. Stripe Webhook server is running:
    -   Install Stripe CLI via this [link](https://github.com/stripe/stripe-cli)
    -   Run :
    ```bash
    stripe login
    stripe listen --forward-to localhost:8082/webhook
    ```

1. Ensure the frontend has a Stripe publishable key. Check for a .env.local inside the "frontend" folder
    -   Make sure it contains a NEXT_PUBLIC_PUBLISHABLE_KEY
    -   Run:
    ```bash
    pnpm run dev
    ```

#### During development

If adjustments made to protobuff, run:

```bash
pnpm run proto:gen
```

This compiles the protobuff and exports useable Typescript types into the ./proto folder

If issues popup (during development or otherwise), run: `deallocate all;` in the psql server

### Frontend

Ensure that you have a .env.local in the frontend folder, and that it contains NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY

```bash
pnpm run dev
```

## Initial Set up postgresql and ent schemas

This applies also when there is a change to the schema

1. First, install the ent related (skip if have done before)

    ```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    export PATH="$PATH:$(go env GOPATH)/bin"
    go install entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc@master
    ```

1. Update the ent schema (ensure that the path is correct)

    ```bash
    go run entgo.io/ent/cmd/ent describe ./ent/schema
    ```

1. After all the schema has been changed, go to cmd folder (or where there is an automatic migration)

    ```bash
    go test
    ```

### If no PostgreSQL online

```bash
CREATE ROLE ticket_admin WITH LOGIN PASSWORD ‘aDmiNTickET’;
ALTER ROLE ticket_admin CREATEDB;

psql postgres -U ticket_admin
CREATE DATABASE ticketing WITH OWNER=ticket_admin;

DATABASE_URL=postgresql://ticket_admin:aDmiNTickET@localhost:5432/ticketing
```

# Dependencies

Go version: 1.21.0
Node version: ?
