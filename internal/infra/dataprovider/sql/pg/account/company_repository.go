package entity

import (
	"context"
	"database/sql"
)

type AccountPGSQLRepository struct {
	ctx context.Context
	db *sql.DB
}

func NewAccountPGSQLRepository(
	ctx context.Context,
	db *sql.DB,
) *AccountPGSQLRepository {
	return &AccountPGSQLRepository{
		ctx: ctx,
		db: db,
	}
}

func (r *AccountPGSQLRepository) GetAccount(accountId int) (*Account, error) {
	var output Account

	return &output, nil
}

func (r *AccountPGSQLRepository) LoginAsCustomer(email, password string) (*Account, error) {
	var output Account
	stmt := `SELECT * FROM account a LEFT JOIN person p ON a.person_id =p.id LEFT JOIN customer c ON p.customer_id =c.id LEFT JOIN city cty ON p.city_id=cty.id
		WHERE email= $1 AND password=$2`
	
	row := r.db.QueryRow(stmt, email, password)

	err := row.Scan(
		&output.ID,
	)

	if err != nil {
		return &output, err
	}

	return &output, nil
}