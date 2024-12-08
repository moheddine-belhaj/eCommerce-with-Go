# eCommerce with Go

This project is an e-commerce application built using the Go programming language, the Gin web framework, and MongoDB for data persistence.

## Features

- RESTful API for managing products, users, and orders.
- Built with **Gin**, a lightweight and fast web framework for Go.
- Secure user authentication using JWT.
- Data validation with `go-playground/validator`.
- Integration with MongoDB using the official MongoDB Go driver.
- Docker setup for development and testing.

## Prerequisites

- [Go](https://go.dev/doc/install) (version 1.22.3 or later)
- [Docker](https://www.docker.com/products/docker-desktop) (for running MongoDB and Mongo Express)

## Project Structure

```
eCommerce-with-Go/
│
├── controllers/          # API controllers
├── database              # Database Configuration
├── middleware            # Authentication function using the Gin framework
├── models/               # Database models
├── routes/               # API routes
├── tokens                # handle generating and validating the JWT tokens.
├── docker-compose.yml    # Docker setup for MongoDB and Mongo Express
├── go.mod                # Module dependencies
├── go.sum                # Dependency checksums
└── main.go               # Main for the application
```

## Setup

### Clone the Repository

```bash
git clone https://github.com/moheddine-belhaj/eCommerce-with-Go.git
cd eCommerce-with-Go
```

### Install Dependencies

Run the following command to download the Go module dependencies:

```bash
go mod tidy
```

### Run the Application

1. Start MongoDB and Mongo Express using Docker:
   ```bash
   docker-compose up -d
   ```
   - MongoDB will be available at `mongodb://localhost:27017`.
   - Mongo Express can be accessed via `http://localhost:8081`.

2. Run the Go application:
   ```bash
   go run main.go
   ```

### API Endpoints

Here are some key API endpoints (adjust as per your implementation):

- **User Authentication**
  - `POST /users/login` - Login a user.
  - `POST /users/signup` - Register a new user.

- **Product Management**
  - `GET /users/productview` - Get a list of products.
  - `POST /admin/products` - Create a new product.

- **Order Management**
  - `POST /addtocart?id=xxxproduct_idxxx&userID=xxxxxxuser_idxxxxxx` - Create an order.

## Contact

For any inquiries or issues, feel free to reach out to [moheddine-belhaj](https://github.com/moheddine-belhaj).

