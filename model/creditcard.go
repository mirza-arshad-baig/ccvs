package model

type AddCreditCardReq struct {
	Name    string `json:"name"`
	Number  string `json:"credit_card_number"`
	Country string `json:"country"`
}

type CreditCard struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Number  string `json:"credit_card_number"`
	Country string `json:"country"`
}
