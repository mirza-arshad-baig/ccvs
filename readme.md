## Credit Card Validation System (CCVS)


Credit Card Validation System (CCVS) is a Go application that allows administrators to submit credit card numbers for validation. The system checks the country of issuance and verifies if it is in a list of banned countries. Valid credit cards are stored securely, and duplicate submissions are prevented.
### Features

- Submit a credit card number for validation.
- Check the country of issuance and compare it against a list of banned countries.
- Store valid credit cards securely.
- Prevent duplicate submissions of the same credit card.
- Retrieve all captured credit cards.
- Retrieve a single credit card by its ID.

### Prerequisites

- Go programming language (version 1.16+)
- SQL database ([click here](https://github.com/mirza-arshad-baig/ccvs/blob/develop/scripts/sql-db-setup/db-setup.sql) for setup)

### Getting Started
1. Clone the repository:

```shell
https://github.com/mirza-arshad-baig/ccvs.git
cd ccvs
Update the database connection configuration in the ./config/app.json file with your SQL credentials.

credit_card_db.username : database user (default: root).
credit_card_db.password : database password.
credit_card_db.host : database host (default: localhost).
credit_card_db.port : database port (default: 3306).
credit_card_db.database_name : database name (default: credit_card).
```

### Configure Banned Countries List
[click here](https://github.com/mirza-arshad-baig/ccvs/blob/develop/config/banned_countries.json) for see format of banned countries list.
We can any countryName in the ./config/banned_countries.json with the value true 
example:
```
{
    "countryName" : true
}
```
open terminal and run below command
```shell
go run ./cmd/main.go
```
The application will start running on http://localhost:8080.

### API Endpoints
```
POST /creditcards: Submit a credit card for validation.
GET /creditcards: Retrieve all captured credit cards.
GET /creditcards/{id}: Retrieve a single credit card by ID.
```

please refer [postman collection](https://github.com/mirza-arshad-baig/ccvs/blob/develop/scripts/postman/credit-card-validation-system.postman_collection.json) for complete request



### Testing
Unit tests are provided for the application's handlers. To run the tests, use the following command:

```shell
cd ./test/
go test -v -run <test case name>
```