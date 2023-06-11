package model

type AddCreditCardReq struct {
	Number  string `json:"credit_card_number"`
	Country string `json:"country"`
}

type CreditCard struct {
	ID      string `json:"id"`
	Number  string `json:"credit_card_number"`
	Country string `json:"country"`
}

type BinLookupResponse struct {
	Country struct {
		Name string `json:"name"`
	} `json:"country"`
}
