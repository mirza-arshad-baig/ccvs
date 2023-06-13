package sql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
	"github.com/spf13/viper"
)

// GetDbInstance returns a singleton instance of the database connection.
// It reads the database configuration from the application settings and initializes the DB connection.
// If an error occurs during the initialization, it logs the error and returns the DB instance (possibly nil).
// It also performs a ping to ensure the connection is established.
func connect() (*sql.DB, error) {
	userName := viper.GetString("credit_card_db.username")
	password := viper.GetString("credit_card_db.password")
	host := viper.GetString("credit_card_db.host")
	port := viper.GetString("credit_card_db.port")
	databaseName := viper.GetString("credit_card_db.database_name")

	// Initialize DB
	dbCOnn, err := sql.Open("mysql", userName+":"+password+"@tcp("+host+":"+port+")/"+databaseName)
	if err != nil {
		return dbCOnn, err
	}

	// Open doesn't open a connection, so we need to ping it to make sure the connection is established
	err = dbCOnn.Ping()
	if err != nil {
		return nil, err
	}
	return dbCOnn, nil
}
