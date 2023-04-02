package database

import (
	"database/sql"
	"errors"
	"time"

	"github.com/maicongiehl/nuvera-api/internal/dto"
	"github.com/maicongiehl/nuvera-api/internal/entity"
)

type BusRepository struct {
	Db *sql.DB
}

func NewBusRepository(db *sql.DB) *BusRepository {
	return &BusRepository{Db: db}
}

func (r *BusRepository) Save(input *entity.Bus) (error) {
	stmt := "INSERT INTO bus (number, max_passengers, account_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"

	rows, err := r.Db.Exec(stmt, &input.Number, &input.MaxPassengers,&input.CompanyID,  time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}

	err = threatNotAffectData(rows)
	if err != nil {
		return err
	}

	return nil
}

func (r *BusRepository) GetById(input *entity.Bus) (*dto.BusOutputDTO, error) {
	var output dto.BusOutputDTO
	stmt := "SELECT id, number, max_passengers, account_id FROM bus WHERE id=$1"

	rows, err := r.Db.Query(stmt, input.Id)
	if err != nil {
		return &output, err
	}
	
	rows.Next()
	err = rows.Scan(&output.ID, &output.Number, &output.MaxPassengers, &output.CompanyID)
	if err != nil {
		return &output, err
	}

	return &output, nil
}

func (r *BusRepository) GetAll() (*[]dto.BusOutputDTO, error) {
	var allBus []dto.BusOutputDTO
	rows, err := r.Db.Query("SELECT id, number, max_passengers, account_id FROM bus")
	if err != nil {
		return &allBus, err
	}
	
	for rows.Next() {
		var bus dto.BusOutputDTO
		err = rows.Scan(&bus.ID, &bus.Number, &bus.MaxPassengers, &bus.CompanyID)
		if err != nil {
			return &[]dto.BusOutputDTO{}, err
		}
		allBus = append(allBus, bus)
	}
	
	return &allBus, err
}

func (r *BusRepository) Delete(input *entity.Bus) error {
	stmt := "DELETE FROM bus WHERE id= $1"

	rows, err := r.Db.Exec(stmt, &input.Id)
	if err != nil {
		return err
	}

	err = threatNotAffectData(rows)
	if err != nil {
		return err
	}

	return nil
}

func (r *BusRepository) Update(input *entity.Bus) error {
	stmt := "UPDATE bus SET number = $2, max_passengers = $3, updated_at = $4 WHERE id = $1"

	rows, err := r.Db.Exec(stmt, input.Id, input.Number, input.MaxPassengers, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}

	err = threatNotAffectData(rows)
	if err != nil {
		return err
	}

	return nil
}

func threatNotAffectData(rows sql.Result) error {
	affectedRows, _ := rows.RowsAffected()
	if affectedRows == 0 {
		return errors.New("data not found")
	}
	return nil
}