# Authentication Microservice

The Authentication Microservice is a component of your application responsible for user authentication and user management. It provides endpoints for user login, OTP (One-Time Password) verification, user creation, user information retrieval, and user deletion. This README.md will guide you through the setup and usage of this microservice.

## Prerequisites

Before using the Authentication Microservice, make sure you have the following prerequisites installed and configured:

- Go (Golang) - [Installation Guide](https://golang.org/doc/install)
- PostgreSQL Database
- Redis Server
- A secret key for JWT (JSON Web Token) generation

## Environment Variables

To configure the Authentication Microservice, set the following environment variables in the .env file:

- REDIS_ADDRESS: The address of the Redis server.
- POSTGRES_ADDRESS: The address of the PostgreSQL database.
- SECRET_KEY: The secret key used for JWT token generation.

## Running the Microservice

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd auth
   ```

2. Install the required Go dependencies:

    ```bash
    go mod tidy
    ```

3. Create a .env file in the project directory and set the required environment variables (see Environment Variables).

4. Build and run the microservice:

    ```bash
    go run auth.go
    ```

The microservice should now be running on http://localhost:8080.

## Development

To run ent to generate a new database schema, run the following commands:

```bash
export GOWORK=off`
go generate -x ./ent
```

## Usage

You can interact with the Authentication Microservice using HTTP requests. Below are the available endpoints and their descriptions.

### Endpoints

#### Login

- Endpoint: /login
- Method: POST
- Description: Logs in a user and sends an OTP to their email.
- Request Body:

```json
{
    "email": "user@example.com",
    "password": "user_password"
}
```

- Response
  - Status 200 OK: OTP sent successfully.
  - Status 400 Bad Request: Invalid JSON request or credentials.
  - Status 401 Unauthorized: Invalid credentials.

#### OTP Verification

- Endpoint: /otp
- Method: POST
- Description: Verifies the OTP sent to the user and returns a JWT token upon successful verification.
- Request Body:

```json
{
  "email": "user@example.com",
  "otpCode": "123456"
}
```

- Response
  - Status 200 OK: JWT token returned.
  - Status 400 Bad Request: Invalid JSON request or OTP.
  - Status 401 Unauthorized: Invalid OTP or expired OTP.

### User Creation

- Endpoint: /user
- Method: POST
- Description: Creates a new user account.
- Request Body:

```json
{
  "name": "John Doe",
  "email": "newuser@example.com",
  "password": "new_user_password",
  "phone": "1234567890"
}
```

- Response
  - Status 200 OK: User created successfully.
  - Status 400 Bad Request: Invalid JSON request or invalid email, password, or phone format.
  - Status 500 Internal Server Error: Error creating user (e.g., email already exists).

### Get User Information

- Endpoint: /user
- Method: GET
- Description: Retrieves user information based on the JWT token.
- Request Headers:
  - Authorization: Bearer JWT_TOKEN
- Response:
  - Status 200 OK: User information returned.
  - Status 401 Unauthorized: Missing or invalid JWT token.
  - Status 500 Internal Server Error: Error fetching user information.

### Delete User Account

- Endpoint: /user
- Method: DELETE
- Description: Deletes the user account based on the JWT token.
- Request Headers:
  - Authorization: Bearer JWT_TOKEN
- Response:
  - Status 200 OK: User account deleted.
  - Status 401 Unauthorized: Missing or invalid JWT token.
  - Status 500 Internal Server Error: Error deleting user account.
 
### JWT

Contains
- `userId`: ID of the user
- `role`: Role of the user, could be either admin or user
- `exp`: Expires in 24 hours
