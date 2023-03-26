package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/entity"
)

type TravelRepository struct {
	Db *sql.DB
}

func NewTravelRepository(db *sql.DB) *TravelRepository {
	return &TravelRepository{Db: db}
}

func (r *TravelRepository) GetAll() (*[]dto.TravelOutputDTO, error) {
	var output []dto.TravelOutputDTO
	rows, err := r.Db.Query(
		`SELECT 
			t.id,
			t.price,
			ac.id,
			ac.fantasy_name,
			ct.id, ct.description,
			ac
		FROM
			public.travel AS t 
		JOIN 
			public.account AS ac 
			ON t.company_account_id=ac.id
		JOIN 
			public.company_type AS ct
			ON ac.company_type_id=ct.id`,
	)
	if err != nil {
		return &output, err
	}
	
	for rows.Next() {
		var travel dto.TravelOutputDTO
		err = rows.Scan(
			&travel.ID, 
			&travel.Price,
			&travel.Company.ID,
			&travel.Company.FantasyName,
			// &travel.Company.CompanyType.ID,
			// &travel.Company.CompanyType.Description,
			&travel.Bus.ID, 
			&travel.Bus.Number,
			&travel.Bus.MaxPassengers,
			&travel.DepartureTime,
			&travel.DepartureCity.ID,
			&travel.DepartureCity.Name,
			&travel.ArrivalTime,
			&travel.ArrivalCity.Name,
		)

		if err != nil {
			return &[]dto.TravelOutputDTO{}, err
		}

		output = append(output, travel)
	}

	return &output, err
}

func (r *TravelRepository) Save(input *entity.Travel) (error) {
	stmt := "INSERT INTO public.travel (price, company_account_id, bus_id, departure_time, departure_city_id, arrival_time, arrival_city_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	rows, err := r.Db.Exec(stmt, input.Price, input.CompanyAccountId, input.BusID, input.DepartureTime, input.DepartureCityId, input.ArrivalTime, input.ArrivalCityId,
		time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	fmt.Print(err)
	if err != nil {
		return err
	}


	err = threatNotAffectData(rows)
	if err != nil {
		return err
	}

	return nil
}

func (r *TravelRepository) Delete(input *entity.Travel) error {
	stmt := "DELETE FROM public.travel WHERE id= $1"

	rows, err := r.Db.Exec(stmt, &input.ID)
	if err != nil {
		return err
	}

	err = threatNotAffectData(rows)
	if err != nil {
		return err
	}

	return nil
}