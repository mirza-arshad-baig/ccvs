package init

import (
	"ccvs/common/libs/sql"
	"ccvs/data"
)

var (
	datastore    data.ICreditCardData
	creditCardDB data.DB
)

func InitializeDependencies() {
	// Read DB config from app config file which is app.json
	sql := sql.GetDbInstance()
	creditCardDB.CreditCardData = sql
	datastore = creditCardDB
}
