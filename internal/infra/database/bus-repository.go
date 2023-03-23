package database

import (
	"database/sql"
	"errors"
	"time"

	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/entity"
)

type BusRepository struct {
	Db *sql.DB
}

func NewBusRepository(db *sql.DB) *BusRepository {
	return &BusRepository{Db: db}
}

func (r *BusRepository) Save(input *entity.Bus) (error) {
	stmt := "INSERT INTO public.onibus (numero, max_passageiros, criado_em, atualizado_em) VALUES ($1, $2, $3, $4)"

	rows, err := r.Db.Exec(stmt, &input.Number, &input.MaxPassengers, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
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
	stmt := "SELECT id, numero, max_passageiros FROM public.onibus WHERE id=$1"

	rows, err := r.Db.Query(stmt, input.Id)
	if err != nil {
		return &output, err
	}
	
	rows.Next()
	err = rows.Scan(&output.ID, &output.Number, &output.MaxPassengers)
	if err != nil {
		return &output, err
	}

	return &output, nil
}

func (r *BusRepository) GetAll() (*[]dto.BusOutputDTO, error) {
	var allBus []dto.BusOutputDTO
	rows, err := r.Db.Query("SELECT id, numero, max_passageiros FROM public.onibus")
	if err != nil {
		return &allBus, err
	}
	
	for rows.Next() {
		var bus dto.BusOutputDTO
		err = rows.Scan(&bus.ID, &bus.Number, &bus.MaxPassengers)
		if err != nil {
			return &[]dto.BusOutputDTO{}, err
		}
		allBus = append(allBus, bus)
	}
	
	return &allBus, err
}

func (r *BusRepository) Delete(input *entity.Bus) error {
	stmt := "DELETE FROM public.onibus WHERE id= $1"

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
	stmt := "UPDATE public.onibus SET numero = $2, max_passageiros = $3, atualizado_em = $4 WHERE id = $1"

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