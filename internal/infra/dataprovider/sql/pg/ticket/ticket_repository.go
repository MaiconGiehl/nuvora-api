package entity

import (
	"context"
	"database/sql"
)

type TicketPGSQLRepository struct {
	ctx context.Context
	db *sql.DB
}

func NewTicketPGSQLRepository(
	ctx context.Context,
	db *sql.DB,
) *TicketPGSQLRepository {
	return &TicketPGSQLRepository{
		ctx: ctx,
		db: db,
	}
}

func (r *TicketPGSQLRepository) CreateTicket() error {
	return nil
}

func (r *TicketPGSQLRepository) GetLastPurchases(accountId int) (*[]Ticket, error) {
	var output []Ticket

	stmt := "SELECT * FROM ticket WHERE account_id = $1 ORDER BY created_at DESC"

	
	rows, err := r.db.Query(stmt, accountId)
	if err != nil {
		return &output, err
	}

	for rows.Next() {
		var travel Ticket
		err = rows.Scan(
			&travel.ID,
		)
		if err != nil {
			return &output, err
		}
		output = append(output, travel)
	}
	
	return &output, nil
}

func (r *TicketPGSQLRepository) GetEmployeesTickets(accountId int) (*[]EmployeeTravelTicket, error) {
	var output []EmployeeTravelTicket

	stmt := "SELECT * FROM ticket WHERE account_id = $1 ORDER BY created_at DESC"

	
	rows, err := r.db.Query(stmt, accountId)
	if err != nil {
		return &output, err
	}

	for rows.Next() {
		var travel EmployeeTravelTicket
		err = rows.Scan(
			&travel.Name,
		)
		if err != nil {
			return &output, err
		}
		output = append(output, travel)
	}
	
	return &output, nil
}
