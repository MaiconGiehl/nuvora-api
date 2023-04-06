package entity

import (
	"context"
	"database/sql"
	"fmt"
)

type CustomerPGSQLRepository struct {
	ctx context.Context
	db *sql.DB
}

func NewCustomerPGSQLRepository(
	ctx context.Context,
	db *sql.DB,
) *CustomerPGSQLRepository {

	return &CustomerPGSQLRepository{
		ctx: ctx,
		db: db,
	}
}

func (r *CustomerPGSQLRepository) Login(email, password string) (*Customer, error) {
	var output Customer
	stmt := `SELECT a.id AS id, c.name AS name FROM account a LEFT JOIN person p ON a.person_id =p.id LEFT JOIN customer c ON p.customer_id =c.id LEFT JOIN city cty ON p.city_id=cty.id
		WHERE email= $1 AND password=$2`
	
	rows := r.db.QueryRow(stmt, email, password)

	err := rows.Scan(
		&output.id,
		&output.name,
	)

	fmt.Print(output.id)
	fmt.Print(output.name)
	
	if err != nil {
		return &output, err
	}

	return &output, nil
}