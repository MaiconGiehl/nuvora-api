package database

import (
	"database/sql"
	"time"

	"github.com/maicongiehl/techtur-api/internal/dto"
	"github.com/maicongiehl/techtur-api/internal/entity"
)

type CompanyRepository struct {
	Db *sql.DB
}

func NewCompanyRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{Db: db}
}

func (r *CompanyRepository) GetAll() (*[]dto.CompanyOutputDTO, error) {
	var output []dto.CompanyOutputDTO
	rows, err := r.Db.Query("SELECT c.id, phone, fantasy_name, ct.description FROM company AS c JOIN company_type AS ct ON c.company_type_id=ct.id")
	if err != nil {
		return &output, err
	}

	for rows.Next() {
		var entity dto.CompanyOutputDTO
		err = rows.Scan(&entity.ID, &entity.Phone, &entity.FantasyName, &entity.CompanyType )
		if err != nil {
			return &[]dto.CompanyOutputDTO{}, err
		}
		output = append(output, entity)
	}
	
	return &output, err
}


func (r *CompanyRepository) Save(input *entity.Company) (error) {
	stmt := "INSERT INTO public.company (phone, cnpj, social_reason, fantasy_name, company_type_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	rows, err := r.Db.Exec(stmt, &input.Phone, &input.Cnpj, &input.SocialReason, &input.FantasyName, &input.CompanyTypeId,
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

func (r *CompanyRepository) Delete(input *entity.Company) error {
	stmt := "DELETE FROM public.company WHERE id= $1"

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

func (r *CompanyRepository) GetEmployees(input *entity.Company) (*[]dto.EmployeeOutputDTO, error) {
	var output []dto.EmployeeOutputDTO
	stmt := "SELECT a.id AS account_id, name FROM account a JOIN person p ON  a.person_id =p.id JOIN customer c  ON p.customer_id =c.id WHERE c.company_id = $1;"

	rows, err := r.Db.Query(stmt, input.ID)
	if err != nil {
		return &output, err
	}

	for rows.Next() {
		var entity dto.EmployeeOutputDTO
		err = rows.Scan(&entity.ID, &entity.Name)
		if err != nil {
			return &[]dto.EmployeeOutputDTO{}, err
		}
		output = append(output, entity)
	}
	
	return &output, err
}

func (r *CompanyRepository) GetLastMonthTickets(input *entity.Company) (*[]dto.EmployeesTicketsOutputDTO, error) {
	var output []dto.EmployeesTicketsOutputDTO
	stmt := `SELECT c.name, c.cpf, t.price, t.departure_time, dpt_city.name, t.arrival_time, arv_city."name", ts.description FROM ticket tkt 
		JOIN travel t ON tkt.travel_id=t.id 
		JOIN account a ON tkt.account_id =a.id
		JOIN person p ON a.person_id =a.person_id
		JOIN customer c ON p.customer_id =c.id
		JOIN city dpt_city ON t.departure_city_id =dpt_city.id
		JOIN city arv_city ON t.arrival_city_id=arv_city.id
		JOIN ticket_status ts ON tkt.status_id=ts.id 
		WHERE c.company_id = $1
		ORDER BY c."name", t.departure_time, t.arrival_time
		`
	
	rows, err := r.Db.Query(stmt, input.ID)
	if err != nil {
		return &output, err
	}

	for rows.Next() {
		var entity dto.EmployeesTicketsOutputDTO
		err = rows.Scan(
			&entity.Name,
			&entity.Cpf,
			&entity.Travel.Price,
			&entity.Travel.Departure.Time,
			&entity.Travel.Departure.City,
			&entity.Travel.Arrival.Time,
			&entity.Travel.Arrival.City,
			&entity.Status,
		)

		if err != nil {
			return &output, err
		}
		output = append(output, entity)
	}
	
	return &output, err
}