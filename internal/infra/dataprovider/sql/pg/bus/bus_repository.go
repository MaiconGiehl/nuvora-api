package entity

import (
	"context"
	"database/sql"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
)

type BusPGSQLRepository struct {
	ctx    context.Context
	db     *sql.DB
	logger logger.Logger
}

func NewBusPGSQLRepository(
	ctx context.Context,
	db *sql.DB,
	logger logger.Logger,
) *BusPGSQLRepository {
	return &BusPGSQLRepository{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

func (r *BusPGSQLRepository) FindBusByID(id int) (*Bus, error) {
	var output Bus

	stmt := "SELECT * FROM bus b WHERE b.id= $1"

	row := r.db.QueryRow(stmt, id)

	err := row.Scan(
		&output.ID,
		&output.Number,
		&output.MaxPassengers,
		&output.CompanyID,
		&output.CreatedAt,
		&output.UpdatedAt,
	)

	if err != nil {
		r.logger.Errorf("BusRepository.FindBusByID: Unable to find bus, %s", err)
		return &output, err
	}

	return &output, err
}

func (r *BusPGSQLRepository) FindBusByCompanyAccountID(id int) ([]*Bus, error) {

	var output []*Bus
	stmt := "SELECT * FROM bus b WHERE b.account_id= $1"

	rows, err := r.db.Query(stmt, id)
	if err != nil {
		return output, err
	}

	for rows.Next() {
		var bus Bus
		err = rows.Scan(
			&bus.ID,
			&bus.Number,
			&bus.MaxPassengers,
			&bus.CompanyID,
			&bus.CreatedAt,
			&bus.UpdatedAt,
		)

		if err != nil {
			return output, err
		}
		output = append(output, &bus)
	}

	return output, err
}
