package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	get_possible_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-possible-travels"

	di "github.com/maicongiehl/nuvora-api/configs/di"
)

type TravelHandler struct {
	app *di.App
}

func NewTravelHandler(
	app *di.App,
) *TravelHandler {
	return &TravelHandler{
		app: app,
	}
}

// Travel godoc
// @Summary      Get customer possible travels
// @Description  Get travels using customer account id
// @Tags         Travel
// @Accept       json
// @Produce      json
// @Param        id   				path      int  true  "Customer account id"
// @Success      200  										{object}   	[]dto.TravelOutputDTO
// @Failure      404
// @Router       /travel/avaiables/{id} [get]
func (h *TravelHandler) CustomerPossibleTravels(w http.ResponseWriter, r *http.Request) {
	customerId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := get_possible_command.With(customerId)
	output, err := h.app.GetPossibleTravelsUseCase.Execute(command)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(output)
}
