package database

import (
	"database/sql"

	"github.com/maicongiehl/techtur-api/internal/entity"
)

type CityRepository struct {
	Db *sql.DB
}

func NewCityRepository(db *sql.DB) *CityRepository {
	return &CityRepository{Db: db}
}

func (r *CityRepository) Save(input *entity.City) (error) {
	stmt := "INSERT INTO public.city (name) VALUES ($1)"

	rows, err := r.Db.Exec(stmt, &input.Name)
	if err != nil {
		return err
	}

	err = threatNotAffectData(rows)
	if err != nil {
		return err
	}

	return nil
}

func (r *CityRepository) Delete(input *entity.City) error {
	stmt := "DELETE FROM public.onicity WHERE id= $1"

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