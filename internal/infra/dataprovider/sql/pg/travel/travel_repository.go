package entity

import (
	"context"
	"database/sql"
)

type TravelPGSQLRepository struct {
	ctx context.Context
	db *sql.DB
}

func NewTravelPGSQLRepository(
	ctx context.Context,
	db *sql.DB,
) *TravelPGSQLRepository {
	return &TravelPGSQLRepository{
		ctx: ctx,
		db: db,
	}
}

func (r *TravelPGSQLRepository) GetpPossibleTravels(departureCityId, arrivalCityId int, ) (*Travel, error) {
	var output Travel

	

	return &output, nil
}