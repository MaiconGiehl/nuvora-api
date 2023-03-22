package database

import (
	"database/sql"
	"errors"
	"fmt"
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

func (r *BusRepository) Save(bus *dto.BusInputDTO) (error) {
	stmt := "INSERT INTO public.onibus (numero, max_passageiros, criado_em, atualizado_em) VALUES ($1, $2, $3, $4)"

	rows, err := r.Db.Exec(stmt, &bus.Number, &bus.MaxPassengers, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}

	err = threatNotAffectData(rows)
	if err != nil {
		return err
	}

	return nil
}

func (r *BusRepository) GetById(id int) (*entity.Bus, error) {
	var bus entity.Bus
	stmt := "SELECT id, numero, max_passageiros, criado_em, atualizado_em FROM public.onibus WHERE id=$1"


	rows, err := r.Db.Query(stmt, id)
	if err != nil {
		return &bus, err
	}
	
	rows.Next()
	err = rows.Scan(&bus.Id, &bus.Number, &bus.MaxPassengers, &bus.CreatedIn, &bus.UpdatedIn)
	if err != nil {
		return &bus, errors.New(err.Error())
	}

	return &bus, nil
}

func (r *BusRepository) GetAll() (*[]entity.Bus, error) {
	var allBus []entity.Bus
	rows, err := r.Db.Query("SELECT id, numero, max_passageiros FROM public.onibus")
	if err != nil {
		return &allBus, err
	}
	
	for rows.Next() {
		var bus entity.Bus
		err = rows.Scan(&bus.Id, &bus.Number, &bus.MaxPassengers)
		if err != nil {
			fmt.Print(err)
			break
		}
		allBus = append(allBus, bus)
	}
	
	return &allBus, err
}

func (r *BusRepository) Delete(id int) error {
	stmt := "DELETE FROM public.onibus WHERE id= $1"

	rows, err := r.Db.Exec(stmt, id)
	if err != nil {
		return err
	}

	err = threatNotAffectData(rows)
	if err != nil {
		return err
	}

	return nil
}

func (r *BusRepository) Update(id int, bus *dto.BusInputDTO) error {
	stmt := "UPDATE public.onibus SET numero = $2, max_passageiros = $3, atualizado_em = $4 WHERE id = $1"

	rows, err := r.Db.Exec(stmt, id, bus.Number, bus.MaxPassengers, time.Now().Format("2006-01-02 15:04:05"))
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