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

func (r *AccountPGSQLRepository) GetAccountByID(accountId int) (*Account, error) {
	var output Account

	return &output, nil
}

func (r *AccountPGSQLRepository) LoginAsCustomer(email, password string) (*Account, error) {
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
		r.logger.Errorf("AccountRepository.LoginAsCustomer: Unable to find account, %s", err)
		err = errors.New("invalid credentials")
		return &output, err
	}

	return &output, nil
}


func (r *AccountPGSQLRepository) LoginAsCompany(email, password string) (*Account, error) {
	var output Account

	

	return &output, nil
}