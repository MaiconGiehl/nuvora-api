package entity

import (
	"context"
	"database/sql"
	"errors"

	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
)

type TicketPGSQLRepository struct {
	ctx    context.Context
	db     *sql.DB
	logger logger.Logger
}

func NewTicketPGSQLRepository(
	ctx context.Context,
	db *sql.DB,
	logger logger.Logger,
) *TicketPGSQLRepository {
	return &TicketPGSQLRepository{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

func (r *TicketPGSQLRepository) CreateTicket(accountId, travelId int) error {
	stmt := "INSERT INTO ticket (account_id, status_id, travel_id, created_at) VALUES ($1, 0, $2, CURRENT_DATE)"

	_, err := r.db.Exec(stmt, accountId, travelId)

	if err != nil {
		r.logger.Errorf("TicketRepository.BuyTicket: Unable to buy a ticket, %s", err)
		err = errors.New("internal error, please try again in some minutes")
		return err
	}

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

func (r *TicketPGSQLRepository) UpdateTickets(accountId int) (*[]EmployeeTravelTicket, error) {
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
