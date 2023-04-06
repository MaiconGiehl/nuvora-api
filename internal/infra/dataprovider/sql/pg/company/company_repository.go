package entity

import (
	"context"
	"database/sql"
	"fmt"
)

type CompanyPGSQLRepository struct {
	ctx context.Context
	db *sql.DB
}

func NewCompanyPGSQLRepository(
	ctx context.Context,
	db *sql.DB,
) *CompanyPGSQLRepository {

	return &CompanyPGSQLRepository{
		ctx: ctx,
		db: db,
	}
}

func (r *CompanyPGSQLRepository) Login(email, password string) (*Company, error) {
	var output Company

	return &output, nil
}

func (r *CompanyPGSQLRepository) GetCompany(companyId int) (*Company, error) {
	var output Company
	stmt := `SELECT * FROM account a WHERE p.company_id =$1`
	
	row := r.db.QueryRow(stmt, companyId)

	err := row.Scan(
		&output.ID,
	)

	if err != nil {
		fmt.Print(err)
		return &output, err
	}

	return &output, nil
}