package model

type AddCreditCardReq struct {
	Name      string
	Number    string
	Country   string
	ValidUpto string
}

type CreditCard struct {
	ID        string
	Name      string
	Number    string
	Country   string
	ValidUpto string
}
