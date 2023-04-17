package entity

import (
	"context"
	"database/sql"
	"errors"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
)

type AccountPGSQLRepository struct {
	ctx context.Context
	db *sql.DB
	logger logger.Logger
}

func NewAccountPGSQLRepository(
	ctx context.Context,
	db *sql.DB,
	logger logger.Logger,
) *AccountPGSQLRepository {
	return &AccountPGSQLRepository{
		ctx: ctx,
		db: db,
		logger: logger,
	}
}

func (r *AccountPGSQLRepository) FindAccountByID(accountId int) (*Account, error) {
	var output Account

	stmt := `SELECT * FROM account a WHERE id = $1`
	
	row := r.db.QueryRow(stmt, accountId)

	err := row.Scan(
		&output.ID,
		&output.Email,
		&output.Password,
		&output.PersonID,
		&output.LastAccess,
		&output.TicketsLeft,
		&output.DailyTickets,
		&output.CreatedAt,
		&output.UpdatedAt,
	)

	if err != nil {
		r.logger.Errorf("AccountRepository.FindAccountByID: Unable to find account, %s", err)
		err = errors.New("invalid account")
		return &output, err
	}

	return &output, nil
}

func (r *AccountPGSQLRepository) Login(email, password string) (*Account, error) {
	var output Account
	stmt := `SELECT * FROM account a WHERE email= $1 AND password=$2`
	
	row := r.db.QueryRow(stmt, email, password)

	err := row.Scan(
		&output.ID,
		&output.Email,
		&output.Password,
		&output.PersonID,
		&output.LastAccess,
		&output.TicketsLeft,
		&output.DailyTickets,
		&output.CreatedAt,
		&output.UpdatedAt,
	)

	if err != nil {
		r.logger.Errorf("AccountRepository.Login: Unable to find account, %s", err)
		err = errors.New("invalid credentials")
		return &output, err
	}

	return &output, nil
}

func (r *AccountPGSQLRepository) RemoveTicket(id int) error {
	stmt := `UPDATE account SET tickets_left = tickets_left - 1 WHERE id = $1`
	
	_, err := r.db.Exec(stmt, id)

	if err != nil {
		r.logger.Errorf("AccountRepository.RemoveTicket: Unable to remove ticket, %s", err)
		return err
	}

	return nil
}

func (r *AccountPGSQLRepository) FindAccountsByCompanyID(id int) ([]*Account, error) {
	var output []*Account
	stmt := `SELECT * FROM account a WHERE a.id IN (SELECT a.id FROM account a JOIN person p ON a.person_id=p.id JOIN customer c ON p.customer_id=c.id WHERE c.company_id = $1 )`
	
	rows, err := r.db.Query(stmt, id)

	if err != nil {
		r.logger.Errorf("AccountRepository.FindAccountsByCompanyID: Unable to find account, %s", err)
		err = errors.New("invalid credentials")
		return output, err
	}

	for rows.Next() {
		var account Account
		err := rows.Scan(
			&account.ID,
			&account.Email,
			&account.Password,
			&account.PersonID,
			&account.LastAccess,
			&account.TicketsLeft,
			&account.DailyTickets,
			&account.CreatedAt,
			&account.UpdatedAt,
		)

		if err != nil {
			return output, err
		}
		output = append(output, &account)
	}

	return output, nil
}


func (r *AccountPGSQLRepository) DeleteAccountByID(id int) error {
	stmt := `DELETE FROM person p WHERE id = $1`
	
	_, err := r.db.Exec(stmt, id)
	if err != nil {
		return err
	}

	return err
}

