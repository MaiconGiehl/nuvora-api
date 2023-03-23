package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/entity"
)

type TicketRepository struct {
	Db *sql.DB
}

func NewTicketRepository(db *sql.DB) *TicketRepository {
	return &TicketRepository{Db: db}
}

func (r *TicketRepository) Save(input *entity.Ticket) (error) {
	stmt := "INSERT INTO public.passagem (cliente_id, preco, situacao, origem_cidade, origem_hora, destino_cidade, destino_hora, onibus_id, criado_em, atualizado_em) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	rows, err := r.Db.Exec(stmt, &input.ClienteID, &input.Price, &input.Status, &input.DepartureCity, &input.DepartureTime, &input.DestinyCity, &input.DestinyTime, &input.BusID,
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

func (r *TicketRepository) GetById(input *entity.Ticket) (*dto.TicketOutputDTO, error) {
	var output dto.TicketOutputDTO
	stmt := "SELECT p.id, p.preco, p.situacao, p.origem_cidade, p.origem_hora, p.destino_cidade, p.destino_hora, o.numero FROM public.passagem AS p JOIN public.onibus AS o ON p.onibus_id=o.id WHERE p.id=$1"


	rows, err := r.Db.Query(stmt, input.ID)
	if err != nil {
		return &output, err
	}
	
	rows.Next()
	err = rows.Scan(&output.ID, &output.Price, &output.Status, &output.DepartureCity, &output.DepartureTime, &output.DestinyCity, &output.DestinyTime, &output.BusNumber)
	if err != nil {
		return &output, errors.New(err.Error())
	}

	return &output, nil
}

func (r *TicketRepository) GetAll() (*[]dto.TicketOutputDTO, error) {
	var output []dto.TicketOutputDTO
	rows, err := r.Db.Query("SELECT p.id, p.preco, p.situacao, p.origem_cidade, p.origem_hora, p.destino_cidade, p.destino_hora, o.numero FROM public.passagem AS p JOIN public.onibus AS o ON p.onibus_id=o.id")
	if err != nil {
		return &output, err
	}
	
	for rows.Next() {
		var ticket dto.TicketOutputDTO
		err = rows.Scan(&ticket.ID, &ticket.Price, &ticket.Status, &ticket.DepartureCity, &ticket.DepartureTime, &ticket.DestinyCity, &ticket.DestinyTime, &ticket.BusNumber)
		if err != nil {
			fmt.Print(err)
			break
		}
		output = append(output, ticket)
	}
	
	return &output, nil
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

func (r *TicketRepository) Update(input *entity.Ticket) error {
	stmt := "UPDATE public.passagem SET preco = $2, situacao = $3, origem_cidade = $4, origem_hora = $5, destino_cidade = $6, destino_hora = $7, onibus_id = $8, atualizado_em = $9 WHERE id = $1"

	rows, err := r.Db.Exec(stmt, &input.ID, &input.Price, &input.Status, &input.DepartureCity, &input.DepartureTime, &input.DestinyCity, &input.DestinyTime, &input.BusID,
		time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}

	err = threatNotAffectData(rows)
	if err != nil {
		fmt.Print(err)
		return err
	}

	return nil
}