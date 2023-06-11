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

	"github.com/spf13/viper"
)

type CreditCardData struct {
	Datastore data.ICreditCardData
}

func NewCreditCardData(data data.ICreditCardData) *CreditCardData {
	return &CreditCardData{Datastore: data}
}

func (cd *CreditCardData) AddCreditCard(ctx context.Context, addCreditCardReq model.AddCreditCardReq) error {
	if addCreditCardReq.Number == "" {
		return errors.New("credit number can't be blank")
	}

	// retrive country name by credit card number
	countryName, err := VerifyCountry(addCreditCardReq.Number)
	if err != nil {
		return err
	}

	// check if country not banned
	if viper.GetBool(countryName) {
		return errors.New("the card is issued in a list of banned countries")
	}

	fmt.Println("countryName:", countryName)

	addCreditCardReq.Country = countryName

	ccData, err := cd.Datastore.GetCreditCardByCCNumber(ctx, addCreditCardReq.Number)
	if err != nil {
		return err
	}
	if ccData.ID != "" {
		return errors.New("credit card already exist")
	}

	err = cd.Datastore.AddCreditCard(ctx, addCreditCardReq)
	if err != nil {
		return err
	}
	return nil
}

func (cd *CreditCardData) GetCreditCard(ctx context.Context, creditCardID string) (model.CreditCard, error) {
	if creditCardID == "" {
		return model.CreditCard{}, errors.New("credit card id can't be blank")
	}
	ccData, err := cd.Datastore.GetCreditCard(ctx, creditCardID)
	if err != nil {
		return model.CreditCard{}, err
	}

	return ccData, nil
}

func (cd *CreditCardData) GetCreditCards(ctx context.Context) ([]model.CreditCard, error) {
	ccDatas, err := cd.Datastore.GetCreditCards(ctx)
	if err != nil {
		return []model.CreditCard{}, err
	}

	return ccDatas, nil
}

func VerifyCountry(cardNumber string) (string, error) {
	bin, err := utils.ExtractBin(cardNumber)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://lookup.binlist.net/%s", bin)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Failed to fetch BIN information. Status code: %d", response.StatusCode)
	}

	var binResponse model.BinLookupResponse
	err = json.NewDecoder(response.Body).Decode(&binResponse)
	if err != nil {
		return "", err
	}

	return binResponse.Country.Name, nil
}
