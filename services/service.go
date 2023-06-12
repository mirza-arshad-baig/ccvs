package services

import (
	"ccvs/model"
	"context"
)

type ICreditCard interface {
	// AddCreditCard adds a credit card to the database.
	AddCreditCard(ctx context.Context, addCreditCardReq model.AddCreditCardReq) error

	// GetCreditCard retrieves a credit card by its ID.
	GetCreditCard(ctx context.Context, creditCardID string) (model.CreditCard, error)

	// GetCreditCards retrieves all credit cards from the database.
	GetCreditCards(ctx context.Context) ([]model.CreditCard, error)
}
