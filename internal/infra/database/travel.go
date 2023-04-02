package database

import (
	"database/sql"
	"time"

	"github.com/maicongiehl/nuvera-api/internal/dto"
	"github.com/maicongiehl/nuvera-api/internal/entity"
)

type TravelRepository struct {
	Db *sql.DB
}

func NewTravelRepository(db *sql.DB) *TravelRepository {
	return &TravelRepository{Db: db}
}

func (r *TravelRepository) GetAllByDestiny(entity *entity.Travel) (*[]dto.TravelOutputDTO, error) {
	var output []dto.TravelOutputDTO
	stmt := 
	`SELECT 
		t.id,
		price, 
		c.fantasy_name, 
		b."number" , 
		b.max_passengers, 
		t.departure_time,
		dpt_c.name,
		t.arrival_time,
		arv_c.name
	FROM
		travel t 
	LEFT JOIN
		account a ON t.account_id=a.id 
	LEFT JOIN
		person p ON a.person_id =p.id  
	LEFT JOIN
		company c ON p.company_id =c.id
	LEFT JOIN
		bus b  ON t.bus_id =b.id
	LEFT JOIN
		city dpt_c ON t.departure_city_id=dpt_c.id
	LEFT JOIN
		city arv_c ON t.arrival_city_id =arv_c.id
	WHERE
		arv_c.id = $1 AND dpt_c.id = $2
	ORDER BY
		t.departure_time
	`

	rows, err := r.Db.Query(stmt, entity.DepartureCityID, entity.ArrivalCityID)
	if err != nil {
		return &output, err
	}
	

	for rows.Next() {
		var travel dto.TravelOutputDTO
		err = rows.Scan(
			&travel.ID,
			&travel.Price,
			&travel.FantasyName,
			&travel.BusNumber,
			&travel.MaxPassengers,
			&travel.DepartureTime,
			&travel.DepartureCity,
			&travel.ArrivalTime,
			&travel.ArrivalCity,
		)

		if err != nil {
			return &[]dto.TravelOutputDTO{}, err
		}

		output = append(output, travel)
	}

	return &output, err
}

func (r *TravelRepository) Save(input *entity.Travel) (error) {
	stmt := "INSERT INTO public.travel (price, account_id, bus_id, departure_time, departure_city_id, arrival_time, arrival_city_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"

	rows, err := r.Db.Exec(stmt, input.Price, input.AccountID, input.BusID, input.DepartureTime, input.DepartureCityID, input.ArrivalTime, input.ArrivalCityID,
		time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
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