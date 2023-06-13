package init

import (
	"ccvs/data"

	dataSQL "ccvs/data/sql"

	"github.com/sirupsen/logrus"
)

var (
	datastore data.ICreditCardData
)

// InitializeDependencies initializes the dependencies for the application.
// It reads the DB configuration from the app config file (app.json),
// gets the DB instance using the GetDbInstance function from the sql package,
// and sets up the creditCardDB and datastore variables.
func InitializeDependencies() {
	creditCard, err := dataSQL.NewSQLCreditCard()
	if err != nil {
		logrus.Fatalf("fatal: ", err)
	}

	// Set the datastore to creditCardDB
	datastore = creditCard
}
