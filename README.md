# gRPC CRUD API with Go and MongoDB

A CRUD API built with **Go, gRPC, Protocol Buffers, and MongoDB**. This project demonstrates how to build a gRPC server that handles user management operations and stores user records in MongoDB.

## Features

- Create user records
- Retrieve users by ID
- Update existing user details
- Delete user records
- Communicate with clients using gRPC
- Define service methods and messages using Protocol Buffers
- Store and manage data with MongoDB

## Tech Stack

- **Language:** Go
- **API:** gRPC
- **Serialization:** Protocol Buffers
- **Database:** MongoDB
- **Database Driver:** MongoDB Go Driver

## Project Structure

```text
grpc_mongo_crud/
├── db/          # MongoDB connection and configuration
├── models/      # User data model
├── proto/       # Protocol Buffer definitions and generated code
├── server/      # gRPC service implementations
├── main.go      # Application entry point
├── go.mod
└── README.md
```

## Prerequisites

- [Go](https://go.dev/dl/)
- [MongoDB Community Server](https://www.mongodb.com/try/download/community)
- Protocol Buffers and gRPC tools if regenerating code from `.proto` files

## Getting Started

### 1. Start MongoDB

Ensure MongoDB is running locally. The application is configured to connect to:

```text
mongodb://localhost:27017
```

The database is named `grpc`, and the collection is named `users`.

### 2. Clone the Repository

```bash
git clone <your-repository-url>
cd grpc_mongo_crud
```

Replace `<your-repository-url>` with your GitHub repository URL.

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Run the Server

```bash
go run .
```

The gRPC server listens on port `50051`.

## CRUD Operations

| Operation | Description |
|---|---|
| CreateUser | Creates a new user record |
| GetUser | Retrieves a user by ID |
| UpdateUser | Updates an existing user |
| DeleteUser | Deletes a user by ID |

You can test the gRPC service using a compatible client, such as [Postman](https://www.postman.com/) or `grpcurl`.

## What I Learned

- Building gRPC services with Go
- Defining service contracts using Protocol Buffers
- Connecting Go applications to MongoDB
- Implementing database CRUD operations
- Organizing code into separate packages
- Handling errors in API and database operations

## Future Improvements

- Add input validation and appropriate gRPC status codes
- Add automated tests for CRUD operations
- Use environment variables for database configuration
- Implement authentication and authorization
- Add structured logging

## License

A license can be added if the project is intended for public reuse.
