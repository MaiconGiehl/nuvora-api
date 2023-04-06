package entity

import (
	"context"
	"database/sql"
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

func (r *CustomerPGSQLRepository) GetCustomerByPerson(id int) (*Customer, error) {
	var output Customer

	stmt := `SELECT * FROM customer c WHERE c.id=$1`
	
	row := r.db.QueryRow(stmt, id)

	err := row.Scan(
		&output.ID,
		&output.Name,
		&output.CompanyID,
	)

	if err != nil {
		return &output, err
	}

	return &output, nil
}
