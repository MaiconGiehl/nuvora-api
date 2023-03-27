package database

import (
	"database/sql"
	"time"

	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/entity"
)

type CustomerRepository struct {
	Db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{Db: db}
}

func (r *CustomerRepository) Save(input *entity.Customer) (error) {
	stmt := "INSERT INTO customer (cpf, name, phone, birth_date, company_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	rows, err := r.Db.Exec(stmt, &input.Cpf, &input.Name, &input.Phone, &input.BirthDate, &input.CompanyID, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}

	err = threatNotAffectData(rows)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepository) GetAll() (*[]dto.CustomerOutputDTO, error) {
	var allCustomer []dto.CustomerOutputDTO
	rows, err := r.Db.Query("SELECT id, name, phone, company_id FROM customer")
	if err != nil {
		return &allCustomer, err
	}
	
	for rows.Next() {
		var customer dto.CustomerOutputDTO
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Phone, &customer.CompanyID)
		if err != nil {
			return &[]dto.CustomerOutputDTO{}, err
		}
		allCustomer = append(allCustomer, customer)
	}
	
	return &allCustomer, err
}

func (r *CustomerRepository) Delete(input *entity.Customer) error {
	stmt := "DELETE FROM customer WHERE id= $1"

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

func (r *CustomerRepository) GetLastId() (int, error) {
	stmt :=  "SELECT id FROM customer ORDER BY id DESC LIMIT 1"
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