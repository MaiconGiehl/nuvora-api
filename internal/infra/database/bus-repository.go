package database

import (
	"database/sql"
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

func (r *BusRepository) Save(number int, maxPassengers int) (error) {

	stmt := "INSERT INTO public.onibus (numero, max_passageiros, criado_em, atualizado_em) VALUES ($1, $2, $3, $4)"

	_, err := r.Db.Exec(stmt, number, maxPassengers, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}
	fmt.Print(err)

	return err
}

func (r *BusRepository) GetById(id int) (entity.Bus, error) {
	var bus entity.Bus

	rows, err := r.Db.Query(fmt.Sprintf("select id, numero, max_passageiros, criado_em, atualizado_em from public.onibus where id=%d", id))
	if err != nil {
		return bus, err
	}
	
	rows.Next()
	err = rows.Scan(&bus.Id, &bus.Number, &bus.MaxPassengers, &bus.CreatedIn, &bus.UpdatedIn)
	if err != nil {
		fmt.Print(err)
	}
	

	return bus, err
}

func (r *BusRepository) GetAll() ([]entity.Bus, error) {
	var allBus []entity.Bus
	rows, err := r.Db.Query("select id, numero, max_passageiros from public.onibus")
	if err != nil {
		return allBus, err
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
	return allBus, err
}

func (r *BusRepository) Delete(id int) (error) {
	_, err := r.Db.Query(fmt.Sprintf("delete from public.onibus  where id=%d", id))
	if err != nil {
		return err
	}

	return err
}

func (r *BusRepository) Update(id int, bus *dto.BusInputDTO) error {
	stmt := "update public.onibus SET numero = $2, max_passageiros = $3, atualizado_em = $4 where id= $1"

	_, err := r.Db.Exec(stmt, id, bus.Number, bus.MaxPassengers, time.Now().Format("2006-01-02 15:04:05"))
	fmt.Print(err)
	if err != nil {
		return err
	}

	return err
}