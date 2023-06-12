package data

import (
	"ccvs/model"
	"context"
)

// ICreditCardData represents the interface for the credit card data layer.
type ICreditCardData interface {

	// AddCreditCard adds a credit card to the data store.
	AddCreditCard(ctx context.Context, addCreditCardReq model.AddCreditCardReq) error

	// GetCreditCard retrieves a credit card by its ID from the data store.
	GetCreditCard(ctx context.Context, creditCardID string) (model.CreditCard, error)

	// GetCreditCards retrieves all credit cards from the data store.
	GetCreditCards(ctx context.Context) ([]model.CreditCard, error)

	// GetCreditCardByCCNumber retrieves a credit card by its credit card number from the data store.
	GetCreditCardByCCNumber(ctx context.Context, ccNumber string) (model.CreditCard, error)
}
