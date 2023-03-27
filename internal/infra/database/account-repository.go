package database

import (
	"database/sql"
	"time"

	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/entity"
)

type AccountRepository struct {
	Db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{Db: db}
}

func (r *AccountRepository) SaveCustomerAccount(input *entity.Account) (error) {
	stmt := "INSERT INTO account (username, email, password, person_id, daily_tickets, tickets_left, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	rows, err := r.Db.Exec(stmt, 	&input.Username, &input.Email, &input.Password, &input.PersonID, &input.DailyTickets, &input.TicketsLeft, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
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
	stmt := `SELECT a.id, username, a.tickets_left, p.permission_level, cty.name FROM account a JOIN person p ON a.person_id =p.id JOIN customer c ON p.customer_id =c.id JOIN city cty ON p.city_id=cty.id
		WHERE email=$1 AND password=$2`
	
	
	rows, err := r.Db.Query(stmt, input.Email, input.Password)
	if err != nil {
		return &output, err
	}
	
	rows.Next()
	err = rows.Scan(
		&output.ID,
		&output.Username,
		&output.TicketsLeft,
		&output.PermissionLevel,
		&output.City,
	)
	
	if err != nil {
		return &output, err
	}

	return &output, nil
}