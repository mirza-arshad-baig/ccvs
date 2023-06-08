package data

import (
	"ccvs/model"
	"context"
)

type SqlDB struct{}

func (d *SqlDB) AddCreditCard(ctx context.Context, addCreditCardReq model.AddCreditCardReq) error {

	return nil
}

func (d *SqlDB) GetCreditCard(ctx context.Context, creditCardID string) (model.CreditCard, error) {

	return model.CreditCard{}, nil
}

func (d *SqlDB) GetCreditCards(ctx context.Context) ([]model.CreditCard, error) {

	return []model.CreditCard{}, nil
}
