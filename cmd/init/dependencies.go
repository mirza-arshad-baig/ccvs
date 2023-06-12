package init

import (
	"ccvs/common/libs/sql"
	"ccvs/data"
)

var (
	datastore    data.ICreditCardData
	creditCardDB data.DB
)

// InitializeDependencies initializes the dependencies for the application.
// It reads the DB configuration from the app config file (app.json),
// gets the DB instance using the GetDbInstance function from the sql package,
// and sets up the creditCardDB and datastore variables.
func InitializeDependencies() {

	// Read DB config from app config file (app.json)
	sql := sql.GetDbInstance()

	// Set up the creditCardDB with the DB instance
	creditCardDB.CreditCardData = sql

	// Set the datastore to creditCardDB
	datastore = creditCardDB
}
