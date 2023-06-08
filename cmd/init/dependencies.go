package init

import (
	"ccvs/data"
)

var (
	datastore data.ICreditCardData
)

func InitializeDependencies() {
	//Initialized DB
	/* var err error
	ctx := context.Background()
	datastore, err = mongo.NewMongoDataStore(ctx, mongoDbSession)
	if err != nil {
		log.Fatalf("[InitializeDependencies] Failed to initialize datastore %s", err.Error())
	} */
}

func Close() {
	// close sql db connection
}
