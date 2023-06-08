package init

import (
	"ccvs/common/libs/db"
	"ccvs/data"
)

var (
	datastore data.ICreditCardData
)

func InitializeDependencies() {
	// Read DB config from app config file which is app.json
	_ = db.GetDbInstance()
}
