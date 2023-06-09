## Credit Card Validation System (CCVS)


Credit Card Validation System (CCVS) is a Go application that allows administrators to submit credit card numbers for validation. The system checks the country of issuance and verifies if it is in a list of banned countries. Valid credit cards are stored securely, and duplicate submissions are prevented.

### Getting Started

1. Clone the repository:

```shell
git clone https://github.com/your-username/credit-card-system.git
cd credit-card-system
```

### Features

- Submit a credit card number for validation.
- Check the country of issuance and compare it against a list of banned countries.
- Store valid credit cards securely.
- Prevent duplicate submissions of the same credit card.
- Retrieve all captured credit cards.
- Retrieve a single credit card by its ID.

### Prerequisites

- Go programming language (version 1.16+)
- PostgreSQL database

### Getting Started

1. Clone the repository:

```shell
git clone https://github.com/your-username/credit-card-system.git
cd credit-card-system
Set up the PostgreSQL database:
Create a new database named "creditcards".
Update the database connection configuration in the main.go file with your PostgreSQL credentials.
```

Set up the PostgreSQL database:
Create a new database named "creditcards".
Update the database connection configuration in the main.go file with your PostgreSQL credentials.
Install the project dependencies:

```shell
Copy code
go mod download
Run the application:
```

```shell
Copy code
go run main.go
```

The application will start running on http://localhost:8080.


### API Endpoints
```
POST /creditcards: Submit a credit card for validation.
GET /creditcards: Retrieve all captured credit cards.
GET /creditcards/{id}: Retrieve a single credit card by ID.
```

### Configuration
The application supports configuration through environment variables. The following variables can be set:

```
DB_HOST: PostgreSQL database host (default: localhost).
DB_PORT: PostgreSQL database port (default: 5432).
DB_USER: PostgreSQL database user (default: postgres).
DB_PASSWORD: PostgreSQL database password.
DB_NAME: PostgreSQL database name (default: creditcards).
```

### Testing
Unit tests are provided for the application's handlers. To run the tests, use the following command:

```shell
Copy code
go test ./...
```

### Postman Collection
A Postman Collection is provided in the postman directory. It contains sample requests that can be used to test the API endpoints.

### OpenAPI Specification
The OpenAPI Specification (OAS) YAML file openapi.yml is included in the project root. It describes the API endpoints, request/response schemas, and other relevant details.
#   c c v s  
 