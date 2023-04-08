package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
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
// @Summary      CreateTravel
// @Description  Create a new travel
// @Tags         Travel
// @Accept       json
// @Produce      json
// @Param        request   				body      dto.TravelInputDTO  true  "Login info"
// @Success      200  										{object}   	object
// @Failure      404
// @Router       /travel [post]
func (h *TravelHandler) CreateTravel(w http.ResponseWriter, r *http.Request) {
	var input dto.TravelOutputDTO

	fmt.Print(input)
}

// Travel godoc
// @Summary      Get customer possible travels
// @Description  Get travels using customer account id
// @Tags         Travel
// @Accept       json
// @Produce      json
// @Param        id   				path      int  true  "Login info"
// @Success      200  										{object}   	dto.TravelOutputDTO
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