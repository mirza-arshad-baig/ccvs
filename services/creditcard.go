package services

import (
	"ccvs/data"
	"ccvs/model"
	"context"
	"errors"
)

type CreditCardData struct {
	Datastore data.ICreditCardData
}

func NewCreditCardData(data data.ICreditCardData) *CreditCardData {
	return &CreditCardData{Datastore: data}
}

func (cd *CreditCardData) AddCreditCard(ctx context.Context, addCreditCardReq model.AddCreditCardReq) error {
	if addCreditCardReq.Name == "" || addCreditCardReq.Number == "" || addCreditCardReq.Country == "" {
		return errors.New("credit card holder name / number / country can't be blank")
	}
	err := cd.Datastore.AddCreditCard(ctx, addCreditCardReq)
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
