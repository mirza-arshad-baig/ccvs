package test

import (
	"ccvs/common/libs"
	"ccvs/common/libs/sql"
	"ccvs/data"
	"ccvs/model"
	"ccvs/services"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	// ===========================================================================
	// Load Environment Config
	// ===========================================================================
	libs.ConfigFile = "app"

	// ConfigPaths is config file paths
	libs.InitConfig("../config")

	// ===========================================================================
	// Load Banned Countries List Config
	// ===========================================================================
	libs.ConfigFile = "banned_countries"
	libs.InitConfig("../config")
}

var creditCardDB data.DB

func TestAddCreditCard(t *testing.T) {
	ctx := context.Background()
	var err error
	creditCardDB.CreditCardData = sql.GetDbInstance()
	service := services.NewCreditCardData(creditCardDB)

	addCreditCardReq := []model.AddCreditCardReq{
		{
			// An error will occur because the above card number is issued in India, and India is among the banned countries
			Number: "4591570098754163",
		},
		{
			// Success
			Number: "4591150098765161",
		},
	}
	for _, testCase := range addCreditCardReq {
		t.Run(testCase.Number, func(t *testing.T) {
			err = service.AddCreditCard(ctx, testCase)
			if err != nil {
				if testCase.Number == "4591570098754163" {
					assert.Equal(t, "the card is issued in a list of banned countries", err.Error(), "Result should be equal")
				} else {
					require.NoError(t, err, "AddCreditCard should not return an error")
				}
			}
		})
	}
}

func TestGetCreditCard(t *testing.T) {
	ctx := context.Background()
	var err error
	creditCardDB.CreditCardData = sql.GetDbInstance()
	service := services.NewCreditCardData(creditCardDB)

	data, err := service.GetCreditCard(ctx, "1")
	if err != nil {
		require.NoError(t, err, "GetCreditCard should not return an error")
	}

	assert.Equal(t, data, model.CreditCard{
		Number: "1234567891234567",
		ID:     "1",
	}, "Result should be equal")
}

func TestGetallCreditCard(t *testing.T) {
	ctx := context.Background()
	var err error
	creditCardDB.CreditCardData = sql.GetDbInstance()
	service := services.NewCreditCardData(creditCardDB)

	data, err := service.GetCreditCards(ctx)
	if err != nil {
		require.NoError(t, err, "GetCreditCards should not return an error")
	}

	assert.Equal(t, data, []model.CreditCard{
		{
			Number: "1234567891234567",
			ID:     "1",
		},
		{
			Number: "4591570098754163",
			ID:     "2",
		},
	}, "Result should be equal")
}
