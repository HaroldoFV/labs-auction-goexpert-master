# Auction Project

This project is a Go application for managing auctions. It includes functionalities for creating auctions, closing expired auctions, and more.

## Prerequisites

- Go 1.20 or later
- MongoDB
- Docker (optional, for running MongoDB in a container)

## Setup

1. **Clone the repository:**

   ```sh
   git clone https://github.com/HaroldoFV/labs-auction-goexpert-master.git
   cd labs-auction-goexpert-master
   ```

3. **Install dependencies:**

   ```sh
   go mod tidy
   ```

## Running the Application

1. **Construa e execute a aplicação usando Docker Compose**:

    ```bash
    docker-compose up --build
    ```

2. A aplicação estará disponível em `http://localhost:8080`.




## Running Tests

To run the tests, use the following command:

```sh
go test ./...
```





## Creating an Auction

To create an auction, send a POST request to `http://localhost:8080/auction` with the following JSON payload:

```json
{
    "product_name": "Mouse",
    "category": "Technology",
    "description": "Mouse Logitech ...",
    "condition": 0
}
```

You can use `curl` to send the request:

```sh
curl -X POST http://localhost:8080/auction -H "Content-Type: application/json" -d '{
    "product_name": "Mouse",
    "category": "Technology",
    "description": "Mouse Logitech ...",
    "condition": 0
}'
```

Or use a tool like Postman to send the request.