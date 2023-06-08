package services

import (
	"ccvs/model"
	"context"
)

type ICreditCard interface {
	AddCreditCard(ctx context.Context, addCreditCardReq model.AddCreditCardReq) error
	GetCreditCard(ctx context.Context, creditCardID string) (model.CreditCard, error)
	GetCreditCards(ctx context.Context) ([]model.CreditCard, error)
}
