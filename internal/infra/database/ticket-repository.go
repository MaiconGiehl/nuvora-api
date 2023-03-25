package database

import (
	"database/sql"
	"time"

	"github.com/MaiconGiehl/API/internal/entity"
)

type TicketRepository struct {
	Db *sql.DB
}

func NewTicketRepository(db *sql.DB) *TicketRepository {
	return &TicketRepository{Db: db}
}

func (r *TicketRepository) Save(input *entity.Ticket) (error) {
	stmt := "INSERT INTO public.ticket (account_id, price, status, travel_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)"

	rows, err := r.Db.Exec(stmt, &input.AccountID, &input.Status, &input.TravelID,
		time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}

	err = threatNotAffectData(rows)
	if err != nil {
		return err
	}

	return nil
}

func (r *TicketRepository) Delete(input *entity.Ticket) error {
	stmt := "DELETE FROM public.oniticket WHERE id= $1"

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