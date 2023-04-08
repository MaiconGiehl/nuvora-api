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

func (r *TravelPGSQLRepository) CreateTravel() error {
	return nil
}

func (r *TravelPGSQLRepository) GetTravelsByCities(dptCityID, arvCityID int) (*[]Travel, error) {
	var output []Travel

	stmt := `
		SELECT * FROM travel 
		WHERE 
			departure_city_id = $1 OR departure_city_id = $2 
		AND 
			arrival_departure_city_id = $1 OR arrival_departure_city_id = $2 
		ORDER BY 
			departure_city_id`
	
	rows, err := r.db.Query(stmt, dptCityID, arvCityID)
	if err != nil {
		return &output, err
	}

	for rows.Next() {
		var travel Travel
		err = rows.Scan(
			&travel.ID,
			&travel.Price,
			&travel.CompanyID,
			&travel.CompanyFantasyName,
			&travel.Bus.Number,
			&travel.Bus.MaxPassengers,
			&travel.Departure.Time,
			&travel.Departure.CityName,
			&travel.Arrival.Time,
			&travel.Arrival.CityName,
		)
		if err != nil {
			return &output, err
		}
		output = append(output, travel)
	}
	
	return &output, err
}