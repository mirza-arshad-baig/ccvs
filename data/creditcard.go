package data

import (
	"ccvs/model"
	"context"
	"database/sql"
)

type DB struct {
	CreditCardData *sql.DB
}

func (d DB) AddCreditCard(ctx context.Context, addCreditCardReq model.AddCreditCardReq) error {
	sqlstr := "INSERT INTO credit_cards SET name = ? , card_number = ?, country = ?"

	stmt, err := d.CreditCardData.Prepare(sqlstr)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		addCreditCardReq.Name,
		addCreditCardReq.Number,
		addCreditCardReq.Country,
	)
	defer stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (d DB) GetCreditCard(ctx context.Context, creditCardID string) (model.CreditCard, error) {
	var ccData model.CreditCard

	sqlStr := `SELECT id, card_number, name, country FROM credit_cards WHERE id = ?`
	err := d.CreditCardData.QueryRow(sqlStr, creditCardID).Scan(&ccData.ID, &ccData.Number, &ccData.Name, &ccData.Country)
	if sql.ErrNoRows == err {
		return model.CreditCard{}, nil
	} else if err != nil {
		return model.CreditCard{}, err
	}
	return ccData, nil
}

func (d DB) GetCreditCards(ctx context.Context) ([]model.CreditCard, error) {
	var ccDatas []model.CreditCard
	sqlStr := `SELECT id, card_number, name, country FROM credit_cards`
	rows, err := d.CreditCardData.Query(sqlStr) //(&ccData)
	if sql.ErrNoRows == err {
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
			&ccData.Name,
			&ccData.Country,
		)
		if err != nil {
			return nil, err
		}
		ccDatas = append(ccDatas, ccData)
	}
	return ccDatas, nil
}

func (d DB) GetCreditCardByCCNumber(ctx context.Context, ccNumber string) (ccData model.CreditCard, err error) {
	sqlStr := `SELECT id, card_number, name, country FROM credit_cards WHERE card_number = ?`
	err = d.CreditCardData.QueryRow(sqlStr, ccNumber).Scan(&ccData.ID, &ccData.Number, &ccData.Name, &ccData.Country)
	if sql.ErrNoRows == err {
		return model.CreditCard{}, nil
	} else if err != nil {
		return model.CreditCard{}, err
	}
	return ccData, nil
}
