package entity

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
)

type TravelPGSQLRepository struct {
	ctx    context.Context
	db     *sql.DB
	logger logger.Logger
}

func NewTravelPGSQLRepository(
	ctx context.Context,
	db *sql.DB,
	logger logger.Logger,
) *TravelPGSQLRepository {
	return &TravelPGSQLRepository{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

func (r *TravelPGSQLRepository) CreateTravel(
	companyId int,
	price float64,
	busId int,
	departureTime time.Time,
	departureCityId int,
	arrivalTime time.Time,
	arrivalCityId int,
) error {

	stmt := `INSERT INTO travel ( price, account_id, bus_id, status, departure_time, departure_city_id, arrival_time, arrival_city_id, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, CURRENT_DATE)
		`
	_, err := r.db.Exec(stmt, price, companyId, busId, 0, departureTime, departureCityId, arrivalTime, arrivalCityId)

	if err != nil {
		return err
	}

	return nil
}

func (r *TravelPGSQLRepository) FindTravelByID(id int) (*Travel, error) {
	var output Travel

	stmt := `SELECT * FROM travel t WHERE t.id = $1`

	row := r.db.QueryRow(stmt, id)

	err := row.Scan(
		&output.ID,
		&output.Price,
		&output.CompanyID,
		&output.Bus.Number,
		&output.Bus.MaxPassengers,
		&output.Departure.Time,
		&output.Departure.CityName,
		&output.Arrival.Time,
		&output.Arrival.CityName,
		&output.CreatedAt,
		&output.UpdatedAt,
	)

	if err != nil {
		r.logger.Errorf("TravelRepository.FindTravelByID: Unable to find travel, %s", err)
		err = errors.New("internal error, please try again in some minutes")
		return &output, err
	}

	return &output, err
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
