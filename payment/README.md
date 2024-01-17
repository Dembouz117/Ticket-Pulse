# Payment

`sudo npm install -g pnpm`

`pnpm install`

`pnpm run dev`

env file is required

For payment microservice, there are prerequisites:
(a) Ensure Stripe Webhook server is running:
- Install Stripe CLI via this link: https://github.com/stripe/stripe-cli
- if testing on local run : 
```bash
stripe login
stripe listen --forward-to localhost:8082/payment/webhook
```
(b) Ensure that the ticketing server is running:
```bash
cd ticketing/cmd
go run . #if on mac
```
(c) Ensure that gRPC server on ticketing is running:
```bash
cd ticketing/server
go run grpc_server.go
```
(d) Ensure auth is running:
```bash
cd auth
go run .
```
(e) Ensure the frontend has a Stripe publishable key. Check for a .env.local inside the "frontend" folder
- Make sure it contains a NEXT_PUBLIC_PUBLISHABLE_KEY
- Make sure dockerfile contains the publisahble key as well