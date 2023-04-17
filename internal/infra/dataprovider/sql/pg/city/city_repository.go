package entity

import (
	"context"
	"database/sql"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
)

type CityPGSQLRepository struct {
	ctx context.Context
	db *sql.DB
	logger logger.Logger
}

func NewCityPGSQLRepository(
	ctx context.Context,
	db *sql.DB,
	logger logger.Logger,
) *CityPGSQLRepository {
	return &CityPGSQLRepository{
		ctx: ctx,
		db: db,
		logger: logger,
	}
}

func (r *CityPGSQLRepository) FindCityByID(id int) (*City, error) {
	var output City
	stmt := "SELECT * FROM city c WHERE c.id= $1"

	row := r.db.QueryRow(stmt, id)

	err := row.Scan(
		&output.ID,
		&output.Name,
	)

	if err != nil {
		r.logger.Errorf("CityRepository.FindCityByID: Unable to find city, %s", err)
		return &output, err
	}

	return &output, nil
}

func (r *CityPGSQLRepository) FindAll() ([]*City, error) {
	var output []*City

	stmt := "SELECT * FROM city c"

	rows, err := r.db.Query(stmt)

	if err != nil {
		r.logger.Errorf("CityRepository.FindAll: Unable to find cities, %s", err)
		return output, err
	}

	for rows.Next() {
		var city City
		err := rows.Scan(
			&city.ID,
			&city.Name,
		)

		if err != nil {
			return output, err
		}
		output = append(output, &city)
	}

	return output, nil
}