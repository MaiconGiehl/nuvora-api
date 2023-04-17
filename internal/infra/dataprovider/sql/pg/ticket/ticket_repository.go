package entity

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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
	stmt := "INSERT INTO ticket (account_id, status_id, travel_id, created_at) VALUES ($1, 0, $2, NOW())"

	_, err := r.db.Exec(stmt, accountId, travelId)

	if err != nil {
		r.logger.Errorf("TicketRepository.BuyTicket: Unable to buy a ticket, %s", err)
		err = errors.New("internal error, please try again in some minutes")
		return err
	}

	return nil
}

func (r *TicketPGSQLRepository) GetPurchases(accountId int) (*[]Ticket, error) {
	var output []Ticket

	stmt := "SELECT * FROM ticket WHERE account_id = $1 ORDER BY created_at DESC"

	rows, err := r.db.Query(stmt, accountId)
	if err != nil {
		return &output, err
	}

	for rows.Next() {
		var ticket Ticket
		err = rows.Scan(
			&ticket.ID,
			&ticket.AccountID,
			&ticket.StatusID,
			&ticket.TravelID,
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
		)
		if err != nil {
			return &output, err
		}
		output = append(output, ticket)
	}

	return &output, nil
}

func (r *TicketPGSQLRepository) GetEmployeesTickets(accountId int) ([]*Ticket, error) {
	var output []*Ticket

	stmt := `
		SELECT * FROM ticket t WHERE t.account_id IN (
			SELECT a.id FROM account a 
			JOIN person p ON a.person_id =p.id 
			JOIN customer c ON p.customer_id =c.id 
			WHERE c.company_id = $1 ) 
		ORDER BY created_at
`

	rows, err := r.db.Query(stmt, accountId)
	if err != nil {
		return output, err
	}

	for rows.Next() {
		var ticket Ticket
		err = rows.Scan(
			&ticket.ID, 
			&ticket.AccountID, 
			&ticket.StatusID, 
			&ticket.TravelID, 
			&ticket.CreatedAt,
			&ticket.UpdatedAt,
		)
		if err != nil {
			return output, err
		}
		output = append(output, &ticket)
	}
	
	return output, nil
}

func (r *TicketPGSQLRepository) UpdateTicketsStatusByCompanyID(companyId int) (string, error) {
	stmt := `UPDATE ticket t SET status_id = 1, updated_at = NOW() WHERE t.id IN (
		SELECT t.id FROM ticket t 
		JOIN account a ON t.account_id=a.id 
		JOIN person p ON a.person_id =p.id 
		JOIN customer c ON p.customer_id =c.id  
		WHERE c.company_id=$1
	) AND status_id = 0`

	
	rows, err := r.db.Exec(stmt, companyId)
	if err != nil {
		return "", err
	}
	
	affectedRows, err := rows.RowsAffected()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d tickets paid", affectedRows), nil
}
