package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	CreditCardDB *sql.DB
	err          error
)

func GetDbInstance() *sql.DB {
	userName := viper.GetString("credit_card_db.username")
	password := viper.GetString("credit_card_db.password")
	host := viper.GetString("credit_card_db.host")
	port := viper.GetString("credit_card_db.port")
	databaseName := viper.GetString("credit_card_db.database_name")

	//Initialized DB
	CreditCardDB, err = sql.Open("mysql", userName+":"+password+"@tcp("+host+":"+port+")/"+databaseName)
	if err != nil {
		logrus.Fatal(err)
		return CreditCardDB
	}

	// Open doesn't open a connection, so we need to ping it to make sure the connection is established
	err = CreditCardDB.Ping()
	if err != nil {
		fmt.Println("=====================", userName, " = ", password, " = ", host, " = ", port, " = ", databaseName)
		logrus.Fatal(err)
	}
	return CreditCardDB
}
