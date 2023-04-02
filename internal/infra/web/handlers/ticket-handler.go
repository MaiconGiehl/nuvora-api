package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	dto "github.com/maicongiehl/techtur-api/internal/dto"
	"github.com/maicongiehl/techtur-api/internal/infra/database"
	"github.com/maicongiehl/techtur-api/internal/usecase"
)

type TicketHandler struct {
	Ctx context.Context
	TicketRepository *database.TicketRepository
}

func NewTicketHandler(ctx context.Context, ticketRepository *database.TicketRepository) *TicketHandler {
	return &TicketHandler{
		Ctx: ctx,
		TicketRepository: ticketRepository,
	}
}

// CreateTicket godoc
// @Summary      			Add ticket
// @Description  			Create new ticket
// @Tags         			Ticket
// @Accept       			json
// @Produce      			json
// @Param        			request   				body      dto.TicketInputDTO  true  "Ticket Info"
// @Success      			201  											{object}   object
// @Failure      			404
// @Router       			/ticket [post]
func (h *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	input, err := getTicketInput(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	usecase := usecase.NewCreateTicketUseCase(*h.TicketRepository) 
	err = usecase.Execute(input)
	if err != nil {
		returnErrMsg(w, err)
		returnErrMsg(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}


// DeleteTicket godoc
// @Summary      			Delete a specific ticket
// @Description  			Get a ticket
// @Tags         			Ticket
// @Accept       			json
// @Produce      			json
// @Param        			id   										path      		int  true  "Ticket Id"
// @Success      			200  										{object}   		object
// @Failure      			404
// @Router       			/ticket/{id} [delete]
func (h *TicketHandler) DeleteTicket(w http.ResponseWriter, r *http.Request) {
	id, err := getTicketId(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}
	
	usecase := usecase.NewDeleteTicketUseCase(*h.TicketRepository)
	err = usecase.Execute(id)
	if err != nil {
		returnErrMsg(w, err)
		return
	}
}


func getTicketId(w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if (err != nil || id <= 0) {
		w.WriteHeader(http.StatusBadRequest)
		return id, err
	}
	return id, nil
}

func getTicketInput(w http.ResponseWriter, r *http.Request) (*dto.TicketInputDTO, error) {
	var ticket dto.TicketInputDTO
	err := json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return &ticket, err
	}
	return &ticket, nil
}