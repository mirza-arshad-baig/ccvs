package data

import (
	"ccvs/model"
	"context"
	"database/sql"
)

// DB represents the implementation of the ICreditCardData interface using a SQL database.
type DB struct {
	CreditCardData *sql.DB
}

// AddCreditCard adds a credit card to the database.
func (d DB) AddCreditCard(ctx context.Context, addCreditCardReq model.AddCreditCardReq) error {
	sqlstr := "INSERT INTO credit_cards SET card_number = ?, country = ?"

	stmt, err := d.CreditCardData.Prepare(sqlstr)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		addCreditCardReq.Number,
		addCreditCardReq.Country,
	)
	defer stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

// GetCreditCard retrieves a credit card by its ID from the database.
func (d DB) GetCreditCard(ctx context.Context, creditCardID string) (model.CreditCard, error) {
	var ccData model.CreditCard

	sqlStr := `SELECT id, card_number, country FROM credit_cards WHERE id = ?`
	err := d.CreditCardData.QueryRow(sqlStr, creditCardID).Scan(&ccData.ID, &ccData.Number, &ccData.Country)
	if sql.ErrNoRows == err {
		// No credit card found with the given ID
		return model.CreditCard{}, nil
	} else if err != nil {
		return model.CreditCard{}, err
	}
	return ccData, nil
}

// GetCreditCards retrieves all credit cards from the database.
func (d DB) GetCreditCards(ctx context.Context) ([]model.CreditCard, error) {
	var ccDatas []model.CreditCard

	sqlStr := `SELECT id, card_number, country FROM credit_cards`
	rows, err := d.CreditCardData.Query(sqlStr)
	if sql.ErrNoRows == err {
		// No credit cards found
		return []model.CreditCard{}, nil
	} else if err != nil {
		return []model.CreditCard{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var ccData model.CreditCard
		err := rows.Scan(
			&ccData.ID,
			&ccData.Number,
			&ccData.Country,
		)
		if err != nil {
			return nil, err
		}
		ccDatas = append(ccDatas, ccData)
	}
	return ccDatas, nil
}

// GetCreditCardByCCNumber retrieves a credit card by its credit card number from the database.
func (d DB) GetCreditCardByCCNumber(ctx context.Context, ccNumber string) (ccData model.CreditCard, err error) {

	sqlStr := `SELECT id, card_number FROM credit_cards WHERE card_number = ?`
	err = d.CreditCardData.QueryRow(sqlStr, ccNumber).Scan(&ccData.ID, &ccData.Number)
	if sql.ErrNoRows == err {
		// No credit card found with the given credit card number
		return model.CreditCard{}, nil
	} else if err != nil {
		return model.CreditCard{}, err
	}

	return ccData, nil
}
