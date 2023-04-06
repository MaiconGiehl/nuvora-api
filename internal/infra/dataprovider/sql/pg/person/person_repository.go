package entity

import (
	"context"
	"database/sql"
)

type PersonPGSQLRepository struct {
	ctx context.Context
	db *sql.DB
}

func NewPersonPGSQLRepository(
	ctx context.Context,
	db *sql.DB,
) *PersonPGSQLRepository {
	return &PersonPGSQLRepository{
		ctx: ctx,
		db: db,
	}
}

func (r *PersonPGSQLRepository) GetPersonByAccount(personId int) (*Person, error) {
	var output Person

	return &output, nil
}

func (r *PersonPGSQLRepository) GetPersonByCompany(CompanyId int) (*Person, error) {
	var output Person

	// stmt := "SELECT * FROM person p WHERE p.company_id = $1"

	return &output, nil
}