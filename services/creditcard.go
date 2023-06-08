package services

import (
	"ccvs/data"
	"ccvs/model"
	"context"
)

type CreditCardData struct {
	Datastore data.ICreditCardData
}

func NewCreditCardData(data data.ICreditCardData) *CreditCardData {
	return &CreditCardData{Datastore: data}
}

func (cd *CreditCardData) AddCreditCard(ctx context.Context, addCreditCardReq model.AddCreditCardReq) error {
	return nil
}

func (cd *CreditCardData) GetCreditCard(ctx context.Context, creditCardID string) (model.CreditCard, error) {
	return model.CreditCard{}, nil
}

func (cd *CreditCardData) GetCreditCards(ctx context.Context) ([]model.CreditCard, error) {
	return []model.CreditCard{}, nil
}
