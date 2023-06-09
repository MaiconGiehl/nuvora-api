package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	di "github.com/maicongiehl/nuvora-api/configs/di"
	dto "github.com/maicongiehl/nuvora-api/internal/core/application/shared/dto"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
	buy_ticket_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/buy-ticket"
	get_possible_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-possible-travels"
	get_last_purchases_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/get-purchases"
	login_command "github.com/maicongiehl/nuvora-api/internal/core/application/usecase/customer/login"
)

type CustomerHandler struct {
	logger logger.Logger
	app    *di.App
}

func NewCustomerHandler(
	logger logger.Logger,
	app *di.App,
) *CustomerHandler {
	return &CustomerHandler{
		logger: logger,
		app:    app,
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
	customerAccount, err := h.app.LoginAsCustomerUseCase.Execute(command)
	if err != nil {
		h.logger.Errorf("CustomerHandler.Login: Error at searching for customer account, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	customerAccount.SetAccessToken(h.createJWT(r,  customerAccount.PermissionLevel)) 

	h.logger.Infof("CustomerHandler.Login: New connection to account %s", input.Email)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(customerAccount)
}

func (h *CustomerHandler) createJWT(r *http.Request, permission_level int) string {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JwtExpiresIn").(int)

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"exp":              time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
		"permission_level": permission_level,
	})

	return tokenString
}

// Ticket godoc
// @Summary      Buy a ticket
// @Description  Generate a ticket when user buy one
// @Tags         Customer
// @Param        id   							path     		int  true  "Customer Account Id"
// @Param        travelId   				path     		int  true  "Travel Id"
// @Accept       json
// @Produce      json
// @Success      201  										{object}   	object
// @Failure      404
// @Failure      406
// @Router       /customer/{id}/tickets/{travelId} [post]
// @Security ApiKeyAuth
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
// @Security ApiKeyAuth
func (h *CustomerHandler) Purchases(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("CustomerHandler.Purchases: Request received")

	customerId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		h.logger.Errorf("CustomerHandler.Purchases: Error, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := get_last_purchases_command.With(customerId)
	output, err := h.app.GetPurchasesUseCase.Execute(command)

	if err != nil {
		h.logger.Errorf("CustomerHandler.Purchases: Error while buying a ticket, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	h.logger.Infof("CustomerHandler.Purchases: Purchases info delievered")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(output)

}

// Customer godoc
// @Summary      Get customer possible travels
// @Description  Get travels using customer account id
// @Tags         Customer
// @Accept       json
// @Produce      json
// @Param        id   				path      int  true  "Customer account id"
// @Success      200  										{object}   	[]dto.TravelOutputDTO
// @Failure      404
// @Router       /customer/{id}/travels [get]
// @Security ApiKeyAuth
func (h *CustomerHandler) PossibleTravels(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("CustomerHandler.PossibleTravels: Request received")

	customerId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		h.logger.Errorf("CustomerHandler.PossibleTravels: Error at getting possible travels buying a ticket, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	command := get_possible_command.With(customerId)
	output, err := h.app.GetPossibleTravelsUseCase.Execute(command)

	if err != nil {
		h.logger.Errorf("CustomerHandler.PossibleTravels: Error at getting possible travels buying a ticket, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	h.logger.Infof("CustomerHandler.PossibleTravels: PossibleTravels info delievered")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(output)
}
