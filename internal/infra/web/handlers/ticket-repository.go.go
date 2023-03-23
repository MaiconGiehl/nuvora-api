package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	dto "github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/infra/database"
	"github.com/MaiconGiehl/API/internal/usecase"
	"github.com/go-chi/chi/v5"
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

// GetAllTicket godoc
// @Summary      			Get all ticket
// @Description  			Get all ticket
// @Tags         			Ticket
// @Accept       			json
// @Produce      			json
// @Success      			200  						{object}   []dto.TicketOutputDTO
// @Failure      			404
// @Router       			/ticket [get]
func (h *TicketHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	usecase := usecase.NewGetAllTicketUseCase(*h.TicketRepository)
	output, err := usecase.Execute()
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(&output)
	if err != nil {
		returnErrMsg(w, err)
		return
	}
}

// GetTicketById godoc
// @Summary      			Search for a specific ticket
// @Description  			Get a ticket
// @Tags         			Ticket
// @Accept       			json
// @Produce      			json
// @Param        			id   										path      	int  true  "Ticket ID"
// @Success      			200  										{object}   	dto.TicketOutputDTO
// @Failure      			404
// @Router       			/ticket/{id} [get]
func (h *TicketHandler) GetTicket(w http.ResponseWriter, r *http.Request) {
	id, err := getTicketId(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	usecase := usecase.NewGetTicketUseCase(*h.TicketRepository)
	output, err := usecase.Execute(id)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(&output)
	if err != nil {
		returnErrMsg(w, err)
		return
	}
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


// UpdateTicket godoc
// @Summary      			Delete a specific ticket
// @Description  			Get a ticket
// @Tags         			Ticket
// @Accept       			json
// @Produce      			json
// @Param        			id   										path      		int  true  "Account ID"
// @Param        			request   								body      		dto.TicketInputDTO  true  "Ticket info"
// @Success      			200  										{object}   		object
// @Failure      			404
// @Router       			/ticket/{id} [patch]
func (h *TicketHandler) UpdateTicket(w http.ResponseWriter, r *http.Request) {
	id, err := getTicketId(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	ticket, err := getTicketInput(w, r)
	if err != nil {
		returnErrMsg(w, err)
		return
	}

	usecase := usecase.NewUpdateTicketUseCase(*h.TicketRepository)
	err = usecase.Execute(id, ticket)
	
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