package services

import (
	"ccvs/common/utils"
	"ccvs/data"
	"ccvs/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// CreditCardData represents the data layer for credit card operations.
type CreditCardData struct {
	Datastore data.ICreditCardData
}

// NewCreditCardData creates a new instance of CreditCardData with the given ICreditCardData implementation.
func NewCreditCardData(data data.ICreditCardData) *CreditCardData {
	return &CreditCardData{Datastore: data}
}

// AddCreditCard adds a credit card to the datastore.
func (cd *CreditCardData) AddCreditCard(ctx context.Context, addCreditCardReq model.AddCreditCardReq) error {

	// Check if credit card number is empty
	if addCreditCardReq.Number == "" {
		return errors.New("credit card number can't be blank")
	}

	// Retrieve country name by credit card number
	countryName, err := VerifyCountry(addCreditCardReq.Number)
	if err != nil {
		return err
	}

	// Check if the country is not banned
	if viper.GetBool(countryName) {
		return errors.New("the card is issued in a list of banned countries")
	}

	addCreditCardReq.Country = countryName

	ccData, err := cd.Datastore.GetCreditCardByCCNumber(ctx, addCreditCardReq.Number)
	if err != nil {
		return err
	}
	if ccData.ID != "" {
		err = errors.New("credit card already exist")
		logrus.Printf("Error adding credit card: %s", err.Error())
		return err
	}

	// Retrieve credit card data from the datastore
	err = cd.Datastore.AddCreditCard(ctx, addCreditCardReq)
	if err != nil {
		logrus.Errorf("Error adding credit card data: %s", err.Error())
		return err
	}
	return nil
}

// GetCreditCard retrieves a credit card by its ID from the datastore.
func (cd *CreditCardData) GetCreditCard(ctx context.Context, creditCardID string) (model.CreditCard, error) {

	// Check if credit card ID is empty
	if creditCardID == "" {
		return model.CreditCard{}, errors.New("credit card ID can't be blank")
	}
	ccData, err := cd.Datastore.GetCreditCard(ctx, creditCardID)
	if err != nil {
		logrus.Errorf("Error retrieving credit card data: %s", err.Error())
		return model.CreditCard{}, err
	}

	return ccData, nil
}

// GetCreditCards retrieves all credit cards from the datastore.
func (cd *CreditCardData) GetCreditCards(ctx context.Context) ([]model.CreditCard, error) {
	// Retrieve credit card data from the datastore
	ccDatas, err := cd.Datastore.GetCreditCards(ctx)
	if err != nil {
		logrus.Errorf("Error retrieving credit cards: %s", err.Error())
		return []model.CreditCard{}, err
	}

	return ccDatas, nil
}

// VerifyCountry retrieves the country name associated with the provided card number.
func VerifyCountry(cardNumber string) (string, error) {

	// Extract the BIN (Bank Identification Number) from the card number
	bin, err := utils.ExtractBin(cardNumber)
	if err != nil {
		return "", err
	}

	// Create the URL for BIN lookup
	url := fmt.Sprintf(utils.BinUrl+"%s", bin)

	// Send HTTP GET request to retrieve BIN information
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Failed to fetch BIN information. Status code: %d", response.StatusCode)
	}

	// Decode the JSON response into a BinLookupResponse struct
	var binResponse model.BinLookupResponse
	err = json.NewDecoder(response.Body).Decode(&binResponse)
	if err != nil {
		return "", err
	}

	return binResponse.Country.Name, nil
}
