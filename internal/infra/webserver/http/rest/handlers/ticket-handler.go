package handlers

import (
	"encoding/json"
	"net/http"

	di "github.com/maicongiehl/nuvora-api/configs/di"
	dto "github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	buy_ticket_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/buy-ticket"
)

type TicketHandler struct {
	app *di.App
}

func NewTicketHandler(
	app *di.App,
) *TicketHandler {
	return &TicketHandler{
		app: app,
	}
}

// Ticket godoc
// @Summary      Buy a ticket
// @Description  Generate a ticket when user buy one
// @Tags         Ticket
// @Accept       json
// @Produce      json
// @Success      200  										{object}   	dto.TicketOutputDTO
// @Failure      404
// @Router       /ticket [post]
func (h *TravelHandler) BuyTicket(w http.ResponseWriter, r *http.Request) {
	var input dto.TicketInputDTO

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := buy_ticket_command.With()
	output, err := h.app.BuyTicketUseCase.Execute(command)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(output)
}