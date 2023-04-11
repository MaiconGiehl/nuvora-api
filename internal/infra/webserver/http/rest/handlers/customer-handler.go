package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	di "github.com/maicongiehl/nuvora-api/configs/di"
	dto "github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	buy_ticket_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/buy-ticket"
	get_last_purchases_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-purchases"
	login_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/login"
)

type CustomerHandler struct {
	logger logger.Logger
	app *di.App
}

func NewCustomerHandler(
	logger logger.Logger,
	app *di.App,
) *CustomerHandler {
	return &CustomerHandler{
		logger: logger,
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
	h.logger.Infof("CustomerHandler.Login: Request received")
	var input dto.LoginInputDTO

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		h.logger.Errorf("CustomerHandler.Login: Error at decoding request body, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := login_command.With(input.Email, input.Password)
	output, err := h.app.LoginAsCustomerUseCase.Execute(command)
	if err != nil {
		h.logger.Errorf("CustomerHandler.Login: Error at searching for customer account, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	h.logger.Infof("CustomerHandler.Login: New connection to account %s", input.Email)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(output)
}

// Ticket godoc
// @Summary      Buy a ticket
// @Description  Generate a ticket when user buy one
// @Tags         Customer
// @Param        id   							path     		int  true  "Customer Account Id"
// @Param        travelId   				path     		int  true  "Travel Id"
// @Accept       json
// @Produce      json
// @Success      200  										{object}   	object
// @Failure      404
// @Router       /customer/{id}/tickets/{travelId} [post]
func (h *CustomerHandler) BuyTicket(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("CustomerHandler.BuyTicket: Request received")

	customerId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Errorf("CustomerHandler.BuyTicket: Invalid url path, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	travelId, err := strconv.Atoi(chi.URLParam(r, "travelId"))
	if err != nil {
		h.logger.Errorf("CustomerHandler.BuyTicket: Invalid url path, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := buy_ticket_command.With(customerId, travelId)
	err = h.app.BuyTicketUseCase.Execute(command)
	if err != nil {
		h.logger.Errorf("CustomerHandler.BuyTicket: Error while buying a ticket, %s", err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	h.logger.Infof("CustomerHandler.BuyTicket: New ticket bought")
	w.WriteHeader(http.StatusCreated)
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
// @Router       /customer/{id}/tickets [get]
func (h *CustomerHandler) LastPurchases(w http.ResponseWriter, r *http.Request) {
	customerId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := get_last_purchases_command.With(customerId)
	output, err := h.app.GetPurchasesUseCase.Execute(command)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(output)

}