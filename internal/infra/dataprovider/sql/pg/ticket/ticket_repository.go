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

func (r *TicketPGSQLRepository) GetLastPurchases(email, password string) (*Ticket, error) {
	var output Ticket

	return &output, nil
}