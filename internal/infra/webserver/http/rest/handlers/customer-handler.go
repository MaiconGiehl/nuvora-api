package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	di "github.com/maicongiehl/nuvora-api/configs/di"
	get_last_purchases_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-last-purchases"
	login_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/login"
	dto "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/shared/dto"
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
// @Tags         Customer
// @Accept       json
// @Produce      json
// @Param        request   								body     		dto.LoginInputDTO  true  "Login info"
// @Success      200  										{object}   	dto.CustomerAccountOutputDTO
// @Failure      404
// @Router       /customer [post]
func (h *CustomerHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input dto.LoginInputDTO

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

// Customer godoc
// @Summary      Get last purchases
// @Description  Get last purchases with customer account id
// @Tags         Customer
// @Accept       json
// @Produce      json
// @Param        id   				path     		int  true  "Id"
// @Success      200  										{object}   	object
// @Failure      404
// @Router       /customer/last-purchases/{id} [get]
func (h *CustomerHandler) LastPurchases(w http.ResponseWriter, r *http.Request) {
	customerId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := get_last_purchases_command.With(customerId)
	output, err := h.app.GetLastPurchasesUseCase.Execute(command)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(output)

}