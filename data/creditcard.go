package data

import (
	"ccvs/model"
	"context"
	"database/sql"
	"fmt"
)

type DB struct {
	CreditCardData *sql.DB
}

func (d DB) AddCreditCard(ctx context.Context, addCreditCardReq model.AddCreditCardReq) error {
	fmt.Println("AddCreditCard")
	return nil
}

func (d DB) GetCreditCard(ctx context.Context, creditCardID string) (model.CreditCard, error) {
	fmt.Println("GetCreditCard")
	return model.CreditCard{}, nil
}

func (d DB) GetCreditCards(ctx context.Context) ([]model.CreditCard, error) {
	fmt.Println("GetCreditCards")
	return []model.CreditCard{}, nil
}
