package database

import (
	"database/sql"
	"time"

	"github.com/maicongiehl/nuvera-api/internal/entity"
)

type PersonRepository struct {
	Db *sql.DB
}

func NewPersonRepository(db *sql.DB) *PersonRepository {
	return &PersonRepository{Db: db}
}

func (r *PersonRepository) SaveCustomerPerson(input *entity.Person) (error) {
	stmt := "INSERT INTO person (permission_level, customer_id, city_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"

	rows, err := r.Db.Exec(stmt, &input.PermissionLevel, &input.CustomerID, &input.CityID, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}
	err = threatNotAffectData(rows)
	if err != nil {
		return err
	}

	return nil
}

func (r *PersonRepository) Delete(input *entity.Person) error {
	stmt := "DELETE FROM person WHERE id= $1"

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

func (r *PersonRepository) GetLastId() (int, error) {
	stmt :=  "SELECT id FROM person ORDER BY id DESC LIMIT 1"
	row, err := r.Db.Query(stmt)
	var lastId int
	if err != nil {
		return lastId, err
	}

	row.Next()
	err = row.Scan(&lastId)
	if err != nil {
		return lastId, err
	}

	return lastId, nil 
}