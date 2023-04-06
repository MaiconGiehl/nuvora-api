package handlers

import (
	"encoding/json"
	"net/http"

	di "github.com/maicongiehl/nuvora-api/configs/di"
	get_last_purchases_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-last-purchases"
	get_possible_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-possible-travels"
	"github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-possible-travels/dto"
	login_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/login"
	shared_dto "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/shared/dto/login"
)

type CustomerHandler struct {
	app *di.App
}

func NewCustomerHandler(
	app *di.App,
) *CustomerHandler {
	return &CustomerHandler{
		app: app,
	}
}

// Customer godoc
// @Summary      Login as customer
// @Description  Login as customer with email and password
// @Tags         customer
// @Accept       json
// @Produce      json
// @Param        request   				body     shared_dto.LoginAsCustomerDTO  true  "Login info"
// @Success      200  										{object}   	object
// @Failure      404
// @Router       /customer [post]
func (h *CustomerHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input shared_dto.LoginAsCustomerDTO

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := login_command.With(input.Email, input.Password)
	output, err := h.app.LoginAsCustomerUseCase.Execute(command)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(output)
}

func (h *CustomerHandler) PossibleTravels(w http.ResponseWriter, r *http.Request) {
	var input dto.GetPossibleTravelsDTO
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := get_possible_command.With()
	output, err := h.app.GetPossibleTravelsUseCase.Execute(command)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(output)
}

func (h *CustomerHandler) LastPurchases(w http.ResponseWriter, r *http.Request) {
	var input dto.GetPossibleTravelsDTO
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := get_last_purchases_command.With()
	output, err := h.app.GetLastPurchasesUseCase.Execute(command)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(output)

}