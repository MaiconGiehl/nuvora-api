package database

import (
	"database/sql"
	"time"

	"github.com/maicongiehl/nuvera-api/internal/dto"
	"github.com/maicongiehl/nuvera-api/internal/entity"
)

type AccountRepository struct {
	Db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{Db: db}
}

func (r *AccountRepository) SaveCustomerAccount(input *entity.Account) (error) {
	stmt := "INSERT INTO account (email, password, person_id, daily_tickets, tickets_left, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	rows, err := r.Db.Exec(stmt, &input.Email, &input.Password, &input.PersonID, &input.DailyTickets, &input.TicketsLeft, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}

	err = threatNotAffectData(rows)
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) GetCustomerAccount(input *entity.Account) (*dto.CustomerAccountOutputDTO, error) {
	var output dto.CustomerAccountOutputDTO
	stmt := `SELECT a.id, a.tickets_left, p.permission_level, cty.id AS city_id FROM account a LEFT JOIN person p ON a.person_id =p.id LEFT JOIN customer c ON p.customer_id =c.id LEFT JOIN city cty ON p.city_id=cty.id
		WHERE email= $1 AND password=$2`
	
	rows, err := r.Db.Query(stmt, input.Email, input.Password)
	if err != nil {
		return &output, err
	}
	
	rows.Next()
	err = rows.Scan(
		&output.ID,
		&output.PermissionLevel,
		&output.TicketsLeft,
		&output.CityID,
	)
	
	if err != nil {
		return &output, err
	}

	return &output, nil
}

func (r *AccountRepository) GetCompanyAccount(input *entity.Account) (*dto.CompanyAccountOutputDTO, error) {
	var output dto.CompanyAccountOutputDTO
	stmt := `SELECT a.id, p.permission_level, cty.id FROM account a LEFT JOIN person p ON a.person_id =p.id LEFT JOIN company c ON p.company_id =c.id LEFT JOIN city cty ON p.city_id=cty.id
		WHERE email= $1 AND password=$2`
	
	rows, err := r.Db.Query(stmt, input.Email, input.Password)
	if err != nil {
		return &output, err
	}
	
	rows.Next()
	err = rows.Scan(
		&output.ID,
		&output.PermissionLevel,
		&output.CityID,
	)
	
	if err != nil {
		return &output, err
	}

	return &output, nil
}