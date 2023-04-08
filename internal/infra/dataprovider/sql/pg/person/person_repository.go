package entity

import (
	"context"
	"database/sql"
	"errors"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
)

type PersonPGSQLRepository struct {
	ctx context.Context
	db *sql.DB
	logger logger.Logger
}

func NewPersonPGSQLRepository(
	ctx context.Context,
	db *sql.DB,
	logger logger.Logger,
) *PersonPGSQLRepository {
	return &PersonPGSQLRepository{
		ctx: ctx,
		db: db,
		logger: logger,
	}
}

func (r *PersonPGSQLRepository) GetPersonByID(id int) (*Person, error) {
	var output Person

	stmt := "SELECT * FROM person p WHERE p.id= $1"

	row := r.db.QueryRow(stmt, id)

	err := row.Scan(
		&output.ID,
		&output.PermissionLevel,
		&output.CustomerID,
		&output.CompanyID,
		&output.CityID,
		&output.CreatedAt,
		&output.UpdatedAt,
	)

	if err != nil {
		r.logger.Errorf("PersonRepository.GetPersonByID: Unable to find person, %s", err)
		err = errors.New("internal error, please try again in some minutes")
		return &output, err
	}

	return &output, nil
}